package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// MatchupRollupRequest contains parameters for the MatchupRollup endpoint
type MatchupRollupRequest struct {
	Season      *parameters.Season
	SeasonType  *parameters.SeasonType
	PerMode     *parameters.PerMode
	LeagueID    *parameters.LeagueID
	DefPlayerID *string
	OffPlayerID *string
}

// MatchupRollupMatchupRollup represents the MatchupRollup result set for MatchupRollup
type MatchupRollupMatchupRollup struct {
	PERSON_ID         string  `json:"PERSON_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	GP                int     `json:"GP"`
	MATCHUP_MIN       string  `json:"MATCHUP_MIN"`
	PARTIAL_POSS      string  `json:"PARTIAL_POSS"`
	PLAYER_PTS        float64 `json:"PLAYER_PTS"`
	TEAM_PTS          float64 `json:"TEAM_PTS"`
	MATCHUP_AST       string  `json:"MATCHUP_AST"`
	MATCHUP_TOV       string  `json:"MATCHUP_TOV"`
	MATCHUP_BLK       string  `json:"MATCHUP_BLK"`
	MATCHUP_FGM       string  `json:"MATCHUP_FGM"`
	MATCHUP_FGA       string  `json:"MATCHUP_FGA"`
	MATCHUP_FG_PCT    float64 `json:"MATCHUP_FG_PCT"`
	SFL               string  `json:"SFL"`
}

// MatchupRollupResponse contains the response data from the MatchupRollup endpoint
type MatchupRollupResponse struct {
	MatchupRollup []MatchupRollupMatchupRollup
}

// GetMatchupRollup retrieves data from the matchuprollup endpoint
func GetMatchupRollup(ctx context.Context, client *stats.Client, req MatchupRollupRequest) (*models.Response[*MatchupRollupResponse], error) {
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
	if err := client.GetJSON(ctx, "/matchuprollup", params, &rawResp); err != nil {
		return nil, err
	}

	response := &MatchupRollupResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.MatchupRollup = make([]MatchupRollupMatchupRollup, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 15 {
				item := MatchupRollupMatchupRollup{
					PERSON_ID:         toString(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					GP:                toInt(row[3]),
					MATCHUP_MIN:       toString(row[4]),
					PARTIAL_POSS:      toString(row[5]),
					PLAYER_PTS:        toFloat(row[6]),
					TEAM_PTS:          toFloat(row[7]),
					MATCHUP_AST:       toString(row[8]),
					MATCHUP_TOV:       toString(row[9]),
					MATCHUP_BLK:       toString(row[10]),
					MATCHUP_FGM:       toString(row[11]),
					MATCHUP_FGA:       toString(row[12]),
					MATCHUP_FG_PCT:    toFloat(row[13]),
					SFL:               toString(row[14]),
				}
				response.MatchupRollup = append(response.MatchupRollup, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
