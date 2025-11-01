package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamPlayerDashboardRequest contains parameters for the TeamPlayerDashboard endpoint
type TeamPlayerDashboardRequest struct {
	TeamID string
	Season *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode *parameters.PerMode
	LeagueID *parameters.LeagueID
	MeasureType *string
}


// TeamPlayerDashboardPlayersSeasonTotals represents the PlayersSeasonTotals result set for TeamPlayerDashboard
type TeamPlayerDashboardPlayersSeasonTotals struct {
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER_NAME string `json:"PLAYER_NAME"`
	NICKNAME string `json:"NICKNAME"`
	PLAYER_POSITION string `json:"PLAYER_POSITION"`
	AGE int `json:"AGE"`
	GP int `json:"GP"`
	W string `json:"W"`
	L string `json:"L"`
	W_PCT float64 `json:"W_PCT"`
	MIN float64 `json:"MIN"`
	FGM int `json:"FGM"`
	FGA int `json:"FGA"`
	FG_PCT float64 `json:"FG_PCT"`
	FG3M int `json:"FG3M"`
	FG3A int `json:"FG3A"`
	FG3_PCT float64 `json:"FG3_PCT"`
	FTM int `json:"FTM"`
	FTA int `json:"FTA"`
	FT_PCT float64 `json:"FT_PCT"`
	OREB float64 `json:"OREB"`
	DREB float64 `json:"DREB"`
	REB float64 `json:"REB"`
	AST float64 `json:"AST"`
	TOV float64 `json:"TOV"`
	STL float64 `json:"STL"`
	BLK float64 `json:"BLK"`
	BLKA int `json:"BLKA"`
	PF float64 `json:"PF"`
	PFD float64 `json:"PFD"`
	PTS float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}


// TeamPlayerDashboardResponse contains the response data from the TeamPlayerDashboard endpoint
type TeamPlayerDashboardResponse struct {
	PlayersSeasonTotals []TeamPlayerDashboardPlayersSeasonTotals
}

// GetTeamPlayerDashboard retrieves data from the teamplayerdashboard endpoint
func GetTeamPlayerDashboard(ctx context.Context, client *stats.Client, req TeamPlayerDashboardRequest) (*models.Response[*TeamPlayerDashboardResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
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
	if req.MeasureType != nil {
		params.Set("MeasureType", string(*req.MeasureType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamplayerdashboard", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamPlayerDashboardResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayersSeasonTotals = make([]TeamPlayerDashboardPlayersSeasonTotals, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 33 {
				item := TeamPlayerDashboardPlayersSeasonTotals{
					TEAM_ID: toInt(row[0]),
					TEAM_ABBREVIATION: toString(row[1]),
					PLAYER_ID: toInt(row[2]),
					PLAYER_NAME: toString(row[3]),
					NICKNAME: toString(row[4]),
					PLAYER_POSITION: toString(row[5]),
					AGE: toInt(row[6]),
					GP: toInt(row[7]),
					W: toString(row[8]),
					L: toString(row[9]),
					W_PCT: toFloat(row[10]),
					MIN: toFloat(row[11]),
					FGM: toInt(row[12]),
					FGA: toInt(row[13]),
					FG_PCT: toFloat(row[14]),
					FG3M: toInt(row[15]),
					FG3A: toInt(row[16]),
					FG3_PCT: toFloat(row[17]),
					FTM: toInt(row[18]),
					FTA: toInt(row[19]),
					FT_PCT: toFloat(row[20]),
					OREB: toFloat(row[21]),
					DREB: toFloat(row[22]),
					REB: toFloat(row[23]),
					AST: toFloat(row[24]),
					TOV: toFloat(row[25]),
					STL: toFloat(row[26]),
					BLK: toFloat(row[27]),
					BLKA: toInt(row[28]),
					PF: toFloat(row[29]),
					PFD: toFloat(row[30]),
					PTS: toFloat(row[31]),
					PLUS_MINUS: toFloat(row[32]),
				}
				response.PlayersSeasonTotals = append(response.PlayersSeasonTotals, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
