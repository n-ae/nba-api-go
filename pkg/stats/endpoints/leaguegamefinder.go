package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// LeagueGameFinderRequest contains parameters for the LeagueGameFinder endpoint
type LeagueGameFinderRequest struct {
	LeagueID     *parameters.LeagueID
	Season       *parameters.Season
	SeasonType   *parameters.SeasonType
	PlayerOrTeam *string
	PlayerID     *string
	TeamID       *string
	VsTeamID     *string
	Outcome      *string
	Location     *string
	DateFrom     *string
	DateTo       *string
	VsConference *string
	VsDivision   *string
	GameSegment  *string
	Period       *string
	LastNGames   *string
	PORound      *string
}

// LeagueGameFinderLeagueGameFinderResults represents the LeagueGameFinderResults result set for LeagueGameFinder
type LeagueGameFinderLeagueGameFinderResults struct {
	SEASON_ID         string  `json:"SEASON_ID"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	TEAM_NAME         string  `json:"TEAM_NAME"`
	GAME_ID           string  `json:"GAME_ID"`
	GAME_DATE         string  `json:"GAME_DATE"`
	MATCHUP           string  `json:"MATCHUP"`
	WL                string  `json:"WL"`
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
	OREB              int     `json:"OREB"`
	DREB              int     `json:"DREB"`
	REB               int     `json:"REB"`
	AST               int     `json:"AST"`
	STL               int     `json:"STL"`
	BLK               int     `json:"BLK"`
	TOV               int     `json:"TOV"`
	PF                int     `json:"PF"`
	PTS               int     `json:"PTS"`
	PLUS_MINUS        float64 `json:"PLUS_MINUS"`
}

// LeagueGameFinderResponse contains the response data from the LeagueGameFinder endpoint
type LeagueGameFinderResponse struct {
	LeagueGameFinderResults []LeagueGameFinderLeagueGameFinderResults
}

// GetLeagueGameFinder retrieves data from the leaguegamefinder endpoint
func GetLeagueGameFinder(ctx context.Context, client *stats.Client, req LeagueGameFinderRequest) (*models.Response[*LeagueGameFinderResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.PlayerOrTeam != nil {
		params.Set("PlayerOrTeam", string(*req.PlayerOrTeam))
	}
	if req.PlayerID != nil {
		params.Set("PlayerID", string(*req.PlayerID))
	}
	if req.TeamID != nil {
		params.Set("TeamID", string(*req.TeamID))
	}
	if req.VsTeamID != nil {
		params.Set("VsTeamID", string(*req.VsTeamID))
	}
	if req.Outcome != nil {
		params.Set("Outcome", string(*req.Outcome))
	}
	if req.Location != nil {
		params.Set("Location", string(*req.Location))
	}
	if req.DateFrom != nil {
		params.Set("DateFrom", string(*req.DateFrom))
	}
	if req.DateTo != nil {
		params.Set("DateTo", string(*req.DateTo))
	}
	if req.VsConference != nil {
		params.Set("VsConference", string(*req.VsConference))
	}
	if req.VsDivision != nil {
		params.Set("VsDivision", string(*req.VsDivision))
	}
	if req.GameSegment != nil {
		params.Set("GameSegment", string(*req.GameSegment))
	}
	if req.Period != nil {
		params.Set("Period", string(*req.Period))
	}
	if req.LastNGames != nil {
		params.Set("LastNGames", string(*req.LastNGames))
	}
	if req.PORound != nil {
		params.Set("PORound", string(*req.PORound))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "leaguegamefinder", params, &rawResp); err != nil {
		return nil, err
	}

	response := &LeagueGameFinderResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.LeagueGameFinderResults = make([]LeagueGameFinderLeagueGameFinderResults, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 28 {
				item := LeagueGameFinderLeagueGameFinderResults{
					SEASON_ID:         toString(row[0]),
					TEAM_ID:           toInt(row[1]),
					TEAM_ABBREVIATION: toString(row[2]),
					TEAM_NAME:         toString(row[3]),
					GAME_ID:           toString(row[4]),
					GAME_DATE:         toString(row[5]),
					MATCHUP:           toString(row[6]),
					WL:                toString(row[7]),
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
					OREB:              toInt(row[18]),
					DREB:              toInt(row[19]),
					REB:               toInt(row[20]),
					AST:               toInt(row[21]),
					STL:               toInt(row[22]),
					BLK:               toInt(row[23]),
					TOV:               toInt(row[24]),
					PF:                toInt(row[25]),
					PTS:               toInt(row[26]),
					PLUS_MINUS:        toFloat(row[27]),
				}
				response.LeagueGameFinderResults = append(response.LeagueGameFinderResults, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
