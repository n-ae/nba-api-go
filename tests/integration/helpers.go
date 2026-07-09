package integration

import (
	"os"
	"testing"
	"time"
)

// Test constants for known stable IDs
const (
	LeBronJamesID  = "2544"
	NikolaJokicID  = "203999"
	LakersTeamID   = 1610612747
	NuggetsTeamID  = 1610612743
	Season2023     = "2023-24"
	DefaultTimeout = 30 * time.Second
)

// skipIfNotIntegration skips the test if INTEGRATION_TESTS env var is not set
func skipIfNotIntegration(t *testing.T) {
	t.Helper()
	if os.Getenv("INTEGRATION_TESTS") != "1" {
		t.Skip("Skipping integration test (set INTEGRATION_TESTS=1 to run)")
	}
}

// assertNoError fails the test if err is not nil
func assertNoError(t *testing.T, err error, message string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %v", message, err)
	}
}
