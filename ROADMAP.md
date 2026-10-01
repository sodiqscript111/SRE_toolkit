# Reliability Lab Roadmap

This roadmap tracks the development of knowledge modules, tooling guides, and reproducible failure experiments in `SRE_toolkit`.

Status Definitions:
- ✅ **Complete**: Working code and configuration exist, the example is locally runnable, commands are documented, and the implementation has been verified.
- 🚧 **In progress**: Reference implementation or harness is currently being scaffolded or validated.
- 🧪 **Planned**: Target topology and hypothesis designed; implementation scheduled for future release.
- 📚 **Notes only**: Theoretical explanation and architecture notes exist; no executable code yet.

---

## v0.1 Implemented & Verified Modules

| Module | Status | Description | Location |
|---|---|---|---|
| **Prometheus Telemetry** | ✅ Complete | Metric types (Counter, Gauge, Histogram), scrape configs, and alerting rules. | [observability/prometheus/](observability/prometheus/README.md) |
| **Grafana Dashboards** | ✅ Complete | Auto-provisioned datasources and Golden Signals dashboard (`reliability-overview.json`). | [observability/grafana/](observability/grafana/README.md) |
| **Observability Demo Service** | ✅ Complete | Go HTTP service exposing `/health`, `/work`, and `/metrics` with configurable latency/error injection. | [observability/demo-app/](observability/demo-app/main.go) |
| **k6 Testing Suite** | ✅ Complete | Executable smoke, load, stress, and spike test scripts targeting the demo service. | [performance-testing/k6/](performance-testing/k6/README.md) |
| **On-CPU Profiling (Go pprof)** | ✅ Complete | Target app with `/fast` and `/slow` CPU-bound routes; load script, CLI and Web UI pprof workflows. | [profiling/cpu/](profiling/cpu/README.md) |
| **Flame Graphs** | ✅ Complete | Practical diagnostic guide building on CPU lab: sampling, reading frames, optimizing code, and differential flame graphs. | [profiling/flamegraphs/](profiling/flamegraphs/README.md) |
| **Async-Aware Profiling** | ✅ Complete | TypeScript/Node.js service illustrating CPU-bound event loop blocking vs async I/O waiting. | [profiling/async-aware/](profiling/async-aware/README.md) |
| **Chaos Mesh on Kubernetes** | ✅ Complete | Disposable `kind` lab with target deployment, `PodChaos`, and `NetworkChaos` manifests. | [chaos-engineering/chaos-mesh/](chaos-engineering/chaos-mesh/README.md) |
| **Reliability Concepts & Glossary** | 📚 Notes only | Reference documentation covering failure dynamics, mathematical definitions, and learning sequence. | [docs/](docs/reliability-engineering.md) |

---

## Future Modules (Planned Work)

These topics are planned for future versions. Directories will only be created when runnable experiments are implemented:

### Resilience & Failure Handling
- 🧪 **Retries & Retry Amplification**: Multi-tier retry storm demonstrations and mathematical analysis.
- 🧪 **Retry Ownership Pattern**: Single-tier retry assignment across service chains.
- 🧪 **Timeouts, Exponential Backoff & Jitter**: Full jitter vs decorrelated jitter against synchronized thundering herds.
- 🧪 **Circuit Breakers**: State machine transitions (`Closed` $\to$ `Open` $\to$ `Half-Open`) under error bursts.
- 🧪 **Idempotency Keys**: Safe retry mechanisms with atomic request deduplication.
- 🧪 **Dead-Letter Queues (DLQ)**: Poison pill handling and retry queues.
- 🧪 **Graceful Degradation**: Fallback caches and non-critical feature shedding.

### Failure Isolation & Blast Radius
- 🧪 **Bulkheads**: Thread pool and socket pool resource partitioning.
- 🧪 **Shuffle Sharding**: Combinatorial tenant routing to reduce outage blast radius from 100% to < 5%.
- 🧪 **Tenant Isolation**: Noisy-neighbor mitigations and per-tier concurrency limits.

### Kubernetes Reliability
- 🧪 **Probes**: In-depth analysis of startup vs liveness vs readiness probes and cascading restart hazards.
- 🧪 **Resource Requests & Limits**: Linux CFS CPU throttling vs cgroup `OOMKilled` (Exit code 137).
- 🧪 **PodDisruptionBudgets (PDB)**: Protecting capacity during voluntary node draining.
- 🧪 **Graceful Termination**: Handling SIGTERM, in-flight request draining, and endpoint propagation races.
- 🧪 **Autoscaling & KEDA**: Scaling on queue backlog and worker lag instead of raw CPU.

### Networking & Service Mesh
- 🧪 **CoreDNS Reliability**: `ndots:5` search path amplification and DNS caching strategies.
- 🧪 **Envoy Traffic Management**: Outlier detection, active health checks, and bounded retries.
- 🧪 **Service Mesh & eBPF**: Sidecar vs Cilium ambient data planes and operational trade-offs.

### High Availability & Storage
- 🧪 **Replication & Lag**: Synchronous vs asynchronous replication trade-offs and split-brain risks.
- 🧪 **Leader Election & Leases**: Distributed consensus, lease expiration, and fencing tokens.
- 🧪 **PostgreSQL HA**: Patroni and etcd-backed leader failover measuring write downtime (RTO/RPO).

### SRE Fundamentals & Advanced Observability
- 🧪 **SLIs, SLOs & Error Budgets**: Mathematical calculation of availability targets and multi-window multi-burn-rate alerting.
- 🧪 **Distributed Tracing (OpenTelemetry)**: W3C Trace Context propagation and span hierarchies.
- 🧪 **Continuous Profiling**: Always-on profiling in production using Grafana Pyroscope or Parca.
