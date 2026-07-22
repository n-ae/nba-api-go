package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/v3/pkg/models"
	"github.com/n-ae/nba-api-go/v3/pkg/stats"
	"github.com/n-ae/nba-api-go/v3/pkg/stats/parameters"
)

// ShotChartDetailRequest contains parameters for the ShotChartDetail endpoint
type ShotChartDetailRequest struct {
	PlayerID       *string
	TeamID         *string
	GameID         *string
	Season         parameters.Season
	SeasonType     parameters.SeasonType
	LeagueID       *parameters.LeagueID
	ContextMeasure *string
	DateFrom       *string
	DateTo         *string
	OpponentTeamID *string
	Period         *string
	RookieYear     *string
	VsConference   *string
	VsDivision     *string
	Position       *string
	GameSegment    *string
	LastNGames     *string
	Location       *string
	Month          *string
	Outcome        *string
	SeasonSegment  *string
	AheadBehind    *string
	ClutchTime     *string
	PointDiff      *string
	RangeType      *string
	StartPeriod    *string
	EndPeriod      *string
	StartRange     *string
	EndRange       *string
}

// ShotChartDetailShot_Chart_Detail represents the Shot_Chart_Detail result set for ShotChartDetail
type ShotChartDetailShot_Chart_Detail struct {
	GRID_TYPE           string  `json:"GRID_TYPE"`
	GAME_ID             string  `json:"GAME_ID"`
	GAME_EVENT_ID       int     `json:"GAME_EVENT_ID"`
	PLAYER_ID           int     `json:"PLAYER_ID"`
	PLAYER_NAME         string  `json:"PLAYER_NAME"`
	TEAM_ID             int     `json:"TEAM_ID"`
	TEAM_NAME           string  `json:"TEAM_NAME"`
	PERIOD              int     `json:"PERIOD"`
	MINUTES_REMAINING   int     `json:"MINUTES_REMAINING"`
	SECONDS_REMAINING   int     `json:"SECONDS_REMAINING"`
	EVENT_TYPE          string  `json:"EVENT_TYPE"`
	ACTION_TYPE         string  `json:"ACTION_TYPE"`
	SHOT_TYPE           string  `json:"SHOT_TYPE"`
	SHOT_ZONE_BASIC     string  `json:"SHOT_ZONE_BASIC"`
	SHOT_ZONE_AREA      string  `json:"SHOT_ZONE_AREA"`
	SHOT_ZONE_RANGE     string  `json:"SHOT_ZONE_RANGE"`
	SHOT_DISTANCE       float64 `json:"SHOT_DISTANCE"`
	LOC_X               float64 `json:"LOC_X"`
	LOC_Y               float64 `json:"LOC_Y"`
	SHOT_ATTEMPTED_FLAG int     `json:"SHOT_ATTEMPTED_FLAG"`
	SHOT_MADE_FLAG      int     `json:"SHOT_MADE_FLAG"`
	GAME_DATE           string  `json:"GAME_DATE"`
	HTM                 string  `json:"HTM"`
	VTM                 string  `json:"VTM"`
}

// ShotChartDetailLeagueAverages represents the LeagueAverages result set for ShotChartDetail
type ShotChartDetailLeagueAverages struct {
	GRID_TYPE       string  `json:"GRID_TYPE"`
	SHOT_ZONE_BASIC string  `json:"SHOT_ZONE_BASIC"`
	SHOT_ZONE_AREA  string  `json:"SHOT_ZONE_AREA"`
	SHOT_ZONE_RANGE string  `json:"SHOT_ZONE_RANGE"`
	FGA             int     `json:"FGA"`
	FGM             int     `json:"FGM"`
	FG_PCT          float64 `json:"FG_PCT"`
}

// ShotChartDetailResponse contains the response data from the ShotChartDetail endpoint
type ShotChartDetailResponse struct {
	Shot_Chart_Detail []ShotChartDetailShot_Chart_Detail
	LeagueAverages    []ShotChartDetailLeagueAverages
}

// GetShotChartDetail retrieves data from the shotchartdetail endpoint
func GetShotChartDetail(ctx context.Context, client *stats.Client, req ShotChartDetailRequest) (*models.Response[*ShotChartDetailResponse], error) {
	params := url.Values{}
	if req.PlayerID != nil {
		params.Set("PlayerID", *req.PlayerID)
	}
	if req.TeamID != nil {
		params.Set("TeamID", *req.TeamID)
	}
	if req.GameID != nil {
		params.Set("GameID", *req.GameID)
	}
	if req.Season == "" {
		return nil, fmt.Errorf("%s is required", "Season")
	}
	params.Set("Season", string(req.Season))
	if req.SeasonType == "" {
		return nil, fmt.Errorf("%s is required", "SeasonType")
	}
	params.Set("SeasonType", string(req.SeasonType))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.ContextMeasure != nil {
		params.Set("ContextMeasure", *req.ContextMeasure)
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", *req.DateFrom)
	}
	if req.DateTo != nil {
		params.Set("DateTo", *req.DateTo)
	}
	if req.OpponentTeamID != nil {
		params.Set("OpponentTeamID", *req.OpponentTeamID)
	}
	if req.Period != nil {
		params.Set("Period", *req.Period)
	}
	if req.RookieYear != nil {
		params.Set("RookieYear", *req.RookieYear)
	}
	if req.VsConference != nil {
		params.Set("VsConference", *req.VsConference)
	}
	if req.VsDivision != nil {
		params.Set("VsDivision", *req.VsDivision)
	}
	if req.Position != nil {
		params.Set("Position", *req.Position)
	}
	if req.GameSegment != nil {
		params.Set("GameSegment", *req.GameSegment)
	}
	if req.LastNGames != nil {
		params.Set("LastNGames", *req.LastNGames)
	}
	if req.Location != nil {
		params.Set("Location", *req.Location)
	}
	if req.Month != nil {
		params.Set("Month", *req.Month)
	}
	if req.Outcome != nil {
		params.Set("Outcome", *req.Outcome)
	}
	if req.SeasonSegment != nil {
		params.Set("SeasonSegment", *req.SeasonSegment)
	}
	if req.AheadBehind != nil {
		params.Set("AheadBehind", *req.AheadBehind)
	}
	if req.ClutchTime != nil {
		params.Set("ClutchTime", *req.ClutchTime)
	}
	if req.PointDiff != nil {
		params.Set("PointDiff", *req.PointDiff)
	}
	if req.RangeType != nil {
		params.Set("RangeType", *req.RangeType)
	}
	if req.StartPeriod != nil {
		params.Set("StartPeriod", *req.StartPeriod)
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", *req.EndPeriod)
	}
	if req.StartRange != nil {
		params.Set("StartRange", *req.StartRange)
	}
	if req.EndRange != nil {
		params.Set("EndRange", *req.EndRange)
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "shotchartdetail", params, &rawResp); err != nil {
		return nil, err
	}

	response := &ShotChartDetailResponse{}
	if rs, ok := findResultSet(rawResp.ResultSets, "Shot_Chart_Detail"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(ShotChartDetailShot_Chart_Detail{})); err != nil {
			return nil, fmt.Errorf("ShotChartDetail: Shot_Chart_Detail result set: %w", err)
		}
		response.Shot_Chart_Detail = make([]ShotChartDetailShot_Chart_Detail, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 24 {
				item := ShotChartDetailShot_Chart_Detail{
					GRID_TYPE:           toString(row[0]),
					GAME_ID:             toString(row[1]),
					GAME_EVENT_ID:       toInt(row[2]),
					PLAYER_ID:           toInt(row[3]),
					PLAYER_NAME:         toString(row[4]),
					TEAM_ID:             toInt(row[5]),
					TEAM_NAME:           toString(row[6]),
					PERIOD:              toInt(row[7]),
					MINUTES_REMAINING:   toInt(row[8]),
					SECONDS_REMAINING:   toInt(row[9]),
					EVENT_TYPE:          toString(row[10]),
					ACTION_TYPE:         toString(row[11]),
					SHOT_TYPE:           toString(row[12]),
					SHOT_ZONE_BASIC:     toString(row[13]),
					SHOT_ZONE_AREA:      toString(row[14]),
					SHOT_ZONE_RANGE:     toString(row[15]),
					SHOT_DISTANCE:       toFloat(row[16]),
					LOC_X:               toFloat(row[17]),
					LOC_Y:               toFloat(row[18]),
					SHOT_ATTEMPTED_FLAG: toInt(row[19]),
					SHOT_MADE_FLAG:      toInt(row[20]),
					GAME_DATE:           toString(row[21]),
					HTM:                 toString(row[22]),
					VTM:                 toString(row[23]),
				}
				response.Shot_Chart_Detail = append(response.Shot_Chart_Detail, item)
			}
		}
	}
	if rs, ok := findResultSet(rawResp.ResultSets, "LeagueAverages"); ok {
		if err := validateHeaders(rs.Headers, jsonTags(ShotChartDetailLeagueAverages{})); err != nil {
			return nil, fmt.Errorf("ShotChartDetail: LeagueAverages result set: %w", err)
		}
		response.LeagueAverages = make([]ShotChartDetailLeagueAverages, 0, len(rs.RowSet))
		for _, row := range rs.RowSet {
			if len(row) >= 7 {
				item := ShotChartDetailLeagueAverages{
					GRID_TYPE:       toString(row[0]),
					SHOT_ZONE_BASIC: toString(row[1]),
					SHOT_ZONE_AREA:  toString(row[2]),
					SHOT_ZONE_RANGE: toString(row[3]),
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
