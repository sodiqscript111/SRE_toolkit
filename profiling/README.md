# Profiling & Runtime Diagnostics

Profiling is the dynamic measurement of resource consumption—CPU cycles, elapsed wall-clock time, memory allocations, and synchronization blocking—mapped to specific functions and lines of source code.

A critical engineering truth: **Profiling is not only CPU profiling.**

---

## Profiling Taxonomy

| Dimension | Core Question | Primary Bottleneck Indicator | Lab Status |
|---|---|---|---|
| **On-CPU** | *Which code is burning CPU cycles?* | Tight loops, quadratic algorithms, serialization | [cpu/](cpu/README.md) |
| **Flame Graphs** | *How is sampled execution time distributed?* | Wide plateau frames in call stack visualization | [flamegraphs/](flamegraphs/README.md) |
| **Async-Aware** | *Why is latency high when CPU usage is low?* | Event loop delays, async waiting, I/O latency | [async-aware/](async-aware/README.md) |
| **Heap / Allocations** | *Which code triggers GC churn & OOMs?* | Escape analysis failures, slice reallocation churn | Planned in [ROADMAP.md](../ROADMAP.md) |
| **Lock Contention** | *How long do threads/goroutines wait on mutexes?* | Critical section contention, lock hold times | Planned in [ROADMAP.md](../ROADMAP.md) |

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
    Network I/O to Downstream (Elapsed: 93ms) : 2, 95
```

- Total **CPU Time**: $2\text{ms} + 3\text{ms} = 5\text{ms}$ (CPU utilization appears near 0%).
- Total **Wall-Clock Elapsed Time**: $98\text{ms}$.
- A standard CPU profile will inspect only the $5\text{ms}$ and completely miss the $93\text{ms}$ external wait.

---

## Executable Labs

- [`cpu/`](cpu/README.md) – Go pprof on-CPU sampling and identifying hotspots under k6 load.
- [`flamegraphs/`](flamegraphs/README.md) – Reading, interpreting, and comparing differential flame graphs.
- [`async-aware/`](async-aware/README.md) – Node.js/TypeScript event loop lag, CPU-bound blocks vs async I/O waiting.
