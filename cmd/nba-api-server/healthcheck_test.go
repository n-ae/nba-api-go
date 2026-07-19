package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/n-ae/nba-api-go/pkg/stats"
)

func TestNewHealthChecker_InitialStatusIsUnknown(t *testing.T) {
	hc := NewHealthChecker(stats.NewDefaultClient())

	if got := hc.Status(); got != nbaAPIStatusUnknown {
		t.Errorf("expected initial status %q, got %q", nbaAPIStatusUnknown, got)
	}
}

func TestHealthChecker_Start_DoesNotBlock(t *testing.T) {
	hc := NewHealthChecker(stats.NewDefaultClient())

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan struct{})
	go func() {
		hc.Start(ctx, time.Minute)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
		t.Fatal("Start blocked instead of returning immediately")
	}

	// Start's background goroutine runs an initial hc.check(ctx) before
	// entering its ticker loop (see Start's doc comment), so at this
	// point a check is still in flight on another goroutine. Cancel now
	// and wait for it to observe cancellation and settle Status(), so
	// that goroutine's work (and its reads of any test-mutable package
	// state, e.g. nowFunc) can't still be running once this test
	// function returns and the next test starts - the same failure mode
	// TestHealthChecker_Check_ParentCancellationAbortsPromptly asserts
	// resolves quickly on its own, relied on here to bound the wait.
	cancel()
	deadline := time.After(3 * time.Second)
	for hc.Status() == nbaAPIStatusUnknown {
		select {
		case <-deadline:
			t.Fatal("background check did not settle Status() after cancellation")
		case <-time.After(time.Millisecond):
		}
	}
}

// TestHealthChecker_Check_Live exercises a real probe against NBA.com.
// Skipped by default so the rest of the suite stays offline; set
// INTEGRATION_TESTS=1 to run it, matching the convention used elsewhere
// in this repo for network-dependent tests.
func TestHealthChecker_Check_Live(t *testing.T) {
	if os.Getenv("INTEGRATION_TESTS") != "1" {
		t.Skip("Skipping integration test (set INTEGRATION_TESTS=1 to run)")
	}

	hc := NewHealthChecker(stats.NewDefaultClient())
	hc.check(context.Background())

	status := hc.Status()
	if status != nbaAPIStatusOperational && status != nbaAPIStatusDegraded {
		t.Errorf("expected operational or degraded after a live check, got %q", status)
	}
}

// TestHealthChecker_Check_ParentCancellationAbortsPromptly proves check's
// request deadline is derived from the parent context: a parent canceled
// before the call returns immediately (well under the probe's own 3s
// timeout) instead of running to completion. This is what lets server
// shutdown stop an in-flight background probe right away.
func TestHealthChecker_Check_ParentCancellationAbortsPromptly(t *testing.T) {
	hc := NewHealthChecker(stats.NewDefaultClient())

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already canceled before check runs

	start := time.Now()
	hc.check(ctx)
	elapsed := time.Since(start)

	if elapsed >= 3*time.Second {
		t.Errorf("expected check to abort promptly on a canceled parent, took %v", elapsed)
	}
	if got := hc.Status(); got != nbaAPIStatusDegraded {
		t.Errorf("expected status %q after an aborted check, got %q", nbaAPIStatusDegraded, got)
	}
}
