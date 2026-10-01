# Wall-Clock Profiling

Wall-clock profiling measures total elapsed duration from start to finish of an operation, regardless of whether the CPU was actively executing instructions or sitting idle.

---

## The Question Answered

> **Where is elapsed application time going?**

---

## Wall-Clock Time vs CPU Time

| Characteristic | On-CPU Profiler | Wall-Clock Profiler |
|---|---|---|
| **Measures** | CPU cycles consumed by thread instructions | Real-world elapsed wall time |
| **I/O Wait (Socket / Disk)** | Ignored (Thread is descheduled) | Captured (Records duration of wait) |
| **Lock Sleep / Mutex** | Ignored | Captured |
| **External API Wait** | Ignored | Captured |

---

## The Asymmetric Latency Trap

Consider an HTTP handler that spends:
- $1\text{ms}$ parsing request parameters.
- $450\text{ms}$ waiting for a slow PostgreSQL query over the network.
- $1\text{ms}$ serializing JSON.

Total wall-clock duration = $452\text{ms}$.  
Total CPU time = $2\text{ms}$.

If an engineer only looks at an On-CPU profile, the profile shows:
- 50% CPU in `parseRequest`
- 50% CPU in `jsonMarshal`

An engineer who does not understand wall-clock dynamics will spend days optimizing JSON serialization to save $0.5\text{ms}$, while the real bottleneck is the $450\text{ms}$ database network wait.

---

## How to Measure Wall-Clock Time

1. **Distributed Tracing (OpenTelemetry)**: Traces record span start and stop timestamps, capturing end-to-end wall-clock latency per call.
2. **Go Block/Goroutine Profiles**: Go `pprof/block` measures time goroutines spent blocked waiting on synchronization primitives (channels, mutexes).
3. **eBPF Off-CPU Profiling**: Tools like `bcc-tools/offcputime` record kernel stack traces when a thread is put to sleep and when it wakes up.
