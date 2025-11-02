package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamVsTeamRequest contains parameters for the TeamVsTeam endpoint
type TeamVsTeamRequest struct {
	TeamID      string
	VsTeamID    string
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	MeasureType *string
	PerMode     *parameters.PerMode
	LeagueID    *parameters.LeagueID
}

// TeamVsTeamOverall represents the Overall result set for TeamVsTeam
type TeamVsTeamOverall struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
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

// TeamVsTeamResponse contains the response data from the TeamVsTeam endpoint
type TeamVsTeamResponse struct {
	Overall []TeamVsTeamOverall
}

// GetTeamVsTeam retrieves data from the teamvsteam endpoint
func GetTeamVsTeam(ctx context.Context, client *stats.Client, req TeamVsTeamRequest) (*models.Response[*TeamVsTeamResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.VsTeamID == "" {
		return nil, fmt.Errorf("VsTeamID is required")
	}
	params.Set("VsTeamID", string(req.VsTeamID))
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.MeasureType != nil {
		params.Set("MeasureType", string(*req.MeasureType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamvsteam", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamVsTeamResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Overall = make([]TeamVsTeamOverall, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := TeamVsTeamOverall{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					GP:         toInt(row[2]),
					W:          toString(row[3]),
					L:          toString(row[4]),
					W_PCT:      toFloat(row[5]),
					MIN:        toFloat(row[6]),
					FGM:        toInt(row[7]),
					FGA:        toInt(row[8]),
					FG_PCT:     toFloat(row[9]),
					FG3M:       toInt(row[10]),
					FG3A:       toInt(row[11]),
					FG3_PCT:    toFloat(row[12]),
					FTM:        toInt(row[13]),
					FTA:        toInt(row[14]),
					FT_PCT:     toFloat(row[15]),
					OREB:       toFloat(row[16]),
					DREB:       toFloat(row[17]),
					REB:        toFloat(row[18]),
					AST:        toFloat(row[19]),
					TOV:        toFloat(row[20]),
					STL:        toFloat(row[21]),
					BLK:        toFloat(row[22]),
					BLKA:       toInt(row[23]),
					PF:         toFloat(row[24]),
					PFD:        toFloat(row[25]),
					PTS:        toFloat(row[26]),
					PLUS_MINUS: toFloat(row[27]),
				}
				response.Overall = append(response.Overall, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
