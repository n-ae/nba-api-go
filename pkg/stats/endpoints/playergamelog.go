package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

type GameLog struct {
	SeasonID       string  `json:"SEASON_ID"`
	PlayerID       int     `json:"Player_ID"`
	GameID         string  `json:"Game_ID"`
	GameDate       string  `json:"GAME_DATE"`
	Matchup        string  `json:"MATCHUP"`
	WL             string  `json:"WL"`
	MIN            int     `json:"MIN"`
	FGM            int     `json:"FGM"`
	FGA            int     `json:"FGA"`
	FGPct          float64 `json:"FG_PCT"`
	FG3M           int     `json:"FG3M"`
	FG3A           int     `json:"FG3A"`
	FG3Pct         float64 `json:"FG3_PCT"`
	FTM            int     `json:"FTM"`
	FTA            int     `json:"FTA"`
	FTPct          float64 `json:"FT_PCT"`
	OREB           int     `json:"OREB"`
	DREB           int     `json:"DREB"`
	REB            int     `json:"REB"`
	AST            int     `json:"AST"`
	STL            int     `json:"STL"`
	BLK            int     `json:"BLK"`
	TOV            int     `json:"TOV"`
	PF             int     `json:"PF"`
	PTS            int     `json:"PTS"`
	PlusMinus      int     `json:"PLUS_MINUS"`
	VideoAvailable int     `json:"VIDEO_AVAILABLE"`
}

type PlayerGameLogResponse struct {
	PlayerGameLog []GameLog `json:"PlayerGameLog"`
}

type PlayerGameLogRequest struct {
	PlayerID   string
	Season     parameters.Season
	SeasonType parameters.SeasonType
	DateFrom   string
	DateTo     string
	LeagueID   parameters.LeagueID
}

func PlayerGameLog(ctx context.Context, client *stats.Client, req PlayerGameLogRequest) (*models.Response[*PlayerGameLogResponse], error) {
	if req.PlayerID == "" {
		return nil, fmt.Errorf("%w: PlayerID is required", models.ErrInvalidRequest)
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
	params.Set("PlayerID", req.PlayerID)
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
	if err := client.GetJSON(ctx, "/playergamelog", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerGameLogResponse{}
	for _, resultSet := range rawResp.ResultSets {
		if resultSet.Name == "PlayerGameLog" {
			response.PlayerGameLog = parseGameLogs(resultSet.RowSet)
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}

func parseGameLogs(rows [][]interface{}) []GameLog {
	logs := make([]GameLog, 0, len(rows))
	for _, row := range rows {
		if len(row) < 27 {
			continue
		}
		log := GameLog{
			SeasonID:       toString(row[0]),
			PlayerID:       toInt(row[1]),
			GameID:         toString(row[2]),
			GameDate:       toString(row[3]),
			Matchup:        toString(row[4]),
			WL:             toString(row[5]),
			MIN:            toInt(row[6]),
			FGM:            toInt(row[7]),
			FGA:            toInt(row[8]),
			FGPct:          toFloat(row[9]),
			FG3M:           toInt(row[10]),
			FG3A:           toInt(row[11]),
			FG3Pct:         toFloat(row[12]),
			FTM:            toInt(row[13]),
			FTA:            toInt(row[14]),
			FTPct:          toFloat(row[15]),
			OREB:           toInt(row[16]),
			DREB:           toInt(row[17]),
			REB:            toInt(row[18]),
			AST:            toInt(row[19]),
			STL:            toInt(row[20]),
			BLK:            toInt(row[21]),
			TOV:            toInt(row[22]),
			PF:             toInt(row[23]),
			PTS:            toInt(row[24]),
			PlusMinus:      toInt(row[25]),
			VideoAvailable: toInt(row[26]),
		}
		logs = append(logs, log)
	}
	return logs
}
