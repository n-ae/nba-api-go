package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPlayerShotLocationsRequest contains parameters for the LeagueDashPlayerShotLocations endpoint
type LeagueDashPlayerShotLocationsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueDashPlayerShotLocationsShotLocations represents the ShotLocations result set for LeagueDashPlayerShotLocations
type LeagueDashPlayerShotLocationsShotLocations struct {
	PLAYER_ID             int     `json:"PLAYER_ID"`
	PLAYER_NAME           string  `json:"PLAYER_NAME"`
	TEAM_ID               int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION     string  `json:"TEAM_ABBREVIATION"`
	AGE                   int     `json:"AGE"`
	GP                    int     `json:"GP"`
	W                     string  `json:"W"`
	L                     string  `json:"L"`
	W_PCT                 float64 `json:"W_PCT"`
	FGM_RA                int     `json:"FGM_RA"`
	FGA_RA                int     `json:"FGA_RA"`
	FG_PCT_RA             string  `json:"FG_PCT_RA"`
	FGM_IN_PAINT          float64 `json:"FGM_IN_PAINT"`
	FGA_IN_PAINT          float64 `json:"FGA_IN_PAINT"`
	FG_PCT_IN_PAINT       string  `json:"FG_PCT_IN_PAINT"`
	FGM_MID_RANGE         float64 `json:"FGM_MID_RANGE"`
	FGA_MID_RANGE         float64 `json:"FGA_MID_RANGE"`
	FG_PCT_MID_RANGE      int     `json:"FG_PCT_MID_RANGE"`
	FGM_LEFT_CORNER_3     float64 `json:"FGM_LEFT_CORNER_3"`
	FGA_LEFT_CORNER_3     float64 `json:"FGA_LEFT_CORNER_3"`
	FG_PCT_LEFT_CORNER_3  string  `json:"FG_PCT_LEFT_CORNER_3"`
	FGM_RIGHT_CORNER_3    float64 `json:"FGM_RIGHT_CORNER_3"`
	FGA_RIGHT_CORNER_3    float64 `json:"FGA_RIGHT_CORNER_3"`
	FG_PCT_RIGHT_CORNER_3 string  `json:"FG_PCT_RIGHT_CORNER_3"`
	FGM_ABOVE_BREAK_3     float64 `json:"FGM_ABOVE_BREAK_3"`
	FGA_ABOVE_BREAK_3     float64 `json:"FGA_ABOVE_BREAK_3"`
	FG_PCT_ABOVE_BREAK_3  string  `json:"FG_PCT_ABOVE_BREAK_3"`
	FGM_BACKCOURT         float64 `json:"FGM_BACKCOURT"`
	FGA_BACKCOURT         float64 `json:"FGA_BACKCOURT"`
	FG_PCT_BACKCOURT      string  `json:"FG_PCT_BACKCOURT"`
}

// LeagueDashPlayerShotLocationsResponse contains the response data from the LeagueDashPlayerShotLocations endpoint
type LeagueDashPlayerShotLocationsResponse struct {
	ShotLocations []LeagueDashPlayerShotLocationsShotLocations
}

// GetLeagueDashPlayerShotLocations retrieves data from the leaguedashplayershotlocations endpoint
func GetLeagueDashPlayerShotLocations(ctx context.Context, client *stats.Client, req LeagueDashPlayerShotLocationsRequest) (*models.Response[*LeagueDashPlayerShotLocationsResponse], error) {
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
	if err := client.GetJSON(ctx, "/leaguedashplayershotlocations", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPlayerShotLocationsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.ShotLocations = make([]LeagueDashPlayerShotLocationsShotLocations, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 30 {
				item := LeagueDashPlayerShotLocationsShotLocations{
					PLAYER_ID:             toInt(row[0]),
					PLAYER_NAME:           toString(row[1]),
					TEAM_ID:               toInt(row[2]),
					TEAM_ABBREVIATION:     toString(row[3]),
					AGE:                   toInt(row[4]),
					GP:                    toInt(row[5]),
					W:                     toString(row[6]),
					L:                     toString(row[7]),
					W_PCT:                 toFloat(row[8]),
					FGM_RA:                toInt(row[9]),
					FGA_RA:                toInt(row[10]),
					FG_PCT_RA:             toString(row[11]),
					FGM_IN_PAINT:          toFloat(row[12]),
					FGA_IN_PAINT:          toFloat(row[13]),
					FG_PCT_IN_PAINT:       toString(row[14]),
					FGM_MID_RANGE:         toFloat(row[15]),
					FGA_MID_RANGE:         toFloat(row[16]),
					FG_PCT_MID_RANGE:      toInt(row[17]),
					FGM_LEFT_CORNER_3:     toFloat(row[18]),
					FGA_LEFT_CORNER_3:     toFloat(row[19]),
					FG_PCT_LEFT_CORNER_3:  toString(row[20]),
					FGM_RIGHT_CORNER_3:    toFloat(row[21]),
					FGA_RIGHT_CORNER_3:    toFloat(row[22]),
					FG_PCT_RIGHT_CORNER_3: toString(row[23]),
					FGM_ABOVE_BREAK_3:     toFloat(row[24]),
					FGA_ABOVE_BREAK_3:     toFloat(row[25]),
					FG_PCT_ABOVE_BREAK_3:  toString(row[26]),
					FGM_BACKCOURT:         toFloat(row[27]),
					FGA_BACKCOURT:         toFloat(row[28]),
					FG_PCT_BACKCOURT:      toString(row[29]),
				}
				response.ShotLocations = append(response.ShotLocations, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
