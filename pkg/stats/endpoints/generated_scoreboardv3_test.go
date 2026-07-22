package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetScoreboardV3_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint ScoreboardV3` instead.
func TestGetScoreboardV3_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "GameHeader", "headers": ["gameId", "gameCode", "gameStatus", "gameStatusText", "period", "gameClock", "gameTimeUTC", "gameEt", "regulationPeriods", "seriesGameNumber", "seriesText", "homeTeamId", "homeTeamName", "homeTeamCity", "homeTeamTricode", "homeTeamScore", "visitorTeamId", "visitorTeamName", "visitorTeamCity", "visitorTeamTricode", "visitorTeamScore"], "rowSet": [["test", "test", "test", "test", 1, "test", "test", "test", 1, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]}
	]}`

	const wantPath = "/scoreboardv3"
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

	req := ScoreboardV3Request{
		GameDate: "1",
	}

	resp, err := GetScoreboardV3(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetScoreboardV3: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetScoreboardV3 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "scoreboardv3")
	}

	if len(resp.Data.GameHeader) == 0 {
		t.Errorf("expected GameHeader to be populated from the synthesized fixture, got empty")
	}
}
