package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// BoxScoreUsageV2Request contains parameters for the BoxScoreUsageV2 endpoint
type BoxScoreUsageV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreUsageV2PlayerStats represents the PlayerStats result set for BoxScoreUsageV2
type BoxScoreUsageV2PlayerStats struct {
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
	USG_PCT           float64 `json:"USG_PCT"`
	PCT_FGM           float64 `json:"PCT_FGM"`
	PCT_FGA           float64 `json:"PCT_FGA"`
	PCT_FG3M          float64 `json:"PCT_FG3M"`
	PCT_FG3A          float64 `json:"PCT_FG3A"`
	PCT_FTM           float64 `json:"PCT_FTM"`
	PCT_FTA           float64 `json:"PCT_FTA"`
	PCT_OREB          float64 `json:"PCT_OREB"`
	PCT_DREB          float64 `json:"PCT_DREB"`
	PCT_REB           float64 `json:"PCT_REB"`
	PCT_AST           float64 `json:"PCT_AST"`
	PCT_TOV           float64 `json:"PCT_TOV"`
	PCT_STL           float64 `json:"PCT_STL"`
	PCT_BLK           float64 `json:"PCT_BLK"`
	PCT_BLKA          float64 `json:"PCT_BLKA"`
	PCT_PF            float64 `json:"PCT_PF"`
	PCT_PFD           float64 `json:"PCT_PFD"`
	PCT_PTS           float64 `json:"PCT_PTS"`
}

// BoxScoreUsageV2TeamStats represents the TeamStats result set for BoxScoreUsageV2
type BoxScoreUsageV2TeamStats struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	MIN               float64 `json:"MIN"`
	USG_PCT           float64 `json:"USG_PCT"`
	PCT_FGM           float64 `json:"PCT_FGM"`
	PCT_FGA           float64 `json:"PCT_FGA"`
	PCT_FG3M          float64 `json:"PCT_FG3M"`
	PCT_FG3A          float64 `json:"PCT_FG3A"`
	PCT_FTM           float64 `json:"PCT_FTM"`
	PCT_FTA           float64 `json:"PCT_FTA"`
	PCT_OREB          float64 `json:"PCT_OREB"`
	PCT_DREB          float64 `json:"PCT_DREB"`
	PCT_REB           float64 `json:"PCT_REB"`
	PCT_AST           float64 `json:"PCT_AST"`
	PCT_TOV           float64 `json:"PCT_TOV"`
	PCT_STL           float64 `json:"PCT_STL"`
	PCT_BLK           float64 `json:"PCT_BLK"`
	PCT_BLKA          float64 `json:"PCT_BLKA"`
	PCT_PF            float64 `json:"PCT_PF"`
	PCT_PFD           float64 `json:"PCT_PFD"`
	PCT_PTS           float64 `json:"PCT_PTS"`
}

// BoxScoreUsageV2Response contains the response data from the BoxScoreUsageV2 endpoint
type BoxScoreUsageV2Response struct {
	PlayerStats []BoxScoreUsageV2PlayerStats
	TeamStats   []BoxScoreUsageV2TeamStats
}

// GetBoxScoreUsageV2 retrieves data from the boxscoreusagev2 endpoint
func GetBoxScoreUsageV2(ctx context.Context, client *stats.Client, req BoxScoreUsageV2Request) (*models.Response[*BoxScoreUsageV2Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("%s is required", "GameID")
	}
	params.Set("GameID", req.GameID)
	if req.StartPeriod != nil {
		params.Set("StartPeriod", *req.StartPeriod)
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", *req.EndPeriod)
	}
	if req.StartRange != nil {
		params.Set("StartRange", *req.StartRange)
	}
	if req.EndRange != nil {
		params.Set("EndRange", *req.EndRange)
	}
	if req.RangeType != nil {
		params.Set("RangeType", *req.RangeType)
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "boxscoreusagev2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreUsageV2Response{}
	if rs, ok := findResultSet(rawResp.ResultSets, "PlayerStats"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreUsageV2PlayerStats{})); err != nil {
			return nil, fmt.Errorf("BoxScoreUsageV2: PlayerStats result set: %w", err)
		}
		response.PlayerStats = make([]BoxScoreUsageV2PlayerStats, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 28 {
				item := BoxScoreUsageV2PlayerStats{
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
					USG_PCT:           toFloat(row[10]),
					PCT_FGM:           toFloat(row[11]),
					PCT_FGA:           toFloat(row[12]),
					PCT_FG3M:          toFloat(row[13]),
					PCT_FG3A:          toFloat(row[14]),
					PCT_FTM:           toFloat(row[15]),
					PCT_FTA:           toFloat(row[16]),
					PCT_OREB:          toFloat(row[17]),
					PCT_DREB:          toFloat(row[18]),
					PCT_REB:           toFloat(row[19]),
					PCT_AST:           toFloat(row[20]),
					PCT_TOV:           toFloat(row[21]),
					PCT_STL:           toFloat(row[22]),
					PCT_BLK:           toFloat(row[23]),
					PCT_BLKA:          toFloat(row[24]),
					PCT_PF:            toFloat(row[25]),
					PCT_PFD:           toFloat(row[26]),
					PCT_PTS:           toFloat(row[27]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "TeamStats"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreUsageV2TeamStats{})); err != nil {
			return nil, fmt.Errorf("BoxScoreUsageV2: TeamStats result set: %w", err)
		}
		response.TeamStats = make([]BoxScoreUsageV2TeamStats, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 24 {
				item := BoxScoreUsageV2TeamStats{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_NAME:         toString(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					TEAM_CITY:         toString(row[4]),
					MIN:               toFloat(row[5]),
					USG_PCT:           toFloat(row[6]),
					PCT_FGM:           toFloat(row[7]),
					PCT_FGA:           toFloat(row[8]),
					PCT_FG3M:          toFloat(row[9]),
					PCT_FG3A:          toFloat(row[10]),
					PCT_FTM:           toFloat(row[11]),
					PCT_FTA:           toFloat(row[12]),
					PCT_OREB:          toFloat(row[13]),
					PCT_DREB:          toFloat(row[14]),
					PCT_REB:           toFloat(row[15]),
					PCT_AST:           toFloat(row[16]),
					PCT_TOV:           toFloat(row[17]),
					PCT_STL:           toFloat(row[18]),
					PCT_BLK:           toFloat(row[19]),
					PCT_BLKA:          toFloat(row[20]),
					PCT_PF:            toFloat(row[21]),
					PCT_PFD:           toFloat(row[22]),
					PCT_PTS:           toFloat(row[23]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
