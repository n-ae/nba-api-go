package integration

import (
	"os"
	"testing"
	"time"

	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
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

// assertNotEmpty fails the test if a slice is empty
func assertNotEmpty(t *testing.T, slice interface{}, fieldName string) {
	t.Helper()

	switch v := slice.(type) {
	case []interface{}:
		if len(v) == 0 {
			t.Errorf("%s should not be empty", fieldName)
		}
	case nil:
		t.Errorf("%s should not be nil", fieldName)
	}
}

// assertNoError fails the test if err is not nil
func assertNoError(t *testing.T, err error, message string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %v", message, err)
	}
}

// Helper functions to create pointers for parameter types
func seasonPtr(s string) *parameters.Season {
	season := parameters.Season(s)
	return &season
}

func seasonTypePtr(s parameters.SeasonType) *parameters.SeasonType {
	return &s
}

func leagueIDPtr(l parameters.LeagueID) *parameters.LeagueID {
	return &l
}

func perModePtr(p parameters.PerMode) *parameters.PerMode {
	return &p
}
