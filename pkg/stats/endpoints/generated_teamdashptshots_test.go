package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamDashPtShots_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamDashPtShots` instead.
func TestGetTeamDashPtShots_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "OverallShooting", "headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "GP", "G", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1, "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "GeneralShooting", "headers": ["TEAM_ID", "TEAM_NAME", "SHOT_TYPE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]},
		{"name": "ShotClockShooting", "headers": ["TEAM_ID", "TEAM_NAME", "SHOT_CLOCK_RANGE", "FGA_FREQUENCY", "FGM", "FGA", "FG_PCT", "EFG_PCT", "FG2A_FREQUENCY", "FG2M", "FG2A", "FG2_PCT", "FG3A_FREQUENCY", "FG3M", "FG3A", "FG3_PCT"], "rowSet": [[1, "test", "test", 1.5, 1, 1, 1.5, 1.5, "test", "test", "test", 1.5, 1.5, 1, 1, 1.5]]}
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

	req := TeamDashPtShotsRequest{
		TeamID: "1",
	}

	resp, err := GetTeamDashPtShots(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamDashPtShots: %v", err)
	}

	if len(resp.Data.OverallShooting) == 0 {
		t.Errorf("expected OverallShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.GeneralShooting) == 0 {
		t.Errorf("expected GeneralShooting to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.ShotClockShooting) == 0 {
		t.Errorf("expected ShotClockShooting to be populated from the synthesized fixture, got empty")
	}
}
