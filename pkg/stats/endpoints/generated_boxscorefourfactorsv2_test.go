package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScoreFourFactorsV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScoreFourFactorsV2` instead.
func TestGetBoxScoreFourFactorsV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PLAYER_ID", "PLAYER_NAME", "NICKNAME", "START_POSITION", "COMMENT", "MIN", "EFG_PCT", "FTA_RATE", "TM_TOV_PCT", "OREB_PCT", "OPP_EFG_PCT", "OPP_FTA_RATE", "OPP_TOV_PCT", "OPP_OREB_PCT"], "rowSet": [["test", 1, "test", "test", 1, "test", "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "TeamStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CITY", "MIN", "EFG_PCT", "FTA_RATE", "TM_TOV_PCT", "OREB_PCT", "OPP_EFG_PCT", "OPP_FTA_RATE", "OPP_TOV_PCT", "OPP_OREB_PCT"], "rowSet": [["test", 1, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
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

	req := BoxScoreFourFactorsV2Request{
		GameID: "1",
	}

	resp, err := GetBoxScoreFourFactorsV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScoreFourFactorsV2: %v", err)
	}

	if len(resp.Data.PlayerStats) == 0 {
		t.Errorf("expected PlayerStats to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamStats) == 0 {
		t.Errorf("expected TeamStats to be populated from the synthesized fixture, got empty")
	}
}
