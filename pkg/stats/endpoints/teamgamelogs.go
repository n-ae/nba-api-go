package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamGameLogsRequest contains parameters for the TeamGameLogs endpoint
type TeamGameLogsRequest struct {
	Season     parameters.Season
	SeasonType parameters.SeasonType
	LeagueID   *parameters.LeagueID
	TeamID     *string
	DateFrom   *string
	DateTo     *string
}

// TeamGameLogsTeamGameLogs represents the TeamGameLogs result set for TeamGameLogs
type TeamGameLogsTeamGameLogs struct {
	SEASON_YEAR       string  `json:"SEASON_YEAR"`
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
	OREB              int     `json:"OREB"`
	DREB              int     `json:"DREB"`
	REB               int     `json:"REB"`
	AST               int     `json:"AST"`
	TOV               int     `json:"TOV"`
	STL               int     `json:"STL"`
	BLK               int     `json:"BLK"`
	BLKA              int     `json:"BLKA"`
	PF                int     `json:"PF"`
	PFD               int     `json:"PFD"`
	PTS               int     `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
	NBA_FANTASY_PTS   float64 `json:"NBA_FANTASY_PTS"`
	DD2               int     `json:"DD2"`
	TD3               int     `json:"TD3"`
}

// TeamGameLogsResponse contains the response data from the TeamGameLogs endpoint
type TeamGameLogsResponse struct {
	TeamGameLogs []TeamGameLogsTeamGameLogs
}

// GetTeamGameLogs retrieves data from the teamgamelogs endpoint
func GetTeamGameLogs(ctx context.Context, client *stats.Client, req TeamGameLogsRequest) (*models.Response[*TeamGameLogsResponse], error) {
	params := url.Values{}
	if req.Season == "" {
		return nil, fmt.Errorf("Season is required")
	}
	params.Set("Season", string(req.Season))
	if req.SeasonType == "" {
		return nil, fmt.Errorf("SeasonType is required")
	}
	params.Set("SeasonType", string(req.SeasonType))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.TeamID != nil {
		params.Set("TeamID", string(*req.TeamID))
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", string(*req.DateFrom))
	}
	if req.DateTo != nil {
		params.Set("DateTo", string(*req.DateTo))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamgamelogs", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamGameLogsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamGameLogs = make([]TeamGameLogsTeamGameLogs, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 33 {
				item := TeamGameLogsTeamGameLogs{
					SEASON_YEAR:       toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_NAME:         toString(row[3]),
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
					OREB:              toInt(row[18]),
					DREB:              toInt(row[19]),
					REB:               toInt(row[20]),
					AST:               toInt(row[21]),
					TOV:               toInt(row[22]),
					STL:               toInt(row[23]),
					BLK:               toInt(row[24]),
					BLKA:              toInt(row[25]),
					PF:                toInt(row[26]),
					PFD:               toInt(row[27]),
					PTS:               toInt(row[28]),
					PLUS_MINUS:        toFloat(row[29]),
					NBA_FANTASY_PTS:   toFloat(row[30]),
					DD2:               toInt(row[31]),
					TD3:               toInt(row[32]),
				}
				response.TeamGameLogs = append(response.TeamGameLogs, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
