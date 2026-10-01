package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// fastHandler performs minimal O(1) in-memory response
func fastHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"endpoint": "fast",
		"status":   "ok",
	})
}

// expensiveComputation simulates a CPU bottleneck using quadratic operations
func expensiveComputation(n int) int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(10000)
	}

	// Deliberate quadratic O(n^2) bubble sort to burn CPU cycles on-core
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data[0]
}

// slowHandler invokes the CPU bottleneck function
func slowHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Sort 2,500 elements quadratically per request
	result := expensiveComputation(2500)
	elapsed := time.Since(start)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"endpoint":   "slow",
		"result":     result,
		"elapsed_ms": elapsed.Milliseconds(),
	})
}

func main() {
	port := 8085
	mux := http.NewServeMux()

	mux.HandleFunc("/fast", fastHandler)
	mux.HandleFunc("/slow", slowHandler)

	// Note: importing _ "net/http/pprof" automatically registers endpoints
	// on http.DefaultServeMux. We attach DefaultServeMux to our router:
	mux.Handle("/debug/pprof/", http.DefaultServeMux)

	log.Printf("CPU profiling target running on http://localhost:%d", port)
	log.Printf("pprof available at http://localhost:%d/debug/pprof/", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
