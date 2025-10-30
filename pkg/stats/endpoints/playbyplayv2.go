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
	GameID string
	StartPeriod *string
	EndPeriod *string
}


// PlayByPlayV2PlayByPlay represents the PlayByPlay result set for PlayByPlayV2
type PlayByPlayV2PlayByPlay struct {
	GAME_ID interface{}
	EVENTNUM interface{}
	EVENTMSGTYPE interface{}
	EVENTMSGACTIONTYPE interface{}
	PERIOD interface{}
	WCTIMESTRING interface{}
	PCTIMESTRING interface{}
	HOMEDESCRIPTION interface{}
	NEUTRALDESCRIPTION interface{}
	VISITORDESCRIPTION interface{}
	SCORE interface{}
	SCOREMARGIN interface{}
	PERSON1TYPE interface{}
	PLAYER1_ID interface{}
	PLAYER1_NAME interface{}
	PLAYER1_TEAM_ID interface{}
	PLAYER1_TEAM_CITY interface{}
	PLAYER1_TEAM_NICKNAME interface{}
	PLAYER1_TEAM_ABBREVIATION interface{}
	PERSON2TYPE interface{}
	PLAYER2_ID interface{}
	PLAYER2_NAME interface{}
	PLAYER2_TEAM_ID interface{}
	PLAYER2_TEAM_CITY interface{}
	PLAYER2_TEAM_NICKNAME interface{}
	PLAYER2_TEAM_ABBREVIATION interface{}
	PERSON3TYPE interface{}
	PLAYER3_ID interface{}
	PLAYER3_NAME interface{}
	PLAYER3_TEAM_ID interface{}
	PLAYER3_TEAM_CITY interface{}
	PLAYER3_TEAM_NICKNAME interface{}
	PLAYER3_TEAM_ABBREVIATION interface{}
	VIDEO_AVAILABLE_FLAG interface{}
}

// PlayByPlayV2AvailableVideo represents the AvailableVideo result set for PlayByPlayV2
type PlayByPlayV2AvailableVideo struct {
	GAME_ID interface{}
	VIDEO_AVAILABLE_FLAG interface{}
}


// PlayByPlayV2Response contains the response data from the PlayByPlayV2 endpoint
type PlayByPlayV2Response struct {
	PlayByPlay []PlayByPlayV2PlayByPlay
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
		response.PlayByPlay = make([]PlayByPlayV2PlayByPlay, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 34 {
				response.PlayByPlay[i] = PlayByPlayV2PlayByPlay{
					GAME_ID: row[0],
					EVENTNUM: row[1],
					EVENTMSGTYPE: row[2],
					EVENTMSGACTIONTYPE: row[3],
					PERIOD: row[4],
					WCTIMESTRING: row[5],
					PCTIMESTRING: row[6],
					HOMEDESCRIPTION: row[7],
					NEUTRALDESCRIPTION: row[8],
					VISITORDESCRIPTION: row[9],
					SCORE: row[10],
					SCOREMARGIN: row[11],
					PERSON1TYPE: row[12],
					PLAYER1_ID: row[13],
					PLAYER1_NAME: row[14],
					PLAYER1_TEAM_ID: row[15],
					PLAYER1_TEAM_CITY: row[16],
					PLAYER1_TEAM_NICKNAME: row[17],
					PLAYER1_TEAM_ABBREVIATION: row[18],
					PERSON2TYPE: row[19],
					PLAYER2_ID: row[20],
					PLAYER2_NAME: row[21],
					PLAYER2_TEAM_ID: row[22],
					PLAYER2_TEAM_CITY: row[23],
					PLAYER2_TEAM_NICKNAME: row[24],
					PLAYER2_TEAM_ABBREVIATION: row[25],
					PERSON3TYPE: row[26],
					PLAYER3_ID: row[27],
					PLAYER3_NAME: row[28],
					PLAYER3_TEAM_ID: row[29],
					PLAYER3_TEAM_CITY: row[30],
					PLAYER3_TEAM_NICKNAME: row[31],
					PLAYER3_TEAM_ABBREVIATION: row[32],
					VIDEO_AVAILABLE_FLAG: row[33],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.AvailableVideo = make([]PlayByPlayV2AvailableVideo, len(rawResp.ResultSets[1].RowSet))
		for i, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 2 {
				response.AvailableVideo[i] = PlayByPlayV2AvailableVideo{
					GAME_ID: row[0],
					VIDEO_AVAILABLE_FLAG: row[1],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
