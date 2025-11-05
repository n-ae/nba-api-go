package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// BoxScoreMiscV2Request contains parameters for the BoxScoreMiscV2 endpoint
type BoxScoreMiscV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreMiscV2PlayerStats represents the PlayerStats result set for BoxScoreMiscV2
type BoxScoreMiscV2PlayerStats struct {
	GAME_ID            string  `json:"GAME_ID"`
	TEAM_ID            int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION  string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY          string  `json:"TEAM_CITY"`
	PLAYER_ID          int     `json:"PLAYER_ID"`
	PLAYER_NAME        string  `json:"PLAYER_NAME"`
	NICKNAME           string  `json:"NICKNAME"`
	START_POSITION     string  `json:"START_POSITION"`
	COMMENT            string  `json:"COMMENT"`
	MIN                float64 `json:"MIN"`
	PTS_OFF_TOV        float64 `json:"PTS_OFF_TOV"`
	PTS_2ND_CHANCE     float64 `json:"PTS_2ND_CHANCE"`
	PTS_FB             float64 `json:"PTS_FB"`
	PTS_PAINT          float64 `json:"PTS_PAINT"`
	OPP_PTS_OFF_TOV    float64 `json:"OPP_PTS_OFF_TOV"`
	OPP_PTS_2ND_CHANCE float64 `json:"OPP_PTS_2ND_CHANCE"`
	OPP_PTS_FB         float64 `json:"OPP_PTS_FB"`
	OPP_PTS_PAINT      float64 `json:"OPP_PTS_PAINT"`
	BLK                float64 `json:"BLK"`
	BLKA               int     `json:"BLKA"`
	PF                 float64 `json:"PF"`
	PFD                float64 `json:"PFD"`
}

// BoxScoreMiscV2TeamStats represents the TeamStats result set for BoxScoreMiscV2
type BoxScoreMiscV2TeamStats struct {
	GAME_ID            string  `json:"GAME_ID"`
	TEAM_ID            int     `json:"TEAM_ID"`
	TEAM_NAME          string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION  string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY          string  `json:"TEAM_CITY"`
	MIN                float64 `json:"MIN"`
	PTS_OFF_TOV        float64 `json:"PTS_OFF_TOV"`
	PTS_2ND_CHANCE     float64 `json:"PTS_2ND_CHANCE"`
	PTS_FB             float64 `json:"PTS_FB"`
	PTS_PAINT          float64 `json:"PTS_PAINT"`
	OPP_PTS_OFF_TOV    float64 `json:"OPP_PTS_OFF_TOV"`
	OPP_PTS_2ND_CHANCE float64 `json:"OPP_PTS_2ND_CHANCE"`
	OPP_PTS_FB         float64 `json:"OPP_PTS_FB"`
	OPP_PTS_PAINT      float64 `json:"OPP_PTS_PAINT"`
	BLK                float64 `json:"BLK"`
	BLKA               int     `json:"BLKA"`
	PF                 float64 `json:"PF"`
	PFD                float64 `json:"PFD"`
}

// BoxScoreMiscV2Response contains the response data from the BoxScoreMiscV2 endpoint
type BoxScoreMiscV2Response struct {
	PlayerStats []BoxScoreMiscV2PlayerStats
	TeamStats   []BoxScoreMiscV2TeamStats
}

// GetBoxScoreMiscV2 retrieves data from the boxscoremiscv2 endpoint
func GetBoxScoreMiscV2(ctx context.Context, client *stats.Client, req BoxScoreMiscV2Request) (*models.Response[*BoxScoreMiscV2Response], error) {
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
	if err := client.GetJSON(ctx, "boxscoremiscv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreMiscV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreMiscV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 22 {
				item := BoxScoreMiscV2PlayerStats{
					GAME_ID:            toString(row[0]),
					TEAM_ID:            toInt(row[1]),
					TEAM_ABBREVIATION:  toString(row[2]),
					TEAM_CITY:          toString(row[3]),
					PLAYER_ID:          toInt(row[4]),
					PLAYER_NAME:        toString(row[5]),
					NICKNAME:           toString(row[6]),
					START_POSITION:     toString(row[7]),
					COMMENT:            toString(row[8]),
					MIN:                toFloat(row[9]),
					PTS_OFF_TOV:        toFloat(row[10]),
					PTS_2ND_CHANCE:     toFloat(row[11]),
					PTS_FB:             toFloat(row[12]),
					PTS_PAINT:          toFloat(row[13]),
					OPP_PTS_OFF_TOV:    toFloat(row[14]),
					OPP_PTS_2ND_CHANCE: toFloat(row[15]),
					OPP_PTS_FB:         toFloat(row[16]),
					OPP_PTS_PAINT:      toFloat(row[17]),
					BLK:                toFloat(row[18]),
					BLKA:               toInt(row[19]),
					PF:                 toFloat(row[20]),
					PFD:                toFloat(row[21]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreMiscV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 18 {
				item := BoxScoreMiscV2TeamStats{
					GAME_ID:            toString(row[0]),
					TEAM_ID:            toInt(row[1]),
					TEAM_NAME:          toString(row[2]),
					TEAM_ABBREVIATION:  toString(row[3]),
					TEAM_CITY:          toString(row[4]),
					MIN:                toFloat(row[5]),
					PTS_OFF_TOV:        toFloat(row[6]),
					PTS_2ND_CHANCE:     toFloat(row[7]),
					PTS_FB:             toFloat(row[8]),
					PTS_PAINT:          toFloat(row[9]),
					OPP_PTS_OFF_TOV:    toFloat(row[10]),
					OPP_PTS_2ND_CHANCE: toFloat(row[11]),
					OPP_PTS_FB:         toFloat(row[12]),
					OPP_PTS_PAINT:      toFloat(row[13]),
					BLK:                toFloat(row[14]),
					BLKA:               toInt(row[15]),
					PF:                 toFloat(row[16]),
					PFD:                toFloat(row[17]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
