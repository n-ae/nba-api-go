package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueGameLogRequest contains parameters for the LeagueGameLog endpoint
type LeagueGameLogRequest struct {
	Season       parameters.Season
	SeasonType   *parameters.SeasonType
	LeagueID     *parameters.LeagueID
	PlayerOrTeam *string
	Counter      *string
	Sorter       *string
	Direction    *string
	DateFrom     *string
	DateTo       *string
}

// LeagueGameLogLeagueGameLog represents the LeagueGameLog result set for LeagueGameLog
type LeagueGameLogLeagueGameLog struct {
	SEASON_ID         string  `json:"SEASON_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	GAME_ID           string  `json:"GAME_ID"`
	GAME_DATE         string  `json:"GAME_DATE"`
	MATCHUP           string  `json:"MATCHUP"`
	WL                string  `json:"WL"`
	MIN               float64 `json:"MIN"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	FTM               int     `json:"FTM"`
	FTA               int     `json:"FTA"`
	FT_PCT            float64 `json:"FT_PCT"`
	OREB              float64 `json:"OREB"`
	DREB              float64 `json:"DREB"`
	REB               float64 `json:"REB"`
	AST               float64 `json:"AST"`
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	TOV               float64 `json:"TOV"`
	PF                float64 `json:"PF"`
	PTS               float64 `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
	VIDEO_AVAILABLE   string  `json:"VIDEO_AVAILABLE"`
}

// LeagueGameLogResponse contains the response data from the LeagueGameLog endpoint
type LeagueGameLogResponse struct {
	LeagueGameLog []LeagueGameLogLeagueGameLog
}

// GetLeagueGameLog retrieves data from the leaguegamelog endpoint
func GetLeagueGameLog(ctx context.Context, client *stats.Client, req LeagueGameLogRequest) (*models.Response[*LeagueGameLogResponse], error) {
	params := url.Values{}
	if req.Season == "" {
		return nil, fmt.Errorf("Season is required")
	}
	params.Set("Season", string(req.Season))
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.PlayerOrTeam != nil {
		params.Set("PlayerOrTeam", string(*req.PlayerOrTeam))
	}
	if req.Counter != nil {
		params.Set("Counter", string(*req.Counter))
	}
	if req.Sorter != nil {
		params.Set("Sorter", string(*req.Sorter))
	}
	if req.Direction != nil {
		params.Set("Direction", string(*req.Direction))
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", string(*req.DateFrom))
	}
	if req.DateTo != nil {
		params.Set("DateTo", string(*req.DateTo))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "leaguegamelog", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueGameLogResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueGameLog = make([]LeagueGameLogLeagueGameLog, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 29 {
				item := LeagueGameLogLeagueGameLog{
					SEASON_ID:         toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_NAME:         toString(row[3]),
					GAME_ID:           toString(row[4]),
					GAME_DATE:         toString(row[5]),
					MATCHUP:           toString(row[6]),
					WL:                toString(row[7]),
					MIN:               toFloat(row[8]),
					FGM:               toInt(row[9]),
					FGA:               toInt(row[10]),
					FG_PCT:            toFloat(row[11]),
					FG3M:              toInt(row[12]),
					FG3A:              toInt(row[13]),
					FG3_PCT:           toFloat(row[14]),
					FTM:               toInt(row[15]),
					FTA:               toInt(row[16]),
					FT_PCT:            toFloat(row[17]),
					OREB:              toFloat(row[18]),
					DREB:              toFloat(row[19]),
					REB:               toFloat(row[20]),
					AST:               toFloat(row[21]),
					STL:               toFloat(row[22]),
					BLK:               toFloat(row[23]),
					TOV:               toFloat(row[24]),
					PF:                toFloat(row[25]),
					PTS:               toFloat(row[26]),
					PLUS_MINUS:        toFloat(row[27]),
					VIDEO_AVAILABLE:   toString(row[28]),
				}
				response.LeagueGameLog = append(response.LeagueGameLog, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
