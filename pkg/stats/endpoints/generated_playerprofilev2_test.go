package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetPlayerProfileV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayerProfileV2` instead.
func TestGetPlayerProfileV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "SeasonTotalsRegularSeason", "headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION", "PLAYER_AGE", "GP", "GS", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "STL", "BLK", "TOV", "PF", "PTS"], "rowSet": [[1, "test", "test", 1, "test", 1, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]},
		{"name": "CareerTotalsRegularSeason", "headers": ["PLAYER_ID", "LEAGUE_ID", "Team_ID", "GP", "GS", "MIN", "FGM", "FGA", "FG_PCT", "FG3M", "FG3A", "FG3_PCT", "FTM", "FTA", "FT_PCT", "OREB", "DREB", "REB", "AST", "STL", "BLK", "TOV", "PF", "PTS"], "rowSet": [[1, "test", 1, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1, 1, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5, 1.5]]}
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

	req := PlayerProfileV2Request{
		PlayerID: "1",
	}

	resp, err := GetPlayerProfileV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayerProfileV2: %v", err)
	}

	if len(resp.Data.SeasonTotalsRegularSeason) == 0 {
		t.Errorf("expected SeasonTotalsRegularSeason to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.CareerTotalsRegularSeason) == 0 {
		t.Errorf("expected CareerTotalsRegularSeason to be populated from the synthesized fixture, got empty")
	}
}
