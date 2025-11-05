package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingDrivesRequest contains parameters for the PlayerTrackingDrives endpoint
type PlayerTrackingDrivesRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingDrivesPlayerTrackingDrives represents the PlayerTrackingDrives result set for PlayerTrackingDrives
type PlayerTrackingDrivesPlayerTrackingDrives struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	DRIVES            string  `json:"DRIVES"`
	DRIVE_FGM         int     `json:"DRIVE_FGM"`
	DRIVE_FGA         int     `json:"DRIVE_FGA"`
	DRIVE_FG_PCT      float64 `json:"DRIVE_FG_PCT"`
	DRIVE_FTM         int     `json:"DRIVE_FTM"`
	DRIVE_FTA         int     `json:"DRIVE_FTA"`
	DRIVE_FT_PCT      float64 `json:"DRIVE_FT_PCT"`
	DRIVE_PTS         float64 `json:"DRIVE_PTS"`
	DRIVE_PASS        string  `json:"DRIVE_PASS"`
	DRIVE_AST         float64 `json:"DRIVE_AST"`
	DRIVE_TOV         float64 `json:"DRIVE_TOV"`
	DRIVE_PF          float64 `json:"DRIVE_PF"`
}

// PlayerTrackingDrivesResponse contains the response data from the PlayerTrackingDrives endpoint
type PlayerTrackingDrivesResponse struct {
	PlayerTrackingDrives []PlayerTrackingDrivesPlayerTrackingDrives
}

// GetPlayerTrackingDrives retrieves data from the playertrackingdrives endpoint
func GetPlayerTrackingDrives(ctx context.Context, client *stats.Client, req PlayerTrackingDrivesRequest) (*models.Response[*PlayerTrackingDrivesResponse], error) {
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
	if err := client.GetJSON(ctx, "playertrackingdrives", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingDrivesResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingDrives = make([]PlayerTrackingDrivesPlayerTrackingDrives, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 18 {
				item := PlayerTrackingDrivesPlayerTrackingDrives{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					DRIVES:            toString(row[6]),
					DRIVE_FGM:         toInt(row[7]),
					DRIVE_FGA:         toInt(row[8]),
					DRIVE_FG_PCT:      toFloat(row[9]),
					DRIVE_FTM:         toInt(row[10]),
					DRIVE_FTA:         toInt(row[11]),
					DRIVE_FT_PCT:      toFloat(row[12]),
					DRIVE_PTS:         toFloat(row[13]),
					DRIVE_PASS:        toString(row[14]),
					DRIVE_AST:         toFloat(row[15]),
					DRIVE_TOV:         toFloat(row[16]),
					DRIVE_PF:          toFloat(row[17]),
				}
				response.PlayerTrackingDrives = append(response.PlayerTrackingDrives, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
