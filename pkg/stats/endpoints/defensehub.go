package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// DefenseHubRequest contains parameters for the DefenseHub endpoint
type DefenseHubRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// DefenseHubDefenseHub represents the DefenseHub result set for DefenseHub
type DefenseHubDefenseHub struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	DREB              float64 `json:"DREB"`
	DEF_RIM_FGM       int     `json:"DEF_RIM_FGM"`
	DEF_RIM_FGA       int     `json:"DEF_RIM_FGA"`
	DEF_RIM_FG_PCT    float64 `json:"DEF_RIM_FG_PCT"`
}

// DefenseHubResponse contains the response data from the DefenseHub endpoint
type DefenseHubResponse struct {
	DefenseHub []DefenseHubDefenseHub
}

// GetDefenseHub retrieves data from the defensehub endpoint
func GetDefenseHub(ctx context.Context, client *stats.Client, req DefenseHubRequest) (*models.Response[*DefenseHubResponse], error) {
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
	if err := client.GetJSON(ctx, "/defensehub", params, &rawResp); err != nil {
		return nil, err
	}

	response := &DefenseHubResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.DefenseHub = make([]DefenseHubDefenseHub, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				item := DefenseHubDefenseHub{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					STL:               toFloat(row[6]),
					BLK:               toFloat(row[7]),
					DREB:              toFloat(row[8]),
					DEF_RIM_FGM:       toInt(row[9]),
					DEF_RIM_FGA:       toInt(row[10]),
					DEF_RIM_FG_PCT:    toFloat(row[11]),
				}
				response.DefenseHub = append(response.DefenseHub, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
