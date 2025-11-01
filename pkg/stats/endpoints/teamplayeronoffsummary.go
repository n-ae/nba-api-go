package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamPlayerOnOffSummaryRequest contains parameters for the TeamPlayerOnOffSummary endpoint
type TeamPlayerOnOffSummaryRequest struct {
	TeamID      string
	MeasureType *string
	PerMode     *parameters.PerMode
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// TeamPlayerOnOffSummaryTeamPlayerOnOffSummary represents the TeamPlayerOnOffSummary result set for TeamPlayerOnOffSummary
type TeamPlayerOnOffSummaryTeamPlayerOnOffSummary struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	VS_PLAYER_ID      int     `json:"VS_PLAYER_ID"`
	VS_PLAYER_NAME    string  `json:"VS_PLAYER_NAME"`
	COURT_STATUS      string  `json:"COURT_STATUS"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
	NET_RATING        string  `json:"NET_RATING"`
	OFF_RATING        string  `json:"OFF_RATING"`
	DEF_RATING        string  `json:"DEF_RATING"`
}

// TeamPlayerOnOffSummaryResponse contains the response data from the TeamPlayerOnOffSummary endpoint
type TeamPlayerOnOffSummaryResponse struct {
	TeamPlayerOnOffSummary []TeamPlayerOnOffSummaryTeamPlayerOnOffSummary
}

// GetTeamPlayerOnOffSummary retrieves data from the teamplayeronoffsummary endpoint
func GetTeamPlayerOnOffSummary(ctx context.Context, client *stats.Client, req TeamPlayerOnOffSummaryRequest) (*models.Response[*TeamPlayerOnOffSummaryResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.MeasureType != nil {
		params.Set("MeasureType", string(*req.MeasureType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamplayeronoffsummary", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamPlayerOnOffSummaryResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamPlayerOnOffSummary = make([]TeamPlayerOnOffSummaryTeamPlayerOnOffSummary, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				item := TeamPlayerOnOffSummaryTeamPlayerOnOffSummary{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					VS_PLAYER_ID:      toInt(row[3]),
					VS_PLAYER_NAME:    toString(row[4]),
					COURT_STATUS:      toString(row[5]),
					GP:                toInt(row[6]),
					MIN:               toFloat(row[7]),
					PLUS_MINUS:        toFloat(row[8]),
					NET_RATING:        toString(row[9]),
					OFF_RATING:        toString(row[10]),
					DEF_RATING:        toString(row[11]),
				}
				response.TeamPlayerOnOffSummary = append(response.TeamPlayerOnOffSummary, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
