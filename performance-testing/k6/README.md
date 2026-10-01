# k6 Load Testing Suite

This directory contains executable [k6](https://k6.io/) load testing scenarios targeting the observability demo service (`http://localhost:8080`).

---

## Test Scenarios Overview

| Script | Profile | Virtual Users | Primary Reliability Purpose |
|---|---|---|---|
| [`smoke.js`](smoke.js) | Sanity check | 2 VUs (30s) | Verifies `/health` responds with `200 OK` before heavy load. |
| [`load.js`](load.js) | Realistic peak | 20 VUs (1m45s) | Measures $p95$/$p99$ tail latency and error rate under steady traffic. |
| [`stress.js`](stress.js) | Saturation search | Up to 300 VUs (3m) | Pushes past capacity to discover saturation cliff and queue delays. |
| [`spike.js`](spike.js) | Shock surge | 5 $\to$ 200 $\to$ 5 VUs | Injects sudden $40\times$ surge to observe thread/socket pool recovery. |

---

## Prerequisites

- [k6 installed locally](https://k6.io/docs/get-started/installation/) (`k6 version`)
- Observability demo app running:
  ```bash
  docker compose -f observability/docker-compose.yml up -d
  ```

---

## Running the Tests

### 1. Smoke Test
```bash
k6 run performance-testing/k6/smoke.js
```

### 2. Standard Load Test
```bash
k6 run performance-testing/k6/load.js
```

### 3. Stress Test
```bash
k6 run performance-testing/k6/stress.js
```

### 4. Spike Test
```bash
k6 run performance-testing/k6/spike.js
```

---

## Observing Telemetry in Real-Time

While running `load.js` or `stress.js`, open Grafana:
- **Grafana Dashboard**: `http://localhost:3000` (User: `admin` / Password: `admin`)
- Open the provisioned **Reliability Overview - Golden Signals** dashboard to watch:
  1. **Request Rate**: Matches k6 throughput.
  2. **Tail Latency**: Visualizes divergence between $p95$ and $p99$.
  3. **In-Flight Requests**: Measures concurrent queueing.
