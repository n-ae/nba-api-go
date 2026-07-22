package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueSeasonMatchups_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueSeasonMatchups` instead.
func TestGetLeagueSeasonMatchups_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "SeasonMatchups", "headers": ["SEASON_ID", "OFF_PLAYER_ID", "OFF_PLAYER_NAME", "DEF_PLAYER_ID", "DEF_PLAYER_NAME", "GP", "MATCHUP_MIN", "PARTIAL_POSS", "PLAYER_PTS", "TEAM_PTS", "MATCHUP_AST", "MATCHUP_TOV", "MATCHUP_BLK", "MATCHUP_FGM", "MATCHUP_FGA", "MATCHUP_FG_PCT", "SFL"], "rowSet": [["test", 1, "test", 1, "test", 1, "test", "test", 1.5, 1.5, "test", "test", "test", "test", "test", 1.5, "test"]]}
	]}`

	const wantPath = "/leagueseasonmatchups"
	var gotPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: server.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}

	req := LeagueSeasonMatchupsRequest{}

	resp, err := GetLeagueSeasonMatchups(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueSeasonMatchups: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetLeagueSeasonMatchups requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "leagueseasonmatchups")
	}

	if len(resp.Data.SeasonMatchups) == 0 {
		t.Errorf("expected SeasonMatchups to be populated from the synthesized fixture, got empty")
	}
}
