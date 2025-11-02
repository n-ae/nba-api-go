package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamDashboardByGeneralSplitsRequest contains parameters for the TeamDashboardByGeneralSplits endpoint
type TeamDashboardByGeneralSplitsRequest struct {
	TeamID         string
	Season         parameters.Season
	SeasonType     parameters.SeasonType
	MeasureType    *string
	PerMode        *parameters.PerMode
	PlusMinus      *string
	PaceAdjust     *string
	Rank           *string
	LeagueID       *parameters.LeagueID
	Outcome        *string
	Location       *string
	Month          *string
	SeasonSegment  *string
	DateFrom       *string
	DateTo         *string
	OpponentTeamID *string
	VsConference   *string
	VsDivision     *string
	GameSegment    *string
	Period         *string
	LastNGames     *string
}

// TeamDashboardByGeneralSplitsOverallTeamDashboard represents the OverallTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsOverallTeamDashboard struct {
	GROUP_SET       interface{}
	GROUP_VALUE     interface{}
	GP              interface{}
	W               interface{}
	L               interface{}
	W_PCT           interface{}
	MIN             interface{}
	FGM             interface{}
	FGA             interface{}
	FG_PCT          interface{}
	FG3M            interface{}
	FG3A            interface{}
	FG3_PCT         interface{}
	FTM             interface{}
	FTA             interface{}
	FT_PCT          interface{}
	OREB            interface{}
	DREB            interface{}
	REB             interface{}
	AST             interface{}
	TOV             interface{}
	STL             interface{}
	BLK             interface{}
	BLKA            interface{}
	PF              interface{}
	PFD             interface{}
	PTS             interface{}
	PLUS_MINUS      interface{}
	GP_RANK         interface{}
	W_RANK          interface{}
	L_RANK          interface{}
	W_PCT_RANK      interface{}
	MIN_RANK        interface{}
	FGM_RANK        interface{}
	FGA_RANK        interface{}
	FG_PCT_RANK     interface{}
	FG3M_RANK       interface{}
	FG3A_RANK       interface{}
	FG3_PCT_RANK    interface{}
	FTM_RANK        interface{}
	FTA_RANK        interface{}
	FT_PCT_RANK     interface{}
	OREB_RANK       interface{}
	DREB_RANK       interface{}
	REB_RANK        interface{}
	AST_RANK        interface{}
	TOV_RANK        interface{}
	STL_RANK        interface{}
	BLK_RANK        interface{}
	BLKA_RANK       interface{}
	PF_RANK         interface{}
	PFD_RANK        interface{}
	PTS_RANK        interface{}
	PLUS_MINUS_RANK interface{}
}

// TeamDashboardByGeneralSplitsLocationTeamDashboard represents the LocationTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsLocationTeamDashboard struct {
	GROUP_SET       interface{}
	GROUP_VALUE     interface{}
	GP              interface{}
	W               interface{}
	L               interface{}
	W_PCT           interface{}
	MIN             interface{}
	FGM             interface{}
	FGA             interface{}
	FG_PCT          interface{}
	FG3M            interface{}
	FG3A            interface{}
	FG3_PCT         interface{}
	FTM             interface{}
	FTA             interface{}
	FT_PCT          interface{}
	OREB            interface{}
	DREB            interface{}
	REB             interface{}
	AST             interface{}
	TOV             interface{}
	STL             interface{}
	BLK             interface{}
	BLKA            interface{}
	PF              interface{}
	PFD             interface{}
	PTS             interface{}
	PLUS_MINUS      interface{}
	GP_RANK         interface{}
	W_RANK          interface{}
	L_RANK          interface{}
	W_PCT_RANK      interface{}
	MIN_RANK        interface{}
	FGM_RANK        interface{}
	FGA_RANK        interface{}
	FG_PCT_RANK     interface{}
	FG3M_RANK       interface{}
	FG3A_RANK       interface{}
	FG3_PCT_RANK    interface{}
	FTM_RANK        interface{}
	FTA_RANK        interface{}
	FT_PCT_RANK     interface{}
	OREB_RANK       interface{}
	DREB_RANK       interface{}
	REB_RANK        interface{}
	AST_RANK        interface{}
	TOV_RANK        interface{}
	STL_RANK        interface{}
	BLK_RANK        interface{}
	BLKA_RANK       interface{}
	PF_RANK         interface{}
	PFD_RANK        interface{}
	PTS_RANK        interface{}
	PLUS_MINUS_RANK interface{}
}

// TeamDashboardByGeneralSplitsWinsLossesTeamDashboard represents the WinsLossesTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsWinsLossesTeamDashboard struct {
	GROUP_SET       interface{}
	GROUP_VALUE     interface{}
	GP              interface{}
	W               interface{}
	L               interface{}
	W_PCT           interface{}
	MIN             interface{}
	FGM             interface{}
	FGA             interface{}
	FG_PCT          interface{}
	FG3M            interface{}
	FG3A            interface{}
	FG3_PCT         interface{}
	FTM             interface{}
	FTA             interface{}
	FT_PCT          interface{}
	OREB            interface{}
	DREB            interface{}
	REB             interface{}
	AST             interface{}
	TOV             interface{}
	STL             interface{}
	BLK             interface{}
	BLKA            interface{}
	PF              interface{}
	PFD             interface{}
	PTS             interface{}
	PLUS_MINUS      interface{}
	GP_RANK         interface{}
	W_RANK          interface{}
	L_RANK          interface{}
	W_PCT_RANK      interface{}
	MIN_RANK        interface{}
	FGM_RANK        interface{}
	FGA_RANK        interface{}
	FG_PCT_RANK     interface{}
	FG3M_RANK       interface{}
	FG3A_RANK       interface{}
	FG3_PCT_RANK    interface{}
	FTM_RANK        interface{}
	FTA_RANK        interface{}
	FT_PCT_RANK     interface{}
	OREB_RANK       interface{}
	DREB_RANK       interface{}
	REB_RANK        interface{}
	AST_RANK        interface{}
	TOV_RANK        interface{}
	STL_RANK        interface{}
	BLK_RANK        interface{}
	BLKA_RANK       interface{}
	PF_RANK         interface{}
	PFD_RANK        interface{}
	PTS_RANK        interface{}
	PLUS_MINUS_RANK interface{}
}

// TeamDashboardByGeneralSplitsMonthTeamDashboard represents the MonthTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsMonthTeamDashboard struct {
	GROUP_SET       interface{}
	GROUP_VALUE     interface{}
	GP              interface{}
	W               interface{}
	L               interface{}
	W_PCT           interface{}
	MIN             interface{}
	FGM             interface{}
	FGA             interface{}
	FG_PCT          interface{}
	FG3M            interface{}
	FG3A            interface{}
	FG3_PCT         interface{}
	FTM             interface{}
	FTA             interface{}
	FT_PCT          interface{}
	OREB            interface{}
	DREB            interface{}
	REB             interface{}
	AST             interface{}
	TOV             interface{}
	STL             interface{}
	BLK             interface{}
	BLKA            interface{}
	PF              interface{}
	PFD             interface{}
	PTS             interface{}
	PLUS_MINUS      interface{}
	GP_RANK         interface{}
	W_RANK          interface{}
	L_RANK          interface{}
	W_PCT_RANK      interface{}
	MIN_RANK        interface{}
	FGM_RANK        interface{}
	FGA_RANK        interface{}
	FG_PCT_RANK     interface{}
	FG3M_RANK       interface{}
	FG3A_RANK       interface{}
	FG3_PCT_RANK    interface{}
	FTM_RANK        interface{}
	FTA_RANK        interface{}
	FT_PCT_RANK     interface{}
	OREB_RANK       interface{}
	DREB_RANK       interface{}
	REB_RANK        interface{}
	AST_RANK        interface{}
	TOV_RANK        interface{}
	STL_RANK        interface{}
	BLK_RANK        interface{}
	BLKA_RANK       interface{}
	PF_RANK         interface{}
	PFD_RANK        interface{}
	PTS_RANK        interface{}
	PLUS_MINUS_RANK interface{}
}

// TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard represents the PrePostAllStarTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard struct {
	GROUP_SET       interface{}
	GROUP_VALUE     interface{}
	GP              interface{}
	W               interface{}
	L               interface{}
	W_PCT           interface{}
	MIN             interface{}
	FGM             interface{}
	FGA             interface{}
	FG_PCT          interface{}
	FG3M            interface{}
	FG3A            interface{}
	FG3_PCT         interface{}
	FTM             interface{}
	FTA             interface{}
	FT_PCT          interface{}
	OREB            interface{}
	DREB            interface{}
	REB             interface{}
	AST             interface{}
	TOV             interface{}
	STL             interface{}
	BLK             interface{}
	BLKA            interface{}
	PF              interface{}
	PFD             interface{}
	PTS             interface{}
	PLUS_MINUS      interface{}
	GP_RANK         interface{}
	W_RANK          interface{}
	L_RANK          interface{}
	W_PCT_RANK      interface{}
	MIN_RANK        interface{}
	FGM_RANK        interface{}
	FGA_RANK        interface{}
	FG_PCT_RANK     interface{}
	FG3M_RANK       interface{}
	FG3A_RANK       interface{}
	FG3_PCT_RANK    interface{}
	FTM_RANK        interface{}
	FTA_RANK        interface{}
	FT_PCT_RANK     interface{}
	OREB_RANK       interface{}
	DREB_RANK       interface{}
	REB_RANK        interface{}
	AST_RANK        interface{}
	TOV_RANK        interface{}
	STL_RANK        interface{}
	BLK_RANK        interface{}
	BLKA_RANK       interface{}
	PF_RANK         interface{}
	PFD_RANK        interface{}
	PTS_RANK        interface{}
	PLUS_MINUS_RANK interface{}
}

// TeamDashboardByGeneralSplitsResponse contains the response data from the TeamDashboardByGeneralSplits endpoint
type TeamDashboardByGeneralSplitsResponse struct {
	OverallTeamDashboard        []TeamDashboardByGeneralSplitsOverallTeamDashboard
	LocationTeamDashboard       []TeamDashboardByGeneralSplitsLocationTeamDashboard
	WinsLossesTeamDashboard     []TeamDashboardByGeneralSplitsWinsLossesTeamDashboard
	MonthTeamDashboard          []TeamDashboardByGeneralSplitsMonthTeamDashboard
	PrePostAllStarTeamDashboard []TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard
}

// GetTeamDashboardByGeneralSplits retrieves data from the teamdashboardbygeneralsplits endpoint
func GetTeamDashboardByGeneralSplits(ctx context.Context, client *stats.Client, req TeamDashboardByGeneralSplitsRequest) (*models.Response[*TeamDashboardByGeneralSplitsResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.Season == "" {
		return nil, fmt.Errorf("Season is required")
	}
	params.Set("Season", string(req.Season))
	if req.SeasonType == "" {
		return nil, fmt.Errorf("SeasonType is required")
	}
	params.Set("SeasonType", string(req.SeasonType))
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
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Outcome != nil {
		params.Set("Outcome", string(*req.Outcome))
	}
	if req.Location != nil {
		params.Set("Location", string(*req.Location))
	}
	if req.Month != nil {
		params.Set("Month", string(*req.Month))
	}
	if req.SeasonSegment != nil {
		params.Set("SeasonSegment", string(*req.SeasonSegment))
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", string(*req.DateFrom))
	}
	if req.DateTo != nil {
		params.Set("DateTo", string(*req.DateTo))
	}
	if req.OpponentTeamID != nil {
		params.Set("OpponentTeamID", string(*req.OpponentTeamID))
	}
	if req.VsConference != nil {
		params.Set("VsConference", string(*req.VsConference))
	}
	if req.VsDivision != nil {
		params.Set("VsDivision", string(*req.VsDivision))
	}
	if req.GameSegment != nil {
		params.Set("GameSegment", string(*req.GameSegment))
	}
	if req.Period != nil {
		params.Set("Period", string(*req.Period))
	}
	if req.LastNGames != nil {
		params.Set("LastNGames", string(*req.LastNGames))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamdashboardbygeneralsplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDashboardByGeneralSplitsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallTeamDashboard = make([]TeamDashboardByGeneralSplitsOverallTeamDashboard, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 54 {
				response.OverallTeamDashboard[i] = TeamDashboardByGeneralSplitsOverallTeamDashboard{
					GROUP_SET:       row[0],
					GROUP_VALUE:     row[1],
					GP:              row[2],
					W:               row[3],
					L:               row[4],
					W_PCT:           row[5],
					MIN:             row[6],
					FGM:             row[7],
					FGA:             row[8],
					FG_PCT:          row[9],
					FG3M:            row[10],
					FG3A:            row[11],
					FG3_PCT:         row[12],
					FTM:             row[13],
					FTA:             row[14],
					FT_PCT:          row[15],
					OREB:            row[16],
					DREB:            row[17],
					REB:             row[18],
					AST:             row[19],
					TOV:             row[20],
					STL:             row[21],
					BLK:             row[22],
					BLKA:            row[23],
					PF:              row[24],
					PFD:             row[25],
					PTS:             row[26],
					PLUS_MINUS:      row[27],
					GP_RANK:         row[28],
					W_RANK:          row[29],
					L_RANK:          row[30],
					W_PCT_RANK:      row[31],
					MIN_RANK:        row[32],
					FGM_RANK:        row[33],
					FGA_RANK:        row[34],
					FG_PCT_RANK:     row[35],
					FG3M_RANK:       row[36],
					FG3A_RANK:       row[37],
					FG3_PCT_RANK:    row[38],
					FTM_RANK:        row[39],
					FTA_RANK:        row[40],
					FT_PCT_RANK:     row[41],
					OREB_RANK:       row[42],
					DREB_RANK:       row[43],
					REB_RANK:        row[44],
					AST_RANK:        row[45],
					TOV_RANK:        row[46],
					STL_RANK:        row[47],
					BLK_RANK:        row[48],
					BLKA_RANK:       row[49],
					PF_RANK:         row[50],
					PFD_RANK:        row[51],
					PTS_RANK:        row[52],
					PLUS_MINUS_RANK: row[53],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.LocationTeamDashboard = make([]TeamDashboardByGeneralSplitsLocationTeamDashboard, len(rawResp.ResultSets[1].RowSet))
		for i, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 54 {
				response.LocationTeamDashboard[i] = TeamDashboardByGeneralSplitsLocationTeamDashboard{
					GROUP_SET:       row[0],
					GROUP_VALUE:     row[1],
					GP:              row[2],
					W:               row[3],
					L:               row[4],
					W_PCT:           row[5],
					MIN:             row[6],
					FGM:             row[7],
					FGA:             row[8],
					FG_PCT:          row[9],
					FG3M:            row[10],
					FG3A:            row[11],
					FG3_PCT:         row[12],
					FTM:             row[13],
					FTA:             row[14],
					FT_PCT:          row[15],
					OREB:            row[16],
					DREB:            row[17],
					REB:             row[18],
					AST:             row[19],
					TOV:             row[20],
					STL:             row[21],
					BLK:             row[22],
					BLKA:            row[23],
					PF:              row[24],
					PFD:             row[25],
					PTS:             row[26],
					PLUS_MINUS:      row[27],
					GP_RANK:         row[28],
					W_RANK:          row[29],
					L_RANK:          row[30],
					W_PCT_RANK:      row[31],
					MIN_RANK:        row[32],
					FGM_RANK:        row[33],
					FGA_RANK:        row[34],
					FG_PCT_RANK:     row[35],
					FG3M_RANK:       row[36],
					FG3A_RANK:       row[37],
					FG3_PCT_RANK:    row[38],
					FTM_RANK:        row[39],
					FTA_RANK:        row[40],
					FT_PCT_RANK:     row[41],
					OREB_RANK:       row[42],
					DREB_RANK:       row[43],
					REB_RANK:        row[44],
					AST_RANK:        row[45],
					TOV_RANK:        row[46],
					STL_RANK:        row[47],
					BLK_RANK:        row[48],
					BLKA_RANK:       row[49],
					PF_RANK:         row[50],
					PFD_RANK:        row[51],
					PTS_RANK:        row[52],
					PLUS_MINUS_RANK: row[53],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.WinsLossesTeamDashboard = make([]TeamDashboardByGeneralSplitsWinsLossesTeamDashboard, len(rawResp.ResultSets[2].RowSet))
		for i, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 54 {
				response.WinsLossesTeamDashboard[i] = TeamDashboardByGeneralSplitsWinsLossesTeamDashboard{
					GROUP_SET:       row[0],
					GROUP_VALUE:     row[1],
					GP:              row[2],
					W:               row[3],
					L:               row[4],
					W_PCT:           row[5],
					MIN:             row[6],
					FGM:             row[7],
					FGA:             row[8],
					FG_PCT:          row[9],
					FG3M:            row[10],
					FG3A:            row[11],
					FG3_PCT:         row[12],
					FTM:             row[13],
					FTA:             row[14],
					FT_PCT:          row[15],
					OREB:            row[16],
					DREB:            row[17],
					REB:             row[18],
					AST:             row[19],
					TOV:             row[20],
					STL:             row[21],
					BLK:             row[22],
					BLKA:            row[23],
					PF:              row[24],
					PFD:             row[25],
					PTS:             row[26],
					PLUS_MINUS:      row[27],
					GP_RANK:         row[28],
					W_RANK:          row[29],
					L_RANK:          row[30],
					W_PCT_RANK:      row[31],
					MIN_RANK:        row[32],
					FGM_RANK:        row[33],
					FGA_RANK:        row[34],
					FG_PCT_RANK:     row[35],
					FG3M_RANK:       row[36],
					FG3A_RANK:       row[37],
					FG3_PCT_RANK:    row[38],
					FTM_RANK:        row[39],
					FTA_RANK:        row[40],
					FT_PCT_RANK:     row[41],
					OREB_RANK:       row[42],
					DREB_RANK:       row[43],
					REB_RANK:        row[44],
					AST_RANK:        row[45],
					TOV_RANK:        row[46],
					STL_RANK:        row[47],
					BLK_RANK:        row[48],
					BLKA_RANK:       row[49],
					PF_RANK:         row[50],
					PFD_RANK:        row[51],
					PTS_RANK:        row[52],
					PLUS_MINUS_RANK: row[53],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.MonthTeamDashboard = make([]TeamDashboardByGeneralSplitsMonthTeamDashboard, len(rawResp.ResultSets[3].RowSet))
		for i, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 54 {
				response.MonthTeamDashboard[i] = TeamDashboardByGeneralSplitsMonthTeamDashboard{
					GROUP_SET:       row[0],
					GROUP_VALUE:     row[1],
					GP:              row[2],
					W:               row[3],
					L:               row[4],
					W_PCT:           row[5],
					MIN:             row[6],
					FGM:             row[7],
					FGA:             row[8],
					FG_PCT:          row[9],
					FG3M:            row[10],
					FG3A:            row[11],
					FG3_PCT:         row[12],
					FTM:             row[13],
					FTA:             row[14],
					FT_PCT:          row[15],
					OREB:            row[16],
					DREB:            row[17],
					REB:             row[18],
					AST:             row[19],
					TOV:             row[20],
					STL:             row[21],
					BLK:             row[22],
					BLKA:            row[23],
					PF:              row[24],
					PFD:             row[25],
					PTS:             row[26],
					PLUS_MINUS:      row[27],
					GP_RANK:         row[28],
					W_RANK:          row[29],
					L_RANK:          row[30],
					W_PCT_RANK:      row[31],
					MIN_RANK:        row[32],
					FGM_RANK:        row[33],
					FGA_RANK:        row[34],
					FG_PCT_RANK:     row[35],
					FG3M_RANK:       row[36],
					FG3A_RANK:       row[37],
					FG3_PCT_RANK:    row[38],
					FTM_RANK:        row[39],
					FTA_RANK:        row[40],
					FT_PCT_RANK:     row[41],
					OREB_RANK:       row[42],
					DREB_RANK:       row[43],
					REB_RANK:        row[44],
					AST_RANK:        row[45],
					TOV_RANK:        row[46],
					STL_RANK:        row[47],
					BLK_RANK:        row[48],
					BLKA_RANK:       row[49],
					PF_RANK:         row[50],
					PFD_RANK:        row[51],
					PTS_RANK:        row[52],
					PLUS_MINUS_RANK: row[53],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.PrePostAllStarTeamDashboard = make([]TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard, len(rawResp.ResultSets[4].RowSet))
		for i, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 54 {
				response.PrePostAllStarTeamDashboard[i] = TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard{
					GROUP_SET:       row[0],
					GROUP_VALUE:     row[1],
					GP:              row[2],
					W:               row[3],
					L:               row[4],
					W_PCT:           row[5],
					MIN:             row[6],
					FGM:             row[7],
					FGA:             row[8],
					FG_PCT:          row[9],
					FG3M:            row[10],
					FG3A:            row[11],
					FG3_PCT:         row[12],
					FTM:             row[13],
					FTA:             row[14],
					FT_PCT:          row[15],
					OREB:            row[16],
					DREB:            row[17],
					REB:             row[18],
					AST:             row[19],
					TOV:             row[20],
					STL:             row[21],
					BLK:             row[22],
					BLKA:            row[23],
					PF:              row[24],
					PFD:             row[25],
					PTS:             row[26],
					PLUS_MINUS:      row[27],
					GP_RANK:         row[28],
					W_RANK:          row[29],
					L_RANK:          row[30],
					W_PCT_RANK:      row[31],
					MIN_RANK:        row[32],
					FGM_RANK:        row[33],
					FGA_RANK:        row[34],
					FG_PCT_RANK:     row[35],
					FG3M_RANK:       row[36],
					FG3A_RANK:       row[37],
					FG3_PCT_RANK:    row[38],
					FTM_RANK:        row[39],
					FTA_RANK:        row[40],
					FT_PCT_RANK:     row[41],
					OREB_RANK:       row[42],
					DREB_RANK:       row[43],
					REB_RANK:        row[44],
					AST_RANK:        row[45],
					TOV_RANK:        row[46],
					STL_RANK:        row[47],
					BLK_RANK:        row[48],
					BLKA_RANK:       row[49],
					PF_RANK:         row[50],
					PFD_RANK:        row[51],
					PTS_RANK:        row[52],
					PLUS_MINUS_RANK: row[53],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
