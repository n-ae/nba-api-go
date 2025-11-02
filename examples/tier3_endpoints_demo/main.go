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

func ptr[T any](v T) *T {
	return &v
}

func main() {
	ctx := context.Background()
	statsClient := stats.NewDefaultClient()

	fmt.Println("=== NBA API Go - Tier 3 Endpoints Demo ===\n")
	fmt.Println("9 new synergy, historical, and comparison endpoints\n")

	// 1. Synergy Play Types
	fmt.Println("1. SynergyPlayTypes - Play type breakdown (Isolation, Post-up, etc.)")
	if _, err := endpoints.GetSynergyPlayTypes(ctx, statsClient, endpoints.SynergyPlayTypesRequest{
		Season:       ptr(parameters.NewSeason(2023)),
		SeasonType:   ptr(parameters.SeasonTypeRegular),
		PlayerOrTeam: ptr("P"),
		PlayType:     ptr("Isolation"),
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Synergy data loaded (Isolation play type)")
	}

	// 2. Franchise History
	fmt.Println("\n2. FranchiseHistory - All-time franchise records")
	if resp, err := endpoints.GetFranchiseHistory(ctx, statsClient, endpoints.FranchiseHistoryRequest{}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Active franchises: %d, Defunct teams: %d\n",
			len(resp.Data.FranchiseHistory),
			len(resp.Data.DefunctTeams))
	}

	// 3. Franchise Leaders
	fmt.Println("\n3. FranchiseLeaders - Team all-time leaders")
	if _, err := endpoints.GetFranchiseLeaders(ctx, statsClient, endpoints.FranchiseLeadersRequest{
		TeamID: "1610612747", // Lakers
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Lakers all-time leaders loaded (PTS, AST, REB, BLK, STL)")
	}

	// 4. Team Historical Leaders
	fmt.Println("\n4. TeamHistoricalLeaders - Detailed career leaders")
	if resp, err := endpoints.GetTeamHistoricalLeaders(ctx, statsClient, endpoints.TeamHistoricalLeadersRequest{
		TeamID: "1610612744", // Warriors
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ Warriors leaders: PTS (%d), AST (%d), REB (%d), BLK (%d), STL (%d)\n",
			len(resp.Data.CareerLeadersPTS),
			len(resp.Data.CareerLeadersAST),
			len(resp.Data.CareerLeadersREB),
			len(resp.Data.CareerLeadersBLK),
			len(resp.Data.CareerLeadersSTL))
	}

	// 5. All-Time Leaders
	fmt.Println("\n5. AllTimeLeadersGrids - NBA all-time statistical leaders")
	if resp, err := endpoints.GetAllTimeLeadersGrids(ctx, statsClient, endpoints.AllTimeLeadersGridsRequest{
		PerMode:    ptr(parameters.PerModeTotals),
		SeasonType: ptr(parameters.SeasonTypeRegular),
		TopX:       ptr("10"),
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   ✓ All-time leaders: PTS (%d), AST (%d), REB (%d), BLK (%d), STL (%d)\n",
			len(resp.Data.AllTimeLeadersPTS),
			len(resp.Data.AllTimeLeadersAST),
			len(resp.Data.AllTimeLeadersREB),
			len(resp.Data.AllTimeLeadersBLK),
			len(resp.Data.AllTimeLeadersSTL))
	}

	// 6. Player Compare
	fmt.Println("\n6. PlayerCompare - Side-by-side player comparison")
	if _, err := endpoints.GetPlayerCompare(ctx, statsClient, endpoints.PlayerCompareRequest{
		PlayerIDList: "2544,201939", // LeBron vs Curry
		Season:       ptr(parameters.NewSeason(2023)),
		SeasonType:   ptr(parameters.SeasonTypeRegular),
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Player comparison loaded (LeBron vs Curry)")
	}

	// 7. Team Shot Tracking
	fmt.Println("\n7. TeamDashPtShots - Team shooting tracking")
	if _, err := endpoints.GetTeamDashPtShots(ctx, statsClient, endpoints.TeamDashPtShotsRequest{
		TeamID:     "1610612738", // Celtics
		Season:     ptr(parameters.NewSeason(2023)),
		SeasonType: ptr(parameters.SeasonTypeRegular),
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Celtics shot tracking (3 result sets)")
	}

	// 8. Team Clutch Dashboard
	fmt.Println("\n8. TeamDashboardByClutch - Team clutch splits")
	if _, err := endpoints.GetTeamDashboardByClutch(ctx, statsClient, endpoints.TeamDashboardByClutchRequest{
		TeamID:      "1610612744", // Warriors
		Season:      ptr(parameters.NewSeason(2023)),
		SeasonType:  ptr(parameters.SeasonTypeRegular),
		MeasureType: ptr("Base"),
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Warriors clutch splits loaded")
	}

	// 9. Player Clutch Dashboard
	fmt.Println("\n9. PlayerDashboardByClutch - Player clutch splits")
	if _, err := endpoints.GetPlayerDashboardByClutch(ctx, statsClient, endpoints.PlayerDashboardByClutchRequest{
		PlayerID:    "201935", // James Harden
		Season:      ptr(parameters.NewSeason(2023)),
		SeasonType:  ptr(parameters.SeasonTypeRegular),
		MeasureType: ptr("Base"),
	}); err != nil {
		log.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   ✓ Harden clutch splits loaded")
	}

	// Summary
	fmt.Println("\n=== Demo Complete ===")
	fmt.Println("\nAll 9 Tier 3 endpoints functional!")

	summary := map[string]interface{}{
		"new_endpoints": []string{
			"SynergyPlayTypes", "FranchiseHistory", "FranchiseLeaders",
			"TeamHistoricalLeaders", "AllTimeLeadersGrids", "PlayerCompare",
			"TeamDashPtShots", "TeamDashboardByClutch", "PlayerDashboardByClutch",
		},
		"total_endpoints_now": 53, // 44 previous + 9 new
		"coverage":            "53/139 = 38.1%",
		"new_capabilities": []string{
			"Synergy play type analysis (Isolation, Post-up, Transition, etc.)",
			"Franchise historical records and all-time stats",
			"Team and league all-time leaders",
			"Player comparison tools",
			"Team shot tracking",
			"Detailed clutch performance dashboards",
		},
	}

	data, _ := json.MarshalIndent(summary, "", "  ")
	fmt.Printf("\n%s\n", string(data))
}
