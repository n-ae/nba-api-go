package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerEstimatedAdvancedStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerEstimatedAdvancedStats` instead.
func TestGetPlayerEstimatedAdvancedStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerEstimatedAdvancedStats", "headers": ["PLAYER_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "GP", "MIN", "E_OFF_RATING", "E_DEF_RATING", "E_NET_RATING", "E_AST_RATIO", "E_OREB_PCT", "E_DREB_PCT", "E_REB_PCT", "E_TOV_PCT", "E_USG_PCT", "E_PACE"], "rowSet": [[1, "test", 1, "test", 1, 1.5, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, "test"]]}
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

	req := PlayerEstimatedAdvancedStatsRequest{}

	resp, err := GetPlayerEstimatedAdvancedStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerEstimatedAdvancedStats: %v", err)
	}

	if len(resp.Data.PlayerEstimatedAdvancedStats) == 0 {
		t.Errorf("expected PlayerEstimatedAdvancedStats to be populated from the synthesized fixture, got empty")
	}
}
