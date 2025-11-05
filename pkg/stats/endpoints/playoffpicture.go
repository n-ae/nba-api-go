package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/n-ae/nba-api-go/pkg/models"
	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

// PlayoffPictureRequest contains parameters for the PlayoffPicture endpoint
type PlayoffPictureRequest struct {
	LeagueID *parameters.LeagueID
	SeasonID parameters.Season
}

// PlayoffPictureEastConfPlayoffPicture represents the EastConfPlayoffPicture result set for PlayoffPicture
type PlayoffPictureEastConfPlayoffPicture struct {
	TEAM_ID                  int     `json:"TEAM_ID"`
	LEAGUE_ID                string  `json:"LEAGUE_ID"`
	SEASON_ID                string  `json:"SEASON_ID"`
	CONFERENCE               string  `json:"CONFERENCE"`
	RANK                     int     `json:"RANK"`
	TEAM                     string  `json:"TEAM"`
	WINS                     string  `json:"WINS"`
	LOSSES                   string  `json:"LOSSES"`
	WIN_PCT                  float64 `json:"WIN_PCT"`
	GAMES_BACK               string  `json:"GAMES_BACK"`
	CLINCHED                 string  `json:"CLINCHED"`
	ELIMINATED_FROM_PLAYOFFS float64 `json:"ELIMINATED_FROM_PLAYOFFS"`
	CAN_WIN_CONF             string  `json:"CAN_WIN_CONF"`
	CAN_WIN_DIV              string  `json:"CAN_WIN_DIV"`
}

// PlayoffPictureWestConfPlayoffPicture represents the WestConfPlayoffPicture result set for PlayoffPicture
type PlayoffPictureWestConfPlayoffPicture struct {
	TEAM_ID                  int     `json:"TEAM_ID"`
	LEAGUE_ID                string  `json:"LEAGUE_ID"`
	SEASON_ID                string  `json:"SEASON_ID"`
	CONFERENCE               string  `json:"CONFERENCE"`
	RANK                     int     `json:"RANK"`
	TEAM                     string  `json:"TEAM"`
	WINS                     string  `json:"WINS"`
	LOSSES                   string  `json:"LOSSES"`
	WIN_PCT                  float64 `json:"WIN_PCT"`
	GAMES_BACK               string  `json:"GAMES_BACK"`
	CLINCHED                 string  `json:"CLINCHED"`
	ELIMINATED_FROM_PLAYOFFS float64 `json:"ELIMINATED_FROM_PLAYOFFS"`
	CAN_WIN_CONF             string  `json:"CAN_WIN_CONF"`
	CAN_WIN_DIV              string  `json:"CAN_WIN_DIV"`
}

// PlayoffPictureResponse contains the response data from the PlayoffPicture endpoint
type PlayoffPictureResponse struct {
	EastConfPlayoffPicture []PlayoffPictureEastConfPlayoffPicture
	WestConfPlayoffPicture []PlayoffPictureWestConfPlayoffPicture
}

// GetPlayoffPicture retrieves data from the playoffpicture endpoint
func GetPlayoffPicture(ctx context.Context, client *stats.Client, req PlayoffPictureRequest) (*models.Response[*PlayoffPictureResponse], error) {
	params := url.Values{}
	if req.LeagueID != nil {
		params.Set("LeagueID", string(*req.LeagueID))
	}
	if req.SeasonID == "" {
		return nil, fmt.Errorf("SeasonID is required")
	}
	params.Set("SeasonID", string(req.SeasonID))

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "playoffpicture", params, &rawResp); err != nil {
		return nil, err
	}

	response := &PlayoffPictureResponse{}
	if len(rawResp.ResultSets) > 0 {
		response.EastConfPlayoffPicture = make([]PlayoffPictureEastConfPlayoffPicture, 0, len(rawResp.ResultSets[0].RowSet))
		for _, row := range rawResp.ResultSets[0].RowSet {
			if len(row) >= 14 {
				item := PlayoffPictureEastConfPlayoffPicture{
					TEAM_ID:                  toInt(row[0]),
					LEAGUE_ID:                toString(row[1]),
					SEASON_ID:                toString(row[2]),
					CONFERENCE:               toString(row[3]),
					RANK:                     toInt(row[4]),
					TEAM:                     toString(row[5]),
					WINS:                     toString(row[6]),
					LOSSES:                   toString(row[7]),
					WIN_PCT:                  toFloat(row[8]),
					GAMES_BACK:               toString(row[9]),
					CLINCHED:                 toString(row[10]),
					ELIMINATED_FROM_PLAYOFFS: toFloat(row[11]),
					CAN_WIN_CONF:             toString(row[12]),
					CAN_WIN_DIV:              toString(row[13]),
				}
				response.EastConfPlayoffPicture = append(response.EastConfPlayoffPicture, item)
			}
		}
	}
	if len(rawResp.ResultSets) > 1 {
		response.WestConfPlayoffPicture = make([]PlayoffPictureWestConfPlayoffPicture, 0, len(rawResp.ResultSets[1].RowSet))
		for _, row := range rawResp.ResultSets[1].RowSet {
			if len(row) >= 14 {
				item := PlayoffPictureWestConfPlayoffPicture{
					TEAM_ID:                  toInt(row[0]),
					LEAGUE_ID:                toString(row[1]),
					SEASON_ID:                toString(row[2]),
					CONFERENCE:               toString(row[3]),
					RANK:                     toInt(row[4]),
					TEAM:                     toString(row[5]),
					WINS:                     toString(row[6]),
					LOSSES:                   toString(row[7]),
					WIN_PCT:                  toFloat(row[8]),
					GAMES_BACK:               toString(row[9]),
					CLINCHED:                 toString(row[10]),
					ELIMINATED_FROM_PLAYOFFS: toFloat(row[11]),
					CAN_WIN_CONF:             toString(row[12]),
					CAN_WIN_DIV:              toString(row[13]),
				}
				response.WestConfPlayoffPicture = append(response.WestConfPlayoffPicture, item)
			}
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}
