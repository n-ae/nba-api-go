package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// CumeStatsTeamRequest contains parameters for the CumeStatsTeam endpoint
type CumeStatsTeamRequest struct {
	TeamID     string
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// CumeStatsTeamGameByGameStats represents the GameByGameStats result set for CumeStatsTeam
type CumeStatsTeamGameByGameStats struct {
	TEAM_ID    int     `json:"TEAM_ID"`
	SEASON_ID  string  `json:"SEASON_ID"`
	GAME_ID    string  `json:"GAME_ID"`
	GAME_DATE  string  `json:"GAME_DATE"`
	MATCHUP    string  `json:"MATCHUP"`
	WL         string  `json:"WL"`
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
	STL        float64 `json:"STL"`
	BLK        float64 `json:"BLK"`
	TOV        float64 `json:"TOV"`
	PF         float64 `json:"PF"`
	PTS        float64 `json:"PTS"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}

// CumeStatsTeamTotalStats represents the TotalStats result set for CumeStatsTeam
type CumeStatsTeamTotalStats struct {
	TEAM_ID   int     `json:"TEAM_ID"`
	SEASON_ID string  `json:"SEASON_ID"`
	GP        int     `json:"GP"`
	MIN       float64 `json:"MIN"`
	FGM       int     `json:"FGM"`
	FGA       int     `json:"FGA"`
	FG_PCT    float64 `json:"FG_PCT"`
	FG3M      int     `json:"FG3M"`
	FG3A      int     `json:"FG3A"`
	FG3_PCT   float64 `json:"FG3_PCT"`
	FTM       int     `json:"FTM"`
	FTA       int     `json:"FTA"`
	FT_PCT    float64 `json:"FT_PCT"`
	OREB      float64 `json:"OREB"`
	DREB      float64 `json:"DREB"`
	REB       float64 `json:"REB"`
	AST       float64 `json:"AST"`
	STL       float64 `json:"STL"`
	BLK       float64 `json:"BLK"`
	TOV       float64 `json:"TOV"`
	PF        float64 `json:"PF"`
	PTS       float64 `json:"PTS"`
}

// CumeStatsTeamResponse contains the response data from the CumeStatsTeam endpoint
type CumeStatsTeamResponse struct {
	GameByGameStats []CumeStatsTeamGameByGameStats
	TotalStats      []CumeStatsTeamTotalStats
}

// GetCumeStatsTeam retrieves data from the cumestatsteam endpoint
func GetCumeStatsTeam(ctx context.Context, client *stats.Client, req CumeStatsTeamRequest) (*models.Response[*CumeStatsTeamResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
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
	if err := client.GetJSON(ctx, "/cumestatsteam", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CumeStatsTeamResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.GameByGameStats = make([]CumeStatsTeamGameByGameStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 26 {
				item := CumeStatsTeamGameByGameStats{
					TEAM_ID:    toInt(row[0]),
					SEASON_ID:  toString(row[1]),
					GAME_ID:    toString(row[2]),
					GAME_DATE:  toString(row[3]),
					MATCHUP:    toString(row[4]),
					WL:         toString(row[5]),
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
					STL:        toFloat(row[20]),
					BLK:        toFloat(row[21]),
					TOV:        toFloat(row[22]),
					PF:         toFloat(row[23]),
					PTS:        toFloat(row[24]),
					PLUS_MINUS: toFloat(row[25]),
				}
				response.GameByGameStats = append(response.GameByGameStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TotalStats = make([]CumeStatsTeamTotalStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 22 {
				item := CumeStatsTeamTotalStats{
					TEAM_ID:   toInt(row[0]),
					SEASON_ID: toString(row[1]),
					GP:        toInt(row[2]),
					MIN:       toFloat(row[3]),
					FGM:       toInt(row[4]),
					FGA:       toInt(row[5]),
					FG_PCT:    toFloat(row[6]),
					FG3M:      toInt(row[7]),
					FG3A:      toInt(row[8]),
					FG3_PCT:   toFloat(row[9]),
					FTM:       toInt(row[10]),
					FTA:       toInt(row[11]),
					FT_PCT:    toFloat(row[12]),
					OREB:      toFloat(row[13]),
					DREB:      toFloat(row[14]),
					REB:       toFloat(row[15]),
					AST:       toFloat(row[16]),
					STL:       toFloat(row[17]),
					BLK:       toFloat(row[18]),
					TOV:       toFloat(row[19]),
					PF:        toFloat(row[20]),
					PTS:       toFloat(row[21]),
				}
				response.TotalStats = append(response.TotalStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
