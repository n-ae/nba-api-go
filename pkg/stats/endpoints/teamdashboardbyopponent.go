package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamDashboardByOpponentRequest contains parameters for the TeamDashboardByOpponent endpoint
type TeamDashboardByOpponentRequest struct {
	TeamID      string
	MeasureType *string
	PerMode     *parameters.PerMode
	PlusMinus   *string
	PaceAdjust  *string
	Rank        *string
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	LeagueID    *parameters.LeagueID
}

// TeamDashboardByOpponentOverallTeamDashboard represents the OverallTeamDashboard result set for TeamDashboardByOpponent
type TeamDashboardByOpponentOverallTeamDashboard struct {
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

// TeamDashboardByOpponentConferenceTeamDashboard represents the ConferenceTeamDashboard result set for TeamDashboardByOpponent
type TeamDashboardByOpponentConferenceTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	CONF       string  `json:"CONF"`
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

// TeamDashboardByOpponentDivisionTeamDashboard represents the DivisionTeamDashboard result set for TeamDashboardByOpponent
type TeamDashboardByOpponentDivisionTeamDashboard struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	DIV        string  `json:"DIV"`
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

// TeamDashboardByOpponentOpponentTeamDashboard represents the OpponentTeamDashboard result set for TeamDashboardByOpponent
type TeamDashboardByOpponentOpponentTeamDashboard struct {
	TEAM_ID              int     `json:"TEAM_ID"`
	TEAM_NAME            string  `json:"TEAM_NAME"`
	VS_TEAM_ID           int     `json:"VS_TEAM_ID"`
	VS_TEAM_ABBREVIATION string  `json:"VS_TEAM_ABBREVIATION"`
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
}

// TeamDashboardByOpponentResponse contains the response data from the TeamDashboardByOpponent endpoint
type TeamDashboardByOpponentResponse struct {
	OverallTeamDashboard    []TeamDashboardByOpponentOverallTeamDashboard
	ConferenceTeamDashboard []TeamDashboardByOpponentConferenceTeamDashboard
	DivisionTeamDashboard   []TeamDashboardByOpponentDivisionTeamDashboard
	OpponentTeamDashboard   []TeamDashboardByOpponentOpponentTeamDashboard
}

// GetTeamDashboardByOpponent retrieves data from the teamdashboardbyopponent endpoint
func GetTeamDashboardByOpponent(ctx context.Context, client *stats.Client, req TeamDashboardByOpponentRequest) (*models.Response[*TeamDashboardByOpponentResponse], error) {
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
	if err := client.GetJSON(ctx, "teamdashboardbyopponent", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDashboardByOpponentResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallTeamDashboard = make([]TeamDashboardByOpponentOverallTeamDashboard, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := TeamDashboardByOpponentOverallTeamDashboard{
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
		response.ConferenceTeamDashboard = make([]TeamDashboardByOpponentConferenceTeamDashboard, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByOpponentConferenceTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					CONF:       toString(row[2]),
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
				response.ConferenceTeamDashboard = append(response.ConferenceTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.DivisionTeamDashboard = make([]TeamDashboardByOpponentDivisionTeamDashboard, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByOpponentDivisionTeamDashboard{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					DIV:        toString(row[2]),
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
				response.DivisionTeamDashboard = append(response.DivisionTeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.OpponentTeamDashboard = make([]TeamDashboardByOpponentOpponentTeamDashboard, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 30 {
				item := TeamDashboardByOpponentOpponentTeamDashboard{
					TEAM_ID:              toInt(row[0]),
					TEAM_NAME:            toString(row[1]),
					VS_TEAM_ID:           toInt(row[2]),
					VS_TEAM_ABBREVIATION: toString(row[3]),
					GP:                   toInt(row[4]),
					W:                    toString(row[5]),
					L:                    toString(row[6]),
					W_PCT:                toFloat(row[7]),
					MIN:                  toFloat(row[8]),
					FGM:                  toInt(row[9]),
					FGA:                  toInt(row[10]),
					FG_PCT:               toFloat(row[11]),
					FG3M:                 toInt(row[12]),
					FG3A:                 toInt(row[13]),
					FG3_PCT:              toFloat(row[14]),
					FTM:                  toInt(row[15]),
					FTA:                  toInt(row[16]),
					FT_PCT:               toFloat(row[17]),
					OREB:                 toFloat(row[18]),
					DREB:                 toFloat(row[19]),
					REB:                  toFloat(row[20]),
					AST:                  toFloat(row[21]),
					TOV:                  toFloat(row[22]),
					STL:                  toFloat(row[23]),
					BLK:                  toFloat(row[24]),
					BLKA:                 toInt(row[25]),
					PF:                   toFloat(row[26]),
					PFD:                  toFloat(row[27]),
					PTS:                  toFloat(row[28]),
					PLUS_MINUS:           toFloat(row[29]),
				}
				response.OpponentTeamDashboard = append(response.OpponentTeamDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
