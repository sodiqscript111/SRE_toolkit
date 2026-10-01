package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

// Instance represents a monitored service node
type Instance struct {
	ID    string
	Name  string
	Alive bool
}

// MonkeyEngine manages randomized instance termination
type MonkeyEngine struct {
	TargetPrefix string
	Probability  float64
	DryRun       bool
	Simulated    bool
	Instances    []*Instance
}

func NewSimulatedEngine(targetPrefix string, probability float64, dryRun bool, count int) *MonkeyEngine {
	instances := make([]*Instance, count)
	for i := 0; i < count; i++ {
		instances[i] = &Instance{
			ID:    fmt.Sprintf("sim-inst-%02d", i+1),
			Name:  fmt.Sprintf("%s%d", targetPrefix, i+1),
			Alive: true,
		}
	}
	return &MonkeyEngine{
		TargetPrefix: targetPrefix,
		Probability:  probability,
		DryRun:       dryRun,
		Simulated:    true,
		Instances:    instances,
	}
}

// GetLiveTargets returns all alive candidate instances
func (m *MonkeyEngine) GetLiveTargets() []*Instance {
	if m.Simulated {
		var live []*Instance
		for _, inst := range m.Instances {
			if inst.Alive {
				live = append(live, inst)
			}
		}
		return live
	}

	// In Docker mode, query running containers matching target prefix
	cmd := exec.Command("docker", "ps", "--filter", fmt.Sprintf("name=%s", m.TargetPrefix), "--format", "{{.ID}}|{{.Names}}")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var live []*Instance
	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			live = append(live, &Instance{
				ID:    strings.TrimSpace(parts[0]),
				Name:  strings.TrimSpace(parts[1]),
				Alive: true,
			})
		}
	}
	return live
}

// TerminateInstance picks and kills one target instance
func (m *MonkeyEngine) TerminateInstance(target *Instance) error {
	if m.DryRun {
		log.Printf("[DRY-RUN] Would terminate instance: %s (%s)", target.Name, target.ID)
		return nil
	}

	if m.Simulated {
		target.Alive = false
		log.Printf("[KILL] Terminated simulated instance: %s (%s)", target.Name, target.ID)
		return nil
	}

	// Real Docker termination
	log.Printf("[KILL] Sending SIGKILL to container: %s (%s)", target.Name, target.ID)
	cmd := exec.Command("docker", "kill", target.ID)
	return cmd.Run()
}

// Step runs a single termination evaluation window
func (m *MonkeyEngine) Step() {
	live := m.GetLiveTargets()
	log.Printf("[STATUS] Active live instances in pool: %d", len(live))

	if len(live) == 0 {
		log.Printf("[WARN] No live targets remaining in candidate pool.")
		return
	}

	// Evaluate probability roll
	roll := rand.Float64()
	if roll > m.Probability {
		log.Printf("[SKIP] Probability check passed (roll: %.2f > threshold: %.2f). No termination this cycle.", roll, m.Probability)
		return
	}

	// Select a single random candidate
	idx := rand.Intn(len(live))
	victim := live[idx]

	if err := m.TerminateInstance(victim); err != nil {
		log.Printf("[ERROR] Failed to terminate %s: %v", victim.Name, err)
	}
}

func main() {
	targetPrefix := flag.String("target", "worker-", "Container or instance name prefix to target")
	intervalSec := flag.Int("interval", 3, "Interval in seconds between chaos evaluation cycles")
	probability := flag.Float64("probability", 0.7, "Probability (0.0 to 1.0) of terminating an instance per cycle")
	dryRun := flag.Bool("dry-run", false, "Simulate evaluation without actually terminating containers")
	simulate := flag.Bool("simulate", false, "Run in simulated in-memory mode without requiring Docker daemon")
	rounds := flag.Int("rounds", 3, "Total cycles to execute (0 for indefinite daemon)")
	flag.Parse()

	log.Printf("==================================================")
	log.Printf("Chaos Monkey local runner initialized")
	log.Printf("Target Prefix: %s | Probability: %.2f | Interval: %ds | Simulate: %t", *targetPrefix, *probability, *intervalSec, *simulate)
	log.Printf("==================================================")

	var engine *MonkeyEngine
	if *simulate {
		engine = NewSimulatedEngine(*targetPrefix, *probability, *dryRun, 4)
	} else {
		engine = &MonkeyEngine{
			TargetPrefix: *targetPrefix,
			Probability:  *probability,
			DryRun:       *dryRun,
			Simulated:    false,
		}
	}

	cycle := 0
	for {
		cycle++
		log.Printf("--- Cycle #%d ---", cycle)
		engine.Step()

		if *rounds > 0 && cycle >= *rounds {
			log.Printf("Completed %d rounds. Chaos Monkey execution finished.", *rounds)
			break
		}

		time.Sleep(time.Duration(*intervalSec) * time.Second)
	}
}
