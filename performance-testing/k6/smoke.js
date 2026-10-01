import http from 'k6/http';
import { check, sleep } from 'k6';

// Smoke Test: Minimal 1-2 VU check to verify endpoint is online and returns 200 OK
export const options = {
  vus: 2,
  duration: '30s',
  thresholds: {
    http_req_failed: ['rate==0.00'], // 0% errors permitted
    http_req_duration: ['p(95)<100'], // 95% of requests must complete under 100ms
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
