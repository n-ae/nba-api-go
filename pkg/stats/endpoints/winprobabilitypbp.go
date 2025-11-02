package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
)

// WinProbabilityPBPRequest contains parameters for the WinProbabilityPBP endpoint
type WinProbabilityPBPRequest struct {
	GameID  string
	RunType *string
}

// WinProbabilityPBPWinProbPBP represents the WinProbPBP result set for WinProbabilityPBP
type WinProbabilityPBPWinProbPBP struct {
	GAME_ID           string  `json:"GAME_ID"`
	EVENT_NUM         string  `json:"EVENT_NUM"`
	HOME_PCT          float64 `json:"HOME_PCT"`
	VISITOR_PCT       float64 `json:"VISITOR_PCT"`
	HOME_SCORE        string  `json:"HOME_SCORE"`
	VISITOR_SCORE     string  `json:"VISITOR_SCORE"`
	SCORE_MARGIN      string  `json:"SCORE_MARGIN"`
	HOME_PTS_EST      float64 `json:"HOME_PTS_EST"`
	VISITOR_PTS_EST   float64 `json:"VISITOR_PTS_EST"`
	HOME_PTS_RANGE    float64 `json:"HOME_PTS_RANGE"`
	VISITOR_PTS_RANGE float64 `json:"VISITOR_PTS_RANGE"`
	PERIOD            int     `json:"PERIOD"`
	SECONDS_REMAINING string  `json:"SECONDS_REMAINING"`
}

// WinProbabilityPBPGameInfo represents the GameInfo result set for WinProbabilityPBP
type WinProbabilityPBPGameInfo struct {
	GAME_ID             string `json:"GAME_ID"`
	HOME_TEAM_ID        int    `json:"HOME_TEAM_ID"`
	VISITOR_TEAM_ID     int    `json:"VISITOR_TEAM_ID"`
	HOME_TEAM_ABR       string `json:"HOME_TEAM_ABR"`
	VISITOR_TEAM_ABR    string `json:"VISITOR_TEAM_ABR"`
	HOME_FINAL_SCORE    string `json:"HOME_FINAL_SCORE"`
	VISITOR_FINAL_SCORE string `json:"VISITOR_FINAL_SCORE"`
}

// WinProbabilityPBPResponse contains the response data from the WinProbabilityPBP endpoint
type WinProbabilityPBPResponse struct {
	WinProbPBP []WinProbabilityPBPWinProbPBP
	GameInfo   []WinProbabilityPBPGameInfo
}

// GetWinProbabilityPBP retrieves data from the winprobabilitypbp endpoint
func GetWinProbabilityPBP(ctx context.Context, client *stats.Client, req WinProbabilityPBPRequest) (*models.Response[*WinProbabilityPBPResponse], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.RunType != nil {
		params.Set("RunType", string(*req.RunType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/winprobabilitypbp", params, &rawResp); err != nil {
		return nil, err
	}

	response := &WinProbabilityPBPResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.WinProbPBP = make([]WinProbabilityPBPWinProbPBP, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 13 {
				item := WinProbabilityPBPWinProbPBP{
					GAME_ID:           toString(row[0]),
					EVENT_NUM:         toString(row[1]),
					HOME_PCT:          toFloat(row[2]),
					VISITOR_PCT:       toFloat(row[3]),
					HOME_SCORE:        toString(row[4]),
					VISITOR_SCORE:     toString(row[5]),
					SCORE_MARGIN:      toString(row[6]),
					HOME_PTS_EST:      toFloat(row[7]),
					VISITOR_PTS_EST:   toFloat(row[8]),
					HOME_PTS_RANGE:    toFloat(row[9]),
					VISITOR_PTS_RANGE: toFloat(row[10]),
					PERIOD:            toInt(row[11]),
					SECONDS_REMAINING: toString(row[12]),
				}
				response.WinProbPBP = append(response.WinProbPBP, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.GameInfo = make([]WinProbabilityPBPGameInfo, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 7 {
				item := WinProbabilityPBPGameInfo{
					GAME_ID:             toString(row[0]),
					HOME_TEAM_ID:        toInt(row[1]),
					VISITOR_TEAM_ID:     toInt(row[2]),
					HOME_TEAM_ABR:       toString(row[3]),
					VISITOR_TEAM_ABR:    toString(row[4]),
					HOME_FINAL_SCORE:    toString(row[5]),
					VISITOR_FINAL_SCORE: toString(row[6]),
				}
				response.GameInfo = append(response.GameInfo, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
