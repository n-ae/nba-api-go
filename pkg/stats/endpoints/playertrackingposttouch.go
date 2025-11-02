package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerTrackingPostTouchRequest contains parameters for the PlayerTrackingPostTouch endpoint
type PlayerTrackingPostTouchRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// PlayerTrackingPostTouchPlayerTrackingPostTouch represents the PlayerTrackingPostTouch result set for PlayerTrackingPostTouch
type PlayerTrackingPostTouchPlayerTrackingPostTouch struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MIN               float64 `json:"MIN"`
	POST_TOUCHES      string  `json:"POST_TOUCHES"`
	POST_TOUCH_FGM    int     `json:"POST_TOUCH_FGM"`
	POST_TOUCH_FGA    int     `json:"POST_TOUCH_FGA"`
	POST_TOUCH_FG_PCT float64 `json:"POST_TOUCH_FG_PCT"`
	POST_TOUCH_FTM    int     `json:"POST_TOUCH_FTM"`
	POST_TOUCH_FTA    int     `json:"POST_TOUCH_FTA"`
	POST_TOUCH_FT_PCT float64 `json:"POST_TOUCH_FT_PCT"`
	POST_TOUCH_PTS    float64 `json:"POST_TOUCH_PTS"`
	POST_TOUCH_PASS   string  `json:"POST_TOUCH_PASS"`
	POST_TOUCH_AST    float64 `json:"POST_TOUCH_AST"`
	POST_TOUCH_TOV    float64 `json:"POST_TOUCH_TOV"`
}

// PlayerTrackingPostTouchResponse contains the response data from the PlayerTrackingPostTouch endpoint
type PlayerTrackingPostTouchResponse struct {
	PlayerTrackingPostTouch []PlayerTrackingPostTouchPlayerTrackingPostTouch
}

// GetPlayerTrackingPostTouch retrieves data from the playertrackingposttouch endpoint
func GetPlayerTrackingPostTouch(ctx context.Context, client *stats.Client, req PlayerTrackingPostTouchRequest) (*models.Response[*PlayerTrackingPostTouchResponse], error) {
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
	if err := client.GetJSON(ctx, "/playertrackingposttouch", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerTrackingPostTouchResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerTrackingPostTouch = make([]PlayerTrackingPostTouchPlayerTrackingPostTouch, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 17 {
				item := PlayerTrackingPostTouchPlayerTrackingPostTouch{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					GP:                toInt(row[4]),
					MIN:               toFloat(row[5]),
					POST_TOUCHES:      toString(row[6]),
					POST_TOUCH_FGM:    toInt(row[7]),
					POST_TOUCH_FGA:    toInt(row[8]),
					POST_TOUCH_FG_PCT: toFloat(row[9]),
					POST_TOUCH_FTM:    toInt(row[10]),
					POST_TOUCH_FTA:    toInt(row[11]),
					POST_TOUCH_FT_PCT: toFloat(row[12]),
					POST_TOUCH_PTS:    toFloat(row[13]),
					POST_TOUCH_PASS:   toString(row[14]),
					POST_TOUCH_AST:    toFloat(row[15]),
					POST_TOUCH_TOV:    toFloat(row[16]),
				}
				response.PlayerTrackingPostTouch = append(response.PlayerTrackingPostTouch, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
