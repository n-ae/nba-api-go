package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// BoxScoreHustleV2Request contains parameters for the BoxScoreHustleV2 endpoint
type BoxScoreHustleV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
	StartRange  *string
	EndRange    *string
	RangeType   *string
}

// BoxScoreHustleV2PlayerStats represents the PlayerStats result set for BoxScoreHustleV2
type BoxScoreHustleV2PlayerStats struct {
	GAME_ID                   string  `json:"GAME_ID"`
	TEAM_ID                   int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION         string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY                 string  `json:"TEAM_CITY"`
	PLAYER_ID                 int     `json:"PLAYER_ID"`
	PLAYER_NAME               string  `json:"PLAYER_NAME"`
	START_POSITION            string  `json:"START_POSITION"`
	COMMENT                   string  `json:"COMMENT"`
	MIN                       float64 `json:"MIN"`
	CONTESTED_SHOTS           string  `json:"CONTESTED_SHOTS"`
	CONTESTED_SHOTS_2PT       string  `json:"CONTESTED_SHOTS_2PT"`
	CONTESTED_SHOTS_3PT       string  `json:"CONTESTED_SHOTS_3PT"`
	DEFLECTIONS               string  `json:"DEFLECTIONS"`
	CHARGES_DRAWN             string  `json:"CHARGES_DRAWN"`
	SCREEN_ASSISTS            string  `json:"SCREEN_ASSISTS"`
	SCREEN_AST_PTS            float64 `json:"SCREEN_AST_PTS"`
	OFF_LOOSE_BALLS_RECOVERED string  `json:"OFF_LOOSE_BALLS_RECOVERED"`
	DEF_LOOSE_BALLS_RECOVERED string  `json:"DEF_LOOSE_BALLS_RECOVERED"`
	LOOSE_BALLS_RECOVERED     string  `json:"LOOSE_BALLS_RECOVERED"`
	OFF_BOXOUTS               string  `json:"OFF_BOXOUTS"`
	DEF_BOXOUTS               string  `json:"DEF_BOXOUTS"`
	BOX_OUTS                  string  `json:"BOX_OUTS"`
}

// BoxScoreHustleV2TeamStats represents the TeamStats result set for BoxScoreHustleV2
type BoxScoreHustleV2TeamStats struct {
	GAME_ID                   string  `json:"GAME_ID"`
	TEAM_ID                   int     `json:"TEAM_ID"`
	TEAM_NAME                 string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION         string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY                 string  `json:"TEAM_CITY"`
	MIN                       float64 `json:"MIN"`
	CONTESTED_SHOTS           string  `json:"CONTESTED_SHOTS"`
	CONTESTED_SHOTS_2PT       string  `json:"CONTESTED_SHOTS_2PT"`
	CONTESTED_SHOTS_3PT       string  `json:"CONTESTED_SHOTS_3PT"`
	DEFLECTIONS               string  `json:"DEFLECTIONS"`
	CHARGES_DRAWN             string  `json:"CHARGES_DRAWN"`
	SCREEN_ASSISTS            string  `json:"SCREEN_ASSISTS"`
	SCREEN_AST_PTS            float64 `json:"SCREEN_AST_PTS"`
	OFF_LOOSE_BALLS_RECOVERED string  `json:"OFF_LOOSE_BALLS_RECOVERED"`
	DEF_LOOSE_BALLS_RECOVERED string  `json:"DEF_LOOSE_BALLS_RECOVERED"`
	LOOSE_BALLS_RECOVERED     string  `json:"LOOSE_BALLS_RECOVERED"`
	OFF_BOXOUTS               string  `json:"OFF_BOXOUTS"`
	DEF_BOXOUTS               string  `json:"DEF_BOXOUTS"`
	BOX_OUTS                  string  `json:"BOX_OUTS"`
}

// BoxScoreHustleV2Response contains the response data from the BoxScoreHustleV2 endpoint
type BoxScoreHustleV2Response struct {
	PlayerStats []BoxScoreHustleV2PlayerStats
	TeamStats   []BoxScoreHustleV2TeamStats
}

// GetBoxScoreHustleV2 retrieves data from the boxscorehustlev2 endpoint
func GetBoxScoreHustleV2(ctx context.Context, client *stats.Client, req BoxScoreHustleV2Request) (*models.Response[*BoxScoreHustleV2Response], error) {
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
	if err := client.GetJSON(ctx, "boxscorehustlev2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreHustleV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]BoxScoreHustleV2PlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 22 {
				item := BoxScoreHustleV2PlayerStats{
					GAME_ID:                   toString(row[0]),
					TEAM_ID:                   toInt(row[1]),
					TEAM_ABBREVIATION:         toString(row[2]),
					TEAM_CITY:                 toString(row[3]),
					PLAYER_ID:                 toInt(row[4]),
					PLAYER_NAME:               toString(row[5]),
					START_POSITION:            toString(row[6]),
					COMMENT:                   toString(row[7]),
					MIN:                       toFloat(row[8]),
					CONTESTED_SHOTS:           toString(row[9]),
					CONTESTED_SHOTS_2PT:       toString(row[10]),
					CONTESTED_SHOTS_3PT:       toString(row[11]),
					DEFLECTIONS:               toString(row[12]),
					CHARGES_DRAWN:             toString(row[13]),
					SCREEN_ASSISTS:            toString(row[14]),
					SCREEN_AST_PTS:            toFloat(row[15]),
					OFF_LOOSE_BALLS_RECOVERED: toString(row[16]),
					DEF_LOOSE_BALLS_RECOVERED: toString(row[17]),
					LOOSE_BALLS_RECOVERED:     toString(row[18]),
					OFF_BOXOUTS:               toString(row[19]),
					DEF_BOXOUTS:               toString(row[20]),
					BOX_OUTS:                  toString(row[21]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TeamStats = make([]BoxScoreHustleV2TeamStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 19 {
				item := BoxScoreHustleV2TeamStats{
					GAME_ID:                   toString(row[0]),
					TEAM_ID:                   toInt(row[1]),
					TEAM_NAME:                 toString(row[2]),
					TEAM_ABBREVIATION:         toString(row[3]),
					TEAM_CITY:                 toString(row[4]),
					MIN:                       toFloat(row[5]),
					CONTESTED_SHOTS:           toString(row[6]),
					CONTESTED_SHOTS_2PT:       toString(row[7]),
					CONTESTED_SHOTS_3PT:       toString(row[8]),
					DEFLECTIONS:               toString(row[9]),
					CHARGES_DRAWN:             toString(row[10]),
					SCREEN_ASSISTS:            toString(row[11]),
					SCREEN_AST_PTS:            toFloat(row[12]),
					OFF_LOOSE_BALLS_RECOVERED: toString(row[13]),
					DEF_LOOSE_BALLS_RECOVERED: toString(row[14]),
					LOOSE_BALLS_RECOVERED:     toString(row[15]),
					OFF_BOXOUTS:               toString(row[16]),
					DEF_BOXOUTS:               toString(row[17]),
					BOX_OUTS:                  toString(row[18]),
				}
				response.TeamStats = append(response.TeamStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
