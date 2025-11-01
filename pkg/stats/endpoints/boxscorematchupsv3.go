package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// BoxScoreMatchupsV3Request contains parameters for the BoxScoreMatchupsV3 endpoint
type BoxScoreMatchupsV3Request struct {
	GameID      string
	StartPeriod *string
	EndPeriod   *string
}

// BoxScoreMatchupsV3HomeTeamPlayerMatchups represents the HomeTeamPlayerMatchups result set for BoxScoreMatchupsV3
type BoxScoreMatchupsV3HomeTeamPlayerMatchups struct {
	GAME_ID              string  `json:"GAME_ID"`
	PERSON_ID            string  `json:"PERSON_ID"`
	PLAYER_NAME          string  `json:"PLAYER_NAME"`
	TEAM_ID              int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION    string  `json:"TEAM_ABBREVIATION"`
	MATCHUP_MIN_PTS      string  `json:"MATCHUP_MIN_PTS"`
	PARTIAL_POSS         string  `json:"PARTIAL_POSS"`
	PLAYER_PTS           float64 `json:"PLAYER_PTS"`
	TEAM_PTS             float64 `json:"TEAM_PTS"`
	MATCHUP_AST          string  `json:"MATCHUP_AST"`
	MATCHUP_TOV          string  `json:"MATCHUP_TOV"`
	MATCHUP_BLK          string  `json:"MATCHUP_BLK"`
	MATCHUP_FGM          string  `json:"MATCHUP_FGM"`
	MATCHUP_FGA          string  `json:"MATCHUP_FGA"`
	MATCHUP_FG_PCT       float64 `json:"MATCHUP_FG_PCT"`
	MATCHUP_FG3M         string  `json:"MATCHUP_FG3M"`
	MATCHUP_FG3A         string  `json:"MATCHUP_FG3A"`
	MATCHUP_FG3_PCT      float64 `json:"MATCHUP_FG3_PCT"`
	HELP_BLK             float64 `json:"HELP_BLK"`
	HELP_FGM             int     `json:"HELP_FGM"`
	HELP_FGA             int     `json:"HELP_FGA"`
	HELP_FG_PCT          float64 `json:"HELP_FG_PCT"`
	SHOOTER_PLAYER_ID    int     `json:"SHOOTER_PLAYER_ID"`
	SHOOTER_PLAYER_NAME  string  `json:"SHOOTER_PLAYER_NAME"`
	DEFENDER_PLAYER_ID   int     `json:"DEFENDER_PLAYER_ID"`
	DEFENDER_PLAYER_NAME string  `json:"DEFENDER_PLAYER_NAME"`
	SFL                  string  `json:"SFL"`
}

// BoxScoreMatchupsV3AwayTeamPlayerMatchups represents the AwayTeamPlayerMatchups result set for BoxScoreMatchupsV3
type BoxScoreMatchupsV3AwayTeamPlayerMatchups struct {
	GAME_ID              string  `json:"GAME_ID"`
	PERSON_ID            string  `json:"PERSON_ID"`
	PLAYER_NAME          string  `json:"PLAYER_NAME"`
	TEAM_ID              int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION    string  `json:"TEAM_ABBREVIATION"`
	MATCHUP_MIN_PTS      string  `json:"MATCHUP_MIN_PTS"`
	PARTIAL_POSS         string  `json:"PARTIAL_POSS"`
	PLAYER_PTS           float64 `json:"PLAYER_PTS"`
	TEAM_PTS             float64 `json:"TEAM_PTS"`
	MATCHUP_AST          string  `json:"MATCHUP_AST"`
	MATCHUP_TOV          string  `json:"MATCHUP_TOV"`
	MATCHUP_BLK          string  `json:"MATCHUP_BLK"`
	MATCHUP_FGM          string  `json:"MATCHUP_FGM"`
	MATCHUP_FGA          string  `json:"MATCHUP_FGA"`
	MATCHUP_FG_PCT       float64 `json:"MATCHUP_FG_PCT"`
	MATCHUP_FG3M         string  `json:"MATCHUP_FG3M"`
	MATCHUP_FG3A         string  `json:"MATCHUP_FG3A"`
	MATCHUP_FG3_PCT      float64 `json:"MATCHUP_FG3_PCT"`
	HELP_BLK             float64 `json:"HELP_BLK"`
	HELP_FGM             int     `json:"HELP_FGM"`
	HELP_FGA             int     `json:"HELP_FGA"`
	HELP_FG_PCT          float64 `json:"HELP_FG_PCT"`
	SHOOTER_PLAYER_ID    int     `json:"SHOOTER_PLAYER_ID"`
	SHOOTER_PLAYER_NAME  string  `json:"SHOOTER_PLAYER_NAME"`
	DEFENDER_PLAYER_ID   int     `json:"DEFENDER_PLAYER_ID"`
	DEFENDER_PLAYER_NAME string  `json:"DEFENDER_PLAYER_NAME"`
	SFL                  string  `json:"SFL"`
}

// BoxScoreMatchupsV3Response contains the response data from the BoxScoreMatchupsV3 endpoint
type BoxScoreMatchupsV3Response struct {
	HomeTeamPlayerMatchups []BoxScoreMatchupsV3HomeTeamPlayerMatchups
	AwayTeamPlayerMatchups []BoxScoreMatchupsV3AwayTeamPlayerMatchups
}

// GetBoxScoreMatchupsV3 retrieves data from the boxscorematchupsv3 endpoint
func GetBoxScoreMatchupsV3(ctx context.Context, client *stats.Client, req BoxScoreMatchupsV3Request) (*models.Response[*BoxScoreMatchupsV3Response], error) {
	params := url.Values{}
	if req.GameID == "" {
		return nil, fmt.Errorf("GameID is required")
	}
	params.Set("GameID", string(req.GameID))
	if req.StartPeriod != nil {
		params.Set("StartPeriod", string(*req.StartPeriod))
	}
	if req.EndPeriod != nil {
		params.Set("EndPeriod", string(*req.EndPeriod))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/boxscorematchupsv3", params, &rawResp); err != nil {
		return nil, err
	}

	response := &BoxScoreMatchupsV3Response{}
	if len(rawResp.ResultSets) > 0 {
		response.HomeTeamPlayerMatchups = make([]BoxScoreMatchupsV3HomeTeamPlayerMatchups, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 27 {
				item := BoxScoreMatchupsV3HomeTeamPlayerMatchups{
					GAME_ID:              toString(row[0]),
					PERSON_ID:            toString(row[1]),
					PLAYER_NAME:          toString(row[2]),
					TEAM_ID:              toInt(row[3]),
					TEAM_ABBREVIATION:    toString(row[4]),
					MATCHUP_MIN_PTS:      toString(row[5]),
					PARTIAL_POSS:         toString(row[6]),
					PLAYER_PTS:           toFloat(row[7]),
					TEAM_PTS:             toFloat(row[8]),
					MATCHUP_AST:          toString(row[9]),
					MATCHUP_TOV:          toString(row[10]),
					MATCHUP_BLK:          toString(row[11]),
					MATCHUP_FGM:          toString(row[12]),
					MATCHUP_FGA:          toString(row[13]),
					MATCHUP_FG_PCT:       toFloat(row[14]),
					MATCHUP_FG3M:         toString(row[15]),
					MATCHUP_FG3A:         toString(row[16]),
					MATCHUP_FG3_PCT:      toFloat(row[17]),
					HELP_BLK:             toFloat(row[18]),
					HELP_FGM:             toInt(row[19]),
					HELP_FGA:             toInt(row[20]),
					HELP_FG_PCT:          toFloat(row[21]),
					SHOOTER_PLAYER_ID:    toInt(row[22]),
					SHOOTER_PLAYER_NAME:  toString(row[23]),
					DEFENDER_PLAYER_ID:   toInt(row[24]),
					DEFENDER_PLAYER_NAME: toString(row[25]),
					SFL:                  toString(row[26]),
				}
				response.HomeTeamPlayerMatchups = append(response.HomeTeamPlayerMatchups, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.AwayTeamPlayerMatchups = make([]BoxScoreMatchupsV3AwayTeamPlayerMatchups, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 27 {
				item := BoxScoreMatchupsV3AwayTeamPlayerMatchups{
					GAME_ID:              toString(row[0]),
					PERSON_ID:            toString(row[1]),
					PLAYER_NAME:          toString(row[2]),
					TEAM_ID:              toInt(row[3]),
					TEAM_ABBREVIATION:    toString(row[4]),
					MATCHUP_MIN_PTS:      toString(row[5]),
					PARTIAL_POSS:         toString(row[6]),
					PLAYER_PTS:           toFloat(row[7]),
					TEAM_PTS:             toFloat(row[8]),
					MATCHUP_AST:          toString(row[9]),
					MATCHUP_TOV:          toString(row[10]),
					MATCHUP_BLK:          toString(row[11]),
					MATCHUP_FGM:          toString(row[12]),
					MATCHUP_FGA:          toString(row[13]),
					MATCHUP_FG_PCT:       toFloat(row[14]),
					MATCHUP_FG3M:         toString(row[15]),
					MATCHUP_FG3A:         toString(row[16]),
					MATCHUP_FG3_PCT:      toFloat(row[17]),
					HELP_BLK:             toFloat(row[18]),
					HELP_FGM:             toInt(row[19]),
					HELP_FGA:             toInt(row[20]),
					HELP_FG_PCT:          toFloat(row[21]),
					SHOOTER_PLAYER_ID:    toInt(row[22]),
					SHOOTER_PLAYER_NAME:  toString(row[23]),
					DEFENDER_PLAYER_ID:   toInt(row[24]),
					DEFENDER_PLAYER_NAME: toString(row[25]),
					SFL:                  toString(row[26]),
				}
				response.AwayTeamPlayerMatchups = append(response.AwayTeamPlayerMatchups, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
