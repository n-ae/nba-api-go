package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPtTeamDefendRequest contains parameters for the LeagueDashPtTeamDefend endpoint
type LeagueDashPtTeamDefendRequest struct {
	Season          *parameters.Season
	SeasonType      *parameters.SeasonType
	PerMode         *parameters.PerMode
	LeagueID        *parameters.LeagueID
	DefenseCategory *string
}

// LeagueDashPtTeamDefendLeagueDashPtTeamDefend represents the LeagueDashPtTeamDefend result set for LeagueDashPtTeamDefend
type LeagueDashPtTeamDefendLeagueDashPtTeamDefend struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	G                 string  `json:"G"`
	FREQ              string  `json:"FREQ"`
	D_FGM             int     `json:"D_FGM"`
	D_FGA             int     `json:"D_FGA"`
	D_FG_PCT          float64 `json:"D_FG_PCT"`
	NORMAL_FG_PCT     float64 `json:"NORMAL_FG_PCT"`
	PCT_PLUSMINUS     float64 `json:"PCT_PLUSMINUS"`
}

// LeagueDashPtTeamDefendResponse contains the response data from the LeagueDashPtTeamDefend endpoint
type LeagueDashPtTeamDefendResponse struct {
	LeagueDashPtTeamDefend []LeagueDashPtTeamDefendLeagueDashPtTeamDefend
}

// GetLeagueDashPtTeamDefend retrieves data from the leaguedashptteamdefend endpoint
func GetLeagueDashPtTeamDefend(ctx context.Context, client *stats.Client, req LeagueDashPtTeamDefendRequest) (*models.Response[*LeagueDashPtTeamDefendResponse], error) {
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
	if req.DefenseCategory != nil {
		params.Set("DefenseCategory", string(*req.DefenseCategory))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "leaguedashptteamdefend", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPtTeamDefendResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPtTeamDefend = make([]LeagueDashPtTeamDefendLeagueDashPtTeamDefend, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 11 {
				item := LeagueDashPtTeamDefendLeagueDashPtTeamDefend{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					GP:                toInt(row[3]),
					G:                 toString(row[4]),
					FREQ:              toString(row[5]),
					D_FGM:             toInt(row[6]),
					D_FGA:             toInt(row[7]),
					D_FG_PCT:          toFloat(row[8]),
					NORMAL_FG_PCT:     toFloat(row[9]),
					PCT_PLUSMINUS:     toFloat(row[10]),
				}
				response.LeagueDashPtTeamDefend = append(response.LeagueDashPtTeamDefend, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
