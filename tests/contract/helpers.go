package contract

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
)

const fixturesDir = "fixtures"

// loadFixture reads a fixture file and returns its contents
// If fixture doesn't exist and UPDATE_FIXTURES is not set, skips the test
func loadFixture(t *testing.T, filename string) []byte {
	t.Helper()

	path := filepath.Join(fixturesDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) && !shouldUpdateFixtures() {
			t.Skipf("Fixture %s not found (run with UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 to record)", filename)
		}
		t.Fatalf("Failed to read fixture %s: %v", filename, err)
	}

	return data
}

// saveFixture writes data to a fixture file
func saveFixture(t *testing.T, filename string, data []byte) {
	t.Helper()

	// Create fixtures directory if it doesn't exist
	if err := os.MkdirAll(fixturesDir, 0755); err != nil {
		t.Fatalf("Failed to create fixtures directory: %v", err)
	}

	path := filepath.Join(fixturesDir, filename)
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatalf("Failed to write fixture %s: %v", filename, err)
	}

	t.Logf("Saved fixture: %s (%d bytes)", filename, len(data))
}

// shouldUpdateFixtures returns true if we should record new fixtures
func shouldUpdateFixtures() bool {
	return os.Getenv("UPDATE_FIXTURES") == "1"
}

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

// assert fails the test if condition is false
func assert(t *testing.T, condition bool, message string) {
	t.Helper()
	if !condition {
		t.Fatalf("Assertion failed: %s", message)
	}
}

// Helper functions for pointer conversions

// stringPtr returns a pointer to the given string
func stringPtr(s string) *string {
	return &s
}

// perModePtr returns a pointer to PerMode
func perModePtr(pm parameters.PerMode) *parameters.PerMode {
	return &pm
}

// leagueIDPtr returns a pointer to LeagueID
func leagueIDPtr(id parameters.LeagueID) *parameters.LeagueID {
	return &id
}

// seasonPtr returns a pointer to Season
func seasonPtr(s parameters.Season) *parameters.Season {
	return &s
}

// seasonTypePtr returns a pointer to SeasonType
func seasonTypePtr(st parameters.SeasonType) *parameters.SeasonType {
	return &st
}
