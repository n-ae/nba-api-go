package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamVsPlayerRequest contains parameters for the TeamVsPlayer endpoint
type TeamVsPlayerRequest struct {
	TeamID     string
	VsPlayerID string
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// TeamVsPlayerOverall represents the Overall result set for TeamVsPlayer
type TeamVsPlayerOverall struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	TEAM_NAME  string  `json:"TEAM_NAME"`
	SORT_ORDER string  `json:"SORT_ORDER"`
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

// TeamVsPlayerOnOffCourt represents the OnOffCourt result set for TeamVsPlayer
type TeamVsPlayerOnOffCourt struct {
	TEAM_ID        int     `json:"TEAM_ID"`
	TEAM_NAME      string  `json:"TEAM_NAME"`
	SORT_ORDER     string  `json:"SORT_ORDER"`
	VS_PLAYER_ID   int     `json:"VS_PLAYER_ID"`
	VS_PLAYER_NAME string  `json:"VS_PLAYER_NAME"`
	COURT_STATUS   string  `json:"COURT_STATUS"`
	GP             int     `json:"GP"`
	W              string  `json:"W"`
	L              string  `json:"L"`
	W_PCT          float64 `json:"W_PCT"`
	MIN            float64 `json:"MIN"`
	FGM            int     `json:"FGM"`
	FGA            int     `json:"FGA"`
	FG_PCT         float64 `json:"FG_PCT"`
	FG3M           int     `json:"FG3M"`
	FG3A           int     `json:"FG3A"`
	FG3_PCT        float64 `json:"FG3_PCT"`
	FTM            int     `json:"FTM"`
	FTA            int     `json:"FTA"`
	FT_PCT         float64 `json:"FT_PCT"`
	OREB           float64 `json:"OREB"`
	DREB           float64 `json:"DREB"`
	REB            float64 `json:"REB"`
	AST            float64 `json:"AST"`
	TOV            float64 `json:"TOV"`
	STL            float64 `json:"STL"`
	BLK            float64 `json:"BLK"`
	BLKA           int     `json:"BLKA"`
	PF             float64 `json:"PF"`
	PFD            float64 `json:"PFD"`
	PTS            float64 `json:"PTS"`
	PLUS_MINUS     float64 `json:"PLUS_MINUS"`
}

// TeamVsPlayerResponse contains the response data from the TeamVsPlayer endpoint
type TeamVsPlayerResponse struct {
	Overall    []TeamVsPlayerOverall
	OnOffCourt []TeamVsPlayerOnOffCourt
}

// GetTeamVsPlayer retrieves data from the teamvsplayer endpoint
func GetTeamVsPlayer(ctx context.Context, client *stats.Client, req TeamVsPlayerRequest) (*models.Response[*TeamVsPlayerResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
	if req.VsPlayerID == "" {
		return nil, fmt.Errorf("VsPlayerID is required")
	}
	params.Set("VsPlayerID", string(req.VsPlayerID))
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamvsplayer", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamVsPlayerResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Overall = make([]TeamVsPlayerOverall, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 29 {
				item := TeamVsPlayerOverall{
					TEAM_ID:    toInt(row[0]),
					TEAM_NAME:  toString(row[1]),
					SORT_ORDER: toString(row[2]),
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
				response.Overall = append(response.Overall, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.OnOffCourt = make([]TeamVsPlayerOnOffCourt, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 32 {
				item := TeamVsPlayerOnOffCourt{
					TEAM_ID:        toInt(row[0]),
					TEAM_NAME:      toString(row[1]),
					SORT_ORDER:     toString(row[2]),
					VS_PLAYER_ID:   toInt(row[3]),
					VS_PLAYER_NAME: toString(row[4]),
					COURT_STATUS:   toString(row[5]),
					GP:             toInt(row[6]),
					W:              toString(row[7]),
					L:              toString(row[8]),
					W_PCT:          toFloat(row[9]),
					MIN:            toFloat(row[10]),
					FGM:            toInt(row[11]),
					FGA:            toInt(row[12]),
					FG_PCT:         toFloat(row[13]),
					FG3M:           toInt(row[14]),
					FG3A:           toInt(row[15]),
					FG3_PCT:        toFloat(row[16]),
					FTM:            toInt(row[17]),
					FTA:            toInt(row[18]),
					FT_PCT:         toFloat(row[19]),
					OREB:           toFloat(row[20]),
					DREB:           toFloat(row[21]),
					REB:            toFloat(row[22]),
					AST:            toFloat(row[23]),
					TOV:            toFloat(row[24]),
					STL:            toFloat(row[25]),
					BLK:            toFloat(row[26]),
					BLKA:           toInt(row[27]),
					PF:             toFloat(row[28]),
					PFD:            toFloat(row[29]),
					PTS:            toFloat(row[30]),
					PLUS_MINUS:     toFloat(row[31]),
				}
				response.OnOffCourt = append(response.OnOffCourt, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
