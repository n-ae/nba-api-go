package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetCommonPlayerInfoV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint CommonPlayerInfoV2` instead.
func TestGetCommonPlayerInfoV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "CommonPlayerInfo", "headers": ["PERSON_ID", "FIRST_NAME", "LAST_NAME", "DISPLAY_FIRST_LAST", "DISPLAY_LAST_COMMA_FIRST", "DISPLAY_FI_LAST", "PLAYER_SLUG", "BIRTHDATE", "SCHOOL", "COUNTRY", "LAST_AFFILIATION", "HEIGHT", "WEIGHT", "SEASON_EXP", "JERSEY", "POSITION", "ROSTERSTATUS", "GAMES_PLAYED_CURRENT_SEASON_FLAG", "TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CODE", "TEAM_CITY", "PLAYERCODE", "FROM_YEAR", "TO_YEAR", "DLEAGUE_FLAG", "NBA_FLAG", "GAMES_PLAYED_FLAG", "DRAFT_YEAR", "DRAFT_ROUND", "DRAFT_NUMBER", "GREATEST_75_FLAG"], "rowSet": [["test", "test", "test", "test", "test", "test", "test", "test", "test", "test", 1.5, "test", "test", "test", "test", "test", "test", "test", 1, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]},
		{"name": "PlayerHeadlineStats", "headers": ["PLAYER_ID", "PLAYER_NAME", "TimeFrame", "PTS", "AST", "REB", "PIE"], "rowSet": [[1, "test", "test", 1.5, 1.5, 1.5, "test"]]}
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

	req := CommonPlayerInfoV2Request{
		PlayerID: "1",
	}

	resp, err := GetCommonPlayerInfoV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetCommonPlayerInfoV2: %v", err)
	}

	if len(resp.Data.CommonPlayerInfo) == 0 {
		t.Errorf("expected CommonPlayerInfo to be populated from the synthesized fixture, got empty")
	}
	if len(resp.Data.PlayerHeadlineStats) == 0 {
		t.Errorf("expected PlayerHeadlineStats to be populated from the synthesized fixture, got empty")
	}
}
