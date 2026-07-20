package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// ScoreboardV3Request contains parameters for the ScoreboardV3 endpoint
type ScoreboardV3Request struct {
	GameDate string
	LeagueID *parameters.LeagueID
}

// ScoreboardV3GameHeader represents the GameHeader result set for ScoreboardV3
type ScoreboardV3GameHeader struct {
	GameID             string `json:"gameId"`
	GameCode           string `json:"gameCode"`
	GameStatus         string `json:"gameStatus"`
	GameStatusText     string `json:"gameStatusText"`
	Period             int    `json:"period"`
	GameClock          string `json:"gameClock"`
	GameTimeUTC        string `json:"gameTimeUTC"`
	GameEt             string `json:"gameEt"`
	RegulationPeriods  int    `json:"regulationPeriods"`
	SeriesGameNumber   string `json:"seriesGameNumber"`
	SeriesText         string `json:"seriesText"`
	HomeTeamID         string `json:"homeTeamId"`
	HomeTeamName       string `json:"homeTeamName"`
	HomeTeamCity       string `json:"homeTeamCity"`
	HomeTeamTricode    string `json:"homeTeamTricode"`
	HomeTeamScore      string `json:"homeTeamScore"`
	VisitorTeamID      string `json:"visitorTeamId"`
	VisitorTeamName    string `json:"visitorTeamName"`
	VisitorTeamCity    string `json:"visitorTeamCity"`
	VisitorTeamTricode string `json:"visitorTeamTricode"`
	VisitorTeamScore   string `json:"visitorTeamScore"`
}

// ScoreboardV3Response contains the response data from the ScoreboardV3 endpoint
type ScoreboardV3Response struct {
	GameHeader []ScoreboardV3GameHeader
}

// GetScoreboardV3 retrieves data from the scoreboardv3 endpoint
func GetScoreboardV3(ctx context.Context, client *stats.Client, req ScoreboardV3Request) (*models.Response[*ScoreboardV3Response], error) {
	params := url.Values{}
	if req.GameDate == "" {
		return nil, fmt.Errorf("%s is required", "GameDate")
	}
	params.Set("GameDate", req.GameDate)
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "scoreboardv3", params, &rawResp); err != nil {
		return nil, err
	}

	response := &ScoreboardV3Response{}
	if len(rawResp.ResultSets) > 0 {
		response.GameHeader = make([]ScoreboardV3GameHeader, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 21 {
				item := ScoreboardV3GameHeader{
					GameID:             toString(row[0]),
					GameCode:           toString(row[1]),
					GameStatus:         toString(row[2]),
					GameStatusText:     toString(row[3]),
					Period:             toInt(row[4]),
					GameClock:          toString(row[5]),
					GameTimeUTC:        toString(row[6]),
					GameEt:             toString(row[7]),
					RegulationPeriods:  toInt(row[8]),
					SeriesGameNumber:   toString(row[9]),
					SeriesText:         toString(row[10]),
					HomeTeamID:         toString(row[11]),
					HomeTeamName:       toString(row[12]),
					HomeTeamCity:       toString(row[13]),
					HomeTeamTricode:    toString(row[14]),
					HomeTeamScore:      toString(row[15]),
					VisitorTeamID:      toString(row[16]),
					VisitorTeamName:    toString(row[17]),
					VisitorTeamCity:    toString(row[18]),
					VisitorTeamTricode: toString(row[19]),
					VisitorTeamScore:   toString(row[20]),
				}
				response.GameHeader = append(response.GameHeader, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
