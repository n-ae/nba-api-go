package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScoreMatchupsV3_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScoreMatchupsV3` instead.
func TestGetBoxScoreMatchupsV3_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "HomeTeamPlayerMatchups", "headers": ["GAME_ID", "PERSON_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "MATCHUP_MIN_PTS", "PARTIAL_POSS", "PLAYER_PTS", "TEAM_PTS", "MATCHUP_AST", "MATCHUP_TOV", "MATCHUP_BLK", "MATCHUP_FGM", "MATCHUP_FGA", "MATCHUP_FG_PCT", "MATCHUP_FG3M", "MATCHUP_FG3A", "MATCHUP_FG3_PCT", "HELP_BLK", "HELP_FGM", "HELP_FGA", "HELP_FG_PCT", "SHOOTER_PLAYER_ID", "SHOOTER_PLAYER_NAME", "DEFENDER_PLAYER_ID", "DEFENDER_PLAYER_NAME", "SFL"], "rowSet": [["test", "test", "test", 1, "test", "test", "test", 1.5, 1.5, "test", "test", "test", "test", "test", 1.5, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, "test", 1, "test", "test"]]},
		{"name": "AwayTeamPlayerMatchups", "headers": ["GAME_ID", "PERSON_ID", "PLAYER_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "MATCHUP_MIN_PTS", "PARTIAL_POSS", "PLAYER_PTS", "TEAM_PTS", "MATCHUP_AST", "MATCHUP_TOV", "MATCHUP_BLK", "MATCHUP_FGM", "MATCHUP_FGA", "MATCHUP_FG_PCT", "MATCHUP_FG3M", "MATCHUP_FG3A", "MATCHUP_FG3_PCT", "HELP_BLK", "HELP_FGM", "HELP_FGA", "HELP_FG_PCT", "SHOOTER_PLAYER_ID", "SHOOTER_PLAYER_NAME", "DEFENDER_PLAYER_ID", "DEFENDER_PLAYER_NAME", "SFL"], "rowSet": [["test", "test", "test", 1, "test", "test", "test", 1.5, 1.5, "test", "test", "test", "test", "test", 1.5, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, "test", 1, "test", "test"]]}
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

	req := BoxScoreMatchupsV3Request{
		GameID: "1",
	}

	resp, err := GetBoxScoreMatchupsV3(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScoreMatchupsV3: %v", err)
	}

	if len(resp.Data.HomeTeamPlayerMatchups) == 0 {
		t.Errorf("expected HomeTeamPlayerMatchups to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AwayTeamPlayerMatchups) == 0 {
		t.Errorf("expected AwayTeamPlayerMatchups to be populated from the synthesized fixture, got empty")
	}
}
