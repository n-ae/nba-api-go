package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerDashboardByYearOverYear_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerDashboardByYearOverYear` instead.
func TestGetPlayerDashboardByYearOverYear_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "OverallPlayerDashboard", "headers": ["PLAYER_ID", "PLAYER_NAME", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "ByYearPlayerDashboard", "headers": ["PLAYER_ID", "PLAYER_NAME", "YEAR", "GP", "W", "L", "W_PCT", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "TOV", "STL", "BLK", "BLKA", "PF", "PFD", "PTS", "PLUS_MINUS"], "rowSet": [[1, "test", "test", 1, "test", "test", 1.5, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1, 1.5, 1.5, 1.5, 1.5]]}
	]}`

	const wantPath = "/playerdashboardbyyearoveryear"
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

	req := PlayerDashboardByYearOverYearRequest{
		PlayerID: "1",
	}

	resp, err := GetPlayerDashboardByYearOverYear(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerDashboardByYearOverYear: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetPlayerDashboardByYearOverYear requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "playerdashboardbyyearoveryear")
	}

	if len(resp.Data.OverallPlayerDashboard) == 0 {
		t.Errorf("expected OverallPlayerDashboard to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.ByYearPlayerDashboard) == 0 {
		t.Errorf("expected ByYearPlayerDashboard to be populated from the synthesized fixture, got empty")
	}
}
