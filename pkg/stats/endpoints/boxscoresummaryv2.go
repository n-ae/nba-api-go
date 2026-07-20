package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v2/pkg/models"
	"github.com/n-ae/nba-api-go/v2/pkg/stats"
)

// BoxScoreSummaryV2Request contains parameters for the BoxScoreSummaryV2 endpoint
type BoxScoreSummaryV2Request struct {
	GameID string
}

// BoxScoreSummaryV2GameSummary represents the GameSummary result set for BoxScoreSummaryV2
type BoxScoreSummaryV2GameSummary struct {
	GAME_DATE_EST                    string  `json:"GAME_DATE_EST"`
	GAME_SEQUENCE                    int     `json:"GAME_SEQUENCE"`
	GAME_ID                          string  `json:"GAME_ID"`
	GAME_STATUS_ID                   string  `json:"GAME_STATUS_ID"`
	GAME_STATUS_TEXT                 string  `json:"GAME_STATUS_TEXT"`
	GAMECODE                         string  `json:"GAMECODE"`
	HOME_TEAM_ID                     int     `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID                  int     `json:"VISITOR_TEAM_ID"`
	SEASON                           string  `json:"SEASON"`
	LIVE_PERIOD                      int     `json:"LIVE_PERIOD"`
	LIVE_PC_TIME                     string  `json:"LIVE_PC_TIME"`
	NATL_TV_BROADCASTER_ABBREVIATION string  `json:"NATL_TV_BROADCASTER_ABBREVIATION"`
	LIVE_PERIOD_TIME_BCAST           float64 `json:"LIVE_PERIOD_TIME_BCAST"`
	WH_STATUS                        string  `json:"WH_STATUS"`
}

// BoxScoreSummaryV2OtherStats represents the OtherStats result set for BoxScoreSummaryV2
type BoxScoreSummaryV2OtherStats struct {
	LEAGUE_ID         string  `json:"LEAGUE_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	PTS_PAINT         float64 `json:"PTS_PAINT"`
	PTS_2ND_CHANCE    float64 `json:"PTS_2ND_CHANCE"`
	PTS_FB            float64 `json:"PTS_FB"`
	LARGEST_LEAD      string  `json:"LARGEST_LEAD"`
	LEAD_CHANGES      string  `json:"LEAD_CHANGES"`
	TIMES_TIED        string  `json:"TIMES_TIED"`
	TEAM_TURNOVERS    string  `json:"TEAM_TURNOVERS"`
	TOTAL_TURNOVERS   string  `json:"TOTAL_TURNOVERS"`
	TEAM_REBOUNDS     float64 `json:"TEAM_REBOUNDS"`
	PTS_OFF_TO        float64 `json:"PTS_OFF_TO"`
}

// BoxScoreSummaryV2Officials represents the Officials result set for BoxScoreSummaryV2
type BoxScoreSummaryV2Officials struct {
	OFFICIAL_ID string `json:"OFFICIAL_ID"`
	FIRST_NAME  string `json:"FIRST_NAME"`
	LAST_NAME   string `json:"LAST_NAME"`
	JERSEY_NUM  string `json:"JERSEY_NUM"`
}

// BoxScoreSummaryV2InactivePlayers represents the InactivePlayers result set for BoxScoreSummaryV2
type BoxScoreSummaryV2InactivePlayers struct {
	PLAYER_ID         int    `json:"PLAYER_ID"`
	FIRST_NAME        string `json:"FIRST_NAME"`
	LAST_NAME         string `json:"LAST_NAME"`
	JERSEY_NUM        string `json:"JERSEY_NUM"`
	TEAM_ID           int    `json:"TEAM_ID"`
	TEAM_CITY         string `json:"TEAM_CITY"`
	TEAM_NAME         string `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
}

// BoxScoreSummaryV2GameInfo represents the GameInfo result set for BoxScoreSummaryV2
type BoxScoreSummaryV2GameInfo struct {
	GAME_DATE  string `json:"GAME_DATE"`
	ATTENDANCE string `json:"ATTENDANCE"`
	GAME_TIME  string `json:"GAME_TIME"`
}

// BoxScoreSummaryV2LineScore represents the LineScore result set for BoxScoreSummaryV2
type BoxScoreSummaryV2LineScore struct {
	GAME_DATE_EST     string  `json:"GAME_DATE_EST"`
	GAME_SEQUENCE     int     `json:"GAME_SEQUENCE"`
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY_NAME    string  `json:"TEAM_CITY_NAME"`
	TEAM_WINS_LOSSES  string  `json:"TEAM_WINS_LOSSES"`
	PTS_QTR1          float64 `json:"PTS_QTR1"`
	PTS_QTR2          float64 `json:"PTS_QTR2"`
	PTS_QTR3          float64 `json:"PTS_QTR3"`
	PTS_QTR4          float64 `json:"PTS_QTR4"`
	PTS_OT1           float64 `json:"PTS_OT1"`
	PTS_OT2           float64 `json:"PTS_OT2"`
	PTS_OT3           float64 `json:"PTS_OT3"`
	PTS_OT4           float64 `json:"PTS_OT4"`
	PTS_OT5           float64 `json:"PTS_OT5"`
	PTS_OT6           float64 `json:"PTS_OT6"`
	PTS_OT7           float64 `json:"PTS_OT7"`
	PTS_OT8           float64 `json:"PTS_OT8"`
	PTS_OT9           float64 `json:"PTS_OT9"`
	PTS_OT10          float64 `json:"PTS_OT10"`
	PTS               float64 `json:"PTS"`
	FG_PCT            float64 `json:"FG_PCT"`
	FT_PCT            float64 `json:"FT_PCT"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	AST               float64 `json:"AST"`
	REB               float64 `json:"REB"`
	TOV               float64 `json:"TOV"`
}

// BoxScoreSummaryV2LastMeeting represents the LastMeeting result set for BoxScoreSummaryV2
type BoxScoreSummaryV2LastMeeting struct {
	GAME_ID                   string `json:"GAME_ID"`
	GAME_DATE_EST             string `json:"GAME_DATE_EST"`
	GAME_DATE_TIME_EST        string `json:"GAME_DATE_TIME_EST"`
	HOME_TEAM_ID              int    `json:"HOME_TEAM_ID"`
	HOME_TEAM_CITY            string `json:"HOME_TEAM_CITY"`
	HOME_TEAM_NAME            string `json:"HOME_TEAM_NAME"`
	HOME_TEAM_ABBREVIATION    string `json:"HOME_TEAM_ABBREVIATION"`
	HOME_TEAM_POINTS          string `json:"HOME_TEAM_POINTS"`
	VISITOR_TEAM_ID           int    `json:"VISITOR_TEAM_ID"`
	VISITOR_TEAM_CITY         string `json:"VISITOR_TEAM_CITY"`
	VISITOR_TEAM_NAME         string `json:"VISITOR_TEAM_NAME"`
	VISITOR_TEAM_ABBREVIATION string `json:"VISITOR_TEAM_ABBREVIATION"`
	VISITOR_TEAM_POINTS       string `json:"VISITOR_TEAM_POINTS"`
}

// BoxScoreSummaryV2SeasonSeries represents the SeasonSeries result set for BoxScoreSummaryV2
type BoxScoreSummaryV2SeasonSeries struct {
	GAME_ID          string `json:"GAME_ID"`
	HOME_TEAM_ID     int    `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID  int    `json:"VISITOR_TEAM_ID"`
	GAME_DATE_EST    string `json:"GAME_DATE_EST"`
	HOME_TEAM_WINS   string `json:"HOME_TEAM_WINS"`
	HOME_TEAM_LOSSES string `json:"HOME_TEAM_LOSSES"`
	SERIES_LEADER    string `json:"SERIES_LEADER"`
}

// BoxScoreSummaryV2AvailableVideo represents the AvailableVideo result set for BoxScoreSummaryV2
type BoxScoreSummaryV2AvailableVideo struct {
	GAME_ID              string `json:"GAME_ID"`
	VIDEO_AVAILABLE_FLAG int    `json:"VIDEO_AVAILABLE_FLAG"`
}

// BoxScoreSummaryV2Response contains the response data from the BoxScoreSummaryV2 endpoint
type BoxScoreSummaryV2Response struct {
	GameSummary     []BoxScoreSummaryV2GameSummary
	OtherStats      []BoxScoreSummaryV2OtherStats
	Officials       []BoxScoreSummaryV2Officials
	InactivePlayers []BoxScoreSummaryV2InactivePlayers
	GameInfo        []BoxScoreSummaryV2GameInfo
	LineScore       []BoxScoreSummaryV2LineScore
	LastMeeting     []BoxScoreSummaryV2LastMeeting
	SeasonSeries    []BoxScoreSummaryV2SeasonSeries
	AvailableVideo  []BoxScoreSummaryV2AvailableVideo
}

// GetBoxScoreSummaryV2 retrieves data from the boxscoresummaryv2 endpoint
func GetBoxScoreSummaryV2(ctx context.Context, client *stats.Client, req BoxScoreSummaryV2Request) (*models.Response[*BoxScoreSummaryV2Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("%s is required", "GameID")
	}
	params.Set("GameID", req.GameID)

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "boxscoresummaryv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreSummaryV2Response{}
	if rs, ok := findResultSet(rawResp.ResultSets, "GameSummary"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2GameSummary{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: GameSummary result set: %w", err)
		}
		response.GameSummary = make([]BoxScoreSummaryV2GameSummary, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 14 {
				item := BoxScoreSummaryV2GameSummary{
					GAME_DATE_EST:                    toString(row[0]),
					GAME_SEQUENCE:                    toInt(row[1]),
					GAME_ID:                          toString(row[2]),
					GAME_STATUS_ID:                   toString(row[3]),
					GAME_STATUS_TEXT:                 toString(row[4]),
					GAMECODE:                         toString(row[5]),
					HOME_TEAM_ID:                     toInt(row[6]),
					VISITOR_TEAM_ID:                  toInt(row[7]),
					SEASON:                           toString(row[8]),
					LIVE_PERIOD:                      toInt(row[9]),
					LIVE_PC_TIME:                     toString(row[10]),
					NATL_TV_BROADCASTER_ABBREVIATION: toString(row[11]),
					LIVE_PERIOD_TIME_BCAST:           toFloat(row[12]),
					WH_STATUS:                        toString(row[13]),
				}
				response.GameSummary = append(response.GameSummary, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "OtherStats"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2OtherStats{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: OtherStats result set: %w", err)
		}
		response.OtherStats = make([]BoxScoreSummaryV2OtherStats, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 14 {
				item := BoxScoreSummaryV2OtherStats{
					LEAGUE_ID:         toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_CITY:         toString(row[3]),
					PTS_PAINT:         toFloat(row[4]),
					PTS_2ND_CHANCE:    toFloat(row[5]),
					PTS_FB:            toFloat(row[6]),
					LARGEST_LEAD:      toString(row[7]),
					LEAD_CHANGES:      toString(row[8]),
					TIMES_TIED:        toString(row[9]),
					TEAM_TURNOVERS:    toString(row[10]),
					TOTAL_TURNOVERS:   toString(row[11]),
					TEAM_REBOUNDS:     toFloat(row[12]),
					PTS_OFF_TO:        toFloat(row[13]),
				}
				response.OtherStats = append(response.OtherStats, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "Officials"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2Officials{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: Officials result set: %w", err)
		}
		response.Officials = make([]BoxScoreSummaryV2Officials, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 4 {
				item := BoxScoreSummaryV2Officials{
					OFFICIAL_ID: toString(row[0]),
					FIRST_NAME:  toString(row[1]),
					LAST_NAME:   toString(row[2]),
					JERSEY_NUM:  toString(row[3]),
				}
				response.Officials = append(response.Officials, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "InactivePlayers"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2InactivePlayers{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: InactivePlayers result set: %w", err)
		}
		response.InactivePlayers = make([]BoxScoreSummaryV2InactivePlayers, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 8 {
				item := BoxScoreSummaryV2InactivePlayers{
					PLAYER_ID:         toInt(row[0]),
					FIRST_NAME:        toString(row[1]),
					LAST_NAME:         toString(row[2]),
					JERSEY_NUM:        toString(row[3]),
					TEAM_ID:           toInt(row[4]),
					TEAM_CITY:         toString(row[5]),
					TEAM_NAME:         toString(row[6]),
					TEAM_ABBREVIATION: toString(row[7]),
				}
				response.InactivePlayers = append(response.InactivePlayers, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "GameInfo"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2GameInfo{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: GameInfo result set: %w", err)
		}
		response.GameInfo = make([]BoxScoreSummaryV2GameInfo, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 3 {
				item := BoxScoreSummaryV2GameInfo{
					GAME_DATE:  toString(row[0]),
					ATTENDANCE: toString(row[1]),
					GAME_TIME:  toString(row[2]),
				}
				response.GameInfo = append(response.GameInfo, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "LineScore"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2LineScore{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: LineScore result set: %w", err)
		}
		response.LineScore = make([]BoxScoreSummaryV2LineScore, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 28 {
				item := BoxScoreSummaryV2LineScore{
					GAME_DATE_EST:     toString(row[0]),
					GAME_SEQUENCE:     toInt(row[1]),
					GAME_ID:           toString(row[2]),
					TEAM_ID:           toInt(row[3]),
					TEAM_ABBREVIATION: toString(row[4]),
					TEAM_CITY_NAME:    toString(row[5]),
					TEAM_WINS_LOSSES:  toString(row[6]),
					PTS_QTR1:          toFloat(row[7]),
					PTS_QTR2:          toFloat(row[8]),
					PTS_QTR3:          toFloat(row[9]),
					PTS_QTR4:          toFloat(row[10]),
					PTS_OT1:           toFloat(row[11]),
					PTS_OT2:           toFloat(row[12]),
					PTS_OT3:           toFloat(row[13]),
					PTS_OT4:           toFloat(row[14]),
					PTS_OT5:           toFloat(row[15]),
					PTS_OT6:           toFloat(row[16]),
					PTS_OT7:           toFloat(row[17]),
					PTS_OT8:           toFloat(row[18]),
					PTS_OT9:           toFloat(row[19]),
					PTS_OT10:          toFloat(row[20]),
					PTS:               toFloat(row[21]),
					FG_PCT:            toFloat(row[22]),
					FT_PCT:            toFloat(row[23]),
					FG3_PCT:           toFloat(row[24]),
					AST:               toFloat(row[25]),
					REB:               toFloat(row[26]),
					TOV:               toFloat(row[27]),
				}
				response.LineScore = append(response.LineScore, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "LastMeeting"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2LastMeeting{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: LastMeeting result set: %w", err)
		}
		response.LastMeeting = make([]BoxScoreSummaryV2LastMeeting, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 13 {
				item := BoxScoreSummaryV2LastMeeting{
					GAME_ID:                   toString(row[0]),
					GAME_DATE_EST:             toString(row[1]),
					GAME_DATE_TIME_EST:        toString(row[2]),
					HOME_TEAM_ID:              toInt(row[3]),
					HOME_TEAM_CITY:            toString(row[4]),
					HOME_TEAM_NAME:            toString(row[5]),
					HOME_TEAM_ABBREVIATION:    toString(row[6]),
					HOME_TEAM_POINTS:          toString(row[7]),
					VISITOR_TEAM_ID:           toInt(row[8]),
					VISITOR_TEAM_CITY:         toString(row[9]),
					VISITOR_TEAM_NAME:         toString(row[10]),
					VISITOR_TEAM_ABBREVIATION: toString(row[11]),
					VISITOR_TEAM_POINTS:       toString(row[12]),
				}
				response.LastMeeting = append(response.LastMeeting, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "SeasonSeries"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2SeasonSeries{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: SeasonSeries result set: %w", err)
		}
		response.SeasonSeries = make([]BoxScoreSummaryV2SeasonSeries, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 7 {
				item := BoxScoreSummaryV2SeasonSeries{
					GAME_ID:          toString(row[0]),
					HOME_TEAM_ID:     toInt(row[1]),
					VISITOR_TEAM_ID:  toInt(row[2]),
					GAME_DATE_EST:    toString(row[3]),
					HOME_TEAM_WINS:   toString(row[4]),
					HOME_TEAM_LOSSES: toString(row[5]),
					SERIES_LEADER:    toString(row[6]),
				}
				response.SeasonSeries = append(response.SeasonSeries, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "AvailableVideo"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(BoxScoreSummaryV2AvailableVideo{})); err != nil {
			return nil, fmt.Errorf("BoxScoreSummaryV2: AvailableVideo result set: %w", err)
		}
		response.AvailableVideo = make([]BoxScoreSummaryV2AvailableVideo, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 2 {
				item := BoxScoreSummaryV2AvailableVideo{
					GAME_ID:              toString(row[0]),
					VIDEO_AVAILABLE_FLAG: toInt(row[1]),
				}
				response.AvailableVideo = append(response.AvailableVideo, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
