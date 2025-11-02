package integration

import (
	"testing"

	"github.com/yourusername/nba-api-go/pkg/stats/endpoints"
	"github.com/yourusername/nba-api-go/pkg/stats/parameters"
)

// TestPlayerEndpoints runs all player-related endpoint tests
func TestPlayerEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	suite := TestSuite{
		Name:        "PlayerEndpoints",
		Description: "Integration tests for all player-related endpoints",
		Tests: []TestEndpoint{
			{
				Name:        "PlayerCareerStats",
				Description: "Test PlayerCareerStats endpoint",
				TestFunc:    testPlayerCareerStats,
			},
			{
				Name:        "PlayerGameLog",
				Description: "Test PlayerGameLog endpoint",
				TestFunc:    testPlayerGameLog,
			},
			{
				Name:        "PlayerYearByYearStats",
				Description: "Test PlayerYearByYearStats endpoint",
				TestFunc:    testPlayerYearByYearStats,
			},
			{
				Name:        "PlayerDashboardByGeneralSplits",
				Description: "Test PlayerDashboardByGeneralSplits endpoint",
				TestFunc:    testPlayerDashboardByGeneralSplits,
			},
			{
				Name:        "PlayerAwards",
				Description: "Test PlayerAwards endpoint",
				TestFunc:    testPlayerAwards,
			},
			{
				Name:        "PlayerIndex",
				Description: "Test PlayerIndex endpoint",
				TestFunc:    testPlayerIndex,
			},
			{
				Name:        "CommonPlayerInfo",
				Description: "Test CommonPlayerInfo endpoint",
				TestFunc:    testCommonPlayerInfo,
			},
			{
				Name:        "PlayerEstimatedMetrics",
				Description: "Test PlayerEstimatedMetrics endpoint",
				TestFunc:    testPlayerEstimatedMetrics,
			},
			{
				Name:        "PlayerProfileV2",
				Description: "Test PlayerProfileV2 endpoint",
				TestFunc:    testPlayerProfileV2,
			},
			{
				Name:        "PlayerDashboardByYearOverYear",
				Description: "Test PlayerDashboardByYearOverYear endpoint",
				TestFunc:    testPlayerDashboardByYearOverYear,
			},
			{
				Name:        "PlayerDashboardByShootingSplits",
				Description: "Test PlayerDashboardByShootingSplits endpoint",
				TestFunc:    testPlayerDashboardByShootingSplits,
			},
			{
				Name:        "PlayerDashboardByOpponent",
				Description: "Test PlayerDashboardByOpponent endpoint",
				TestFunc:    testPlayerDashboardByOpponent,
			},
			{
				Name:        "PlayerDashboardByClutch",
				Description: "Test PlayerDashboardByClutch endpoint",
				TestFunc:    testPlayerDashboardByClutch,
			},
			{
				Name:        "PlayerDashboardByLastNGames",
				Description: "Test PlayerDashboardByLastNGames endpoint",
				TestFunc:    testPlayerDashboardByLastNGames,
			},
			{
				Name:        "PlayerDashboardByTeamPerformance",
				Description: "Test PlayerDashboardByTeamPerformance endpoint",
				TestFunc:    testPlayerDashboardByTeamPerformance,
			},
			{
				Name:        "PlayerDashboardByGameSplits",
				Description: "Test PlayerDashboardByGameSplits endpoint",
				TestFunc:    testPlayerDashboardByGameSplits,
			},
			{
				Name:        "PlayerDashPtShots",
				Description: "Test PlayerDashPtShots endpoint",
				TestFunc:    testPlayerDashPtShots,
			},
			{
				Name:        "PlayerCompare",
				Description: "Test PlayerCompare endpoint",
				TestFunc:    testPlayerCompare,
			},
			{
				Name:        "PlayerVsPlayer",
				Description: "Test PlayerVsPlayer endpoint",
				TestFunc:    testPlayerVsPlayer,
			},
			{
				Name:        "PlayerGameStreakFinder",
				Description: "Test PlayerGameStreakFinder endpoint",
				TestFunc:    testPlayerGameStreakFinder,
			},
			{
				Name:        "PlayerGameLogs",
				Description: "Test PlayerGameLogs endpoint",
				TestFunc:    testPlayerGameLogs,
			},
			{
				Name:        "PlayerNextNGames",
				Description: "Test PlayerNextNGames endpoint",
				TestFunc:    testPlayerNextNGames,
			},
			{
				Name:        "PlayerCareerByCollege",
				Description: "Test PlayerCareerByCollege endpoint",
				TestFunc:    testPlayerCareerByCollege,
			},
			{
				Name:        "PlayerCareerByCollegeRollup",
				Description: "Test PlayerCareerByCollegeRollup endpoint",
				TestFunc:    testPlayerCareerByCollegeRollup,
			},
			{
				Name:        "PlayerFantasyProfile",
				Description: "Test PlayerFantasyProfile endpoint",
				TestFunc:    testPlayerFantasyProfile,
			},
			{
				Name:        "PlayerEstimatedAdvancedStats",
				Description: "Test PlayerEstimatedAdvancedStats endpoint",
				TestFunc:    testPlayerEstimatedAdvancedStats,
			},
			{
				Name:        "PlayerTrackingSpeedDistance",
				Description: "Test PlayerTrackingSpeedDistance endpoint",
				TestFunc:    testPlayerTrackingSpeedDistance,
			},
			{
				Name:        "PlayerTrackingRebounding",
				Description: "Test PlayerTrackingRebounding endpoint",
				TestFunc:    testPlayerTrackingRebounding,
			},
			{
				Name:        "PlayerTrackingPasses",
				Description: "Test PlayerTrackingPasses endpoint",
				TestFunc:    testPlayerTrackingPasses,
			},
			{
				Name:        "PlayerTrackingDefense",
				Description: "Test PlayerTrackingDefense endpoint",
				TestFunc:    testPlayerTrackingDefense,
			},
			{
				Name:        "PlayerTrackingCatchShoot",
				Description: "Test PlayerTrackingCatchShoot endpoint",
				TestFunc:    testPlayerTrackingCatchShoot,
			},
			{
				Name:        "PlayerTrackingDrives",
				Description: "Test PlayerTrackingDrives endpoint",
				TestFunc:    testPlayerTrackingDrives,
			},
			{
				Name:        "PlayerTrackingElbowTouch",
				Description: "Test PlayerTrackingElbowTouch endpoint",
				TestFunc:    testPlayerTrackingElbowTouch,
			},
			{
				Name:        "PlayerTrackingPostTouch",
				Description: "Test PlayerTrackingPostTouch endpoint",
				TestFunc:    testPlayerTrackingPostTouch,
			},
			{
				Name:        "PlayerTrackingPaintTouch",
				Description: "Test PlayerTrackingPaintTouch endpoint",
				TestFunc:    testPlayerTrackingPaintTouch,
			},
			{
				Name:        "PlayerTrackingPullUpShot",
				Description: "Test PlayerTrackingPullUpShot endpoint",
				TestFunc:    testPlayerTrackingPullUpShot,
			},
			{
				Name:        "PlayerTrackingShootingEfficiency",
				Description: "Test PlayerTrackingShootingEfficiency endpoint",
				TestFunc:    testPlayerTrackingShootingEfficiency,
			},
		},
	}

	suite.Run(t)
}

func testPlayerCareerStats(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerCareerStatsRequest{
		PlayerID: TestPlayerID,
		PerMode:  parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerCareerStats(ctx, et.client, req)
	if err != nil {
		return err
	}

	// Validate response has data
	if len(resp.CareerTotalsRegularSeason) == 0 {
		return fmt.Errorf("expected career stats, got empty response")
	}

	return nil
}

func testPlayerGameLog(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerGameLogRequest{
		PlayerID:   TestPlayerID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
	}

	resp, err := endpoints.GetPlayerGameLog(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerGameLog) == 0 {
		return fmt.Errorf("expected game log entries, got empty response")
	}

	return nil
}

func testPlayerYearByYearStats(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerYearByYearStatsRequest{
		PlayerID: TestPlayerID,
		PerMode:  parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerYearByYearStats(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerStats) == 0 {
		return fmt.Errorf("expected year-by-year stats, got empty response")
	}

	return nil
}

func testPlayerDashboardByGeneralSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByGeneralSplitsRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByGeneralSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	// Validate at least one result set has data
	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected dashboard data, got empty response")
	}

	return nil
}

func testPlayerAwards(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerAwardsRequest{
		PlayerID: TestPlayerID,
	}

	resp, err := endpoints.GetPlayerAwards(ctx, et.client, req)
	if err != nil {
		return err
	}

	// Note: Not all players have awards, so we just check for successful response
	// The response struct should exist even if empty
	_ = resp

	return nil
}

func testPlayerIndex(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerIndexRequest{
		Season:   TestSeason,
		LeagueID: TestLeagueID,
	}

	resp, err := endpoints.GetPlayerIndex(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerIndex) == 0 {
		return fmt.Errorf("expected player index data, got empty response")
	}

	return nil
}

func testCommonPlayerInfo(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.CommonPlayerInfoRequest{
		PlayerID: TestPlayerID,
	}

	resp, err := endpoints.GetCommonPlayerInfo(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.CommonPlayerInfo) == 0 {
		return fmt.Errorf("expected player info, got empty response")
	}

	return nil
}

func testPlayerEstimatedMetrics(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerEstimatedMetricsRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerEstimatedMetrics(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerEstimatedMetrics) == 0 {
		return fmt.Errorf("expected estimated metrics, got empty response")
	}

	return nil
}
