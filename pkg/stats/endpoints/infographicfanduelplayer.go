package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// InfographicFanDuelPlayerRequest contains parameters for the InfographicFanDuelPlayer endpoint
type InfographicFanDuelPlayerRequest struct {
	PlayerID string
}

// InfographicFanDuelPlayerFanDuelPlayer represents the FanDuelPlayer result set for InfographicFanDuelPlayer
type InfographicFanDuelPlayerFanDuelPlayer struct {
	PLAYER_ID   int     `json:"PLAYER_ID"`
	PLAYER_NAME string  `json:"PLAYER_NAME"`
	FD_POSITION string  `json:"FD_POSITION"`
	FD_SALARY   string  `json:"FD_SALARY"`
	FD_MINUTES  float64 `json:"FD_MINUTES"`
	FD_FG_PCT   float64 `json:"FD_FG_PCT"`
	FD_FT_PCT   float64 `json:"FD_FT_PCT"`
	FD_FG3_PCT  float64 `json:"FD_FG3_PCT"`
	FD_PTS      float64 `json:"FD_PTS"`
	FD_REB      float64 `json:"FD_REB"`
	FD_AST      float64 `json:"FD_AST"`
	FD_STL      float64 `json:"FD_STL"`
	FD_BLK      float64 `json:"FD_BLK"`
	FD_TOV      float64 `json:"FD_TOV"`
}

// InfographicFanDuelPlayerResponse contains the response data from the InfographicFanDuelPlayer endpoint
type InfographicFanDuelPlayerResponse struct {
	FanDuelPlayer []InfographicFanDuelPlayerFanDuelPlayer
}

// GetInfographicFanDuelPlayer retrieves data from the infographicfanduelplayer endpoint
func GetInfographicFanDuelPlayer(ctx context.Context, client *stats.Client, req InfographicFanDuelPlayerRequest) (*models.Response[*InfographicFanDuelPlayerResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "infographicfanduelplayer", params, &rawResp); err != nil {
		return nil, err
	}

	response := &InfographicFanDuelPlayerResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.FanDuelPlayer = make([]InfographicFanDuelPlayerFanDuelPlayer, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := InfographicFanDuelPlayerFanDuelPlayer{
					PLAYER_ID:   toInt(row[0]),
					PLAYER_NAME: toString(row[1]),
					FD_POSITION: toString(row[2]),
					FD_SALARY:   toString(row[3]),
					FD_MINUTES:  toFloat(row[4]),
					FD_FG_PCT:   toFloat(row[5]),
					FD_FT_PCT:   toFloat(row[6]),
					FD_FG3_PCT:  toFloat(row[7]),
					FD_PTS:      toFloat(row[8]),
					FD_REB:      toFloat(row[9]),
					FD_AST:      toFloat(row[10]),
					FD_STL:      toFloat(row[11]),
					FD_BLK:      toFloat(row[12]),
					FD_TOV:      toFloat(row[13]),
				}
				response.FanDuelPlayer = append(response.FanDuelPlayer, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
