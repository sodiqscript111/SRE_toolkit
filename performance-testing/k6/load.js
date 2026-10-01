import http from 'k6/http';
import { check, sleep } from 'k6';

// Load Test: Simulates sustained realistic peak traffic against the demo service
export const options = {
  stages: [
    { duration: '30s', target: 20 }, // Ramp up to 20 concurrent Virtual Users
    { duration: '1m', target: 20 },  // Maintain peak load
    { duration: '15s', target: 0 },  // Ramp down
  ],
  thresholds: {
    // Allow up to 10% errors due to simulated 5% server fault injection
    http_req_failed: ['rate<0.10'],
    // 95% of responses should complete under 250ms
    http_req_duration: ['p(95)<250', 'p(99)<500'],
  },
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/work';
  const res = http.get(target);

  check(res, {
    'status is 200 or 500': (r) => r.status === 200 || r.status === 500,
  });

  sleep(0.1);
}
