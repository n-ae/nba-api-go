package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// LeagueDashTeamShotLocationsRequest contains parameters for the LeagueDashTeamShotLocations endpoint
type LeagueDashTeamShotLocationsRequest struct {
	Season *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode *parameters.PerMode
	LeagueID *parameters.LeagueID
}


// LeagueDashTeamShotLocationsShotLocations represents the ShotLocations result set for LeagueDashTeamShotLocations
type LeagueDashTeamShotLocationsShotLocations struct {
	TEAM_ID int `json:"TEAM_ID"`
	TEAM_NAME string `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	GP int `json:"GP"`
	W string `json:"W"`
	L string `json:"L"`
	W_PCT float64 `json:"W_PCT"`
	FGM_RA int `json:"FGM_RA"`
	FGA_RA int `json:"FGA_RA"`
	FG_PCT_RA string `json:"FG_PCT_RA"`
	FGM_IN_PAINT float64 `json:"FGM_IN_PAINT"`
	FGA_IN_PAINT float64 `json:"FGA_IN_PAINT"`
	FG_PCT_IN_PAINT string `json:"FG_PCT_IN_PAINT"`
	FGM_MID_RANGE float64 `json:"FGM_MID_RANGE"`
	FGA_MID_RANGE float64 `json:"FGA_MID_RANGE"`
	FG_PCT_MID_RANGE int `json:"FG_PCT_MID_RANGE"`
	FGM_LEFT_CORNER_3 float64 `json:"FGM_LEFT_CORNER_3"`
	FGA_LEFT_CORNER_3 float64 `json:"FGA_LEFT_CORNER_3"`
	FG_PCT_LEFT_CORNER_3 string `json:"FG_PCT_LEFT_CORNER_3"`
	FGM_RIGHT_CORNER_3 float64 `json:"FGM_RIGHT_CORNER_3"`
	FGA_RIGHT_CORNER_3 float64 `json:"FGA_RIGHT_CORNER_3"`
	FG_PCT_RIGHT_CORNER_3 string `json:"FG_PCT_RIGHT_CORNER_3"`
	FGM_ABOVE_BREAK_3 float64 `json:"FGM_ABOVE_BREAK_3"`
	FGA_ABOVE_BREAK_3 float64 `json:"FGA_ABOVE_BREAK_3"`
	FG_PCT_ABOVE_BREAK_3 string `json:"FG_PCT_ABOVE_BREAK_3"`
	FGM_BACKCOURT float64 `json:"FGM_BACKCOURT"`
	FGA_BACKCOURT float64 `json:"FGA_BACKCOURT"`
	FG_PCT_BACKCOURT string `json:"FG_PCT_BACKCOURT"`
}


// LeagueDashTeamShotLocationsResponse contains the response data from the LeagueDashTeamShotLocations endpoint
type LeagueDashTeamShotLocationsResponse struct {
	ShotLocations []LeagueDashTeamShotLocationsShotLocations
}

// GetLeagueDashTeamShotLocations retrieves data from the leaguedashteamshotlocations endpoint
func GetLeagueDashTeamShotLocations(ctx context.Context, client *stats.Client, req LeagueDashTeamShotLocationsRequest) (*models.Response[*LeagueDashTeamShotLocationsResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leaguedashteamshotlocations", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashTeamShotLocationsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.ShotLocations = make([]LeagueDashTeamShotLocationsShotLocations, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := LeagueDashTeamShotLocationsShotLocations{
					TEAM_ID: toInt(row[0]),
					TEAM_NAME: toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					GP: toInt(row[3]),
					W: toString(row[4]),
					L: toString(row[5]),
					W_PCT: toFloat(row[6]),
					FGM_RA: toInt(row[7]),
					FGA_RA: toInt(row[8]),
					FG_PCT_RA: toString(row[9]),
					FGM_IN_PAINT: toFloat(row[10]),
					FGA_IN_PAINT: toFloat(row[11]),
					FG_PCT_IN_PAINT: toString(row[12]),
					FGM_MID_RANGE: toFloat(row[13]),
					FGA_MID_RANGE: toFloat(row[14]),
					FG_PCT_MID_RANGE: toInt(row[15]),
					FGM_LEFT_CORNER_3: toFloat(row[16]),
					FGA_LEFT_CORNER_3: toFloat(row[17]),
					FG_PCT_LEFT_CORNER_3: toString(row[18]),
					FGM_RIGHT_CORNER_3: toFloat(row[19]),
					FGA_RIGHT_CORNER_3: toFloat(row[20]),
					FG_PCT_RIGHT_CORNER_3: toString(row[21]),
					FGM_ABOVE_BREAK_3: toFloat(row[22]),
					FGA_ABOVE_BREAK_3: toFloat(row[23]),
					FG_PCT_ABOVE_BREAK_3: toString(row[24]),
					FGM_BACKCOURT: toFloat(row[25]),
					FGA_BACKCOURT: toFloat(row[26]),
					FG_PCT_BACKCOURT: toString(row[27]),
				}
				response.ShotLocations = append(response.ShotLocations, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
