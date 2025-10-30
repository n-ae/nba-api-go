package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// BoxScoreSummaryV2Request contains parameters for the BoxScoreSummaryV2 endpoint
type BoxScoreSummaryV2Request struct {
	GameID string
}


// BoxScoreSummaryV2GameSummary represents the GameSummary result set for BoxScoreSummaryV2
type BoxScoreSummaryV2GameSummary struct {
	GAME_DATE_EST interface{}
	GAME_SEQUENCE interface{}
	GAME_ID interface{}
	GAME_STATUS_ID interface{}
	GAME_STATUS_TEXT interface{}
	GAMECODE interface{}
	HOME_TEAM_ID interface{}
	VISITOR_TEAM_ID interface{}
	SEASON interface{}
	LIVE_PERIOD interface{}
	LIVE_PC_TIME interface{}
	NATL_TV_BROADCASTER_ABBREVIATION interface{}
	LIVE_PERIOD_TIME_BCAST interface{}
	WH_STATUS interface{}
}

// BoxScoreSummaryV2OtherStats represents the OtherStats result set for BoxScoreSummaryV2
type BoxScoreSummaryV2OtherStats struct {
	LEAGUE_ID interface{}
	TEAM_ID interface{}
	TEAM_ABBREVIATION interface{}
	TEAM_CITY interface{}
	PTS_PAINT interface{}
	PTS_2ND_CHANCE interface{}
	PTS_FB interface{}
	LARGEST_LEAD interface{}
	LEAD_CHANGES interface{}
	TIMES_TIED interface{}
	TEAM_TURNOVERS interface{}
	TOTAL_TURNOVERS interface{}
	TEAM_REBOUNDS interface{}
	PTS_OFF_TO interface{}
}

// BoxScoreSummaryV2Officials represents the Officials result set for BoxScoreSummaryV2
type BoxScoreSummaryV2Officials struct {
	OFFICIAL_ID interface{}
	FIRST_NAME interface{}
	LAST_NAME interface{}
	JERSEY_NUM interface{}
}

// BoxScoreSummaryV2InactivePlayers represents the InactivePlayers result set for BoxScoreSummaryV2
type BoxScoreSummaryV2InactivePlayers struct {
	PLAYER_ID interface{}
	FIRST_NAME interface{}
	LAST_NAME interface{}
	JERSEY_NUM interface{}
	TEAM_ID interface{}
	TEAM_CITY interface{}
	TEAM_NAME interface{}
	TEAM_ABBREVIATION interface{}
}

// BoxScoreSummaryV2GameInfo represents the GameInfo result set for BoxScoreSummaryV2
type BoxScoreSummaryV2GameInfo struct {
	GAME_DATE interface{}
	ATTENDANCE interface{}
	GAME_TIME interface{}
}

// BoxScoreSummaryV2LineScore represents the LineScore result set for BoxScoreSummaryV2
type BoxScoreSummaryV2LineScore struct {
	GAME_DATE_EST interface{}
	GAME_SEQUENCE interface{}
	GAME_ID interface{}
	TEAM_ID interface{}
	TEAM_ABBREVIATION interface{}
	TEAM_CITY_NAME interface{}
	TEAM_WINS_LOSSES interface{}
	PTS_QTR1 interface{}
	PTS_QTR2 interface{}
	PTS_QTR3 interface{}
	PTS_QTR4 interface{}
	PTS_OT1 interface{}
	PTS_OT2 interface{}
	PTS_OT3 interface{}
	PTS_OT4 interface{}
	PTS_OT5 interface{}
	PTS_OT6 interface{}
	PTS_OT7 interface{}
	PTS_OT8 interface{}
	PTS_OT9 interface{}
	PTS_OT10 interface{}
	PTS interface{}
	FG_PCT interface{}
	FT_PCT interface{}
	FG3_PCT interface{}
	AST interface{}
	REB interface{}
	TOV interface{}
}

// BoxScoreSummaryV2LastMeeting represents the LastMeeting result set for BoxScoreSummaryV2
type BoxScoreSummaryV2LastMeeting struct {
	GAME_ID interface{}
	GAME_DATE_EST interface{}
	GAME_DATE_TIME_EST interface{}
	HOME_TEAM_ID interface{}
	HOME_TEAM_CITY interface{}
	HOME_TEAM_NAME interface{}
	HOME_TEAM_ABBREVIATION interface{}
	HOME_TEAM_POINTS interface{}
	VISITOR_TEAM_ID interface{}
	VISITOR_TEAM_CITY interface{}
	VISITOR_TEAM_NAME interface{}
	VISITOR_TEAM_ABBREVIATION interface{}
	VISITOR_TEAM_POINTS interface{}
}

// BoxScoreSummaryV2SeasonSeries represents the SeasonSeries result set for BoxScoreSummaryV2
type BoxScoreSummaryV2SeasonSeries struct {
	GAME_ID interface{}
	HOME_TEAM_ID interface{}
	VISITOR_TEAM_ID interface{}
	GAME_DATE_EST interface{}
	HOME_TEAM_WINS interface{}
	HOME_TEAM_LOSSES interface{}
	SERIES_LEADER interface{}
}

// BoxScoreSummaryV2AvailableVideo represents the AvailableVideo result set for BoxScoreSummaryV2
type BoxScoreSummaryV2AvailableVideo struct {
	GAME_ID interface{}
	VIDEO_AVAILABLE_FLAG interface{}
}


// BoxScoreSummaryV2Response contains the response data from the BoxScoreSummaryV2 endpoint
type BoxScoreSummaryV2Response struct {
	GameSummary []BoxScoreSummaryV2GameSummary
	OtherStats []BoxScoreSummaryV2OtherStats
	Officials []BoxScoreSummaryV2Officials
	InactivePlayers []BoxScoreSummaryV2InactivePlayers
	GameInfo []BoxScoreSummaryV2GameInfo
	LineScore []BoxScoreSummaryV2LineScore
	LastMeeting []BoxScoreSummaryV2LastMeeting
	SeasonSeries []BoxScoreSummaryV2SeasonSeries
	AvailableVideo []BoxScoreSummaryV2AvailableVideo
}

// GetBoxScoreSummaryV2 retrieves data from the boxscoresummaryv2 endpoint
func GetBoxScoreSummaryV2(ctx context.Context, client *stats.Client, req BoxScoreSummaryV2Request) (*models.Response[*BoxScoreSummaryV2Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/boxscoresummaryv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreSummaryV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.GameSummary = make([]BoxScoreSummaryV2GameSummary, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				response.GameSummary[i] = BoxScoreSummaryV2GameSummary{
					GAME_DATE_EST: row[0],
					GAME_SEQUENCE: row[1],
					GAME_ID: row[2],
					GAME_STATUS_ID: row[3],
					GAME_STATUS_TEXT: row[4],
					GAMECODE: row[5],
					HOME_TEAM_ID: row[6],
					VISITOR_TEAM_ID: row[7],
					SEASON: row[8],
					LIVE_PERIOD: row[9],
					LIVE_PC_TIME: row[10],
					NATL_TV_BROADCASTER_ABBREVIATION: row[11],
					LIVE_PERIOD_TIME_BCAST: row[12],
					WH_STATUS: row[13],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.OtherStats = make([]BoxScoreSummaryV2OtherStats, len(rawResp.ResultSets[1].RowSet))
		for i, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 14 {
				response.OtherStats[i] = BoxScoreSummaryV2OtherStats{
					LEAGUE_ID: row[0],
					TEAM_ID: row[1],
					TEAM_ABBREVIATION: row[2],
					TEAM_CITY: row[3],
					PTS_PAINT: row[4],
					PTS_2ND_CHANCE: row[5],
					PTS_FB: row[6],
					LARGEST_LEAD: row[7],
					LEAD_CHANGES: row[8],
					TIMES_TIED: row[9],
					TEAM_TURNOVERS: row[10],
					TOTAL_TURNOVERS: row[11],
					TEAM_REBOUNDS: row[12],
					PTS_OFF_TO: row[13],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.Officials = make([]BoxScoreSummaryV2Officials, len(rawResp.ResultSets[2].RowSet))
		for i, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 4 {
				response.Officials[i] = BoxScoreSummaryV2Officials{
					OFFICIAL_ID: row[0],
					FIRST_NAME: row[1],
					LAST_NAME: row[2],
					JERSEY_NUM: row[3],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.InactivePlayers = make([]BoxScoreSummaryV2InactivePlayers, len(rawResp.ResultSets[3].RowSet))
		for i, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 8 {
				response.InactivePlayers[i] = BoxScoreSummaryV2InactivePlayers{
					PLAYER_ID: row[0],
					FIRST_NAME: row[1],
					LAST_NAME: row[2],
					JERSEY_NUM: row[3],
					TEAM_ID: row[4],
					TEAM_CITY: row[5],
					TEAM_NAME: row[6],
					TEAM_ABBREVIATION: row[7],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.GameInfo = make([]BoxScoreSummaryV2GameInfo, len(rawResp.ResultSets[4].RowSet))
		for i, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 3 {
				response.GameInfo[i] = BoxScoreSummaryV2GameInfo{
					GAME_DATE: row[0],
					ATTENDANCE: row[1],
					GAME_TIME: row[2],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 5 {
		response.LineScore = make([]BoxScoreSummaryV2LineScore, len(rawResp.ResultSets[5].RowSet))
		for i, row := range rawResp.ResultSets[5].RowSet {
			if len(row) >= 28 {
				response.LineScore[i] = BoxScoreSummaryV2LineScore{
					GAME_DATE_EST: row[0],
					GAME_SEQUENCE: row[1],
					GAME_ID: row[2],
					TEAM_ID: row[3],
					TEAM_ABBREVIATION: row[4],
					TEAM_CITY_NAME: row[5],
					TEAM_WINS_LOSSES: row[6],
					PTS_QTR1: row[7],
					PTS_QTR2: row[8],
					PTS_QTR3: row[9],
					PTS_QTR4: row[10],
					PTS_OT1: row[11],
					PTS_OT2: row[12],
					PTS_OT3: row[13],
					PTS_OT4: row[14],
					PTS_OT5: row[15],
					PTS_OT6: row[16],
					PTS_OT7: row[17],
					PTS_OT8: row[18],
					PTS_OT9: row[19],
					PTS_OT10: row[20],
					PTS: row[21],
					FG_PCT: row[22],
					FT_PCT: row[23],
					FG3_PCT: row[24],
					AST: row[25],
					REB: row[26],
					TOV: row[27],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 6 {
		response.LastMeeting = make([]BoxScoreSummaryV2LastMeeting, len(rawResp.ResultSets[6].RowSet))
		for i, row := range rawResp.ResultSets[6].RowSet {
			if len(row) >= 13 {
				response.LastMeeting[i] = BoxScoreSummaryV2LastMeeting{
					GAME_ID: row[0],
					GAME_DATE_EST: row[1],
					GAME_DATE_TIME_EST: row[2],
					HOME_TEAM_ID: row[3],
					HOME_TEAM_CITY: row[4],
					HOME_TEAM_NAME: row[5],
					HOME_TEAM_ABBREVIATION: row[6],
					HOME_TEAM_POINTS: row[7],
					VISITOR_TEAM_ID: row[8],
					VISITOR_TEAM_CITY: row[9],
					VISITOR_TEAM_NAME: row[10],
					VISITOR_TEAM_ABBREVIATION: row[11],
					VISITOR_TEAM_POINTS: row[12],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 7 {
		response.SeasonSeries = make([]BoxScoreSummaryV2SeasonSeries, len(rawResp.ResultSets[7].RowSet))
		for i, row := range rawResp.ResultSets[7].RowSet {
			if len(row) >= 7 {
				response.SeasonSeries[i] = BoxScoreSummaryV2SeasonSeries{
					GAME_ID: row[0],
					HOME_TEAM_ID: row[1],
					VISITOR_TEAM_ID: row[2],
					GAME_DATE_EST: row[3],
					HOME_TEAM_WINS: row[4],
					HOME_TEAM_LOSSES: row[5],
					SERIES_LEADER: row[6],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 8 {
		response.AvailableVideo = make([]BoxScoreSummaryV2AvailableVideo, len(rawResp.ResultSets[8].RowSet))
		for i, row := range rawResp.ResultSets[8].RowSet {
			if len(row) >= 2 {
				response.AvailableVideo[i] = BoxScoreSummaryV2AvailableVideo{
					GAME_ID: row[0],
					VIDEO_AVAILABLE_FLAG: row[1],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
