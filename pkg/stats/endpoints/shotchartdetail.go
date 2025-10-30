package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// ShotChartDetailRequest contains parameters for the ShotChartDetail endpoint
type ShotChartDetailRequest struct {
	PlayerID *string
	TeamID *string
	GameID *string
	Season parameters.Season
	SeasonType parameters.SeasonType
	LeagueID *parameters.LeagueID
	ContextMeasure *string
	DateFrom *string
	DateTo *string
	OpponentTeamID *string
	Period *string
	RookieYear *string
	VsConference *string
	VsDivision *string
	Position *string
	GameSegment *string
	LastNGames *string
	Location *string
	Month *string
	Outcome *string
	SeasonSegment *string
	AheadBehind *string
	ClutchTime *string
	PointDiff *string
	RangeType *string
	StartPeriod *string
	EndPeriod *string
	StartRange *string
	EndRange *string
}


// ShotChartDetailShot_Chart_Detail represents the Shot_Chart_Detail result set for ShotChartDetail
type ShotChartDetailShot_Chart_Detail struct {
	GRID_TYPE interface{}
	GAME_ID interface{}
	GAME_EVENT_ID interface{}
	PLAYER_ID interface{}
	PLAYER_NAME interface{}
	TEAM_ID interface{}
	TEAM_NAME interface{}
	PERIOD interface{}
	MINUTES_REMAINING interface{}
	SECONDS_REMAINING interface{}
	EVENT_TYPE interface{}
	ACTION_TYPE interface{}
	SHOT_TYPE interface{}
	SHOT_ZONE_BASIC interface{}
	SHOT_ZONE_AREA interface{}
	SHOT_ZONE_RANGE interface{}
	SHOT_DISTANCE interface{}
	LOC_X interface{}
	LOC_Y interface{}
	SHOT_ATTEMPTED_FLAG interface{}
	SHOT_MADE_FLAG interface{}
	GAME_DATE interface{}
	HTM interface{}
	VTM interface{}
}

// ShotChartDetailLeagueAverages represents the LeagueAverages result set for ShotChartDetail
type ShotChartDetailLeagueAverages struct {
	GRID_TYPE interface{}
	SHOT_ZONE_BASIC interface{}
	SHOT_ZONE_AREA interface{}
	SHOT_ZONE_RANGE interface{}
	FGA interface{}
	FGM interface{}
	FG_PCT interface{}
}


// ShotChartDetailResponse contains the response data from the ShotChartDetail endpoint
type ShotChartDetailResponse struct {
	Shot_Chart_Detail []ShotChartDetailShot_Chart_Detail
	LeagueAverages []ShotChartDetailLeagueAverages
}

// GetShotChartDetail retrieves data from the shotchartdetail endpoint
func GetShotChartDetail(ctx context.Context, client *stats.Client, req ShotChartDetailRequest) (*models.Response[*ShotChartDetailResponse], error) {
	params := url.Values{}
	if req.PlayerID != nil {
		params.Set("PlayerID", string(*req.PlayerID))
	}
	if req.TeamID != nil {
		params.Set("TeamID", string(*req.TeamID))
	}
	if req.GameID != nil {
		params.Set("GameID", string(*req.GameID))
	}
	if req.Season == "" {
		return nil, fmt.Errorf("Season is required")
	}
	params.Set("Season", string(req.Season))
	if req.SeasonType == "" {
		return nil, fmt.Errorf("SeasonType is required")
	}
	params.Set("SeasonType", string(req.SeasonType))
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.ContextMeasure != nil {
		params.Set("ContextMeasure", string(*req.ContextMeasure))
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", string(*req.DateFrom))
	}
	if req.DateTo != nil {
		params.Set("DateTo", string(*req.DateTo))
	}
	if req.OpponentTeamID != nil {
		params.Set("OpponentTeamID", string(*req.OpponentTeamID))
	}
	if req.Period != nil {
		params.Set("Period", string(*req.Period))
	}
	if req.RookieYear != nil {
		params.Set("RookieYear", string(*req.RookieYear))
	}
	if req.VsConference != nil {
		params.Set("VsConference", string(*req.VsConference))
	}
	if req.VsDivision != nil {
		params.Set("VsDivision", string(*req.VsDivision))
	}
	if req.Position != nil {
		params.Set("Position", string(*req.Position))
	}
	if req.GameSegment != nil {
		params.Set("GameSegment", string(*req.GameSegment))
	}
	if req.LastNGames != nil {
		params.Set("LastNGames", string(*req.LastNGames))
	}
	if req.Location != nil {
		params.Set("Location", string(*req.Location))
	}
	if req.Month != nil {
		params.Set("Month", string(*req.Month))
	}
	if req.Outcome != nil {
		params.Set("Outcome", string(*req.Outcome))
	}
	if req.SeasonSegment != nil {
		params.Set("SeasonSegment", string(*req.SeasonSegment))
	}
	if req.AheadBehind != nil {
		params.Set("AheadBehind", string(*req.AheadBehind))
	}
	if req.ClutchTime != nil {
		params.Set("ClutchTime", string(*req.ClutchTime))
	}
	if req.PointDiff != nil {
		params.Set("PointDiff", string(*req.PointDiff))
	}
	if req.RangeType != nil {
		params.Set("RangeType", string(*req.RangeType))
	}
	if req.StartPeriod != nil {
		params.Set("StartPeriod", string(*req.StartPeriod))
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", string(*req.EndPeriod))
	}
	if req.StartRange != nil {
		params.Set("StartRange", string(*req.StartRange))
	}
	if req.EndRange != nil {
		params.Set("EndRange", string(*req.EndRange))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/shotchartdetail", params, &rawResp); err != nil {
		return nil, err
	}

	response := &ShotChartDetailResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.Shot_Chart_Detail = make([]ShotChartDetailShot_Chart_Detail, len(rawResp.ResultSets[0].RowSet))
		for i, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 24 {
				response.Shot_Chart_Detail[i] = ShotChartDetailShot_Chart_Detail{
					GRID_TYPE: row[0],
					GAME_ID: row[1],
					GAME_EVENT_ID: row[2],
					PLAYER_ID: row[3],
					PLAYER_NAME: row[4],
					TEAM_ID: row[5],
					TEAM_NAME: row[6],
					PERIOD: row[7],
					MINUTES_REMAINING: row[8],
					SECONDS_REMAINING: row[9],
					EVENT_TYPE: row[10],
					ACTION_TYPE: row[11],
					SHOT_TYPE: row[12],
					SHOT_ZONE_BASIC: row[13],
					SHOT_ZONE_AREA: row[14],
					SHOT_ZONE_RANGE: row[15],
					SHOT_DISTANCE: row[16],
					LOC_X: row[17],
					LOC_Y: row[18],
					SHOT_ATTEMPTED_FLAG: row[19],
					SHOT_MADE_FLAG: row[20],
					GAME_DATE: row[21],
					HTM: row[22],
					VTM: row[23],
				}
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.LeagueAverages = make([]ShotChartDetailLeagueAverages, len(rawResp.ResultSets[1].RowSet))
		for i, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 7 {
				response.LeagueAverages[i] = ShotChartDetailLeagueAverages{
					GRID_TYPE: row[0],
					SHOT_ZONE_BASIC: row[1],
					SHOT_ZONE_AREA: row[2],
					SHOT_ZONE_RANGE: row[3],
					FGA: row[4],
					FGM: row[5],
					FG_PCT: row[6],
				}
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
