package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// BoxScoreAdvancedV2Request contains parameters for the BoxScoreAdvancedV2 endpoint
type BoxScoreAdvancedV2Request struct {
	GameID string
	StartPeriod *string
	EndPeriod *string
	StartRange *string
	EndRange *string
	RangeType *string
}


// BoxScoreAdvancedV2PlayerStats represents the PlayerStats result set for BoxScoreAdvancedV2
type BoxScoreAdvancedV2PlayerStats struct {
	GAME_ID string `json:"GAME_ID"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	TEAM_CITY string `json:"TEAM_CITY"`
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER_NAME string `json:"PLAYER_NAME"`
	NICKNAME string `json:"NICKNAME"`
	START_POSITION string `json:"START_POSITION"`
	COMMENT string `json:"COMMENT"`
	MIN float64 `json:"MIN"`
	E_OFF_RATING string `json:"E_OFF_RATING"`
	OFF_RATING string `json:"OFF_RATING"`
	E_DEF_RATING string `json:"E_DEF_RATING"`
	DEF_RATING string `json:"DEF_RATING"`
	E_NET_RATING string `json:"E_NET_RATING"`
	NET_RATING string `json:"NET_RATING"`
	AST_PCT float64 `json:"AST_PCT"`
	AST_TOV float64 `json:"AST_TOV"`
	AST_RATIO float64 `json:"AST_RATIO"`
	OREB_PCT float64 `json:"OREB_PCT"`
	DREB_PCT float64 `json:"DREB_PCT"`
	REB_PCT float64 `json:"REB_PCT"`
	TM_TOV_PCT float64 `json:"TM_TOV_PCT"`
	EFG_PCT float64 `json:"EFG_PCT"`
	TS_PCT float64 `json:"TS_PCT"`
	USG_PCT float64 `json:"USG_PCT"`
	E_USG_PCT float64 `json:"E_USG_PCT"`
	E_PACE string `json:"E_PACE"`
	PACE string `json:"PACE"`
	PACE_PER40 string `json:"PACE_PER40"`
	POSS string `json:"POSS"`
	PIE string `json:"PIE"`
}

// BoxScoreAdvancedV2TeamStats represents the TeamStats result set for BoxScoreAdvancedV2
type BoxScoreAdvancedV2TeamStats struct {
	GAME_ID string `json:"GAME_ID"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_NAME string `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	TEAM_CITY string `json:"TEAM_CITY"`
	MIN float64 `json:"MIN"`
	E_OFF_RATING string `json:"E_OFF_RATING"`
	OFF_RATING string `json:"OFF_RATING"`
	E_DEF_RATING string `json:"E_DEF_RATING"`
	DEF_RATING string `json:"DEF_RATING"`
	E_NET_RATING string `json:"E_NET_RATING"`
	NET_RATING string `json:"NET_RATING"`
	AST_PCT float64 `json:"AST_PCT"`
	AST_TOV float64 `json:"AST_TOV"`
	AST_RATIO float64 `json:"AST_RATIO"`
	OREB_PCT float64 `json:"OREB_PCT"`
	DREB_PCT float64 `json:"DREB_PCT"`
	REB_PCT float64 `json:"REB_PCT"`
	E_TM_TOV_PCT float64 `json:"E_TM_TOV_PCT"`
	TM_TOV_PCT float64 `json:"TM_TOV_PCT"`
	EFG_PCT float64 `json:"EFG_PCT"`
	TS_PCT float64 `json:"TS_PCT"`
	USG_PCT float64 `json:"USG_PCT"`
	E_USG_PCT float64 `json:"E_USG_PCT"`
	E_PACE string `json:"E_PACE"`
	PACE string `json:"PACE"`
	PACE_PER40 string `json:"PACE_PER40"`
	POSS string `json:"POSS"`
	PIE string `json:"PIE"`
}


// BoxScoreAdvancedV2Response contains the response data from the BoxScoreAdvancedV2 endpoint
type BoxScoreAdvancedV2Response struct {
	PlayerStats []BoxScoreAdvancedV2PlayerStats
	TeamStats []BoxScoreAdvancedV2TeamStats
}

// GetBoxScoreAdvancedV2 retrieves data from the boxscoreadvancedv2 endpoint
func GetBoxScoreAdvancedV2(ctx context.Context, client *stats.Client, req BoxScoreAdvancedV2Request) (*models.Response[*BoxScoreAdvancedV2Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.StartPeriod != nil {
		params.Set("StartPeriod", string(*req.StartPeriod))
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", string(*req.EndPeriod))
	}
	if req.StartRange != nil {
		params.Set("StartRange", string(*req.StartRange))
	}
	if req.EndRange != nil {
		params.Set("EndRange", string(*req.EndRange))
	}
	if req.RangeType != nil {
		params.Set("RangeType", string(*req.RangeType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/boxscoreadvancedv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreAdvancedV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreAdvancedV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 32 {
				item := BoxScoreAdvancedV2PlayerStats{
					GAME_ID: toString(row[0]),
					TEAM_ID: toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_CITY: toString(row[3]),
					PLAYER_ID: toInt(row[4]),
					PLAYER_NAME: toString(row[5]),
					NICKNAME: toString(row[6]),
					START_POSITION: toString(row[7]),
					COMMENT: toString(row[8]),
					MIN: toFloat(row[9]),
					E_OFF_RATING: toString(row[10]),
					OFF_RATING: toString(row[11]),
					E_DEF_RATING: toString(row[12]),
					DEF_RATING: toString(row[13]),
					E_NET_RATING: toString(row[14]),
					NET_RATING: toString(row[15]),
					AST_PCT: toFloat(row[16]),
					AST_TOV: toFloat(row[17]),
					AST_RATIO: toFloat(row[18]),
					OREB_PCT: toFloat(row[19]),
					DREB_PCT: toFloat(row[20]),
					REB_PCT: toFloat(row[21]),
					TM_TOV_PCT: toFloat(row[22]),
					EFG_PCT: toFloat(row[23]),
					TS_PCT: toFloat(row[24]),
					USG_PCT: toFloat(row[25]),
					E_USG_PCT: toFloat(row[26]),
					E_PACE: toString(row[27]),
					PACE: toString(row[28]),
					PACE_PER40: toString(row[29]),
					POSS: toString(row[30]),
					PIE: toString(row[31]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreAdvancedV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 29 {
				item := BoxScoreAdvancedV2TeamStats{
					GAME_ID: toString(row[0]),
					TEAM_ID: toInt(row[1]),
					TEAM_NAME: toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY: toString(row[4]),
					MIN: toFloat(row[5]),
					E_OFF_RATING: toString(row[6]),
					OFF_RATING: toString(row[7]),
					E_DEF_RATING: toString(row[8]),
					DEF_RATING: toString(row[9]),
					E_NET_RATING: toString(row[10]),
					NET_RATING: toString(row[11]),
					AST_PCT: toFloat(row[12]),
					AST_TOV: toFloat(row[13]),
					AST_RATIO: toFloat(row[14]),
					OREB_PCT: toFloat(row[15]),
					DREB_PCT: toFloat(row[16]),
					REB_PCT: toFloat(row[17]),
					E_TM_TOV_PCT: toFloat(row[18]),
					TM_TOV_PCT: toFloat(row[19]),
					EFG_PCT: toFloat(row[20]),
					TS_PCT: toFloat(row[21]),
					USG_PCT: toFloat(row[22]),
					E_USG_PCT: toFloat(row[23]),
					E_PACE: toString(row[24]),
					PACE: toString(row[25]),
					PACE_PER40: toString(row[26]),
					POSS: toString(row[27]),
					PIE: toString(row[28]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
