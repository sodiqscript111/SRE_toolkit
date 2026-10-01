package main

import (
	"testing"
)

func TestSimulatedChaosMonkey(t *testing.T) {
	engine := NewSimulatedEngine("worker-", 1.0, false, 4)

	if len(engine.GetLiveTargets()) != 4 {
		t.Fatalf("expected 4 initial instances, got %d", len(engine.GetLiveTargets()))
	}

	// 100% probability step should terminate exactly 1 instance
	engine.Step()

	remaining := len(engine.GetLiveTargets())
	if remaining != 3 {
		t.Fatalf("expected 3 remaining instances after 1 kill, got %d", remaining)
	}
}

func TestChaosMonkeyDryRun(t *testing.T) {
	engine := NewSimulatedEngine("worker-", 1.0, true, 4)

	// In dry-run mode, no instances should be terminated
	engine.Step()

	remaining := len(engine.GetLiveTargets())
	if remaining != 4 {
		t.Fatalf("expected 4 instances to remain in dry-run mode, got %d", remaining)
	}
}
