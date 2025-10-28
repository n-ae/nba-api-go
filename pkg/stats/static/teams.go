package static

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

//go:embed data/teams.json
var teamsData []byte

type Team struct {
	ID           int    `json:"id"`
	FullName     string `json:"full_name"`
	Abbreviation string `json:"abbreviation"`
	Nickname     string `json:"nickname"`
	City         string `json:"city"`
	State        string `json:"state"`
	YearFounded  int    `json:"year_founded"`
}

var (
	allTeams   []Team
	teamsByID  map[int]Team
	teamsIndex sync.Once
	teamsErr   error
)

func initTeams() {
	teamsIndex.Do(func() {
		if err := json.Unmarshal(teamsData, &allTeams); err != nil {
			teamsErr = fmt.Errorf("failed to unmarshal teams data: %w", err)
			return
		}

		teamsByID = make(map[int]Team, len(allTeams))
		for _, team := range allTeams {
			teamsByID[team.ID] = team
		}
	})
}

func GetAllTeams() ([]Team, error) {
	initTeams()
	if teamsErr != nil {
		return nil, teamsErr
	}
	result := make([]Team, len(allTeams))
	copy(result, allTeams)
	return result, nil
}

func FindTeamByID(teamID int) (*Team, error) {
	initTeams()
	if teamsErr != nil {
		return nil, teamsErr
	}

	team, exists := teamsByID[teamID]
	if !exists {
		return nil, nil
	}
	return &team, nil
}

func FindTeamByAbbreviation(abbreviation string) (*Team, error) {
	initTeams()
	if teamsErr != nil {
		return nil, teamsErr
	}

	abbr := strings.ToUpper(abbreviation)
	for _, team := range allTeams {
		if strings.ToUpper(team.Abbreviation) == abbr {
			return &team, nil
		}
	}
	return nil, nil
}

func FindTeamsByNickname(nickname string) ([]Team, error) {
	initTeams()
	if teamsErr != nil {
		return nil, teamsErr
	}

	query := strings.ToLower(nickname)
	matches := make([]Team, 0)

	for _, team := range allTeams {
		if strings.Contains(strings.ToLower(team.Nickname), query) {
			matches = append(matches, team)
		}
	}

	return matches, nil
}

func SearchTeams(query string) ([]Team, error) {
	initTeams()
	if teamsErr != nil {
		return nil, teamsErr
	}

	normalizedQuery := strings.ToLower(query)
	matches := make([]Team, 0)

	for _, team := range allTeams {
		if strings.Contains(strings.ToLower(team.FullName), normalizedQuery) ||
			strings.Contains(strings.ToLower(team.Nickname), normalizedQuery) ||
			strings.Contains(strings.ToLower(team.Abbreviation), normalizedQuery) ||
			strings.Contains(strings.ToLower(team.City), normalizedQuery) {
			matches = append(matches, team)
		}
	}

	return matches, nil
}
