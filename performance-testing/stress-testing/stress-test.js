import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 50 },
    { duration: '1m', target: 150 },
    { duration: '1m', target: 300 },
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
