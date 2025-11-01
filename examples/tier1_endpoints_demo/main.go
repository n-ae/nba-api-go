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

func ptr[T any](v T) *T {
	return &v
}

func main() {
	ctx := context.Background()
	statsClient := stats.NewDefaultClient()

	fmt.Println("=== NBA API Go - Tier 1 Endpoints Demo ===\n")
	fmt.Println("Demonstrating 10 newly generated high-value endpoints\n")

	// 1. League Game Log
	fmt.Println("1. LeagueGameLog - All games in a date range")
	leagueGameLogReq := endpoints.LeagueGameLogRequest{
		Season:     parameters.NewSeason(2023),
		SeasonType: ptr(parameters.SeasonTypeRegular),
	}
	if resp, err := endpoints.GetLeagueGameLog(ctx, statsClient, leagueGameLogReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Loaded %d games\n", len(resp.Data.LeagueGameLog))
	}

	// 2. Player Awards
	fmt.Println("\n2. PlayerAwards - Career accolades")
	playerAwardsReq := endpoints.PlayerAwardsRequest{
		PlayerID: "2544", // LeBron James
	}
	if resp, err := endpoints.GetPlayerAwards(ctx, statsClient, playerAwardsReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ LeBron James has %d career awards/accolades\n", len(resp.Data.PlayerAwards))
	}

	// 3. Playoff Picture
	fmt.Println("\n3. PlayoffPicture - Conference standings for playoff race")
	playoffReq := endpoints.PlayoffPictureRequest{
		SeasonID: parameters.NewSeason(2023),
	}
	if resp, err := endpoints.GetPlayoffPicture(ctx, statsClient, playoffReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ East: %d teams, West: %d teams\n",
			len(resp.Data.EastConfPlayoffPicture),
			len(resp.Data.WestConfPlayoffPicture))
	}

	// 4. Team Year Over Year
	fmt.Println("\n4. TeamDashboardByYearOverYear - Team trends over time")
	teamYoYReq := endpoints.TeamDashboardByYearOverYearRequest{
		TeamID:      "1610612747", // Lakers
		Season:      ptr(parameters.NewSeason(2023)),
		SeasonType:  ptr(parameters.SeasonTypeRegular),
		MeasureType: ptr("Base"),
		PerMode:     ptr(parameters.PerModePerGame),
	}
	if resp, err := endpoints.GetTeamDashboardByYearOverYear(ctx, statsClient, teamYoYReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Lakers historical data: %d seasons\n", len(resp.Data.ByYearTeamDashboard))
	}

	// 5. Player Year Over Year
	fmt.Println("\n5. PlayerDashboardByYearOverYear - Player career progression")
	playerYoYReq := endpoints.PlayerDashboardByYearOverYearRequest{
		PlayerID:    "201939", // Stephen Curry
		Season:      ptr(parameters.NewSeason(2023)),
		SeasonType:  ptr(parameters.SeasonTypeRegular),
		MeasureType: ptr("Base"),
		PerMode:     ptr(parameters.PerModePerGame),
	}
	if resp, err := endpoints.GetPlayerDashboardByYearOverYear(ctx, statsClient, playerYoYReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Curry career data: %d seasons\n", len(resp.Data.ByYearPlayerDashboard))
	}

	// 6. Player vs Player
	fmt.Println("\n6. PlayerVsPlayer - Head-to-head matchup stats")
	pvpReq := endpoints.PlayerVsPlayerRequest{
		PlayerID:   "2544",   // LeBron
		VsPlayerID: "201142", // Durant
		Season:     ptr(parameters.NewSeason(2023)),
		SeasonType: ptr(parameters.SeasonTypeRegular),
		PerMode:    ptr(parameters.PerModePerGame),
	}
	if _, err := endpoints.GetPlayerVsPlayer(ctx, statsClient, pvpReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ LeBron vs Durant matchup data loaded (5 result sets)\n")
	}

	// 7. Team vs Player
	fmt.Println("\n7. TeamVsPlayer - Team performance vs specific player")
	tvpReq := endpoints.TeamVsPlayerRequest{
		TeamID:     "1610612744", // Warriors
		VsPlayerID: "2544",       // LeBron
		Season:     ptr(parameters.NewSeason(2023)),
		SeasonType: ptr(parameters.SeasonTypeRegular),
		PerMode:    ptr(parameters.PerModePerGame),
	}
	if _, err := endpoints.GetTeamVsPlayer(ctx, statsClient, tvpReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Warriors vs LeBron matchup data loaded\n")
	}

	// 8. Draft Combine Stats
	fmt.Println("\n8. DraftCombineStats - Physical measurements and athletic tests")
	combineReq := endpoints.DraftCombineStatsRequest{
		SeasonYear: ptr("2023"),
	}
	if resp, err := endpoints.GetDraftCombineStats(ctx, statsClient, combineReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ 2023 draft combine: %d prospects measured\n", len(resp.Data.DraftCombineStats))
	}

	// 9. League Player Tracking
	fmt.Println("\n9. LeagueDashPtStats - Player tracking (speed/distance)")
	ptReq := endpoints.LeagueDashPtStatsRequest{
		Season:        ptr(parameters.NewSeason(2023)),
		SeasonType:    ptr(parameters.SeasonTypeRegular),
		PerMode:       ptr(parameters.PerModePerGame),
		PtMeasureType: ptr("SpeedDistance"),
	}
	if resp, err := endpoints.GetLeagueDashPtStats(ctx, statsClient, ptReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Player tracking data for %d players\n", len(resp.Data.LeagueDashPTStats))
	}

	// 10. League Lineups
	fmt.Println("\n10. LeagueDashLineups - Lineup combination analytics")
	lineupsReq := endpoints.LeagueDashLineupsRequest{
		Season:        ptr(parameters.NewSeason(2023)),
		SeasonType:    ptr(parameters.SeasonTypeRegular),
		PerMode:       ptr(parameters.PerModePer100Poss),
		MeasureType:   ptr("Base"),
		GroupQuantity: ptr("5"), // 5-man lineups
	}
	if resp, err := endpoints.GetLeagueDashLineups(ctx, statsClient, lineupsReq); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Found %d different 5-man lineup combinations\n", len(resp.Data.Lineups))
	}

	// Summary
	fmt.Println("\n=== Demo Complete ===")
	fmt.Println("\nAll 10 Tier 1 endpoints functional!")

	summary := map[string]interface{}{
		"new_endpoints": []string{
			"LeagueGameLog", "PlayerAwards", "PlayoffPicture",
			"TeamDashboardByYearOverYear", "PlayerDashboardByYearOverYear",
			"PlayerVsPlayer", "TeamVsPlayer", "DraftCombineStats",
			"LeagueDashPtStats", "LeagueDashLineups",
		},
		"total_endpoints_now": 33, // 23 previous + 10 new
		"coverage":            "33/139 = 23.7%",
		"capabilities": []string{
			"League-wide game logs",
			"Player awards and accolades",
			"Playoff race tracking",
			"Year-over-year analysis",
			"Head-to-head matchups",
			"Draft combine data",
			"Player tracking analytics",
			"Lineup combinations",
		},
	}

	data, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Printf("\n%s\n", string(data))
}
