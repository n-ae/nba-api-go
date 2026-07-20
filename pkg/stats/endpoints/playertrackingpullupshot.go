package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
	"github.com/n-ae/nba-api-go/v2/pkg/stats"
	"github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
)

// PlayerTrackingPullUpShotRequest contains parameters for the PlayerTrackingPullUpShot endpoint
type PlayerTrackingPullUpShotRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingPullUpShotPlayerTrackingPullUpShot represents the PlayerTrackingPullUpShot result set for PlayerTrackingPullUpShot
type PlayerTrackingPullUpShotPlayerTrackingPullUpShot struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	PULL_UP_FGM       int     `json:"PULL_UP_FGM"`
	PULL_UP_FGA       int     `json:"PULL_UP_FGA"`
	PULL_UP_FG_PCT    float64 `json:"PULL_UP_FG_PCT"`
	PULL_UP_PTS       float64 `json:"PULL_UP_PTS"`
	PULL_UP_FG3M      int     `json:"PULL_UP_FG3M"`
	PULL_UP_FG3A      int     `json:"PULL_UP_FG3A"`
	PULL_UP_FG3_PCT   float64 `json:"PULL_UP_FG3_PCT"`
	PULL_UP_EFG_PCT   float64 `json:"PULL_UP_EFG_PCT"`
}

// PlayerTrackingPullUpShotResponse contains the response data from the PlayerTrackingPullUpShot endpoint
type PlayerTrackingPullUpShotResponse struct {
	PlayerTrackingPullUpShot []PlayerTrackingPullUpShotPlayerTrackingPullUpShot
}

// GetPlayerTrackingPullUpShot retrieves data from the playertrackingpullupshot endpoint
func GetPlayerTrackingPullUpShot(ctx context.Context, client *stats.Client, req PlayerTrackingPullUpShotRequest) (*models.Response[*PlayerTrackingPullUpShotResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingpullupshot", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingPullUpShotResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "PlayerTrackingPullUpShot"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayerTrackingPullUpShotPlayerTrackingPullUpShot{})); err != nil {
			return nil, fmt.Errorf("PlayerTrackingPullUpShot: PlayerTrackingPullUpShot result set: %w", err)
		}
		response.PlayerTrackingPullUpShot = make([]PlayerTrackingPullUpShotPlayerTrackingPullUpShot, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 14 {
				item := PlayerTrackingPullUpShotPlayerTrackingPullUpShot{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					PULL_UP_FGM:       toInt(row[6]),
					PULL_UP_FGA:       toInt(row[7]),
					PULL_UP_FG_PCT:    toFloat(row[8]),
					PULL_UP_PTS:       toFloat(row[9]),
					PULL_UP_FG3M:      toInt(row[10]),
					PULL_UP_FG3A:      toInt(row[11]),
					PULL_UP_FG3_PCT:   toFloat(row[12]),
					PULL_UP_EFG_PCT:   toFloat(row[13]),
				}
				response.PlayerTrackingPullUpShot = append(response.PlayerTrackingPullUpShot, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
