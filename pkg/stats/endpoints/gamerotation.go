package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// GameRotationRequest contains parameters for the GameRotation endpoint
type GameRotationRequest struct {
	GameID   string
	LeagueID *parameters.LeagueID
}

// GameRotationAwayTeam represents the AwayTeam result set for GameRotation
type GameRotationAwayTeam struct {
	GAME_ID       string  `json:"GAME_ID"`
	TEAM_ID       int     `json:"TEAM_ID"`
	TEAM_NAME     string  `json:"TEAM_NAME"`
	PERSON_ID     string  `json:"PERSON_ID"`
	PLAYER_FIRST  string  `json:"PLAYER_FIRST"`
	PLAYER_LAST   float64 `json:"PLAYER_LAST"`
	IN_TIME_REAL  string  `json:"IN_TIME_REAL"`
	OUT_TIME_REAL string  `json:"OUT_TIME_REAL"`
	PLAYER_PTS    float64 `json:"PLAYER_PTS"`
	PT_DIFF       string  `json:"PT_DIFF"`
	USG_PCT       float64 `json:"USG_PCT"`
}

// GameRotationHomeTeam represents the HomeTeam result set for GameRotation
type GameRotationHomeTeam struct {
	GAME_ID       string  `json:"GAME_ID"`
	TEAM_ID       int     `json:"TEAM_ID"`
	TEAM_NAME     string  `json:"TEAM_NAME"`
	PERSON_ID     string  `json:"PERSON_ID"`
	PLAYER_FIRST  string  `json:"PLAYER_FIRST"`
	PLAYER_LAST   float64 `json:"PLAYER_LAST"`
	IN_TIME_REAL  string  `json:"IN_TIME_REAL"`
	OUT_TIME_REAL string  `json:"OUT_TIME_REAL"`
	PLAYER_PTS    float64 `json:"PLAYER_PTS"`
	PT_DIFF       string  `json:"PT_DIFF"`
	USG_PCT       float64 `json:"USG_PCT"`
}

// GameRotationResponse contains the response data from the GameRotation endpoint
type GameRotationResponse struct {
	AwayTeam []GameRotationAwayTeam
	HomeTeam []GameRotationHomeTeam
}

// GetGameRotation retrieves data from the gamerotation endpoint
func GetGameRotation(ctx context.Context, client *stats.Client, req GameRotationRequest) (*models.Response[*GameRotationResponse], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "gamerotation", params, &rawResp); err != nil {
		return nil, err
	}

	response := &GameRotationResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.AwayTeam = make([]GameRotationAwayTeam, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 11 {
				item := GameRotationAwayTeam{
					GAME_ID:       toString(row[0]),
					TEAM_ID:       toInt(row[1]),
					TEAM_NAME:     toString(row[2]),
					PERSON_ID:     toString(row[3]),
					PLAYER_FIRST:  toString(row[4]),
					PLAYER_LAST:   toFloat(row[5]),
					IN_TIME_REAL:  toString(row[6]),
					OUT_TIME_REAL: toString(row[7]),
					PLAYER_PTS:    toFloat(row[8]),
					PT_DIFF:       toString(row[9]),
					USG_PCT:       toFloat(row[10]),
				}
				response.AwayTeam = append(response.AwayTeam, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.HomeTeam = make([]GameRotationHomeTeam, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 11 {
				item := GameRotationHomeTeam{
					GAME_ID:       toString(row[0]),
					TEAM_ID:       toInt(row[1]),
					TEAM_NAME:     toString(row[2]),
					PERSON_ID:     toString(row[3]),
					PLAYER_FIRST:  toString(row[4]),
					PLAYER_LAST:   toFloat(row[5]),
					IN_TIME_REAL:  toString(row[6]),
					OUT_TIME_REAL: toString(row[7]),
					PLAYER_PTS:    toFloat(row[8]),
					PT_DIFF:       toString(row[9]),
					USG_PCT:       toFloat(row[10]),
				}
				response.HomeTeam = append(response.HomeTeam, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
