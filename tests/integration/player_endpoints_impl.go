package integration

import (
	"fmt"

	"github.com/yourn-ae/nba-api-go/pkg/stats/endpoints"
	"github.com/yourn-ae/nba-api-go/pkg/stats/parameters"
)

// Additional player endpoint test implementations

func testPlayerProfileV2(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerProfileV2Request{
		PlayerID: TestPlayerID,
	}

	resp, err := endpoints.GetPlayerProfileV2(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.SeasonTotalsRegularSeason) == 0 {
		return fmt.Errorf("expected season totals, got empty response")
	}

	return nil
}

func testPlayerDashboardByYearOverYear(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByYearOverYearRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByYearOverYear(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected dashboard data, got empty response")
	}

	return nil
}

func testPlayerDashboardByShootingSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByShootingSplitsRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByShootingSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected shooting splits, got empty response")
	}

	return nil
}

func testPlayerDashboardByOpponent(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByOpponentRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByOpponent(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected opponent data, got empty response")
	}

	return nil
}

func testPlayerDashboardByClutch(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByClutchRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByClutch(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected clutch data, got empty response")
	}

	return nil
}

func testPlayerDashboardByLastNGames(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByLastNGamesRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByLastNGames(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected last N games data, got empty response")
	}

	return nil
}

func testPlayerDashboardByTeamPerformance(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByTeamPerformanceRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByTeamPerformance(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected team performance data, got empty response")
	}

	return nil
}

func testPlayerDashboardByGameSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashboardByGameSplitsRequest{
		PlayerID:    TestPlayerID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetPlayerDashboardByGameSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallPlayerDashboard) == 0 {
		return fmt.Errorf("expected game splits data, got empty response")
	}

	return nil
}

func testPlayerDashPtShots(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerDashPtShotsRequest{
		PlayerID:   TestPlayerID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerDashPtShots(ctx, et.client, req)
	if err != nil {
		return err
	}

	// This endpoint has multiple result sets
	_ = resp

	return nil
}

func testPlayerCompare(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	// Compare two players
	req := endpoints.PlayerCompareRequest{
		PlayerIDList: []int{TestPlayerID, "2544"}, // Jokic vs LeBron
		Season:       TestSeason,
		SeasonType:   TestSeasonType,
	}

	resp, err := endpoints.GetPlayerCompare(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testPlayerVsPlayer(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerVsPlayerRequest{
		PlayerID:   TestPlayerID,
		VsPlayerID: "2544", // vs LeBron
		Season:     TestSeason,
		SeasonType: TestSeasonType,
	}

	resp, err := endpoints.GetPlayerVsPlayer(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testPlayerGameStreakFinder(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerGameStreakFinderRequest{
		PlayerID:   TestPlayerID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
	}

	resp, err := endpoints.GetPlayerGameStreakFinder(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testPlayerGameLogs(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerGameLogsRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerGameLogs(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerGameLogs) == 0 {
		return fmt.Errorf("expected game logs, got empty response")
	}

	return nil
}

func testPlayerNextNGames(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerNextNGamesRequest{
		PlayerID:      TestPlayerID,
		Season:        TestSeason,
		SeasonType:    TestSeasonType,
		NumberOfGames: 5,
	}

	resp, err := endpoints.GetPlayerNextNGames(ctx, et.client, req)
	if err != nil {
		return err
	}

	// May be empty if no upcoming games
	_ = resp
	return nil
}

func testPlayerCareerByCollege(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerCareerByCollegeRequest{
		LeagueID: TestLeagueID,
		College:  "Duke",
	}

	resp, err := endpoints.GetPlayerCareerByCollege(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerCareerByCollege) == 0 {
		return fmt.Errorf("expected college data, got empty response")
	}

	return nil
}

func testPlayerCareerByCollegeRollup(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerCareerByCollegeRollupRequest{
		LeagueID: TestLeagueID,
		PerMode:  parameters.PerModeTotals,
	}

	resp, err := endpoints.GetPlayerCareerByCollegeRollup(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.CollegeStats) == 0 {
		return fmt.Errorf("expected college rollup data, got empty response")
	}

	return nil
}

func testPlayerFantasyProfile(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerFantasyProfileRequest{
		PlayerID: TestPlayerID,
	}

	resp, err := endpoints.GetPlayerFantasyProfile(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.LastNGames) == 0 {
		return fmt.Errorf("expected fantasy data, got empty response")
	}

	return nil
}

func testPlayerEstimatedAdvancedStats(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerEstimatedAdvancedStatsRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerEstimatedAdvancedStats(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerEstimatedAdvancedStats) == 0 {
		return fmt.Errorf("expected estimated advanced stats, got empty response")
	}

	return nil
}

// Player Tracking Endpoints

func testPlayerTrackingSpeedDistance(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingSpeedDistanceRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingSpeedDistance(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingSpeedDistance) == 0 {
		return fmt.Errorf("expected tracking data, got empty response")
	}

	return nil
}

func testPlayerTrackingRebounding(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingReboundingRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingRebounding(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingRebounding) == 0 {
		return fmt.Errorf("expected rebounding data, got empty response")
	}

	return nil
}

func testPlayerTrackingPasses(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingPassesRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingPasses(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingPasses) == 0 {
		return fmt.Errorf("expected passes data, got empty response")
	}

	return nil
}

func testPlayerTrackingDefense(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingDefenseRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingDefense(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingDefense) == 0 {
		return fmt.Errorf("expected defense data, got empty response")
	}

	return nil
}

func testPlayerTrackingCatchShoot(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingCatchShootRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingCatchShoot(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingCatchShoot) == 0 {
		return fmt.Errorf("expected catch & shoot data, got empty response")
	}

	return nil
}

func testPlayerTrackingDrives(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingDrivesRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingDrives(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingDrives) == 0 {
		return fmt.Errorf("expected drives data, got empty response")
	}

	return nil
}

func testPlayerTrackingElbowTouch(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingElbowTouchRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingElbowTouch(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingElbowTouch) == 0 {
		return fmt.Errorf("expected elbow touch data, got empty response")
	}

	return nil
}

func testPlayerTrackingPostTouch(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingPostTouchRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingPostTouch(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingPostTouch) == 0 {
		return fmt.Errorf("expected post touch data, got empty response")
	}

	return nil
}

func testPlayerTrackingPaintTouch(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingPaintTouchRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingPaintTouch(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingPaintTouch) == 0 {
		return fmt.Errorf("expected paint touch data, got empty response")
	}

	return nil
}

func testPlayerTrackingPullUpShot(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingPullUpShotRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingPullUpShot(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingPullUpShot) == 0 {
		return fmt.Errorf("expected pull-up shot data, got empty response")
	}

	return nil
}

func testPlayerTrackingShootingEfficiency(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.PlayerTrackingShootingEfficiencyRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetPlayerTrackingShootingEfficiency(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayerTrackingShootingEfficiency) == 0 {
		return fmt.Errorf("expected shooting efficiency data, got empty response")
	}

	return nil
}
