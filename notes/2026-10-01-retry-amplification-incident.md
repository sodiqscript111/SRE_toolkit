# Retry Amplification and Cascading Collapse

## What I learned
When multiple layers in a service call chain independently configure retries, the downstream load multiplies exponentially rather than linearly. In a chain with $N$ hops where each tier retries $R$ times, the maximum possible requests reaching the bottom tier for a single user request is:
$$\text{Max Downstream Requests} = (1 + R)^N$$

For a 4-tier chain (`Gateway -> Checkout -> Inventory -> Stock`) with $R=2$ (initial request + 2 retries = 3 attempts total per hop), a single failed request produces up to $3^3 = 27$ requests at the `Stock` service.

## What surprised me
Adding retries to "improve reliability" was the direct cause of the outage. The database did not fail on its own; it had a brief 200ms lock delay. Upstream timeouts fired, triggering retries simultaneously across all tiers. The database was hit with $20\times$ baseline throughput, exhausting its connection pool and causing complete failure.

## Mental model
Think of retries like megaphone echoes in an enclosed canyon: a single shout reflects back and forth until the sound pressure shatters the windows.

## Example
Envoy naive retry policy applied at every hop:
```yaml
retry_policy:
  retry_on: "5xx,connect-failure,refused-stream"
  num_retries: 2
```
If each hop retries twice, 1 root failure results in:
- Gateway sends 3 requests to Checkout.
- Each Checkout request sends 3 requests to Inventory (total 9).
- Each Inventory request sends 3 requests to Stock (total 27).

## Mistakes / misconceptions I had
I assumed "retrying 2 times" was a safe, conservative default for HTTP clients. I failed to consider the compound multiplication across distributed microservice boundaries.

## Things I still don't understand
How Envoy's adaptive concurrency and retry budgets interact in practice when downstream errors are a mix of HTTP 503s (overload) and HTTP 504s (gateway timeout).

## Experiment idea
Build a 4-tier Go + Envoy service chain in Docker Compose. Inject 10% artificial errors at the bottom leaf service, and measure total request count at the leaf with naive retries vs single-tier retry ownership.
*(Implemented in [`experiments/retry-ownership/`](file:///experiments/retry-ownership/README.md))*.
