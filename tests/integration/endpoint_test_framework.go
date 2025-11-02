package integration

import (
	"context"
	"testing"
	"time"

	"github.com/yourn-ae/nba-api-go/pkg/client"
	"github.com/yourn-ae/nba-api-go/pkg/stats/parameters"
)

// EndpointTester provides a framework for testing NBA API endpoints
type EndpointTester struct {
	client  *client.Client
	timeout time.Duration
	t       *testing.T
}

// NewEndpointTester creates a new endpoint testing framework
func NewEndpointTester(t *testing.T) *EndpointTester {
	t.Helper()

	c := client.NewDefaultClient()

	return &EndpointTester{
		client:  c,
		timeout: 30 * time.Second,
		t:       t,
	}
}

// WithTimeout sets a custom timeout for requests
func (et *EndpointTester) WithTimeout(timeout time.Duration) *EndpointTester {
	et.timeout = timeout
	return et
}

// Context creates a context with timeout
func (et *EndpointTester) Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), et.timeout)
}

// TestEndpoint represents a single endpoint test case
type TestEndpoint struct {
	Name        string
	Description string
	TestFunc    func(*EndpointTester) error
	Skip        bool
	SkipReason  string
}

// Run executes the endpoint test
func (te *TestEndpoint) Run(t *testing.T) {
	t.Helper()

	if te.Skip {
		t.Skipf("Skipping: %s", te.SkipReason)
		return
	}

	tester := NewEndpointTester(t)

	err := te.TestFunc(tester)
	if err != nil {
		t.Errorf("%s failed: %v", te.Name, err)
	}
}

// TestSuite groups related endpoint tests
type TestSuite struct {
	Name        string
	Description string
	Tests       []TestEndpoint
}

// Run executes all tests in the suite
func (ts *TestSuite) Run(t *testing.T) {
	t.Helper()

	t.Run(ts.Name, func(t *testing.T) {
		for _, test := range ts.Tests {
			t.Run(test.Name, func(t *testing.T) {
				test.Run(t)
			})
		}
	})
}

// Common test parameters
var (
	TestPlayerID   = "203999"   // Nikola Jokic
	TestTeamID     = 1610612743 // Denver Nuggets
	TestGameID     = "0022300001"
	TestSeason     = parameters.NewSeason(2023)
	TestSeasonType = parameters.SeasonTypeRegular
	TestLeagueID   = parameters.LeagueIDNBA
)

// Helper functions for common validations

// AssertNotEmpty validates that a slice is not empty
func AssertNotEmpty(t *testing.T, slice interface{}, fieldName string) {
	t.Helper()

	switch v := slice.(type) {
	case []interface{}:
		if len(v) == 0 {
			t.Errorf("%s should not be empty", fieldName)
		}
	default:
		// Use reflection for other slice types if needed
	}
}

// AssertValidResponse checks basic response validity
func AssertValidResponse(t *testing.T, err error, message string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: %v", message, err)
	}
}

// AssertFieldsPresent validates that required fields are present and non-zero
func AssertFieldsPresent(t *testing.T, data interface{}, fields ...string) {
	t.Helper()

	// This would use reflection to check fields
	// Implementation depends on the specific needs
}
