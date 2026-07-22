package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetHomepageV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint HomepageV2` instead.
func TestGetHomepageV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "GameHeader", "headers": ["GAME_ID", "GAME_DATE", "HOME_TEAM_ID", "HOME_TEAM_NAME", "HOME_TEAM_ABBREVIATION", "HOME_TEAM_SCORE", "VISITOR_TEAM_ID", "VISITOR_TEAM_NAME", "VISITOR_TEAM_ABBREVIATION", "VISITOR_TEAM_SCORE", "GAME_STATUS_TEXT"], "rowSet": [["test", "test", 1, "test", "test", "test", 1, "test", "test", "test", "test"]]}
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

	req := HomepageV2Request{}

	resp, err := GetHomepageV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetHomepageV2: %v", err)
	}

	if len(resp.Data.GameHeader) == 0 {
		t.Errorf("expected GameHeader to be populated from the synthesized fixture, got empty")
	}
}
