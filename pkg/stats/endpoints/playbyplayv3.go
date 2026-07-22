package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// PlayByPlayV3Request contains parameters for the PlayByPlayV3 endpoint
type PlayByPlayV3Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
}

// PlayByPlayV3PlayByPlay represents the PlayByPlay result set for PlayByPlayV3
type PlayByPlayV3PlayByPlay struct {
	GameID                  string  `json:"gameId"`
	ActionNumber            string  `json:"actionNumber"`
	Clock                   string  `json:"clock"`
	TimeActual              string  `json:"timeActual"`
	Period                  int     `json:"period"`
	PeriodType              int     `json:"periodType"`
	TeamID                  string  `json:"teamId"`
	TeamTricode             string  `json:"teamTricode"`
	ActionType              string  `json:"actionType"`
	SubType                 string  `json:"subType"`
	Descriptor              string  `json:"descriptor"`
	Qualifiers              string  `json:"qualifiers"`
	PersonID                string  `json:"personId"`
	PlayerName              string  `json:"playerName"`
	PlayerNameI             string  `json:"playerNameI"`
	JerseyNum               string  `json:"jerseyNum"`
	AssistPersonID          string  `json:"assistPersonId"`
	AssistPlayerNameI       string  `json:"assistPlayerNameI"`
	AssistTotal             string  `json:"assistTotal"`
	OfficialID              string  `json:"officialId"`
	Description             string  `json:"description"`
	ShotDistance            string  `json:"shotDistance"`
	ShotResult              string  `json:"shotResult"`
	IsFieldGoal             string  `json:"isFieldGoal"`
	ScoreHome               string  `json:"scoreHome"`
	ScoreAway               string  `json:"scoreAway"`
	PointsTotal             string  `json:"pointsTotal"`
	Location                string  `json:"location"`
	XLegacy                 string  `json:"xLegacy"`
	YLegacy                 string  `json:"yLegacy"`
	IsTargetScoreLastPeriod float64 `json:"isTargetScoreLastPeriod"`
	OrderNumber             string  `json:"orderNumber"`
	Edited                  string  `json:"edited"`
}

// PlayByPlayV3Response contains the response data from the PlayByPlayV3 endpoint
type PlayByPlayV3Response struct {
	PlayByPlay []PlayByPlayV3PlayByPlay
}

// GetPlayByPlayV3 retrieves data from the playbyplayv3 endpoint
func GetPlayByPlayV3(ctx context.Context, client *stats.Client, req PlayByPlayV3Request) (*models.Response[*PlayByPlayV3Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("%s is required", "GameID")
	}
	params.Set("GameID", req.GameID)
	if req.StartPeriod != nil {
		params.Set("StartPeriod", *req.StartPeriod)
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", *req.EndPeriod)
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "playbyplayv3", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayByPlayV3Response{}
	if rs, ok := findResultSet(rawResp.ResultSets, "PlayByPlay"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(PlayByPlayV3PlayByPlay{})); err != nil {
			return nil, fmt.Errorf("PlayByPlayV3: PlayByPlay result set: %w", err)
		}
		response.PlayByPlay = make([]PlayByPlayV3PlayByPlay, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 33 {
				item := PlayByPlayV3PlayByPlay{
					GameID:                  toString(row[0]),
					ActionNumber:            toString(row[1]),
					Clock:                   toString(row[2]),
					TimeActual:              toString(row[3]),
					Period:                  toInt(row[4]),
					PeriodType:              toInt(row[5]),
					TeamID:                  toString(row[6]),
					TeamTricode:             toString(row[7]),
					ActionType:              toString(row[8]),
					SubType:                 toString(row[9]),
					Descriptor:              toString(row[10]),
					Qualifiers:              toString(row[11]),
					PersonID:                toString(row[12]),
					PlayerName:              toString(row[13]),
					PlayerNameI:             toString(row[14]),
					JerseyNum:               toString(row[15]),
					AssistPersonID:          toString(row[16]),
					AssistPlayerNameI:       toString(row[17]),
					AssistTotal:             toString(row[18]),
					OfficialID:              toString(row[19]),
					Description:             toString(row[20]),
					ShotDistance:            toString(row[21]),
					ShotResult:              toString(row[22]),
					IsFieldGoal:             toString(row[23]),
					ScoreHome:               toString(row[24]),
					ScoreAway:               toString(row[25]),
					PointsTotal:             toString(row[26]),
					Location:                toString(row[27]),
					XLegacy:                 toString(row[28]),
					YLegacy:                 toString(row[29]),
					IsTargetScoreLastPeriod: toFloat(row[30]),
					OrderNumber:             toString(row[31]),
					Edited:                  toString(row[32]),
				}
				response.PlayByPlay = append(response.PlayByPlay, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
