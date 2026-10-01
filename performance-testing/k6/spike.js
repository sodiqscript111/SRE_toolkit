import http from 'k6/http';
import { check, sleep } from 'k6';

// Spike Test: Injects sudden violent burst from 5 to 200 VUs in seconds
export const options = {
  stages: [
    { duration: '10s', target: 5 },   // Low baseline
    { duration: '10s', target: 200 }, // 40x instantaneous surge
    { duration: '30s', target: 200 }, // Hold burst
    { duration: '10s', target: 5 },   // Drop to baseline
    { duration: '15s', target: 5 },   // Observe recovery
  ],
};

export default function () {
  const target = __ENV.TARGET_URL || 'http://localhost:8080/work';
  const res = http.get(target);

  check(res, {
    'received response': (r) => r.status !== 0,
  });

  sleep(0.05);
}
