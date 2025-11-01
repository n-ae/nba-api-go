package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashTeamBioStatsRequest contains parameters for the LeagueDashTeamBioStats endpoint
type LeagueDashTeamBioStatsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueDashTeamBioStatsLeagueDashTeamBioStats represents the LeagueDashTeamBioStats result set for LeagueDashTeamBioStats
type LeagueDashTeamBioStatsLeagueDashTeamBioStats struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	W                 string  `json:"W"`
	L                 string  `json:"L"`
	W_PCT             float64 `json:"W_PCT"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
	OFF_RATING        string  `json:"OFF_RATING"`
	DEF_RATING        string  `json:"DEF_RATING"`
	NET_RATING        string  `json:"NET_RATING"`
	AST_PCT           float64 `json:"AST_PCT"`
	AST_TO            float64 `json:"AST_TO"`
	AST_RATIO         float64 `json:"AST_RATIO"`
	OREB_PCT          float64 `json:"OREB_PCT"`
	DREB_PCT          float64 `json:"DREB_PCT"`
	REB_PCT           float64 `json:"REB_PCT"`
	TM_TOV_PCT        float64 `json:"TM_TOV_PCT"`
	EFG_PCT           float64 `json:"EFG_PCT"`
	TS_PCT            float64 `json:"TS_PCT"`
	PACE              string  `json:"PACE"`
	PIE               string  `json:"PIE"`
}

// LeagueDashTeamBioStatsResponse contains the response data from the LeagueDashTeamBioStats endpoint
type LeagueDashTeamBioStatsResponse struct {
	LeagueDashTeamBioStats []LeagueDashTeamBioStatsLeagueDashTeamBioStats
}

// GetLeagueDashTeamBioStats retrieves data from the leaguedashteambiostats endpoint
func GetLeagueDashTeamBioStats(ctx context.Context, client *stats.Client, req LeagueDashTeamBioStatsRequest) (*models.Response[*LeagueDashTeamBioStatsResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguedashteambiostats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashTeamBioStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashTeamBioStats = make([]LeagueDashTeamBioStatsLeagueDashTeamBioStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 22 {
				item := LeagueDashTeamBioStatsLeagueDashTeamBioStats{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					GP:                toInt(row[3]),
					W:                 toString(row[4]),
					L:                 toString(row[5]),
					W_PCT:             toFloat(row[6]),
					PLUS_MINUS:        toFloat(row[7]),
					OFF_RATING:        toString(row[8]),
					DEF_RATING:        toString(row[9]),
					NET_RATING:        toString(row[10]),
					AST_PCT:           toFloat(row[11]),
					AST_TO:            toFloat(row[12]),
					AST_RATIO:         toFloat(row[13]),
					OREB_PCT:          toFloat(row[14]),
					DREB_PCT:          toFloat(row[15]),
					REB_PCT:           toFloat(row[16]),
					TM_TOV_PCT:        toFloat(row[17]),
					EFG_PCT:           toFloat(row[18]),
					TS_PCT:            toFloat(row[19]),
					PACE:              toString(row[20]),
					PIE:               toString(row[21]),
				}
				response.LeagueDashTeamBioStats = append(response.LeagueDashTeamBioStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
