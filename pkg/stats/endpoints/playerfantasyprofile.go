package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// PlayerFantasyProfileRequest contains parameters for the PlayerFantasyProfile endpoint
type PlayerFantasyProfileRequest struct {
	PlayerID string
}

// PlayerFantasyProfileLastNGames represents the LastNGames result set for PlayerFantasyProfile
type PlayerFantasyProfileLastNGames struct {
	PLAYER_ID       int     `json:"PLAYER_ID"`
	PLAYER_NAME     string  `json:"PLAYER_NAME"`
	LAST_N_GAMES    float64 `json:"LAST_N_GAMES"`
	GP              int     `json:"GP"`
	MIN             float64 `json:"MIN"`
	FGM             int     `json:"FGM"`
	FGA             int     `json:"FGA"`
	FG_PCT          float64 `json:"FG_PCT"`
	FG3M            int     `json:"FG3M"`
	FG3A            int     `json:"FG3A"`
	FG3_PCT         float64 `json:"FG3_PCT"`
	FTM             int     `json:"FTM"`
	FTA             int     `json:"FTA"`
	FT_PCT          float64 `json:"FT_PCT"`
	REB             float64 `json:"REB"`
	AST             float64 `json:"AST"`
	TOV             float64 `json:"TOV"`
	STL             float64 `json:"STL"`
	BLK             float64 `json:"BLK"`
	PTS             float64 `json:"PTS"`
	NBA_FANTASY_PTS float64 `json:"NBA_FANTASY_PTS"`
}

// PlayerFantasyProfileResponse contains the response data from the PlayerFantasyProfile endpoint
type PlayerFantasyProfileResponse struct {
	LastNGames []PlayerFantasyProfileLastNGames
}

// GetPlayerFantasyProfile retrieves data from the playerfantasyprofile endpoint
func GetPlayerFantasyProfile(ctx context.Context, client *stats.Client, req PlayerFantasyProfileRequest) (*models.Response[*PlayerFantasyProfileResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playerfantasyprofile", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerFantasyProfileResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LastNGames = make([]PlayerFantasyProfileLastNGames, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 21 {
				item := PlayerFantasyProfileLastNGames{
					PLAYER_ID:       toInt(row[0]),
					PLAYER_NAME:     toString(row[1]),
					LAST_N_GAMES:    toFloat(row[2]),
					GP:              toInt(row[3]),
					MIN:             toFloat(row[4]),
					FGM:             toInt(row[5]),
					FGA:             toInt(row[6]),
					FG_PCT:          toFloat(row[7]),
					FG3M:            toInt(row[8]),
					FG3A:            toInt(row[9]),
					FG3_PCT:         toFloat(row[10]),
					FTM:             toInt(row[11]),
					FTA:             toInt(row[12]),
					FT_PCT:          toFloat(row[13]),
					REB:             toFloat(row[14]),
					AST:             toFloat(row[15]),
					TOV:             toFloat(row[16]),
					STL:             toFloat(row[17]),
					BLK:             toFloat(row[18]),
					PTS:             toFloat(row[19]),
					NBA_FANTASY_PTS: toFloat(row[20]),
				}
				response.LastNGames = append(response.LastNGames, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
