package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// TestGetPlayoffPicture_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint PlayoffPicture` instead.
func TestGetPlayoffPicture_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "EastConfPlayoffPicture", "headers": ["TEAM_ID", "LEAGUE_ID", "SEASON_ID", "CONFERENCE", "RANK", "TEAM", "WINS", "LOSSES", "WIN_PCT", "GAMES_BACK", "CLINCHED", "ELIMINATED_FROM_PLAYOFFS", "CAN_WIN_CONF", "CAN_WIN_DIV"], "rowSet": [[1, "test", "test", "test", 1, "test", "test", "test", 1.5, "test", "test", 1.5, "test", "test"]]},
		{"name": "WestConfPlayoffPicture", "headers": ["TEAM_ID", "LEAGUE_ID", "SEASON_ID", "CONFERENCE", "RANK", "TEAM", "WINS", "LOSSES", "WIN_PCT", "GAMES_BACK", "CLINCHED", "ELIMINATED_FROM_PLAYOFFS", "CAN_WIN_CONF", "CAN_WIN_DIV"], "rowSet": [[1, "test", "test", "test", 1, "test", "test", "test", 1.5, "test", "test", 1.5, "test", "test"]]}
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

	req := PlayoffPictureRequest{
		SeasonID: parameters.Season("2023-24"),
	}

	resp, err := GetPlayoffPicture(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetPlayoffPicture: %v", err)
	}

	if len(resp.Data.EastConfPlayoffPicture) == 0 {
		t.Errorf("expected EastConfPlayoffPicture to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.WestConfPlayoffPicture) == 0 {
		t.Errorf("expected WestConfPlayoffPicture to be populated from the synthesized fixture, got empty")
	}
}
