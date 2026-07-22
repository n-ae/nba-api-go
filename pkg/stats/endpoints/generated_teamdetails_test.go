package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetTeamDetails_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint TeamDetails` instead.
func TestGetTeamDetails_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "TeamBackground", "headers": ["TEAM_ID", "ABBREVIATION", "NICKNAME", "YEARFOUNDED", "CITY", "ARENA", "ARENACAPACITY", "OWNER", "GENERALMANAGER", "HEADCOACH", "DLEAGUEAFFILIATION"], "rowSet": [[1, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]},
		{"name": "TeamHistory", "headers": ["TEAM_ID", "CITY", "NICKNAME", "YEARFOUNDED", "YEARACTIVETILL"], "rowSet": [[1, "test", "test", "test", "test"]]},
		{"name": "TeamSocialSites", "headers": ["ACCOUNTTYPE", "WEBSITE_LINK"], "rowSet": [["test", "test"]]},
		{"name": "TeamAwardsChampionships", "headers": ["YEARAWARDED", "OPPOSITETEAM"], "rowSet": [["test", "test"]]},
		{"name": "TeamAwardsConf", "headers": ["YEARAWARDED", "OPPOSITETEAM"], "rowSet": [["test", "test"]]},
		{"name": "TeamAwardsDiv", "headers": ["YEARAWARDED", "OPPOSITETEAM"], "rowSet": [["test", "test"]]},
		{"name": "TeamHof", "headers": ["PLAYERID", "PLAYER", "POSITION", "JERSEY", "SEASONSWITHTEAM", "YEAR"], "rowSet": [["test", "test", "test", "test", "test", "test"]]},
		{"name": "TeamRetired", "headers": ["PLAYERID", "PLAYER", "POSITION", "JERSEY", "SEASONSWITHTEAM", "YEAR"], "rowSet": [["test", "test", "test", "test", "test", "test"]]}
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

	req := TeamDetailsRequest{
		TeamID: "1",
	}

	resp, err := GetTeamDetails(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetTeamDetails: %v", err)
	}

	if len(resp.Data.TeamBackground) == 0 {
		t.Errorf("expected TeamBackground to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamHistory) == 0 {
		t.Errorf("expected TeamHistory to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamSocialSites) == 0 {
		t.Errorf("expected TeamSocialSites to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamAwardsChampionships) == 0 {
		t.Errorf("expected TeamAwardsChampionships to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamAwardsConf) == 0 {
		t.Errorf("expected TeamAwardsConf to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamAwardsDiv) == 0 {
		t.Errorf("expected TeamAwardsDiv to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamHof) == 0 {
		t.Errorf("expected TeamHof to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.TeamRetired) == 0 {
		t.Errorf("expected TeamRetired to be populated from the synthesized fixture, got empty")
	}
}
