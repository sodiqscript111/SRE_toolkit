# Stress Testing

A stress test pushes the system past its designed capacity until it breaks. The goal is to identify:
1. The exact concurrency or throughput where saturation occurs.
2. How the system degrades (graceful backpressure/HTTP 429/503 vs process crash/OOMKill).
3. Whether the system recovers automatically when load subsides.

---

## Stress Script: `stress-test.js`

```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 50 },
    { duration: '1m', target: 150 },
    { duration: '1m', target: 300 }, // Overload threshold
    { duration: '30s', target: 0 },
  ],
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/workload';
  const res = http.get(target);
  check(res, {
    'status is 200 or 429': (r) => r.status === 200 || r.status === 429,
  });
  sleep(0.1);
}
```

## Running the Stress Test
```bash
k6 run performance-testing/stress-testing/stress-test.js
```
