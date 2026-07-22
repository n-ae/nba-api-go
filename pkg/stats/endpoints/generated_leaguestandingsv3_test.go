package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// TestGetLeagueStandingsV3_Generated is generated from tools/generator/metadata;
// see tools/generator/templates/endpoint_test.tmpl. Verifies real response
// parsing (header validation + positional row decoding) against a
// synthesized fixture matching this endpoint's actual result-set column
// names - not just request construction, which TestGeneratedHandlers
// (cmd/nba-api-server) already exercises indirectly for every endpoint.
// Do not hand-edit - regenerate via `cd tools/generator && go run . -endpoint LeagueStandingsV3` instead.
func TestGetLeagueStandingsV3_Generated(t *testing.T) {
	const responseBody = `{"resultSets": [
		{"name": "Standings", "headers": ["TeamID", "TeamCity", "TeamName", "Conference", "ConferenceRecord", "PlayoffRank", "ClinchIndicator", "DivisionRecord", "DivisionRank", "WINS", "LOSSES", "WinPCT", "LeagueRank", "Record", "HOME", "ROAD", "L10", "Last10Home", "Last10Road", "OT", "ThreePTSOrLess", "TenPTSOrMore", "LongHomeStreak", "strLongHomeStreak", "LongRoadStreak", "strLongRoadStreak", "LongWinStreak", "LongLossStreak", "CurrentHomeStreak", "strCurrentHomeStreak", "CurrentRoadStreak", "strCurrentRoadStreak", "CurrentStreak", "strCurrentStreak", "ConferenceGamesBack", "ClinchedConferenceTitle", "ClinchedDivisionTitle", "ClinchedPlayoffBirth", "EliminatedConference", "EliminatedDivision", "AheadAtHalf", "BehindAtHalf", "TiedAtHalf", "AheadAtThird", "BehindAtThird", "TiedAtThird", "Score100PTS", "OppScore100PTS", "OppOver500", "LeadInFGPCT", "LeadInReb", "FewerTurnovers", "PointsPG", "OppPointsPG", "DiffPointsPG", "vsEast", "vsAtlantic", "vsCentral", "vsSoutheast", "vsWest", "vsNorthwest", "vsPacific", "vsSouthwest", "Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"], "rowSet": [["test", "test", "test", "test", "test", 1, "test", "test", 1, "test", "test", 1.5, 1, "test", "test", "test", "test", 1.5, 1.5, "test", 1.5, 1.5, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", 1.5, 1.5, "test", "test", "test", "test", "test", "test", 1.5, 1.5, "test", 1.5, 1.5, "test", "test", "test", "test", 1.5, "test", "test", 1.5, "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test", "test"]]}
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

	req := LeagueStandingsV3Request{}

	resp, err := GetLeagueStandingsV3(context.Background(), client, req)
	if err != nil {
		t.Fatalf("GetLeagueStandingsV3: %v", err)
	}

	if len(resp.Data.Standings) == 0 {
		t.Errorf("expected Standings to be populated from the synthesized fixture, got empty")
	}
}
