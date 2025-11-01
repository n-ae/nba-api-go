package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerDashboardByGameSplitsRequest contains parameters for the PlayerDashboardByGameSplits endpoint
type PlayerDashboardByGameSplitsRequest struct {
	PlayerID    string
	MeasureType *string
	PerMode     *parameters.PerMode
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// PlayerDashboardByGameSplitsOverallPlayerDashboard represents the OverallPlayerDashboard result set for PlayerDashboardByGameSplits
type PlayerDashboardByGameSplitsOverallPlayerDashboard struct {
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

// PlayerDashboardByGameSplitsLocationPlayerDashboard represents the LocationPlayerDashboard result set for PlayerDashboardByGameSplits
type PlayerDashboardByGameSplitsLocationPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	LOCATION    string  `json:"LOCATION"`
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

// PlayerDashboardByGameSplitsWinsLossesPlayerDashboard represents the WinsLossesPlayerDashboard result set for PlayerDashboardByGameSplits
type PlayerDashboardByGameSplitsWinsLossesPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	W_L         string  `json:"W_L"`
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

// PlayerDashboardByGameSplitsMonthPlayerDashboard represents the MonthPlayerDashboard result set for PlayerDashboardByGameSplits
type PlayerDashboardByGameSplitsMonthPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	MONTH       string  `json:"MONTH"`
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

// PlayerDashboardByGameSplitsPrePostAllStarPlayerDashboard represents the PrePostAllStarPlayerDashboard result set for PlayerDashboardByGameSplits
type PlayerDashboardByGameSplitsPrePostAllStarPlayerDashboard struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	PRE_POST_ALL_STAR string  `json:"PRE_POST_ALL_STAR"`
	GP                int     `json:"GP"`
	W                 string  `json:"W"`
	L                 string  `json:"L"`
	W_PCT             float64 `json:"W_PCT"`
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
	OREB              float64 `json:"OREB"`
	DREB              float64 `json:"DREB"`
	REB               float64 `json:"REB"`
	AST               float64 `json:"AST"`
	TOV               float64 `json:"TOV"`
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	BLKA              int     `json:"BLKA"`
	PF                float64 `json:"PF"`
	PFD               float64 `json:"PFD"`
	PTS               float64 `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// PlayerDashboardByGameSplitsDaysRestPlayerDashboard represents the DaysRestPlayerDashboard result set for PlayerDashboardByGameSplits
type PlayerDashboardByGameSplitsDaysRestPlayerDashboard struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	DAYS_REST   string  `json:"DAYS_REST"`
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

// PlayerDashboardByGameSplitsResponse contains the response data from the PlayerDashboardByGameSplits endpoint
type PlayerDashboardByGameSplitsResponse struct {
	OverallPlayerDashboard        []PlayerDashboardByGameSplitsOverallPlayerDashboard
	LocationPlayerDashboard       []PlayerDashboardByGameSplitsLocationPlayerDashboard
	WinsLossesPlayerDashboard     []PlayerDashboardByGameSplitsWinsLossesPlayerDashboard
	MonthPlayerDashboard          []PlayerDashboardByGameSplitsMonthPlayerDashboard
	PrePostAllStarPlayerDashboard []PlayerDashboardByGameSplitsPrePostAllStarPlayerDashboard
	DaysRestPlayerDashboard       []PlayerDashboardByGameSplitsDaysRestPlayerDashboard
}

// GetPlayerDashboardByGameSplits retrieves data from the playerdashboardbygamesplits endpoint
func GetPlayerDashboardByGameSplits(ctx context.Context, client *stats.Client, req PlayerDashboardByGameSplitsRequest) (*models.Response[*PlayerDashboardByGameSplitsResponse], error) {
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
	if err := client.GetJSON(ctx, "/playerdashboardbygamesplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerDashboardByGameSplitsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallPlayerDashboard = make([]PlayerDashboardByGameSplitsOverallPlayerDashboard, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := PlayerDashboardByGameSplitsOverallPlayerDashboard{
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
		response.LocationPlayerDashboard = make([]PlayerDashboardByGameSplitsLocationPlayerDashboard, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByGameSplitsLocationPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					LOCATION:    toString(row[2]),
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
				response.LocationPlayerDashboard = append(response.LocationPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.WinsLossesPlayerDashboard = make([]PlayerDashboardByGameSplitsWinsLossesPlayerDashboard, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByGameSplitsWinsLossesPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					W_L:         toString(row[2]),
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
				response.WinsLossesPlayerDashboard = append(response.WinsLossesPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.MonthPlayerDashboard = make([]PlayerDashboardByGameSplitsMonthPlayerDashboard, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByGameSplitsMonthPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					MONTH:       toString(row[2]),
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
				response.MonthPlayerDashboard = append(response.MonthPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.PrePostAllStarPlayerDashboard = make([]PlayerDashboardByGameSplitsPrePostAllStarPlayerDashboard, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByGameSplitsPrePostAllStarPlayerDashboard{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					PRE_POST_ALL_STAR: toString(row[2]),
					GP:                toInt(row[3]),
					W:                 toString(row[4]),
					L:                 toString(row[5]),
					W_PCT:             toFloat(row[6]),
					MIN:               toFloat(row[7]),
					FGM:               toInt(row[8]),
					FGA:               toInt(row[9]),
					FG_PCT:            toFloat(row[10]),
					FG3M:              toInt(row[11]),
					FG3A:              toInt(row[12]),
					FG3_PCT:           toFloat(row[13]),
					FTM:               toInt(row[14]),
					FTA:               toInt(row[15]),
					FT_PCT:            toFloat(row[16]),
					OREB:              toFloat(row[17]),
					DREB:              toFloat(row[18]),
					REB:               toFloat(row[19]),
					AST:               toFloat(row[20]),
					TOV:               toFloat(row[21]),
					STL:               toFloat(row[22]),
					BLK:               toFloat(row[23]),
					BLKA:              toInt(row[24]),
					PF:                toFloat(row[25]),
					PFD:               toFloat(row[26]),
					PTS:               toFloat(row[27]),
					PLUS_MINUS:        toFloat(row[28]),
				}
				response.PrePostAllStarPlayerDashboard = append(response.PrePostAllStarPlayerDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 5 {
		response.DaysRestPlayerDashboard = make([]PlayerDashboardByGameSplitsDaysRestPlayerDashboard, 0, len(rawResp.ResultSets[5].RowSet))
		for _, row := range rawResp.ResultSets[5].RowSet {
			if len(row) >= 29 {
				item := PlayerDashboardByGameSplitsDaysRestPlayerDashboard{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					DAYS_REST:   toString(row[2]),
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
				response.DaysRestPlayerDashboard = append(response.DaysRestPlayerDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
