import http from 'k6/http';
import { check, sleep } from 'k6';

// Stress Test: Pushes beyond capacity to locate saturation and breaking limits
export const options = {
  stages: [
    { duration: '30s', target: 50 },  // Moderate load
    { duration: '1m', target: 150 },  // High stress
    { duration: '1m', target: 300 },  // Extreme saturation
    { duration: '30s', target: 0 },   // Cool down
  ],
  thresholds: {
    // Monitor degradation boundary without failing script early
    http_req_duration: ['p(95)<2000'],
  },
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/work';
  const res = http.get(target);

  check(res, {
    'received response': (r) => r.status !== 0,
  });

  sleep(0.05);
}
