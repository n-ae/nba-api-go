package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamYearByYearStatsRequest contains parameters for the TeamYearByYearStats endpoint
type TeamYearByYearStatsRequest struct {
	TeamID string
	LeagueID *parameters.LeagueID
	PerMode *parameters.PerMode
	SeasonType *parameters.SeasonType
}


// TeamYearByYearStatsTeamStats represents the TeamStats result set for TeamYearByYearStats
type TeamYearByYearStatsTeamStats struct {
	TEAM_ID interface{}
	TEAM_CITY interface{}
	TEAM_NAME interface{}
	YEAR interface{}
	GP interface{}
	WINS interface{}
	LOSSES interface{}
	WIN_PCT interface{}
	CONF_RANK interface{}
	DIV_RANK interface{}
	PO_WINS interface{}
	PO_LOSSES interface{}
	CONF_COUNT interface{}
	DIV_COUNT interface{}
	NBA_FINALS_APPEARANCE interface{}
	FGM interface{}
	FGA interface{}
	FG_PCT interface{}
	FG3M interface{}
	FG3A interface{}
	FG3_PCT interface{}
	FTM interface{}
	FTA interface{}
	FT_PCT interface{}
	OREB interface{}
	DREB interface{}
	REB interface{}
	AST interface{}
	PF interface{}
	STL interface{}
	TOV interface{}
	BLK interface{}
	PTS interface{}
	PTS_RANK interface{}
}


// TeamYearByYearStatsResponse contains the response data from the TeamYearByYearStats endpoint
type TeamYearByYearStatsResponse struct {
	TeamStats []TeamYearByYearStatsTeamStats
}

// GetTeamYearByYearStats retrieves data from the teamyearbyyearstats endpoint
func GetTeamYearByYearStats(ctx context.Context, client *stats.Client, req TeamYearByYearStatsRequest) (*models.Response[*TeamYearByYearStatsResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamyearbyyearstats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamYearByYearStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamStats = make([]TeamYearByYearStatsTeamStats, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 34 {
				response.TeamStats[i] = TeamYearByYearStatsTeamStats{
					TEAM_ID: row[0],
					TEAM_CITY: row[1],
					TEAM_NAME: row[2],
					YEAR: row[3],
					GP: row[4],
					WINS: row[5],
					LOSSES: row[6],
					WIN_PCT: row[7],
					CONF_RANK: row[8],
					DIV_RANK: row[9],
					PO_WINS: row[10],
					PO_LOSSES: row[11],
					CONF_COUNT: row[12],
					DIV_COUNT: row[13],
					NBA_FINALS_APPEARANCE: row[14],
					FGM: row[15],
					FGA: row[16],
					FG_PCT: row[17],
					FG3M: row[18],
					FG3A: row[19],
					FG3_PCT: row[20],
					FTM: row[21],
					FTA: row[22],
					FT_PCT: row[23],
					OREB: row[24],
					DREB: row[25],
					REB: row[26],
					AST: row[27],
					PF: row[28],
					STL: row[29],
					TOV: row[30],
					BLK: row[31],
					PTS: row[32],
					PTS_RANK: row[33],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
