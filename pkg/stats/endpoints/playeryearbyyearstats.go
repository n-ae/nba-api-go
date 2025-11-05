package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayerYearByYearStatsRequest contains parameters for the PlayerYearByYearStats endpoint
type PlayerYearByYearStatsRequest struct {
	PlayerID string
	PerMode  *parameters.PerMode
	LeagueID *parameters.LeagueID
}

// PlayerYearByYearStatsPlayerStats represents the PlayerStats result set for PlayerYearByYearStats
type PlayerYearByYearStatsPlayerStats struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	SEASON_ID         string  `json:"SEASON_ID"`
	LEAGUE_ID         string  `json:"LEAGUE_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	PLAYER_AGE        int     `json:"PLAYER_AGE"`
	GP                int     `json:"GP"`
	GS                int     `json:"GS"`
	MIN               float64 `json:"MIN"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	FG3M              int     `json:"FG3M"`
	FG3A              int     `json:"FG3A"`
	FG3_PCT           float64 `json:"FG3_PCT"`
	FTM               int     `json:"FTM"`
	FTA               int     `json:"FTA"`
	FT_PCT            float64 `json:"FT_PCT"`
	OREB              float64 `json:"OREB"`
	DREB              float64 `json:"DREB"`
	REB               float64 `json:"REB"`
	AST               float64 `json:"AST"`
	STL               float64 `json:"STL"`
	BLK               float64 `json:"BLK"`
	TOV               float64 `json:"TOV"`
	PF                float64 `json:"PF"`
	PTS               float64 `json:"PTS"`
}

// PlayerYearByYearStatsResponse contains the response data from the PlayerYearByYearStats endpoint
type PlayerYearByYearStatsResponse struct {
	PlayerStats []PlayerYearByYearStatsPlayerStats
}

// GetPlayerYearByYearStats retrieves data from the playeryearbyyearstats endpoint
func GetPlayerYearByYearStats(ctx context.Context, client *stats.Client, req PlayerYearByYearStatsRequest) (*models.Response[*PlayerYearByYearStatsResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))
	if req.PerMode != nil {
		params.Set("PerMode", string(*req.PerMode))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "playeryearbyyearstats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerYearByYearStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerStats = make([]PlayerYearByYearStatsPlayerStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 27 {
				item := PlayerYearByYearStatsPlayerStats{
					PLAYER_ID:         toInt(row[0]),
					SEASON_ID:         toString(row[1]),
					LEAGUE_ID:         toString(row[2]),
					TEAM_ID:           toInt(row[3]),
					TEAM_ABBREVIATION: toString(row[4]),
					PLAYER_AGE:        toInt(row[5]),
					GP:                toInt(row[6]),
					GS:                toInt(row[7]),
					MIN:               toFloat(row[8]),
					FGM:               toInt(row[9]),
					FGA:               toInt(row[10]),
					FG_PCT:            toFloat(row[11]),
					FG3M:              toInt(row[12]),
					FG3A:              toInt(row[13]),
					FG3_PCT:           toFloat(row[14]),
					FTM:               toInt(row[15]),
					FTA:               toInt(row[16]),
					FT_PCT:            toFloat(row[17]),
					OREB:              toFloat(row[18]),
					DREB:              toFloat(row[19]),
					REB:               toFloat(row[20]),
					AST:               toFloat(row[21]),
					STL:               toFloat(row[22]),
					BLK:               toFloat(row[23]),
					TOV:               toFloat(row[24]),
					PF:                toFloat(row[25]),
					PTS:               toFloat(row[26]),
				}
				response.PlayerStats = append(response.PlayerStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
