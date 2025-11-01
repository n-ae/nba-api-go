package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerDashboardByGeneralSplitsRequest contains parameters for the PlayerDashboardByGeneralSplits endpoint
type PlayerDashboardByGeneralSplitsRequest struct {
	PlayerID       string
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

// PlayerDashboardByGeneralSplitsOverallPlayerDashboard represents the OverallPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsOverallPlayerDashboard struct {
	GROUP_SET            interface{}
	GROUP_VALUE          interface{}
	GP                   interface{}
	W                    interface{}
	L                    interface{}
	W_PCT                interface{}
	MIN                  interface{}
	FGM                  interface{}
	FGA                  interface{}
	FG_PCT               interface{}
	FG3M                 interface{}
	FG3A                 interface{}
	FG3_PCT              interface{}
	FTM                  interface{}
	FTA                  interface{}
	FT_PCT               interface{}
	OREB                 interface{}
	DREB                 interface{}
	REB                  interface{}
	AST                  interface{}
	TOV                  interface{}
	STL                  interface{}
	BLK                  interface{}
	BLKA                 interface{}
	PF                   interface{}
	PFD                  interface{}
	PTS                  interface{}
	PLUS_MINUS           interface{}
	NBA_FANTASY_PTS      interface{}
	DD2                  interface{}
	TD3                  interface{}
	GP_RANK              interface{}
	W_RANK               interface{}
	L_RANK               interface{}
	W_PCT_RANK           interface{}
	MIN_RANK             interface{}
	FGM_RANK             interface{}
	FGA_RANK             interface{}
	FG_PCT_RANK          interface{}
	FG3M_RANK            interface{}
	FG3A_RANK            interface{}
	FG3_PCT_RANK         interface{}
	FTM_RANK             interface{}
	FTA_RANK             interface{}
	FT_PCT_RANK          interface{}
	OREB_RANK            interface{}
	DREB_RANK            interface{}
	REB_RANK             interface{}
	AST_RANK             interface{}
	TOV_RANK             interface{}
	STL_RANK             interface{}
	BLK_RANK             interface{}
	BLKA_RANK            interface{}
	PF_RANK              interface{}
	PFD_RANK             interface{}
	PTS_RANK             interface{}
	PLUS_MINUS_RANK      interface{}
	NBA_FANTASY_PTS_RANK interface{}
	DD2_RANK             interface{}
	TD3_RANK             interface{}
}

// PlayerDashboardByGeneralSplitsLocationPlayerDashboard represents the LocationPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsLocationPlayerDashboard struct {
	GROUP_SET            interface{}
	GROUP_VALUE          interface{}
	GP                   interface{}
	W                    interface{}
	L                    interface{}
	W_PCT                interface{}
	MIN                  interface{}
	FGM                  interface{}
	FGA                  interface{}
	FG_PCT               interface{}
	FG3M                 interface{}
	FG3A                 interface{}
	FG3_PCT              interface{}
	FTM                  interface{}
	FTA                  interface{}
	FT_PCT               interface{}
	OREB                 interface{}
	DREB                 interface{}
	REB                  interface{}
	AST                  interface{}
	TOV                  interface{}
	STL                  interface{}
	BLK                  interface{}
	BLKA                 interface{}
	PF                   interface{}
	PFD                  interface{}
	PTS                  interface{}
	PLUS_MINUS           interface{}
	NBA_FANTASY_PTS      interface{}
	DD2                  interface{}
	TD3                  interface{}
	GP_RANK              interface{}
	W_RANK               interface{}
	L_RANK               interface{}
	W_PCT_RANK           interface{}
	MIN_RANK             interface{}
	FGM_RANK             interface{}
	FGA_RANK             interface{}
	FG_PCT_RANK          interface{}
	FG3M_RANK            interface{}
	FG3A_RANK            interface{}
	FG3_PCT_RANK         interface{}
	FTM_RANK             interface{}
	FTA_RANK             interface{}
	FT_PCT_RANK          interface{}
	OREB_RANK            interface{}
	DREB_RANK            interface{}
	REB_RANK             interface{}
	AST_RANK             interface{}
	TOV_RANK             interface{}
	STL_RANK             interface{}
	BLK_RANK             interface{}
	BLKA_RANK            interface{}
	PF_RANK              interface{}
	PFD_RANK             interface{}
	PTS_RANK             interface{}
	PLUS_MINUS_RANK      interface{}
	NBA_FANTASY_PTS_RANK interface{}
	DD2_RANK             interface{}
	TD3_RANK             interface{}
}

// PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard represents the WinsLossesPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard struct {
	GROUP_SET            interface{}
	GROUP_VALUE          interface{}
	GP                   interface{}
	W                    interface{}
	L                    interface{}
	W_PCT                interface{}
	MIN                  interface{}
	FGM                  interface{}
	FGA                  interface{}
	FG_PCT               interface{}
	FG3M                 interface{}
	FG3A                 interface{}
	FG3_PCT              interface{}
	FTM                  interface{}
	FTA                  interface{}
	FT_PCT               interface{}
	OREB                 interface{}
	DREB                 interface{}
	REB                  interface{}
	AST                  interface{}
	TOV                  interface{}
	STL                  interface{}
	BLK                  interface{}
	BLKA                 interface{}
	PF                   interface{}
	PFD                  interface{}
	PTS                  interface{}
	PLUS_MINUS           interface{}
	NBA_FANTASY_PTS      interface{}
	DD2                  interface{}
	TD3                  interface{}
	GP_RANK              interface{}
	W_RANK               interface{}
	L_RANK               interface{}
	W_PCT_RANK           interface{}
	MIN_RANK             interface{}
	FGM_RANK             interface{}
	FGA_RANK             interface{}
	FG_PCT_RANK          interface{}
	FG3M_RANK            interface{}
	FG3A_RANK            interface{}
	FG3_PCT_RANK         interface{}
	FTM_RANK             interface{}
	FTA_RANK             interface{}
	FT_PCT_RANK          interface{}
	OREB_RANK            interface{}
	DREB_RANK            interface{}
	REB_RANK             interface{}
	AST_RANK             interface{}
	TOV_RANK             interface{}
	STL_RANK             interface{}
	BLK_RANK             interface{}
	BLKA_RANK            interface{}
	PF_RANK              interface{}
	PFD_RANK             interface{}
	PTS_RANK             interface{}
	PLUS_MINUS_RANK      interface{}
	NBA_FANTASY_PTS_RANK interface{}
	DD2_RANK             interface{}
	TD3_RANK             interface{}
}

// PlayerDashboardByGeneralSplitsMonthPlayerDashboard represents the MonthPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsMonthPlayerDashboard struct {
	GROUP_SET            interface{}
	GROUP_VALUE          interface{}
	GP                   interface{}
	W                    interface{}
	L                    interface{}
	W_PCT                interface{}
	MIN                  interface{}
	FGM                  interface{}
	FGA                  interface{}
	FG_PCT               interface{}
	FG3M                 interface{}
	FG3A                 interface{}
	FG3_PCT              interface{}
	FTM                  interface{}
	FTA                  interface{}
	FT_PCT               interface{}
	OREB                 interface{}
	DREB                 interface{}
	REB                  interface{}
	AST                  interface{}
	TOV                  interface{}
	STL                  interface{}
	BLK                  interface{}
	BLKA                 interface{}
	PF                   interface{}
	PFD                  interface{}
	PTS                  interface{}
	PLUS_MINUS           interface{}
	NBA_FANTASY_PTS      interface{}
	DD2                  interface{}
	TD3                  interface{}
	GP_RANK              interface{}
	W_RANK               interface{}
	L_RANK               interface{}
	W_PCT_RANK           interface{}
	MIN_RANK             interface{}
	FGM_RANK             interface{}
	FGA_RANK             interface{}
	FG_PCT_RANK          interface{}
	FG3M_RANK            interface{}
	FG3A_RANK            interface{}
	FG3_PCT_RANK         interface{}
	FTM_RANK             interface{}
	FTA_RANK             interface{}
	FT_PCT_RANK          interface{}
	OREB_RANK            interface{}
	DREB_RANK            interface{}
	REB_RANK             interface{}
	AST_RANK             interface{}
	TOV_RANK             interface{}
	STL_RANK             interface{}
	BLK_RANK             interface{}
	BLKA_RANK            interface{}
	PF_RANK              interface{}
	PFD_RANK             interface{}
	PTS_RANK             interface{}
	PLUS_MINUS_RANK      interface{}
	NBA_FANTASY_PTS_RANK interface{}
	DD2_RANK             interface{}
	TD3_RANK             interface{}
}

// PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard represents the PrePostAllStarPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard struct {
	GROUP_SET            interface{}
	GROUP_VALUE          interface{}
	GP                   interface{}
	W                    interface{}
	L                    interface{}
	W_PCT                interface{}
	MIN                  interface{}
	FGM                  interface{}
	FGA                  interface{}
	FG_PCT               interface{}
	FG3M                 interface{}
	FG3A                 interface{}
	FG3_PCT              interface{}
	FTM                  interface{}
	FTA                  interface{}
	FT_PCT               interface{}
	OREB                 interface{}
	DREB                 interface{}
	REB                  interface{}
	AST                  interface{}
	TOV                  interface{}
	STL                  interface{}
	BLK                  interface{}
	BLKA                 interface{}
	PF                   interface{}
	PFD                  interface{}
	PTS                  interface{}
	PLUS_MINUS           interface{}
	NBA_FANTASY_PTS      interface{}
	DD2                  interface{}
	TD3                  interface{}
	GP_RANK              interface{}
	W_RANK               interface{}
	L_RANK               interface{}
	W_PCT_RANK           interface{}
	MIN_RANK             interface{}
	FGM_RANK             interface{}
	FGA_RANK             interface{}
	FG_PCT_RANK          interface{}
	FG3M_RANK            interface{}
	FG3A_RANK            interface{}
	FG3_PCT_RANK         interface{}
	FTM_RANK             interface{}
	FTA_RANK             interface{}
	FT_PCT_RANK          interface{}
	OREB_RANK            interface{}
	DREB_RANK            interface{}
	REB_RANK             interface{}
	AST_RANK             interface{}
	TOV_RANK             interface{}
	STL_RANK             interface{}
	BLK_RANK             interface{}
	BLKA_RANK            interface{}
	PF_RANK              interface{}
	PFD_RANK             interface{}
	PTS_RANK             interface{}
	PLUS_MINUS_RANK      interface{}
	NBA_FANTASY_PTS_RANK interface{}
	DD2_RANK             interface{}
	TD3_RANK             interface{}
}

// PlayerDashboardByGeneralSplitsResponse contains the response data from the PlayerDashboardByGeneralSplits endpoint
type PlayerDashboardByGeneralSplitsResponse struct {
	OverallPlayerDashboard        []PlayerDashboardByGeneralSplitsOverallPlayerDashboard
	LocationPlayerDashboard       []PlayerDashboardByGeneralSplitsLocationPlayerDashboard
	WinsLossesPlayerDashboard     []PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard
	MonthPlayerDashboard          []PlayerDashboardByGeneralSplitsMonthPlayerDashboard
	PrePostAllStarPlayerDashboard []PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard
}

// GetPlayerDashboardByGeneralSplits retrieves data from the playerdashboardbygeneralsplits endpoint
func GetPlayerDashboardByGeneralSplits(ctx context.Context, client *stats.Client, req PlayerDashboardByGeneralSplitsRequest) (*models.Response[*PlayerDashboardByGeneralSplitsResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
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
	if err := client.GetJSON(ctx, "/playerdashboardbygeneralsplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerDashboardByGeneralSplitsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallPlayerDashboard = make([]PlayerDashboardByGeneralSplitsOverallPlayerDashboard, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 60 {
				response.OverallPlayerDashboard[i] = PlayerDashboardByGeneralSplitsOverallPlayerDashboard{
					GROUP_SET:            row[0],
					GROUP_VALUE:          row[1],
					GP:                   row[2],
					W:                    row[3],
					L:                    row[4],
					W_PCT:                row[5],
					MIN:                  row[6],
					FGM:                  row[7],
					FGA:                  row[8],
					FG_PCT:               row[9],
					FG3M:                 row[10],
					FG3A:                 row[11],
					FG3_PCT:              row[12],
					FTM:                  row[13],
					FTA:                  row[14],
					FT_PCT:               row[15],
					OREB:                 row[16],
					DREB:                 row[17],
					REB:                  row[18],
					AST:                  row[19],
					TOV:                  row[20],
					STL:                  row[21],
					BLK:                  row[22],
					BLKA:                 row[23],
					PF:                   row[24],
					PFD:                  row[25],
					PTS:                  row[26],
					PLUS_MINUS:           row[27],
					NBA_FANTASY_PTS:      row[28],
					DD2:                  row[29],
					TD3:                  row[30],
					GP_RANK:              row[31],
					W_RANK:               row[32],
					L_RANK:               row[33],
					W_PCT_RANK:           row[34],
					MIN_RANK:             row[35],
					FGM_RANK:             row[36],
					FGA_RANK:             row[37],
					FG_PCT_RANK:          row[38],
					FG3M_RANK:            row[39],
					FG3A_RANK:            row[40],
					FG3_PCT_RANK:         row[41],
					FTM_RANK:             row[42],
					FTA_RANK:             row[43],
					FT_PCT_RANK:          row[44],
					OREB_RANK:            row[45],
					DREB_RANK:            row[46],
					REB_RANK:             row[47],
					AST_RANK:             row[48],
					TOV_RANK:             row[49],
					STL_RANK:             row[50],
					BLK_RANK:             row[51],
					BLKA_RANK:            row[52],
					PF_RANK:              row[53],
					PFD_RANK:             row[54],
					PTS_RANK:             row[55],
					PLUS_MINUS_RANK:      row[56],
					NBA_FANTASY_PTS_RANK: row[57],
					DD2_RANK:             row[58],
					TD3_RANK:             row[59],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.LocationPlayerDashboard = make([]PlayerDashboardByGeneralSplitsLocationPlayerDashboard, len(rawResp.ResultSets[1].RowSet))
		for i, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 60 {
				response.LocationPlayerDashboard[i] = PlayerDashboardByGeneralSplitsLocationPlayerDashboard{
					GROUP_SET:            row[0],
					GROUP_VALUE:          row[1],
					GP:                   row[2],
					W:                    row[3],
					L:                    row[4],
					W_PCT:                row[5],
					MIN:                  row[6],
					FGM:                  row[7],
					FGA:                  row[8],
					FG_PCT:               row[9],
					FG3M:                 row[10],
					FG3A:                 row[11],
					FG3_PCT:              row[12],
					FTM:                  row[13],
					FTA:                  row[14],
					FT_PCT:               row[15],
					OREB:                 row[16],
					DREB:                 row[17],
					REB:                  row[18],
					AST:                  row[19],
					TOV:                  row[20],
					STL:                  row[21],
					BLK:                  row[22],
					BLKA:                 row[23],
					PF:                   row[24],
					PFD:                  row[25],
					PTS:                  row[26],
					PLUS_MINUS:           row[27],
					NBA_FANTASY_PTS:      row[28],
					DD2:                  row[29],
					TD3:                  row[30],
					GP_RANK:              row[31],
					W_RANK:               row[32],
					L_RANK:               row[33],
					W_PCT_RANK:           row[34],
					MIN_RANK:             row[35],
					FGM_RANK:             row[36],
					FGA_RANK:             row[37],
					FG_PCT_RANK:          row[38],
					FG3M_RANK:            row[39],
					FG3A_RANK:            row[40],
					FG3_PCT_RANK:         row[41],
					FTM_RANK:             row[42],
					FTA_RANK:             row[43],
					FT_PCT_RANK:          row[44],
					OREB_RANK:            row[45],
					DREB_RANK:            row[46],
					REB_RANK:             row[47],
					AST_RANK:             row[48],
					TOV_RANK:             row[49],
					STL_RANK:             row[50],
					BLK_RANK:             row[51],
					BLKA_RANK:            row[52],
					PF_RANK:              row[53],
					PFD_RANK:             row[54],
					PTS_RANK:             row[55],
					PLUS_MINUS_RANK:      row[56],
					NBA_FANTASY_PTS_RANK: row[57],
					DD2_RANK:             row[58],
					TD3_RANK:             row[59],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.WinsLossesPlayerDashboard = make([]PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard, len(rawResp.ResultSets[2].RowSet))
		for i, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 60 {
				response.WinsLossesPlayerDashboard[i] = PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard{
					GROUP_SET:            row[0],
					GROUP_VALUE:          row[1],
					GP:                   row[2],
					W:                    row[3],
					L:                    row[4],
					W_PCT:                row[5],
					MIN:                  row[6],
					FGM:                  row[7],
					FGA:                  row[8],
					FG_PCT:               row[9],
					FG3M:                 row[10],
					FG3A:                 row[11],
					FG3_PCT:              row[12],
					FTM:                  row[13],
					FTA:                  row[14],
					FT_PCT:               row[15],
					OREB:                 row[16],
					DREB:                 row[17],
					REB:                  row[18],
					AST:                  row[19],
					TOV:                  row[20],
					STL:                  row[21],
					BLK:                  row[22],
					BLKA:                 row[23],
					PF:                   row[24],
					PFD:                  row[25],
					PTS:                  row[26],
					PLUS_MINUS:           row[27],
					NBA_FANTASY_PTS:      row[28],
					DD2:                  row[29],
					TD3:                  row[30],
					GP_RANK:              row[31],
					W_RANK:               row[32],
					L_RANK:               row[33],
					W_PCT_RANK:           row[34],
					MIN_RANK:             row[35],
					FGM_RANK:             row[36],
					FGA_RANK:             row[37],
					FG_PCT_RANK:          row[38],
					FG3M_RANK:            row[39],
					FG3A_RANK:            row[40],
					FG3_PCT_RANK:         row[41],
					FTM_RANK:             row[42],
					FTA_RANK:             row[43],
					FT_PCT_RANK:          row[44],
					OREB_RANK:            row[45],
					DREB_RANK:            row[46],
					REB_RANK:             row[47],
					AST_RANK:             row[48],
					TOV_RANK:             row[49],
					STL_RANK:             row[50],
					BLK_RANK:             row[51],
					BLKA_RANK:            row[52],
					PF_RANK:              row[53],
					PFD_RANK:             row[54],
					PTS_RANK:             row[55],
					PLUS_MINUS_RANK:      row[56],
					NBA_FANTASY_PTS_RANK: row[57],
					DD2_RANK:             row[58],
					TD3_RANK:             row[59],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.MonthPlayerDashboard = make([]PlayerDashboardByGeneralSplitsMonthPlayerDashboard, len(rawResp.ResultSets[3].RowSet))
		for i, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 60 {
				response.MonthPlayerDashboard[i] = PlayerDashboardByGeneralSplitsMonthPlayerDashboard{
					GROUP_SET:            row[0],
					GROUP_VALUE:          row[1],
					GP:                   row[2],
					W:                    row[3],
					L:                    row[4],
					W_PCT:                row[5],
					MIN:                  row[6],
					FGM:                  row[7],
					FGA:                  row[8],
					FG_PCT:               row[9],
					FG3M:                 row[10],
					FG3A:                 row[11],
					FG3_PCT:              row[12],
					FTM:                  row[13],
					FTA:                  row[14],
					FT_PCT:               row[15],
					OREB:                 row[16],
					DREB:                 row[17],
					REB:                  row[18],
					AST:                  row[19],
					TOV:                  row[20],
					STL:                  row[21],
					BLK:                  row[22],
					BLKA:                 row[23],
					PF:                   row[24],
					PFD:                  row[25],
					PTS:                  row[26],
					PLUS_MINUS:           row[27],
					NBA_FANTASY_PTS:      row[28],
					DD2:                  row[29],
					TD3:                  row[30],
					GP_RANK:              row[31],
					W_RANK:               row[32],
					L_RANK:               row[33],
					W_PCT_RANK:           row[34],
					MIN_RANK:             row[35],
					FGM_RANK:             row[36],
					FGA_RANK:             row[37],
					FG_PCT_RANK:          row[38],
					FG3M_RANK:            row[39],
					FG3A_RANK:            row[40],
					FG3_PCT_RANK:         row[41],
					FTM_RANK:             row[42],
					FTA_RANK:             row[43],
					FT_PCT_RANK:          row[44],
					OREB_RANK:            row[45],
					DREB_RANK:            row[46],
					REB_RANK:             row[47],
					AST_RANK:             row[48],
					TOV_RANK:             row[49],
					STL_RANK:             row[50],
					BLK_RANK:             row[51],
					BLKA_RANK:            row[52],
					PF_RANK:              row[53],
					PFD_RANK:             row[54],
					PTS_RANK:             row[55],
					PLUS_MINUS_RANK:      row[56],
					NBA_FANTASY_PTS_RANK: row[57],
					DD2_RANK:             row[58],
					TD3_RANK:             row[59],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.PrePostAllStarPlayerDashboard = make([]PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard, len(rawResp.ResultSets[4].RowSet))
		for i, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 60 {
				response.PrePostAllStarPlayerDashboard[i] = PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard{
					GROUP_SET:            row[0],
					GROUP_VALUE:          row[1],
					GP:                   row[2],
					W:                    row[3],
					L:                    row[4],
					W_PCT:                row[5],
					MIN:                  row[6],
					FGM:                  row[7],
					FGA:                  row[8],
					FG_PCT:               row[9],
					FG3M:                 row[10],
					FG3A:                 row[11],
					FG3_PCT:              row[12],
					FTM:                  row[13],
					FTA:                  row[14],
					FT_PCT:               row[15],
					OREB:                 row[16],
					DREB:                 row[17],
					REB:                  row[18],
					AST:                  row[19],
					TOV:                  row[20],
					STL:                  row[21],
					BLK:                  row[22],
					BLKA:                 row[23],
					PF:                   row[24],
					PFD:                  row[25],
					PTS:                  row[26],
					PLUS_MINUS:           row[27],
					NBA_FANTASY_PTS:      row[28],
					DD2:                  row[29],
					TD3:                  row[30],
					GP_RANK:              row[31],
					W_RANK:               row[32],
					L_RANK:               row[33],
					W_PCT_RANK:           row[34],
					MIN_RANK:             row[35],
					FGM_RANK:             row[36],
					FGA_RANK:             row[37],
					FG_PCT_RANK:          row[38],
					FG3M_RANK:            row[39],
					FG3A_RANK:            row[40],
					FG3_PCT_RANK:         row[41],
					FTM_RANK:             row[42],
					FTA_RANK:             row[43],
					FT_PCT_RANK:          row[44],
					OREB_RANK:            row[45],
					DREB_RANK:            row[46],
					REB_RANK:             row[47],
					AST_RANK:             row[48],
					TOV_RANK:             row[49],
					STL_RANK:             row[50],
					BLK_RANK:             row[51],
					BLKA_RANK:            row[52],
					PF_RANK:              row[53],
					PFD_RANK:             row[54],
					PTS_RANK:             row[55],
					PLUS_MINUS_RANK:      row[56],
					NBA_FANTASY_PTS_RANK: row[57],
					DD2_RANK:             row[58],
					TD3_RANK:             row[59],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
