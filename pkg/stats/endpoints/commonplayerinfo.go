package endpoints

import (
	"context"
	"fmt"
	"net/url"

	"github.com/username/nba-api-go/pkg/models"
	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

type PlayerInfo struct {
	PersonID               int    `json:"PERSON_ID"`
	FirstName              string `json:"FIRST_NAME"`
	LastName               string `json:"LAST_NAME"`
	DisplayFirstLast       string `json:"DISPLAY_FIRST_LAST"`
	DisplayLastCommaFirst  string `json:"DISPLAY_LAST_COMMA_FIRST"`
	DisplayFILast          string `json:"DISPLAY_FI_LAST"`
	PlayerSlug             string `json:"PLAYER_SLUG"`
	Birthdate              string `json:"BIRTHDATE"`
	School                 string `json:"SCHOOL"`
	Country                string `json:"COUNTRY"`
	LastAffiliation        string `json:"LAST_AFFILIATION"`
	Height                 string `json:"HEIGHT"`
	Weight                 string `json:"WEIGHT"`
	SeasonExp              int    `json:"SEASON_EXP"`
	Jersey                 string `json:"JERSEY"`
	Position               string `json:"POSITION"`
	RosterStatus           string `json:"ROSTERSTATUS"`
	TeamID                 int    `json:"TEAM_ID"`
	TeamName               string `json:"TEAM_NAME"`
	TeamAbbreviation       string `json:"TEAM_ABBREVIATION"`
	TeamCode               string `json:"TEAM_CODE"`
	TeamCity               string `json:"TEAM_CITY"`
	PlayerCode             string `json:"PLAYERCODE"`
	FromYear               string `json:"FROM_YEAR"`
	ToYear                 string `json:"TO_YEAR"`
	DLeagueFlag            string `json:"DLEAGUE_FLAG"`
	NBAFlag                string `json:"NBA_FLAG"`
	GamesPlayedFlag        string `json:"GAMES_PLAYED_FLAG"`
	DraftYear              string `json:"DRAFT_YEAR"`
	DraftRound             string `json:"DRAFT_ROUND"`
	DraftNumber            string `json:"DRAFT_NUMBER"`
}

type HeadlineStats struct {
	PlayerID   int     `json:"PLAYER_ID"`
	PlayerName string  `json:"PLAYER_NAME"`
	TimeFrame  string  `json:"TimeFrame"`
	PTS        float64 `json:"PTS"`
	AST        float64 `json:"AST"`
	REB        float64 `json:"REB"`
	PIE        float64 `json:"PIE"`
}

type AvailableSeason struct {
	SeasonID string `json:"SEASON_ID"`
}

type CommonPlayerInfoResponse struct {
	CommonPlayerInfo   []PlayerInfo      `json:"CommonPlayerInfo"`
	PlayerHeadlineStats []HeadlineStats   `json:"PlayerHeadlineStats"`
	AvailableSeasons   []AvailableSeason `json:"AvailableSeasons"`
}

type CommonPlayerInfoRequest struct {
	PlayerID string
	LeagueID parameters.LeagueID
}

func CommonPlayerInfo(ctx context.Context, client *stats.Client, req CommonPlayerInfoRequest) (*models.Response[*CommonPlayerInfoResponse], error) {
	if req.PlayerID == "" {
		return nil, fmt.Errorf("%w: PlayerID is required", models.ErrInvalidRequest)
	}

	if err := req.LeagueID.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrInvalidRequest, err)
	}

	params := url.Values{}
	params.Set("PlayerID", req.PlayerID)
	if req.LeagueID != "" {
		params.Set("LeagueID", req.LeagueID.String())
	}

	var rawResp rawStatsResponse
	if err := client.GetJSON(ctx, "/commonplayerinfo", params, &rawResp); err != nil {
		return nil, err
	}

	response := &CommonPlayerInfoResponse{}
	for _, resultSet := range rawResp.ResultSets {
		switch resultSet.Name {
		case "CommonPlayerInfo":
			response.CommonPlayerInfo = parsePlayerInfo(resultSet.RowSet)
		case "PlayerHeadlineStats":
			response.PlayerHeadlineStats = parseHeadlineStats(resultSet.RowSet)
		case "AvailableSeasons":
			response.AvailableSeasons = parseAvailableSeasons(resultSet.RowSet)
		}
	}

	return models.NewResponse(response, 200, "", nil), nil
}

func parsePlayerInfo(rows [][]interface{}) []PlayerInfo {
	infos := make([]PlayerInfo, 0, len(rows))
	for _, row := range rows {
		if len(row) < 31 {
			continue
		}
		info := PlayerInfo{
			PersonID:              toInt(row[0]),
			FirstName:             toString(row[1]),
			LastName:              toString(row[2]),
			DisplayFirstLast:      toString(row[3]),
			DisplayLastCommaFirst: toString(row[4]),
			DisplayFILast:         toString(row[5]),
			PlayerSlug:            toString(row[6]),
			Birthdate:             toString(row[7]),
			School:                toString(row[8]),
			Country:               toString(row[9]),
			LastAffiliation:       toString(row[10]),
			Height:                toString(row[11]),
			Weight:                toString(row[12]),
			SeasonExp:             toInt(row[13]),
			Jersey:                toString(row[14]),
			Position:              toString(row[15]),
			RosterStatus:          toString(row[16]),
			TeamID:                toInt(row[17]),
			TeamName:              toString(row[18]),
			TeamAbbreviation:      toString(row[19]),
			TeamCode:              toString(row[20]),
			TeamCity:              toString(row[21]),
			PlayerCode:            toString(row[22]),
			FromYear:              toString(row[23]),
			ToYear:                toString(row[24]),
			DLeagueFlag:           toString(row[25]),
			NBAFlag:               toString(row[26]),
			GamesPlayedFlag:       toString(row[27]),
			DraftYear:             toString(row[28]),
			DraftRound:            toString(row[29]),
			DraftNumber:           toString(row[30]),
		}
		infos = append(infos, info)
	}
	return infos
}

func parseHeadlineStats(rows [][]interface{}) []HeadlineStats {
	stats := make([]HeadlineStats, 0, len(rows))
	for _, row := range rows {
		if len(row) < 7 {
			continue
		}
		stat := HeadlineStats{
			PlayerID:   toInt(row[0]),
			PlayerName: toString(row[1]),
			TimeFrame:  toString(row[2]),
			PTS:        toFloat(row[3]),
			AST:        toFloat(row[4]),
			REB:        toFloat(row[5]),
			PIE:        toFloat(row[6]),
		}
		stats = append(stats, stat)
	}
	return stats
}

func parseAvailableSeasons(rows [][]interface{}) []AvailableSeason {
	seasons := make([]AvailableSeason, 0, len(rows))
	for _, row := range rows {
		if len(row) < 1 {
			continue
		}
		season := AvailableSeason{
			SeasonID: toString(row[0]),
		}
		seasons = append(seasons, season)
	}
	return seasons
}
