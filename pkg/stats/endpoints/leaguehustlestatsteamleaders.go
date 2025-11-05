package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueHustleStatsTeamLeadersRequest contains parameters for the LeagueHustleStatsTeamLeaders endpoint
type LeagueHustleStatsTeamLeadersRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueHustleStatsTeamLeadersHustleStatsTeamLeaders represents the HustleStatsTeamLeaders result set for LeagueHustleStatsTeamLeaders
type LeagueHustleStatsTeamLeadersHustleStatsTeamLeaders struct {
	TEAM_ID                   int     `json:"TEAM_ID"`
	TEAM_NAME                 string  `json:"TEAM_NAME"`
	GP                        int     `json:"GP"`
	MIN                       float64 `json:"MIN"`
	SCREEN_ASSISTS            string  `json:"SCREEN_ASSISTS"`
	SCREEN_AST_PTS            float64 `json:"SCREEN_AST_PTS"`
	OFF_LOOSE_BALLS_RECOVERED string  `json:"OFF_LOOSE_BALLS_RECOVERED"`
	DEF_LOOSE_BALLS_RECOVERED string  `json:"DEF_LOOSE_BALLS_RECOVERED"`
	LOOSE_BALLS_RECOVERED     string  `json:"LOOSE_BALLS_RECOVERED"`
	OFF_BOXOUTS               string  `json:"OFF_BOXOUTS"`
	DEF_BOXOUTS               string  `json:"DEF_BOXOUTS"`
	BOX_OUTS                  string  `json:"BOX_OUTS"`
	CONTESTED_SHOTS           string  `json:"CONTESTED_SHOTS"`
	CONTESTED_SHOTS_2PT       string  `json:"CONTESTED_SHOTS_2PT"`
	CONTESTED_SHOTS_3PT       string  `json:"CONTESTED_SHOTS_3PT"`
	CHARGES_DRAWN             string  `json:"CHARGES_DRAWN"`
	DEFLECTIONS               string  `json:"DEFLECTIONS"`
}

// LeagueHustleStatsTeamLeadersResponse contains the response data from the LeagueHustleStatsTeamLeaders endpoint
type LeagueHustleStatsTeamLeadersResponse struct {
	HustleStatsTeamLeaders []LeagueHustleStatsTeamLeadersHustleStatsTeamLeaders
}

// GetLeagueHustleStatsTeamLeaders retrieves data from the leaguehustlestatsTeamleaders endpoint
func GetLeagueHustleStatsTeamLeaders(ctx context.Context, client *stats.Client, req LeagueHustleStatsTeamLeadersRequest) (*models.Response[*LeagueHustleStatsTeamLeadersResponse], error) {
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
	if err := client.GetJSON(ctx, "leaguehustlestatsTeamleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueHustleStatsTeamLeadersResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.HustleStatsTeamLeaders = make([]LeagueHustleStatsTeamLeadersHustleStatsTeamLeaders, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 17 {
				item := LeagueHustleStatsTeamLeadersHustleStatsTeamLeaders{
					TEAM_ID:                   toInt(row[0]),
					TEAM_NAME:                 toString(row[1]),
					GP:                        toInt(row[2]),
					MIN:                       toFloat(row[3]),
					SCREEN_ASSISTS:            toString(row[4]),
					SCREEN_AST_PTS:            toFloat(row[5]),
					OFF_LOOSE_BALLS_RECOVERED: toString(row[6]),
					DEF_LOOSE_BALLS_RECOVERED: toString(row[7]),
					LOOSE_BALLS_RECOVERED:     toString(row[8]),
					OFF_BOXOUTS:               toString(row[9]),
					DEF_BOXOUTS:               toString(row[10]),
					BOX_OUTS:                  toString(row[11]),
					CONTESTED_SHOTS:           toString(row[12]),
					CONTESTED_SHOTS_2PT:       toString(row[13]),
					CONTESTED_SHOTS_3PT:       toString(row[14]),
					CHARGES_DRAWN:             toString(row[15]),
					DEFLECTIONS:               toString(row[16]),
				}
				response.HustleStatsTeamLeaders = append(response.HustleStatsTeamLeaders, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
