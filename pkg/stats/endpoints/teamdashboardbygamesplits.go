package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamDashboardByGameSplitsRequest contains parameters for the TeamDashboardByGameSplits endpoint
type TeamDashboardByGameSplitsRequest struct {
	TeamID      string
	MeasureType *string
	PerMode     *parameters.PerMode
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// TeamDashboardByGameSplitsOverallTeamDashboard represents the OverallTeamDashboard result set for TeamDashboardByGameSplits
type TeamDashboardByGameSplitsOverallTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	GP         int     `json:"GP"`
	W          string  `json:"W"`
	L          string  `json:"L"`
	W_PCT      float64 `json:"W_PCT"`
	MIN        float64 `json:"MIN"`
	FGM        int     `json:"FGM"`
	FGA        int     `json:"FGA"`
	FG_PCT     float64 `json:"FG_PCT"`
	FG3M       int     `json:"FG3M"`
	FG3A       int     `json:"FG3A"`
	FG3_PCT    float64 `json:"FG3_PCT"`
	FTM        int     `json:"FTM"`
	FTA        int     `json:"FTA"`
	FT_PCT     float64 `json:"FT_PCT"`
	OREB       float64 `json:"OREB"`
	DREB       float64 `json:"DREB"`
	REB        float64 `json:"REB"`
	AST        float64 `json:"AST"`
	TOV        float64 `json:"TOV"`
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	BLKA       int     `json:"BLKA"`
	PF         float64 `json:"PF"`
	PFD        float64 `json:"PFD"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByGameSplitsLocationTeamDashboard represents the LocationTeamDashboard result set for TeamDashboardByGameSplits
type TeamDashboardByGameSplitsLocationTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	LOCATION   string  `json:"LOCATION"`
	GP         int     `json:"GP"`
	W          string  `json:"W"`
	L          string  `json:"L"`
	W_PCT      float64 `json:"W_PCT"`
	MIN        float64 `json:"MIN"`
	FGM        int     `json:"FGM"`
	FGA        int     `json:"FGA"`
	FG_PCT     float64 `json:"FG_PCT"`
	FG3M       int     `json:"FG3M"`
	FG3A       int     `json:"FG3A"`
	FG3_PCT    float64 `json:"FG3_PCT"`
	FTM        int     `json:"FTM"`
	FTA        int     `json:"FTA"`
	FT_PCT     float64 `json:"FT_PCT"`
	OREB       float64 `json:"OREB"`
	DREB       float64 `json:"DREB"`
	REB        float64 `json:"REB"`
	AST        float64 `json:"AST"`
	TOV        float64 `json:"TOV"`
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	BLKA       int     `json:"BLKA"`
	PF         float64 `json:"PF"`
	PFD        float64 `json:"PFD"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByGameSplitsWinsLossesTeamDashboard represents the WinsLossesTeamDashboard result set for TeamDashboardByGameSplits
type TeamDashboardByGameSplitsWinsLossesTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	W_L        string  `json:"W_L"`
	GP         int     `json:"GP"`
	W          string  `json:"W"`
	L          string  `json:"L"`
	W_PCT      float64 `json:"W_PCT"`
	MIN        float64 `json:"MIN"`
	FGM        int     `json:"FGM"`
	FGA        int     `json:"FGA"`
	FG_PCT     float64 `json:"FG_PCT"`
	FG3M       int     `json:"FG3M"`
	FG3A       int     `json:"FG3A"`
	FG3_PCT    float64 `json:"FG3_PCT"`
	FTM        int     `json:"FTM"`
	FTA        int     `json:"FTA"`
	FT_PCT     float64 `json:"FT_PCT"`
	OREB       float64 `json:"OREB"`
	DREB       float64 `json:"DREB"`
	REB        float64 `json:"REB"`
	AST        float64 `json:"AST"`
	TOV        float64 `json:"TOV"`
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	BLKA       int     `json:"BLKA"`
	PF         float64 `json:"PF"`
	PFD        float64 `json:"PFD"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByGameSplitsMonthTeamDashboard represents the MonthTeamDashboard result set for TeamDashboardByGameSplits
type TeamDashboardByGameSplitsMonthTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	MONTH      string  `json:"MONTH"`
	GP         int     `json:"GP"`
	W          string  `json:"W"`
	L          string  `json:"L"`
	W_PCT      float64 `json:"W_PCT"`
	MIN        float64 `json:"MIN"`
	FGM        int     `json:"FGM"`
	FGA        int     `json:"FGA"`
	FG_PCT     float64 `json:"FG_PCT"`
	FG3M       int     `json:"FG3M"`
	FG3A       int     `json:"FG3A"`
	FG3_PCT    float64 `json:"FG3_PCT"`
	FTM        int     `json:"FTM"`
	FTA        int     `json:"FTA"`
	FT_PCT     float64 `json:"FT_PCT"`
	OREB       float64 `json:"OREB"`
	DREB       float64 `json:"DREB"`
	REB        float64 `json:"REB"`
	AST        float64 `json:"AST"`
	TOV        float64 `json:"TOV"`
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	BLKA       int     `json:"BLKA"`
	PF         float64 `json:"PF"`
	PFD        float64 `json:"PFD"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByGameSplitsPrePostAllStarTeamDashboard represents the PrePostAllStarTeamDashboard result set for TeamDashboardByGameSplits
type TeamDashboardByGameSplitsPrePostAllStarTeamDashboard struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
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

// TeamDashboardByGameSplitsDaysRestTeamDashboard represents the DaysRestTeamDashboard result set for TeamDashboardByGameSplits
type TeamDashboardByGameSplitsDaysRestTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	DAYS_REST  string  `json:"DAYS_REST"`
	GP         int     `json:"GP"`
	W          string  `json:"W"`
	L          string  `json:"L"`
	W_PCT      float64 `json:"W_PCT"`
	MIN        float64 `json:"MIN"`
	FGM        int     `json:"FGM"`
	FGA        int     `json:"FGA"`
	FG_PCT     float64 `json:"FG_PCT"`
	FG3M       int     `json:"FG3M"`
	FG3A       int     `json:"FG3A"`
	FG3_PCT    float64 `json:"FG3_PCT"`
	FTM        int     `json:"FTM"`
	FTA        int     `json:"FTA"`
	FT_PCT     float64 `json:"FT_PCT"`
	OREB       float64 `json:"OREB"`
	DREB       float64 `json:"DREB"`
	REB        float64 `json:"REB"`
	AST        float64 `json:"AST"`
	TOV        float64 `json:"TOV"`
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	BLKA       int     `json:"BLKA"`
	PF         float64 `json:"PF"`
	PFD        float64 `json:"PFD"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByGameSplitsResponse contains the response data from the TeamDashboardByGameSplits endpoint
type TeamDashboardByGameSplitsResponse struct {
	OverallTeamDashboard        []TeamDashboardByGameSplitsOverallTeamDashboard
	LocationTeamDashboard       []TeamDashboardByGameSplitsLocationTeamDashboard
	WinsLossesTeamDashboard     []TeamDashboardByGameSplitsWinsLossesTeamDashboard
	MonthTeamDashboard          []TeamDashboardByGameSplitsMonthTeamDashboard
	PrePostAllStarTeamDashboard []TeamDashboardByGameSplitsPrePostAllStarTeamDashboard
	DaysRestTeamDashboard       []TeamDashboardByGameSplitsDaysRestTeamDashboard
}

// GetTeamDashboardByGameSplits retrieves data from the teamdashboardbygamesplits endpoint
func GetTeamDashboardByGameSplits(ctx context.Context, client *stats.Client, req TeamDashboardByGameSplitsRequest) (*models.Response[*TeamDashboardByGameSplitsResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
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
	if err := client.GetJSON(ctx, "/teamdashboardbygamesplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDashboardByGameSplitsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallTeamDashboard = make([]TeamDashboardByGameSplitsOverallTeamDashboard, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := TeamDashboardByGameSplitsOverallTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					GP:         toInt(row[2]),
					W:          toString(row[3]),
					L:          toString(row[4]),
					W_PCT:      toFloat(row[5]),
					MIN:        toFloat(row[6]),
					FGM:        toInt(row[7]),
					FGA:        toInt(row[8]),
					FG_PCT:     toFloat(row[9]),
					FG3M:       toInt(row[10]),
					FG3A:       toInt(row[11]),
					FG3_PCT:    toFloat(row[12]),
					FTM:        toInt(row[13]),
					FTA:        toInt(row[14]),
					FT_PCT:     toFloat(row[15]),
					OREB:       toFloat(row[16]),
					DREB:       toFloat(row[17]),
					REB:        toFloat(row[18]),
					AST:        toFloat(row[19]),
					TOV:        toFloat(row[20]),
					STL:        toFloat(row[21]),
					BLK:        toFloat(row[22]),
					BLKA:       toInt(row[23]),
					PF:         toFloat(row[24]),
					PFD:        toFloat(row[25]),
					PTS:        toFloat(row[26]),
					PLUS_MINUS: toFloat(row[27]),
				}
				response.OverallTeamDashboard = append(response.OverallTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.LocationTeamDashboard = make([]TeamDashboardByGameSplitsLocationTeamDashboard, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByGameSplitsLocationTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					LOCATION:   toString(row[2]),
					GP:         toInt(row[3]),
					W:          toString(row[4]),
					L:          toString(row[5]),
					W_PCT:      toFloat(row[6]),
					MIN:        toFloat(row[7]),
					FGM:        toInt(row[8]),
					FGA:        toInt(row[9]),
					FG_PCT:     toFloat(row[10]),
					FG3M:       toInt(row[11]),
					FG3A:       toInt(row[12]),
					FG3_PCT:    toFloat(row[13]),
					FTM:        toInt(row[14]),
					FTA:        toInt(row[15]),
					FT_PCT:     toFloat(row[16]),
					OREB:       toFloat(row[17]),
					DREB:       toFloat(row[18]),
					REB:        toFloat(row[19]),
					AST:        toFloat(row[20]),
					TOV:        toFloat(row[21]),
					STL:        toFloat(row[22]),
					BLK:        toFloat(row[23]),
					BLKA:       toInt(row[24]),
					PF:         toFloat(row[25]),
					PFD:        toFloat(row[26]),
					PTS:        toFloat(row[27]),
					PLUS_MINUS: toFloat(row[28]),
				}
				response.LocationTeamDashboard = append(response.LocationTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.WinsLossesTeamDashboard = make([]TeamDashboardByGameSplitsWinsLossesTeamDashboard, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByGameSplitsWinsLossesTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					W_L:        toString(row[2]),
					GP:         toInt(row[3]),
					W:          toString(row[4]),
					L:          toString(row[5]),
					W_PCT:      toFloat(row[6]),
					MIN:        toFloat(row[7]),
					FGM:        toInt(row[8]),
					FGA:        toInt(row[9]),
					FG_PCT:     toFloat(row[10]),
					FG3M:       toInt(row[11]),
					FG3A:       toInt(row[12]),
					FG3_PCT:    toFloat(row[13]),
					FTM:        toInt(row[14]),
					FTA:        toInt(row[15]),
					FT_PCT:     toFloat(row[16]),
					OREB:       toFloat(row[17]),
					DREB:       toFloat(row[18]),
					REB:        toFloat(row[19]),
					AST:        toFloat(row[20]),
					TOV:        toFloat(row[21]),
					STL:        toFloat(row[22]),
					BLK:        toFloat(row[23]),
					BLKA:       toInt(row[24]),
					PF:         toFloat(row[25]),
					PFD:        toFloat(row[26]),
					PTS:        toFloat(row[27]),
					PLUS_MINUS: toFloat(row[28]),
				}
				response.WinsLossesTeamDashboard = append(response.WinsLossesTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.MonthTeamDashboard = make([]TeamDashboardByGameSplitsMonthTeamDashboard, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByGameSplitsMonthTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					MONTH:      toString(row[2]),
					GP:         toInt(row[3]),
					W:          toString(row[4]),
					L:          toString(row[5]),
					W_PCT:      toFloat(row[6]),
					MIN:        toFloat(row[7]),
					FGM:        toInt(row[8]),
					FGA:        toInt(row[9]),
					FG_PCT:     toFloat(row[10]),
					FG3M:       toInt(row[11]),
					FG3A:       toInt(row[12]),
					FG3_PCT:    toFloat(row[13]),
					FTM:        toInt(row[14]),
					FTA:        toInt(row[15]),
					FT_PCT:     toFloat(row[16]),
					OREB:       toFloat(row[17]),
					DREB:       toFloat(row[18]),
					REB:        toFloat(row[19]),
					AST:        toFloat(row[20]),
					TOV:        toFloat(row[21]),
					STL:        toFloat(row[22]),
					BLK:        toFloat(row[23]),
					BLKA:       toInt(row[24]),
					PF:         toFloat(row[25]),
					PFD:        toFloat(row[26]),
					PTS:        toFloat(row[27]),
					PLUS_MINUS: toFloat(row[28]),
				}
				response.MonthTeamDashboard = append(response.MonthTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.PrePostAllStarTeamDashboard = make([]TeamDashboardByGameSplitsPrePostAllStarTeamDashboard, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByGameSplitsPrePostAllStarTeamDashboard{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
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
				response.PrePostAllStarTeamDashboard = append(response.PrePostAllStarTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 5 {
		response.DaysRestTeamDashboard = make([]TeamDashboardByGameSplitsDaysRestTeamDashboard, 0, len(rawResp.ResultSets[5].RowSet))
		for _, row := range rawResp.ResultSets[5].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByGameSplitsDaysRestTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					DAYS_REST:  toString(row[2]),
					GP:         toInt(row[3]),
					W:          toString(row[4]),
					L:          toString(row[5]),
					W_PCT:      toFloat(row[6]),
					MIN:        toFloat(row[7]),
					FGM:        toInt(row[8]),
					FGA:        toInt(row[9]),
					FG_PCT:     toFloat(row[10]),
					FG3M:       toInt(row[11]),
					FG3A:       toInt(row[12]),
					FG3_PCT:    toFloat(row[13]),
					FTM:        toInt(row[14]),
					FTA:        toInt(row[15]),
					FT_PCT:     toFloat(row[16]),
					OREB:       toFloat(row[17]),
					DREB:       toFloat(row[18]),
					REB:        toFloat(row[19]),
					AST:        toFloat(row[20]),
					TOV:        toFloat(row[21]),
					STL:        toFloat(row[22]),
					BLK:        toFloat(row[23]),
					BLKA:       toInt(row[24]),
					PF:         toFloat(row[25]),
					PFD:        toFloat(row[26]),
					PTS:        toFloat(row[27]),
					PLUS_MINUS: toFloat(row[28]),
				}
				response.DaysRestTeamDashboard = append(response.DaysRestTeamDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
