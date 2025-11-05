package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingCatchShootRequest contains parameters for the PlayerTrackingCatchShoot endpoint
type PlayerTrackingCatchShootRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingCatchShootPlayerTrackingCatchShoot represents the PlayerTrackingCatchShoot result set for PlayerTrackingCatchShoot
type PlayerTrackingCatchShootPlayerTrackingCatchShoot struct {
	PLAYER_ID           int     `json:"PLAYER_ID"`
	PLAYER_NAME         string  `json:"PLAYER_NAME"`
	TEAM_ID             int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION   string  `json:"TEAM_ABBREVIATION"`
	GP                  int     `json:"GP"`
	MIN                 float64 `json:"MIN"`
	CATCH_SHOOT_FGM     int     `json:"CATCH_SHOOT_FGM"`
	CATCH_SHOOT_FGA     int     `json:"CATCH_SHOOT_FGA"`
	CATCH_SHOOT_FG_PCT  float64 `json:"CATCH_SHOOT_FG_PCT"`
	CATCH_SHOOT_PTS     float64 `json:"CATCH_SHOOT_PTS"`
	CATCH_SHOOT_FG3M    int     `json:"CATCH_SHOOT_FG3M"`
	CATCH_SHOOT_FG3A    int     `json:"CATCH_SHOOT_FG3A"`
	CATCH_SHOOT_FG3_PCT float64 `json:"CATCH_SHOOT_FG3_PCT"`
	CATCH_SHOOT_EFG_PCT float64 `json:"CATCH_SHOOT_EFG_PCT"`
}

// PlayerTrackingCatchShootResponse contains the response data from the PlayerTrackingCatchShoot endpoint
type PlayerTrackingCatchShootResponse struct {
	PlayerTrackingCatchShoot []PlayerTrackingCatchShootPlayerTrackingCatchShoot
}

// GetPlayerTrackingCatchShoot retrieves data from the playertrackingcatchshoot endpoint
func GetPlayerTrackingCatchShoot(ctx context.Context, client *stats.Client, req PlayerTrackingCatchShootRequest) (*models.Response[*PlayerTrackingCatchShootResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingcatchshoot", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingCatchShootResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingCatchShoot = make([]PlayerTrackingCatchShootPlayerTrackingCatchShoot, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := PlayerTrackingCatchShootPlayerTrackingCatchShoot{
					PLAYER_ID:           toInt(row[0]),
					PLAYER_NAME:         toString(row[1]),
					TEAM_ID:             toInt(row[2]),
					TEAM_ABBREVIATION:   toString(row[3]),
					GP:                  toInt(row[4]),
					MIN:                 toFloat(row[5]),
					CATCH_SHOOT_FGM:     toInt(row[6]),
					CATCH_SHOOT_FGA:     toInt(row[7]),
					CATCH_SHOOT_FG_PCT:  toFloat(row[8]),
					CATCH_SHOOT_PTS:     toFloat(row[9]),
					CATCH_SHOOT_FG3M:    toInt(row[10]),
					CATCH_SHOOT_FG3A:    toInt(row[11]),
					CATCH_SHOOT_FG3_PCT: toFloat(row[12]),
					CATCH_SHOOT_EFG_PCT: toFloat(row[13]),
				}
				response.PlayerTrackingCatchShoot = append(response.PlayerTrackingCatchShoot, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
