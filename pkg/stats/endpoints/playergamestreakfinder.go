package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerGameStreakFinderRequest contains parameters for the PlayerGameStreakFinder endpoint
type PlayerGameStreakFinderRequest struct {
	Season *parameters.Season
	SeasonType *parameters.SeasonType
	LeagueID *parameters.LeagueID
}


// PlayerGameStreakFinderPlayerGameStreakFinder represents the PlayerGameStreakFinder result set for PlayerGameStreakFinder
type PlayerGameStreakFinderPlayerGameStreakFinder struct {
	PLAYER_ID int `json:"PLAYER_ID"`
	PLAYER_NAME string `json:"PLAYER_NAME"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	GAME_ID string `json:"GAME_ID"`
	GAME_DATE string `json:"GAME_DATE"`
	MATCHUP string `json:"MATCHUP"`
	WL string `json:"WL"`
	MIN float64 `json:"MIN"`
	PTS float64 `json:"PTS"`
	FGM int `json:"FGM"`
	FGA int `json:"FGA"`
	FG_PCT float64 `json:"FG_PCT"`
	FG3M int `json:"FG3M"`
	FG3A int `json:"FG3A"`
	FG3_PCT float64 `json:"FG3_PCT"`
	FTM int `json:"FTM"`
	FTA int `json:"FTA"`
	FT_PCT float64 `json:"FT_PCT"`
	OREB float64 `json:"OREB"`
	DREB float64 `json:"DREB"`
	REB float64 `json:"REB"`
	AST float64 `json:"AST"`
	TOV float64 `json:"TOV"`
	STL float64 `json:"STL"`
	BLK float64 `json:"BLK"`
	PF float64 `json:"PF"`
	PLUS_MINUS float64 `json:"PLUS_MINUS"`
}


// PlayerGameStreakFinderResponse contains the response data from the PlayerGameStreakFinder endpoint
type PlayerGameStreakFinderResponse struct {
	PlayerGameStreakFinder []PlayerGameStreakFinderPlayerGameStreakFinder
}

// GetPlayerGameStreakFinder retrieves data from the playergamestreakfinder endpoint
func GetPlayerGameStreakFinder(ctx context.Context, client *stats.Client, req PlayerGameStreakFinderRequest) (*models.Response[*PlayerGameStreakFinderResponse], error) {
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
	if err := client.GetJSON(ctx, "/playergamestreakfinder", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerGameStreakFinderResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerGameStreakFinder = make([]PlayerGameStreakFinderPlayerGameStreakFinder, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := PlayerGameStreakFinderPlayerGameStreakFinder{
					PLAYER_ID: toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					TEAM_ID: toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GAME_ID: toString(row[4]),
					GAME_DATE: toString(row[5]),
					MATCHUP: toString(row[6]),
					WL: toString(row[7]),
					MIN: toFloat(row[8]),
					PTS: toFloat(row[9]),
					FGM: toInt(row[10]),
					FGA: toInt(row[11]),
					FG_PCT: toFloat(row[12]),
					FG3M: toInt(row[13]),
					FG3A: toInt(row[14]),
					FG3_PCT: toFloat(row[15]),
					FTM: toInt(row[16]),
					FTA: toInt(row[17]),
					FT_PCT: toFloat(row[18]),
					OREB: toFloat(row[19]),
					DREB: toFloat(row[20]),
					REB: toFloat(row[21]),
					AST: toFloat(row[22]),
					TOV: toFloat(row[23]),
					STL: toFloat(row[24]),
					BLK: toFloat(row[25]),
					PF: toFloat(row[26]),
					PLUS_MINUS: toFloat(row[27]),
				}
				response.PlayerGameStreakFinder = append(response.PlayerGameStreakFinder, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
