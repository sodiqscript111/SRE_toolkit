# Reliability Engineering Glossary

Concise definitions of critical concepts in SRE, distributed systems, and production engineering.

---

### Availability
The proportion of time a system is functional and capable of fulfilling requests correctly, commonly expressed as a percentage (e.g., $99.9\%$ or "three nines"):
$$\text{Availability} = \frac{\text{Successful Requests}}{\text{Total Valid Requests}}$$

### Reliability
The probability that a system performs its intended function without failure over a specified period under defined environmental conditions. Unlike raw availability, reliability accounts for correct output, timeliness, and consistency.

### Resilience
The ability of a system to withstand, adapt to, and recover gracefully from adverse events, component degradation, or unexpected load spikes without complete failure.

### Fault Tolerance
The architectural property enabling a system to continue normal operation without interruption despite the failure of one or more internal components (e.g., via redundancy, consensus, or error masking).

### Redundancy
Duplication of critical components or data paths to prevent a single point of failure (SPOF). Redundancy is necessary for high availability, but insufficient on its own (e.g., redundant nodes can fail simultaneously due to software bugs or misconfigurations).

### Service Level Indicator (SLI)
A quantifiable metric that measures service performance in real time (e.g., HTTP request latency $\le 200\text{ms}$ or percentage of successful HTTP responses).

### Service Level Objective (SLO)
An agreed internal target reliability level for an SLI over a defined time window (e.g., "99.9% of HTTP requests return status $< 500$ within 250ms over a rolling 30-day window").

### Service Level Agreement (SLA)
A formal, legally binding commitment made to external customers regarding service availability, typically including contractual or financial penalties if the commitment is breached.

### Error Budget
The allowed amount of unreliability over a time window, derived directly from the SLO:
$$\text{Error Budget} = 100\% - \text{SLO}$$
If an SLO is $99.9\%$, the error budget is $0.1\%$.

### Mean Time to Recovery (MTTR)
The average elapsed time required to restore a failed component or service to normal operational status after an incident is detected.

### Recovery Time Objective (RTO)
The maximum acceptable duration of time a system or business process can remain unavailable following a disaster before catastrophic damage occurs.

### Recovery Point Objective (RPO)
The maximum acceptable amount of data loss measured in time (e.g., 5 minutes of transaction logs) that a business can tolerate following a disruption.

### Percentile Latency (p50, p95, p99, p99.9)
A statistical boundary indicating that $X\%$ of requests were served faster than that value. For example, a $p99$ of $120\text{ms}$ means $99\%$ of requests completed in $120\text{ms}$ or less, and $1\%$ took longer. Percentiles expose tail latency that arithmetic means hide.

### Saturation
The degree to which a constrained system resource (CPU, memory, disk I/O bandwidth, network buffers, connection pools) is occupied relative to its total capacity. Once saturation approaches $100\%$, latency climbs exponentially due to queueing delays.

### Retry
Re-sending a failed request across a network in anticipation that the failure was transient (e.g., temporary packet loss or brief load spike).

### Exponential Backoff
A delay strategy where the wait time between successive retries increases exponentially (e.g., $100\text{ms}, 200\text{ms}, 400\text{ms}, 800\text{ms}$) to allow an overloaded dependency time to recover:
$$t_{\text{wait}} = t_{\text{base}} \times 2^{\text{attempt}}$$

### Jitter
Random noise added to backoff intervals to prevent multiple concurrent clients from retrying at identical synchronized moments, avoiding thundering herd retry storms.

### Circuit Breaker
A proxy pattern that detects repeated failures and trips into an `Open` state, failing fast immediately without executing remote calls. After a cool-down period, it enters `Half-Open` to probe whether the downstream dependency has recovered.

### Bulkhead
An isolation pattern that partitions resources (threads, sockets, CPU quotas, memory) into distinct pools so that a failure or exhaustion in one pool cannot starve or crash other unrelated workloads.

### Blast Radius
The total scope, number of tenants, or percentage of traffic impacted when a specific component, service, or availability zone fails.

### Shuffle Sharding
An advanced routing and assignment technique that deterministically assigns each customer or tenant to a unique, small subset (shard) of total capacity. While individual workers are shared, the combination assigned to each tenant is unique, drastically reducing the probability of two tenants sharing the exact same failing shard.

### Idempotency
A mathematical and API property where executing an operation multiple times produces the exact same outcome as executing it once (e.g., HTTP `PUT` or `POST` with a deduplicating `Idempotency-Key` header).

### Failover
The automatic or manual process of redirecting workload and routing traffic from a failed primary node or region to a designated standby replica.

### Split-Brain
A critical distributed systems failure state where network partitioning isolates nodes into two or more subsets, each incorrectly believing the other is dead and electing their own primary leader, leading to concurrent writes and silent data corruption.

### Profiling
The measurement and dynamic analysis of a program's resource consumption (CPU instructions, wall-clock time, memory allocations, thread blocking, lock contention) mapped back to specific functions and lines of code.

### Tracing
The telemetry technique of tracking the lifecycle of an end-to-end request as it flows through multiple microservices, recording causal parent-child relationships and span timings using context propagation.

### High Cardinality
A metric dimension characterized by a very large or unbounded number of unique values (e.g., user IDs, order IDs, timestamps). In time-series databases like Prometheus, high cardinality explodes memory consumption and query times.
