# Reliability Lab Roadmap

This roadmap tracks the development of knowledge modules, tooling guides, and reproducible failure experiments in `reliability-lab`.

Statuses:
- ✅ **Complete**: Fully documented with runnable lab code, manifests, and reproduction scripts.
- 🚧 **In progress**: Architectural docs written; reference implementation or experiment harness being tested.
- 🧪 **Experiment planned**: Hypothesis and test topology designed; harness pending.
- 📚 **To learn**: Theoretical foundations being mapped; deeper production validation needed.

---

## 1. Observability Stack

| Topic | Status | Description |
|---|---|---|
| Prometheus Fundamentals & PromQL | ✅ Complete | Metric types, scraping, counters, gauges, histograms, rate/increase/histogram_quantile. |
| Grafana Dashboards | ✅ Complete | Provisioned datasources, Golden Signals dashboard, panel configurations. |
| Structured Logging & Correlation IDs | 🚧 In progress | Context propagation with trace/request correlation across services. |
| Distributed Tracing (OpenTelemetry) | 📚 To learn | Context propagation over W3C Trace Context, span taxonomy, trace sampling trade-offs. |
| Continuous Profiling (Pyroscope / Parca) | 📚 To learn | Always-on profiling in production, overhead considerations, differential analysis. |

---

## 2. Profiling & Performance Diagnostics

| Topic | Status | Description |
|---|---|---|
| On-CPU Profiling (Go pprof) | ✅ Complete | Sampling CPU profiles, identifying algorithmic bottlenecks, flat vs cum metrics. |
| Flame Graphs & Stack Visualization | ✅ Complete | Reading frames, frame width interpretation, identifying hot paths and false leads. |
| Wall-Clock & Async-Aware Profiling | ✅ Complete | CPU time vs elapsed time in async runtimes (Node.js/Python/Go I/O wait). |
| Memory Profiling & Allocations | 🚧 In progress | Heap in-use vs alloc_space, escape analysis, garbage collection pressure. |
| Mutex & Goroutine Blocking Contention | 🚧 In progress | Lock hold times, scheduling latency, goroutine leaks under backpressure. |

---

## 3. Performance & Load Testing

| Topic | Status | Description |
|---|---|---|
| k6 Testing Suite | ✅ Complete | Smoke, load, stress, spike, and soak test scripts with percentile thresholds. |
| Metric Interpretation | ✅ Complete | Why averages hide tail latency; p50 vs p95 vs p99; saturation vs throughput. |
| Load Testing + Profiling Closed Loop | ✅ Complete | Unified workflow linking k6 injection -> Prometheus scraping -> pprof capture. |

---

## 4. Failure Handling & Resilience Patterns

| Topic | Status | Description |
|---|---|---|
| Retries & Amplification | ✅ Complete | How multi-tier retries cause exponential traffic amplification and cascaded collapse. |
| Retry Ownership Pattern | ✅ Complete | Designing single-layer retry ownership across an Envoy + Go microservice chain. |
| Backoff & Jitter | ✅ Complete | Constant vs exponential backoff; Full Jitter vs Decorrelated Jitter against herds. |
| Circuit Breakers | ✅ Complete | State machine (Closed, Open, Half-Open), failure thresholds, cool-down timers. |
| Idempotency Patterns | 🚧 In progress | Idempotency keys, atomic deduplication, at-least-once delivery safety. |
| Dead-Letter Queues (DLQ) | 🚧 In progress | Poison-pill isolation, retry headers, message replay mechanisms. |
| Graceful Degradation | 🚧 In progress | Fallbacks, shed non-critical features, static cache fallbacks. |

---

## 5. Failure Isolation & Blast Radius

| Topic | Status | Description |
|---|---|---|
| Blast Radius & Bulkheads | 🚧 In progress | Thread pool / connection pool segregation, failure domain containment. |
| Shuffle Sharding Simulation | ✅ Complete | Mathematical simulation and routing demonstration of blast radius reduction. |
| Tenant Isolation | 🚧 In progress | Multi-tenant noisy-neighbor isolation, rate-limiting per customer tier. |

---

## 6. Kubernetes Reliability

| Topic | Status | Description |
|---|---|---|
| Health Probes | ✅ Complete | Startup vs Liveness vs Readiness; deadlocks, cascading restarts, route removal. |
| Resource Requests & Limits | ✅ Complete | CPU throttling (CFS quota) vs Memory OOMKill (cgroup kill, code 137). |
| PodDisruptionBudgets (PDB) | 🚧 In progress | Protecting quorum and minimum capacity during voluntary node draining. |
| Graceful Termination | 🚧 In progress | SIGTERM propagation, endpoint deregistration race condition, connection draining. |
| Topology Spread & Scheduling | 🚧 In progress | Zone anti-affinity, spread constraints across failure domains. |
| Node Eviction & Failure | 🧪 Experiment planned | Behavior of StatefulSets and Deployments during abrupt worker node loss. |

---

## 7. Chaos Engineering

| Topic | Status | Description |
|---|---|---|
| Chaos Engineering Methodology | ✅ Complete | Hypothesis formulation, steady-state definition, blast radius control. |
| Chaos Mesh Pod Disruption | ✅ Complete | PodKill and PodFailure experiments on Kubernetes under steady load. |
| Chaos Mesh Network Latency & Loss | ✅ Complete | Injecting millisecond delays and packet drop rates to trigger tail latency cliffs. |
| Chaos Monkey / VM Level Chaos | 🚧 In progress | Instance-level randomized termination patterns in cloud infrastructure. |
| StressChaos (CPU/Memory) | 🧪 Experiment planned | Kernel memory pressure, page thrashing, noisy neighbor chaos. |
| Advanced Chaos (IOChaos, DNSChaos) | 📚 To learn | Corrupting disk block writes, dropping DNS UDP packets to observe retries. |

---

## 8. Networking Reliability

| Topic | Status | Description |
|---|---|---|
| DNS & CoreDNS Reliability | 🚧 In progress | ndots:5 query explosion, UDP packet drops, DNS caching strategies. |
| L4 vs L7 Load Balancing | 🚧 In progress | TCP connection reuse, HTTP/2 connection coalescing, least-request routing. |
| Envoy Traffic Management | ✅ Complete | Envoy configuration for bounded timeouts, retries, and outlier detection. |
| Service Mesh Architecture | 🚧 In progress | Sidecar vs Ambient/eBPF data plane trade-offs and operational overhead. |
| Network Partitions & Asymmetric Loss | 🧪 Experiment planned | Split-brain simulation, unidirectional packet drops. |

---

## 9. High Availability & Data Systems

| Topic | Status | Description |
|---|---|---|
| Replication & Consistency | 🚧 In progress | Synchronous vs asynchronous replication, replication lag, split-brain risks. |
| Leader Election & Leases | 🚧 In progress | Consensus leases, heartbeat timeouts, fencing tokens. |
| PostgreSQL HA with Patroni | 📚 To learn | DCS-backed failover (etcd/Consul), measuring RTO/RPO during primary kill. |

---

## 10. Autoscaling & Capacity Planning

| Topic | Status | Description |
|---|---|---|
| CPU/Memory HPA | 🚧 In progress | Scaling limitations when CPU does not correlate with business workload. |
| Queue-Based Scaling (KEDA) | 🚧 In progress | Scaling on backlog depth, processing lag, consumer saturation. |
| Capacity Modeling (Little's Law) | 🚧 In progress | $L = \lambda W$, concurrency limits, queueing theory under saturation. |

---

## 11. SRE Fundamentals & SLO Engineering

| Topic | Status | Description |
|---|---|---|
| Four Golden Signals | ✅ Complete | Latency, Traffic, Errors, Saturation definitions and telemetry mappings. |
| SLI / SLO / SLA Taxonomy | ✅ Complete | Distinguishing measurement, target, and legal consequence. |
| Error Budget & Multi-Window Burn Rates | 📚 To learn | Mathematical burn-rate calculation ($14.4\times$, $6\times$, $1\times$) and pager alerting. |
| Reliability Metrics Critique | ✅ Complete | Analysis of MTTR, MTTF, MTBF fallacies and valid recovery metrics. |
