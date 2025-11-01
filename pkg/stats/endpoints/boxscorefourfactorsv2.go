package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// BoxScoreFourFactorsV2Request contains parameters for the BoxScoreFourFactorsV2 endpoint
type BoxScoreFourFactorsV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreFourFactorsV2PlayerStats represents the PlayerStats result set for BoxScoreFourFactorsV2
type BoxScoreFourFactorsV2PlayerStats struct {
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
	EFG_PCT           float64 `json:"EFG_PCT"`
	FTA_RATE          float64 `json:"FTA_RATE"`
	TM_TOV_PCT        float64 `json:"TM_TOV_PCT"`
	OREB_PCT          float64 `json:"OREB_PCT"`
	OPP_EFG_PCT       float64 `json:"OPP_EFG_PCT"`
	OPP_FTA_RATE      float64 `json:"OPP_FTA_RATE"`
	OPP_TOV_PCT       float64 `json:"OPP_TOV_PCT"`
	OPP_OREB_PCT      float64 `json:"OPP_OREB_PCT"`
}

// BoxScoreFourFactorsV2TeamStats represents the TeamStats result set for BoxScoreFourFactorsV2
type BoxScoreFourFactorsV2TeamStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	MIN               float64 `json:"MIN"`
	EFG_PCT           float64 `json:"EFG_PCT"`
	FTA_RATE          float64 `json:"FTA_RATE"`
	TM_TOV_PCT        float64 `json:"TM_TOV_PCT"`
	OREB_PCT          float64 `json:"OREB_PCT"`
	OPP_EFG_PCT       float64 `json:"OPP_EFG_PCT"`
	OPP_FTA_RATE      float64 `json:"OPP_FTA_RATE"`
	OPP_TOV_PCT       float64 `json:"OPP_TOV_PCT"`
	OPP_OREB_PCT      float64 `json:"OPP_OREB_PCT"`
}

// BoxScoreFourFactorsV2Response contains the response data from the BoxScoreFourFactorsV2 endpoint
type BoxScoreFourFactorsV2Response struct {
	PlayerStats []BoxScoreFourFactorsV2PlayerStats
	TeamStats   []BoxScoreFourFactorsV2TeamStats
}

// GetBoxScoreFourFactorsV2 retrieves data from the boxscorefourfactorsv2 endpoint
func GetBoxScoreFourFactorsV2(ctx context.Context, client *stats.Client, req BoxScoreFourFactorsV2Request) (*models.Response[*BoxScoreFourFactorsV2Response], error) {
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
	if err := client.GetJSON(ctx, "/boxscorefourfactorsv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreFourFactorsV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreFourFactorsV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 18 {
				item := BoxScoreFourFactorsV2PlayerStats{
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
					EFG_PCT:           toFloat(row[10]),
					FTA_RATE:          toFloat(row[11]),
					TM_TOV_PCT:        toFloat(row[12]),
					OREB_PCT:          toFloat(row[13]),
					OPP_EFG_PCT:       toFloat(row[14]),
					OPP_FTA_RATE:      toFloat(row[15]),
					OPP_TOV_PCT:       toFloat(row[16]),
					OPP_OREB_PCT:      toFloat(row[17]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreFourFactorsV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 14 {
				item := BoxScoreFourFactorsV2TeamStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_NAME:         toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY:         toString(row[4]),
					MIN:               toFloat(row[5]),
					EFG_PCT:           toFloat(row[6]),
					FTA_RATE:          toFloat(row[7]),
					TM_TOV_PCT:        toFloat(row[8]),
					OREB_PCT:          toFloat(row[9]),
					OPP_EFG_PCT:       toFloat(row[10]),
					OPP_FTA_RATE:      toFloat(row[11]),
					OPP_TOV_PCT:       toFloat(row[12]),
					OPP_OREB_PCT:      toFloat(row[13]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
