package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFastHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/fast", nil)
	rr := httptest.NewRecorder()

	fastHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
}

func TestExpensiveComputation(t *testing.T) {
	// Quick test on a small slice size
	res := expensiveComputation(10)
	if res < 0 || res > 10000 {
		t.Fatalf("unexpected sorted result %d", res)
	}
}
