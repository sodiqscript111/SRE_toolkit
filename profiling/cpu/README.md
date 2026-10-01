# On-CPU Profiling with Go `pprof`

On-CPU profiling samples the call stack of active OS threads at a fixed frequency (default: 100 Hz, or every 10ms) to determine where execution time is spent.

---

## The Question Answered

> **What code is consuming CPU time right now?**

---

## Enabling `pprof` in Go

Import `net/http/pprof` for automatic HTTP handler registration:

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // registers /debug/pprof handlers on DefaultServeMux
)

func cpuHeavyWorkload(w http.ResponseWriter, r *http.Request) {
	// Inefficient computation simulating quadratic regex or nested hashing
	var total uint64
	for i := 0; i < 50_000_000; i++ {
		total += uint64(i ^ (i >> 3))
	}
	fmt.Fprintf(w, "Result: %d\n", total)
}

func main() {
	http.HandleFunc("/workload", cpuHeavyWorkload)
	http.ListenAndServe(":8080", nil)
}
```

---

## Capturing and Inspecting Profiles

### 1. Capture a 30-Second CPU Profile
```bash
go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
```

### 2. Interactive Analysis in pprof CLI
```text
(pprof) top 10
Showing nodes accounting for 4.82s, 95.82% of 5.03s total
      flat  flat%   sum%        cum   cum%
     4.50s 89.46% 89.46%      4.50s 89.46%  main.cpuHeavyWorkload
     0.20s  3.98% 93.44%      0.20s  3.98%  runtime.kevent
     0.12s  2.39% 95.82%      4.62s 91.85%  net/http.(*conn).serve
```

### Understanding `flat` vs `cum`
- **`flat`**: Time spent strictly within this specific function's instructions (excluding downstream functions it calls). High flat time indicates a hot loop or heavy internal computation.
- **`cum` (cumulative)**: Total time spent in this function **plus all downstream functions it called**. High cumulative time with low flat time means this function is an orchestrator calling a slow helper.

### 3. Launch Web UI with Visual Graphs
```bash
go tool pprof -http=:8081 profile.pb.gz
```
Navigates to `http://localhost:8081` to view interactive call trees and flame graphs.
