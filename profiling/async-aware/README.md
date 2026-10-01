# Async-Aware Profiling: CPU Time vs Elapsed Time

In asynchronous runtimes (Node.js/TypeScript, Python `asyncio`), traditional thread-based and on-CPU profilers can be misleading.

A service can experience seconds of tail latency while CPU utilization remains below 5%. Conversely, a single synchronous CPU-heavy block will freeze the entire event loop for all concurrent requests.

---

## Core Principle

$$\text{CPU Time} \ne \text{Elapsed (Wall-Clock) Time}$$

- **On-CPU Time**: The number of clock cycles CPU cores spent executing instructions in user space or kernel space.
- **Elapsed Wall-Clock Time**: The real-world duration experienced by the client from sending the request to receiving the response.

---

## Lab Architecture

This TypeScript service ([src/server.ts](src/server.ts)) provides two contrasting endpoints:

| Endpoint | Behavior | CPU Consumption | Elapsed Time | Concurrency Impact |
|---|---|---|---|---|
| `/cpu` | Tight synchronous calculation loop (150ms) | **100% on 1 core** | 150ms | **Blocks event loop**; queues all other requests |
| `/io` | Asynchronous timer wait (150ms) | **~0%** | 150ms | **Yields event loop**; handles thousands concurrently |

---

## Running the Lab

### 1. Build and Start the Server
```bash
cd profiling/async-aware
npm install
npm run build
npm start
```
The server listens on `http://localhost:8086`.

### 2. Test the Asynchronous I/O Endpoint
In another terminal, send concurrent requests to `/io`:
```bash
curl http://localhost:8086/io
```
Response:
```json
{
  "workload": "async_io_wait",
  "elapsed_ms": 152,
  "cpu_burned": false
}
```
**Observation**: The request took 152ms of elapsed time. However, checking CPU usage with `top` or Task Manager shows CPU utilization is near 0%. The thread was idle waiting on the timer callback.

### 3. Test the CPU-Bound Synchronous Endpoint
Send a request to `/cpu`:
```bash
curl http://localhost:8086/cpu
```
Response:
```json
{
  "workload": "cpu_bound",
  "elapsed_ms": 150,
  "cpu_burned": true,
  "iterations": 1450210
}
```
**Observation**: The CPU core was 100% pegged during this 150ms window. If 10 clients hit `/cpu` at the same time, the 10th request will take $1.5\text{s}$ because each request synchronously blocks the event loop.

---

## Why Async-Aware Profiling Matters Across Runtimes

This dynamic applies universally to single-threaded event loops and cooperative coroutines:
- **Node.js / TypeScript**: Promises, async/await, libuv event loop.
- **Python `asyncio`**: Coroutine task scheduling, `await asyncio.sleep()`.

When diagnosing slow async systems:
1. If **Latency is High and CPU is High**: An on-CPU profiler (`pprof`, `py-spy`, Node `--cpu-prof`) will pinpoint synchronous loops, regexes, or serialization routines blocking the thread.
2. If **Latency is High and CPU is Low**: An on-CPU profiler will be useless. The bottleneck is off-CPU waiting (database queries, network socket reads, downstream timeouts, or event loop queue delay). You need distributed traces or wall-clock/event-loop lag metrics.
