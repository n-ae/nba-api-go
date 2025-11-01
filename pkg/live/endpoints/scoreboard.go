package endpoints

import (
	"context"

	"github.com/username/nba-api-go/pkg/live"
	"github.com/username/nba-api-go/pkg/models"
)

type TeamScore struct {
	TeamID            int    `json:"teamId"`
	TeamName          string `json:"teamName"`
	TeamCity          string `json:"teamCity"`
	TeamTricode       string `json:"teamTricode"`
	Wins              int    `json:"wins"`
	Losses            int    `json:"losses"`
	Score             int    `json:"score"`
	InBonus           *bool  `json:"inBonus"`
	TimeoutsRemaining int    `json:"timeoutsRemaining"`
	Periods           []struct {
		Period     int    `json:"period"`
		PeriodType string `json:"periodType"`
		Score      int    `json:"score"`
	} `json:"periods"`
}

type GameLeader struct {
	PersonID    int     `json:"personId"`
	Name        string  `json:"name"`
	JerseyNum   string  `json:"jerseyNum"`
	Position    string  `json:"position"`
	TeamTricode string  `json:"teamTricode"`
	PlayerSlug  *string `json:"playerSlug"`
	Points      int     `json:"points"`
	Rebounds    int     `json:"rebounds"`
	Assists     int     `json:"assists"`
}

type Game struct {
	GameID            string    `json:"gameId"`
	GameCode          string    `json:"gameCode"`
	GameStatus        int       `json:"gameStatus"`
	GameStatusText    string    `json:"gameStatusText"`
	Period            int       `json:"period"`
	GameClock         string    `json:"gameClock"`
	GameTimeUTC       string    `json:"gameTimeUTC"`
	GameEt            string    `json:"gameEt"`
	RegulationPeriods int       `json:"regulationPeriods"`
	SeriesGameNumber  string    `json:"seriesGameNumber"`
	SeriesText        string    `json:"seriesText"`
	HomeTeam          TeamScore `json:"homeTeam"`
	AwayTeam          TeamScore `json:"awayTeam"`
	GameLeaders       struct {
		HomeLeaders GameLeader `json:"homeLeaders"`
		AwayLeaders GameLeader `json:"awayLeaders"`
	} `json:"gameLeaders"`
	PBOdds *struct {
		Team      *string `json:"team"`
		Odds      float64 `json:"odds"`
		Suspended int     `json:"suspended"`
	} `json:"pbOdds"`
}

type ScoreboardResponse struct {
	Meta struct {
		Version int    `json:"version"`
		Request string `json:"request"`
		Time    string `json:"time"`
		Code    int    `json:"code"`
	} `json:"meta"`
	Scoreboard struct {
		GameDate   string `json:"gameDate"`
		LeagueID   string `json:"leagueId"`
		LeagueName string `json:"leagueName"`
		Games      []Game `json:"games"`
	} `json:"scoreboard"`
}

func Scoreboard(ctx context.Context, client *live.Client) (*models.Response[*ScoreboardResponse], error) {
	var resp ScoreboardResponse
	if err := client.GetJSON(ctx, "/scoreboard/todaysScoreboard_00.json", nil, &resp); err != nil {
		return nil, err
	}

	return models.NewResponse(&resp, 200, "", nil), nil
}

func ScoreboardByDate(ctx context.Context, client *live.Client, date string) (*models.Response[*ScoreboardResponse], error) {
	endpoint := "/scoreboard/scoreboard_" + date + ".json"

	var resp ScoreboardResponse
	if err := client.GetJSON(ctx, endpoint, nil, &resp); err != nil {
		return nil, err
	}

	return models.NewResponse(&resp, 200, "", nil), nil
}
