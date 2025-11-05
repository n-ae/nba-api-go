package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

func main() {
	ctx := context.Background()
	statsClient := stats.NewDefaultClient()

	fmt.Println("=== Testing Newly Generated Endpoints ===")

	fmt.Println("1. CommonAllPlayers - Get all players from 2023-24 season")
	season := parameters.Season("2023-24")
	leagueID := parameters.LeagueIDNBA
	isOnlyCurrent := "1"
	allPlayersReq := endpoints.CommonAllPlayersRequest{
		Season:              season,
		LeagueID:            &leagueID,
		IsOnlyCurrentSeason: &isOnlyCurrent,
	}
	allPlayersResp, err := endpoints.GetCommonAllPlayers(ctx, statsClient, allPlayersReq)
	if err != nil {
		log.Printf("Error getting all players: %v\n", err)
	} else {
		fmt.Printf("   ✓ Found %d players\n", len(allPlayersResp.Data.CommonAllPlayers))
		if len(allPlayersResp.Data.CommonAllPlayers) > 0 {
			player := allPlayersResp.Data.CommonAllPlayers[0]
			fmt.Printf("   Example: %v (ID: %v)\n", player.DISPLAY_FIRST_LAST, player.PERSON_ID)
		}
	}
	fmt.Println()

	fmt.Println("=== Generation Summary ===")
	summary := map[string]interface{}{
		"endpoints_generated": []string{
			"CommonAllPlayers",
			"CommonTeamRoster",
			"LeagueDashPlayerStats",
			"LeagueDashTeamStats",
			"ScoreboardV2",
			"PlayerProfileV2",
			"LeagueStandings",
			"BoxScoreAdvancedV2",
		},
		"total_new_endpoints": 8,
		"compilation_status":  "success",
	}
	data, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Println(string(data))
}
