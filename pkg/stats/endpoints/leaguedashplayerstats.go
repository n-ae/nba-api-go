package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPlayerStatsRequest contains parameters for the LeagueDashPlayerStats endpoint
type LeagueDashPlayerStatsRequest struct {
	Season *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode *parameters.PerMode
	LeagueID *parameters.LeagueID
}


// LeagueDashPlayerStatsLeagueDashPlayerStats represents the LeagueDashPlayerStats result set for LeagueDashPlayerStats
type LeagueDashPlayerStatsLeagueDashPlayerStats struct {
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER_NAME string `json:"PLAYER_NAME"`
	NICKNAME string `json:"NICKNAME"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
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
	NBA_FANTASY_PTS float64 `json:"NBA_FANTASY_PTS"`
	DD2 float64 `json:"DD2"`
	TD3 float64 `json:"TD3"`
	GP_RANK float64 `json:"GP_RANK"`
	W_RANK float64 `json:"W_RANK"`
	L_RANK float64 `json:"L_RANK"`
	W_PCT_RANK float64 `json:"W_PCT_RANK"`
	MIN_RANK float64 `json:"MIN_RANK"`
	FGM_RANK float64 `json:"FGM_RANK"`
	FGA_RANK float64 `json:"FGA_RANK"`
	FG_PCT_RANK float64 `json:"FG_PCT_RANK"`
	FG3M_RANK float64 `json:"FG3M_RANK"`
	FG3A_RANK float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK float64 `json:"FG3_PCT_RANK"`
	FTM_RANK float64 `json:"FTM_RANK"`
	FTA_RANK float64 `json:"FTA_RANK"`
	FT_PCT_RANK float64 `json:"FT_PCT_RANK"`
	OREB_RANK float64 `json:"OREB_RANK"`
	DREB_RANK float64 `json:"DREB_RANK"`
	REB_RANK float64 `json:"REB_RANK"`
	AST_RANK float64 `json:"AST_RANK"`
	TOV_RANK float64 `json:"TOV_RANK"`
	STL_RANK float64 `json:"STL_RANK"`
	BLK_RANK float64 `json:"BLK_RANK"`
	BLKA_RANK float64 `json:"BLKA_RANK"`
	PF_RANK float64 `json:"PF_RANK"`
	PFD_RANK float64 `json:"PFD_RANK"`
	PTS_RANK float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK float64 `json:"PLUS_MINUS_RANK"`
	NBA_FANTASY_PTS_RANK float64 `json:"NBA_FANTASY_PTS_RANK"`
	DD2_RANK float64 `json:"DD2_RANK"`
	TD3_RANK float64 `json:"TD3_RANK"`
	CFID string `json:"CFID"`
	CFPARAMS string `json:"CFPARAMS"`
}


// LeagueDashPlayerStatsResponse contains the response data from the LeagueDashPlayerStats endpoint
type LeagueDashPlayerStatsResponse struct {
	LeagueDashPlayerStats []LeagueDashPlayerStatsLeagueDashPlayerStats
}

// GetLeagueDashPlayerStats retrieves data from the leaguedashplayerstats endpoint
func GetLeagueDashPlayerStats(ctx context.Context, client *stats.Client, req LeagueDashPlayerStatsRequest) (*models.Response[*LeagueDashPlayerStatsResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguedashplayerstats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPlayerStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPlayerStats = make([]LeagueDashPlayerStatsLeagueDashPlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 66 {
				item := LeagueDashPlayerStatsLeagueDashPlayerStats{
					PLAYER_ID: toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					NICKNAME: toString(row[2]),
					TEAM_ID: toInt(row[3]),
					TEAM_ABBREVIATION: toString(row[4]),
					AGE: toInt(row[5]),
					GP: toInt(row[6]),
					W: toString(row[7]),
					L: toString(row[8]),
					W_PCT: toFloat(row[9]),
					MIN: toFloat(row[10]),
					FGM: toInt(row[11]),
					FGA: toInt(row[12]),
					FG_PCT: toFloat(row[13]),
					FG3M: toInt(row[14]),
					FG3A: toInt(row[15]),
					FG3_PCT: toFloat(row[16]),
					FTM: toInt(row[17]),
					FTA: toInt(row[18]),
					FT_PCT: toFloat(row[19]),
					OREB: toFloat(row[20]),
					DREB: toFloat(row[21]),
					REB: toFloat(row[22]),
					AST: toFloat(row[23]),
					TOV: toFloat(row[24]),
					STL: toFloat(row[25]),
					BLK: toFloat(row[26]),
					BLKA: toInt(row[27]),
					PF: toFloat(row[28]),
					PFD: toFloat(row[29]),
					PTS: toFloat(row[30]),
					PLUS_MINUS: toFloat(row[31]),
					NBA_FANTASY_PTS: toFloat(row[32]),
					DD2: toFloat(row[33]),
					TD3: toFloat(row[34]),
					GP_RANK: toFloat(row[35]),
					W_RANK: toFloat(row[36]),
					L_RANK: toFloat(row[37]),
					W_PCT_RANK: toFloat(row[38]),
					MIN_RANK: toFloat(row[39]),
					FGM_RANK: toFloat(row[40]),
					FGA_RANK: toFloat(row[41]),
					FG_PCT_RANK: toFloat(row[42]),
					FG3M_RANK: toFloat(row[43]),
					FG3A_RANK: toFloat(row[44]),
					FG3_PCT_RANK: toFloat(row[45]),
					FTM_RANK: toFloat(row[46]),
					FTA_RANK: toFloat(row[47]),
					FT_PCT_RANK: toFloat(row[48]),
					OREB_RANK: toFloat(row[49]),
					DREB_RANK: toFloat(row[50]),
					REB_RANK: toFloat(row[51]),
					AST_RANK: toFloat(row[52]),
					TOV_RANK: toFloat(row[53]),
					STL_RANK: toFloat(row[54]),
					BLK_RANK: toFloat(row[55]),
					BLKA_RANK: toFloat(row[56]),
					PF_RANK: toFloat(row[57]),
					PFD_RANK: toFloat(row[58]),
					PTS_RANK: toFloat(row[59]),
					PLUS_MINUS_RANK: toFloat(row[60]),
					NBA_FANTASY_PTS_RANK: toFloat(row[61]),
					DD2_RANK: toFloat(row[62]),
					TD3_RANK: toFloat(row[63]),
					CFID: toString(row[64]),
					CFPARAMS: toString(row[65]),
				}
				response.LeagueDashPlayerStats = append(response.LeagueDashPlayerStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
