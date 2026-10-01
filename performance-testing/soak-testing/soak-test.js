import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '2m', target: 20 },
    { duration: '30m', target: 20 },
    { duration: '2m', target: 0 },
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
