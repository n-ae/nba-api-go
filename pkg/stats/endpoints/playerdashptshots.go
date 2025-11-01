package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerDashPtShotsRequest contains parameters for the PlayerDashPtShots endpoint
type PlayerDashPtShotsRequest struct {
	PlayerID   string
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerDashPtShotsOverallShooting represents the OverallShooting result set for PlayerDashPtShots
type PlayerDashPtShotsOverallShooting struct {
	PLAYER_ID              int     `json:"PLAYER_ID"`
	PLAYER_NAME_LAST_FIRST float64 `json:"PLAYER_NAME_LAST_FIRST"`
	SORT_ORDER             string  `json:"SORT_ORDER"`
	GP                     int     `json:"GP"`
	G                      string  `json:"G"`
	FGA_FREQUENCY          float64 `json:"FGA_FREQUENCY"`
	FGM                    int     `json:"FGM"`
	FGA                    int     `json:"FGA"`
	FG_PCT                 float64 `json:"FG_PCT"`
	EFG_PCT                float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY         string  `json:"FG2A_FREQUENCY"`
	FG2M                   string  `json:"FG2M"`
	FG2A                   string  `json:"FG2A"`
	FG2_PCT                float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY         float64 `json:"FG3A_FREQUENCY"`
	FG3M                   int     `json:"FG3M"`
	FG3A                   int     `json:"FG3A"`
	FG3_PCT                float64 `json:"FG3_PCT"`
}

// PlayerDashPtShotsGeneralShooting represents the GeneralShooting result set for PlayerDashPtShots
type PlayerDashPtShotsGeneralShooting struct {
	PLAYER_ID              int     `json:"PLAYER_ID"`
	PLAYER_NAME_LAST_FIRST float64 `json:"PLAYER_NAME_LAST_FIRST"`
	SHOT_TYPE              string  `json:"SHOT_TYPE"`
	FGA_FREQUENCY          float64 `json:"FGA_FREQUENCY"`
	FGM                    int     `json:"FGM"`
	FGA                    int     `json:"FGA"`
	FG_PCT                 float64 `json:"FG_PCT"`
	EFG_PCT                float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY         string  `json:"FG2A_FREQUENCY"`
	FG2M                   string  `json:"FG2M"`
	FG2A                   string  `json:"FG2A"`
	FG2_PCT                float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY         float64 `json:"FG3A_FREQUENCY"`
	FG3M                   int     `json:"FG3M"`
	FG3A                   int     `json:"FG3A"`
	FG3_PCT                float64 `json:"FG3_PCT"`
}

// PlayerDashPtShotsShotClockShooting represents the ShotClockShooting result set for PlayerDashPtShots
type PlayerDashPtShotsShotClockShooting struct {
	PLAYER_ID              int     `json:"PLAYER_ID"`
	PLAYER_NAME_LAST_FIRST float64 `json:"PLAYER_NAME_LAST_FIRST"`
	SHOT_CLOCK_RANGE       int     `json:"SHOT_CLOCK_RANGE"`
	FGA_FREQUENCY          float64 `json:"FGA_FREQUENCY"`
	FGM                    int     `json:"FGM"`
	FGA                    int     `json:"FGA"`
	FG_PCT                 float64 `json:"FG_PCT"`
	EFG_PCT                float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY         string  `json:"FG2A_FREQUENCY"`
	FG2M                   string  `json:"FG2M"`
	FG2A                   string  `json:"FG2A"`
	FG2_PCT                float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY         float64 `json:"FG3A_FREQUENCY"`
	FG3M                   int     `json:"FG3M"`
	FG3A                   int     `json:"FG3A"`
	FG3_PCT                float64 `json:"FG3_PCT"`
}

// PlayerDashPtShotsDribbleShooting represents the DribbleShooting result set for PlayerDashPtShots
type PlayerDashPtShotsDribbleShooting struct {
	PLAYER_ID              int     `json:"PLAYER_ID"`
	PLAYER_NAME_LAST_FIRST float64 `json:"PLAYER_NAME_LAST_FIRST"`
	DRIBBLE_RANGE          int     `json:"DRIBBLE_RANGE"`
	FGA_FREQUENCY          float64 `json:"FGA_FREQUENCY"`
	FGM                    int     `json:"FGM"`
	FGA                    int     `json:"FGA"`
	FG_PCT                 float64 `json:"FG_PCT"`
	EFG_PCT                float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY         string  `json:"FG2A_FREQUENCY"`
	FG2M                   string  `json:"FG2M"`
	FG2A                   string  `json:"FG2A"`
	FG2_PCT                float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY         float64 `json:"FG3A_FREQUENCY"`
	FG3M                   int     `json:"FG3M"`
	FG3A                   int     `json:"FG3A"`
	FG3_PCT                float64 `json:"FG3_PCT"`
}

// PlayerDashPtShotsClosestDefenderShooting represents the ClosestDefenderShooting result set for PlayerDashPtShots
type PlayerDashPtShotsClosestDefenderShooting struct {
	PLAYER_ID              int     `json:"PLAYER_ID"`
	PLAYER_NAME_LAST_FIRST float64 `json:"PLAYER_NAME_LAST_FIRST"`
	CLOSE_DEF_DIST_RANGE   int     `json:"CLOSE_DEF_DIST_RANGE"`
	FGA_FREQUENCY          float64 `json:"FGA_FREQUENCY"`
	FGM                    int     `json:"FGM"`
	FGA                    int     `json:"FGA"`
	FG_PCT                 float64 `json:"FG_PCT"`
	EFG_PCT                float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY         string  `json:"FG2A_FREQUENCY"`
	FG2M                   string  `json:"FG2M"`
	FG2A                   string  `json:"FG2A"`
	FG2_PCT                float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY         float64 `json:"FG3A_FREQUENCY"`
	FG3M                   int     `json:"FG3M"`
	FG3A                   int     `json:"FG3A"`
	FG3_PCT                float64 `json:"FG3_PCT"`
}

// PlayerDashPtShotsTouchTimeShooting represents the TouchTimeShooting result set for PlayerDashPtShots
type PlayerDashPtShotsTouchTimeShooting struct {
	PLAYER_ID              int     `json:"PLAYER_ID"`
	PLAYER_NAME_LAST_FIRST float64 `json:"PLAYER_NAME_LAST_FIRST"`
	TOUCH_TIME_RANGE       int     `json:"TOUCH_TIME_RANGE"`
	FGA_FREQUENCY          float64 `json:"FGA_FREQUENCY"`
	FGM                    int     `json:"FGM"`
	FGA                    int     `json:"FGA"`
	FG_PCT                 float64 `json:"FG_PCT"`
	EFG_PCT                float64 `json:"EFG_PCT"`
	FG2A_FREQUENCY         string  `json:"FG2A_FREQUENCY"`
	FG2M                   string  `json:"FG2M"`
	FG2A                   string  `json:"FG2A"`
	FG2_PCT                float64 `json:"FG2_PCT"`
	FG3A_FREQUENCY         float64 `json:"FG3A_FREQUENCY"`
	FG3M                   int     `json:"FG3M"`
	FG3A                   int     `json:"FG3A"`
	FG3_PCT                float64 `json:"FG3_PCT"`
}

// PlayerDashPtShotsResponse contains the response data from the PlayerDashPtShots endpoint
type PlayerDashPtShotsResponse struct {
	OverallShooting         []PlayerDashPtShotsOverallShooting
	GeneralShooting         []PlayerDashPtShotsGeneralShooting
	ShotClockShooting       []PlayerDashPtShotsShotClockShooting
	DribbleShooting         []PlayerDashPtShotsDribbleShooting
	ClosestDefenderShooting []PlayerDashPtShotsClosestDefenderShooting
	TouchTimeShooting       []PlayerDashPtShotsTouchTimeShooting
}

// GetPlayerDashPtShots retrieves data from the playerdashptshots endpoint
func GetPlayerDashPtShots(ctx context.Context, client *stats.Client, req PlayerDashPtShotsRequest) (*models.Response[*PlayerDashPtShotsResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
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
	if err := client.GetJSON(ctx, "/playerdashptshots", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerDashPtShotsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.OverallShooting = make([]PlayerDashPtShotsOverallShooting, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 18 {
				item := PlayerDashPtShotsOverallShooting{
					PLAYER_ID:              toInt(row[0]),
					PLAYER_NAME_LAST_FIRST: toFloat(row[1]),
					SORT_ORDER:             toString(row[2]),
					GP:                     toInt(row[3]),
					G:                      toString(row[4]),
					FGA_FREQUENCY:          toFloat(row[5]),
					FGM:                    toInt(row[6]),
					FGA:                    toInt(row[7]),
					FG_PCT:                 toFloat(row[8]),
					EFG_PCT:                toFloat(row[9]),
					FG2A_FREQUENCY:         toString(row[10]),
					FG2M:                   toString(row[11]),
					FG2A:                   toString(row[12]),
					FG2_PCT:                toFloat(row[13]),
					FG3A_FREQUENCY:         toFloat(row[14]),
					FG3M:                   toInt(row[15]),
					FG3A:                   toInt(row[16]),
					FG3_PCT:                toFloat(row[17]),
				}
				response.OverallShooting = append(response.OverallShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.GeneralShooting = make([]PlayerDashPtShotsGeneralShooting, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 16 {
				item := PlayerDashPtShotsGeneralShooting{
					PLAYER_ID:              toInt(row[0]),
					PLAYER_NAME_LAST_FIRST: toFloat(row[1]),
					SHOT_TYPE:              toString(row[2]),
					FGA_FREQUENCY:          toFloat(row[3]),
					FGM:                    toInt(row[4]),
					FGA:                    toInt(row[5]),
					FG_PCT:                 toFloat(row[6]),
					EFG_PCT:                toFloat(row[7]),
					FG2A_FREQUENCY:         toString(row[8]),
					FG2M:                   toString(row[9]),
					FG2A:                   toString(row[10]),
					FG2_PCT:                toFloat(row[11]),
					FG3A_FREQUENCY:         toFloat(row[12]),
					FG3M:                   toInt(row[13]),
					FG3A:                   toInt(row[14]),
					FG3_PCT:                toFloat(row[15]),
				}
				response.GeneralShooting = append(response.GeneralShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 2 {
		response.ShotClockShooting = make([]PlayerDashPtShotsShotClockShooting, 0, len(rawResp.ResultSets[2].RowSet))
		for _, row := range rawResp.ResultSets[2].RowSet {
			if len(row) >= 16 {
				item := PlayerDashPtShotsShotClockShooting{
					PLAYER_ID:              toInt(row[0]),
					PLAYER_NAME_LAST_FIRST: toFloat(row[1]),
					SHOT_CLOCK_RANGE:       toInt(row[2]),
					FGA_FREQUENCY:          toFloat(row[3]),
					FGM:                    toInt(row[4]),
					FGA:                    toInt(row[5]),
					FG_PCT:                 toFloat(row[6]),
					EFG_PCT:                toFloat(row[7]),
					FG2A_FREQUENCY:         toString(row[8]),
					FG2M:                   toString(row[9]),
					FG2A:                   toString(row[10]),
					FG2_PCT:                toFloat(row[11]),
					FG3A_FREQUENCY:         toFloat(row[12]),
					FG3M:                   toInt(row[13]),
					FG3A:                   toInt(row[14]),
					FG3_PCT:                toFloat(row[15]),
				}
				response.ShotClockShooting = append(response.ShotClockShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 3 {
		response.DribbleShooting = make([]PlayerDashPtShotsDribbleShooting, 0, len(rawResp.ResultSets[3].RowSet))
		for _, row := range rawResp.ResultSets[3].RowSet {
			if len(row) >= 16 {
				item := PlayerDashPtShotsDribbleShooting{
					PLAYER_ID:              toInt(row[0]),
					PLAYER_NAME_LAST_FIRST: toFloat(row[1]),
					DRIBBLE_RANGE:          toInt(row[2]),
					FGA_FREQUENCY:          toFloat(row[3]),
					FGM:                    toInt(row[4]),
					FGA:                    toInt(row[5]),
					FG_PCT:                 toFloat(row[6]),
					EFG_PCT:                toFloat(row[7]),
					FG2A_FREQUENCY:         toString(row[8]),
					FG2M:                   toString(row[9]),
					FG2A:                   toString(row[10]),
					FG2_PCT:                toFloat(row[11]),
					FG3A_FREQUENCY:         toFloat(row[12]),
					FG3M:                   toInt(row[13]),
					FG3A:                   toInt(row[14]),
					FG3_PCT:                toFloat(row[15]),
				}
				response.DribbleShooting = append(response.DribbleShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 4 {
		response.ClosestDefenderShooting = make([]PlayerDashPtShotsClosestDefenderShooting, 0, len(rawResp.ResultSets[4].RowSet))
		for _, row := range rawResp.ResultSets[4].RowSet {
			if len(row) >= 16 {
				item := PlayerDashPtShotsClosestDefenderShooting{
					PLAYER_ID:              toInt(row[0]),
					PLAYER_NAME_LAST_FIRST: toFloat(row[1]),
					CLOSE_DEF_DIST_RANGE:   toInt(row[2]),
					FGA_FREQUENCY:          toFloat(row[3]),
					FGM:                    toInt(row[4]),
					FGA:                    toInt(row[5]),
					FG_PCT:                 toFloat(row[6]),
					EFG_PCT:                toFloat(row[7]),
					FG2A_FREQUENCY:         toString(row[8]),
					FG2M:                   toString(row[9]),
					FG2A:                   toString(row[10]),
					FG2_PCT:                toFloat(row[11]),
					FG3A_FREQUENCY:         toFloat(row[12]),
					FG3M:                   toInt(row[13]),
					FG3A:                   toInt(row[14]),
					FG3_PCT:                toFloat(row[15]),
				}
				response.ClosestDefenderShooting = append(response.ClosestDefenderShooting, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 5 {
		response.TouchTimeShooting = make([]PlayerDashPtShotsTouchTimeShooting, 0, len(rawResp.ResultSets[5].RowSet))
		for _, row := range rawResp.ResultSets[5].RowSet {
			if len(row) >= 16 {
				item := PlayerDashPtShotsTouchTimeShooting{
					PLAYER_ID:              toInt(row[0]),
					PLAYER_NAME_LAST_FIRST: toFloat(row[1]),
					TOUCH_TIME_RANGE:       toInt(row[2]),
					FGA_FREQUENCY:          toFloat(row[3]),
					FGM:                    toInt(row[4]),
					FGA:                    toInt(row[5]),
					FG_PCT:                 toFloat(row[6]),
					EFG_PCT:                toFloat(row[7]),
					FG2A_FREQUENCY:         toString(row[8]),
					FG2M:                   toString(row[9]),
					FG2A:                   toString(row[10]),
					FG2_PCT:                toFloat(row[11]),
					FG3A_FREQUENCY:         toFloat(row[12]),
					FG3M:                   toInt(row[13]),
					FG3A:                   toInt(row[14]),
					FG3_PCT:                toFloat(row[15]),
				}
				response.TouchTimeShooting = append(response.TouchTimeShooting, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
