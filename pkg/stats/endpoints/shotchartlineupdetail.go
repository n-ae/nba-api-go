package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// ShotChartLineupDetailRequest contains parameters for the ShotChartLineupDetail endpoint
type ShotChartLineupDetailRequest struct {
	Season         *parameters.Season
	SeasonType     *parameters.SeasonType
	TeamID         *string
	GroupID        *string
	LeagueID       *parameters.LeagueID
	ContextMeasure *string
}

// ShotChartLineupDetailShot_Chart_Detail represents the Shot_Chart_Detail result set for ShotChartLineupDetail
type ShotChartLineupDetailShot_Chart_Detail struct {
	GRID_TYPE           string  `json:"GRID_TYPE"`
	GAME_ID             string  `json:"GAME_ID"`
	GAME_EVENT_ID       string  `json:"GAME_EVENT_ID"`
	PLAYER_ID           int     `json:"PLAYER_ID"`
	PLAYER_NAME         string  `json:"PLAYER_NAME"`
	TEAM_ID             int     `json:"TEAM_ID"`
	TEAM_NAME           string  `json:"TEAM_NAME"`
	PERIOD              int     `json:"PERIOD"`
	MINUTES_REMAINING   float64 `json:"MINUTES_REMAINING"`
	SECONDS_REMAINING   string  `json:"SECONDS_REMAINING"`
	EVENT_TYPE          string  `json:"EVENT_TYPE"`
	ACTION_TYPE         string  `json:"ACTION_TYPE"`
	SHOT_TYPE           string  `json:"SHOT_TYPE"`
	SHOT_ZONE_BASIC     string  `json:"SHOT_ZONE_BASIC"`
	SHOT_ZONE_AREA      string  `json:"SHOT_ZONE_AREA"`
	SHOT_ZONE_RANGE     int     `json:"SHOT_ZONE_RANGE"`
	SHOT_DISTANCE       string  `json:"SHOT_DISTANCE"`
	LOC_X               string  `json:"LOC_X"`
	LOC_Y               string  `json:"LOC_Y"`
	SHOT_ATTEMPTED_FLAG string  `json:"SHOT_ATTEMPTED_FLAG"`
	SHOT_MADE_FLAG      string  `json:"SHOT_MADE_FLAG"`
	GAME_DATE           string  `json:"GAME_DATE"`
	HTM                 string  `json:"HTM"`
	VTM                 string  `json:"VTM"`
}

// ShotChartLineupDetailLeagueAverages represents the LeagueAverages result set for ShotChartLineupDetail
type ShotChartLineupDetailLeagueAverages struct {
	GRID_TYPE       string  `json:"GRID_TYPE"`
	SHOT_ZONE_BASIC string  `json:"SHOT_ZONE_BASIC"`
	SHOT_ZONE_AREA  string  `json:"SHOT_ZONE_AREA"`
	SHOT_ZONE_RANGE int     `json:"SHOT_ZONE_RANGE"`
	FGA             int     `json:"FGA"`
	FGM             int     `json:"FGM"`
	FG_PCT          float64 `json:"FG_PCT"`
}

// ShotChartLineupDetailResponse contains the response data from the ShotChartLineupDetail endpoint
type ShotChartLineupDetailResponse struct {
	Shot_Chart_Detail []ShotChartLineupDetailShot_Chart_Detail
	LeagueAverages    []ShotChartLineupDetailLeagueAverages
}

// GetShotChartLineupDetail retrieves data from the shotchartlineupdetail endpoint
func GetShotChartLineupDetail(ctx context.Context, client *stats.Client, req ShotChartLineupDetailRequest) (*models.Response[*ShotChartLineupDetailResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.TeamID != nil {
		params.Set("TeamID", string(*req.TeamID))
	}
	if req.GroupID != nil {
		params.Set("GroupID", string(*req.GroupID))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.ContextMeasure != nil {
		params.Set("ContextMeasure", string(*req.ContextMeasure))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "shotchartlineupdetail", params, &rawResp); err != nil {
		return nil, err
	}

	response := &ShotChartLineupDetailResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Shot_Chart_Detail = make([]ShotChartLineupDetailShot_Chart_Detail, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 24 {
				item := ShotChartLineupDetailShot_Chart_Detail{
					GRID_TYPE:           toString(row[0]),
					GAME_ID:             toString(row[1]),
					GAME_EVENT_ID:       toString(row[2]),
					PLAYER_ID:           toInt(row[3]),
					PLAYER_NAME:         toString(row[4]),
					TEAM_ID:             toInt(row[5]),
					TEAM_NAME:           toString(row[6]),
					PERIOD:              toInt(row[7]),
					MINUTES_REMAINING:   toFloat(row[8]),
					SECONDS_REMAINING:   toString(row[9]),
					EVENT_TYPE:          toString(row[10]),
					ACTION_TYPE:         toString(row[11]),
					SHOT_TYPE:           toString(row[12]),
					SHOT_ZONE_BASIC:     toString(row[13]),
					SHOT_ZONE_AREA:      toString(row[14]),
					SHOT_ZONE_RANGE:     toInt(row[15]),
					SHOT_DISTANCE:       toString(row[16]),
					LOC_X:               toString(row[17]),
					LOC_Y:               toString(row[18]),
					SHOT_ATTEMPTED_FLAG: toString(row[19]),
					SHOT_MADE_FLAG:      toString(row[20]),
					GAME_DATE:           toString(row[21]),
					HTM:                 toString(row[22]),
					VTM:                 toString(row[23]),
				}
				response.Shot_Chart_Detail = append(response.Shot_Chart_Detail, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.LeagueAverages = make([]ShotChartLineupDetailLeagueAverages, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 7 {
				item := ShotChartLineupDetailLeagueAverages{
					GRID_TYPE:       toString(row[0]),
					SHOT_ZONE_BASIC: toString(row[1]),
					SHOT_ZONE_AREA:  toString(row[2]),
					SHOT_ZONE_RANGE: toInt(row[3]),
					FGA:             toInt(row[4]),
					FGM:             toInt(row[5]),
					FG_PCT:          toFloat(row[6]),
				}
				response.LeagueAverages = append(response.LeagueAverages, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
