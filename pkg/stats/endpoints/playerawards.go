package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
)

// PlayerAwardsRequest contains parameters for the PlayerAwards endpoint
type PlayerAwardsRequest struct {
	PlayerID string
}

// PlayerAwardsPlayerAwards represents the PlayerAwards result set for PlayerAwards
type PlayerAwardsPlayerAwards struct {
	PERSON_ID           string `json:"PERSON_ID"`
	FIRST_NAME          string `json:"FIRST_NAME"`
	LAST_NAME           string `json:"LAST_NAME"`
	TEAM                string `json:"TEAM"`
	DESCRIPTION         string `json:"DESCRIPTION"`
	ALL_NBA_TEAM_NUMBER string `json:"ALL_NBA_TEAM_NUMBER"`
	SEASON              string `json:"SEASON"`
	MONTH               string `json:"MONTH"`
	WEEK                string `json:"WEEK"`
	CONFERENCE          string `json:"CONFERENCE"`
	TYPE                string `json:"TYPE"`
	SUBTYPE1            string `json:"SUBTYPE1"`
	SUBTYPE2            string `json:"SUBTYPE2"`
	SUBTYPE3            string `json:"SUBTYPE3"`
}

// PlayerAwardsResponse contains the response data from the PlayerAwards endpoint
type PlayerAwardsResponse struct {
	PlayerAwards []PlayerAwardsPlayerAwards
}

// GetPlayerAwards retrieves data from the playerawards endpoint
func GetPlayerAwards(ctx context.Context, client *stats.Client, req PlayerAwardsRequest) (*models.Response[*PlayerAwardsResponse], error) {
	params := url.Values{}
	if req.PlayerID == "" {
		return nil, fmt.Errorf("PlayerID is required")
	}
	params.Set("PlayerID", string(req.PlayerID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/playerawards", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayerAwardsResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.PlayerAwards = make([]PlayerAwardsPlayerAwards, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := PlayerAwardsPlayerAwards{
					PERSON_ID:           toString(row[0]),
					FIRST_NAME:          toString(row[1]),
					LAST_NAME:           toString(row[2]),
					TEAM:                toString(row[3]),
					DESCRIPTION:         toString(row[4]),
					ALL_NBA_TEAM_NUMBER: toString(row[5]),
					SEASON:              toString(row[6]),
					MONTH:               toString(row[7]),
					WEEK:                toString(row[8]),
					CONFERENCE:          toString(row[9]),
					TYPE:                toString(row[10]),
					SUBTYPE1:            toString(row[11]),
					SUBTYPE2:            toString(row[12]),
					SUBTYPE3:            toString(row[13]),
				}
				response.PlayerAwards = append(response.PlayerAwards, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
