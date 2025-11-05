package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueDashPlayerBioStatsRequest contains parameters for the LeagueDashPlayerBioStats endpoint
type LeagueDashPlayerBioStatsRequest struct {
	Season     *parameters.Season
	SeasonType *parameters.SeasonType
	PerMode    *parameters.PerMode
	LeagueID   *parameters.LeagueID
}

// LeagueDashPlayerBioStatsLeagueDashPlayerBioStats represents the LeagueDashPlayerBioStats result set for LeagueDashPlayerBioStats
type LeagueDashPlayerBioStatsLeagueDashPlayerBioStats struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	AGE               int     `json:"AGE"`
	PLAYER_HEIGHT     string  `json:"PLAYER_HEIGHT"`
	PLAYER_WEIGHT     string  `json:"PLAYER_WEIGHT"`
	COLLEGE           string  `json:"COLLEGE"`
	COUNTRY           string  `json:"COUNTRY"`
	DRAFT_YEAR        string  `json:"DRAFT_YEAR"`
	DRAFT_ROUND       string  `json:"DRAFT_ROUND"`
	DRAFT_NUMBER      string  `json:"DRAFT_NUMBER"`
	GP                int     `json:"GP"`
	PTS               float64 `json:"PTS"`
	REB               float64 `json:"REB"`
	AST               float64 `json:"AST"`
	NET_RATING        string  `json:"NET_RATING"`
	OREB_PCT          float64 `json:"OREB_PCT"`
	DREB_PCT          float64 `json:"DREB_PCT"`
	USG_PCT           float64 `json:"USG_PCT"`
	TS_PCT            float64 `json:"TS_PCT"`
	AST_PCT           float64 `json:"AST_PCT"`
}

// LeagueDashPlayerBioStatsResponse contains the response data from the LeagueDashPlayerBioStats endpoint
type LeagueDashPlayerBioStatsResponse struct {
	LeagueDashPlayerBioStats []LeagueDashPlayerBioStatsLeagueDashPlayerBioStats
}

// GetLeagueDashPlayerBioStats retrieves data from the leaguedashplayerbiostats endpoint
func GetLeagueDashPlayerBioStats(ctx context.Context, client *stats.Client, req LeagueDashPlayerBioStatsRequest) (*models.Response[*LeagueDashPlayerBioStatsResponse], error) {
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
	if err := client.GetJSON(ctx, "leaguedashplayerbiostats", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueDashPlayerBioStatsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueDashPlayerBioStats = make([]LeagueDashPlayerBioStatsLeagueDashPlayerBioStats, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 22 {
				item := LeagueDashPlayerBioStatsLeagueDashPlayerBioStats{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					AGE:               toInt(row[4]),
					PLAYER_HEIGHT:     toString(row[5]),
					PLAYER_WEIGHT:     toString(row[6]),
					COLLEGE:           toString(row[7]),
					COUNTRY:           toString(row[8]),
					DRAFT_YEAR:        toString(row[9]),
					DRAFT_ROUND:       toString(row[10]),
					DRAFT_NUMBER:      toString(row[11]),
					GP:                toInt(row[12]),
					PTS:               toFloat(row[13]),
					REB:               toFloat(row[14]),
					AST:               toFloat(row[15]),
					NET_RATING:        toString(row[16]),
					OREB_PCT:          toFloat(row[17]),
					DREB_PCT:          toFloat(row[18]),
					USG_PCT:           toFloat(row[19]),
					TS_PCT:            toFloat(row[20]),
					AST_PCT:           toFloat(row[21]),
				}
				response.LeagueDashPlayerBioStats = append(response.LeagueDashPlayerBioStats, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
