package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// BoxScoreTraditionalV2Request contains parameters for the BoxScoreTraditionalV2 endpoint
type BoxScoreTraditionalV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreTraditionalV2PlayerStats represents the PlayerStats result set for BoxScoreTraditionalV2
type BoxScoreTraditionalV2PlayerStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	NICKNAME          string  `json:"NICKNAME"`
	START_POSITION    string  `json:"START_POSITION"`
	COMMENT           string  `json:"COMMENT"`
	MIN               float64 `json:"MIN"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	FTM               int     `json:"FTM"`
	FTA               int     `json:"FTA"`
	FT_PCT            float64 `json:"FT_PCT"`
	OREB              int     `json:"OREB"`
	DREB              int     `json:"DREB"`
	REB               int     `json:"REB"`
	AST               int     `json:"AST"`
	STL               int     `json:"STL"`
	BLK               int     `json:"BLK"`
	TO                int     `json:"TO"`
	PF                int     `json:"PF"`
	PTS               int     `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// BoxScoreTraditionalV2TeamStats represents the TeamStats result set for BoxScoreTraditionalV2
type BoxScoreTraditionalV2TeamStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	MIN               float64 `json:"MIN"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	FTM               int     `json:"FTM"`
	FTA               int     `json:"FTA"`
	FT_PCT            float64 `json:"FT_PCT"`
	OREB              int     `json:"OREB"`
	DREB              int     `json:"DREB"`
	REB               int     `json:"REB"`
	AST               int     `json:"AST"`
	STL               int     `json:"STL"`
	BLK               int     `json:"BLK"`
	TO                int     `json:"TO"`
	PF                int     `json:"PF"`
	PTS               int     `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// BoxScoreTraditionalV2TeamStarterBenchStats represents the TeamStarterBenchStats result set for BoxScoreTraditionalV2
type BoxScoreTraditionalV2TeamStarterBenchStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	STARTERS_BENCH    string  `json:"STARTERS_BENCH"`
	MIN               float64 `json:"MIN"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	FTM               int     `json:"FTM"`
	FTA               int     `json:"FTA"`
	FT_PCT            float64 `json:"FT_PCT"`
	OREB              int     `json:"OREB"`
	DREB              int     `json:"DREB"`
	REB               int     `json:"REB"`
	AST               int     `json:"AST"`
	STL               int     `json:"STL"`
	BLK               int     `json:"BLK"`
	TO                int     `json:"TO"`
	PF                int     `json:"PF"`
	PTS               int     `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// BoxScoreTraditionalV2Response contains the response data from the BoxScoreTraditionalV2 endpoint
type BoxScoreTraditionalV2Response struct {
	PlayerStats           []BoxScoreTraditionalV2PlayerStats
	TeamStats             []BoxScoreTraditionalV2TeamStats
	TeamStarterBenchStats []BoxScoreTraditionalV2TeamStarterBenchStats
}

// GetBoxScoreTraditionalV2 retrieves data from the boxscoretraditionalv2 endpoint
func GetBoxScoreTraditionalV2(ctx context.Context, client *stats.Client, req BoxScoreTraditionalV2Request) (*models.Response[*BoxScoreTraditionalV2Response], error) {
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
	if err := client.GetJSON(ctx, "/boxscoretraditionalv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreTraditionalV2Response{}

	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreTraditionalV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 29 {
				item := BoxScoreTraditionalV2PlayerStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_CITY:         toString(row[3]),
					PLAYER_ID:         toInt(row[4]),
					PLAYER_NAME:       toString(row[5]),
					NICKNAME:          toString(row[6]),
					START_POSITION:    toString(row[7]),
					COMMENT:           toString(row[8]),
					MIN:               toFloat(row[9]),
					FGM:               toInt(row[10]),
					FGA:               toInt(row[11]),
					FG_PCT:            toFloat(row[12]),
					FG3M:              toInt(row[13]),
					FG3A:              toInt(row[14]),
					FG3_PCT:           toFloat(row[15]),
					FTM:               toInt(row[16]),
					FTA:               toInt(row[17]),
					FT_PCT:            toFloat(row[18]),
					OREB:              toInt(row[19]),
					DREB:              toInt(row[20]),
					REB:               toInt(row[21]),
					AST:               toInt(row[22]),
					STL:               toInt(row[23]),
					BLK:               toInt(row[24]),
					TO:                toInt(row[25]),
					PF:                toInt(row[26]),
					PTS:               toInt(row[27]),
					PLUS_MINUS:        toFloat(row[28]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}

	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreTraditionalV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 25 {
				item := BoxScoreTraditionalV2TeamStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_NAME:         toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY:         toString(row[4]),
					MIN:               toFloat(row[5]),
					FGM:               toInt(row[6]),
					FGA:               toInt(row[7]),
					FG_PCT:            toFloat(row[8]),
					FG3M:              toInt(row[9]),
					FG3A:              toInt(row[10]),
					FG3_PCT:           toFloat(row[11]),
					FTM:               toInt(row[12]),
					FTA:               toInt(row[13]),
					FT_PCT:            toFloat(row[14]),
					OREB:              toInt(row[15]),
					DREB:              toInt(row[16]),
					REB:               toInt(row[17]),
					AST:               toInt(row[18]),
					STL:               toInt(row[19]),
					BLK:               toInt(row[20]),
					TO:                toInt(row[21]),
					PF:                toInt(row[22]),
					PTS:               toInt(row[23]),
					PLUS_MINUS:        toFloat(row[24]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	if len(rawResp.ResultSets) > 2 {
		response.TeamStarterBenchStats = make([]BoxScoreTraditionalV2TeamStarterBenchStats, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 26 {
				item := BoxScoreTraditionalV2TeamStarterBenchStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_NAME:         toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY:         toString(row[4]),
					STARTERS_BENCH:    toString(row[5]),
					MIN:               toFloat(row[6]),
					FGM:               toInt(row[7]),
					FGA:               toInt(row[8]),
					FG_PCT:            toFloat(row[9]),
					FG3M:              toInt(row[10]),
					FG3A:              toInt(row[11]),
					FG3_PCT:           toFloat(row[12]),
					FTM:               toInt(row[13]),
					FTA:               toInt(row[14]),
					FT_PCT:            toFloat(row[15]),
					OREB:              toInt(row[16]),
					DREB:              toInt(row[17]),
					REB:               toInt(row[18]),
					AST:               toInt(row[19]),
					STL:               toInt(row[20]),
					BLK:               toInt(row[21]),
					TO:                toInt(row[22]),
					PF:                toInt(row[23]),
					PTS:               toInt(row[24]),
					PLUS_MINUS:        toFloat(row[25]),
				}
				response.TeamStarterBenchStats = append(response.TeamStarterBenchStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
