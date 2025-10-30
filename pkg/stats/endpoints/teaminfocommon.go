package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamInfoCommonRequest contains parameters for the TeamInfoCommon endpoint
type TeamInfoCommonRequest struct {
	TeamID string
	LeagueID *parameters.LeagueID
	SeasonType *parameters.SeasonType
}


// TeamInfoCommonTeamInfoCommon represents the TeamInfoCommon result set for TeamInfoCommon
type TeamInfoCommonTeamInfoCommon struct {
	TEAM_ID interface{}
	SEASON_YEAR interface{}
	TEAM_CITY interface{}
	TEAM_NAME interface{}
	TEAM_ABBREVIATION interface{}
	TEAM_CONFERENCE interface{}
	TEAM_DIVISION interface{}
	W interface{}
	L interface{}
	PCT interface{}
	MIN_YEAR interface{}
	MAX_YEAR interface{}
}

// TeamInfoCommonTeamSeasonRanks represents the TeamSeasonRanks result set for TeamInfoCommon
type TeamInfoCommonTeamSeasonRanks struct {
	LEAGUE_ID interface{}
	SEASON_ID interface{}
	TEAM_ID interface{}
	PTS_RANK interface{}
	PTS_PG interface{}
	REB_RANK interface{}
	REB_PG interface{}
	AST_RANK interface{}
	AST_PG interface{}
	OPP_PTS_RANK interface{}
	OPP_PTS_PG interface{}
}


// TeamInfoCommonResponse contains the response data from the TeamInfoCommon endpoint
type TeamInfoCommonResponse struct {
	TeamInfoCommon []TeamInfoCommonTeamInfoCommon
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
	if err := client.GetJSON(ctx, "/teaminfocommon", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamInfoCommonResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamInfoCommon = make([]TeamInfoCommonTeamInfoCommon, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 12 {
				response.TeamInfoCommon[i] = TeamInfoCommonTeamInfoCommon{
					TEAM_ID: row[0],
					SEASON_YEAR: row[1],
					TEAM_CITY: row[2],
					TEAM_NAME: row[3],
					TEAM_ABBREVIATION: row[4],
					TEAM_CONFERENCE: row[5],
					TEAM_DIVISION: row[6],
					W: row[7],
					L: row[8],
					PCT: row[9],
					MIN_YEAR: row[10],
					MAX_YEAR: row[11],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamSeasonRanks = make([]TeamInfoCommonTeamSeasonRanks, len(rawResp.ResultSets[1].RowSet))
		for i, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				response.TeamSeasonRanks[i] = TeamInfoCommonTeamSeasonRanks{
					LEAGUE_ID: row[0],
					SEASON_ID: row[1],
					TEAM_ID: row[2],
					PTS_RANK: row[3],
					PTS_PG: row[4],
					REB_RANK: row[5],
					REB_PG: row[6],
					AST_RANK: row[7],
					AST_PG: row[8],
					OPP_PTS_RANK: row[9],
					OPP_PTS_PG: row[10],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
