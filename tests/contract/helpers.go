package contract

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
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

// assertEqual fails if expected != actual
func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s: expected %v, got %v", message, expected, actual)
	}
}

// assertGreaterThan fails if value <= threshold
func assertGreaterThan(t *testing.T, value, threshold int, message string) {
	t.Helper()
	if value <= threshold {
		t.Fatalf("%s: expected > %d, got %d", message, threshold, value)
	}
}

// compareSchemas compares two JSON structures and reports differences
func compareSchemas(t *testing.T, expected, actual []byte) {
	t.Helper()

	var expectedMap, actualMap map[string]interface{}

	if err := json.Unmarshal(expected, &expectedMap); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}

	if err := json.Unmarshal(actual, &actualMap); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	differences := findSchemaDifferences(expectedMap, actualMap, "")
	if len(differences) > 0 {
		t.Errorf("Schema differences detected:\n")
		for _, diff := range differences {
			t.Errorf("  - %s\n", diff)
		}
	}
}

// findSchemaDifferences recursively compares two maps and returns differences
func findSchemaDifferences(expected, actual map[string]interface{}, path string) []string {
	var diffs []string

	// Check for missing keys in actual
	for key := range expected {
		newPath := path + "." + key
		if path == "" {
			newPath = key
		}

		if _, exists := actual[key]; !exists {
			diffs = append(diffs, newPath+" missing in actual response")
			continue
		}

		// Check types match
		expectedType := reflect.TypeOf(expected[key])
		actualType := reflect.TypeOf(actual[key])

		if expectedType != actualType {
			diffs = append(diffs, newPath+" type changed: "+expectedType.String()+" -> "+actualType.String())
			continue
		}

		// Recursively check nested objects
		if expectedMap, ok := expected[key].(map[string]interface{}); ok {
			if actualMap, ok := actual[key].(map[string]interface{}); ok {
				diffs = append(diffs, findSchemaDifferences(expectedMap, actualMap, newPath)...)
			}
		}

		// Recursively check arrays
		if expectedSlice, ok := expected[key].([]interface{}); ok {
			if actualSlice, ok := actual[key].([]interface{}); ok {
				if len(expectedSlice) > 0 && len(actualSlice) > 0 {
					// Compare first element schemas
					if expMap, ok := expectedSlice[0].(map[string]interface{}); ok {
						if actMap, ok := actualSlice[0].(map[string]interface{}); ok {
							diffs = append(diffs, findSchemaDifferences(expMap, actMap, newPath+"[0]")...)
						}
					}
				}
			}
		}
	}

	// Check for new keys in actual
	for key := range actual {
		newPath := path + "." + key
		if path == "" {
			newPath = key
		}

		if _, exists := expected[key]; !exists {
			diffs = append(diffs, newPath+" added in actual response (new field)")
		}
	}

	return diffs
}

// prettyJSON formats JSON for readable output
func prettyJSON(t *testing.T, data []byte) string {
	t.Helper()

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return string(data)
	}

	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return string(data)
	}

	return string(pretty)
}

// Helper functions for pointer conversions

// stringPtr returns a pointer to the given string
func stringPtr(s string) *string {
	return &s
}

// intPtr returns a pointer to the given int
func intPtr(i int) *int {
	return &i
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
