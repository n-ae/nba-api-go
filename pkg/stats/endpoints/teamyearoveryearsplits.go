package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// TeamYearOverYearSplitsRequest contains parameters for the TeamYearOverYearSplits endpoint
type TeamYearOverYearSplitsRequest struct {
	TeamID      string
	MeasureType *string
	PerMode     *parameters.PerMode
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// TeamYearOverYearSplitsByYearTeamDashboard represents the ByYearTeamDashboard result set for TeamYearOverYearSplits
type TeamYearOverYearSplitsByYearTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	YEAR       string  `json:"YEAR"`
	GP         int     `json:"GP"`
	W          string  `json:"W"`
	L          string  `json:"L"`
	W_PCT      float64 `json:"W_PCT"`
	MIN        float64 `json:"MIN"`
	FGM        int     `json:"FGM"`
	FGA        int     `json:"FGA"`
	FG_PCT     float64 `json:"FG_PCT"`
	FG3M       int     `json:"FG3M"`
	FG3A       int     `json:"FG3A"`
	FG3_PCT    float64 `json:"FG3_PCT"`
	FTM        int     `json:"FTM"`
	FTA        int     `json:"FTA"`
	FT_PCT     float64 `json:"FT_PCT"`
	OREB       float64 `json:"OREB"`
	DREB       float64 `json:"DREB"`
	REB        float64 `json:"REB"`
	AST        float64 `json:"AST"`
	TOV        float64 `json:"TOV"`
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	BLKA       int     `json:"BLKA"`
	PF         float64 `json:"PF"`
	PFD        float64 `json:"PFD"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// TeamYearOverYearSplitsResponse contains the response data from the TeamYearOverYearSplits endpoint
type TeamYearOverYearSplitsResponse struct {
	ByYearTeamDashboard []TeamYearOverYearSplitsByYearTeamDashboard
}

// GetTeamYearOverYearSplits retrieves data from the teamdashboardbyyearoveryearsplits endpoint
func GetTeamYearOverYearSplits(ctx context.Context, client *stats.Client, req TeamYearOverYearSplitsRequest) (*models.Response[*TeamYearOverYearSplitsResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("%s is required", "TeamID")
	}
	params.Set("TeamID", req.TeamID)
	if req.MeasureType != nil {
		params.Set("MeasureType", *req.MeasureType)
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "teamdashboardbyyearoveryearsplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamYearOverYearSplitsResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "ByYearTeamDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamYearOverYearSplitsByYearTeamDashboard{})); err != nil {
			return nil, fmt.Errorf("TeamYearOverYearSplits: ByYearTeamDashboard result set: %w", err)
		}
		response.ByYearTeamDashboard = make([]TeamYearOverYearSplitsByYearTeamDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 29 {
				item := TeamYearOverYearSplitsByYearTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					YEAR:       toString(row[2]),
					GP:         toInt(row[3]),
					W:          toString(row[4]),
					L:          toString(row[5]),
					W_PCT:      toFloat(row[6]),
					MIN:        toFloat(row[7]),
					FGM:        toInt(row[8]),
					FGA:        toInt(row[9]),
					FG_PCT:     toFloat(row[10]),
					FG3M:       toInt(row[11]),
					FG3A:       toInt(row[12]),
					FG3_PCT:    toFloat(row[13]),
					FTM:        toInt(row[14]),
					FTA:        toInt(row[15]),
					FT_PCT:     toFloat(row[16]),
					OREB:       toFloat(row[17]),
					DREB:       toFloat(row[18]),
					REB:        toFloat(row[19]),
					AST:        toFloat(row[20]),
					TOV:        toFloat(row[21]),
					STL:        toFloat(row[22]),
					BLK:        toFloat(row[23]),
					BLKA:       toInt(row[24]),
					PF:         toFloat(row[25]),
					PFD:        toFloat(row[26]),
					PTS:        toFloat(row[27]),
					PLUS_MINUS: toFloat(row[28]),
				}
				response.ByYearTeamDashboard = append(response.ByYearTeamDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
