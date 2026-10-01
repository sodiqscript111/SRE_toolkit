package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	healthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestWorkHandler(t *testing.T) {
	// Request with 0% error rate and 1ms delay for deterministic fast test
	req := httptest.NewRequest(http.MethodGet, "/work?error_rate=0.0&delay_ms=1", nil)
	rr := httptest.NewRecorder()

	workHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}
