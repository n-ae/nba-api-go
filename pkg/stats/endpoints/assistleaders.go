package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
	"github.com/n-ae/nba-api-go/v2/pkg/stats"
	"github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
)

// AssistLeadersRequest contains parameters for the AssistLeaders endpoint
type AssistLeadersRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// AssistLeadersAssistLeaders represents the AssistLeaders result set for AssistLeaders
type AssistLeadersAssistLeaders struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	AST               float64 `json:"AST"`
}

// AssistLeadersResponse contains the response data from the AssistLeaders endpoint
type AssistLeadersResponse struct {
	AssistLeaders []AssistLeadersAssistLeaders
}

// GetAssistLeaders retrieves data from the assistleaders endpoint
func GetAssistLeaders(ctx context.Context, client *stats.Client, req AssistLeadersRequest) (*models.Response[*AssistLeadersResponse], error) {
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
	if err := client.GetJSON(ctx, "assistleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &AssistLeadersResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "AssistLeaders"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(AssistLeadersAssistLeaders{})); err != nil {
			return nil, fmt.Errorf("AssistLeaders: AssistLeaders result set: %w", err)
		}
		response.AssistLeaders = make([]AssistLeadersAssistLeaders, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 7 {
				item := AssistLeadersAssistLeaders{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					AST:               toFloat(row[6]),
				}
				response.AssistLeaders = append(response.AssistLeaders, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
