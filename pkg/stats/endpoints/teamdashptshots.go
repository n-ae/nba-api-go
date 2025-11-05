package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// TeamDashPtShotsRequest contains parameters for the TeamDashPtShots endpoint
type TeamDashPtShotsRequest struct {
	TeamID     string
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// TeamDashPtShotsOverallShooting represents the OverallShooting result set for TeamDashPtShots
type TeamDashPtShotsOverallShooting struct {
	TEAM_ID        int     `json:"TEAM_ID"`
	TEAM_NAME      string  `json:"TEAM_NAME"`
	SORT_ORDER     string  `json:"SORT_ORDER"`
	GP             int     `json:"GP"`
	G              string  `json:"G"`
	FGA_FREQUENCY  float64 `json:"FGA_FREQUENCY"`
	FGM            int     `json:"FGM"`
	FGA            int     `json:"FGA"`
	FG_PCT         float64 `json:"FG_PCT"`
	EFG_PCT        float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY string  `json:"FG2A_FREQUENCY"`
	FG2M           string  `json:"FG2M"`
	FG2A           string  `json:"FG2A"`
	FG2_PCT        float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY float64 `json:"FG3A_FREQUENCY"`
	FG3M           int     `json:"FG3M"`
	FG3A           int     `json:"FG3A"`
	FG3_PCT        float64 `json:"FG3_PCT"`
}

// TeamDashPtShotsGeneralShooting represents the GeneralShooting result set for TeamDashPtShots
type TeamDashPtShotsGeneralShooting struct {
	TEAM_ID        int     `json:"TEAM_ID"`
	TEAM_NAME      string  `json:"TEAM_NAME"`
	SHOT_TYPE      string  `json:"SHOT_TYPE"`
	FGA_FREQUENCY  float64 `json:"FGA_FREQUENCY"`
	FGM            int     `json:"FGM"`
	FGA            int     `json:"FGA"`
	FG_PCT         float64 `json:"FG_PCT"`
	EFG_PCT        float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY string  `json:"FG2A_FREQUENCY"`
	FG2M           string  `json:"FG2M"`
	FG2A           string  `json:"FG2A"`
	FG2_PCT        float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY float64 `json:"FG3A_FREQUENCY"`
	FG3M           int     `json:"FG3M"`
	FG3A           int     `json:"FG3A"`
	FG3_PCT        float64 `json:"FG3_PCT"`
}

// TeamDashPtShotsShotClockShooting represents the ShotClockShooting result set for TeamDashPtShots
type TeamDashPtShotsShotClockShooting struct {
	TEAM_ID          int     `json:"TEAM_ID"`
	TEAM_NAME        string  `json:"TEAM_NAME"`
	SHOT_CLOCK_RANGE int     `json:"SHOT_CLOCK_RANGE"`
	FGA_FREQUENCY    float64 `json:"FGA_FREQUENCY"`
	FGM              int     `json:"FGM"`
	FGA              int     `json:"FGA"`
	FG_PCT           float64 `json:"FG_PCT"`
	EFG_PCT          float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY   string  `json:"FG2A_FREQUENCY"`
	FG2M             string  `json:"FG2M"`
	FG2A             string  `json:"FG2A"`
	FG2_PCT          float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY   float64 `json:"FG3A_FREQUENCY"`
	FG3M             int     `json:"FG3M"`
	FG3A             int     `json:"FG3A"`
	FG3_PCT          float64 `json:"FG3_PCT"`
}

// TeamDashPtShotsResponse contains the response data from the TeamDashPtShots endpoint
type TeamDashPtShotsResponse struct {
	OverallShooting   []TeamDashPtShotsOverallShooting
	GeneralShooting   []TeamDashPtShotsGeneralShooting
	ShotClockShooting []TeamDashPtShotsShotClockShooting
}

// GetTeamDashPtShots retrieves data from the teamdashptshots endpoint
func GetTeamDashPtShots(ctx context.Context, client *stats.Client, req TeamDashPtShotsRequest) (*models.Response[*TeamDashPtShotsResponse], error) {
	params := url.Values{}
	if req.TeamID == "" {
		return nil, fmt.Errorf("TeamID is required")
	}
	params.Set("TeamID", string(req.TeamID))
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
	if err := client.GetJSON(ctx, "teamdashptshots", params, &rawResp); err != nil {
		return nil, err
	}

	response := &TeamDashPtShotsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallShooting = make([]TeamDashPtShotsOverallShooting, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 18 {
				item := TeamDashPtShotsOverallShooting{
					TEAM_ID:        toInt(row[0]),
					TEAM_NAME:      toString(row[1]),
					SORT_ORDER:     toString(row[2]),
					GP:             toInt(row[3]),
					G:              toString(row[4]),
					FGA_FREQUENCY:  toFloat(row[5]),
					FGM:            toInt(row[6]),
					FGA:            toInt(row[7]),
					FG_PCT:         toFloat(row[8]),
					EFG_PCT:        toFloat(row[9]),
					FG2A_FREQUENCY: toString(row[10]),
					FG2M:           toString(row[11]),
					FG2A:           toString(row[12]),
					FG2_PCT:        toFloat(row[13]),
					FG3A_FREQUENCY: toFloat(row[14]),
					FG3M:           toInt(row[15]),
					FG3A:           toInt(row[16]),
					FG3_PCT:        toFloat(row[17]),
				}
				response.OverallShooting = append(response.OverallShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.GeneralShooting = make([]TeamDashPtShotsGeneralShooting, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 16 {
				item := TeamDashPtShotsGeneralShooting{
					TEAM_ID:        toInt(row[0]),
					TEAM_NAME:      toString(row[1]),
					SHOT_TYPE:      toString(row[2]),
					FGA_FREQUENCY:  toFloat(row[3]),
					FGM:            toInt(row[4]),
					FGA:            toInt(row[5]),
					FG_PCT:         toFloat(row[6]),
					EFG_PCT:        toFloat(row[7]),
					FG2A_FREQUENCY: toString(row[8]),
					FG2M:           toString(row[9]),
					FG2A:           toString(row[10]),
					FG2_PCT:        toFloat(row[11]),
					FG3A_FREQUENCY: toFloat(row[12]),
					FG3M:           toInt(row[13]),
					FG3A:           toInt(row[14]),
					FG3_PCT:        toFloat(row[15]),
				}
				response.GeneralShooting = append(response.GeneralShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.ShotClockShooting = make([]TeamDashPtShotsShotClockShooting, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 16 {
				item := TeamDashPtShotsShotClockShooting{
					TEAM_ID:          toInt(row[0]),
					TEAM_NAME:        toString(row[1]),
					SHOT_CLOCK_RANGE: toInt(row[2]),
					FGA_FREQUENCY:    toFloat(row[3]),
					FGM:              toInt(row[4]),
					FGA:              toInt(row[5]),
					FG_PCT:           toFloat(row[6]),
					EFG_PCT:          toFloat(row[7]),
					FG2A_FREQUENCY:   toString(row[8]),
					FG2M:             toString(row[9]),
					FG2A:             toString(row[10]),
					FG2_PCT:          toFloat(row[11]),
					FG3A_FREQUENCY:   toFloat(row[12]),
					FG3M:             toInt(row[13]),
					FG3A:             toInt(row[14]),
					FG3_PCT:          toFloat(row[15]),
				}
				response.ShotClockShooting = append(response.ShotClockShooting, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
