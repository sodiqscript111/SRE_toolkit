import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 2,
  duration: '30s',
  thresholds: {
    http_req_failed: ['rate==0.00'],
    http_req_duration: ['p(95)<100'],
  },
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/health';
  const res = http.get(target);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
