package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
	"github.com/n-ae/nba-api-go/v2/pkg/stats"
	"github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
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
	GROUP_SET       string  `json:"GROUP_SET"`
	GROUP_VALUE     string  `json:"GROUP_VALUE"`
	GP              int     `json:"GP"`
	W               string  `json:"W"`
	L               string  `json:"L"`
	W_PCT           float64 `json:"W_PCT"`
	MIN             float64 `json:"MIN"`
	FGM             int     `json:"FGM"`
	FGA             int     `json:"FGA"`
	FG_PCT          float64 `json:"FG_PCT"`
	FG3M            int     `json:"FG3M"`
	FG3A            int     `json:"FG3A"`
	FG3_PCT         float64 `json:"FG3_PCT"`
	FTM             int     `json:"FTM"`
	FTA             int     `json:"FTA"`
	FT_PCT          float64 `json:"FT_PCT"`
	OREB            float64 `json:"OREB"`
	DREB            float64 `json:"DREB"`
	REB             float64 `json:"REB"`
	AST             float64 `json:"AST"`
	TOV             float64 `json:"TOV"`
	STL             float64 `json:"STL"`
	BLK             float64 `json:"BLK"`
	BLKA            int     `json:"BLKA"`
	PF              float64 `json:"PF"`
	PFD             float64 `json:"PFD"`
	PTS             float64 `json:"PTS"`
	PLUS_MINUS      float64 `json:"PLUS_MINUS"`
	GP_RANK         float64 `json:"GP_RANK"`
	W_RANK          float64 `json:"W_RANK"`
	L_RANK          float64 `json:"L_RANK"`
	W_PCT_RANK      float64 `json:"W_PCT_RANK"`
	MIN_RANK        float64 `json:"MIN_RANK"`
	FGM_RANK        float64 `json:"FGM_RANK"`
	FGA_RANK        float64 `json:"FGA_RANK"`
	FG_PCT_RANK     float64 `json:"FG_PCT_RANK"`
	FG3M_RANK       float64 `json:"FG3M_RANK"`
	FG3A_RANK       float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK    float64 `json:"FG3_PCT_RANK"`
	FTM_RANK        float64 `json:"FTM_RANK"`
	FTA_RANK        float64 `json:"FTA_RANK"`
	FT_PCT_RANK     float64 `json:"FT_PCT_RANK"`
	OREB_RANK       float64 `json:"OREB_RANK"`
	DREB_RANK       float64 `json:"DREB_RANK"`
	REB_RANK        float64 `json:"REB_RANK"`
	AST_RANK        float64 `json:"AST_RANK"`
	TOV_RANK        float64 `json:"TOV_RANK"`
	STL_RANK        float64 `json:"STL_RANK"`
	BLK_RANK        float64 `json:"BLK_RANK"`
	BLKA_RANK       float64 `json:"BLKA_RANK"`
	PF_RANK         float64 `json:"PF_RANK"`
	PFD_RANK        float64 `json:"PFD_RANK"`
	PTS_RANK        float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK float64 `json:"PLUS_MINUS_RANK"`
}

// TeamDashboardByGeneralSplitsLocationTeamDashboard represents the LocationTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsLocationTeamDashboard struct {
	GROUP_SET       string  `json:"GROUP_SET"`
	GROUP_VALUE     string  `json:"GROUP_VALUE"`
	GP              int     `json:"GP"`
	W               string  `json:"W"`
	L               string  `json:"L"`
	W_PCT           float64 `json:"W_PCT"`
	MIN             float64 `json:"MIN"`
	FGM             int     `json:"FGM"`
	FGA             int     `json:"FGA"`
	FG_PCT          float64 `json:"FG_PCT"`
	FG3M            int     `json:"FG3M"`
	FG3A            int     `json:"FG3A"`
	FG3_PCT         float64 `json:"FG3_PCT"`
	FTM             int     `json:"FTM"`
	FTA             int     `json:"FTA"`
	FT_PCT          float64 `json:"FT_PCT"`
	OREB            float64 `json:"OREB"`
	DREB            float64 `json:"DREB"`
	REB             float64 `json:"REB"`
	AST             float64 `json:"AST"`
	TOV             float64 `json:"TOV"`
	STL             float64 `json:"STL"`
	BLK             float64 `json:"BLK"`
	BLKA            int     `json:"BLKA"`
	PF              float64 `json:"PF"`
	PFD             float64 `json:"PFD"`
	PTS             float64 `json:"PTS"`
	PLUS_MINUS      float64 `json:"PLUS_MINUS"`
	GP_RANK         float64 `json:"GP_RANK"`
	W_RANK          float64 `json:"W_RANK"`
	L_RANK          float64 `json:"L_RANK"`
	W_PCT_RANK      float64 `json:"W_PCT_RANK"`
	MIN_RANK        float64 `json:"MIN_RANK"`
	FGM_RANK        float64 `json:"FGM_RANK"`
	FGA_RANK        float64 `json:"FGA_RANK"`
	FG_PCT_RANK     float64 `json:"FG_PCT_RANK"`
	FG3M_RANK       float64 `json:"FG3M_RANK"`
	FG3A_RANK       float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK    float64 `json:"FG3_PCT_RANK"`
	FTM_RANK        float64 `json:"FTM_RANK"`
	FTA_RANK        float64 `json:"FTA_RANK"`
	FT_PCT_RANK     float64 `json:"FT_PCT_RANK"`
	OREB_RANK       float64 `json:"OREB_RANK"`
	DREB_RANK       float64 `json:"DREB_RANK"`
	REB_RANK        float64 `json:"REB_RANK"`
	AST_RANK        float64 `json:"AST_RANK"`
	TOV_RANK        float64 `json:"TOV_RANK"`
	STL_RANK        float64 `json:"STL_RANK"`
	BLK_RANK        float64 `json:"BLK_RANK"`
	BLKA_RANK       float64 `json:"BLKA_RANK"`
	PF_RANK         float64 `json:"PF_RANK"`
	PFD_RANK        float64 `json:"PFD_RANK"`
	PTS_RANK        float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK float64 `json:"PLUS_MINUS_RANK"`
}

// TeamDashboardByGeneralSplitsWinsLossesTeamDashboard represents the WinsLossesTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsWinsLossesTeamDashboard struct {
	GROUP_SET       string  `json:"GROUP_SET"`
	GROUP_VALUE     string  `json:"GROUP_VALUE"`
	GP              int     `json:"GP"`
	W               string  `json:"W"`
	L               string  `json:"L"`
	W_PCT           float64 `json:"W_PCT"`
	MIN             float64 `json:"MIN"`
	FGM             int     `json:"FGM"`
	FGA             int     `json:"FGA"`
	FG_PCT          float64 `json:"FG_PCT"`
	FG3M            int     `json:"FG3M"`
	FG3A            int     `json:"FG3A"`
	FG3_PCT         float64 `json:"FG3_PCT"`
	FTM             int     `json:"FTM"`
	FTA             int     `json:"FTA"`
	FT_PCT          float64 `json:"FT_PCT"`
	OREB            float64 `json:"OREB"`
	DREB            float64 `json:"DREB"`
	REB             float64 `json:"REB"`
	AST             float64 `json:"AST"`
	TOV             float64 `json:"TOV"`
	STL             float64 `json:"STL"`
	BLK             float64 `json:"BLK"`
	BLKA            int     `json:"BLKA"`
	PF              float64 `json:"PF"`
	PFD             float64 `json:"PFD"`
	PTS             float64 `json:"PTS"`
	PLUS_MINUS      float64 `json:"PLUS_MINUS"`
	GP_RANK         float64 `json:"GP_RANK"`
	W_RANK          float64 `json:"W_RANK"`
	L_RANK          float64 `json:"L_RANK"`
	W_PCT_RANK      float64 `json:"W_PCT_RANK"`
	MIN_RANK        float64 `json:"MIN_RANK"`
	FGM_RANK        float64 `json:"FGM_RANK"`
	FGA_RANK        float64 `json:"FGA_RANK"`
	FG_PCT_RANK     float64 `json:"FG_PCT_RANK"`
	FG3M_RANK       float64 `json:"FG3M_RANK"`
	FG3A_RANK       float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK    float64 `json:"FG3_PCT_RANK"`
	FTM_RANK        float64 `json:"FTM_RANK"`
	FTA_RANK        float64 `json:"FTA_RANK"`
	FT_PCT_RANK     float64 `json:"FT_PCT_RANK"`
	OREB_RANK       float64 `json:"OREB_RANK"`
	DREB_RANK       float64 `json:"DREB_RANK"`
	REB_RANK        float64 `json:"REB_RANK"`
	AST_RANK        float64 `json:"AST_RANK"`
	TOV_RANK        float64 `json:"TOV_RANK"`
	STL_RANK        float64 `json:"STL_RANK"`
	BLK_RANK        float64 `json:"BLK_RANK"`
	BLKA_RANK       float64 `json:"BLKA_RANK"`
	PF_RANK         float64 `json:"PF_RANK"`
	PFD_RANK        float64 `json:"PFD_RANK"`
	PTS_RANK        float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK float64 `json:"PLUS_MINUS_RANK"`
}

// TeamDashboardByGeneralSplitsMonthTeamDashboard represents the MonthTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsMonthTeamDashboard struct {
	GROUP_SET       string  `json:"GROUP_SET"`
	GROUP_VALUE     string  `json:"GROUP_VALUE"`
	GP              int     `json:"GP"`
	W               string  `json:"W"`
	L               string  `json:"L"`
	W_PCT           float64 `json:"W_PCT"`
	MIN             float64 `json:"MIN"`
	FGM             int     `json:"FGM"`
	FGA             int     `json:"FGA"`
	FG_PCT          float64 `json:"FG_PCT"`
	FG3M            int     `json:"FG3M"`
	FG3A            int     `json:"FG3A"`
	FG3_PCT         float64 `json:"FG3_PCT"`
	FTM             int     `json:"FTM"`
	FTA             int     `json:"FTA"`
	FT_PCT          float64 `json:"FT_PCT"`
	OREB            float64 `json:"OREB"`
	DREB            float64 `json:"DREB"`
	REB             float64 `json:"REB"`
	AST             float64 `json:"AST"`
	TOV             float64 `json:"TOV"`
	STL             float64 `json:"STL"`
	BLK             float64 `json:"BLK"`
	BLKA            int     `json:"BLKA"`
	PF              float64 `json:"PF"`
	PFD             float64 `json:"PFD"`
	PTS             float64 `json:"PTS"`
	PLUS_MINUS      float64 `json:"PLUS_MINUS"`
	GP_RANK         float64 `json:"GP_RANK"`
	W_RANK          float64 `json:"W_RANK"`
	L_RANK          float64 `json:"L_RANK"`
	W_PCT_RANK      float64 `json:"W_PCT_RANK"`
	MIN_RANK        float64 `json:"MIN_RANK"`
	FGM_RANK        float64 `json:"FGM_RANK"`
	FGA_RANK        float64 `json:"FGA_RANK"`
	FG_PCT_RANK     float64 `json:"FG_PCT_RANK"`
	FG3M_RANK       float64 `json:"FG3M_RANK"`
	FG3A_RANK       float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK    float64 `json:"FG3_PCT_RANK"`
	FTM_RANK        float64 `json:"FTM_RANK"`
	FTA_RANK        float64 `json:"FTA_RANK"`
	FT_PCT_RANK     float64 `json:"FT_PCT_RANK"`
	OREB_RANK       float64 `json:"OREB_RANK"`
	DREB_RANK       float64 `json:"DREB_RANK"`
	REB_RANK        float64 `json:"REB_RANK"`
	AST_RANK        float64 `json:"AST_RANK"`
	TOV_RANK        float64 `json:"TOV_RANK"`
	STL_RANK        float64 `json:"STL_RANK"`
	BLK_RANK        float64 `json:"BLK_RANK"`
	BLKA_RANK       float64 `json:"BLKA_RANK"`
	PF_RANK         float64 `json:"PF_RANK"`
	PFD_RANK        float64 `json:"PFD_RANK"`
	PTS_RANK        float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK float64 `json:"PLUS_MINUS_RANK"`
}

// TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard represents the PrePostAllStarTeamDashboard result set for TeamDashboardByGeneralSplits
type TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard struct {
	GROUP_SET       string  `json:"GROUP_SET"`
	GROUP_VALUE     string  `json:"GROUP_VALUE"`
	GP              int     `json:"GP"`
	W               string  `json:"W"`
	L               string  `json:"L"`
	W_PCT           float64 `json:"W_PCT"`
	MIN             float64 `json:"MIN"`
	FGM             int     `json:"FGM"`
	FGA             int     `json:"FGA"`
	FG_PCT          float64 `json:"FG_PCT"`
	FG3M            int     `json:"FG3M"`
	FG3A            int     `json:"FG3A"`
	FG3_PCT         float64 `json:"FG3_PCT"`
	FTM             int     `json:"FTM"`
	FTA             int     `json:"FTA"`
	FT_PCT          float64 `json:"FT_PCT"`
	OREB            float64 `json:"OREB"`
	DREB            float64 `json:"DREB"`
	REB             float64 `json:"REB"`
	AST             float64 `json:"AST"`
	TOV             float64 `json:"TOV"`
	STL             float64 `json:"STL"`
	BLK             float64 `json:"BLK"`
	BLKA            int     `json:"BLKA"`
	PF              float64 `json:"PF"`
	PFD             float64 `json:"PFD"`
	PTS             float64 `json:"PTS"`
	PLUS_MINUS      float64 `json:"PLUS_MINUS"`
	GP_RANK         float64 `json:"GP_RANK"`
	W_RANK          float64 `json:"W_RANK"`
	L_RANK          float64 `json:"L_RANK"`
	W_PCT_RANK      float64 `json:"W_PCT_RANK"`
	MIN_RANK        float64 `json:"MIN_RANK"`
	FGM_RANK        float64 `json:"FGM_RANK"`
	FGA_RANK        float64 `json:"FGA_RANK"`
	FG_PCT_RANK     float64 `json:"FG_PCT_RANK"`
	FG3M_RANK       float64 `json:"FG3M_RANK"`
	FG3A_RANK       float64 `json:"FG3A_RANK"`
	FG3_PCT_RANK    float64 `json:"FG3_PCT_RANK"`
	FTM_RANK        float64 `json:"FTM_RANK"`
	FTA_RANK        float64 `json:"FTA_RANK"`
	FT_PCT_RANK     float64 `json:"FT_PCT_RANK"`
	OREB_RANK       float64 `json:"OREB_RANK"`
	DREB_RANK       float64 `json:"DREB_RANK"`
	REB_RANK        float64 `json:"REB_RANK"`
	AST_RANK        float64 `json:"AST_RANK"`
	TOV_RANK        float64 `json:"TOV_RANK"`
	STL_RANK        float64 `json:"STL_RANK"`
	BLK_RANK        float64 `json:"BLK_RANK"`
	BLKA_RANK       float64 `json:"BLKA_RANK"`
	PF_RANK         float64 `json:"PF_RANK"`
	PFD_RANK        float64 `json:"PFD_RANK"`
	PTS_RANK        float64 `json:"PTS_RANK"`
	PLUS_MINUS_RANK float64 `json:"PLUS_MINUS_RANK"`
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
		return nil, fmt.Errorf("%s is required", "TeamID")
	}
	params.Set("TeamID", req.TeamID)
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
	if err := client.GetJSON(ctx, "teamdashboardbygeneralsplits", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDashboardByGeneralSplitsResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "OverallTeamDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDashboardByGeneralSplitsOverallTeamDashboard{})); err != nil {
			return nil, fmt.Errorf("TeamDashboardByGeneralSplits: OverallTeamDashboard result set: %w", err)
		}
		response.OverallTeamDashboard = make([]TeamDashboardByGeneralSplitsOverallTeamDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 54 {
				item := TeamDashboardByGeneralSplitsOverallTeamDashboard{
					GROUP_SET:       toString(row[0]),
					GROUP_VALUE:     toString(row[1]),
					GP:              toInt(row[2]),
					W:               toString(row[3]),
					L:               toString(row[4]),
					W_PCT:           toFloat(row[5]),
					MIN:             toFloat(row[6]),
					FGM:             toInt(row[7]),
					FGA:             toInt(row[8]),
					FG_PCT:          toFloat(row[9]),
					FG3M:            toInt(row[10]),
					FG3A:            toInt(row[11]),
					FG3_PCT:         toFloat(row[12]),
					FTM:             toInt(row[13]),
					FTA:             toInt(row[14]),
					FT_PCT:          toFloat(row[15]),
					OREB:            toFloat(row[16]),
					DREB:            toFloat(row[17]),
					REB:             toFloat(row[18]),
					AST:             toFloat(row[19]),
					TOV:             toFloat(row[20]),
					STL:             toFloat(row[21]),
					BLK:             toFloat(row[22]),
					BLKA:            toInt(row[23]),
					PF:              toFloat(row[24]),
					PFD:             toFloat(row[25]),
					PTS:             toFloat(row[26]),
					PLUS_MINUS:      toFloat(row[27]),
					GP_RANK:         toFloat(row[28]),
					W_RANK:          toFloat(row[29]),
					L_RANK:          toFloat(row[30]),
					W_PCT_RANK:      toFloat(row[31]),
					MIN_RANK:        toFloat(row[32]),
					FGM_RANK:        toFloat(row[33]),
					FGA_RANK:        toFloat(row[34]),
					FG_PCT_RANK:     toFloat(row[35]),
					FG3M_RANK:       toFloat(row[36]),
					FG3A_RANK:       toFloat(row[37]),
					FG3_PCT_RANK:    toFloat(row[38]),
					FTM_RANK:        toFloat(row[39]),
					FTA_RANK:        toFloat(row[40]),
					FT_PCT_RANK:     toFloat(row[41]),
					OREB_RANK:       toFloat(row[42]),
					DREB_RANK:       toFloat(row[43]),
					REB_RANK:        toFloat(row[44]),
					AST_RANK:        toFloat(row[45]),
					TOV_RANK:        toFloat(row[46]),
					STL_RANK:        toFloat(row[47]),
					BLK_RANK:        toFloat(row[48]),
					BLKA_RANK:       toFloat(row[49]),
					PF_RANK:         toFloat(row[50]),
					PFD_RANK:        toFloat(row[51]),
					PTS_RANK:        toFloat(row[52]),
					PLUS_MINUS_RANK: toFloat(row[53]),
				}
				response.OverallTeamDashboard = append(response.OverallTeamDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "LocationTeamDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDashboardByGeneralSplitsLocationTeamDashboard{})); err != nil {
			return nil, fmt.Errorf("TeamDashboardByGeneralSplits: LocationTeamDashboard result set: %w", err)
		}
		response.LocationTeamDashboard = make([]TeamDashboardByGeneralSplitsLocationTeamDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 54 {
				item := TeamDashboardByGeneralSplitsLocationTeamDashboard{
					GROUP_SET:       toString(row[0]),
					GROUP_VALUE:     toString(row[1]),
					GP:              toInt(row[2]),
					W:               toString(row[3]),
					L:               toString(row[4]),
					W_PCT:           toFloat(row[5]),
					MIN:             toFloat(row[6]),
					FGM:             toInt(row[7]),
					FGA:             toInt(row[8]),
					FG_PCT:          toFloat(row[9]),
					FG3M:            toInt(row[10]),
					FG3A:            toInt(row[11]),
					FG3_PCT:         toFloat(row[12]),
					FTM:             toInt(row[13]),
					FTA:             toInt(row[14]),
					FT_PCT:          toFloat(row[15]),
					OREB:            toFloat(row[16]),
					DREB:            toFloat(row[17]),
					REB:             toFloat(row[18]),
					AST:             toFloat(row[19]),
					TOV:             toFloat(row[20]),
					STL:             toFloat(row[21]),
					BLK:             toFloat(row[22]),
					BLKA:            toInt(row[23]),
					PF:              toFloat(row[24]),
					PFD:             toFloat(row[25]),
					PTS:             toFloat(row[26]),
					PLUS_MINUS:      toFloat(row[27]),
					GP_RANK:         toFloat(row[28]),
					W_RANK:          toFloat(row[29]),
					L_RANK:          toFloat(row[30]),
					W_PCT_RANK:      toFloat(row[31]),
					MIN_RANK:        toFloat(row[32]),
					FGM_RANK:        toFloat(row[33]),
					FGA_RANK:        toFloat(row[34]),
					FG_PCT_RANK:     toFloat(row[35]),
					FG3M_RANK:       toFloat(row[36]),
					FG3A_RANK:       toFloat(row[37]),
					FG3_PCT_RANK:    toFloat(row[38]),
					FTM_RANK:        toFloat(row[39]),
					FTA_RANK:        toFloat(row[40]),
					FT_PCT_RANK:     toFloat(row[41]),
					OREB_RANK:       toFloat(row[42]),
					DREB_RANK:       toFloat(row[43]),
					REB_RANK:        toFloat(row[44]),
					AST_RANK:        toFloat(row[45]),
					TOV_RANK:        toFloat(row[46]),
					STL_RANK:        toFloat(row[47]),
					BLK_RANK:        toFloat(row[48]),
					BLKA_RANK:       toFloat(row[49]),
					PF_RANK:         toFloat(row[50]),
					PFD_RANK:        toFloat(row[51]),
					PTS_RANK:        toFloat(row[52]),
					PLUS_MINUS_RANK: toFloat(row[53]),
				}
				response.LocationTeamDashboard = append(response.LocationTeamDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "WinsLossesTeamDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDashboardByGeneralSplitsWinsLossesTeamDashboard{})); err != nil {
			return nil, fmt.Errorf("TeamDashboardByGeneralSplits: WinsLossesTeamDashboard result set: %w", err)
		}
		response.WinsLossesTeamDashboard = make([]TeamDashboardByGeneralSplitsWinsLossesTeamDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 54 {
				item := TeamDashboardByGeneralSplitsWinsLossesTeamDashboard{
					GROUP_SET:       toString(row[0]),
					GROUP_VALUE:     toString(row[1]),
					GP:              toInt(row[2]),
					W:               toString(row[3]),
					L:               toString(row[4]),
					W_PCT:           toFloat(row[5]),
					MIN:             toFloat(row[6]),
					FGM:             toInt(row[7]),
					FGA:             toInt(row[8]),
					FG_PCT:          toFloat(row[9]),
					FG3M:            toInt(row[10]),
					FG3A:            toInt(row[11]),
					FG3_PCT:         toFloat(row[12]),
					FTM:             toInt(row[13]),
					FTA:             toInt(row[14]),
					FT_PCT:          toFloat(row[15]),
					OREB:            toFloat(row[16]),
					DREB:            toFloat(row[17]),
					REB:             toFloat(row[18]),
					AST:             toFloat(row[19]),
					TOV:             toFloat(row[20]),
					STL:             toFloat(row[21]),
					BLK:             toFloat(row[22]),
					BLKA:            toInt(row[23]),
					PF:              toFloat(row[24]),
					PFD:             toFloat(row[25]),
					PTS:             toFloat(row[26]),
					PLUS_MINUS:      toFloat(row[27]),
					GP_RANK:         toFloat(row[28]),
					W_RANK:          toFloat(row[29]),
					L_RANK:          toFloat(row[30]),
					W_PCT_RANK:      toFloat(row[31]),
					MIN_RANK:        toFloat(row[32]),
					FGM_RANK:        toFloat(row[33]),
					FGA_RANK:        toFloat(row[34]),
					FG_PCT_RANK:     toFloat(row[35]),
					FG3M_RANK:       toFloat(row[36]),
					FG3A_RANK:       toFloat(row[37]),
					FG3_PCT_RANK:    toFloat(row[38]),
					FTM_RANK:        toFloat(row[39]),
					FTA_RANK:        toFloat(row[40]),
					FT_PCT_RANK:     toFloat(row[41]),
					OREB_RANK:       toFloat(row[42]),
					DREB_RANK:       toFloat(row[43]),
					REB_RANK:        toFloat(row[44]),
					AST_RANK:        toFloat(row[45]),
					TOV_RANK:        toFloat(row[46]),
					STL_RANK:        toFloat(row[47]),
					BLK_RANK:        toFloat(row[48]),
					BLKA_RANK:       toFloat(row[49]),
					PF_RANK:         toFloat(row[50]),
					PFD_RANK:        toFloat(row[51]),
					PTS_RANK:        toFloat(row[52]),
					PLUS_MINUS_RANK: toFloat(row[53]),
				}
				response.WinsLossesTeamDashboard = append(response.WinsLossesTeamDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "MonthTeamDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDashboardByGeneralSplitsMonthTeamDashboard{})); err != nil {
			return nil, fmt.Errorf("TeamDashboardByGeneralSplits: MonthTeamDashboard result set: %w", err)
		}
		response.MonthTeamDashboard = make([]TeamDashboardByGeneralSplitsMonthTeamDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 54 {
				item := TeamDashboardByGeneralSplitsMonthTeamDashboard{
					GROUP_SET:       toString(row[0]),
					GROUP_VALUE:     toString(row[1]),
					GP:              toInt(row[2]),
					W:               toString(row[3]),
					L:               toString(row[4]),
					W_PCT:           toFloat(row[5]),
					MIN:             toFloat(row[6]),
					FGM:             toInt(row[7]),
					FGA:             toInt(row[8]),
					FG_PCT:          toFloat(row[9]),
					FG3M:            toInt(row[10]),
					FG3A:            toInt(row[11]),
					FG3_PCT:         toFloat(row[12]),
					FTM:             toInt(row[13]),
					FTA:             toInt(row[14]),
					FT_PCT:          toFloat(row[15]),
					OREB:            toFloat(row[16]),
					DREB:            toFloat(row[17]),
					REB:             toFloat(row[18]),
					AST:             toFloat(row[19]),
					TOV:             toFloat(row[20]),
					STL:             toFloat(row[21]),
					BLK:             toFloat(row[22]),
					BLKA:            toInt(row[23]),
					PF:              toFloat(row[24]),
					PFD:             toFloat(row[25]),
					PTS:             toFloat(row[26]),
					PLUS_MINUS:      toFloat(row[27]),
					GP_RANK:         toFloat(row[28]),
					W_RANK:          toFloat(row[29]),
					L_RANK:          toFloat(row[30]),
					W_PCT_RANK:      toFloat(row[31]),
					MIN_RANK:        toFloat(row[32]),
					FGM_RANK:        toFloat(row[33]),
					FGA_RANK:        toFloat(row[34]),
					FG_PCT_RANK:     toFloat(row[35]),
					FG3M_RANK:       toFloat(row[36]),
					FG3A_RANK:       toFloat(row[37]),
					FG3_PCT_RANK:    toFloat(row[38]),
					FTM_RANK:        toFloat(row[39]),
					FTA_RANK:        toFloat(row[40]),
					FT_PCT_RANK:     toFloat(row[41]),
					OREB_RANK:       toFloat(row[42]),
					DREB_RANK:       toFloat(row[43]),
					REB_RANK:        toFloat(row[44]),
					AST_RANK:        toFloat(row[45]),
					TOV_RANK:        toFloat(row[46]),
					STL_RANK:        toFloat(row[47]),
					BLK_RANK:        toFloat(row[48]),
					BLKA_RANK:       toFloat(row[49]),
					PF_RANK:         toFloat(row[50]),
					PFD_RANK:        toFloat(row[51]),
					PTS_RANK:        toFloat(row[52]),
					PLUS_MINUS_RANK: toFloat(row[53]),
				}
				response.MonthTeamDashboard = append(response.MonthTeamDashboard, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "PrePostAllStarTeamDashboard"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard{})); err != nil {
			return nil, fmt.Errorf("TeamDashboardByGeneralSplits: PrePostAllStarTeamDashboard result set: %w", err)
		}
		response.PrePostAllStarTeamDashboard = make([]TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 54 {
				item := TeamDashboardByGeneralSplitsPrePostAllStarTeamDashboard{
					GROUP_SET:       toString(row[0]),
					GROUP_VALUE:     toString(row[1]),
					GP:              toInt(row[2]),
					W:               toString(row[3]),
					L:               toString(row[4]),
					W_PCT:           toFloat(row[5]),
					MIN:             toFloat(row[6]),
					FGM:             toInt(row[7]),
					FGA:             toInt(row[8]),
					FG_PCT:          toFloat(row[9]),
					FG3M:            toInt(row[10]),
					FG3A:            toInt(row[11]),
					FG3_PCT:         toFloat(row[12]),
					FTM:             toInt(row[13]),
					FTA:             toInt(row[14]),
					FT_PCT:          toFloat(row[15]),
					OREB:            toFloat(row[16]),
					DREB:            toFloat(row[17]),
					REB:             toFloat(row[18]),
					AST:             toFloat(row[19]),
					TOV:             toFloat(row[20]),
					STL:             toFloat(row[21]),
					BLK:             toFloat(row[22]),
					BLKA:            toInt(row[23]),
					PF:              toFloat(row[24]),
					PFD:             toFloat(row[25]),
					PTS:             toFloat(row[26]),
					PLUS_MINUS:      toFloat(row[27]),
					GP_RANK:         toFloat(row[28]),
					W_RANK:          toFloat(row[29]),
					L_RANK:          toFloat(row[30]),
					W_PCT_RANK:      toFloat(row[31]),
					MIN_RANK:        toFloat(row[32]),
					FGM_RANK:        toFloat(row[33]),
					FGA_RANK:        toFloat(row[34]),
					FG_PCT_RANK:     toFloat(row[35]),
					FG3M_RANK:       toFloat(row[36]),
					FG3A_RANK:       toFloat(row[37]),
					FG3_PCT_RANK:    toFloat(row[38]),
					FTM_RANK:        toFloat(row[39]),
					FTA_RANK:        toFloat(row[40]),
					FT_PCT_RANK:     toFloat(row[41]),
					OREB_RANK:       toFloat(row[42]),
					DREB_RANK:       toFloat(row[43]),
					REB_RANK:        toFloat(row[44]),
					AST_RANK:        toFloat(row[45]),
					TOV_RANK:        toFloat(row[46]),
					STL_RANK:        toFloat(row[47]),
					BLK_RANK:        toFloat(row[48]),
					BLKA_RANK:       toFloat(row[49]),
					PF_RANK:         toFloat(row[50]),
					PFD_RANK:        toFloat(row[51]),
					PTS_RANK:        toFloat(row[52]),
					PLUS_MINUS_RANK: toFloat(row[53]),
				}
				response.PrePostAllStarTeamDashboard = append(response.PrePostAllStarTeamDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
