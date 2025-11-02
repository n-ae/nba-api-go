package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPtDefendRequest contains parameters for the LeagueDashPtDefend endpoint
type LeagueDashPtDefendRequest struct {
	Season          *parameters.Season
	SeasonType      *parameters.SeasonType
	PerMode         *parameters.PerMode
	LeagueID        *parameters.LeagueID
	DefenseCategory *string
}

// LeagueDashPtDefendLeagueDashPtDefend represents the LeagueDashPtDefend result set for LeagueDashPtDefend
type LeagueDashPtDefendLeagueDashPtDefend struct {
	CLOSE_DEF_PERSON_ID           string  `json:"CLOSE_DEF_PERSON_ID"`
	PLAYER_NAME                   string  `json:"PLAYER_NAME"`
	PLAYER_LAST_TEAM_ID           int     `json:"PLAYER_LAST_TEAM_ID"`
	PLAYER_LAST_TEAM_ABBREVIATION string  `json:"PLAYER_LAST_TEAM_ABBREVIATION"`
	PLAYER_POSITION               string  `json:"PLAYER_POSITION"`
	AGE                           int     `json:"AGE"`
	GP                            int     `json:"GP"`
	G                             string  `json:"G"`
	FREQ                          string  `json:"FREQ"`
	D_FGM                         int     `json:"D_FGM"`
	D_FGA                         int     `json:"D_FGA"`
	D_FG_PCT                      float64 `json:"D_FG_PCT"`
	NORMAL_FG_PCT                 float64 `json:"NORMAL_FG_PCT"`
	PCT_PLUSMINUS                 float64 `json:"PCT_PLUSMINUS"`
}

// LeagueDashPtDefendResponse contains the response data from the LeagueDashPtDefend endpoint
type LeagueDashPtDefendResponse struct {
	LeagueDashPtDefend []LeagueDashPtDefendLeagueDashPtDefend
}

// GetLeagueDashPtDefend retrieves data from the leaguedashptdefend endpoint
func GetLeagueDashPtDefend(ctx context.Context, client *stats.Client, req LeagueDashPtDefendRequest) (*models.Response[*LeagueDashPtDefendResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguedashptdefend", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPtDefendResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPtDefend = make([]LeagueDashPtDefendLeagueDashPtDefend, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := LeagueDashPtDefendLeagueDashPtDefend{
					CLOSE_DEF_PERSON_ID:           toString(row[0]),
					PLAYER_NAME:                   toString(row[1]),
					PLAYER_LAST_TEAM_ID:           toInt(row[2]),
					PLAYER_LAST_TEAM_ABBREVIATION: toString(row[3]),
					PLAYER_POSITION:               toString(row[4]),
					AGE:                           toInt(row[5]),
					GP:                            toInt(row[6]),
					G:                             toString(row[7]),
					FREQ:                          toString(row[8]),
					D_FGM:                         toInt(row[9]),
					D_FGA:                         toInt(row[10]),
					D_FG_PCT:                      toFloat(row[11]),
					NORMAL_FG_PCT:                 toFloat(row[12]),
					PCT_PLUSMINUS:                 toFloat(row[13]),
				}
				response.LeagueDashPtDefend = append(response.LeagueDashPtDefend, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
