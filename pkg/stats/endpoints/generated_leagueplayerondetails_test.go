package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeaguePlayerOnDetails_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeaguePlayerOnDetails` instead.
func TestGetLeaguePlayerOnDetails_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeaguePlayerOnDetails", "headers": ["TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "VS_PLAYER_ID", "VS_PLAYER_NAME", "COURT_STATUS", "GP", "MIN", "PLUS_MINUS", "NET_RATING", "OFF_RATING", "DEF_RATING"], "rowSet": [[1, "test", "test", 1, "test", "test", 1, 1.5, 1.5, "test", "test", "test"]]}
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

	req := LeaguePlayerOnDetailsRequest{}

	resp, err := GetLeaguePlayerOnDetails(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeaguePlayerOnDetails: %v", err)
	}

	if len(resp.Data.LeaguePlayerOnDetails) == 0 {
		t.Errorf("expected LeaguePlayerOnDetails to be populated from the synthesized fixture, got empty")
	}
}
