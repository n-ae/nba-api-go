package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayByPlayV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayByPlayV2` instead.
func TestGetPlayByPlayV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "PlayByPlay", "headers": ["GAME_ID", "EVENTNUM", "EVENTMSGTYPE", "EVENTMSGACTIONTYPE", "PERIOD", "WCTIMESTRING", "PCTIMESTRING", "HOMEDESCRIPTION", "NEUTRALDESCRIPTION", "VISITORDESCRIPTION", "SCORE", "SCOREMARGIN", "PERSON1TYPE", "PLAYER1_ID", "PLAYER1_NAME", "PLAYER1_TEAM_ID", "PLAYER1_TEAM_CITY", "PLAYER1_TEAM_NICKNAME", "PLAYER1_TEAM_ABBREVIATION", "PERSON2TYPE", "PLAYER2_ID", "PLAYER2_NAME", "PLAYER2_TEAM_ID", "PLAYER2_TEAM_CITY", "PLAYER2_TEAM_NICKNAME", "PLAYER2_TEAM_ABBREVIATION", "PERSON3TYPE", "PLAYER3_ID", "PLAYER3_NAME", "PLAYER3_TEAM_ID", "PLAYER3_TEAM_CITY", "PLAYER3_TEAM_NICKNAME", "PLAYER3_TEAM_ABBREVIATION", "VIDEO_AVAILABLE_FLAG"], "rowSet": [["test", 1, 1, 1, 1, "test", "test", "test", "test", "test", "test", "test", 1, 1, "test", 1, "test", "test", "test", 1, 1, "test", 1, "test", "test", "test", 1, 1, "test", 1, "test", "test", "test", 1]]},
		{"name": "AvailableVideo", "headers": ["GAME_ID", "VIDEO_AVAILABLE_FLAG"], "rowSet": [["test", 1]]}
	]}`

	const wantPath = "/playbyplayv2"
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

	req := PlayByPlayV2Request{
		GameID: "1",
	}

	resp, err := GetPlayByPlayV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayByPlayV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayByPlayV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playbyplayv2")
	}

	if len(resp.Data.PlayByPlay) == 0 {
		t.Errorf("expected PlayByPlay to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.AvailableVideo) == 0 {
		t.Errorf("expected AvailableVideo to be populated from the synthesized fixture, got empty")
	}
}
