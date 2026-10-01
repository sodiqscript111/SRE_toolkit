# k6 Reusable Harness & Prometheus Export

k6 is a developer-centric, scriptable load testing tool written in Go that executes tests scripted in JavaScript (ES6).

---

## Outputting Metrics Directly to Prometheus

Run k6 with native Prometheus remote write output to visualize k6 metrics directly alongside infrastructure metrics in Grafana:

```bash
K6_PROMETHEUS_REMOTE_URL=http://localhost:9090/api/v1/write \
k6 run -o experimental-prometheus-rw performance-testing/k6/load-test.js
```

---

## Defining Strict Percentile Thresholds

In k6, thresholds halt the test or mark CI/CD runs as failed if performance drops below SLO targets:

```javascript
export const options = {
  thresholds: {
    // Failure rate must remain strictly below 0.1%
    http_req_failed: ['rate<0.001'],
    // 90% of requests must finish within 150ms, 99% within 400ms
    http_req_duration: ['p(90)<150', 'p(99)<400'],
  },
};
```

---

## Test Execution Command

```bash
k6 run performance-testing/k6/load-test.js
```
