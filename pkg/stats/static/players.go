package static

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//go:embed data/players.json
var playersData []byte

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	IsActive  bool   `json:"is_active"`
}

var (
	allPlayers   []Player
	playersByID  map[int]Player
	playersIndex sync.Once
	indexErr     error
)

func initPlayers() {
	playersIndex.Do(func() {
		if err := json.Unmarshal(playersData, &allPlayers); err != nil {
			indexErr = fmt.Errorf("failed to unmarshal players data: %w", err)
			return
		}

		playersByID = make(map[int]Player, len(allPlayers))
		for _, player := range allPlayers {
			playersByID[player.ID] = player
		}
	})
}

func GetAllPlayers() ([]Player, error) {
	initPlayers()
	if indexErr != nil {
		return nil, indexErr
	}
	result := make([]Player, len(allPlayers))
	copy(result, allPlayers)
	return result, nil
}

func GetActivePlayers() ([]Player, error) {
	initPlayers()
	if indexErr != nil {
		return nil, indexErr
	}

	active := make([]Player, 0)
	for _, player := range allPlayers {
		if player.IsActive {
			active = append(active, player)
		}
	}
	return active, nil
}

func GetInactivePlayers() ([]Player, error) {
	initPlayers()
	if indexErr != nil {
		return nil, indexErr
	}

	inactive := make([]Player, 0)
	for _, player := range allPlayers {
		if !player.IsActive {
			inactive = append(inactive, player)
		}
	}
	return inactive, nil
}

func FindPlayerByID(playerID int) (*Player, error) {
	initPlayers()
	if indexErr != nil {
		return nil, indexErr
	}

	player, exists := playersByID[playerID]
	if !exists {
		return nil, nil
	}
	return &player, nil
}

func FindPlayersByFullName(fullName string) ([]Player, error) {
	return findPlayers(fullName, func(p Player) string { return p.FullName })
}

func FindPlayersByLastName(lastName string) ([]Player, error) {
	return findPlayers(lastName, func(p Player) string { return p.LastName })
}

func FindPlayersByFirstName(firstName string) ([]Player, error) {
	return findPlayers(firstName, func(p Player) string { return p.FirstName })
}

func findPlayers(pattern string, fieldExtractor func(Player) string) ([]Player, error) {
	initPlayers()
	if indexErr != nil {
		return nil, indexErr
	}

	normalizedPattern := stripAccents(pattern)

	regex, err := regexp.Compile("(?i)" + normalizedPattern)
	if err != nil {
		return nil, fmt.Errorf("invalid pattern: %w", err)
	}

	matches := make([]Player, 0)
	for _, player := range allPlayers {
		field := fieldExtractor(player)
		normalizedField := stripAccents(field)

		if regex.MatchString(normalizedField) {
			matches = append(matches, player)
		}
	}

	return matches, nil
}

func stripAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

func SearchPlayers(query string) ([]Player, error) {
	initPlayers()
	if indexErr != nil {
		return nil, indexErr
	}

	normalizedQuery := strings.ToLower(stripAccents(query))

	matches := make([]Player, 0)
	for _, player := range allPlayers {
		normalizedName := strings.ToLower(stripAccents(player.FullName))
		normalizedFirst := strings.ToLower(stripAccents(player.FirstName))
		normalizedLast := strings.ToLower(stripAccents(player.LastName))

		if strings.Contains(normalizedName, normalizedQuery) ||
			strings.Contains(normalizedFirst, normalizedQuery) ||
			strings.Contains(normalizedLast, normalizedQuery) {
			matches = append(matches, player)
		}
	}

	return matches, nil
}
