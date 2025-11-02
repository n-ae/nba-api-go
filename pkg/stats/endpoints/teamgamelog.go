package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

type TeamGameLog struct {
	TeamID   int     `json:"Team_ID"`
	GameID   string  `json:"Game_ID"`
	GameDate string  `json:"GAME_DATE"`
	Matchup  string  `json:"MATCHUP"`
	WL       string  `json:"WL"`
	W        int     `json:"W"`
	L        int     `json:"L"`
	WPct     float64 `json:"W_PCT"`
	MIN      int     `json:"MIN"`
	FGM      int     `json:"FGM"`
	FGA      int     `json:"FGA"`
	FGPct    float64 `json:"FG_PCT"`
	FG3M     int     `json:"FG3M"`
	FG3A     int     `json:"FG3A"`
	FG3Pct   float64 `json:"FG3_PCT"`
	FTM      int     `json:"FTM"`
	FTA      int     `json:"FTA"`
	FTPct    float64 `json:"FT_PCT"`
	OREB     int     `json:"OREB"`
	DREB     int     `json:"DREB"`
	REB      int     `json:"REB"`
	AST      int     `json:"AST"`
	STL      int     `json:"STL"`
	BLK      int     `json:"BLK"`
	TOV      int     `json:"TOV"`
	PF       int     `json:"PF"`
	PTS      int     `json:"PTS"`
}

type TeamGameLogResponse struct {
	TeamGameLog []TeamGameLog `json:"TeamGameLog"`
}

type TeamGameLogRequest struct {
	TeamID     string
	Season     parameters.Season
	SeasonType parameters.SeasonType
	DateFrom   string
	DateTo     string
	LeagueID   parameters.LeagueID
}

func GetTeamGameLog(ctx context.Context, client *stats.Client, req TeamGameLogRequest) (*models.Response[*TeamGameLogResponse], error) {
	if req.TeamID == "" {
		return nil, fmt.Errorf("%w: TeamID is required", models.ErrInvalidRequest)
	}

	if req.SeasonType != "" {
		if err := req.SeasonType.Validate(); err != nil {
			return nil, fmt.Errorf("%w: %v", models.ErrInvalidRequest, err)
		}
	} else {
		req.SeasonType = parameters.SeasonTypeRegular
	}

	if err := req.LeagueID.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrInvalidRequest, err)
	}

	params := url.Values{}
	params.Set("TeamID", req.TeamID)
	if req.Season != "" {
		params.Set("Season", req.Season.String())
	}
	params.Set("SeasonType", req.SeasonType.String())
	if req.DateFrom != "" {
		params.Set("DateFrom", req.DateFrom)
	}
	if req.DateTo != "" {
		params.Set("DateTo", req.DateTo)
	}
	if req.LeagueID != "" {
		params.Set("LeagueID", req.LeagueID.String())
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamgamelog", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamGameLogResponse{}
	for _, resultSet := range rawResp.ResultSets {
		if resultSet.Name == "TeamGameLog" {
			response.TeamGameLog = parseTeamGameLogs(resultSet.RowSet)
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}

func parseTeamGameLogs(rows [][]interface{}) []TeamGameLog {
	logs := make([]TeamGameLog, 0, len(rows))
	for _, row := range rows {
		if len(row) < 27 {
			continue
		}
		log := TeamGameLog{
			TeamID:   toInt(row[0]),
			GameID:   toString(row[1]),
			GameDate: toString(row[2]),
			Matchup:  toString(row[3]),
			WL:       toString(row[4]),
			W:        toInt(row[5]),
			L:        toInt(row[6]),
			WPct:     toFloat(row[7]),
			MIN:      toInt(row[8]),
			FGM:      toInt(row[9]),
			FGA:      toInt(row[10]),
			FGPct:    toFloat(row[11]),
			FG3M:     toInt(row[12]),
			FG3A:     toInt(row[13]),
			FG3Pct:   toFloat(row[14]),
			FTM:      toInt(row[15]),
			FTA:      toInt(row[16]),
			FTPct:    toFloat(row[17]),
			OREB:     toInt(row[18]),
			DREB:     toInt(row[19]),
			REB:      toInt(row[20]),
			AST:      toInt(row[21]),
			STL:      toInt(row[22]),
			BLK:      toInt(row[23]),
			TOV:      toInt(row[24]),
			PF:       toInt(row[25]),
			PTS:      toInt(row[26]),
		}
		logs = append(logs, log)
	}
	return logs
}
