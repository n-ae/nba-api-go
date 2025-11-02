package endpoints

import (
	"context"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// HomepageLeadersRequest contains parameters for the HomepageLeaders endpoint
type HomepageLeadersRequest struct {
	Season       *parameters.Season
	SeasonType   *parameters.SeasonType
	LeagueID     *parameters.LeagueID
	PlayerOrTeam *string
	GameScope    *string
	PlayerScope  *string
	Stat         *string
}

// HomepageLeadersHomepageLeaders represents the HomepageLeaders result set for HomepageLeaders
type HomepageLeadersHomepageLeaders struct {
	PLAYER_ID int     `json:"PLAYER_ID"`
	RANK      int     `json:"RANK"`
	PLAYER    string  `json:"PLAYER"`
	TEAM_ID   int     `json:"TEAM_ID"`
	TEAM      string  `json:"TEAM"`
	GP        int     `json:"GP"`
	MIN       float64 `json:"MIN"`
	FGM       int     `json:"FGM"`
	FGA       int     `json:"FGA"`
	FG_PCT    float64 `json:"FG_PCT"`
	FG3M      int     `json:"FG3M"`
	FG3A      int     `json:"FG3A"`
	FG3_PCT   float64 `json:"FG3_PCT"`
	FTM       int     `json:"FTM"`
	FTA       int     `json:"FTA"`
	FT_PCT    float64 `json:"FT_PCT"`
	OREB      float64 `json:"OREB"`
	DREB      float64 `json:"DREB"`
	REB       float64 `json:"REB"`
	AST       float64 `json:"AST"`
	STL       float64 `json:"STL"`
	BLK       float64 `json:"BLK"`
	TOV       float64 `json:"TOV"`
	PTS       float64 `json:"PTS"`
	EFF       string  `json:"EFF"`
}

// HomepageLeadersResponse contains the response data from the HomepageLeaders endpoint
type HomepageLeadersResponse struct {
	HomepageLeaders []HomepageLeadersHomepageLeaders
}

// GetHomepageLeaders retrieves data from the homepageleaders endpoint
func GetHomepageLeaders(ctx context.Context, client *stats.Client, req HomepageLeadersRequest) (*models.Response[*HomepageLeadersResponse], error) {
	params := url.Values{}
	if req.Season != nil {
		params.Set("Season", string(*req.Season))
	}
	if req.SeasonType != nil {
		params.Set("SeasonType", string(*req.SeasonType))
	}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.PlayerOrTeam != nil {
		params.Set("PlayerOrTeam", string(*req.PlayerOrTeam))
	}
	if req.GameScope != nil {
		params.Set("GameScope", string(*req.GameScope))
	}
	if req.PlayerScope != nil {
		params.Set("PlayerScope", string(*req.PlayerScope))
	}
	if req.Stat != nil {
		params.Set("Stat", string(*req.Stat))
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/homepageleaders", params, &rawResp); err != nil {
		return nil, err
	}

	response := &HomepageLeadersResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.HomepageLeaders = make([]HomepageLeadersHomepageLeaders, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 25 {
				item := HomepageLeadersHomepageLeaders{
					PLAYER_ID: toInt(row[0]),
					RANK:      toInt(row[1]),
					PLAYER:    toString(row[2]),
					TEAM_ID:   toInt(row[3]),
					TEAM:      toString(row[4]),
					GP:        toInt(row[5]),
					MIN:       toFloat(row[6]),
					FGM:       toInt(row[7]),
					FGA:       toInt(row[8]),
					FG_PCT:    toFloat(row[9]),
					FG3M:      toInt(row[10]),
					FG3A:      toInt(row[11]),
					FG3_PCT:   toFloat(row[12]),
					FTM:       toInt(row[13]),
					FTA:       toInt(row[14]),
					FT_PCT:    toFloat(row[15]),
					OREB:      toFloat(row[16]),
					DREB:      toFloat(row[17]),
					REB:       toFloat(row[18]),
					AST:       toFloat(row[19]),
					STL:       toFloat(row[20]),
					BLK:       toFloat(row[21]),
					TOV:       toFloat(row[22]),
					PTS:       toFloat(row[23]),
					EFF:       toString(row[24]),
				}
				response.HomepageLeaders = append(response.HomepageLeaders, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
