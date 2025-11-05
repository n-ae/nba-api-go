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
	gameId             string `json:"gameId"`
	gameCode           string `json:"gameCode"`
	gameStatus         string `json:"gameStatus"`
	gameStatusText     string `json:"gameStatusText"`
	period             int    `json:"period"`
	gameClock          string `json:"gameClock"`
	gameTimeUTC        string `json:"gameTimeUTC"`
	gameEt             string `json:"gameEt"`
	regulationPeriods  int    `json:"regulationPeriods"`
	seriesGameNumber   string `json:"seriesGameNumber"`
	seriesText         string `json:"seriesText"`
	homeTeamId         string `json:"homeTeamId"`
	homeTeamName       string `json:"homeTeamName"`
	homeTeamCity       string `json:"homeTeamCity"`
	homeTeamTricode    string `json:"homeTeamTricode"`
	homeTeamScore      string `json:"homeTeamScore"`
	visitorTeamId      string `json:"visitorTeamId"`
	visitorTeamName    string `json:"visitorTeamName"`
	visitorTeamCity    string `json:"visitorTeamCity"`
	visitorTeamTricode string `json:"visitorTeamTricode"`
	visitorTeamScore   string `json:"visitorTeamScore"`
}

// ScoreboardV3Response contains the response data from the ScoreboardV3 endpoint
type ScoreboardV3Response struct {
	GameHeader []ScoreboardV3GameHeader
}

// GetScoreboardV3 retrieves data from the scoreboardv3 endpoint
func GetScoreboardV3(ctx context.Context, client *stats.Client, req ScoreboardV3Request) (*models.Response[*ScoreboardV3Response], error) {
	params := url.Values{}
	if req.GameDate == "" {
		return nil, fmt.Errorf("GameDate is required")
	}
	params.Set("GameDate", string(req.GameDate))
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
					gameId:             toString(row[0]),
					gameCode:           toString(row[1]),
					gameStatus:         toString(row[2]),
					gameStatusText:     toString(row[3]),
					period:             toInt(row[4]),
					gameClock:          toString(row[5]),
					gameTimeUTC:        toString(row[6]),
					gameEt:             toString(row[7]),
					regulationPeriods:  toInt(row[8]),
					seriesGameNumber:   toString(row[9]),
					seriesText:         toString(row[10]),
					homeTeamId:         toString(row[11]),
					homeTeamName:       toString(row[12]),
					homeTeamCity:       toString(row[13]),
					homeTeamTricode:    toString(row[14]),
					homeTeamScore:      toString(row[15]),
					visitorTeamId:      toString(row[16]),
					visitorTeamName:    toString(row[17]),
					visitorTeamCity:    toString(row[18]),
					visitorTeamTricode: toString(row[19]),
					visitorTeamScore:   toString(row[20]),
				}
				response.GameHeader = append(response.GameHeader, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
