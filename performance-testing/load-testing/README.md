# Load Testing

A load test evaluates how a service behaves under anticipated normal and peak production concurrency.

---

## Load Test Profile: Stages

```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 20 },  // Ramp-up to 20 VUs
    { duration: '3m', target: 20 },  // Steady-state peak load
    { duration: '30s', target: 0 },  // Ramp-down
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'],    // Error rate < 1%
    http_req_duration: ['p(95)<300', 'p(99)<600'], // p95 < 300ms, p99 < 600ms
  },
};

export default function () {
  const res = http.get('http://localhost:8080/api/v1/orders');
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(0.5);
}
```

## Running the Load Test
```bash
k6 run performance-testing/load-testing/load-test.js
```
