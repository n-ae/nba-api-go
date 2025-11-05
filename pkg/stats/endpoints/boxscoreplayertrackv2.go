package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// BoxScorePlayerTrackV2Request contains parameters for the BoxScorePlayerTrackV2 endpoint
type BoxScorePlayerTrackV2Request struct {
	GameID string
}

// BoxScorePlayerTrackV2PlayerTrack represents the PlayerTrack result set for BoxScorePlayerTrackV2
type BoxScorePlayerTrackV2PlayerTrack struct {
	GAME_ID           string  `json:"GAME_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_CITY         string  `json:"TEAM_CITY"`
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	START_POSITION    string  `json:"START_POSITION"`
	COMMENT           string  `json:"COMMENT"`
	MIN               float64 `json:"MIN"`
	SPD               string  `json:"SPD"`
	DIST              string  `json:"DIST"`
	ORBC              string  `json:"ORBC"`
	DRBC              string  `json:"DRBC"`
	RBC               string  `json:"RBC"`
	TCHS              string  `json:"TCHS"`
	SAST              float64 `json:"SAST"`
	FTAST             float64 `json:"FTAST"`
	PASS              string  `json:"PASS"`
	AST               float64 `json:"AST"`
	CFGM              int     `json:"CFGM"`
	CFGA              int     `json:"CFGA"`
	CFG_PCT           float64 `json:"CFG_PCT"`
	UFGM              int     `json:"UFGM"`
	UFGA              int     `json:"UFGA"`
	UFG_PCT           float64 `json:"UFG_PCT"`
	FG_PCT            float64 `json:"FG_PCT"`
	DFGM              int     `json:"DFGM"`
	DFGA              int     `json:"DFGA"`
	DFG_PCT           float64 `json:"DFG_PCT"`
}

// BoxScorePlayerTrackV2Response contains the response data from the BoxScorePlayerTrackV2 endpoint
type BoxScorePlayerTrackV2Response struct {
	PlayerTrack []BoxScorePlayerTrackV2PlayerTrack
}

// GetBoxScorePlayerTrackV2 retrieves data from the boxscoreplayertrackv2 endpoint
func GetBoxScorePlayerTrackV2(ctx context.Context, client *stats.Client, req BoxScorePlayerTrackV2Request) (*models.Response[*BoxScorePlayerTrackV2Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "boxscoreplayertrackv2", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScorePlayerTrackV2Response{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrack = make([]BoxScorePlayerTrackV2PlayerTrack, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 29 {
				item := BoxScorePlayerTrackV2PlayerTrack{
					GAME_ID:           toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_CITY:         toString(row[3]),
					PLAYER_ID:         toInt(row[4]),
					PLAYER_NAME:       toString(row[5]),
					START_POSITION:    toString(row[6]),
					COMMENT:           toString(row[7]),
					MIN:               toFloat(row[8]),
					SPD:               toString(row[9]),
					DIST:              toString(row[10]),
					ORBC:              toString(row[11]),
					DRBC:              toString(row[12]),
					RBC:               toString(row[13]),
					TCHS:              toString(row[14]),
					SAST:              toFloat(row[15]),
					FTAST:             toFloat(row[16]),
					PASS:              toString(row[17]),
					AST:               toFloat(row[18]),
					CFGM:              toInt(row[19]),
					CFGA:              toInt(row[20]),
					CFG_PCT:           toFloat(row[21]),
					UFGM:              toInt(row[22]),
					UFGA:              toInt(row[23]),
					UFG_PCT:           toFloat(row[24]),
					FG_PCT:            toFloat(row[25]),
					DFGM:              toInt(row[26]),
					DFGA:              toInt(row[27]),
					DFG_PCT:           toFloat(row[28]),
				}
				response.PlayerTrack = append(response.PlayerTrack, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
