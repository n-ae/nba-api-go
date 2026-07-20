package endpoints

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n-ae/nba-api-go/v2/pkg/stats"
)

// TestGetGameRotation_ColumnOffsets regression-tests the 12-column
// gamerotation response (with TEAM_CITY at index 2) parsing correctly
// instead of the previous 11-column assumption, which shifted every
// field from TEAM_NAME onward and silently dropped USG_PCT.
func TestGetGameRotation_ColumnOffsets(t *testing.T) {
	const responseBody = `{
		"resultSets": [
			{
				"name": "AwayTeam",
				"headers": ["GAME_ID", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "PERSON_ID", "PLAYER_FIRST", "PLAYER_LAST", "IN_TIME_REAL", "OUT_TIME_REAL", "PLAYER_PTS", "PT_DIFF", "USG_PCT"],
				"rowSet": [
					["0022300001", 1610612745, "Houston", "Rockets", 201142, "Kevin", "Durant", "0", "5830", 5, -3, 24.6],
					["0022300001", 1610612745, "Houston", "Rockets"]
				]
			},
			{
				"name": "HomeTeam",
				"headers": ["GAME_ID", "TEAM_ID", "TEAM_CITY", "TEAM_NAME", "PERSON_ID", "PLAYER_FIRST", "PLAYER_LAST", "IN_TIME_REAL", "OUT_TIME_REAL", "PLAYER_PTS", "PT_DIFF", "USG_PCT"],
				"rowSet": [
					["0022300001", 1610612744, "Golden State", "Warriors", 201939, "Stephen", "Curry", "120", "7200", 12, 5, 31.2]
				]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(responseBody))
	}))
	defer server.Close()

	client := stats.NewClient(stats.Config{BaseURL: server.URL})

	resp, err := GetGameRotation(context.Background(), client, GameRotationRequest{GameID: "0022300001"})
	if err != nil {
		t.Fatalf("GetGameRotation() unexpected error: %v", err)
	}

	if len(resp.Data.AwayTeam) != 1 {
		t.Fatalf("AwayTeam: got %d rows, want 1 (short row must be skipped)", len(resp.Data.AwayTeam))
	}
	away := resp.Data.AwayTeam[0]

	wantAway := GameRotationAwayTeam{
		GAME_ID:       "0022300001",
		TEAM_ID:       1610612745,
		TEAM_CITY:     "Houston",
		TEAM_NAME:     "Rockets",
		PERSON_ID:     "201142",
		PLAYER_FIRST:  "Kevin",
		PLAYER_LAST:   "Durant",
		IN_TIME_REAL:  "0",
		OUT_TIME_REAL: "5830",
		PLAYER_PTS:    5,
		PT_DIFF:       -3,
		USG_PCT:       24.6,
	}
	if away != wantAway {
		t.Errorf("AwayTeam[0] = %+v, want %+v", away, wantAway)
	}

	if len(resp.Data.HomeTeam) != 1 {
		t.Fatalf("HomeTeam: got %d rows, want 1", len(resp.Data.HomeTeam))
	}
	home := resp.Data.HomeTeam[0]

	wantHome := GameRotationHomeTeam{
		GAME_ID:       "0022300001",
		TEAM_ID:       1610612744,
		TEAM_CITY:     "Golden State",
		TEAM_NAME:     "Warriors",
		PERSON_ID:     "201939",
		PLAYER_FIRST:  "Stephen",
		PLAYER_LAST:   "Curry",
		IN_TIME_REAL:  "120",
		OUT_TIME_REAL: "7200",
		PLAYER_PTS:    12,
		PT_DIFF:       5,
		USG_PCT:       31.2,
	}
	if home != wantHome {
		t.Errorf("HomeTeam[0] = %+v, want %+v", home, wantHome)
	}
}

func TestGetGameRotation_RequiresGameID(t *testing.T) {
	client := stats.NewClient(stats.Config{})
	if _, err := GetGameRotation(context.Background(), client, GameRotationRequest{}); err == nil {
		t.Error("GetGameRotation() with empty GameID: got nil error, want error")
	}
}
