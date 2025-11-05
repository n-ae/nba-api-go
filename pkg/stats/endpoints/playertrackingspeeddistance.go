package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingSpeedDistanceRequest contains parameters for the PlayerTrackingSpeedDistance endpoint
type PlayerTrackingSpeedDistanceRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingSpeedDistancePlayerTrackingSpeedDistance represents the PlayerTrackingSpeedDistance result set for PlayerTrackingSpeedDistance
type PlayerTrackingSpeedDistancePlayerTrackingSpeedDistance struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	DIST_FEET         string  `json:"DIST_FEET"`
	DIST_MILES        string  `json:"DIST_MILES"`
	DIST_MILES_OFF    string  `json:"DIST_MILES_OFF"`
	DIST_MILES_DEF    string  `json:"DIST_MILES_DEF"`
	AVG_SPEED         string  `json:"AVG_SPEED"`
	AVG_SPEED_OFF     string  `json:"AVG_SPEED_OFF"`
	AVG_SPEED_DEF     string  `json:"AVG_SPEED_DEF"`
}

// PlayerTrackingSpeedDistanceResponse contains the response data from the PlayerTrackingSpeedDistance endpoint
type PlayerTrackingSpeedDistanceResponse struct {
	PlayerTrackingSpeedDistance []PlayerTrackingSpeedDistancePlayerTrackingSpeedDistance
}

// GetPlayerTrackingSpeedDistance retrieves data from the playertrackingspeeddistance endpoint
func GetPlayerTrackingSpeedDistance(ctx context.Context, client *stats.Client, req PlayerTrackingSpeedDistanceRequest) (*models.Response[*PlayerTrackingSpeedDistanceResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingspeeddistance", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingSpeedDistanceResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingSpeedDistance = make([]PlayerTrackingSpeedDistancePlayerTrackingSpeedDistance, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 13 {
				item := PlayerTrackingSpeedDistancePlayerTrackingSpeedDistance{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					DIST_FEET:         toString(row[6]),
					DIST_MILES:        toString(row[7]),
					DIST_MILES_OFF:    toString(row[8]),
					DIST_MILES_DEF:    toString(row[9]),
					AVG_SPEED:         toString(row[10]),
					AVG_SPEED_OFF:     toString(row[11]),
					AVG_SPEED_DEF:     toString(row[12]),
				}
				response.PlayerTrackingSpeedDistance = append(response.PlayerTrackingSpeedDistance, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
