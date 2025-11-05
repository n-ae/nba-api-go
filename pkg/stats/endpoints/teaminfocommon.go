package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamInfoCommonRequest contains parameters for the TeamInfoCommon endpoint
type TeamInfoCommonRequest struct {
	TeamID     string
	LeagueID   *parameters.LeagueID
	SeasonType *parameters.SeasonType
}

// TeamInfoCommonTeamInfoCommon represents the TeamInfoCommon result set for TeamInfoCommon
type TeamInfoCommonTeamInfoCommon struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	SEASON_YEAR       string  `json:"SEASON_YEAR"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CONFERENCE   string  `json:"TEAM_CONFERENCE"`
	TEAM_DIVISION     string  `json:"TEAM_DIVISION"`
	W                 int     `json:"W"`
	L                 int     `json:"L"`
	PCT               float64 `json:"PCT"`
	MIN_YEAR          string  `json:"MIN_YEAR"`
	MAX_YEAR          string  `json:"MAX_YEAR"`
}

// TeamInfoCommonTeamSeasonRanks represents the TeamSeasonRanks result set for TeamInfoCommon
type TeamInfoCommonTeamSeasonRanks struct {
	LEAGUE_ID    string  `json:"LEAGUE_ID"`
	SEASON_ID    string  `json:"SEASON_ID"`
	TEAM_ID      int     `json:"TEAM_ID"`
	PTS_RANK     int     `json:"PTS_RANK"`
	PTS_PG       float64 `json:"PTS_PG"`
	REB_RANK     int     `json:"REB_RANK"`
	REB_PG       float64 `json:"REB_PG"`
	AST_RANK     int     `json:"AST_RANK"`
	AST_PG       float64 `json:"AST_PG"`
	OPP_PTS_RANK int     `json:"OPP_PTS_RANK"`
	OPP_PTS_PG   float64 `json:"OPP_PTS_PG"`
}

// TeamInfoCommonResponse contains the response data from the TeamInfoCommon endpoint
type TeamInfoCommonResponse struct {
	TeamInfoCommon  []TeamInfoCommonTeamInfoCommon
	TeamSeasonRanks []TeamInfoCommonTeamSeasonRanks
}

// GetTeamInfoCommon retrieves data from the teaminfocommon endpoint
func GetTeamInfoCommon(ctx context.Context, client *stats.Client, req TeamInfoCommonRequest) (*models.Response[*TeamInfoCommonResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "teaminfocommon", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamInfoCommonResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamInfoCommon = make([]TeamInfoCommonTeamInfoCommon, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				item := TeamInfoCommonTeamInfoCommon{
					TEAM_ID:           toInt(row[0]),
					SEASON_YEAR:       toString(row[1]),
					TEAM_CITY:         toString(row[2]),
					TEAM_NAME:         toString(row[3]),
					TEAM_ABBREVIATION: toString(row[4]),
					TEAM_CONFERENCE:   toString(row[5]),
					TEAM_DIVISION:     toString(row[6]),
					W:                 toInt(row[7]),
					L:                 toInt(row[8]),
					PCT:               toFloat(row[9]),
					MIN_YEAR:          toString(row[10]),
					MAX_YEAR:          toString(row[11]),
				}
				response.TeamInfoCommon = append(response.TeamInfoCommon, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamSeasonRanks = make([]TeamInfoCommonTeamSeasonRanks, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				item := TeamInfoCommonTeamSeasonRanks{
					LEAGUE_ID:    toString(row[0]),
					SEASON_ID:    toString(row[1]),
					TEAM_ID:      toInt(row[2]),
					PTS_RANK:     toInt(row[3]),
					PTS_PG:       toFloat(row[4]),
					REB_RANK:     toInt(row[5]),
					REB_PG:       toFloat(row[6]),
					AST_RANK:     toInt(row[7]),
					AST_PG:       toFloat(row[8]),
					OPP_PTS_RANK: toInt(row[9]),
					OPP_PTS_PG:   toFloat(row[10]),
				}
				response.TeamSeasonRanks = append(response.TeamSeasonRanks, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
