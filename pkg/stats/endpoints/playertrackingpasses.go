package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingPassesRequest contains parameters for the PlayerTrackingPasses endpoint
type PlayerTrackingPassesRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingPassesPlayerTrackingPasses represents the PlayerTrackingPasses result set for PlayerTrackingPasses
type PlayerTrackingPassesPlayerTrackingPasses struct {
	PLAYER_ID           int     `json:"PLAYER_ID"`
	PLAYER_NAME         string  `json:"PLAYER_NAME"`
	TEAM_ID             int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION   string  `json:"TEAM_ABBREVIATION"`
	GP                  int     `json:"GP"`
	MIN                 float64 `json:"MIN"`
	PASSES_MADE         string  `json:"PASSES_MADE"`
	PASSES_RECEIVED     string  `json:"PASSES_RECEIVED"`
	AST                 float64 `json:"AST"`
	SECONDARY_AST       float64 `json:"SECONDARY_AST"`
	POTENTIAL_AST       float64 `json:"POTENTIAL_AST"`
	AST_POINTS_CREATED  float64 `json:"AST_POINTS_CREATED"`
	AST_ADJ             float64 `json:"AST_ADJ"`
	AST_TO_PASS_PCT     float64 `json:"AST_TO_PASS_PCT"`
	AST_TO_PASS_PCT_ADJ float64 `json:"AST_TO_PASS_PCT_ADJ"`
}

// PlayerTrackingPassesResponse contains the response data from the PlayerTrackingPasses endpoint
type PlayerTrackingPassesResponse struct {
	PlayerTrackingPasses []PlayerTrackingPassesPlayerTrackingPasses
}

// GetPlayerTrackingPasses retrieves data from the playertrackingpasses endpoint
func GetPlayerTrackingPasses(ctx context.Context, client *stats.Client, req PlayerTrackingPassesRequest) (*models.Response[*PlayerTrackingPassesResponse], error) {
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
	if err := client.GetJSON(ctx, "/playertrackingpasses", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingPassesResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingPasses = make([]PlayerTrackingPassesPlayerTrackingPasses, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 15 {
				item := PlayerTrackingPassesPlayerTrackingPasses{
					PLAYER_ID:           toInt(row[0]),
					PLAYER_NAME:         toString(row[1]),
					TEAM_ID:             toInt(row[2]),
					TEAM_ABBREVIATION:   toString(row[3]),
					GP:                  toInt(row[4]),
					MIN:                 toFloat(row[5]),
					PASSES_MADE:         toString(row[6]),
					PASSES_RECEIVED:     toString(row[7]),
					AST:                 toFloat(row[8]),
					SECONDARY_AST:       toFloat(row[9]),
					POTENTIAL_AST:       toFloat(row[10]),
					AST_POINTS_CREATED:  toFloat(row[11]),
					AST_ADJ:             toFloat(row[12]),
					AST_TO_PASS_PCT:     toFloat(row[13]),
					AST_TO_PASS_PCT_ADJ: toFloat(row[14]),
				}
				response.PlayerTrackingPasses = append(response.PlayerTrackingPasses, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
