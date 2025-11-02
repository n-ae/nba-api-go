package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// AssistTrackerRequest contains parameters for the AssistTracker endpoint
type AssistTrackerRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// AssistTrackerAssistTracker represents the AssistTracker result set for AssistTracker
type AssistTrackerAssistTracker struct {
	PLAYER_ID                int     `json:"PLAYER_ID"`
	PLAYER_NAME              string  `json:"PLAYER_NAME"`
	TEAM_ID                  int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION        string  `json:"TEAM_ABBREVIATION"`
	GP                       int     `json:"GP"`
	W                        string  `json:"W"`
	L                        string  `json:"L"`
	W_PCT                    float64 `json:"W_PCT"`
	MIN                      float64 `json:"MIN"`
	AST                      float64 `json:"AST"`
	PASS_TO                  string  `json:"PASS_TO"`
	AST_PTS_CREATED          float64 `json:"AST_PTS_CREATED"`
	AST_PTS_CREATED_PER_PASS float64 `json:"AST_PTS_CREATED_PER_PASS"`
	AST_PCT                  float64 `json:"AST_PCT"`
	AST_ADJ                  float64 `json:"AST_ADJ"`
}

// AssistTrackerResponse contains the response data from the AssistTracker endpoint
type AssistTrackerResponse struct {
	AssistTracker []AssistTrackerAssistTracker
}

// GetAssistTracker retrieves data from the assisttracker endpoint
func GetAssistTracker(ctx context.Context, client *stats.Client, req AssistTrackerRequest) (*models.Response[*AssistTrackerResponse], error) {
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
	if err := client.GetJSON(ctx, "/assisttracker", params, &rawResp); err != nil {
		return nil, err
	}

	response := &AssistTrackerResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.AssistTracker = make([]AssistTrackerAssistTracker, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 15 {
				item := AssistTrackerAssistTracker{
					PLAYER_ID:                toInt(row[0]),
					PLAYER_NAME:              toString(row[1]),
					TEAM_ID:                  toInt(row[2]),
					TEAM_ABBREVIATION:        toString(row[3]),
					GP:                       toInt(row[4]),
					W:                        toString(row[5]),
					L:                        toString(row[6]),
					W_PCT:                    toFloat(row[7]),
					MIN:                      toFloat(row[8]),
					AST:                      toFloat(row[9]),
					PASS_TO:                  toString(row[10]),
					AST_PTS_CREATED:          toFloat(row[11]),
					AST_PTS_CREATED_PER_PASS: toFloat(row[12]),
					AST_PCT:                  toFloat(row[13]),
					AST_ADJ:                  toFloat(row[14]),
				}
				response.AssistTracker = append(response.AssistTracker, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
