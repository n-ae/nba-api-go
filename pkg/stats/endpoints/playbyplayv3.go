package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// PlayByPlayV3Request contains parameters for the PlayByPlayV3 endpoint
type PlayByPlayV3Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
}

// PlayByPlayV3PlayByPlay represents the PlayByPlay result set for PlayByPlayV3
type PlayByPlayV3PlayByPlay struct {
	gameId                  string  `json:"gameId"`
	actionNumber            string  `json:"actionNumber"`
	clock                   string  `json:"clock"`
	timeActual              string  `json:"timeActual"`
	period                  int     `json:"period"`
	periodType              int     `json:"periodType"`
	teamId                  string  `json:"teamId"`
	teamTricode             string  `json:"teamTricode"`
	actionType              string  `json:"actionType"`
	subType                 string  `json:"subType"`
	descriptor              string  `json:"descriptor"`
	qualifiers              string  `json:"qualifiers"`
	personId                string  `json:"personId"`
	playerName              string  `json:"playerName"`
	playerNameI             string  `json:"playerNameI"`
	jerseyNum               string  `json:"jerseyNum"`
	assistPersonId          string  `json:"assistPersonId"`
	assistPlayerNameI       string  `json:"assistPlayerNameI"`
	assistTotal             string  `json:"assistTotal"`
	officialId              string  `json:"officialId"`
	description             string  `json:"description"`
	shotDistance            string  `json:"shotDistance"`
	shotResult              string  `json:"shotResult"`
	isFieldGoal             string  `json:"isFieldGoal"`
	scoreHome               string  `json:"scoreHome"`
	scoreAway               string  `json:"scoreAway"`
	pointsTotal             string  `json:"pointsTotal"`
	location                string  `json:"location"`
	xLegacy                 string  `json:"xLegacy"`
	yLegacy                 string  `json:"yLegacy"`
	isTargetScoreLastPeriod float64 `json:"isTargetScoreLastPeriod"`
	orderNumber             string  `json:"orderNumber"`
	edited                  string  `json:"edited"`
}

// PlayByPlayV3Response contains the response data from the PlayByPlayV3 endpoint
type PlayByPlayV3Response struct {
	PlayByPlay []PlayByPlayV3PlayByPlay
}

// GetPlayByPlayV3 retrieves data from the playbyplayv3 endpoint
func GetPlayByPlayV3(ctx context.Context, client *stats.Client, req PlayByPlayV3Request) (*models.Response[*PlayByPlayV3Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.StartPeriod != nil {
		params.Set("StartPeriod", string(*req.StartPeriod))
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", string(*req.EndPeriod))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playbyplayv3", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayByPlayV3Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayByPlay = make([]PlayByPlayV3PlayByPlay, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 33 {
				item := PlayByPlayV3PlayByPlay{
					gameId:                  toString(row[0]),
					actionNumber:            toString(row[1]),
					clock:                   toString(row[2]),
					timeActual:              toString(row[3]),
					period:                  toInt(row[4]),
					periodType:              toInt(row[5]),
					teamId:                  toString(row[6]),
					teamTricode:             toString(row[7]),
					actionType:              toString(row[8]),
					subType:                 toString(row[9]),
					descriptor:              toString(row[10]),
					qualifiers:              toString(row[11]),
					personId:                toString(row[12]),
					playerName:              toString(row[13]),
					playerNameI:             toString(row[14]),
					jerseyNum:               toString(row[15]),
					assistPersonId:          toString(row[16]),
					assistPlayerNameI:       toString(row[17]),
					assistTotal:             toString(row[18]),
					officialId:              toString(row[19]),
					description:             toString(row[20]),
					shotDistance:            toString(row[21]),
					shotResult:              toString(row[22]),
					isFieldGoal:             toString(row[23]),
					scoreHome:               toString(row[24]),
					scoreAway:               toString(row[25]),
					pointsTotal:             toString(row[26]),
					location:                toString(row[27]),
					xLegacy:                 toString(row[28]),
					yLegacy:                 toString(row[29]),
					isTargetScoreLastPeriod: toFloat(row[30]),
					orderNumber:             toString(row[31]),
					edited:                  toString(row[32]),
				}
				response.PlayByPlay = append(response.PlayByPlay, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
