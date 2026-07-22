package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScoreAdvancedV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScoreAdvancedV2` instead.
func TestGetBoxScoreAdvancedV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayerStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PLAYER_ID", "PLAYER_NAME", "NICKNAME", "START_POSITION", "COMMENT", "MIN", "E_OFF_RATING", "OFF_RATING", "E_DEF_RATING", "DEF_RATING", "E_NET_RATING", "NET_RATING", "AST_PCT", "AST_TOV", "AST_RATIO", "OREB_PCT", "DREB_PCT", "REB_PCT", "TM_TOV_PCT", "EFG_PCT", "TS_PCT", "USG_PCT", "E_USG_PCT", "E_PACE", "PACE", "PACE_PER40", "POSS", "PIE"], "rowSet": [["test", 1, "test", "test", 1, "test", "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, "test", "test", "test", "test", "test"]]},
		{"name": "TeamStats", "headers": ["GAME_ID", "TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CITY", "MIN", "E_OFF_RATING", "OFF_RATING", "E_DEF_RATING", "DEF_RATING", "E_NET_RATING", "NET_RATING", "AST_PCT", "AST_TOV", "AST_RATIO", "OREB_PCT", "DREB_PCT", "REB_PCT", "E_TM_TOV_PCT", "TM_TOV_PCT", "EFG_PCT", "TS_PCT", "USG_PCT", "E_USG_PCT", "E_PACE", "PACE", "PACE_PER40", "POSS", "PIE"], "rowSet": [["test", 1, "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, "test", "test", "test", "test", "test"]]}
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

	req := BoxScoreAdvancedV2Request{
		GameID: "1",
	}

	resp, err := GetBoxScoreAdvancedV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScoreAdvancedV2: %v", err)
	}

	if len(resp.Data.PlayerStats) == 0 {
		t.Errorf("expected PlayerStats to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamStats) == 0 {
		t.Errorf("expected TeamStats to be populated from the synthesized fixture, got empty")
	}
}
