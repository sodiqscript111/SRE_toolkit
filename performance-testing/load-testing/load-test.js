import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 },
    { duration: '1m', target: 20 },
    { duration: '15s', target: 0 },
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'],
    http_req_duration: ['p(95)<300', 'p(99)<600'],
  },
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/api/v1/orders';
  const res = http.get(target);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(0.5);
}
