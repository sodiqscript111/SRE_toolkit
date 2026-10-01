# Soak Testing

A soak test (or endurance test) runs moderate load (typically 60-80% of rated capacity) continuously for hours or days.

---

## Why Soak Tests Are Critical

Short load tests (10 minutes) cannot detect:
1. **Slow Memory Leaks**: An allocation leak of 1MB per hour will not crash a service during a 5-minute benchmark, but will cause `OOMKilled` on day 3 in production.
2. **Connection Leaks**: Leaking 1 database connection or unclosed HTTP response body every 100 requests slowly exhausts socket pools.
3. **Log Disk Saturation**: Rotated logs consuming root filesystem disk space over 48 hours.

---

## Soak Script: `soak-test.js`

```javascript
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '5m', target: 20 },
    { duration: '2h', target: 20 }, // Sustained 2-hour soak
    { duration: '5m', target: 0 },
  ],
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/api/v1/orders';
  const res = http.get(target);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
```

## Running the Soak Test
```bash
k6 run performance-testing/soak-testing/soak-test.js
```
