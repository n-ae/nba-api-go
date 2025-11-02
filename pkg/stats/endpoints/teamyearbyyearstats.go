package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamYearByYearStatsRequest contains parameters for the TeamYearByYearStats endpoint
type TeamYearByYearStatsRequest struct {
	TeamID     string
	LeagueID   *parameters.LeagueID
	PerMode    *parameters.PerMode
	SeasonType *parameters.SeasonType
}

// TeamYearByYearStatsTeamStats represents the TeamStats result set for TeamYearByYearStats
type TeamYearByYearStatsTeamStats struct {
	TEAM_ID               int     `json:"TEAM_ID"`
	TEAM_CITY             string  `json:"TEAM_CITY"`
	TEAM_NAME             string  `json:"TEAM_NAME"`
	YEAR                  string  `json:"YEAR"`
	GP                    int     `json:"GP"`
	WINS                  int     `json:"WINS"`
	LOSSES                int     `json:"LOSSES"`
	WIN_PCT               float64 `json:"WIN_PCT"`
	CONF_RANK             int     `json:"CONF_RANK"`
	DIV_RANK              int     `json:"DIV_RANK"`
	PO_WINS               int     `json:"PO_WINS"`
	PO_LOSSES             int     `json:"PO_LOSSES"`
	CONF_COUNT            int     `json:"CONF_COUNT"`
	DIV_COUNT             int     `json:"DIV_COUNT"`
	NBA_FINALS_APPEARANCE string  `json:"NBA_FINALS_APPEARANCE"`
	FGM                   int     `json:"FGM"`
	FGA                   int     `json:"FGA"`
	FG_PCT                float64 `json:"FG_PCT"`
	FG3M                  int     `json:"FG3M"`
	FG3A                  int     `json:"FG3A"`
	FG3_PCT               float64 `json:"FG3_PCT"`
	FTM                   int     `json:"FTM"`
	FTA                   int     `json:"FTA"`
	FT_PCT                float64 `json:"FT_PCT"`
	OREB                  int     `json:"OREB"`
	DREB                  int     `json:"DREB"`
	REB                   int     `json:"REB"`
	AST                   int     `json:"AST"`
	PF                    int     `json:"PF"`
	STL                   int     `json:"STL"`
	TOV                   int     `json:"TOV"`
	BLK                   int     `json:"BLK"`
	PTS                   int     `json:"PTS"`
	PTS_RANK              int     `json:"PTS_RANK"`
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
		response.TeamStats = make([]TeamYearByYearStatsTeamStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 34 {
				item := TeamYearByYearStatsTeamStats{
					TEAM_ID:               toInt(row[0]),
					TEAM_CITY:             toString(row[1]),
					TEAM_NAME:             toString(row[2]),
					YEAR:                  toString(row[3]),
					GP:                    toInt(row[4]),
					WINS:                  toInt(row[5]),
					LOSSES:                toInt(row[6]),
					WIN_PCT:               toFloat(row[7]),
					CONF_RANK:             toInt(row[8]),
					DIV_RANK:              toInt(row[9]),
					PO_WINS:               toInt(row[10]),
					PO_LOSSES:             toInt(row[11]),
					CONF_COUNT:            toInt(row[12]),
					DIV_COUNT:             toInt(row[13]),
					NBA_FINALS_APPEARANCE: toString(row[14]),
					FGM:                   toInt(row[15]),
					FGA:                   toInt(row[16]),
					FG_PCT:                toFloat(row[17]),
					FG3M:                  toInt(row[18]),
					FG3A:                  toInt(row[19]),
					FG3_PCT:               toFloat(row[20]),
					FTM:                   toInt(row[21]),
					FTA:                   toInt(row[22]),
					FT_PCT:                toFloat(row[23]),
					OREB:                  toInt(row[24]),
					DREB:                  toInt(row[25]),
					REB:                   toInt(row[26]),
					AST:                   toInt(row[27]),
					PF:                    toInt(row[28]),
					STL:                   toInt(row[29]),
					TOV:                   toInt(row[30]),
					BLK:                   toInt(row[31]),
					PTS:                   toInt(row[32]),
					PTS_RANK:              toInt(row[33]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
