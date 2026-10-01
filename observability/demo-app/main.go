package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed, partitioned by status and path.",
		},
		[]string{"path", "status"},
	)

	errorsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total number of HTTP request errors encountered.",
		},
		[]string{"path"},
	)

	inFlightGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of in-flight HTTP requests being handled.",
		},
	)

	durationHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response latency for HTTP requests in seconds.",
			Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0},
		},
		[]string{"path", "status"},
	)
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	inFlightGauge.Inc()
	defer inFlightGauge.Dec()

	start := time.Now()
	status := http.StatusOK

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})

	elapsed := time.Since(start).Seconds()
	requestsTotal.WithLabelValues("/health", strconv.Itoa(status)).Inc()
	durationHistogram.WithLabelValues("/health", strconv.Itoa(status)).Observe(elapsed)
}

func workHandler(w http.ResponseWriter, r *http.Request) {
	inFlightGauge.Inc()
	defer inFlightGauge.Dec()

	start := time.Now()

	// Optional query parameters for testing: ?error_rate=0.2 & ?delay_ms=50
	errorRateStr := r.URL.Query().Get("error_rate")
	delayMsStr := r.URL.Query().Get("delay_ms")

	errorRate := 0.05 // default 5% error rate
	if errorRateStr != "" {
		if val, err := strconv.ParseFloat(errorRateStr, 64); err == nil {
			errorRate = val
		}
	}

	delayMs := 15 // default 15ms simulated processing
	if delayMsStr != "" {
		if val, err := strconv.Atoi(delayMsStr); err == nil {
			delayMs = val
		}
	}

	// Add random jitter to simulated latency: delayMs +/- 5ms
	jitter := rand.Intn(10) - 5
	actualDelay := delayMs + jitter
	if actualDelay < 1 {
		actualDelay = 1
	}
	time.Sleep(time.Duration(actualDelay) * time.Millisecond)

	status := http.StatusOK
	if rand.Float64() < errorRate {
		status = http.StatusInternalServerError
		errorsTotal.WithLabelValues("/work").Inc()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if status == http.StatusOK {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "ok",
			"delay_ms": actualDelay,
		})
	} else {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error":    "simulated_worker_failure",
			"delay_ms": actualDelay,
		})
	}

	elapsed := time.Since(start).Seconds()
	requestsTotal.WithLabelValues("/work", strconv.Itoa(status)).Inc()
	durationHistogram.WithLabelValues("/work", strconv.Itoa(status)).Observe(elapsed)
}

func main() {
	port := 8080
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/work", workHandler)
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Demo app listening on http://localhost:%d", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
