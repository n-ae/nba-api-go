package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingReboundingRequest contains parameters for the PlayerTrackingRebounding endpoint
type PlayerTrackingReboundingRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingReboundingPlayerTrackingRebounding represents the PlayerTrackingRebounding result set for PlayerTrackingRebounding
type PlayerTrackingReboundingPlayerTrackingRebounding struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	OREB              float64 `json:"OREB"`
	DREB              float64 `json:"DREB"`
	REB               float64 `json:"REB"`
	OREB_CONTEST      float64 `json:"OREB_CONTEST"`
	OREB_UNCONTEST    float64 `json:"OREB_UNCONTEST"`
	OREB_CONTEST_PCT  float64 `json:"OREB_CONTEST_PCT"`
	DREB_CONTEST      float64 `json:"DREB_CONTEST"`
	DREB_UNCONTEST    float64 `json:"DREB_UNCONTEST"`
	DREB_CONTEST_PCT  float64 `json:"DREB_CONTEST_PCT"`
	REB_CONTEST       float64 `json:"REB_CONTEST"`
	REB_UNCONTEST     float64 `json:"REB_UNCONTEST"`
	REB_CONTEST_PCT   float64 `json:"REB_CONTEST_PCT"`
	AVG_REB_DIST      float64 `json:"AVG_REB_DIST"`
}

// PlayerTrackingReboundingResponse contains the response data from the PlayerTrackingRebounding endpoint
type PlayerTrackingReboundingResponse struct {
	PlayerTrackingRebounding []PlayerTrackingReboundingPlayerTrackingRebounding
}

// GetPlayerTrackingRebounding retrieves data from the playertrackingebounding endpoint
func GetPlayerTrackingRebounding(ctx context.Context, client *stats.Client, req PlayerTrackingReboundingRequest) (*models.Response[*PlayerTrackingReboundingResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingebounding", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingReboundingResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingRebounding = make([]PlayerTrackingReboundingPlayerTrackingRebounding, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 19 {
				item := PlayerTrackingReboundingPlayerTrackingRebounding{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					OREB:              toFloat(row[6]),
					DREB:              toFloat(row[7]),
					REB:               toFloat(row[8]),
					OREB_CONTEST:      toFloat(row[9]),
					OREB_UNCONTEST:    toFloat(row[10]),
					OREB_CONTEST_PCT:  toFloat(row[11]),
					DREB_CONTEST:      toFloat(row[12]),
					DREB_UNCONTEST:    toFloat(row[13]),
					DREB_CONTEST_PCT:  toFloat(row[14]),
					REB_CONTEST:       toFloat(row[15]),
					REB_UNCONTEST:     toFloat(row[16]),
					REB_CONTEST_PCT:   toFloat(row[17]),
					AVG_REB_DIST:      toFloat(row[18]),
				}
				response.PlayerTrackingRebounding = append(response.PlayerTrackingRebounding, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
