package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamInfoCommonV2Request contains parameters for the TeamInfoCommonV2 endpoint
type TeamInfoCommonV2Request struct {
	TeamID     string
	LeagueID   *parameters.LeagueID
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
}

// TeamInfoCommonV2TeamInfoCommon represents the TeamInfoCommon result set for TeamInfoCommonV2
type TeamInfoCommonV2TeamInfoCommon struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	SEASON_YEAR       string  `json:"SEASON_YEAR"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CONFERENCE   string  `json:"TEAM_CONFERENCE"`
	TEAM_DIVISION     string  `json:"TEAM_DIVISION"`
	TEAM_CODE         string  `json:"TEAM_CODE"`
	W                 string  `json:"W"`
	L                 string  `json:"L"`
	PCT               string  `json:"PCT"`
	CONF_RANK         float64 `json:"CONF_RANK"`
	DIV_RANK          float64 `json:"DIV_RANK"`
	MIN_YEAR          float64 `json:"MIN_YEAR"`
	MAX_YEAR          string  `json:"MAX_YEAR"`
}

// TeamInfoCommonV2TeamSeasonRanks represents the TeamSeasonRanks result set for TeamInfoCommonV2
type TeamInfoCommonV2TeamSeasonRanks struct {
	LEAGUE_ID    string  `json:"LEAGUE_ID"`
	SEASON_ID    string  `json:"SEASON_ID"`
	TEAM_ID      int     `json:"TEAM_ID"`
	PTS_RANK     float64 `json:"PTS_RANK"`
	PTS_PG       float64 `json:"PTS_PG"`
	REB_RANK     float64 `json:"REB_RANK"`
	REB_PG       float64 `json:"REB_PG"`
	AST_RANK     float64 `json:"AST_RANK"`
	AST_PG       float64 `json:"AST_PG"`
	OPP_PTS_RANK float64 `json:"OPP_PTS_RANK"`
	OPP_PTS_PG   float64 `json:"OPP_PTS_PG"`
}

// TeamInfoCommonV2Response contains the response data from the TeamInfoCommonV2 endpoint
type TeamInfoCommonV2Response struct {
	TeamInfoCommon  []TeamInfoCommonV2TeamInfoCommon
	TeamSeasonRanks []TeamInfoCommonV2TeamSeasonRanks
}

// GetTeamInfoCommonV2 retrieves data from the teaminfocommonv2 endpoint
func GetTeamInfoCommonV2(ctx context.Context, client *stats.Client, req TeamInfoCommonV2Request) (*models.Response[*TeamInfoCommonV2Response], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teaminfocommonv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamInfoCommonV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamInfoCommon = make([]TeamInfoCommonV2TeamInfoCommon, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 15 {
				item := TeamInfoCommonV2TeamInfoCommon{
					TEAM_ID:           toInt(row[0]),
					SEASON_YEAR:       toString(row[1]),
					TEAM_CITY:         toString(row[2]),
					TEAM_NAME:         toString(row[3]),
					TEAM_ABBREVIATION: toString(row[4]),
					TEAM_CONFERENCE:   toString(row[5]),
					TEAM_DIVISION:     toString(row[6]),
					TEAM_CODE:         toString(row[7]),
					W:                 toString(row[8]),
					L:                 toString(row[9]),
					PCT:               toString(row[10]),
					CONF_RANK:         toFloat(row[11]),
					DIV_RANK:          toFloat(row[12]),
					MIN_YEAR:          toFloat(row[13]),
					MAX_YEAR:          toString(row[14]),
				}
				response.TeamInfoCommon = append(response.TeamInfoCommon, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamSeasonRanks = make([]TeamInfoCommonV2TeamSeasonRanks, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				item := TeamInfoCommonV2TeamSeasonRanks{
					LEAGUE_ID:    toString(row[0]),
					SEASON_ID:    toString(row[1]),
					TEAM_ID:      toInt(row[2]),
					PTS_RANK:     toFloat(row[3]),
					PTS_PG:       toFloat(row[4]),
					REB_RANK:     toFloat(row[5]),
					REB_PG:       toFloat(row[6]),
					AST_RANK:     toFloat(row[7]),
					AST_PG:       toFloat(row[8]),
					OPP_PTS_RANK: toFloat(row[9]),
					OPP_PTS_PG:   toFloat(row[10]),
				}
				response.TeamSeasonRanks = append(response.TeamSeasonRanks, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
