# Profiling & Runtime Diagnostics

Profiling is the dynamic measurement of resource consumption—CPU cycles, wall-clock duration, memory allocations, and synchronization blocking—mapped to specific functions and source lines.

A critical engineering truth: **Profiling is not just CPU profiling.**

---

## Profiling Taxonomy

| Profiling Dimension | Core Question | Primary Bottleneck Indicator | Tools |
|---|---|---|---|
| **On-CPU** | *Which code is burning CPU cycles?* | Tight loops, regex compilation, JSON serialization | Go `pprof/profile`, Linux `perf`, `py-spy` |
| **Wall-Clock** | *Where is elapsed request time going?* | I/O wait, network latency, slow database queries | Tracing spans, wall-clock profilers |
| **Async-Aware** | *Why is the event loop blocked or delayed?* | Unhandled microtasks, queue delays, starvation | Node Clinic.js, Python `asyncio` debug mode |
| **Heap Memory** | *Which objects remain in memory?* | Unbounded caches, retained references, leaks | Go `pprof/heap`, Node heapdump, V8 snapshot |
| **Allocations** | *Which code triggers garbage collection churn?* | Temporary slice reallocation, string conversions | Go `pprof/allocs`, `-benchmem` |
| **Mutex Contention** | *How long do threads/goroutines wait for locks?* | Shared lock contention under concurrency | Go `pprof/mutex` |
| **Goroutine / Block** | *Why are goroutines stalled?* | Unbuffered channel full, socket read blocking | Go `pprof/block`, `pprof/goroutine` |

---

## CPU Time vs Elapsed Time

In asynchronous and distributed services:
$$\text{CPU Time} \ne \text{Elapsed (Wall-Clock) Time}$$

```mermaid
gantt
    title CPU Time vs Elapsed Time Breakdown
    dateFormat X
    axisFormat %s ms
    section On-CPU Execution
    Parse Request (CPU: 2ms) : 0, 2
    Serialize Response (CPU: 3ms) : 95, 98
    section Wall-Clock Wait (Zero CPU)
    Network I/O to Payment Gateway (Elapsed: 93ms) : 2, 95
```

- Total **CPU Time**: $2\text{ms} + 3\text{ms} = 5\text{ms}$ (CPU utilization appears near 0%).
- Total **Wall-Clock Elapsed Time**: $98\text{ms}$.
- A standard CPU profile will only inspect the $5\text{ms}$ and completely miss the $93\text{ms}$ network delay.

---

## Directory Index

- [`cpu/`](file:///profiling/cpu/README.md) – Go pprof on-CPU sampling and identifying hotspots.
- [`wall-clock/`](file:///profiling/wall-clock/README.md) – Tracing elapsed time vs CPU time.
- [`async-aware/`](file:///profiling/async-aware/README.md) – Node.js and Python event loop starvation and async wait bottlenecks.
- [`memory/`](file:///profiling/memory/README.md) – Heap growth, in-use objects, and leaks.
- [`allocations/`](file:///profiling/allocations/README.md) – High-churn temporary allocations and GC pressure.
- [`blocking/`](file:///profiling/blocking/README.md) – Goroutine channel stalls and thread descheduling.
- [`mutex-contention/`](file:///profiling/mutex-contention/README.md) – Lock contention analysis.
- [`flamegraphs/`](file:///profiling/flamegraphs/README.md) – Reading, interpreting, and comparing flame graphs.
