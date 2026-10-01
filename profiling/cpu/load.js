import http from 'k6/http';
import { check, sleep } from 'k6';

// Generates load across /fast and /slow endpoints to profile CPU bottlenecks
export const options = {
  stages: [
    { duration: '10s', target: 5 },
    { duration: '40s', target: 15 },
    { duration: '10s', target: 0 },
  ],
};

export default function () {
  // 80% traffic hits /slow, 20% hits /fast
  if (Math.random() < 0.8) {
    const resSlow = http.get('http://localhost:8085/slow');
    check(resSlow, { 'status is 200': (r) => r.status === 200 });
  } else {
    const resFast = http.get('http://localhost:8085/fast');
    check(resFast, { 'status is 200': (r) => r.status === 200 });
  }

  sleep(0.05);
}
