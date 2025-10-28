package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/endpoints"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

func main() {
	client := stats.NewDefaultClient()

	req := endpoints.PlayerCareerStatsRequest{
		PlayerID: "203999",
		PerMode:  parameters.PerModePerGame,
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := endpoints.PlayerCareerStats(context.Background(), client, req)
	if err != nil {
		log.Fatalf("Failed to get player career stats: %v", err)
	}

	fmt.Println("=== Regular Season Stats ===")
	for _, season := range resp.Data.SeasonTotalsRegularSeason {
		fmt.Printf("Season %s (%s): %.1f PPG, %.1f RPG, %.1f APG in %d games\n",
			season.SeasonID,
			season.TeamAbbreviation,
			season.PTS,
			season.REB,
			season.AST,
			season.GP,
		)
	}

	if len(resp.Data.CareerTotalsRegularSeason) > 0 {
		career := resp.Data.CareerTotalsRegularSeason[0]
		fmt.Printf("\n=== Career Totals ===\n")
		fmt.Printf("Games: %d\n", career.GP)
		fmt.Printf("Points: %.0f\n", career.PTS)
		fmt.Printf("Rebounds: %.0f\n", career.REB)
		fmt.Printf("Assists: %.0f\n", career.AST)
	}

	jsonData, err := json.MarshalIndent(resp.Data, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Printf("\n=== Full JSON Response ===\n%s\n", string(jsonData))
}
