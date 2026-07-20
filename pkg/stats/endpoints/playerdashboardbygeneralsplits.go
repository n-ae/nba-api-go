package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
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
	GROUP_SET            string  `json:"GROUP_SET"`
	GROUP_VALUE          string  `json:"GROUP_VALUE"`
	GP                   int     `json:"GP"`
	W                    string  `json:"W"`
	L                    string  `json:"L"`
	W_PCT                float64 `json:"W_PCT"`
	MIN                  float64 `json:"MIN"`
	FGM                  int     `json:"FGM"`
	FGA                  int     `json:"FGA"`
	FG_PCT               float64 `json:"FG_PCT"`
	FG3M                 int     `json:"FG3M"`
	FG3A                 int     `json:"FG3A"`
	FG3_PCT              float64 `json:"FG3_PCT"`
	FTM                  int     `json:"FTM"`
	FTA                  int     `json:"FTA"`
	FT_PCT               float64 `json:"FT_PCT"`
	OREB                 float64 `json:"OREB"`
	DREB                 float64 `json:"DREB"`
	REB                  float64 `json:"REB"`
	AST                  float64 `json:"AST"`
	TOV                  float64 `json:"TOV"`
	STL                  float64 `json:"STL"`
	BLK                  float64 `json:"BLK"`
	BLKA                 int     `json:"BLKA"`
	PF                   float64 `json:"PF"`
	PFD                  float64 `json:"PFD"`
	PTS                  float64 `json:"PTS"`
	PLUS_MINUS           float64 `json:"PLUS_MINUS"`
	NBA_FANTASY_PTS      float64 `json:"NBA_FANTASY_PTS"`
	DD2                  float64 `json:"DD2"`
	TD3                  float64 `json:"TD3"`
	GP_RANK              float64 `json:"GP_RANK"`
	W_RANK               float64 `json:"W_RANK"`
	L_RANK               float64 `json:"L_RANK"`
	W_PCT_RANK           float64 `json:"W_PCT_RANK"`
	MIN_RANK             float64 `json:"MIN_RANK"`
	FGM_RANK             float64 `json:"FGM_RANK"`
	FGA_RANK             float64 `json:"FGA_RANK"`
	FG_PCT_RANK          float64 `json:"FG_PCT_RANK"`
	FG3M_RANK            float64 `json:"FG3M_RANK"`
	FG3A_RANK            float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK         float64 `json:"FG3_PCT_RANK"`
	FTM_RANK             float64 `json:"FTM_RANK"`
	FTA_RANK             float64 `json:"FTA_RANK"`
	FT_PCT_RANK          float64 `json:"FT_PCT_RANK"`
	OREB_RANK            float64 `json:"OREB_RANK"`
	DREB_RANK            float64 `json:"DREB_RANK"`
	REB_RANK             float64 `json:"REB_RANK"`
	AST_RANK             float64 `json:"AST_RANK"`
	TOV_RANK             float64 `json:"TOV_RANK"`
	STL_RANK             float64 `json:"STL_RANK"`
	BLK_RANK             float64 `json:"BLK_RANK"`
	BLKA_RANK            float64 `json:"BLKA_RANK"`
	PF_RANK              float64 `json:"PF_RANK"`
	PFD_RANK             float64 `json:"PFD_RANK"`
	PTS_RANK             float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK      float64 `json:"PLUS_MINUS_RANK"`
	NBA_FANTASY_PTS_RANK float64 `json:"NBA_FANTASY_PTS_RANK"`
	DD2_RANK             float64 `json:"DD2_RANK"`
	TD3_RANK             float64 `json:"TD3_RANK"`
}

// PlayerDashboardByGeneralSplitsLocationPlayerDashboard represents the LocationPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsLocationPlayerDashboard struct {
	GROUP_SET            string  `json:"GROUP_SET"`
	GROUP_VALUE          string  `json:"GROUP_VALUE"`
	GP                   int     `json:"GP"`
	W                    string  `json:"W"`
	L                    string  `json:"L"`
	W_PCT                float64 `json:"W_PCT"`
	MIN                  float64 `json:"MIN"`
	FGM                  int     `json:"FGM"`
	FGA                  int     `json:"FGA"`
	FG_PCT               float64 `json:"FG_PCT"`
	FG3M                 int     `json:"FG3M"`
	FG3A                 int     `json:"FG3A"`
	FG3_PCT              float64 `json:"FG3_PCT"`
	FTM                  int     `json:"FTM"`
	FTA                  int     `json:"FTA"`
	FT_PCT               float64 `json:"FT_PCT"`
	OREB                 float64 `json:"OREB"`
	DREB                 float64 `json:"DREB"`
	REB                  float64 `json:"REB"`
	AST                  float64 `json:"AST"`
	TOV                  float64 `json:"TOV"`
	STL                  float64 `json:"STL"`
	BLK                  float64 `json:"BLK"`
	BLKA                 int     `json:"BLKA"`
	PF                   float64 `json:"PF"`
	PFD                  float64 `json:"PFD"`
	PTS                  float64 `json:"PTS"`
	PLUS_MINUS           float64 `json:"PLUS_MINUS"`
	NBA_FANTASY_PTS      float64 `json:"NBA_FANTASY_PTS"`
	DD2                  float64 `json:"DD2"`
	TD3                  float64 `json:"TD3"`
	GP_RANK              float64 `json:"GP_RANK"`
	W_RANK               float64 `json:"W_RANK"`
	L_RANK               float64 `json:"L_RANK"`
	W_PCT_RANK           float64 `json:"W_PCT_RANK"`
	MIN_RANK             float64 `json:"MIN_RANK"`
	FGM_RANK             float64 `json:"FGM_RANK"`
	FGA_RANK             float64 `json:"FGA_RANK"`
	FG_PCT_RANK          float64 `json:"FG_PCT_RANK"`
	FG3M_RANK            float64 `json:"FG3M_RANK"`
	FG3A_RANK            float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK         float64 `json:"FG3_PCT_RANK"`
	FTM_RANK             float64 `json:"FTM_RANK"`
	FTA_RANK             float64 `json:"FTA_RANK"`
	FT_PCT_RANK          float64 `json:"FT_PCT_RANK"`
	OREB_RANK            float64 `json:"OREB_RANK"`
	DREB_RANK            float64 `json:"DREB_RANK"`
	REB_RANK             float64 `json:"REB_RANK"`
	AST_RANK             float64 `json:"AST_RANK"`
	TOV_RANK             float64 `json:"TOV_RANK"`
	STL_RANK             float64 `json:"STL_RANK"`
	BLK_RANK             float64 `json:"BLK_RANK"`
	BLKA_RANK            float64 `json:"BLKA_RANK"`
	PF_RANK              float64 `json:"PF_RANK"`
	PFD_RANK             float64 `json:"PFD_RANK"`
	PTS_RANK             float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK      float64 `json:"PLUS_MINUS_RANK"`
	NBA_FANTASY_PTS_RANK float64 `json:"NBA_FANTASY_PTS_RANK"`
	DD2_RANK             float64 `json:"DD2_RANK"`
	TD3_RANK             float64 `json:"TD3_RANK"`
}

// PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard represents the WinsLossesPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard struct {
	GROUP_SET            string  `json:"GROUP_SET"`
	GROUP_VALUE          string  `json:"GROUP_VALUE"`
	GP                   int     `json:"GP"`
	W                    string  `json:"W"`
	L                    string  `json:"L"`
	W_PCT                float64 `json:"W_PCT"`
	MIN                  float64 `json:"MIN"`
	FGM                  int     `json:"FGM"`
	FGA                  int     `json:"FGA"`
	FG_PCT               float64 `json:"FG_PCT"`
	FG3M                 int     `json:"FG3M"`
	FG3A                 int     `json:"FG3A"`
	FG3_PCT              float64 `json:"FG3_PCT"`
	FTM                  int     `json:"FTM"`
	FTA                  int     `json:"FTA"`
	FT_PCT               float64 `json:"FT_PCT"`
	OREB                 float64 `json:"OREB"`
	DREB                 float64 `json:"DREB"`
	REB                  float64 `json:"REB"`
	AST                  float64 `json:"AST"`
	TOV                  float64 `json:"TOV"`
	STL                  float64 `json:"STL"`
	BLK                  float64 `json:"BLK"`
	BLKA                 int     `json:"BLKA"`
	PF                   float64 `json:"PF"`
	PFD                  float64 `json:"PFD"`
	PTS                  float64 `json:"PTS"`
	PLUS_MINUS           float64 `json:"PLUS_MINUS"`
	NBA_FANTASY_PTS      float64 `json:"NBA_FANTASY_PTS"`
	DD2                  float64 `json:"DD2"`
	TD3                  float64 `json:"TD3"`
	GP_RANK              float64 `json:"GP_RANK"`
	W_RANK               float64 `json:"W_RANK"`
	L_RANK               float64 `json:"L_RANK"`
	W_PCT_RANK           float64 `json:"W_PCT_RANK"`
	MIN_RANK             float64 `json:"MIN_RANK"`
	FGM_RANK             float64 `json:"FGM_RANK"`
	FGA_RANK             float64 `json:"FGA_RANK"`
	FG_PCT_RANK          float64 `json:"FG_PCT_RANK"`
	FG3M_RANK            float64 `json:"FG3M_RANK"`
	FG3A_RANK            float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK         float64 `json:"FG3_PCT_RANK"`
	FTM_RANK             float64 `json:"FTM_RANK"`
	FTA_RANK             float64 `json:"FTA_RANK"`
	FT_PCT_RANK          float64 `json:"FT_PCT_RANK"`
	OREB_RANK            float64 `json:"OREB_RANK"`
	DREB_RANK            float64 `json:"DREB_RANK"`
	REB_RANK             float64 `json:"REB_RANK"`
	AST_RANK             float64 `json:"AST_RANK"`
	TOV_RANK             float64 `json:"TOV_RANK"`
	STL_RANK             float64 `json:"STL_RANK"`
	BLK_RANK             float64 `json:"BLK_RANK"`
	BLKA_RANK            float64 `json:"BLKA_RANK"`
	PF_RANK              float64 `json:"PF_RANK"`
	PFD_RANK             float64 `json:"PFD_RANK"`
	PTS_RANK             float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK      float64 `json:"PLUS_MINUS_RANK"`
	NBA_FANTASY_PTS_RANK float64 `json:"NBA_FANTASY_PTS_RANK"`
	DD2_RANK             float64 `json:"DD2_RANK"`
	TD3_RANK             float64 `json:"TD3_RANK"`
}

// PlayerDashboardByGeneralSplitsMonthPlayerDashboard represents the MonthPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsMonthPlayerDashboard struct {
	GROUP_SET            string  `json:"GROUP_SET"`
	GROUP_VALUE          string  `json:"GROUP_VALUE"`
	GP                   int     `json:"GP"`
	W                    string  `json:"W"`
	L                    string  `json:"L"`
	W_PCT                float64 `json:"W_PCT"`
	MIN                  float64 `json:"MIN"`
	FGM                  int     `json:"FGM"`
	FGA                  int     `json:"FGA"`
	FG_PCT               float64 `json:"FG_PCT"`
	FG3M                 int     `json:"FG3M"`
	FG3A                 int     `json:"FG3A"`
	FG3_PCT              float64 `json:"FG3_PCT"`
	FTM                  int     `json:"FTM"`
	FTA                  int     `json:"FTA"`
	FT_PCT               float64 `json:"FT_PCT"`
	OREB                 float64 `json:"OREB"`
	DREB                 float64 `json:"DREB"`
	REB                  float64 `json:"REB"`
	AST                  float64 `json:"AST"`
	TOV                  float64 `json:"TOV"`
	STL                  float64 `json:"STL"`
	BLK                  float64 `json:"BLK"`
	BLKA                 int     `json:"BLKA"`
	PF                   float64 `json:"PF"`
	PFD                  float64 `json:"PFD"`
	PTS                  float64 `json:"PTS"`
	PLUS_MINUS           float64 `json:"PLUS_MINUS"`
	NBA_FANTASY_PTS      float64 `json:"NBA_FANTASY_PTS"`
	DD2                  float64 `json:"DD2"`
	TD3                  float64 `json:"TD3"`
	GP_RANK              float64 `json:"GP_RANK"`
	W_RANK               float64 `json:"W_RANK"`
	L_RANK               float64 `json:"L_RANK"`
	W_PCT_RANK           float64 `json:"W_PCT_RANK"`
	MIN_RANK             float64 `json:"MIN_RANK"`
	FGM_RANK             float64 `json:"FGM_RANK"`
	FGA_RANK             float64 `json:"FGA_RANK"`
	FG_PCT_RANK          float64 `json:"FG_PCT_RANK"`
	FG3M_RANK            float64 `json:"FG3M_RANK"`
	FG3A_RANK            float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK         float64 `json:"FG3_PCT_RANK"`
	FTM_RANK             float64 `json:"FTM_RANK"`
	FTA_RANK             float64 `json:"FTA_RANK"`
	FT_PCT_RANK          float64 `json:"FT_PCT_RANK"`
	OREB_RANK            float64 `json:"OREB_RANK"`
	DREB_RANK            float64 `json:"DREB_RANK"`
	REB_RANK             float64 `json:"REB_RANK"`
	AST_RANK             float64 `json:"AST_RANK"`
	TOV_RANK             float64 `json:"TOV_RANK"`
	STL_RANK             float64 `json:"STL_RANK"`
	BLK_RANK             float64 `json:"BLK_RANK"`
	BLKA_RANK            float64 `json:"BLKA_RANK"`
	PF_RANK              float64 `json:"PF_RANK"`
	PFD_RANK             float64 `json:"PFD_RANK"`
	PTS_RANK             float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK      float64 `json:"PLUS_MINUS_RANK"`
	NBA_FANTASY_PTS_RANK float64 `json:"NBA_FANTASY_PTS_RANK"`
	DD2_RANK             float64 `json:"DD2_RANK"`
	TD3_RANK             float64 `json:"TD3_RANK"`
}

// PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard represents the PrePostAllStarPlayerDashboard result set for PlayerDashboardByGeneralSplits
type PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard struct {
	GROUP_SET            string  `json:"GROUP_SET"`
	GROUP_VALUE          string  `json:"GROUP_VALUE"`
	GP                   int     `json:"GP"`
	W                    string  `json:"W"`
	L                    string  `json:"L"`
	W_PCT                float64 `json:"W_PCT"`
	MIN                  float64 `json:"MIN"`
	FGM                  int     `json:"FGM"`
	FGA                  int     `json:"FGA"`
	FG_PCT               float64 `json:"FG_PCT"`
	FG3M                 int     `json:"FG3M"`
	FG3A                 int     `json:"FG3A"`
	FG3_PCT              float64 `json:"FG3_PCT"`
	FTM                  int     `json:"FTM"`
	FTA                  int     `json:"FTA"`
	FT_PCT               float64 `json:"FT_PCT"`
	OREB                 float64 `json:"OREB"`
	DREB                 float64 `json:"DREB"`
	REB                  float64 `json:"REB"`
	AST                  float64 `json:"AST"`
	TOV                  float64 `json:"TOV"`
	STL                  float64 `json:"STL"`
	BLK                  float64 `json:"BLK"`
	BLKA                 int     `json:"BLKA"`
	PF                   float64 `json:"PF"`
	PFD                  float64 `json:"PFD"`
	PTS                  float64 `json:"PTS"`
	PLUS_MINUS           float64 `json:"PLUS_MINUS"`
	NBA_FANTASY_PTS      float64 `json:"NBA_FANTASY_PTS"`
	DD2                  float64 `json:"DD2"`
	TD3                  float64 `json:"TD3"`
	GP_RANK              float64 `json:"GP_RANK"`
	W_RANK               float64 `json:"W_RANK"`
	L_RANK               float64 `json:"L_RANK"`
	W_PCT_RANK           float64 `json:"W_PCT_RANK"`
	MIN_RANK             float64 `json:"MIN_RANK"`
	FGM_RANK             float64 `json:"FGM_RANK"`
	FGA_RANK             float64 `json:"FGA_RANK"`
	FG_PCT_RANK          float64 `json:"FG_PCT_RANK"`
	FG3M_RANK            float64 `json:"FG3M_RANK"`
	FG3A_RANK            float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK         float64 `json:"FG3_PCT_RANK"`
	FTM_RANK             float64 `json:"FTM_RANK"`
	FTA_RANK             float64 `json:"FTA_RANK"`
	FT_PCT_RANK          float64 `json:"FT_PCT_RANK"`
	OREB_RANK            float64 `json:"OREB_RANK"`
	DREB_RANK            float64 `json:"DREB_RANK"`
	REB_RANK             float64 `json:"REB_RANK"`
	AST_RANK             float64 `json:"AST_RANK"`
	TOV_RANK             float64 `json:"TOV_RANK"`
	STL_RANK             float64 `json:"STL_RANK"`
	BLK_RANK             float64 `json:"BLK_RANK"`
	BLKA_RANK            float64 `json:"BLKA_RANK"`
	PF_RANK              float64 `json:"PF_RANK"`
	PFD_RANK             float64 `json:"PFD_RANK"`
	PTS_RANK             float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK      float64 `json:"PLUS_MINUS_RANK"`
	NBA_FANTASY_PTS_RANK float64 `json:"NBA_FANTASY_PTS_RANK"`
	DD2_RANK             float64 `json:"DD2_RANK"`
	TD3_RANK             float64 `json:"TD3_RANK"`
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
		return nil, fmt.Errorf("%s is required", "PlayerID")
	}
	params.Set("PlayerID", req.PlayerID)
	if req.Season == "" {
		return nil, fmt.Errorf("%s is required", "Season")
	}
	params.Set("Season", string(req.Season))
	if req.SeasonType == "" {
		return nil, fmt.Errorf("%s is required", "SeasonType")
	}
	params.Set("SeasonType", string(req.SeasonType))
	if req.MeasureType != nil {
		params.Set("MeasureType", *req.MeasureType)
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.PlusMinus != nil {
		params.Set("PlusMinus", *req.PlusMinus)
	}
	if req.PaceAdjust != nil {
		params.Set("PaceAdjust", *req.PaceAdjust)
	}
	if req.Rank != nil {
		params.Set("Rank", *req.Rank)
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Outcome != nil {
		params.Set("Outcome", *req.Outcome)
	}
	if req.Location != nil {
		params.Set("Location", *req.Location)
	}
	if req.Month != nil {
		params.Set("Month", *req.Month)
	}
	if req.SeasonSegment != nil {
		params.Set("SeasonSegment", *req.SeasonSegment)
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", *req.DateFrom)
	}
	if req.DateTo != nil {
		params.Set("DateTo", *req.DateTo)
	}
	if req.OpponentTeamID != nil {
		params.Set("OpponentTeamID", *req.OpponentTeamID)
	}
	if req.VsConference != nil {
		params.Set("VsConference", *req.VsConference)
	}
	if req.VsDivision != nil {
		params.Set("VsDivision", *req.VsDivision)
	}
	if req.GameSegment != nil {
		params.Set("GameSegment", *req.GameSegment)
	}
	if req.Period != nil {
		params.Set("Period", *req.Period)
	}
	if req.LastNGames != nil {
		params.Set("LastNGames", *req.LastNGames)
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "playerdashboardbygeneralsplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerDashboardByGeneralSplitsResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "OverallPlayerDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayerDashboardByGeneralSplitsOverallPlayerDashboard{})); err != nil {
			return nil, fmt.Errorf("PlayerDashboardByGeneralSplits: OverallPlayerDashboard result set: %w", err)
		}
		response.OverallPlayerDashboard = make([]PlayerDashboardByGeneralSplitsOverallPlayerDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 60 {
				item := PlayerDashboardByGeneralSplitsOverallPlayerDashboard{
					GROUP_SET:            toString(row[0]),
					GROUP_VALUE:          toString(row[1]),
					GP:                   toInt(row[2]),
					W:                    toString(row[3]),
					L:                    toString(row[4]),
					W_PCT:                toFloat(row[5]),
					MIN:                  toFloat(row[6]),
					FGM:                  toInt(row[7]),
					FGA:                  toInt(row[8]),
					FG_PCT:               toFloat(row[9]),
					FG3M:                 toInt(row[10]),
					FG3A:                 toInt(row[11]),
					FG3_PCT:              toFloat(row[12]),
					FTM:                  toInt(row[13]),
					FTA:                  toInt(row[14]),
					FT_PCT:               toFloat(row[15]),
					OREB:                 toFloat(row[16]),
					DREB:                 toFloat(row[17]),
					REB:                  toFloat(row[18]),
					AST:                  toFloat(row[19]),
					TOV:                  toFloat(row[20]),
					STL:                  toFloat(row[21]),
					BLK:                  toFloat(row[22]),
					BLKA:                 toInt(row[23]),
					PF:                   toFloat(row[24]),
					PFD:                  toFloat(row[25]),
					PTS:                  toFloat(row[26]),
					PLUS_MINUS:           toFloat(row[27]),
					NBA_FANTASY_PTS:      toFloat(row[28]),
					DD2:                  toFloat(row[29]),
					TD3:                  toFloat(row[30]),
					GP_RANK:              toFloat(row[31]),
					W_RANK:               toFloat(row[32]),
					L_RANK:               toFloat(row[33]),
					W_PCT_RANK:           toFloat(row[34]),
					MIN_RANK:             toFloat(row[35]),
					FGM_RANK:             toFloat(row[36]),
					FGA_RANK:             toFloat(row[37]),
					FG_PCT_RANK:          toFloat(row[38]),
					FG3M_RANK:            toFloat(row[39]),
					FG3A_RANK:            toFloat(row[40]),
					FG3_PCT_RANK:         toFloat(row[41]),
					FTM_RANK:             toFloat(row[42]),
					FTA_RANK:             toFloat(row[43]),
					FT_PCT_RANK:          toFloat(row[44]),
					OREB_RANK:            toFloat(row[45]),
					DREB_RANK:            toFloat(row[46]),
					REB_RANK:             toFloat(row[47]),
					AST_RANK:             toFloat(row[48]),
					TOV_RANK:             toFloat(row[49]),
					STL_RANK:             toFloat(row[50]),
					BLK_RANK:             toFloat(row[51]),
					BLKA_RANK:            toFloat(row[52]),
					PF_RANK:              toFloat(row[53]),
					PFD_RANK:             toFloat(row[54]),
					PTS_RANK:             toFloat(row[55]),
					PLUS_MINUS_RANK:      toFloat(row[56]),
					NBA_FANTASY_PTS_RANK: toFloat(row[57]),
					DD2_RANK:             toFloat(row[58]),
					TD3_RANK:             toFloat(row[59]),
				}
				response.OverallPlayerDashboard = append(response.OverallPlayerDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "LocationPlayerDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayerDashboardByGeneralSplitsLocationPlayerDashboard{})); err != nil {
			return nil, fmt.Errorf("PlayerDashboardByGeneralSplits: LocationPlayerDashboard result set: %w", err)
		}
		response.LocationPlayerDashboard = make([]PlayerDashboardByGeneralSplitsLocationPlayerDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 60 {
				item := PlayerDashboardByGeneralSplitsLocationPlayerDashboard{
					GROUP_SET:            toString(row[0]),
					GROUP_VALUE:          toString(row[1]),
					GP:                   toInt(row[2]),
					W:                    toString(row[3]),
					L:                    toString(row[4]),
					W_PCT:                toFloat(row[5]),
					MIN:                  toFloat(row[6]),
					FGM:                  toInt(row[7]),
					FGA:                  toInt(row[8]),
					FG_PCT:               toFloat(row[9]),
					FG3M:                 toInt(row[10]),
					FG3A:                 toInt(row[11]),
					FG3_PCT:              toFloat(row[12]),
					FTM:                  toInt(row[13]),
					FTA:                  toInt(row[14]),
					FT_PCT:               toFloat(row[15]),
					OREB:                 toFloat(row[16]),
					DREB:                 toFloat(row[17]),
					REB:                  toFloat(row[18]),
					AST:                  toFloat(row[19]),
					TOV:                  toFloat(row[20]),
					STL:                  toFloat(row[21]),
					BLK:                  toFloat(row[22]),
					BLKA:                 toInt(row[23]),
					PF:                   toFloat(row[24]),
					PFD:                  toFloat(row[25]),
					PTS:                  toFloat(row[26]),
					PLUS_MINUS:           toFloat(row[27]),
					NBA_FANTASY_PTS:      toFloat(row[28]),
					DD2:                  toFloat(row[29]),
					TD3:                  toFloat(row[30]),
					GP_RANK:              toFloat(row[31]),
					W_RANK:               toFloat(row[32]),
					L_RANK:               toFloat(row[33]),
					W_PCT_RANK:           toFloat(row[34]),
					MIN_RANK:             toFloat(row[35]),
					FGM_RANK:             toFloat(row[36]),
					FGA_RANK:             toFloat(row[37]),
					FG_PCT_RANK:          toFloat(row[38]),
					FG3M_RANK:            toFloat(row[39]),
					FG3A_RANK:            toFloat(row[40]),
					FG3_PCT_RANK:         toFloat(row[41]),
					FTM_RANK:             toFloat(row[42]),
					FTA_RANK:             toFloat(row[43]),
					FT_PCT_RANK:          toFloat(row[44]),
					OREB_RANK:            toFloat(row[45]),
					DREB_RANK:            toFloat(row[46]),
					REB_RANK:             toFloat(row[47]),
					AST_RANK:             toFloat(row[48]),
					TOV_RANK:             toFloat(row[49]),
					STL_RANK:             toFloat(row[50]),
					BLK_RANK:             toFloat(row[51]),
					BLKA_RANK:            toFloat(row[52]),
					PF_RANK:              toFloat(row[53]),
					PFD_RANK:             toFloat(row[54]),
					PTS_RANK:             toFloat(row[55]),
					PLUS_MINUS_RANK:      toFloat(row[56]),
					NBA_FANTASY_PTS_RANK: toFloat(row[57]),
					DD2_RANK:             toFloat(row[58]),
					TD3_RANK:             toFloat(row[59]),
				}
				response.LocationPlayerDashboard = append(response.LocationPlayerDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "WinsLossesPlayerDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard{})); err != nil {
			return nil, fmt.Errorf("PlayerDashboardByGeneralSplits: WinsLossesPlayerDashboard result set: %w", err)
		}
		response.WinsLossesPlayerDashboard = make([]PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 60 {
				item := PlayerDashboardByGeneralSplitsWinsLossesPlayerDashboard{
					GROUP_SET:            toString(row[0]),
					GROUP_VALUE:          toString(row[1]),
					GP:                   toInt(row[2]),
					W:                    toString(row[3]),
					L:                    toString(row[4]),
					W_PCT:                toFloat(row[5]),
					MIN:                  toFloat(row[6]),
					FGM:                  toInt(row[7]),
					FGA:                  toInt(row[8]),
					FG_PCT:               toFloat(row[9]),
					FG3M:                 toInt(row[10]),
					FG3A:                 toInt(row[11]),
					FG3_PCT:              toFloat(row[12]),
					FTM:                  toInt(row[13]),
					FTA:                  toInt(row[14]),
					FT_PCT:               toFloat(row[15]),
					OREB:                 toFloat(row[16]),
					DREB:                 toFloat(row[17]),
					REB:                  toFloat(row[18]),
					AST:                  toFloat(row[19]),
					TOV:                  toFloat(row[20]),
					STL:                  toFloat(row[21]),
					BLK:                  toFloat(row[22]),
					BLKA:                 toInt(row[23]),
					PF:                   toFloat(row[24]),
					PFD:                  toFloat(row[25]),
					PTS:                  toFloat(row[26]),
					PLUS_MINUS:           toFloat(row[27]),
					NBA_FANTASY_PTS:      toFloat(row[28]),
					DD2:                  toFloat(row[29]),
					TD3:                  toFloat(row[30]),
					GP_RANK:              toFloat(row[31]),
					W_RANK:               toFloat(row[32]),
					L_RANK:               toFloat(row[33]),
					W_PCT_RANK:           toFloat(row[34]),
					MIN_RANK:             toFloat(row[35]),
					FGM_RANK:             toFloat(row[36]),
					FGA_RANK:             toFloat(row[37]),
					FG_PCT_RANK:          toFloat(row[38]),
					FG3M_RANK:            toFloat(row[39]),
					FG3A_RANK:            toFloat(row[40]),
					FG3_PCT_RANK:         toFloat(row[41]),
					FTM_RANK:             toFloat(row[42]),
					FTA_RANK:             toFloat(row[43]),
					FT_PCT_RANK:          toFloat(row[44]),
					OREB_RANK:            toFloat(row[45]),
					DREB_RANK:            toFloat(row[46]),
					REB_RANK:             toFloat(row[47]),
					AST_RANK:             toFloat(row[48]),
					TOV_RANK:             toFloat(row[49]),
					STL_RANK:             toFloat(row[50]),
					BLK_RANK:             toFloat(row[51]),
					BLKA_RANK:            toFloat(row[52]),
					PF_RANK:              toFloat(row[53]),
					PFD_RANK:             toFloat(row[54]),
					PTS_RANK:             toFloat(row[55]),
					PLUS_MINUS_RANK:      toFloat(row[56]),
					NBA_FANTASY_PTS_RANK: toFloat(row[57]),
					DD2_RANK:             toFloat(row[58]),
					TD3_RANK:             toFloat(row[59]),
				}
				response.WinsLossesPlayerDashboard = append(response.WinsLossesPlayerDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "MonthPlayerDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayerDashboardByGeneralSplitsMonthPlayerDashboard{})); err != nil {
			return nil, fmt.Errorf("PlayerDashboardByGeneralSplits: MonthPlayerDashboard result set: %w", err)
		}
		response.MonthPlayerDashboard = make([]PlayerDashboardByGeneralSplitsMonthPlayerDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 60 {
				item := PlayerDashboardByGeneralSplitsMonthPlayerDashboard{
					GROUP_SET:            toString(row[0]),
					GROUP_VALUE:          toString(row[1]),
					GP:                   toInt(row[2]),
					W:                    toString(row[3]),
					L:                    toString(row[4]),
					W_PCT:                toFloat(row[5]),
					MIN:                  toFloat(row[6]),
					FGM:                  toInt(row[7]),
					FGA:                  toInt(row[8]),
					FG_PCT:               toFloat(row[9]),
					FG3M:                 toInt(row[10]),
					FG3A:                 toInt(row[11]),
					FG3_PCT:              toFloat(row[12]),
					FTM:                  toInt(row[13]),
					FTA:                  toInt(row[14]),
					FT_PCT:               toFloat(row[15]),
					OREB:                 toFloat(row[16]),
					DREB:                 toFloat(row[17]),
					REB:                  toFloat(row[18]),
					AST:                  toFloat(row[19]),
					TOV:                  toFloat(row[20]),
					STL:                  toFloat(row[21]),
					BLK:                  toFloat(row[22]),
					BLKA:                 toInt(row[23]),
					PF:                   toFloat(row[24]),
					PFD:                  toFloat(row[25]),
					PTS:                  toFloat(row[26]),
					PLUS_MINUS:           toFloat(row[27]),
					NBA_FANTASY_PTS:      toFloat(row[28]),
					DD2:                  toFloat(row[29]),
					TD3:                  toFloat(row[30]),
					GP_RANK:              toFloat(row[31]),
					W_RANK:               toFloat(row[32]),
					L_RANK:               toFloat(row[33]),
					W_PCT_RANK:           toFloat(row[34]),
					MIN_RANK:             toFloat(row[35]),
					FGM_RANK:             toFloat(row[36]),
					FGA_RANK:             toFloat(row[37]),
					FG_PCT_RANK:          toFloat(row[38]),
					FG3M_RANK:            toFloat(row[39]),
					FG3A_RANK:            toFloat(row[40]),
					FG3_PCT_RANK:         toFloat(row[41]),
					FTM_RANK:             toFloat(row[42]),
					FTA_RANK:             toFloat(row[43]),
					FT_PCT_RANK:          toFloat(row[44]),
					OREB_RANK:            toFloat(row[45]),
					DREB_RANK:            toFloat(row[46]),
					REB_RANK:             toFloat(row[47]),
					AST_RANK:             toFloat(row[48]),
					TOV_RANK:             toFloat(row[49]),
					STL_RANK:             toFloat(row[50]),
					BLK_RANK:             toFloat(row[51]),
					BLKA_RANK:            toFloat(row[52]),
					PF_RANK:              toFloat(row[53]),
					PFD_RANK:             toFloat(row[54]),
					PTS_RANK:             toFloat(row[55]),
					PLUS_MINUS_RANK:      toFloat(row[56]),
					NBA_FANTASY_PTS_RANK: toFloat(row[57]),
					DD2_RANK:             toFloat(row[58]),
					TD3_RANK:             toFloat(row[59]),
				}
				response.MonthPlayerDashboard = append(response.MonthPlayerDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "PrePostAllStarPlayerDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard{})); err != nil {
			return nil, fmt.Errorf("PlayerDashboardByGeneralSplits: PrePostAllStarPlayerDashboard result set: %w", err)
		}
		response.PrePostAllStarPlayerDashboard = make([]PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 60 {
				item := PlayerDashboardByGeneralSplitsPrePostAllStarPlayerDashboard{
					GROUP_SET:            toString(row[0]),
					GROUP_VALUE:          toString(row[1]),
					GP:                   toInt(row[2]),
					W:                    toString(row[3]),
					L:                    toString(row[4]),
					W_PCT:                toFloat(row[5]),
					MIN:                  toFloat(row[6]),
					FGM:                  toInt(row[7]),
					FGA:                  toInt(row[8]),
					FG_PCT:               toFloat(row[9]),
					FG3M:                 toInt(row[10]),
					FG3A:                 toInt(row[11]),
					FG3_PCT:              toFloat(row[12]),
					FTM:                  toInt(row[13]),
					FTA:                  toInt(row[14]),
					FT_PCT:               toFloat(row[15]),
					OREB:                 toFloat(row[16]),
					DREB:                 toFloat(row[17]),
					REB:                  toFloat(row[18]),
					AST:                  toFloat(row[19]),
					TOV:                  toFloat(row[20]),
					STL:                  toFloat(row[21]),
					BLK:                  toFloat(row[22]),
					BLKA:                 toInt(row[23]),
					PF:                   toFloat(row[24]),
					PFD:                  toFloat(row[25]),
					PTS:                  toFloat(row[26]),
					PLUS_MINUS:           toFloat(row[27]),
					NBA_FANTASY_PTS:      toFloat(row[28]),
					DD2:                  toFloat(row[29]),
					TD3:                  toFloat(row[30]),
					GP_RANK:              toFloat(row[31]),
					W_RANK:               toFloat(row[32]),
					L_RANK:               toFloat(row[33]),
					W_PCT_RANK:           toFloat(row[34]),
					MIN_RANK:             toFloat(row[35]),
					FGM_RANK:             toFloat(row[36]),
					FGA_RANK:             toFloat(row[37]),
					FG_PCT_RANK:          toFloat(row[38]),
					FG3M_RANK:            toFloat(row[39]),
					FG3A_RANK:            toFloat(row[40]),
					FG3_PCT_RANK:         toFloat(row[41]),
					FTM_RANK:             toFloat(row[42]),
					FTA_RANK:             toFloat(row[43]),
					FT_PCT_RANK:          toFloat(row[44]),
					OREB_RANK:            toFloat(row[45]),
					DREB_RANK:            toFloat(row[46]),
					REB_RANK:             toFloat(row[47]),
					AST_RANK:             toFloat(row[48]),
					TOV_RANK:             toFloat(row[49]),
					STL_RANK:             toFloat(row[50]),
					BLK_RANK:             toFloat(row[51]),
					BLKA_RANK:            toFloat(row[52]),
					PF_RANK:              toFloat(row[53]),
					PFD_RANK:             toFloat(row[54]),
					PTS_RANK:             toFloat(row[55]),
					PLUS_MINUS_RANK:      toFloat(row[56]),
					NBA_FANTASY_PTS_RANK: toFloat(row[57]),
					DD2_RANK:             toFloat(row[58]),
					TD3_RANK:             toFloat(row[59]),
				}
				response.PrePostAllStarPlayerDashboard = append(response.PrePostAllStarPlayerDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
