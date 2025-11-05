package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// ShootingEfficiencyRequest contains parameters for the ShootingEfficiency endpoint
type ShootingEfficiencyRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// ShootingEfficiencyShootingEfficiency represents the ShootingEfficiency result set for ShootingEfficiency
type ShootingEfficiencyShootingEfficiency struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	PTS               float64 `json:"PTS"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG2_PCT           float64 `json:"FG2_PCT"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	EFG_PCT           float64 `json:"EFG_PCT"`
	FT_PCT            float64 `json:"FT_PCT"`
	TS_PCT            float64 `json:"TS_PCT"`
}

// ShootingEfficiencyResponse contains the response data from the ShootingEfficiency endpoint
type ShootingEfficiencyResponse struct {
	ShootingEfficiency []ShootingEfficiencyShootingEfficiency
}

// GetShootingEfficiency retrieves data from the shootingefficiency endpoint
func GetShootingEfficiency(ctx context.Context, client *stats.Client, req ShootingEfficiencyRequest) (*models.Response[*ShootingEfficiencyResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "shootingefficiency", params, &rawResp); err != nil {
		return nil, err
	}

	response := &ShootingEfficiencyResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.ShootingEfficiency = make([]ShootingEfficiencyShootingEfficiency, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 13 {
				item := ShootingEfficiencyShootingEfficiency{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					PTS:               toFloat(row[6]),
					FG_PCT:            toFloat(row[7]),
					FG2_PCT:           toFloat(row[8]),
					FG3_PCT:           toFloat(row[9]),
					EFG_PCT:           toFloat(row[10]),
					FT_PCT:            toFloat(row[11]),
					TS_PCT:            toFloat(row[12]),
				}
				response.ShootingEfficiency = append(response.ShootingEfficiency, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
