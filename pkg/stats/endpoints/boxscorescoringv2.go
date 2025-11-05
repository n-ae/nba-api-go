package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// BoxScoreScoringV2Request contains parameters for the BoxScoreScoringV2 endpoint
type BoxScoreScoringV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreScoringV2PlayerStats represents the PlayerStats result set for BoxScoreScoringV2
type BoxScoreScoringV2PlayerStats struct {
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
	PCT_FGA_2PT       float64 `json:"PCT_FGA_2PT"`
	PCT_FGA_3PT       float64 `json:"PCT_FGA_3PT"`
	PCT_PTS_2PT       float64 `json:"PCT_PTS_2PT"`
	PCT_PTS_2PT_MR    float64 `json:"PCT_PTS_2PT_MR"`
	PCT_PTS_3PT       float64 `json:"PCT_PTS_3PT"`
	PCT_PTS_FB        float64 `json:"PCT_PTS_FB"`
	PCT_PTS_FT        float64 `json:"PCT_PTS_FT"`
	PCT_PTS_OFF_TOV   float64 `json:"PCT_PTS_OFF_TOV"`
	PCT_PTS_PAINT     float64 `json:"PCT_PTS_PAINT"`
	PCT_AST_2PM       int     `json:"PCT_AST_2PM"`
	PCT_UAST_2PM      int     `json:"PCT_UAST_2PM"`
	PCT_AST_3PM       int     `json:"PCT_AST_3PM"`
	PCT_UAST_3PM      int     `json:"PCT_UAST_3PM"`
	PCT_AST_FGM       int     `json:"PCT_AST_FGM"`
	PCT_UAST_FGM      int     `json:"PCT_UAST_FGM"`
}

// BoxScoreScoringV2TeamStats represents the TeamStats result set for BoxScoreScoringV2
type BoxScoreScoringV2TeamStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	MIN               float64 `json:"MIN"`
	PCT_FGA_2PT       float64 `json:"PCT_FGA_2PT"`
	PCT_FGA_3PT       float64 `json:"PCT_FGA_3PT"`
	PCT_PTS_2PT       float64 `json:"PCT_PTS_2PT"`
	PCT_PTS_2PT_MR    float64 `json:"PCT_PTS_2PT_MR"`
	PCT_PTS_3PT       float64 `json:"PCT_PTS_3PT"`
	PCT_PTS_FB        float64 `json:"PCT_PTS_FB"`
	PCT_PTS_FT        float64 `json:"PCT_PTS_FT"`
	PCT_PTS_OFF_TOV   float64 `json:"PCT_PTS_OFF_TOV"`
	PCT_PTS_PAINT     float64 `json:"PCT_PTS_PAINT"`
	PCT_AST_2PM       int     `json:"PCT_AST_2PM"`
	PCT_UAST_2PM      int     `json:"PCT_UAST_2PM"`
	PCT_AST_3PM       int     `json:"PCT_AST_3PM"`
	PCT_UAST_3PM      int     `json:"PCT_UAST_3PM"`
	PCT_AST_FGM       int     `json:"PCT_AST_FGM"`
	PCT_UAST_FGM      int     `json:"PCT_UAST_FGM"`
}

// BoxScoreScoringV2Response contains the response data from the BoxScoreScoringV2 endpoint
type BoxScoreScoringV2Response struct {
	PlayerStats []BoxScoreScoringV2PlayerStats
	TeamStats   []BoxScoreScoringV2TeamStats
}

// GetBoxScoreScoringV2 retrieves data from the boxscorescoringv2 endpoint
func GetBoxScoreScoringV2(ctx context.Context, client *stats.Client, req BoxScoreScoringV2Request) (*models.Response[*BoxScoreScoringV2Response], error) {
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
	if err := client.GetJSON(ctx, "boxscorescoringv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreScoringV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreScoringV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 25 {
				item := BoxScoreScoringV2PlayerStats{
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
					PCT_FGA_2PT:       toFloat(row[10]),
					PCT_FGA_3PT:       toFloat(row[11]),
					PCT_PTS_2PT:       toFloat(row[12]),
					PCT_PTS_2PT_MR:    toFloat(row[13]),
					PCT_PTS_3PT:       toFloat(row[14]),
					PCT_PTS_FB:        toFloat(row[15]),
					PCT_PTS_FT:        toFloat(row[16]),
					PCT_PTS_OFF_TOV:   toFloat(row[17]),
					PCT_PTS_PAINT:     toFloat(row[18]),
					PCT_AST_2PM:       toInt(row[19]),
					PCT_UAST_2PM:      toInt(row[20]),
					PCT_AST_3PM:       toInt(row[21]),
					PCT_UAST_3PM:      toInt(row[22]),
					PCT_AST_FGM:       toInt(row[23]),
					PCT_UAST_FGM:      toInt(row[24]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreScoringV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 21 {
				item := BoxScoreScoringV2TeamStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_NAME:         toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY:         toString(row[4]),
					MIN:               toFloat(row[5]),
					PCT_FGA_2PT:       toFloat(row[6]),
					PCT_FGA_3PT:       toFloat(row[7]),
					PCT_PTS_2PT:       toFloat(row[8]),
					PCT_PTS_2PT_MR:    toFloat(row[9]),
					PCT_PTS_3PT:       toFloat(row[10]),
					PCT_PTS_FB:        toFloat(row[11]),
					PCT_PTS_FT:        toFloat(row[12]),
					PCT_PTS_OFF_TOV:   toFloat(row[13]),
					PCT_PTS_PAINT:     toFloat(row[14]),
					PCT_AST_2PM:       toInt(row[15]),
					PCT_UAST_2PM:      toInt(row[16]),
					PCT_AST_3PM:       toInt(row[17]),
					PCT_UAST_3PM:      toInt(row[18]),
					PCT_AST_FGM:       toInt(row[19]),
					PCT_UAST_FGM:      toInt(row[20]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
