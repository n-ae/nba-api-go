package main

import (
	"fmt"
	"log"

	"github.com/username/nba-api-go/pkg/stats/static"
)

func main() {
	fmt.Println("=== NBA Player Search Examples ===")

	fmt.Println("1. Search by player name:")
	players, err := static.SearchPlayers("lebron")
	if err != nil {
		log.Fatalf("Failed to search players: %v", err)
	}
	for _, player := range players {
		status := "Inactive"
		if player.IsActive {
			status = "Active"
		}
		fmt.Printf("  %s (ID: %d) - %s\n", player.FullName, player.ID, status)
	}

	fmt.Println("\n2. Find player by ID:")
	player, err := static.FindPlayerByID(203999)
	if err != nil {
		log.Fatalf("Failed to find player: %v", err)
	}
	if player != nil {
		fmt.Printf("  Found: %s (ID: %d)\n", player.FullName, player.ID)
	}

	fmt.Println("\n3. Find players by last name (regex):")
	players, err = static.FindPlayersByLastName("^Jordan$")
	if err != nil {
		log.Fatalf("Failed to find players: %v", err)
	}
	fmt.Printf("  Found %d players named Jordan:\n", len(players))
	for _, player := range players {
		fmt.Printf("    %s (ID: %d)\n", player.FullName, player.ID)
	}

	fmt.Println("\n4. Get active players:")
	activePlayers, err := static.GetActivePlayers()
	if err != nil {
		log.Fatalf("Failed to get active players: %v", err)
	}
	fmt.Printf("  Total active players: %d\n", len(activePlayers))

	fmt.Println("\n5. Search teams:")
	teams, err := static.SearchTeams("lakers")
	if err != nil {
		log.Fatalf("Failed to search teams: %v", err)
	}
	for _, team := range teams {
		fmt.Printf("  %s (%s) - Founded: %d\n", team.FullName, team.Abbreviation, team.YearFounded)
	}

	fmt.Println("\n6. Find team by abbreviation:")
	team, err := static.FindTeamByAbbreviation("LAL")
	if err != nil {
		log.Fatalf("Failed to find team: %v", err)
	}
	if team != nil {
		fmt.Printf("  %s - %s, %s\n", team.FullName, team.City, team.State)
	}
}
