package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetCommonAllPlayersV2_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint -
// and asserts the outbound request path matches this endpoint's own
// metadata exactly, the class of bug ten endpoints shipped with before a
// live-reachability sweep caught it (see CHANGELOG.md's [3.1.0] section).
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint CommonAllPlayersV2` instead.
func TestGetCommonAllPlayersV2_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "CommonAllPlayers", "headers": ["PERSON_ID", "DISPLAY_LAST_COMMA_FIRST", "DISPLAY_FIRST_LAST", "ROSTERSTATUS", "FROM_YEAR", "TO_YEAR", "PLAYERCODE", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CODE", "GAMES_PLAYED_FLAG", "OTHERLEAGUE_EXPERIENCE_CH"], "rowSet": [["test", "test", "test", "test", "test", "test", "test", 1, "test", "test", "test", "test", "test", "test"]]}
	]}`

	const wantPath = "/commonallplayersv2"
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

	req := CommonAllPlayersV2Request{}

	resp, err := GetCommonAllPlayersV2(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetCommonAllPlayersV2: %v", err)
	}

	if gotPath != wantPath {
		t.Errorf("GetCommonAllPlayersV2 requested path %q, want %q (endpoint metadata says %q)", gotPath, wantPath, "commonallplayersv2")
	}

	if len(resp.Data.CommonAllPlayers) == 0 {
		t.Errorf("expected CommonAllPlayers to be populated from the synthesized fixture, got empty")
	}
}
