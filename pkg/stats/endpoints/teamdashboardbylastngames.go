package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamDashboardByLastNGamesRequest contains parameters for the TeamDashboardByLastNGames endpoint
type TeamDashboardByLastNGamesRequest struct {
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

// TeamDashboardByLastNGamesOverallTeamDashboard represents the OverallTeamDashboard result set for TeamDashboardByLastNGames
type TeamDashboardByLastNGamesOverallTeamDashboard struct {
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

// TeamDashboardByLastNGamesLast5TeamDashboard represents the Last5TeamDashboard result set for TeamDashboardByLastNGames
type TeamDashboardByLastNGamesLast5TeamDashboard struct {
	TEAM_ID      int     `json:"TEAM_ID"`
	TEAM_NAME    string  `json:"TEAM_NAME"`
	LAST_N_GAMES float64 `json:"LAST_N_GAMES"`
	GP           int     `json:"GP"`
	W            string  `json:"W"`
	L            string  `json:"L"`
	W_PCT        float64 `json:"W_PCT"`
	MIN          float64 `json:"MIN"`
	FGM          int     `json:"FGM"`
	FGA          int     `json:"FGA"`
	FG_PCT       float64 `json:"FG_PCT"`
	FG3M         int     `json:"FG3M"`
	FG3A         int     `json:"FG3A"`
	FG3_PCT      float64 `json:"FG3_PCT"`
	FTM          int     `json:"FTM"`
	FTA          int     `json:"FTA"`
	FT_PCT       float64 `json:"FT_PCT"`
	OREB         float64 `json:"OREB"`
	DREB         float64 `json:"DREB"`
	REB          float64 `json:"REB"`
	AST          float64 `json:"AST"`
	TOV          float64 `json:"TOV"`
	STL          float64 `json:"STL"`
	BLK          float64 `json:"BLK"`
	BLKA         int     `json:"BLKA"`
	PF           float64 `json:"PF"`
	PFD          float64 `json:"PFD"`
	PTS          float64 `json:"PTS"`
	PLUS_MINUS   float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByLastNGamesLast10TeamDashboard represents the Last10TeamDashboard result set for TeamDashboardByLastNGames
type TeamDashboardByLastNGamesLast10TeamDashboard struct {
	TEAM_ID      int     `json:"TEAM_ID"`
	TEAM_NAME    string  `json:"TEAM_NAME"`
	LAST_N_GAMES float64 `json:"LAST_N_GAMES"`
	GP           int     `json:"GP"`
	W            string  `json:"W"`
	L            string  `json:"L"`
	W_PCT        float64 `json:"W_PCT"`
	MIN          float64 `json:"MIN"`
	FGM          int     `json:"FGM"`
	FGA          int     `json:"FGA"`
	FG_PCT       float64 `json:"FG_PCT"`
	FG3M         int     `json:"FG3M"`
	FG3A         int     `json:"FG3A"`
	FG3_PCT      float64 `json:"FG3_PCT"`
	FTM          int     `json:"FTM"`
	FTA          int     `json:"FTA"`
	FT_PCT       float64 `json:"FT_PCT"`
	OREB         float64 `json:"OREB"`
	DREB         float64 `json:"DREB"`
	REB          float64 `json:"REB"`
	AST          float64 `json:"AST"`
	TOV          float64 `json:"TOV"`
	STL          float64 `json:"STL"`
	BLK          float64 `json:"BLK"`
	BLKA         int     `json:"BLKA"`
	PF           float64 `json:"PF"`
	PFD          float64 `json:"PFD"`
	PTS          float64 `json:"PTS"`
	PLUS_MINUS   float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByLastNGamesLast15TeamDashboard represents the Last15TeamDashboard result set for TeamDashboardByLastNGames
type TeamDashboardByLastNGamesLast15TeamDashboard struct {
	TEAM_ID      int     `json:"TEAM_ID"`
	TEAM_NAME    string  `json:"TEAM_NAME"`
	LAST_N_GAMES float64 `json:"LAST_N_GAMES"`
	GP           int     `json:"GP"`
	W            string  `json:"W"`
	L            string  `json:"L"`
	W_PCT        float64 `json:"W_PCT"`
	MIN          float64 `json:"MIN"`
	FGM          int     `json:"FGM"`
	FGA          int     `json:"FGA"`
	FG_PCT       float64 `json:"FG_PCT"`
	FG3M         int     `json:"FG3M"`
	FG3A         int     `json:"FG3A"`
	FG3_PCT      float64 `json:"FG3_PCT"`
	FTM          int     `json:"FTM"`
	FTA          int     `json:"FTA"`
	FT_PCT       float64 `json:"FT_PCT"`
	OREB         float64 `json:"OREB"`
	DREB         float64 `json:"DREB"`
	REB          float64 `json:"REB"`
	AST          float64 `json:"AST"`
	TOV          float64 `json:"TOV"`
	STL          float64 `json:"STL"`
	BLK          float64 `json:"BLK"`
	BLKA         int     `json:"BLKA"`
	PF           float64 `json:"PF"`
	PFD          float64 `json:"PFD"`
	PTS          float64 `json:"PTS"`
	PLUS_MINUS   float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByLastNGamesLast20TeamDashboard represents the Last20TeamDashboard result set for TeamDashboardByLastNGames
type TeamDashboardByLastNGamesLast20TeamDashboard struct {
	TEAM_ID      int     `json:"TEAM_ID"`
	TEAM_NAME    string  `json:"TEAM_NAME"`
	LAST_N_GAMES float64 `json:"LAST_N_GAMES"`
	GP           int     `json:"GP"`
	W            string  `json:"W"`
	L            string  `json:"L"`
	W_PCT        float64 `json:"W_PCT"`
	MIN          float64 `json:"MIN"`
	FGM          int     `json:"FGM"`
	FGA          int     `json:"FGA"`
	FG_PCT       float64 `json:"FG_PCT"`
	FG3M         int     `json:"FG3M"`
	FG3A         int     `json:"FG3A"`
	FG3_PCT      float64 `json:"FG3_PCT"`
	FTM          int     `json:"FTM"`
	FTA          int     `json:"FTA"`
	FT_PCT       float64 `json:"FT_PCT"`
	OREB         float64 `json:"OREB"`
	DREB         float64 `json:"DREB"`
	REB          float64 `json:"REB"`
	AST          float64 `json:"AST"`
	TOV          float64 `json:"TOV"`
	STL          float64 `json:"STL"`
	BLK          float64 `json:"BLK"`
	BLKA         int     `json:"BLKA"`
	PF           float64 `json:"PF"`
	PFD          float64 `json:"PFD"`
	PTS          float64 `json:"PTS"`
	PLUS_MINUS   float64 `json:"PLUS_MINUS"`
}

// TeamDashboardByLastNGamesResponse contains the response data from the TeamDashboardByLastNGames endpoint
type TeamDashboardByLastNGamesResponse struct {
	OverallTeamDashboard []TeamDashboardByLastNGamesOverallTeamDashboard
	Last5TeamDashboard   []TeamDashboardByLastNGamesLast5TeamDashboard
	Last10TeamDashboard  []TeamDashboardByLastNGamesLast10TeamDashboard
	Last15TeamDashboard  []TeamDashboardByLastNGamesLast15TeamDashboard
	Last20TeamDashboard  []TeamDashboardByLastNGamesLast20TeamDashboard
}

// GetTeamDashboardByLastNGames retrieves data from the teamdashboardbylastnGames endpoint
func GetTeamDashboardByLastNGames(ctx context.Context, client *stats.Client, req TeamDashboardByLastNGamesRequest) (*models.Response[*TeamDashboardByLastNGamesResponse], error) {
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
	if err := client.GetJSON(ctx, "/teamdashboardbylastnGames", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDashboardByLastNGamesResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallTeamDashboard = make([]TeamDashboardByLastNGamesOverallTeamDashboard, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := TeamDashboardByLastNGamesOverallTeamDashboard{
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
		response.Last5TeamDashboard = make([]TeamDashboardByLastNGamesLast5TeamDashboard, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByLastNGamesLast5TeamDashboard{
					TEAM_ID:      toInt(row[0]),
					TEAM_NAME:    toString(row[1]),
					LAST_N_GAMES: toFloat(row[2]),
					GP:           toInt(row[3]),
					W:            toString(row[4]),
					L:            toString(row[5]),
					W_PCT:        toFloat(row[6]),
					MIN:          toFloat(row[7]),
					FGM:          toInt(row[8]),
					FGA:          toInt(row[9]),
					FG_PCT:       toFloat(row[10]),
					FG3M:         toInt(row[11]),
					FG3A:         toInt(row[12]),
					FG3_PCT:      toFloat(row[13]),
					FTM:          toInt(row[14]),
					FTA:          toInt(row[15]),
					FT_PCT:       toFloat(row[16]),
					OREB:         toFloat(row[17]),
					DREB:         toFloat(row[18]),
					REB:          toFloat(row[19]),
					AST:          toFloat(row[20]),
					TOV:          toFloat(row[21]),
					STL:          toFloat(row[22]),
					BLK:          toFloat(row[23]),
					BLKA:         toInt(row[24]),
					PF:           toFloat(row[25]),
					PFD:          toFloat(row[26]),
					PTS:          toFloat(row[27]),
					PLUS_MINUS:   toFloat(row[28]),
				}
				response.Last5TeamDashboard = append(response.Last5TeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.Last10TeamDashboard = make([]TeamDashboardByLastNGamesLast10TeamDashboard, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByLastNGamesLast10TeamDashboard{
					TEAM_ID:      toInt(row[0]),
					TEAM_NAME:    toString(row[1]),
					LAST_N_GAMES: toFloat(row[2]),
					GP:           toInt(row[3]),
					W:            toString(row[4]),
					L:            toString(row[5]),
					W_PCT:        toFloat(row[6]),
					MIN:          toFloat(row[7]),
					FGM:          toInt(row[8]),
					FGA:          toInt(row[9]),
					FG_PCT:       toFloat(row[10]),
					FG3M:         toInt(row[11]),
					FG3A:         toInt(row[12]),
					FG3_PCT:      toFloat(row[13]),
					FTM:          toInt(row[14]),
					FTA:          toInt(row[15]),
					FT_PCT:       toFloat(row[16]),
					OREB:         toFloat(row[17]),
					DREB:         toFloat(row[18]),
					REB:          toFloat(row[19]),
					AST:          toFloat(row[20]),
					TOV:          toFloat(row[21]),
					STL:          toFloat(row[22]),
					BLK:          toFloat(row[23]),
					BLKA:         toInt(row[24]),
					PF:           toFloat(row[25]),
					PFD:          toFloat(row[26]),
					PTS:          toFloat(row[27]),
					PLUS_MINUS:   toFloat(row[28]),
				}
				response.Last10TeamDashboard = append(response.Last10TeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.Last15TeamDashboard = make([]TeamDashboardByLastNGamesLast15TeamDashboard, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByLastNGamesLast15TeamDashboard{
					TEAM_ID:      toInt(row[0]),
					TEAM_NAME:    toString(row[1]),
					LAST_N_GAMES: toFloat(row[2]),
					GP:           toInt(row[3]),
					W:            toString(row[4]),
					L:            toString(row[5]),
					W_PCT:        toFloat(row[6]),
					MIN:          toFloat(row[7]),
					FGM:          toInt(row[8]),
					FGA:          toInt(row[9]),
					FG_PCT:       toFloat(row[10]),
					FG3M:         toInt(row[11]),
					FG3A:         toInt(row[12]),
					FG3_PCT:      toFloat(row[13]),
					FTM:          toInt(row[14]),
					FTA:          toInt(row[15]),
					FT_PCT:       toFloat(row[16]),
					OREB:         toFloat(row[17]),
					DREB:         toFloat(row[18]),
					REB:          toFloat(row[19]),
					AST:          toFloat(row[20]),
					TOV:          toFloat(row[21]),
					STL:          toFloat(row[22]),
					BLK:          toFloat(row[23]),
					BLKA:         toInt(row[24]),
					PF:           toFloat(row[25]),
					PFD:          toFloat(row[26]),
					PTS:          toFloat(row[27]),
					PLUS_MINUS:   toFloat(row[28]),
				}
				response.Last15TeamDashboard = append(response.Last15TeamDashboard, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.Last20TeamDashboard = make([]TeamDashboardByLastNGamesLast20TeamDashboard, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 29 {
				item := TeamDashboardByLastNGamesLast20TeamDashboard{
					TEAM_ID:      toInt(row[0]),
					TEAM_NAME:    toString(row[1]),
					LAST_N_GAMES: toFloat(row[2]),
					GP:           toInt(row[3]),
					W:            toString(row[4]),
					L:            toString(row[5]),
					W_PCT:        toFloat(row[6]),
					MIN:          toFloat(row[7]),
					FGM:          toInt(row[8]),
					FGA:          toInt(row[9]),
					FG_PCT:       toFloat(row[10]),
					FG3M:         toInt(row[11]),
					FG3A:         toInt(row[12]),
					FG3_PCT:      toFloat(row[13]),
					FTM:          toInt(row[14]),
					FTA:          toInt(row[15]),
					FT_PCT:       toFloat(row[16]),
					OREB:         toFloat(row[17]),
					DREB:         toFloat(row[18]),
					REB:          toFloat(row[19]),
					AST:          toFloat(row[20]),
					TOV:          toFloat(row[21]),
					STL:          toFloat(row[22]),
					BLK:          toFloat(row[23]),
					BLKA:         toInt(row[24]),
					PF:           toFloat(row[25]),
					PFD:          toFloat(row[26]),
					PTS:          toFloat(row[27]),
					PLUS_MINUS:   toFloat(row[28]),
				}
				response.Last20TeamDashboard = append(response.Last20TeamDashboard, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
