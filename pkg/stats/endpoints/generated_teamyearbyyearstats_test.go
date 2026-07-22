package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamYearByYearStats_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamYearByYearStats` instead.
func TestGetTeamYearByYearStats_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "TeamStats", "headers": ["TEAM_ID", "TEAM_CITY", "TEAM_NAME", "YEAR", "GP", "WINS", "LOSSES", "WIN_PCT", "CONF_RANK", "DIV_RANK", "PO_WINS", "PO_LOSSES", "CONF_COUNT", "DIV_COUNT", "NBA_FINALS_APPEARANCE", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "PF", "STL", "TOV", "BLK", "PTS", "PTS_RANK"], "rowSet": [[1, "test", "test", "test", 1, 1, 1, 1.5, 1, 1, 1, 1, 1, 1, "test", 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]]}
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

	req := TeamYearByYearStatsRequest{
		TeamID: "1",
	}

	resp, err := GetTeamYearByYearStats(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamYearByYearStats: %v", err)
	}

	if len(resp.Data.TeamStats) == 0 {
		t.Errorf("expected TeamStats to be populated from the synthesized fixture, got empty")
	}
}
