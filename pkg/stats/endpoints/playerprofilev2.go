package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerProfileV2Request contains parameters for the PlayerProfileV2 endpoint
type PlayerProfileV2Request struct {
	PlayerID string
	PerMode *parameters.PerMode
	LeagueID *parameters.LeagueID
}


// PlayerProfileV2SeasonTotalsRegularSeason represents the SeasonTotalsRegularSeason result set for PlayerProfileV2
type PlayerProfileV2SeasonTotalsRegularSeason struct {
	PLAYER_ID int `json:"PLAYER_ID"`
	SEASON_ID string `json:"SEASON_ID"`
	LEAGUE_ID string `json:"LEAGUE_ID"`
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	PLAYER_AGE int `json:"PLAYER_AGE"`
	GP int `json:"GP"`
	GS int `json:"GS"`
	MIN float64 `json:"MIN"`
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
	STL float64 `json:"STL"`
	BLK float64 `json:"BLK"`
	TOV float64 `json:"TOV"`
	PF float64 `json:"PF"`
	PTS float64 `json:"PTS"`
}

// PlayerProfileV2CareerTotalsRegularSeason represents the CareerTotalsRegularSeason result set for PlayerProfileV2
type PlayerProfileV2CareerTotalsRegularSeason struct {
	PLAYER_ID int `json:"PLAYER_ID"`
	LEAGUE_ID string `json:"LEAGUE_ID"`
	Team_ID int `json:"Team_ID"`
	GP int `json:"GP"`
	GS int `json:"GS"`
	MIN float64 `json:"MIN"`
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
	STL float64 `json:"STL"`
	BLK float64 `json:"BLK"`
	TOV float64 `json:"TOV"`
	PF float64 `json:"PF"`
	PTS float64 `json:"PTS"`
}


// PlayerProfileV2Response contains the response data from the PlayerProfileV2 endpoint
type PlayerProfileV2Response struct {
	SeasonTotalsRegularSeason []PlayerProfileV2SeasonTotalsRegularSeason
	CareerTotalsRegularSeason []PlayerProfileV2CareerTotalsRegularSeason
}

// GetPlayerProfileV2 retrieves data from the playerprofilev2 endpoint
func GetPlayerProfileV2(ctx context.Context, client *stats.Client, req PlayerProfileV2Request) (*models.Response[*PlayerProfileV2Response], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playerprofilev2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerProfileV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.SeasonTotalsRegularSeason = make([]PlayerProfileV2SeasonTotalsRegularSeason, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 27 {
				item := PlayerProfileV2SeasonTotalsRegularSeason{
					PLAYER_ID: toInt(row[0]),
					SEASON_ID: toString(row[1]),
					LEAGUE_ID: toString(row[2]),
					TEAM_ID: toInt(row[3]),
					TEAM_ABBREVIATION: toString(row[4]),
					PLAYER_AGE: toInt(row[5]),
					GP: toInt(row[6]),
					GS: toInt(row[7]),
					MIN: toFloat(row[8]),
					FGM: toInt(row[9]),
					FGA: toInt(row[10]),
					FG_PCT: toFloat(row[11]),
					FG3M: toInt(row[12]),
					FG3A: toInt(row[13]),
					FG3_PCT: toFloat(row[14]),
					FTM: toInt(row[15]),
					FTA: toInt(row[16]),
					FT_PCT: toFloat(row[17]),
					OREB: toFloat(row[18]),
					DREB: toFloat(row[19]),
					REB: toFloat(row[20]),
					AST: toFloat(row[21]),
					STL: toFloat(row[22]),
					BLK: toFloat(row[23]),
					TOV: toFloat(row[24]),
					PF: toFloat(row[25]),
					PTS: toFloat(row[26]),
				}
				response.SeasonTotalsRegularSeason = append(response.SeasonTotalsRegularSeason, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.CareerTotalsRegularSeason = make([]PlayerProfileV2CareerTotalsRegularSeason, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 24 {
				item := PlayerProfileV2CareerTotalsRegularSeason{
					PLAYER_ID: toInt(row[0]),
					LEAGUE_ID: toString(row[1]),
					Team_ID: toInt(row[2]),
					GP: toInt(row[3]),
					GS: toInt(row[4]),
					MIN: toFloat(row[5]),
					FGM: toInt(row[6]),
					FGA: toInt(row[7]),
					FG_PCT: toFloat(row[8]),
					FG3M: toInt(row[9]),
					FG3A: toInt(row[10]),
					FG3_PCT: toFloat(row[11]),
					FTM: toInt(row[12]),
					FTA: toInt(row[13]),
					FT_PCT: toFloat(row[14]),
					OREB: toFloat(row[15]),
					DREB: toFloat(row[16]),
					REB: toFloat(row[17]),
					AST: toFloat(row[18]),
					STL: toFloat(row[19]),
					BLK: toFloat(row[20]),
					TOV: toFloat(row[21]),
					PF: toFloat(row[22]),
					PTS: toFloat(row[23]),
				}
				response.CareerTotalsRegularSeason = append(response.CareerTotalsRegularSeason, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
