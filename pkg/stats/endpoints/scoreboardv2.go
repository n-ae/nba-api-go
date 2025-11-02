package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// ScoreboardV2Request contains parameters for the ScoreboardV2 endpoint
type ScoreboardV2Request struct {
	GameDate  string
	LeagueID  *parameters.LeagueID
	DayOffset *string
}

// ScoreboardV2GameHeader represents the GameHeader result set for ScoreboardV2
type ScoreboardV2GameHeader struct {
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

// ScoreboardV2LineScore represents the LineScore result set for ScoreboardV2
type ScoreboardV2LineScore struct {
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

// ScoreboardV2SeriesStandings represents the SeriesStandings result set for ScoreboardV2
type ScoreboardV2SeriesStandings struct {
	GAME_ID          string `json:"GAME_ID"`
	HOME_TEAM_ID     int    `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID  int    `json:"VISITOR_TEAM_ID"`
	GAME_DATE_EST    string `json:"GAME_DATE_EST"`
	HOME_TEAM_WINS   string `json:"HOME_TEAM_WINS"`
	HOME_TEAM_LOSSES string `json:"HOME_TEAM_LOSSES"`
	SERIES_LEADER    string `json:"SERIES_LEADER"`
}

// ScoreboardV2LastMeeting represents the LastMeeting result set for ScoreboardV2
type ScoreboardV2LastMeeting struct {
	GAME_ID                             string  `json:"GAME_ID"`
	LAST_GAME_ID                        string  `json:"LAST_GAME_ID"`
	LAST_GAME_DATE_EST                  string  `json:"LAST_GAME_DATE_EST"`
	LAST_GAME_HOME_TEAM_ID              int     `json:"LAST_GAME_HOME_TEAM_ID"`
	LAST_GAME_HOME_TEAM_CITY            string  `json:"LAST_GAME_HOME_TEAM_CITY"`
	LAST_GAME_HOME_TEAM_NAME            string  `json:"LAST_GAME_HOME_TEAM_NAME"`
	LAST_GAME_HOME_TEAM_ABBREVIATION    string  `json:"LAST_GAME_HOME_TEAM_ABBREVIATION"`
	LAST_GAME_HOME_TEAM_POINTS          float64 `json:"LAST_GAME_HOME_TEAM_POINTS"`
	LAST_GAME_VISITOR_TEAM_ID           int     `json:"LAST_GAME_VISITOR_TEAM_ID"`
	LAST_GAME_VISITOR_TEAM_CITY         string  `json:"LAST_GAME_VISITOR_TEAM_CITY"`
	LAST_GAME_VISITOR_TEAM_NAME         string  `json:"LAST_GAME_VISITOR_TEAM_NAME"`
	LAST_GAME_VISITOR_TEAM_ABBREVIATION string  `json:"LAST_GAME_VISITOR_TEAM_ABBREVIATION"`
	LAST_GAME_VISITOR_TEAM_POINTS       float64 `json:"LAST_GAME_VISITOR_TEAM_POINTS"`
}

// ScoreboardV2EastConfStandingsByDay represents the EastConfStandingsByDay result set for ScoreboardV2
type ScoreboardV2EastConfStandingsByDay struct {
	TEAM_ID       int     `json:"TEAM_ID"`
	LEAGUE_ID     string  `json:"LEAGUE_ID"`
	SEASON_ID     string  `json:"SEASON_ID"`
	STANDINGSDATE string  `json:"STANDINGSDATE"`
	CONFERENCE    string  `json:"CONFERENCE"`
	TEAM          string  `json:"TEAM"`
	G             string  `json:"G"`
	W             string  `json:"W"`
	L             string  `json:"L"`
	W_PCT         float64 `json:"W_PCT"`
	HOME_RECORD   string  `json:"HOME_RECORD"`
	ROAD_RECORD   string  `json:"ROAD_RECORD"`
}

// ScoreboardV2WestConfStandingsByDay represents the WestConfStandingsByDay result set for ScoreboardV2
type ScoreboardV2WestConfStandingsByDay struct {
	TEAM_ID       int     `json:"TEAM_ID"`
	LEAGUE_ID     string  `json:"LEAGUE_ID"`
	SEASON_ID     string  `json:"SEASON_ID"`
	STANDINGSDATE string  `json:"STANDINGSDATE"`
	CONFERENCE    string  `json:"CONFERENCE"`
	TEAM          string  `json:"TEAM"`
	G             string  `json:"G"`
	W             string  `json:"W"`
	L             string  `json:"L"`
	W_PCT         float64 `json:"W_PCT"`
	HOME_RECORD   string  `json:"HOME_RECORD"`
	ROAD_RECORD   string  `json:"ROAD_RECORD"`
}

// ScoreboardV2Available represents the Available result set for ScoreboardV2
type ScoreboardV2Available struct {
	GAME_ID      string `json:"GAME_ID"`
	PT_AVAILABLE string `json:"PT_AVAILABLE"`
}

// ScoreboardV2Response contains the response data from the ScoreboardV2 endpoint
type ScoreboardV2Response struct {
	GameHeader             []ScoreboardV2GameHeader
	LineScore              []ScoreboardV2LineScore
	SeriesStandings        []ScoreboardV2SeriesStandings
	LastMeeting            []ScoreboardV2LastMeeting
	EastConfStandingsByDay []ScoreboardV2EastConfStandingsByDay
	WestConfStandingsByDay []ScoreboardV2WestConfStandingsByDay
	Available              []ScoreboardV2Available
}

// GetScoreboardV2 retrieves data from the scoreboardv2 endpoint
func GetScoreboardV2(ctx context.Context, client *stats.Client, req ScoreboardV2Request) (*models.Response[*ScoreboardV2Response], error) {
	params := url.Values{}
	if req.GameDate == "" {
		return nil, fmt.Errorf("GameDate is required")
	}
	params.Set("GameDate", string(req.GameDate))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.DayOffset != nil {
		params.Set("DayOffset", string(*req.DayOffset))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/scoreboardv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &ScoreboardV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.GameHeader = make([]ScoreboardV2GameHeader, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := ScoreboardV2GameHeader{
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
				response.GameHeader = append(response.GameHeader, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.LineScore = make([]ScoreboardV2LineScore, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 28 {
				item := ScoreboardV2LineScore{
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
	if len(rawResp.ResultSets) > 2 {
		response.SeriesStandings = make([]ScoreboardV2SeriesStandings, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 7 {
				item := ScoreboardV2SeriesStandings{
					GAME_ID:          toString(row[0]),
					HOME_TEAM_ID:     toInt(row[1]),
					VISITOR_TEAM_ID:  toInt(row[2]),
					GAME_DATE_EST:    toString(row[3]),
					HOME_TEAM_WINS:   toString(row[4]),
					HOME_TEAM_LOSSES: toString(row[5]),
					SERIES_LEADER:    toString(row[6]),
				}
				response.SeriesStandings = append(response.SeriesStandings, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.LastMeeting = make([]ScoreboardV2LastMeeting, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 13 {
				item := ScoreboardV2LastMeeting{
					GAME_ID:                             toString(row[0]),
					LAST_GAME_ID:                        toString(row[1]),
					LAST_GAME_DATE_EST:                  toString(row[2]),
					LAST_GAME_HOME_TEAM_ID:              toInt(row[3]),
					LAST_GAME_HOME_TEAM_CITY:            toString(row[4]),
					LAST_GAME_HOME_TEAM_NAME:            toString(row[5]),
					LAST_GAME_HOME_TEAM_ABBREVIATION:    toString(row[6]),
					LAST_GAME_HOME_TEAM_POINTS:          toFloat(row[7]),
					LAST_GAME_VISITOR_TEAM_ID:           toInt(row[8]),
					LAST_GAME_VISITOR_TEAM_CITY:         toString(row[9]),
					LAST_GAME_VISITOR_TEAM_NAME:         toString(row[10]),
					LAST_GAME_VISITOR_TEAM_ABBREVIATION: toString(row[11]),
					LAST_GAME_VISITOR_TEAM_POINTS:       toFloat(row[12]),
				}
				response.LastMeeting = append(response.LastMeeting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.EastConfStandingsByDay = make([]ScoreboardV2EastConfStandingsByDay, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 12 {
				item := ScoreboardV2EastConfStandingsByDay{
					TEAM_ID:       toInt(row[0]),
					LEAGUE_ID:     toString(row[1]),
					SEASON_ID:     toString(row[2]),
					STANDINGSDATE: toString(row[3]),
					CONFERENCE:    toString(row[4]),
					TEAM:          toString(row[5]),
					G:             toString(row[6]),
					W:             toString(row[7]),
					L:             toString(row[8]),
					W_PCT:         toFloat(row[9]),
					HOME_RECORD:   toString(row[10]),
					ROAD_RECORD:   toString(row[11]),
				}
				response.EastConfStandingsByDay = append(response.EastConfStandingsByDay, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 5 {
		response.WestConfStandingsByDay = make([]ScoreboardV2WestConfStandingsByDay, 0, len(rawResp.ResultSets[5].RowSet))
		for _, row := range rawResp.ResultSets[5].RowSet {
			if len(row) >= 12 {
				item := ScoreboardV2WestConfStandingsByDay{
					TEAM_ID:       toInt(row[0]),
					LEAGUE_ID:     toString(row[1]),
					SEASON_ID:     toString(row[2]),
					STANDINGSDATE: toString(row[3]),
					CONFERENCE:    toString(row[4]),
					TEAM:          toString(row[5]),
					G:             toString(row[6]),
					W:             toString(row[7]),
					L:             toString(row[8]),
					W_PCT:         toFloat(row[9]),
					HOME_RECORD:   toString(row[10]),
					ROAD_RECORD:   toString(row[11]),
				}
				response.WestConfStandingsByDay = append(response.WestConfStandingsByDay, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 6 {
		response.Available = make([]ScoreboardV2Available, 0, len(rawResp.ResultSets[6].RowSet))
		for _, row := range rawResp.ResultSets[6].RowSet {
			if len(row) >= 2 {
				item := ScoreboardV2Available{
					GAME_ID:      toString(row[0]),
					PT_AVAILABLE: toString(row[1]),
				}
				response.Available = append(response.Available, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
