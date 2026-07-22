package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetAssistLeaders_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint AssistLeaders` instead.
func TestGetAssistLeaders_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "AssistLeaders", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "AST"], "rowSet": [[1, "test", 1, "test", 1, 1.5, 1.5]]}
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

	req := AssistLeadersRequest{}

	resp, err := GetAssistLeaders(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetAssistLeaders: %v", err)
	}

	if len(resp.Data.AssistLeaders) == 0 {
		t.Errorf("expected AssistLeaders to be populated from the synthesized fixture, got empty")
	}
}
