package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetBoxScoreSummaryV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint BoxScoreSummaryV2` instead.
func TestGetBoxScoreSummaryV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "GameSummary", "headers": ["GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "GAME_STATUS_ID", "GAME_STATUS_TEXT", "GAMECODE", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "SEASON", "LIVE_PERIOD", "LIVE_PC_TIME", "NATL_TV_BROADCASTER_ABBREVIATION", "LIVE_PERIOD_TIME_BCAST", "WH_STATUS"], "rowSet": [["test", 1, "test", "test", "test", "test", 1, 1, "test", 1, "test", "test", 1.5, "test"]]},
		{"name": "OtherStats", "headers": ["LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PTS_PAINT", "PTS_2ND_CHANCE", "PTS_FB", "LARGEST_LEAD", "LEAD_CHANGES", "TIMES_TIED", "TEAM_TURNOVERS", "TOTAL_TURNOVERS", "TEAM_REBOUNDS", "PTS_OFF_TO"], "rowSet": [["test", 1, "test", "test", 1.5, 1.5, 1.5, "test", "test", "test", "test", "test", 1.5, 1.5]]},
		{"name": "Officials", "headers": ["OFFICIAL_ID", "FIRST_NAME", "LAST_NAME", "JERSEY_NUM"], "rowSet": [["test", "test", "test", "test"]]},
		{"name": "InactivePlayers", "headers": ["PLAYER_ID", "FIRST_NAME", "LAST_NAME", "JERSEY_NUM", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION"], "rowSet": [[1, "test", "test", "test", 1, "test", "test", "test"]]},
		{"name": "GameInfo", "headers": ["GAME_DATE", "ATTENDANCE", "GAME_TIME"], "rowSet": [["test", "test", "test"]]},
		{"name": "LineScore", "headers": ["GAME_DATE_EST", "GAME_SEQUENCE", "GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY_NAME", "TEAM_WINS_LOSSES", "PTS_QTR1", "PTS_QTR2", "PTS_QTR3", "PTS_QTR4", "PTS_OT1", "PTS_OT2", "PTS_OT3", "PTS_OT4", "PTS_OT5", "PTS_OT6", "PTS_OT7", "PTS_OT8", "PTS_OT9", "PTS_OT10", "PTS", "FG_PCT", "FT_PCT", "FG3_PCT", "AST", "REB", "TOV"], "rowSet": [["test", 1, "test", 1, "test", "test", "test", 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "LastMeeting", "headers": ["GAME_ID", "GAME_DATE_EST", "GAME_DATE_TIME_EST", "HOME_TEAM_ID", "HOME_TEAM_CITY", "HOME_TEAM_NAME", "HOME_TEAM_ABBREVIATION", "HOME_TEAM_POINTS", "VISITOR_TEAM_ID", "VISITOR_TEAM_CITY", "VISITOR_TEAM_NAME", "VISITOR_TEAM_ABBREVIATION", "VISITOR_TEAM_POINTS"], "rowSet": [["test", "test", "test", 1, "test", "test", "test", "test", 1, "test", "test", "test", "test"]]},
		{"name": "SeasonSeries", "headers": ["GAME_ID", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "GAME_DATE_EST", "HOME_TEAM_WINS", "HOME_TEAM_LOSSES", "SERIES_LEADER"], "rowSet": [["test", 1, 1, "test", "test", "test", "test"]]},
		{"name": "AvailableVideo", "headers": ["GAME_ID", "VIDEO_AVAILABLE_FLAG"], "rowSet": [["test", 1]]}
	]}`

	const wantPath = "/boxscoresummaryv2"
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

	req := BoxScoreSummaryV2Request{
		GameID: "1",
	}

	resp, err := GetBoxScoreSummaryV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetBoxScoreSummaryV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetBoxScoreSummaryV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "boxscoresummaryv2")
	}

	if len(resp.Data.GameSummary) == 0 {
		t.Errorf("expected GameSummary to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.OtherStats) == 0 {
		t.Errorf("expected OtherStats to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.Officials) == 0 {
		t.Errorf("expected Officials to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.InactivePlayers) == 0 {
		t.Errorf("expected InactivePlayers to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.GameInfo) == 0 {
		t.Errorf("expected GameInfo to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.LineScore) == 0 {
		t.Errorf("expected LineScore to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.LastMeeting) == 0 {
		t.Errorf("expected LastMeeting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.SeasonSeries) == 0 {
		t.Errorf("expected SeasonSeries to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AvailableVideo) == 0 {
		t.Errorf("expected AvailableVideo to be populated from the synthesized fixture, got empty")
	}
}
