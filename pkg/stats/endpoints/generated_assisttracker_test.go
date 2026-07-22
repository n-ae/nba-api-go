package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetAssistTracker_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint AssistTracker` instead.
func TestGetAssistTracker_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "AssistTracker", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "W", "L", "W_PCT", "MIN", "AST", "PASS_TO", "AST_PTS_CREATED", "AST_PTS_CREATED_PER_PASS", "AST_PCT", "AST_ADJ"], "rowSet": [[1, "test", 1, "test", 1, "test", "test", 1.5, 1.5, 1.5, "test", 1.5, 1.5, 1.5, 1.5]]}
	]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}

	req := AssistTrackerRequest{}

	resp, err := GetAssistTracker(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetAssistTracker: %v", err)
	}

	if len(resp.Data.AssistTracker) == 0 {
		t.Errorf("expected AssistTracker to be populated from the synthesized fixture, got empty")
	}
}
