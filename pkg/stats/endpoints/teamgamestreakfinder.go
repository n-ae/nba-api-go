package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// TeamGameStreakFinderRequest contains parameters for the TeamGameStreakFinder endpoint
type TeamGameStreakFinderRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID   *parameters.LeagueID
}

// TeamGameStreakFinderTeamGameStreakFinder represents the TeamGameStreakFinder result set for TeamGameStreakFinder
type TeamGameStreakFinderTeamGameStreakFinder struct {
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GAME_ID           string  `json:"GAME_ID"`
	GAME_DATE         string  `json:"GAME_DATE"`
	MATCHUP           string  `json:"MATCHUP"`
	WL                string  `json:"WL"`
	MIN               float64 `json:"MIN"`
	PTS               float64 `json:"PTS"`
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
	TOV               float64 `json:"TOV"`
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	PF                float64 `json:"PF"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// TeamGameStreakFinderResponse contains the response data from the TeamGameStreakFinder endpoint
type TeamGameStreakFinderResponse struct {
	TeamGameStreakFinder []TeamGameStreakFinderTeamGameStreakFinder
}

// GetTeamGameStreakFinder retrieves data from the teamgamestreakfinder endpoint
func GetTeamGameStreakFinder(ctx context.Context, client *stats.Client, req TeamGameStreakFinderRequest) (*models.Response[*TeamGameStreakFinderResponse], error) {
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

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/teamgamestreakfinder", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamGameStreakFinderResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.TeamGameStreakFinder = make([]TeamGameStreakFinderTeamGameStreakFinder, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 27 {
				item := TeamGameStreakFinderTeamGameStreakFinder{
					TEAM_ID:           toInt(row[0]),
					TEAM_NAME:         toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					GAME_ID:           toString(row[3]),
					GAME_DATE:         toString(row[4]),
					MATCHUP:           toString(row[5]),
					WL:                toString(row[6]),
					MIN:               toFloat(row[7]),
					PTS:               toFloat(row[8]),
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
					TOV:               toFloat(row[22]),
					STL:               toFloat(row[23]),
					BLK:               toFloat(row[24]),
					PF:                toFloat(row[25]),
					PLUS_MINUS:        toFloat(row[26]),
				}
				response.TeamGameStreakFinder = append(response.TeamGameStreakFinder, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
