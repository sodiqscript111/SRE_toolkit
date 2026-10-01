# Spike Testing

A spike test injects an instantaneous, extreme surge of traffic to observe how autoscaling, connection pools, and thread buffers cope with sudden shock.

---

## Spike Script: `spike-test.js`

```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 5 },   // Normal low baseline
    { duration: '10s', target: 200 }, // Instantaneous 40x spike
    { duration: '30s', target: 200 }, // Hold spike
    { duration: '10s', target: 5 },   // Instant drop
    { duration: '20s', target: 5 },   // Observe recovery
  ],
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/workload';
  const res = http.get(target);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(0.1);
}
```

## Running the Spike Test
```bash
k6 run performance-testing/spike-testing/spike-test.js
```
