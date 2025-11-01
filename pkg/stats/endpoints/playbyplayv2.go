package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// PlayByPlayV2Request contains parameters for the PlayByPlayV2 endpoint
type PlayByPlayV2Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
}

// PlayByPlayV2PlayByPlay represents the PlayByPlay result set for PlayByPlayV2
type PlayByPlayV2PlayByPlay struct {
	GAME_ID                   string `json:"GAME_ID"`
	EVENTNUM                  int    `json:"EVENTNUM"`
	EVENTMSGTYPE              int    `json:"EVENTMSGTYPE"`
	EVENTMSGACTIONTYPE        int    `json:"EVENTMSGACTIONTYPE"`
	PERIOD                    int    `json:"PERIOD"`
	WCTIMESTRING              string `json:"WCTIMESTRING"`
	PCTIMESTRING              string `json:"PCTIMESTRING"`
	HOMEDESCRIPTION           string `json:"HOMEDESCRIPTION"`
	NEUTRALDESCRIPTION        string `json:"NEUTRALDESCRIPTION"`
	VISITORDESCRIPTION        string `json:"VISITORDESCRIPTION"`
	SCORE                     string `json:"SCORE"`
	SCOREMARGIN               string `json:"SCOREMARGIN"`
	PERSON1TYPE               int    `json:"PERSON1TYPE"`
	PLAYER1_ID                int    `json:"PLAYER1_ID"`
	PLAYER1_NAME              string `json:"PLAYER1_NAME"`
	PLAYER1_TEAM_ID           int    `json:"PLAYER1_TEAM_ID"`
	PLAYER1_TEAM_CITY         string `json:"PLAYER1_TEAM_CITY"`
	PLAYER1_TEAM_NICKNAME     string `json:"PLAYER1_TEAM_NICKNAME"`
	PLAYER1_TEAM_ABBREVIATION string `json:"PLAYER1_TEAM_ABBREVIATION"`
	PERSON2TYPE               int    `json:"PERSON2TYPE"`
	PLAYER2_ID                int    `json:"PLAYER2_ID"`
	PLAYER2_NAME              string `json:"PLAYER2_NAME"`
	PLAYER2_TEAM_ID           int    `json:"PLAYER2_TEAM_ID"`
	PLAYER2_TEAM_CITY         string `json:"PLAYER2_TEAM_CITY"`
	PLAYER2_TEAM_NICKNAME     string `json:"PLAYER2_TEAM_NICKNAME"`
	PLAYER2_TEAM_ABBREVIATION string `json:"PLAYER2_TEAM_ABBREVIATION"`
	PERSON3TYPE               int    `json:"PERSON3TYPE"`
	PLAYER3_ID                int    `json:"PLAYER3_ID"`
	PLAYER3_NAME              string `json:"PLAYER3_NAME"`
	PLAYER3_TEAM_ID           int    `json:"PLAYER3_TEAM_ID"`
	PLAYER3_TEAM_CITY         string `json:"PLAYER3_TEAM_CITY"`
	PLAYER3_TEAM_NICKNAME     string `json:"PLAYER3_TEAM_NICKNAME"`
	PLAYER3_TEAM_ABBREVIATION string `json:"PLAYER3_TEAM_ABBREVIATION"`
	VIDEO_AVAILABLE_FLAG      int    `json:"VIDEO_AVAILABLE_FLAG"`
}

// PlayByPlayV2AvailableVideo represents the AvailableVideo result set for PlayByPlayV2
type PlayByPlayV2AvailableVideo struct {
	GAME_ID              string `json:"GAME_ID"`
	VIDEO_AVAILABLE_FLAG int    `json:"VIDEO_AVAILABLE_FLAG"`
}

// PlayByPlayV2Response contains the response data from the PlayByPlayV2 endpoint
type PlayByPlayV2Response struct {
	PlayByPlay     []PlayByPlayV2PlayByPlay
	AvailableVideo []PlayByPlayV2AvailableVideo
}

// GetPlayByPlayV2 retrieves data from the playbyplayv2 endpoint
func GetPlayByPlayV2(ctx context.Context, client *stats.Client, req PlayByPlayV2Request) (*models.Response[*PlayByPlayV2Response], error) {
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
	if err := client.GetJSON(ctx, "/playbyplayv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayByPlayV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayByPlay = make([]PlayByPlayV2PlayByPlay, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 34 {
				item := PlayByPlayV2PlayByPlay{
					GAME_ID:                   toString(row[0]),
					EVENTNUM:                  toInt(row[1]),
					EVENTMSGTYPE:              toInt(row[2]),
					EVENTMSGACTIONTYPE:        toInt(row[3]),
					PERIOD:                    toInt(row[4]),
					WCTIMESTRING:              toString(row[5]),
					PCTIMESTRING:              toString(row[6]),
					HOMEDESCRIPTION:           toString(row[7]),
					NEUTRALDESCRIPTION:        toString(row[8]),
					VISITORDESCRIPTION:        toString(row[9]),
					SCORE:                     toString(row[10]),
					SCOREMARGIN:               toString(row[11]),
					PERSON1TYPE:               toInt(row[12]),
					PLAYER1_ID:                toInt(row[13]),
					PLAYER1_NAME:              toString(row[14]),
					PLAYER1_TEAM_ID:           toInt(row[15]),
					PLAYER1_TEAM_CITY:         toString(row[16]),
					PLAYER1_TEAM_NICKNAME:     toString(row[17]),
					PLAYER1_TEAM_ABBREVIATION: toString(row[18]),
					PERSON2TYPE:               toInt(row[19]),
					PLAYER2_ID:                toInt(row[20]),
					PLAYER2_NAME:              toString(row[21]),
					PLAYER2_TEAM_ID:           toInt(row[22]),
					PLAYER2_TEAM_CITY:         toString(row[23]),
					PLAYER2_TEAM_NICKNAME:     toString(row[24]),
					PLAYER2_TEAM_ABBREVIATION: toString(row[25]),
					PERSON3TYPE:               toInt(row[26]),
					PLAYER3_ID:                toInt(row[27]),
					PLAYER3_NAME:              toString(row[28]),
					PLAYER3_TEAM_ID:           toInt(row[29]),
					PLAYER3_TEAM_CITY:         toString(row[30]),
					PLAYER3_TEAM_NICKNAME:     toString(row[31]),
					PLAYER3_TEAM_ABBREVIATION: toString(row[32]),
					VIDEO_AVAILABLE_FLAG:      toInt(row[33]),
				}
				response.PlayByPlay = append(response.PlayByPlay, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.AvailableVideo = make([]PlayByPlayV2AvailableVideo, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 2 {
				item := PlayByPlayV2AvailableVideo{
					GAME_ID:              toString(row[0]),
					VIDEO_AVAILABLE_FLAG: toInt(row[1]),
				}
				response.AvailableVideo = append(response.AvailableVideo, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
