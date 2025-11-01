package endpoints

import (
	"context"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// SynergyPlayTypesRequest contains parameters for the SynergyPlayTypes endpoint
type SynergyPlayTypesRequest struct {
	Season       *parameters.Season
	SeasonType   *parameters.SeasonType
	PerMode      *parameters.PerMode
	LeagueID     *parameters.LeagueID
	PlayerOrTeam *string
	PlayType     *string
}

// SynergyPlayTypesSynergyPlayType represents the SynergyPlayType result set for SynergyPlayTypes
type SynergyPlayTypesSynergyPlayType struct {
	PLAYER_ID         int     `json:"PLAYER_ID"`
	PLAYER_NAME       string  `json:"PLAYER_NAME"`
	TEAM_ID           int     `json:"TEAM_ID"`
	TEAM_ABBREVIATION string  `json:"TEAM_ABBREVIATION"`
	PLAY_TYPE         string  `json:"PLAY_TYPE"`
	TYPE_GROUPING     string  `json:"TYPE_GROUPING"`
	PERCENTILE        string  `json:"PERCENTILE"`
	GP                int     `json:"GP"`
	POSS              string  `json:"POSS"`
	TIME              string  `json:"TIME"`
	PTS               float64 `json:"PTS"`
	FGM               int     `json:"FGM"`
	FGA               int     `json:"FGA"`
	FG_PCT            float64 `json:"FG_PCT"`
	AEFG_PCT          float64 `json:"AEFG_PCT"`
	FT_PCT            float64 `json:"FT_PCT"`
	TOV               float64 `json:"TOV"`
	SF                string  `json:"SF"`
	PLUSONE           string  `json:"PLUSONE"`
	SCORE             string  `json:"SCORE"`
	EFG_PCT           float64 `json:"EFG_PCT"`
	PPP               string  `json:"PPP"`
	POSS_PCT          float64 `json:"POSS_PCT"`
}

// SynergyPlayTypesResponse contains the response data from the SynergyPlayTypes endpoint
type SynergyPlayTypesResponse struct {
	SynergyPlayType []SynergyPlayTypesSynergyPlayType
}

// GetSynergyPlayTypes retrieves data from the synergyplaytypes endpoint
func GetSynergyPlayTypes(ctx context.Context, client *stats.Client, req SynergyPlayTypesRequest) (*models.Response[*SynergyPlayTypesResponse], error) {
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
	if req.PlayerOrTeam != nil {
		params.Set("PlayerOrTeam", string(*req.PlayerOrTeam))
	}
	if req.PlayType != nil {
		params.Set("PlayType", string(*req.PlayType))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/synergyplaytypes", params, &rawResp); err != nil {
		return nil, err
	}

	response := &SynergyPlayTypesResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.SynergyPlayType = make([]SynergyPlayTypesSynergyPlayType, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 23 {
				item := SynergyPlayTypesSynergyPlayType{
					PLAYER_ID:         toInt(row[0]),
					PLAYER_NAME:       toString(row[1]),
					TEAM_ID:           toInt(row[2]),
					TEAM_ABBREVIATION: toString(row[3]),
					PLAY_TYPE:         toString(row[4]),
					TYPE_GROUPING:     toString(row[5]),
					PERCENTILE:        toString(row[6]),
					GP:                toInt(row[7]),
					POSS:              toString(row[8]),
					TIME:              toString(row[9]),
					PTS:               toFloat(row[10]),
					FGM:               toInt(row[11]),
					FGA:               toInt(row[12]),
					FG_PCT:            toFloat(row[13]),
					AEFG_PCT:          toFloat(row[14]),
					FT_PCT:            toFloat(row[15]),
					TOV:               toFloat(row[16]),
					SF:                toString(row[17]),
					PLUSONE:           toString(row[18]),
					SCORE:             toString(row[19]),
					EFG_PCT:           toFloat(row[20]),
					PPP:               toString(row[21]),
					POSS_PCT:          toFloat(row[22]),
				}
				response.SynergyPlayType = append(response.SynergyPlayType, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
