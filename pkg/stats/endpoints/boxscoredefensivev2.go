package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// BoxScoreDefensiveV2Request contains parameters for the BoxScoreDefensiveV2 endpoint
type BoxScoreDefensiveV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreDefensiveV2PlayerStats represents the PlayerStats result set for BoxScoreDefensiveV2
type BoxScoreDefensiveV2PlayerStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	NICKNAME          string  `json:"NICKNAME"`
	START_POSITION    string  `json:"START_POSITION"`
	COMMENT           string  `json:"COMMENT"`
	MIN               float64 `json:"MIN"`
	DEF_RIM_FGM       int     `json:"DEF_RIM_FGM"`
	DEF_RIM_FGA       int     `json:"DEF_RIM_FGA"`
	DEF_RIM_FG_PCT    float64 `json:"DEF_RIM_FG_PCT"`
}

// BoxScoreDefensiveV2TeamStats represents the TeamStats result set for BoxScoreDefensiveV2
type BoxScoreDefensiveV2TeamStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	MIN               float64 `json:"MIN"`
	DEF_RIM_FGM       int     `json:"DEF_RIM_FGM"`
	DEF_RIM_FGA       int     `json:"DEF_RIM_FGA"`
	DEF_RIM_FG_PCT    float64 `json:"DEF_RIM_FG_PCT"`
}

// BoxScoreDefensiveV2Response contains the response data from the BoxScoreDefensiveV2 endpoint
type BoxScoreDefensiveV2Response struct {
	PlayerStats []BoxScoreDefensiveV2PlayerStats
	TeamStats   []BoxScoreDefensiveV2TeamStats
}

// GetBoxScoreDefensiveV2 retrieves data from the boxscoredefensivev2 endpoint
func GetBoxScoreDefensiveV2(ctx context.Context, client *stats.Client, req BoxScoreDefensiveV2Request) (*models.Response[*BoxScoreDefensiveV2Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.StartPeriod != nil {
		params.Set("StartPeriod", string(*req.StartPeriod))
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", string(*req.EndPeriod))
	}
	if req.StartRange != nil {
		params.Set("StartRange", string(*req.StartRange))
	}
	if req.EndRange != nil {
		params.Set("EndRange", string(*req.EndRange))
	}
	if req.RangeType != nil {
		params.Set("RangeType", string(*req.RangeType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/boxscoredefensivev2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreDefensiveV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreDefensiveV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 13 {
				item := BoxScoreDefensiveV2PlayerStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_CITY:         toString(row[3]),
					PLAYER_ID:         toInt(row[4]),
					PLAYER_NAME:       toString(row[5]),
					NICKNAME:          toString(row[6]),
					START_POSITION:    toString(row[7]),
					COMMENT:           toString(row[8]),
					MIN:               toFloat(row[9]),
					DEF_RIM_FGM:       toInt(row[10]),
					DEF_RIM_FGA:       toInt(row[11]),
					DEF_RIM_FG_PCT:    toFloat(row[12]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreDefensiveV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 9 {
				item := BoxScoreDefensiveV2TeamStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_NAME:         toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY:         toString(row[4]),
					MIN:               toFloat(row[5]),
					DEF_RIM_FGM:       toInt(row[6]),
					DEF_RIM_FGA:       toInt(row[7]),
					DEF_RIM_FG_PCT:    toFloat(row[8]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
