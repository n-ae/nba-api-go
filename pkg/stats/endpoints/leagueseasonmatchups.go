package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueSeasonMatchupsRequest contains parameters for the LeagueSeasonMatchups endpoint
type LeagueSeasonMatchupsRequest struct {
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	PerMode     *parameters.PerMode
	LeagueID    *parameters.LeagueID
	DefPlayerID *string
	OffPlayerID *string
}

// LeagueSeasonMatchupsSeasonMatchups represents the SeasonMatchups result set for LeagueSeasonMatchups
type LeagueSeasonMatchupsSeasonMatchups struct {
	SEASON_ID       string  `json:"SEASON_ID"`
	OFF_PLAYER_ID   int     `json:"OFF_PLAYER_ID"`
	OFF_PLAYER_NAME string  `json:"OFF_PLAYER_NAME"`
	DEF_PLAYER_ID   int     `json:"DEF_PLAYER_ID"`
	DEF_PLAYER_NAME string  `json:"DEF_PLAYER_NAME"`
	GP              int     `json:"GP"`
	MATCHUP_MIN     string  `json:"MATCHUP_MIN"`
	PARTIAL_POSS    string  `json:"PARTIAL_POSS"`
	PLAYER_PTS      float64 `json:"PLAYER_PTS"`
	TEAM_PTS        float64 `json:"TEAM_PTS"`
	MATCHUP_AST     string  `json:"MATCHUP_AST"`
	MATCHUP_TOV     string  `json:"MATCHUP_TOV"`
	MATCHUP_BLK     string  `json:"MATCHUP_BLK"`
	MATCHUP_FGM     string  `json:"MATCHUP_FGM"`
	MATCHUP_FGA     string  `json:"MATCHUP_FGA"`
	MATCHUP_FG_PCT  float64 `json:"MATCHUP_FG_PCT"`
	SFL             string  `json:"SFL"`
}

// LeagueSeasonMatchupsResponse contains the response data from the LeagueSeasonMatchups endpoint
type LeagueSeasonMatchupsResponse struct {
	SeasonMatchups []LeagueSeasonMatchupsSeasonMatchups
}

// GetLeagueSeasonMatchups retrieves data from the leagueseasonmatchups endpoint
func GetLeagueSeasonMatchups(ctx context.Context, client *stats.Client, req LeagueSeasonMatchupsRequest) (*models.Response[*LeagueSeasonMatchupsResponse], error) {
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
	if req.DefPlayerID != nil {
		params.Set("DefPlayerID", string(*req.DefPlayerID))
	}
	if req.OffPlayerID != nil {
		params.Set("OffPlayerID", string(*req.OffPlayerID))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/leagueseasonmatchups", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueSeasonMatchupsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.SeasonMatchups = make([]LeagueSeasonMatchupsSeasonMatchups, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 17 {
				item := LeagueSeasonMatchupsSeasonMatchups{
					SEASON_ID:       toString(row[0]),
					OFF_PLAYER_ID:   toInt(row[1]),
					OFF_PLAYER_NAME: toString(row[2]),
					DEF_PLAYER_ID:   toInt(row[3]),
					DEF_PLAYER_NAME: toString(row[4]),
					GP:              toInt(row[5]),
					MATCHUP_MIN:     toString(row[6]),
					PARTIAL_POSS:    toString(row[7]),
					PLAYER_PTS:      toFloat(row[8]),
					TEAM_PTS:        toFloat(row[9]),
					MATCHUP_AST:     toString(row[10]),
					MATCHUP_TOV:     toString(row[11]),
					MATCHUP_BLK:     toString(row[12]),
					MATCHUP_FGM:     toString(row[13]),
					MATCHUP_FGA:     toString(row[14]),
					MATCHUP_FG_PCT:  toFloat(row[15]),
					SFL:             toString(row[16]),
				}
				response.SeasonMatchups = append(response.SeasonMatchups, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
