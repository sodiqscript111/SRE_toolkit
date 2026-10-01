.PHONY: help up down test lint clean load profile metrics chaos

SHELL := /bin/bash

help: ## Show this help menu
	@echo "Reliability Lab - Make commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ---------------------------------------------------------
# Core Observability Stack
# ---------------------------------------------------------
up: ## Start core observability stack (Prometheus, Grafana, target demo app)
	docker compose -f observability/docker-compose.yml up -d

down: ## Stop core observability stack
	docker compose -f observability/docker-compose.yml down -v

metrics: ## Open Prometheus web UI and scrape targets
	@echo "Prometheus is available at: http://localhost:9090"
	@echo "Grafana is available at:    http://localhost:3000 (admin/admin)"
	@echo "Sample metrics app at:      http://localhost:8080/metrics"

# ---------------------------------------------------------
# Load Testing & Performance
# ---------------------------------------------------------
load: ## Run k6 load test against sample app
	k6 run performance-testing/k6/load-test.js

load-smoke: ## Run k6 smoke test
	k6 run performance-testing/smoke-testing/smoke-test.js

load-stress: ## Run k6 stress test
	k6 run performance-testing/stress-testing/stress-test.js

load-spike: ## Run k6 spike test
	k6 run performance-testing/spike-testing/spike-test.js

# ---------------------------------------------------------
# Profiling Workflows
# ---------------------------------------------------------
profile-cpu: ## Collect 30s CPU profile from sample app via pprof
	@echo "Collecting CPU profile from http://localhost:8080/debug/pprof/profile?seconds=30..."
	go tool pprof -http=:8081 http://localhost:8080/debug/pprof/profile?seconds=30

profile-heap: ## Collect Heap profile from sample app via pprof
	@echo "Collecting Heap profile from http://localhost:8080/debug/pprof/heap..."
	go tool pprof -http=:8082 http://localhost:8080/debug/pprof/heap

# ---------------------------------------------------------
# Flagship Experiments
# ---------------------------------------------------------
experiment-retry-naive: ## Launch Experiment A: Naive multi-hop retries
	docker compose -f experiments/retry-ownership/docker-compose.naive.yml up -d --build

experiment-retry-owned: ## Launch Experiment B: Dedicated retry ownership
	docker compose -f experiments/retry-ownership/docker-compose.owned.yml up -d --build

experiment-retry-down: ## Tear down retry ownership experiment containers
	docker compose -f experiments/retry-ownership/docker-compose.naive.yml down -v
	docker compose -f experiments/retry-ownership/docker-compose.owned.yml down -v

experiment-shuffle-sim: ## Run Shuffle Sharding mathematical simulation
	go run experiments/shuffle-sharding/simulation.go

# ---------------------------------------------------------
# Chaos Experiments (Kubernetes / Local)
# ---------------------------------------------------------
chaos-pod: ## Apply Chaos Mesh Pod Failure experiment
	kubectl apply -f experiments/pod-failure/chaos-pod-kill.yaml

chaos-pod-clean: ## Remove Chaos Mesh Pod Failure experiment
	kubectl delete -f experiments/pod-failure/chaos-pod-kill.yaml --ignore-not-found

chaos-network-latency: ## Apply Chaos Mesh Network Latency experiment
	kubectl apply -f experiments/network-latency/chaos-network-delay.yaml

chaos-network-latency-clean: ## Remove Chaos Mesh Network Latency experiment
	kubectl delete -f experiments/network-latency/chaos-network-delay.yaml --ignore-not-found

# ---------------------------------------------------------
# Maintenance & Clean
# ---------------------------------------------------------
test: ## Run unit tests across all experiment subpackages
	go test -v ./...

clean: ## Clean up local temp files, logs, and profiles
	rm -rf *.pprof *.svg profile.pb.gz bin/ tmp/
