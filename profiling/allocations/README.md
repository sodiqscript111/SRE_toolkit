# Allocations & Garbage Collection Churn

A high rate of temporary object allocations does not necessarily cause a memory leak, but it triggers high **Garbage Collection (GC) CPU overhead** and tail latency spikes.

---

## Why High Allocation Rate Degrades Latency

In garbage-collected runtimes (Go, Java, Node.js):
1. **CPU Stealing**: The runtime must dedicate 25% of background CPU cores to concurrent mark-and-sweep phases when the heap reaches `GOGC` target thresholds.
2. **Stop-The-World (STW) Pauses**: While modern runtimes minimize STW pauses to sub-millisecond durations, thousands of allocations per second force frequent sweep phases, inflating $p99$ tail latency.

---

## Escape Analysis in Go

Go allocates variables on the goroutine stack if their lifetime is strictly bounded to the function. If a reference escapes the function scope, it must be allocated on the heap:

```go
// Stack allocated: zero GC overhead
func computeValue() int {
    x := 42
    return x
}

// Escapes to heap: GC must track and sweep
func computePointer() *int {
    x := 42
    return &x // escapes to heap!
}
```

Check compiler escape decisions:
```bash
go build -gcflags="-m" main.go
```

---

## Allocation Hotspots Profile

Inspect functions generating the highest total allocation rate:
```bash
go tool pprof http://localhost:8080/debug/pprof/allocs
```

### Optimization Techniques
1. **`sync.Pool`**: Reusing temporary byte buffers or struct instances across requests to avoid heap allocations.
2. **Preallocate Slices**: `make([]T, 0, expectedCapacity)` avoids repeated buffer reallocation and copying.
3. **Avoid String Conversions**: Passing `[]byte` directly avoids allocating new immutable string buffers.
