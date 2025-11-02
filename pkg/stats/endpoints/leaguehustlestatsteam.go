package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueHustleStatsTeamRequest contains parameters for the LeagueHustleStatsTeam endpoint
type LeagueHustleStatsTeamRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueHustleStatsTeamHustleStatsTeam represents the HustleStatsTeam result set for LeagueHustleStatsTeam
type LeagueHustleStatsTeamHustleStatsTeam struct {
	TEAM_ID                   int     `json:"TEAM_ID"`
	TEAM_NAME                 string  `json:"TEAM_NAME"`
	GP                        int     `json:"GP"`
	W                         string  `json:"W"`
	L                         string  `json:"L"`
	W_PCT                     float64 `json:"W_PCT"`
	MIN                       float64 `json:"MIN"`
	CONTESTED_SHOTS           string  `json:"CONTESTED_SHOTS"`
	CONTESTED_SHOTS_2PT       string  `json:"CONTESTED_SHOTS_2PT"`
	CONTESTED_SHOTS_3PT       string  `json:"CONTESTED_SHOTS_3PT"`
	DEFLECTIONS               string  `json:"DEFLECTIONS"`
	CHARGES_DRAWN             string  `json:"CHARGES_DRAWN"`
	SCREEN_ASSISTS            string  `json:"SCREEN_ASSISTS"`
	SCREEN_AST_PTS            float64 `json:"SCREEN_AST_PTS"`
	OFF_LOOSE_BALLS_RECOVERED string  `json:"OFF_LOOSE_BALLS_RECOVERED"`
	DEF_LOOSE_BALLS_RECOVERED string  `json:"DEF_LOOSE_BALLS_RECOVERED"`
	LOOSE_BALLS_RECOVERED     string  `json:"LOOSE_BALLS_RECOVERED"`
	OFF_BOXOUTS               string  `json:"OFF_BOXOUTS"`
	DEF_BOXOUTS               string  `json:"DEF_BOXOUTS"`
	BOX_OUT_TEAM_REBS         float64 `json:"BOX_OUT_TEAM_REBS"`
	BOX_OUTS                  string  `json:"BOX_OUTS"`
}

// LeagueHustleStatsTeamResponse contains the response data from the LeagueHustleStatsTeam endpoint
type LeagueHustleStatsTeamResponse struct {
	HustleStatsTeam []LeagueHustleStatsTeamHustleStatsTeam
}

// GetLeagueHustleStatsTeam retrieves data from the leaguehustlestats team endpoint
func GetLeagueHustleStatsTeam(ctx context.Context, client *stats.Client, req LeagueHustleStatsTeamRequest) (*models.Response[*LeagueHustleStatsTeamResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguehustlestats team", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueHustleStatsTeamResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.HustleStatsTeam = make([]LeagueHustleStatsTeamHustleStatsTeam, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 21 {
				item := LeagueHustleStatsTeamHustleStatsTeam{
					TEAM_ID:                   toInt(row[0]),
					TEAM_NAME:                 toString(row[1]),
					GP:                        toInt(row[2]),
					W:                         toString(row[3]),
					L:                         toString(row[4]),
					W_PCT:                     toFloat(row[5]),
					MIN:                       toFloat(row[6]),
					CONTESTED_SHOTS:           toString(row[7]),
					CONTESTED_SHOTS_2PT:       toString(row[8]),
					CONTESTED_SHOTS_3PT:       toString(row[9]),
					DEFLECTIONS:               toString(row[10]),
					CHARGES_DRAWN:             toString(row[11]),
					SCREEN_ASSISTS:            toString(row[12]),
					SCREEN_AST_PTS:            toFloat(row[13]),
					OFF_LOOSE_BALLS_RECOVERED: toString(row[14]),
					DEF_LOOSE_BALLS_RECOVERED: toString(row[15]),
					LOOSE_BALLS_RECOVERED:     toString(row[16]),
					OFF_BOXOUTS:               toString(row[17]),
					DEF_BOXOUTS:               toString(row[18]),
					BOX_OUT_TEAM_REBS:         toFloat(row[19]),
					BOX_OUTS:                  toString(row[20]),
				}
				response.HustleStatsTeam = append(response.HustleStatsTeam, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
