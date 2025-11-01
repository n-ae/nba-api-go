package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingDefenseRequest contains parameters for the PlayerTrackingDefense endpoint
type PlayerTrackingDefenseRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingDefensePlayerTrackingDefense represents the PlayerTrackingDefense result set for PlayerTrackingDefense
type PlayerTrackingDefensePlayerTrackingDefense struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	DEF_RIM_FGM       int     `json:"DEF_RIM_FGM"`
	DEF_RIM_FGA       int     `json:"DEF_RIM_FGA"`
	DEF_RIM_FG_PCT    float64 `json:"DEF_RIM_FG_PCT"`
	DREB              float64 `json:"DREB"`
}

// PlayerTrackingDefenseResponse contains the response data from the PlayerTrackingDefense endpoint
type PlayerTrackingDefenseResponse struct {
	PlayerTrackingDefense []PlayerTrackingDefensePlayerTrackingDefense
}

// GetPlayerTrackingDefense retrieves data from the playertrackingdefense endpoint
func GetPlayerTrackingDefense(ctx context.Context, client *stats.Client, req PlayerTrackingDefenseRequest) (*models.Response[*PlayerTrackingDefenseResponse], error) {
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
	if err := client.GetJSON(ctx, "/playertrackingdefense", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingDefenseResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingDefense = make([]PlayerTrackingDefensePlayerTrackingDefense, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 10 {
				item := PlayerTrackingDefensePlayerTrackingDefense{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					DEF_RIM_FGM:       toInt(row[6]),
					DEF_RIM_FGA:       toInt(row[7]),
					DEF_RIM_FG_PCT:    toFloat(row[8]),
					DREB:              toFloat(row[9]),
				}
				response.PlayerTrackingDefense = append(response.PlayerTrackingDefense, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
