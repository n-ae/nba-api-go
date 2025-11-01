package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// CumeStatsPlayerRequest contains parameters for the CumeStatsPlayer endpoint
type CumeStatsPlayerRequest struct {
	PlayerID   string
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// CumeStatsPlayerGameByGameStats represents the GameByGameStats result set for CumeStatsPlayer
type CumeStatsPlayerGameByGameStats struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	SEASON_ID         string  `json:"SEASON_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GAME_ID           string  `json:"GAME_ID"`
	GAME_DATE         string  `json:"GAME_DATE"`
	MATCHUP           string  `json:"MATCHUP"`
	WL                string  `json:"WL"`
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
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	TOV               float64 `json:"TOV"`
	PF                float64 `json:"PF"`
	PTS               float64 `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// CumeStatsPlayerTotalStats represents the TotalStats result set for CumeStatsPlayer
type CumeStatsPlayerTotalStats struct {
	PLAYER_ID int     `json:"PLAYER_ID"`
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

// CumeStatsPlayerResponse contains the response data from the CumeStatsPlayer endpoint
type CumeStatsPlayerResponse struct {
	GameByGameStats []CumeStatsPlayerGameByGameStats
	TotalStats      []CumeStatsPlayerTotalStats
}

// GetCumeStatsPlayer retrieves data from the cumestatsplayer endpoint
func GetCumeStatsPlayer(ctx context.Context, client *stats.Client, req CumeStatsPlayerRequest) (*models.Response[*CumeStatsPlayerResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
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
	if err := client.GetJSON(ctx, "/cumestatsplayer", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CumeStatsPlayerResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.GameByGameStats = make([]CumeStatsPlayerGameByGameStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := CumeStatsPlayerGameByGameStats{
					PLAYER_ID:         toInt(row[0]),
					SEASON_ID:         toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GAME_ID:           toString(row[4]),
					GAME_DATE:         toString(row[5]),
					MATCHUP:           toString(row[6]),
					WL:                toString(row[7]),
					MIN:               toFloat(row[8]),
					FGM:               toInt(row[9]),
					FGA:               toInt(row[10]),
					FG_PCT:            toFloat(row[11]),
					FG3M:              toInt(row[12]),
					FG3A:              toInt(row[13]),
					FG3_PCT:           toFloat(row[14]),
					FTM:               toInt(row[15]),
					FTA:               toInt(row[16]),
					FT_PCT:            toFloat(row[17]),
					OREB:              toFloat(row[18]),
					DREB:              toFloat(row[19]),
					REB:               toFloat(row[20]),
					AST:               toFloat(row[21]),
					STL:               toFloat(row[22]),
					BLK:               toFloat(row[23]),
					TOV:               toFloat(row[24]),
					PF:                toFloat(row[25]),
					PTS:               toFloat(row[26]),
					PLUS_MINUS:        toFloat(row[27]),
				}
				response.GameByGameStats = append(response.GameByGameStats, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.TotalStats = make([]CumeStatsPlayerTotalStats, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 22 {
				item := CumeStatsPlayerTotalStats{
					PLAYER_ID: toInt(row[0]),
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
