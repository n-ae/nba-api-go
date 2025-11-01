package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerGameLogsRequest contains parameters for the PlayerGameLogs endpoint
type PlayerGameLogsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
	DateFrom   *string
	DateTo     *string
}

// PlayerGameLogsPlayerGameLogs represents the PlayerGameLogs result set for PlayerGameLogs
type PlayerGameLogsPlayerGameLogs struct {
	SEASON_YEAR       string  `json:"SEASON_YEAR"`
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	NICKNAME          string  `json:"NICKNAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
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
	NBA_FANTASY_PTS   float64 `json:"NBA_FANTASY_PTS"`
	DD2               float64 `json:"DD2"`
	TD3               float64 `json:"TD3"`
}

// PlayerGameLogsResponse contains the response data from the PlayerGameLogs endpoint
type PlayerGameLogsResponse struct {
	PlayerGameLogs []PlayerGameLogsPlayerGameLogs
}

// GetPlayerGameLogs retrieves data from the playergamelogs endpoint
func GetPlayerGameLogs(ctx context.Context, client *stats.Client, req PlayerGameLogsRequest) (*models.Response[*PlayerGameLogsResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", string(*req.DateFrom))
	}
	if req.DateTo != nil {
		params.Set("DateTo", string(*req.DateTo))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playergamelogs", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerGameLogsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerGameLogs = make([]PlayerGameLogsPlayerGameLogs, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 34 {
				item := PlayerGameLogsPlayerGameLogs{
					SEASON_YEAR:       toString(row[0]),
					PLAYER_ID:         toInt(row[1]),
					PLAYER_NAME:       toString(row[2]),
					NICKNAME:          toString(row[3]),
					TEAM_ID:           toInt(row[4]),
					TEAM_ABBREVIATION: toString(row[5]),
					TEAM_NAME:         toString(row[6]),
					GAME_ID:           toString(row[7]),
					GAME_DATE:         toString(row[8]),
					MATCHUP:           toString(row[9]),
					WL:                toString(row[10]),
					MIN:               toFloat(row[11]),
					FGM:               toInt(row[12]),
					FGA:               toInt(row[13]),
					FG_PCT:            toFloat(row[14]),
					FG3M:              toInt(row[15]),
					FG3A:              toInt(row[16]),
					FG3_PCT:           toFloat(row[17]),
					FTM:               toInt(row[18]),
					FTA:               toInt(row[19]),
					FT_PCT:            toFloat(row[20]),
					OREB:              toFloat(row[21]),
					DREB:              toFloat(row[22]),
					REB:               toFloat(row[23]),
					AST:               toFloat(row[24]),
					STL:               toFloat(row[25]),
					BLK:               toFloat(row[26]),
					TOV:               toFloat(row[27]),
					PF:                toFloat(row[28]),
					PTS:               toFloat(row[29]),
					PLUS_MINUS:        toFloat(row[30]),
					NBA_FANTASY_PTS:   toFloat(row[31]),
					DD2:               toFloat(row[32]),
					TD3:               toFloat(row[33]),
				}
				response.PlayerGameLogs = append(response.PlayerGameLogs, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
