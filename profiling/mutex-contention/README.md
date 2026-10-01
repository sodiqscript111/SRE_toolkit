# Mutex & Lock Contention

Mutex contention occurs when multiple concurrent execution threads or goroutines attempt to acquire a mutual exclusion lock (`sync.Mutex` or `sync.RWMutex`) held by another worker.

---

## The Question Answered

> **How much wall-clock time are requests losing waiting in queue for shared locks?**

---

## Enabling Mutex Profiling in Go

Mutex profiling is controlled via `runtime.SetMutexProfileFraction(rate)`. A fraction of `5` samples approximately $1/5$ of all contested mutex events:

```go
package main

import (
	"runtime"
	"sync"
)

var (
	mu    sync.Mutex
	state int
)

func init() {
	runtime.SetMutexProfileFraction(5)
}

func updateSharedState() {
	mu.Lock()
	defer mu.Unlock()
	// Contention grows as concurrency increases
	state++
}
```

---

## Capturing and Reading Mutex Profiles

```bash
go tool pprof http://localhost:8080/debug/pprof/mutex
```

Sample output:
```text
(pprof) top
Showing nodes accounting for 14.20s, 98.61% of 14.40s total
      flat  flat%   sum%        cum   cum%
    14.20s 98.61% 98.61%     14.20s 98.61%  sync.(*Mutex).Lock
         0     0% 98.61%     14.20s 98.61%  main.updateSharedState
```

The output reveals that callers spent an aggregate of 14.2 seconds descheduled, waiting in lock contention queues.

---

## Mitigation Strategies

1. **Reduce Critical Section Scope**: Perform I/O, parsing, and serialization outside the lock; hold the lock only for minimal in-memory state updates.
2. **Lock Sharding (Partitioning)**: Partition a single global mutex into $N$ buckets (e.g., hash mod 64) so concurrent requests operate on independent locks.
3. **Atomic Operations**: Replace simple counter or flag locks with `sync/atomic` primitives (`atomic.AddInt64`).
4. **Channel Actor Pattern**: Replace shared memory with a single worker goroutine processing incoming messages over a buffered channel.
