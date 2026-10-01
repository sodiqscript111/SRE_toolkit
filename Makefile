.PHONY: help up down smoke load stress spike profile-cpu kind-up kind-down chaos-app chaos-pod chaos-network chaos-clean test clean

help: ## Show this help menu
	@echo "SRE Toolkit / Reliability Lab Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

# Observability Stack
up: ## Start Prometheus, Grafana, and demo-app in Docker Compose
	docker compose -f observability/docker-compose.yml up -d

down: ## Tear down the Docker Compose observability stack
	docker compose -f observability/docker-compose.yml down -v

# Performance Testing with k6
smoke: ## Run k6 smoke test (2 VUs against /health)
	k6 run performance-testing/k6/smoke.js

load: ## Run k6 standard load test (20 VUs against /work)
	k6 run performance-testing/k6/load.js

stress: ## Run k6 stress test (up to 300 VUs against /work)
	k6 run performance-testing/k6/stress.js

spike: ## Run k6 spike test (instantaneous 200 VU surge)
	k6 run performance-testing/k6/spike.js

# Profiling
profile-cpu: ## Capture 20s CPU profile from cpu profiling target (:8085)
	go tool pprof http://localhost:8085/debug/pprof/profile?seconds=20

# Kubernetes Chaos Mesh (Local kind cluster)
kind-up: ## Create local kind cluster for chaos lab
	kind create cluster --name chaos-lab

kind-down: ## Delete local kind cluster
	kind delete cluster --name chaos-lab

chaos-app: ## Deploy chaos target application to Kubernetes
	kubectl apply -f chaos-engineering/chaos-mesh/k8s/

chaos-pod: ## Inject PodChaos (periodic pod kill)
	kubectl apply -f chaos-engineering/chaos-mesh/experiments/pod-kill.yaml

chaos-network: ## Inject NetworkChaos (150ms network delay)
	kubectl apply -f chaos-engineering/chaos-mesh/experiments/network-delay.yaml

chaos-clean: ## Remove chaos experiments and target application
	kubectl delete -f chaos-engineering/chaos-mesh/experiments/pod-kill.yaml --ignore-not-found
	kubectl delete -f chaos-engineering/chaos-mesh/experiments/network-delay.yaml --ignore-not-found
	kubectl delete -f chaos-engineering/chaos-mesh/k8s/ --ignore-not-found

# Testing & Hygiene
test: ## Run tests across Go packages
	cd observability/demo-app && go test -v ./...
	cd profiling/cpu/app && go test -v ./...
	cd chaos-engineering/chaos-mesh/app && go test -v ./...

clean: ## Clean up temporary profile dumps and build files
	rm -rf *.pprof *.pb.gz bin/ dist/
