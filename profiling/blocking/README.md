# Goroutine & Thread Blocking

Blocking profiles measure the time goroutines or OS threads spend waiting for synchronization events:
- Channel send/receive operations.
- `select` statements waiting on multiple channels.
- OS thread parking during syscalls.

---

## The Question Answered

> **Why are concurrent workers stalled and making zero progress?**

---

## Enabling Block Profiling in Go

Block profiling is disabled by default due to runtime overhead. Enable it selectively with a sampling rate:

```go
package main

import (
	"runtime"
)

func init() {
	// Sample 1 out of every 1000 blocking events (nanoseconds duration)
	runtime.SetBlockProfileRate(10_000)
}
```

---

## Common Goroutine Blocking Hazards

### 1. Unbuffered Channel Deadlocks & Stalls
```go
// Sender blocks until a receiver is ready
ch := make(chan WorkItem) // unbuffered
ch <- item // BLOCKS if consumer is slow
```

### 2. Leaked Goroutines (Channel Wait Forever)
```go
func processRequest() {
    ch := make(chan Result)
    go func() {
        res := slowCall()
        ch <- res // BLOCKS FOREVER if caller timed out and abandoned receiver
    }()
    
    select {
    case r := <-ch:
        // success
    case <-time.After(100 * time.Millisecond):
        // caller returns, goroutine above leaks and holds memory
    }
}
```

---

## Inspecting Block Profiles

```bash
go tool pprof http://localhost:8080/debug/pprof/block
```
The resulting call graph highlights the channels and synchronization primitives with the highest total wait time.
