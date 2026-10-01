# Contributing to Reliability Lab

We welcome contributions that improve explanations, add reproducible failure experiments, optimize profiling setups, or introduce implementations in other languages.

The core principle of this project is:
> **Learn reliability engineering by deliberately breaking, observing, measuring, and improving systems.**

Contributions must emphasize understanding, reproducibility, telemetry, and actionable reliability lessons—not tool collection for its own sake.

---

## Areas Where You Can Contribute

1. **New Experiments**: Reproducible failure cases demonstrating distributed systems failures, memory pressure, connection pool exhaustion, or cascading timeouts.
2. **Clarifications & Corrections**: Improving technical accuracy, fixing misconceptions, or refining explanations of subtle runtime behaviors (e.g., Go scheduler quirks, Linux CFS quota nuances, TCP retransmission semantics).
3. **Dashboards & PromQL**: Refining Grafana panels, alerting rules, and PromQL queries to surface degradation more clearly.
4. **Multi-Language Examples**: Adding equivalent profiling or resilience pattern implementations in Go, TypeScript, Python, Rust, or Java.
5. **Real-World Incident Case Studies**: Adding learning notes dissecting production incidents and how they map to lab experiments.

---

## Experiment Quality Standards

Every experiment added to `experiments/` must follow the standard repository template. An experiment is not complete if it only runs a command; it must explain what was broken, what telemetry showed, why it failed, and what engineering lesson applies.

### Required Experiment Structure

Every experiment directory (`experiments/<name>/README.md`) must adhere to this exact structure:

```markdown
# Experiment Name

## Goal
What are we trying to understand?

## System
What does the architecture look like? (Include Mermaid diagram)

## Hypothesis
What do we expect to happen when the fault is introduced?

## Setup
How do we start the target workload and dependencies?

## Baseline
What does the healthy system look like under normal load?

## Failure Injection
What fault are we injecting (latency, kill, packet loss, thread leak)?

## Observability
What metrics, logs, traces, or profiles should be inspected?

## Run
Exact shell commands required to execute the test and inject the fault.

## Results
What happened during the test?
(If unmeasured in your PR environment, specify: `TODO: Run experiment and record measurements.`)
Do NOT invent benchmark numbers or fabricated charts.

## Explanation
Why did the system behave this way? What runtime or network mechanism was triggered?

## Reliability Lesson
What architectural or operational rule should an engineer take away from this?

## Cleanup
Commands to tear down containers, pods, or background processes.

## Further Experiments
Suggestions for modifying variables (e.g., varying jitter, buffer sizes, timeout budgets).
```

---

## Code and Manifest Standards

1. **Local Execution**: All experiments must run locally using standard tooling (Docker Compose, `kind`, `kubectl`, Go, Node.js, Python, or `k6`). Avoid dependencies on paid cloud services.
2. **Deterministic Configuration**: Ports, timeouts, retry counts, and seed values should be explicitly defined rather than left to implicit runtime defaults.
3. **Observability Integration**: Applications should expose standard `/metrics` (Prometheus) endpoints or pprof endpoints (`net/http/pprof`) where relevant.
4. **No Artificial Bloat**: Keep sample microservices minimal. The code should demonstrate the failure mode with the minimum lines of code required.
5. **Preserve Comments and Docstrings**: Keep comments that explain why non-obvious configurations or algorithms were chosen.

---

## Workflow

1. Fork the repository and create a feature branch (`git checkout -b feat/new-failure-scenario`).
2. Implement your experiment or documentation update.
3. Test commands locally using the root `Makefile`.
4. Ensure all links and markdown syntax are valid.
5. Submit a Pull Request describing:
   - The failure scenario tested.
   - The telemetry used to observe the failure.
   - The reliability lesson demonstrated.
