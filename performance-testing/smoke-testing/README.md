# Smoke Testing

A smoke test is a minimal, low-concurrency execution designed to verify that the target system is functional and returns expected status codes before executing heavier load profiles.

---

## Script: `smoke-test.js`

```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 2,
  duration: '30s',
  thresholds: {
    http_req_failed: ['rate==0.00'], // 0% errors allowed
    http_req_duration: ['p(95)<100'], // 95% under 100ms
  },
};

export default function () {
  const res = http.get('http://localhost:8080/health');
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
```

## Running the Smoke Test

```bash
k6 run performance-testing/smoke-testing/smoke-test.js
```
