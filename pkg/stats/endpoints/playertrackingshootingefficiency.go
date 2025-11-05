package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingShootingEfficiencyRequest contains parameters for the PlayerTrackingShootingEfficiency endpoint
type PlayerTrackingShootingEfficiencyRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingShootingEfficiencyPlayerTrackingShootingEfficiency represents the PlayerTrackingShootingEfficiency result set for PlayerTrackingShootingEfficiency
type PlayerTrackingShootingEfficiencyPlayerTrackingShootingEfficiency struct {
	PLAYER_ID          int     `json:"PLAYER_ID"`
	PLAYER_NAME        string  `json:"PLAYER_NAME"`
	TEAM_ID            int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION  string  `json:"TEAM_ABBREVIATION"`
	GP                 int     `json:"GP"`
	MIN                float64 `json:"MIN"`
	DRIVE_PTS          float64 `json:"DRIVE_PTS"`
	DRIVE_FG_PCT       float64 `json:"DRIVE_FG_PCT"`
	CATCH_SHOOT_PTS    float64 `json:"CATCH_SHOOT_PTS"`
	CATCH_SHOOT_FG_PCT float64 `json:"CATCH_SHOOT_FG_PCT"`
	PULL_UP_PTS        float64 `json:"PULL_UP_PTS"`
	PULL_UP_FG_PCT     float64 `json:"PULL_UP_FG_PCT"`
}

// PlayerTrackingShootingEfficiencyResponse contains the response data from the PlayerTrackingShootingEfficiency endpoint
type PlayerTrackingShootingEfficiencyResponse struct {
	PlayerTrackingShootingEfficiency []PlayerTrackingShootingEfficiencyPlayerTrackingShootingEfficiency
}

// GetPlayerTrackingShootingEfficiency retrieves data from the playertrackingshootingefficiency endpoint
func GetPlayerTrackingShootingEfficiency(ctx context.Context, client *stats.Client, req PlayerTrackingShootingEfficiencyRequest) (*models.Response[*PlayerTrackingShootingEfficiencyResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingshootingefficiency", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingShootingEfficiencyResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingShootingEfficiency = make([]PlayerTrackingShootingEfficiencyPlayerTrackingShootingEfficiency, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				item := PlayerTrackingShootingEfficiencyPlayerTrackingShootingEfficiency{
					PLAYER_ID:          toInt(row[0]),
					PLAYER_NAME:        toString(row[1]),
					TEAM_ID:            toInt(row[2]),
					TEAM_ABBREVIATION:  toString(row[3]),
					GP:                 toInt(row[4]),
					MIN:                toFloat(row[5]),
					DRIVE_PTS:          toFloat(row[6]),
					DRIVE_FG_PCT:       toFloat(row[7]),
					CATCH_SHOOT_PTS:    toFloat(row[8]),
					CATCH_SHOOT_FG_PCT: toFloat(row[9]),
					PULL_UP_PTS:        toFloat(row[10]),
					PULL_UP_FG_PCT:     toFloat(row[11]),
				}
				response.PlayerTrackingShootingEfficiency = append(response.PlayerTrackingShootingEfficiency, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
