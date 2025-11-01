package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerDashboardByShootingSplitsRequest contains parameters for the PlayerDashboardByShootingSplits endpoint
type PlayerDashboardByShootingSplitsRequest struct {
	PlayerID    string
	MeasureType *string
	PerMode     *parameters.PerMode
	PlusMinus   *string
	PaceAdjust  *string
	Rank        *string
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// PlayerDashboardByShootingSplitsOverallPlayerDashboard represents the OverallPlayerDashboard result set for PlayerDashboardByShootingSplits
type PlayerDashboardByShootingSplitsOverallPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	GP          int     `json:"GP"`
	W           string  `json:"W"`
	L           string  `json:"L"`
	W_PCT       float64 `json:"W_PCT"`
	MIN         float64 `json:"MIN"`
	FGM         int     `json:"FGM"`
	FGA         int     `json:"FGA"`
	FG_PCT      float64 `json:"FG_PCT"`
	FG3M        int     `json:"FG3M"`
	FG3A        int     `json:"FG3A"`
	FG3_PCT     float64 `json:"FG3_PCT"`
	FTM         int     `json:"FTM"`
	FTA         int     `json:"FTA"`
	FT_PCT      float64 `json:"FT_PCT"`
	OREB        float64 `json:"OREB"`
	DREB        float64 `json:"DREB"`
	REB         float64 `json:"REB"`
	AST         float64 `json:"AST"`
	TOV         float64 `json:"TOV"`
	STL         float64 `json:"STL"`
	BLK         float64 `json:"BLK"`
	BLKA        int     `json:"BLKA"`
	PF          float64 `json:"PF"`
	PFD         float64 `json:"PFD"`
	PTS         float64 `json:"PTS"`
	PLUS_MINUS  float64 `json:"PLUS_MINUS"`
}

// PlayerDashboardByShootingSplitsShot5FTPlayerDashboard represents the Shot5FTPlayerDashboard result set for PlayerDashboardByShootingSplits
type PlayerDashboardByShootingSplitsShot5FTPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	SHOT_5FT    string  `json:"SHOT_5FT"`
	GP          int     `json:"GP"`
	W           string  `json:"W"`
	L           string  `json:"L"`
	W_PCT       float64 `json:"W_PCT"`
	MIN         float64 `json:"MIN"`
	FGM         int     `json:"FGM"`
	FGA         int     `json:"FGA"`
	FG_PCT      float64 `json:"FG_PCT"`
	FG3M        int     `json:"FG3M"`
	FG3A        int     `json:"FG3A"`
	FG3_PCT     float64 `json:"FG3_PCT"`
	FTM         int     `json:"FTM"`
	FTA         int     `json:"FTA"`
	FT_PCT      float64 `json:"FT_PCT"`
	OREB        float64 `json:"OREB"`
	DREB        float64 `json:"DREB"`
	REB         float64 `json:"REB"`
	AST         float64 `json:"AST"`
	TOV         float64 `json:"TOV"`
	STL         float64 `json:"STL"`
	BLK         float64 `json:"BLK"`
	BLKA        int     `json:"BLKA"`
	PF          float64 `json:"PF"`
	PFD         float64 `json:"PFD"`
	PTS         float64 `json:"PTS"`
	PLUS_MINUS  float64 `json:"PLUS_MINUS"`
}

// PlayerDashboardByShootingSplitsShot8FTPlayerDashboard represents the Shot8FTPlayerDashboard result set for PlayerDashboardByShootingSplits
type PlayerDashboardByShootingSplitsShot8FTPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	SHOT_8FT    string  `json:"SHOT_8FT"`
	GP          int     `json:"GP"`
	W           string  `json:"W"`
	L           string  `json:"L"`
	W_PCT       float64 `json:"W_PCT"`
	MIN         float64 `json:"MIN"`
	FGM         int     `json:"FGM"`
	FGA         int     `json:"FGA"`
	FG_PCT      float64 `json:"FG_PCT"`
	FG3M        int     `json:"FG3M"`
	FG3A        int     `json:"FG3A"`
	FG3_PCT     float64 `json:"FG3_PCT"`
	FTM         int     `json:"FTM"`
	FTA         int     `json:"FTA"`
	FT_PCT      float64 `json:"FT_PCT"`
	OREB        float64 `json:"OREB"`
	DREB        float64 `json:"DREB"`
	REB         float64 `json:"REB"`
	AST         float64 `json:"AST"`
	TOV         float64 `json:"TOV"`
	STL         float64 `json:"STL"`
	BLK         float64 `json:"BLK"`
	BLKA        int     `json:"BLKA"`
	PF          float64 `json:"PF"`
	PFD         float64 `json:"PFD"`
	PTS         float64 `json:"PTS"`
	PLUS_MINUS  float64 `json:"PLUS_MINUS"`
}

// PlayerDashboardByShootingSplitsShotAreaPlayerDashboard represents the ShotAreaPlayerDashboard result set for PlayerDashboardByShootingSplits
type PlayerDashboardByShootingSplitsShotAreaPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	SHOT_AREA   string  `json:"SHOT_AREA"`
	GP          int     `json:"GP"`
	W           string  `json:"W"`
	L           string  `json:"L"`
	W_PCT       float64 `json:"W_PCT"`
	MIN         float64 `json:"MIN"`
	FGM         int     `json:"FGM"`
	FGA         int     `json:"FGA"`
	FG_PCT      float64 `json:"FG_PCT"`
	FG3M        int     `json:"FG3M"`
	FG3A        int     `json:"FG3A"`
	FG3_PCT     float64 `json:"FG3_PCT"`
	FTM         int     `json:"FTM"`
	FTA         int     `json:"FTA"`
	FT_PCT      float64 `json:"FT_PCT"`
	OREB        float64 `json:"OREB"`
	DREB        float64 `json:"DREB"`
	REB         float64 `json:"REB"`
	AST         float64 `json:"AST"`
	TOV         float64 `json:"TOV"`
	STL         float64 `json:"STL"`
	BLK         float64 `json:"BLK"`
	BLKA        int     `json:"BLKA"`
	PF          float64 `json:"PF"`
	PFD         float64 `json:"PFD"`
	PTS         float64 `json:"PTS"`
	PLUS_MINUS  float64 `json:"PLUS_MINUS"`
}

// PlayerDashboardByShootingSplitsAssistedShotsPlayerDashboard represents the AssistedShotsPlayerDashboard result set for PlayerDashboardByShootingSplits
type PlayerDashboardByShootingSplitsAssistedShotsPlayerDashboard struct {
	PLAYER_ID          int     `json:"PLAYER_ID"`
	PLAYER_NAME        string  `json:"PLAYER_NAME"`
	ASSISTED_SHOT_TYPE string  `json:"ASSISTED_SHOT_TYPE"`
	GP                 int     `json:"GP"`
	W                  string  `json:"W"`
	L                  string  `json:"L"`
	W_PCT              float64 `json:"W_PCT"`
	MIN                float64 `json:"MIN"`
	FGM                int     `json:"FGM"`
	FGA                int     `json:"FGA"`
	FG_PCT             float64 `json:"FG_PCT"`
	FG3M               int     `json:"FG3M"`
	FG3A               int     `json:"FG3A"`
	FG3_PCT            float64 `json:"FG3_PCT"`
	FTM                int     `json:"FTM"`
	FTA                int     `json:"FTA"`
	FT_PCT             float64 `json:"FT_PCT"`
	OREB               float64 `json:"OREB"`
	DREB               float64 `json:"DREB"`
	REB                float64 `json:"REB"`
	AST                float64 `json:"AST"`
	TOV                float64 `json:"TOV"`
	STL                float64 `json:"STL"`
	BLK                float64 `json:"BLK"`
	BLKA               int     `json:"BLKA"`
	PF                 float64 `json:"PF"`
	PFD                float64 `json:"PFD"`
	PTS                float64 `json:"PTS"`
	PLUS_MINUS         float64 `json:"PLUS_MINUS"`
}

// PlayerDashboardByShootingSplitsResponse contains the response data from the PlayerDashboardByShootingSplits endpoint
type PlayerDashboardByShootingSplitsResponse struct {
	OverallPlayerDashboard       []PlayerDashboardByShootingSplitsOverallPlayerDashboard
	Shot5FTPlayerDashboard       []PlayerDashboardByShootingSplitsShot5FTPlayerDashboard
	Shot8FTPlayerDashboard       []PlayerDashboardByShootingSplitsShot8FTPlayerDashboard
	ShotAreaPlayerDashboard      []PlayerDashboardByShootingSplitsShotAreaPlayerDashboard
	AssistedShotsPlayerDashboard []PlayerDashboardByShootingSplitsAssistedShotsPlayerDashboard
}

// GetPlayerDashboardByShootingSplits retrieves data from the playerdashboardbyshootingsplits endpoint
func GetPlayerDashboardByShootingSplits(ctx context.Context, client *stats.Client, req PlayerDashboardByShootingSplitsRequest) (*models.Response[*PlayerDashboardByShootingSplitsResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
	if req.MeasureType != nil {
		params.Set("MeasureType", string(*req.MeasureType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.PlusMinus != nil {
		params.Set("PlusMinus", string(*req.PlusMinus))
	}
	if req.PaceAdjust != nil {
		params.Set("PaceAdjust", string(*req.PaceAdjust))
	}
	if req.Rank != nil {
		params.Set("Rank", string(*req.Rank))
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
	if err := client.GetJSON(ctx, "/playerdashboardbyshootingsplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerDashboardByShootingSplitsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallPlayerDashboard = make([]PlayerDashboardByShootingSplitsOverallPlayerDashboard, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := PlayerDashboardByShootingSplitsOverallPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					GP:          toInt(row[2]),
					W:           toString(row[3]),
					L:           toString(row[4]),
					W_PCT:       toFloat(row[5]),
					MIN:         toFloat(row[6]),
					FGM:         toInt(row[7]),
					FGA:         toInt(row[8]),
					FG_PCT:      toFloat(row[9]),
					FG3M:        toInt(row[10]),
					FG3A:        toInt(row[11]),
					FG3_PCT:     toFloat(row[12]),
					FTM:         toInt(row[13]),
					FTA:         toInt(row[14]),
					FT_PCT:      toFloat(row[15]),
					OREB:        toFloat(row[16]),
					DREB:        toFloat(row[17]),
					REB:         toFloat(row[18]),
					AST:         toFloat(row[19]),
					TOV:         toFloat(row[20]),
					STL:         toFloat(row[21]),
					BLK:         toFloat(row[22]),
					BLKA:        toInt(row[23]),
					PF:          toFloat(row[24]),
					PFD:         toFloat(row[25]),
					PTS:         toFloat(row[26]),
					PLUS_MINUS:  toFloat(row[27]),
				}
				response.OverallPlayerDashboard = append(response.OverallPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.Shot5FTPlayerDashboard = make([]PlayerDashboardByShootingSplitsShot5FTPlayerDashboard, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByShootingSplitsShot5FTPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					SHOT_5FT:    toString(row[2]),
					GP:          toInt(row[3]),
					W:           toString(row[4]),
					L:           toString(row[5]),
					W_PCT:       toFloat(row[6]),
					MIN:         toFloat(row[7]),
					FGM:         toInt(row[8]),
					FGA:         toInt(row[9]),
					FG_PCT:      toFloat(row[10]),
					FG3M:        toInt(row[11]),
					FG3A:        toInt(row[12]),
					FG3_PCT:     toFloat(row[13]),
					FTM:         toInt(row[14]),
					FTA:         toInt(row[15]),
					FT_PCT:      toFloat(row[16]),
					OREB:        toFloat(row[17]),
					DREB:        toFloat(row[18]),
					REB:         toFloat(row[19]),
					AST:         toFloat(row[20]),
					TOV:         toFloat(row[21]),
					STL:         toFloat(row[22]),
					BLK:         toFloat(row[23]),
					BLKA:        toInt(row[24]),
					PF:          toFloat(row[25]),
					PFD:         toFloat(row[26]),
					PTS:         toFloat(row[27]),
					PLUS_MINUS:  toFloat(row[28]),
				}
				response.Shot5FTPlayerDashboard = append(response.Shot5FTPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.Shot8FTPlayerDashboard = make([]PlayerDashboardByShootingSplitsShot8FTPlayerDashboard, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByShootingSplitsShot8FTPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					SHOT_8FT:    toString(row[2]),
					GP:          toInt(row[3]),
					W:           toString(row[4]),
					L:           toString(row[5]),
					W_PCT:       toFloat(row[6]),
					MIN:         toFloat(row[7]),
					FGM:         toInt(row[8]),
					FGA:         toInt(row[9]),
					FG_PCT:      toFloat(row[10]),
					FG3M:        toInt(row[11]),
					FG3A:        toInt(row[12]),
					FG3_PCT:     toFloat(row[13]),
					FTM:         toInt(row[14]),
					FTA:         toInt(row[15]),
					FT_PCT:      toFloat(row[16]),
					OREB:        toFloat(row[17]),
					DREB:        toFloat(row[18]),
					REB:         toFloat(row[19]),
					AST:         toFloat(row[20]),
					TOV:         toFloat(row[21]),
					STL:         toFloat(row[22]),
					BLK:         toFloat(row[23]),
					BLKA:        toInt(row[24]),
					PF:          toFloat(row[25]),
					PFD:         toFloat(row[26]),
					PTS:         toFloat(row[27]),
					PLUS_MINUS:  toFloat(row[28]),
				}
				response.Shot8FTPlayerDashboard = append(response.Shot8FTPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.ShotAreaPlayerDashboard = make([]PlayerDashboardByShootingSplitsShotAreaPlayerDashboard, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByShootingSplitsShotAreaPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					SHOT_AREA:   toString(row[2]),
					GP:          toInt(row[3]),
					W:           toString(row[4]),
					L:           toString(row[5]),
					W_PCT:       toFloat(row[6]),
					MIN:         toFloat(row[7]),
					FGM:         toInt(row[8]),
					FGA:         toInt(row[9]),
					FG_PCT:      toFloat(row[10]),
					FG3M:        toInt(row[11]),
					FG3A:        toInt(row[12]),
					FG3_PCT:     toFloat(row[13]),
					FTM:         toInt(row[14]),
					FTA:         toInt(row[15]),
					FT_PCT:      toFloat(row[16]),
					OREB:        toFloat(row[17]),
					DREB:        toFloat(row[18]),
					REB:         toFloat(row[19]),
					AST:         toFloat(row[20]),
					TOV:         toFloat(row[21]),
					STL:         toFloat(row[22]),
					BLK:         toFloat(row[23]),
					BLKA:        toInt(row[24]),
					PF:          toFloat(row[25]),
					PFD:         toFloat(row[26]),
					PTS:         toFloat(row[27]),
					PLUS_MINUS:  toFloat(row[28]),
				}
				response.ShotAreaPlayerDashboard = append(response.ShotAreaPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.AssistedShotsPlayerDashboard = make([]PlayerDashboardByShootingSplitsAssistedShotsPlayerDashboard, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByShootingSplitsAssistedShotsPlayerDashboard{
					PLAYER_ID:          toInt(row[0]),
					PLAYER_NAME:        toString(row[1]),
					ASSISTED_SHOT_TYPE: toString(row[2]),
					GP:                 toInt(row[3]),
					W:                  toString(row[4]),
					L:                  toString(row[5]),
					W_PCT:              toFloat(row[6]),
					MIN:                toFloat(row[7]),
					FGM:                toInt(row[8]),
					FGA:                toInt(row[9]),
					FG_PCT:             toFloat(row[10]),
					FG3M:               toInt(row[11]),
					FG3A:               toInt(row[12]),
					FG3_PCT:            toFloat(row[13]),
					FTM:                toInt(row[14]),
					FTA:                toInt(row[15]),
					FT_PCT:             toFloat(row[16]),
					OREB:               toFloat(row[17]),
					DREB:               toFloat(row[18]),
					REB:                toFloat(row[19]),
					AST:                toFloat(row[20]),
					TOV:                toFloat(row[21]),
					STL:                toFloat(row[22]),
					BLK:                toFloat(row[23]),
					BLKA:               toInt(row[24]),
					PF:                 toFloat(row[25]),
					PFD:                toFloat(row[26]),
					PTS:                toFloat(row[27]),
					PLUS_MINUS:         toFloat(row[28]),
				}
				response.AssistedShotsPlayerDashboard = append(response.AssistedShotsPlayerDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
