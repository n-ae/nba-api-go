package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/endpoints"
	"github.com/username/nba-api-go/pkg/stats/parameters"
)

// This example demonstrates the type safety improvements in nba-api-go
// After type inference, all fields are properly typed with compile-time checking

func main() {
	fmt.Println("🏀 NBA API Go - Type Safety Demo")
	fmt.Println("=================================")
	fmt.Println()

	// Create a client
	client := stats.NewDefaultClient()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Example 1: BoxScoreTraditionalV2 - Type-safe player stats
	demoBoxScore(ctx, client)

	// Example 2: LeagueGameFinder - Type-safe game search
	demoGameFinder(ctx, client)

	// Example 3: TeamGameLogs - Type-safe team performance
	demoTeamLogs(ctx, client)
}

func demoBoxScore(ctx context.Context, client *stats.Client) {
	fmt.Println("📊 Example 1: Box Score with Type Safety")
	fmt.Println("-----------------------------------------")

	// Use a recent game ID (you'll need to update this with a current game)
	req := endpoints.BoxScoreTraditionalV2Request{
		GameID: "0022300001", // Example game ID
	}

	resp, err := endpoints.GetBoxScoreTraditionalV2(ctx, client, req)
	if err != nil {
		log.Printf("Error fetching box score: %v", err)
		return
	}

	fmt.Println("\n✅ Type-Safe Player Stats:")
	fmt.Println()

	// Display top 5 scorers - Notice how clean this is!
	// NO type assertions needed - all fields are properly typed
	count := 0
	for _, player := range resp.Data.PlayerStats {
		if count >= 5 {
			break
		}

		// Direct field access - compiler enforces types!
		// player.PLAYER_NAME is string
		// player.PTS is int
		// player.MIN is float64
		// player.FG_PCT is float64

		fmt.Printf("%-20s | %2d pts | %.1f min | %.1f%% FG\n",
			player.PLAYER_NAME, // string - no assertion!
			player.PTS,         // int - no assertion!
			player.MIN,         // float64 - no assertion!
			player.FG_PCT*100,  // float64 - math works directly!
		)

		count++
	}

	// Show team totals
	if len(resp.Data.TeamStats) > 0 {
		fmt.Println("\n✅ Type-Safe Team Stats:")
		for _, team := range resp.Data.TeamStats {
			fmt.Printf("\n%s %s:\n",
				team.TEAM_CITY, // string
				team.TEAM_NAME, // string
			)
			fmt.Printf("  Points: %d\n", team.PTS)           // int
			fmt.Printf("  FG%%: %.1f%%\n", team.FG_PCT*100)  // float64
			fmt.Printf("  3P%%: %.1f%%\n", team.FG3_PCT*100) // float64
			fmt.Printf("  Rebounds: %d\n", team.REB)         // int
			fmt.Printf("  Assists: %d\n", team.AST)          // int
		}
	}

	fmt.Println()
}

func demoGameFinder(ctx context.Context, client *stats.Client) {
	fmt.Println("\n🔍 Example 2: Game Finder with Type Safety")
	fmt.Println("-------------------------------------------")

	req := endpoints.LeagueGameFinderRequest{
		Season:     stringPtr(string(parameters.NewSeason(2023))),
		SeasonType: stringPtr(string(parameters.SeasonTypeRegular)),
		TeamID:     stringPtr("1610612747"), // Lakers team ID
	}

	resp, err := endpoints.GetLeagueGameFinder(ctx, client, req)
	if err != nil {
		log.Printf("Error finding games: %v", err)
		return
	}

	fmt.Println("\n✅ Type-Safe Game Results:")
	fmt.Println()

	// Display first 5 games
	count := 0
	for _, game := range resp.Data.LeagueGameFinderResults {
		if count >= 5 {
			break
		}

		// All fields properly typed - no assertions!
		// WL is string
		// PTS is int
		// FG_PCT is float64
		// PLUS_MINUS is float64

		outcome := "W"
		if game.WL == "L" {
			outcome = "L"
		}

		fmt.Printf("%s | %s | %s (%s) - %d pts | %.1f%% FG | %+.0f\n",
			game.GAME_DATE,  // string
			outcome,         // string (from WL)
			game.MATCHUP,    // string
			game.WL,         // string
			game.PTS,        // int
			game.FG_PCT*100, // float64
			game.PLUS_MINUS, // float64
		)

		count++
	}

	// Calculate average stats - type safety makes this easy!
	if len(resp.Data.LeagueGameFinderResults) > 0 {
		var totalPts, totalFG, totalReb, totalAst int
		var totalFGPct, totalPlusMinus float64

		for _, game := range resp.Data.LeagueGameFinderResults {
			totalPts += game.PTS              // int addition
			totalFG += game.FGM               // int addition
			totalReb += game.REB              // int addition
			totalAst += game.AST              // int addition
			totalFGPct += game.FG_PCT         // float64 addition
			totalPlusMinus += game.PLUS_MINUS // float64 addition
		}

		games := float64(len(resp.Data.LeagueGameFinderResults))

		fmt.Println("\n✅ Type-Safe Averages:")
		fmt.Printf("  Games: %.0f\n", games)
		fmt.Printf("  Avg Points: %.1f\n", float64(totalPts)/games)
		fmt.Printf("  Avg FG%%: %.1f%%\n", (totalFGPct/games)*100)
		fmt.Printf("  Avg Rebounds: %.1f\n", float64(totalReb)/games)
		fmt.Printf("  Avg Assists: %.1f\n", float64(totalAst)/games)
		fmt.Printf("  Avg +/-: %+.1f\n", totalPlusMinus/games)
	}

	fmt.Println()
}

func demoTeamLogs(ctx context.Context, client *stats.Client) {
	fmt.Println("\n📅 Example 3: Team Game Logs with Type Safety")
	fmt.Println("---------------------------------------------")

	req := endpoints.TeamGameLogsRequest{
		Season:     parameters.NewSeason(2023),
		SeasonType: parameters.SeasonTypeRegular,
	}

	resp, err := endpoints.GetTeamGameLogs(ctx, client, req)
	if err != nil {
		log.Printf("Error fetching team logs: %v", err)
		return
	}

	fmt.Println("\n✅ Type-Safe Team Game Logs:")
	fmt.Println()

	// Group by team and show latest game
	teamLatest := make(map[int]endpoints.TeamGameLogsTeamGameLogs)

	for _, log := range resp.Data.TeamGameLogs {
		// TEAM_ID is int - perfect for map keys!
		if _, exists := teamLatest[log.TEAM_ID]; !exists {
			teamLatest[log.TEAM_ID] = log
		}
	}

	// Display first 5 teams
	count := 0
	for teamID, log := range teamLatest {
		if count >= 5 {
			break
		}

		// All fields properly typed!
		// Type safety enables complex operations

		efficiency := float64(log.PTS) / float64(log.FGA+log.FTA) * 100

		fmt.Printf("%d. %s %s\n",
			count+1,
			log.TEAM_ABBREVIATION, // string
			log.TEAM_NAME,         // string
		)
		fmt.Printf("   Latest: %s vs %s (%s)\n",
			log.GAME_DATE, // string
			log.MATCHUP,   // string
			log.WL,        // string
		)
		fmt.Printf("   Stats: %d pts, %d reb, %d ast (%.1f eff)\n",
			log.PTS,    // int
			log.REB,    // int
			log.AST,    // int
			efficiency, // calculated from ints and floats
		)
		fmt.Printf("   Shooting: %.1f%% FG, %.1f%% 3P, %.1f%% FT\n",
			log.FG_PCT*100,  // float64
			log.FG3_PCT*100, // float64
			log.FT_PCT*100,  // float64
		)
		fmt.Printf("   Advanced: +%.1f, %.1f fantasy pts, %d DD2, %d TD3\n",
			log.PLUS_MINUS,      // float64
			log.NBA_FANTASY_PTS, // float64
			log.DD2,             // int
			log.TD3,             // int
		)
		fmt.Println()

		count++
	}

	fmt.Println()
}

// Helper function
func stringPtr(s string) *string {
	return &s
}

// Key Takeaways from this example:
//
// 1. NO TYPE ASSERTIONS REQUIRED
//    Before: player.PLAYER_NAME.(string)
//    After:  player.PLAYER_NAME
//
// 2. FULL IDE SUPPORT
//    - Autocomplete works for all fields
//    - Type hints show int/float64/string
//    - Go to definition works
//
// 3. COMPILE-TIME CHECKING
//    - Wrong types caught at build time
//    - No runtime panics from bad assertions
//    - Refactoring is safe
//
// 4. CLEAN CODE
//    - No error handling for type assertions
//    - Direct field access
//    - Math operations work naturally
//
// 5. BETTER PERFORMANCE
//    - No reflection at runtime
//    - Direct memory access
//    - Compiler optimizations enabled
//
// This is the power of type inference - production-ready,
// type-safe code that's a joy to work with!
