package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetScoreboardV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint ScoreboardV2` instead.
func TestGetScoreboardV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "GameHeader", "headers": ["GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "GAME_STATUS_ID", "GAME_STATUS_TEXT", "GAMECODE", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "SEASON", "LIVE_PERIOD", "LIVE_PC_TIME", "NATL_TV_BROADCASTER_ABBREVIATION", "LIVE_PERIOD_TIME_BCAST", "WH_STATUS"], "rowSet": [["test", 1, "test", "test", "test", "test", 1, 1, "test", 1, "test", "test", 1.5, "test"]]},
		{"name": "LineScore", "headers": ["GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY_NAME", "TEAM_WINS_LOSSES", "PTS_QTR1", "PTS_QTR2", "PTS_QTR3", "PTS_QTR4", "PTS_OT1", "PTS_OT2", "PTS_OT3", "PTS_OT4", "PTS_OT5", "PTS_OT6", "PTS_OT7", "PTS_OT8", "PTS_OT9", "PTS_OT10", "PTS", "FG_PCT", "FT_PCT", "FG3_PCT", "AST", "REB", "TOV"], "rowSet": [["test", 1, "test", 1, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "SeriesStandings", "headers": ["GAME_ID", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "GAME_DATE_EST", "HOME_TEAM_WINS", "HOME_TEAM_LOSSES", "SERIES_LEADER"], "rowSet": [["test", 1, 1, "test", "test", "test", "test"]]},
		{"name": "LastMeeting", "headers": ["GAME_ID", "LAST_GAME_ID", "LAST_GAME_DATE_EST", "LAST_GAME_HOME_TEAM_ID", "LAST_GAME_HOME_TEAM_CITY", "LAST_GAME_HOME_TEAM_NAME", "LAST_GAME_HOME_TEAM_ABBREVIATION", "LAST_GAME_HOME_TEAM_POINTS", "LAST_GAME_VISITOR_TEAM_ID", "LAST_GAME_VISITOR_TEAM_CITY", "LAST_GAME_VISITOR_TEAM_NAME", "LAST_GAME_VISITOR_TEAM_ABBREVIATION", "LAST_GAME_VISITOR_TEAM_POINTS"], "rowSet": [["test", "test", "test", 1, "test", "test", "test", 1.5, 1, "test", "test", "test", 1.5]]},
		{"name": "EastConfStandingsByDay", "headers": ["TEAM_ID", "LEAGUE_ID", "SEASON_ID", "STANDINGSDATE", "CONFERENCE", "TEAM", "G", "W", "L", "W_PCT", "HOME_RECORD", "ROAD_RECORD"], "rowSet": [[1, "test", "test", "test", "test", "test", "test", "test", "test", 1.5, "test", "test"]]},
		{"name": "WestConfStandingsByDay", "headers": ["TEAM_ID", "LEAGUE_ID", "SEASON_ID", "STANDINGSDATE", "CONFERENCE", "TEAM", "G", "W", "L", "W_PCT", "HOME_RECORD", "ROAD_RECORD"], "rowSet": [[1, "test", "test", "test", "test", "test", "test", "test", "test", 1.5, "test", "test"]]},
		{"name": "Available", "headers": ["GAME_ID", "PT_AVAILABLE"], "rowSet": [["test", "test"]]}
	]}`

	const wantPath = "/scoreboardv2"
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

	req := ScoreboardV2Request{
		GameDate: "1",
	}

	resp, err := GetScoreboardV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetScoreboardV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetScoreboardV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "scoreboardv2")
	}

	if len(resp.Data.GameHeader) == 0 {
		t.Errorf("expected GameHeader to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.LineScore) == 0 {
		t.Errorf("expected LineScore to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.SeriesStandings) == 0 {
		t.Errorf("expected SeriesStandings to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.LastMeeting) == 0 {
		t.Errorf("expected LastMeeting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.EastConfStandingsByDay) == 0 {
		t.Errorf("expected EastConfStandingsByDay to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.WestConfStandingsByDay) == 0 {
		t.Errorf("expected WestConfStandingsByDay to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.Available) == 0 {
		t.Errorf("expected Available to be populated from the synthesized fixture, got empty")
	}
}
