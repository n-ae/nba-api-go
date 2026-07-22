package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueDashPlayerBioStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueDashPlayerBioStats` instead.
func TestGetLeagueDashPlayerBioStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "LeagueDashPlayerBioStats", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "AGE", "PLAYER_HEIGHT", "PLAYER_WEIGHT", "COLLEGE", "COUNTRY", "DRAFT_YEAR", "DRAFT_ROUND", "DRAFT_NUMBER", "GP", "PTS", "REB", "AST", "NET_RATING", "OREB_PCT", "DREB_PCT", "USG_PCT", "TS_PCT", "AST_PCT"], "rowSet": [[1, "test", 1, "test", 1, "test", "test", "test", "test", "test", "test", "test", 1, 1.5, 1.5, 1.5, "test", 1.5, 1.5, 1.5, 1.5, 1.5]]}
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

	req := LeagueDashPlayerBioStatsRequest{}

	resp, err := GetLeagueDashPlayerBioStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueDashPlayerBioStats: %v", err)
	}

	if len(resp.Data.LeagueDashPlayerBioStats) == 0 {
		t.Errorf("expected LeagueDashPlayerBioStats to be populated from the synthesized fixture, got empty")
	}
}
