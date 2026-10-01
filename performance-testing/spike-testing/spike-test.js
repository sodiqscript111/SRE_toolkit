import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 5 },
    { duration: '10s', target: 200 },
    { duration: '30s', target: 200 },
    { duration: '10s', target: 5 },
    { duration: '20s', target: 5 },
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
