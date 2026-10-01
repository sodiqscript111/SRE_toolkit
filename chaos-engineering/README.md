# Chaos Engineering

Chaos Engineering is the discipline of experimenting on a software system to build confidence in its capability to withstand turbulent conditions in production.

It is **not** randomly breaking production; it is a hypothesis-driven empirical process.

---

## The Scientific Method of Chaos

```mermaid
flowchart TD
    A["1. Define Steady State (Normal p95 latency, 0% errors)"] --> B["2. Form Hypothesis ('If 1 pod dies, traffic routes with zero errors')"]
    B --> C["3. Inject Controlled Fault (Terminate pod or inject 150ms delay)"]
    C --> D["4. Observe Telemetry (Check error rate, latency cliff, recovery)"]
    D --> E{"5. Did Hypothesis Hold?"}
    E -->|Yes| F["Confidence validated"]
    E -->|No| G["Identify architectural weakness (Missing readiness probe, retry storm)"]
```

---

## Executable Labs

- [Chaos Mesh on Local Kubernetes](chaos-mesh/README.md) – End-to-end lab on `kind` demonstrating `PodChaos` and `NetworkChaos`.
- [Chaos Monkey Local Runner](chaos-monkey/README.md) – Netflix instance termination pattern and local Go container/simulation runner.

*(Future chaos patterns such as cross-region partitions and disk IO corruption are tracked in [ROADMAP.md](../ROADMAP.md))*.
