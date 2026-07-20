package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/n-ae/nba-api-go/v2/pkg/stats"
	"github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
)

// newStatsFixtureClient serves a raw NBA Stats response so these tests cover
// each endpoint's real decode and header-validation path, rather than merely
// calling validateHeaders in isolation.
func newStatsFixtureClient(t *testing.T, body interface{}) *stats.Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(body); err != nil {
			t.Errorf("encode fixture response: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	return stats.NewClient(stats.Config{BaseURL: server.URL})
}

func fixtureRow(columns int, values ...interface{}) []interface{} {
	row := make([]interface{}, columns)
	copy(row, values)
	return row
}

// arbitraryWrongHeader is a stand-in wrong column name used to corrupt a
// fixture's headers in tests below that don't care which column is wrong,
// only that validateHeaders rejects the mismatch - factored out (rather
// than repeating a literal) to satisfy golangci-lint's goconst. The
// LeagueLeaders test uses leagueLeadersPlayerIDColumn instead, since there
// the value is a real column name from that endpoint's actual header set,
// not an arbitrary one.
const arbitraryWrongHeader = "PLAYER_ID"

func TestCommonPlayerInfoValidatesResultSetHeaders(t *testing.T) {
	valid := rawStatsResponse{ResultSets: []resultSet{
		{Name: "CommonPlayerInfo", Headers: jsonTags(PlayerInfo{}), RowSet: [][]interface{}{fixtureRow(len(jsonTags(PlayerInfo{})), float64(201939), "Stephen")}},
		{Name: "PlayerHeadlineStats", Headers: jsonTags(HeadlineStats{}), RowSet: [][]interface{}{fixtureRow(len(jsonTags(HeadlineStats{})), float64(201939), "Stephen Curry")}},
		{Name: "AvailableSeasons", Headers: jsonTags(AvailableSeason{}), RowSet: [][]interface{}{{"2023-24"}}},
	}}
	resp, err := CommonPlayerInfo(context.Background(), newStatsFixtureClient(t, valid), CommonPlayerInfoRequest{PlayerID: "201939"})
	if err != nil {
		t.Fatalf("CommonPlayerInfo() error = %v", err)
	}
	if got := resp.Data.CommonPlayerInfo[0].PersonID; got != 201939 {
		t.Errorf("CommonPlayerInfo()[0].PersonID = %d, want 201939", got)
	}

	invalid := valid
	invalid.ResultSets = append([]resultSet(nil), valid.ResultSets...)
	invalid.ResultSets[0].Headers = append([]string(nil), valid.ResultSets[0].Headers...)
	invalid.ResultSets[0].Headers[0] = arbitraryWrongHeader
	_, err = CommonPlayerInfo(context.Background(), newStatsFixtureClient(t, invalid), CommonPlayerInfoRequest{PlayerID: "201939"})
	if err == nil || !strings.Contains(err.Error(), "CommonPlayerInfo result set") {
		t.Fatalf("CommonPlayerInfo() error = %v, want a CommonPlayerInfo header error", err)
	}
}

func TestPlayerGameLogValidatesResultSetHeaders(t *testing.T) {
	valid := rawStatsResponse{ResultSets: []resultSet{{
		Name:    "PlayerGameLog",
		Headers: jsonTags(GameLog{}),
		RowSet:  [][]interface{}{fixtureRow(len(jsonTags(GameLog{})), "2023-24", float64(203999))},
	}}}
	request := PlayerGameLogRequest{PlayerID: "203999", Season: "2023-24", SeasonType: parameters.SeasonTypeRegular}
	resp, err := PlayerGameLog(context.Background(), newStatsFixtureClient(t, valid), request)
	if err != nil {
		t.Fatalf("PlayerGameLog() error = %v", err)
	}
	if got := resp.Data.PlayerGameLog[0].PlayerID; got != 203999 {
		t.Errorf("PlayerGameLog()[0].PlayerID = %d, want 203999", got)
	}

	invalid := valid
	invalid.ResultSets = append([]resultSet(nil), valid.ResultSets...)
	invalid.ResultSets[0].Headers = append([]string(nil), valid.ResultSets[0].Headers...)
	invalid.ResultSets[0].Headers[1] = arbitraryWrongHeader
	_, err = PlayerGameLog(context.Background(), newStatsFixtureClient(t, invalid), request)
	if err == nil || !strings.Contains(err.Error(), "PlayerGameLog result set") {
		t.Fatalf("PlayerGameLog() error = %v, want a PlayerGameLog header error", err)
	}
}

func TestTeamGameLogValidatesResultSetHeaders(t *testing.T) {
	valid := rawStatsResponse{ResultSets: []resultSet{{
		Name:    "TeamGameLog",
		Headers: jsonTags(TeamGameLog{}),
		RowSet:  [][]interface{}{fixtureRow(len(jsonTags(TeamGameLog{})), float64(1610612747), "0022300001")},
	}}}
	request := TeamGameLogRequest{TeamID: "1610612747", Season: "2023-24", SeasonType: parameters.SeasonTypeRegular}
	resp, err := GetTeamGameLog(context.Background(), newStatsFixtureClient(t, valid), request)
	if err != nil {
		t.Fatalf("GetTeamGameLog() error = %v", err)
	}
	if got := resp.Data.TeamGameLog[0].TeamID; got != 1610612747 {
		t.Errorf("GetTeamGameLog()[0].TeamID = %d, want 1610612747", got)
	}

	invalid := valid
	invalid.ResultSets = append([]resultSet(nil), valid.ResultSets...)
	invalid.ResultSets[0].Headers = append([]string(nil), valid.ResultSets[0].Headers...)
	invalid.ResultSets[0].Headers[0] = "TEAM_ID"
	_, err = GetTeamGameLog(context.Background(), newStatsFixtureClient(t, invalid), request)
	if err == nil || !strings.Contains(err.Error(), "TeamGameLog result set") {
		t.Fatalf("GetTeamGameLog() error = %v, want a TeamGameLog header error", err)
	}
}

func TestPlayerCareerStatsValidatesResultSetHeaders(t *testing.T) {
	valid := rawStatsResponse{ResultSets: []resultSet{
		{Name: "SeasonTotalsRegularSeason", Headers: jsonTags(SeasonStat{}), RowSet: [][]interface{}{fixtureRow(len(jsonTags(SeasonStat{})), float64(201939), "2023-24")}},
		{Name: "CareerTotalsRegularSeason", Headers: jsonTags(CareerTotalStat{}), RowSet: [][]interface{}{fixtureRow(len(jsonTags(CareerTotalStat{})), float64(201939))}},
	}}
	request := PlayerCareerStatsRequest{PlayerID: "201939"}
	resp, err := PlayerCareerStats(context.Background(), newStatsFixtureClient(t, valid), request)
	if err != nil {
		t.Fatalf("PlayerCareerStats() error = %v", err)
	}
	if got := resp.Data.SeasonTotalsRegularSeason[0].PlayerID; got != 201939 {
		t.Errorf("PlayerCareerStats().SeasonTotalsRegularSeason[0].PlayerID = %d, want 201939", got)
	}
	if got := resp.Data.CareerTotalsRegularSeason[0].PlayerID; got != 201939 {
		t.Errorf("PlayerCareerStats().CareerTotalsRegularSeason[0].PlayerID = %d, want 201939", got)
	}

	invalid := valid
	invalid.ResultSets = append([]resultSet(nil), valid.ResultSets...)
	invalid.ResultSets[0].Headers = append([]string(nil), valid.ResultSets[0].Headers...)
	invalid.ResultSets[0].Headers[0] = "GP"
	_, err = PlayerCareerStats(context.Background(), newStatsFixtureClient(t, invalid), request)
	if err == nil || !strings.Contains(err.Error(), "SeasonTotalsRegularSeason result set") {
		t.Fatalf("PlayerCareerStats() error = %v, want a SeasonTotalsRegularSeason header error", err)
	}
}

func TestLeagueLeadersUsesItsSingularEnvelopeAndHeaderNames(t *testing.T) {
	valid := struct {
		ResultSet resultSet `json:"resultSet"`
	}{ResultSet: resultSet{
		Name:    "LeagueLeaders",
		Headers: []string{"PTS", leagueLeadersPlayerColumn, leagueLeadersRankColumn, leagueLeadersPlayerIDColumn, "TEAM_ID", "GP"},
		RowSet:  [][]interface{}{{float64(33.5), "Test Player", float64(1), float64(42), float64(99), float64(10)}},
	}}
	resp, err := LeagueLeaders(context.Background(), newStatsFixtureClient(t, valid), LeagueLeadersRequest{})
	if err != nil {
		t.Fatalf("LeagueLeaders() error = %v", err)
	}
	leaders := resp.Data.LeagueLeaders
	if len(leaders) != 1 || leaders[0].PlayerID != 42 || leaders[0].TeamID != 99 || leaders[0].PTS != 33.5 {
		t.Errorf("LeagueLeaders() = %+v, want name-keyed values from singular resultSet", leaders)
	}

	invalid := valid
	invalid.ResultSet.Headers = []string{"PTS", leagueLeadersRankColumn, leagueLeadersPlayerIDColumn}
	_, err = LeagueLeaders(context.Background(), newStatsFixtureClient(t, invalid), LeagueLeadersRequest{})
	if err == nil || !strings.Contains(err.Error(), `missing expected column "PLAYER"`) {
		t.Fatalf("LeagueLeaders() error = %v, want missing required-column error", err)
	}
}
