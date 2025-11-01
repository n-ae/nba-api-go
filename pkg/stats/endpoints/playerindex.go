package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// PlayerIndexRequest contains parameters for the PlayerIndex endpoint
type PlayerIndexRequest struct {
	LeagueID   *parameters.LeagueID
	Season     *parameters.Season
	TeamID     *string
	Historical *string
}

// PlayerIndexPlayerIndex represents the PlayerIndex result set for PlayerIndex
type PlayerIndexPlayerIndex struct {
	PERSON_ID         string `json:"PERSON_ID"`
	PLAYER_LAST_NAME  string `json:"PLAYER_LAST_NAME"`
	PLAYER_FIRST_NAME string `json:"PLAYER_FIRST_NAME"`
	PLAYER_SLUG       string `json:"PLAYER_SLUG"`
	TEAM_ID           int    `json:"TEAM_ID"`
	TEAM_SLUG         string `json:"TEAM_SLUG"`
	IS_DEFUNCT        string `json:"IS_DEFUNCT"`
	TEAM_CITY         string `json:"TEAM_CITY"`
	TEAM_NAME         string `json:"TEAM_NAME"`
	TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
	JERSEY_NUMBER     string `json:"JERSEY_NUMBER"`
	POSITION          string `json:"POSITION"`
	HEIGHT            string `json:"HEIGHT"`
	WEIGHT            string `json:"WEIGHT"`
	COLLEGE           string `json:"COLLEGE"`
	COUNTRY           string `json:"COUNTRY"`
	DRAFT_YEAR        string `json:"DRAFT_YEAR"`
	DRAFT_ROUND       string `json:"DRAFT_ROUND"`
	DRAFT_NUMBER      string `json:"DRAFT_NUMBER"`
	ROSTER_STATUS     string `json:"ROSTER_STATUS"`
	FROM_YEAR         string `json:"FROM_YEAR"`
	TO_YEAR           string `json:"TO_YEAR"`
}

// PlayerIndexResponse contains the response data from the PlayerIndex endpoint
type PlayerIndexResponse struct {
	PlayerIndex []PlayerIndexPlayerIndex
}

// GetPlayerIndex retrieves data from the playerindex endpoint
func GetPlayerIndex(ctx context.Context, client *stats.Client, req PlayerIndexRequest) (*models.Response[*PlayerIndexResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.TeamID != nil {
		params.Set("TeamID", string(*req.TeamID))
	}
	if req.Historical != nil {
		params.Set("Historical", string(*req.Historical))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playerindex", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerIndexResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerIndex = make([]PlayerIndexPlayerIndex, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 22 {
				item := PlayerIndexPlayerIndex{
					PERSON_ID:         toString(row[0]),
					PLAYER_LAST_NAME:  toString(row[1]),
					PLAYER_FIRST_NAME: toString(row[2]),
					PLAYER_SLUG:       toString(row[3]),
					TEAM_ID:           toInt(row[4]),
					TEAM_SLUG:         toString(row[5]),
					IS_DEFUNCT:        toString(row[6]),
					TEAM_CITY:         toString(row[7]),
					TEAM_NAME:         toString(row[8]),
					TEAM_ABBREVIATION: toString(row[9]),
					JERSEY_NUMBER:     toString(row[10]),
					POSITION:          toString(row[11]),
					HEIGHT:            toString(row[12]),
					WEIGHT:            toString(row[13]),
					COLLEGE:           toString(row[14]),
					COUNTRY:           toString(row[15]),
					DRAFT_YEAR:        toString(row[16]),
					DRAFT_ROUND:       toString(row[17]),
					DRAFT_NUMBER:      toString(row[18]),
					ROSTER_STATUS:     toString(row[19]),
					FROM_YEAR:         toString(row[20]),
					TO_YEAR:           toString(row[21]),
				}
				response.PlayerIndex = append(response.PlayerIndex, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
