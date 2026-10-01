# SRE Toolkit (`reliability-lab`)

> A hands-on, runnable laboratory for learning how distributed systems fail and how to make them resilient.

Reliability cannot be mastered through passive reading. This repository is built around an empirical feedback loop:

```text
Build  →  Load  →  Observe  →  Break  →  Measure  →  Diagnose  →  Improve
```

Every major module in this v0.1 release includes working code, reproducible load scripts, and clear diagnostic instructions.

---

## What Is in v0.1

| Section | Scope | Description |
|---|---|---|
| [`docs/`](docs/) | Concepts & Glossary | [Reliability Engineering Principles](docs/reliability-engineering.md), [SRE Glossary](docs/glossary.md), [Learning Path](docs/learning-path.md). |
| [`observability/`](observability/README.md) | Metrics & Dashboards | Prometheus metrics collection, custom alerts, auto-provisioned Grafana dashboards, and an instrumented Go demo app. |
| [`performance-testing/`](performance-testing/README.md) | k6 Load Suite | Executable smoke, load, stress, and spike test scripts targeting the demo service. |
| [`profiling/`](profiling/README.md) | Runtime Diagnostics | On-CPU profiling with Go `pprof` ([cpu/](profiling/cpu/README.md)), [Flame Graphs](profiling/flamegraphs/README.md), and [Async-Aware Profiling](profiling/async-aware/README.md) in TypeScript. |
| [`chaos-engineering/`](chaos-engineering/README.md) | Fault Injection | Disposable local Kubernetes lab on `kind` using [Chaos Mesh](chaos-engineering/chaos-mesh/README.md) (`PodChaos`, `NetworkChaos`) and local [Chaos Monkey](chaos-engineering/chaos-monkey/README.md) runner. |
| [`notes/`](notes/README.md) | Field Observations | Lightweight template for recording real incident observations, postmortems, and experiment takeaways. |

---

## Quickstart

### 1. Observability & Load Testing
Start the Prometheus, Grafana, and demo service stack:
```bash
make up
```
- **Prometheus UI**: `http://localhost:9090`
- **Grafana UI**: `http://localhost:3000` (User: `admin` / Password: `admin`)
- **Demo App Metrics**: `http://localhost:8080/metrics`

Generate traffic to observe the Golden Signals in Grafana:
```bash
# Run realistic peak load
make load

# Run stress test to find the saturation cliff
make stress
```

Tear down the observability stack:
```bash
make down
```

### 2. Go CPU Profiling & Flame Graphs
Run the CPU profiling target service:
```bash
cd profiling/cpu
go run app/main.go
```
In another terminal, generate load and capture a CPU profile:
```bash
# Generate traffic
cd profiling/cpu && k6 run load.js

# Capture profile and open web interface (Flame Graph)
go tool pprof -http=:8081 http://localhost:8085/debug/pprof/profile?seconds=20
```
Navigate to `http://localhost:8081/ui/flamegraph` to identify the wide plateau for `expensiveComputation()`. Follow the optimization guide in [profiling/flamegraphs/README.md](profiling/flamegraphs/README.md).

### 3. Async-Aware Profiling (Node.js/TypeScript)
Understand why async systems can suffer tail latency while CPU usage remains near 0%:
```bash
cd profiling/async-aware
npm install
npm run build
npm start
```
Test the difference between synchronous CPU-bound blocks (`curl http://localhost:8086/cpu`) and async I/O waiting (`curl http://localhost:8086/io`). See [profiling/async-aware/README.md](profiling/async-aware/README.md).

### 4. Chaos Engineering (Chaos Mesh & Chaos Monkey)
- **Chaos Mesh on Kubernetes**: Test pod termination and network latency injection in a local `kind` cluster (see [chaos-engineering/chaos-mesh/README.md](chaos-engineering/chaos-mesh/README.md)).
- **Chaos Monkey Local Runner**: Test randomized container termination against a local worker pool or in-memory simulation:
  ```bash
  # In-memory simulation
  go run chaos-engineering/chaos-monkey/main.go -simulate=true -rounds=3

  # Live container termination
  make monkey-up
  make monkey-run
  make monkey-down
  ```
  See [chaos-engineering/chaos-monkey/README.md](chaos-engineering/chaos-monkey/README.md).

---

## Makefile Command Summary

```bash
make help        # Show all available make targets
make up          # Start Prometheus, Grafana, and demo-app in Docker Compose
make down        # Stop the Docker Compose stack
make smoke       # Run k6 smoke test
make load        # Run k6 standard load test
make stress      # Run k6 stress test
make spike       # Run k6 spike test
make profile-cpu # Capture 20s CPU profile from cpu lab
make monkey-sim  # Run Chaos Monkey in simulated mode
make monkey-up   # Start Chaos Monkey local container pool
make monkey-run  # Run Chaos Monkey against local container pool
make monkey-down # Stop Chaos Monkey local container pool
make test        # Run unit tests across Go packages
make clean       # Remove build binaries and profile dumps
```

---

## Roadmap

Future topics—such as retry amplification, circuit breakers, shuffle sharding, PostgreSQL failover, and OpenTelemetry tracing—are tracked in [ROADMAP.md](ROADMAP.md).

## Contributing

Review guidelines in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE)
