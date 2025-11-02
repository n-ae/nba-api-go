package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamPlayerOnOffDetailsRequest contains parameters for the TeamPlayerOnOffDetails endpoint
type TeamPlayerOnOffDetailsRequest struct {
	TeamID      string
	MeasureType *string
	PerMode     *parameters.PerMode
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// TeamPlayerOnOffDetailsOverallOnOffSummary represents the OverallOnOffSummary result set for TeamPlayerOnOffDetails
type TeamPlayerOnOffDetailsOverallOnOffSummary struct {
	TEAM_ID        int     `json:"TEAM_ID"`
	TEAM_NAME      string  `json:"TEAM_NAME"`
	VS_PLAYER_ID   int     `json:"VS_PLAYER_ID"`
	VS_PLAYER_NAME string  `json:"VS_PLAYER_NAME"`
	COURT_STATUS   string  `json:"COURT_STATUS"`
	GP             int     `json:"GP"`
	MIN            float64 `json:"MIN"`
	PLUS_MINUS     float64 `json:"PLUS_MINUS"`
}

// TeamPlayerOnOffDetailsPlayersOnCourtTeamPlayerOnOffDetails represents the PlayersOnCourtTeamPlayerOnOffDetails result set for TeamPlayerOnOffDetails
type TeamPlayerOnOffDetailsPlayersOnCourtTeamPlayerOnOffDetails struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	VS_PLAYER_ID      int     `json:"VS_PLAYER_ID"`
	VS_PLAYER_NAME    string  `json:"VS_PLAYER_NAME"`
	COURT_STATUS      string  `json:"COURT_STATUS"`
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// TeamPlayerOnOffDetailsResponse contains the response data from the TeamPlayerOnOffDetails endpoint
type TeamPlayerOnOffDetailsResponse struct {
	OverallOnOffSummary                  []TeamPlayerOnOffDetailsOverallOnOffSummary
	PlayersOnCourtTeamPlayerOnOffDetails []TeamPlayerOnOffDetailsPlayersOnCourtTeamPlayerOnOffDetails
}

// GetTeamPlayerOnOffDetails retrieves data from the teamplayeronoffdetails endpoint
func GetTeamPlayerOnOffDetails(ctx context.Context, client *stats.Client, req TeamPlayerOnOffDetailsRequest) (*models.Response[*TeamPlayerOnOffDetailsResponse], error) {
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
	if err := client.GetJSON(ctx, "/teamplayeronoffdetails", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamPlayerOnOffDetailsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallOnOffSummary = make([]TeamPlayerOnOffDetailsOverallOnOffSummary, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 8 {
				item := TeamPlayerOnOffDetailsOverallOnOffSummary{
					TEAM_ID:        toInt(row[0]),
					TEAM_NAME:      toString(row[1]),
					VS_PLAYER_ID:   toInt(row[2]),
					VS_PLAYER_NAME: toString(row[3]),
					COURT_STATUS:   toString(row[4]),
					GP:             toInt(row[5]),
					MIN:            toFloat(row[6]),
					PLUS_MINUS:     toFloat(row[7]),
				}
				response.OverallOnOffSummary = append(response.OverallOnOffSummary, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.PlayersOnCourtTeamPlayerOnOffDetails = make([]TeamPlayerOnOffDetailsPlayersOnCourtTeamPlayerOnOffDetails, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				item := TeamPlayerOnOffDetailsPlayersOnCourtTeamPlayerOnOffDetails{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					VS_PLAYER_ID:      toInt(row[3]),
					VS_PLAYER_NAME:    toString(row[4]),
					COURT_STATUS:      toString(row[5]),
					PLAYER_ID:         toInt(row[6]),
					PLAYER_NAME:       toString(row[7]),
					GP:                toInt(row[8]),
					MIN:               toFloat(row[9]),
					PLUS_MINUS:        toFloat(row[10]),
				}
				response.PlayersOnCourtTeamPlayerOnOffDetails = append(response.PlayersOnCourtTeamPlayerOnOffDetails, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
