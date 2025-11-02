package integration

import (
	"fmt"
	"testing"

	"github.com/yourusername/nba-api-go/pkg/stats/endpoints"
	"github.com/yourusername/nba-api-go/pkg/stats/parameters"
)

// TestTeamEndpoints runs all team-related endpoint tests
func TestTeamEndpoints(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	suite := TestSuite{
		Name:        "TeamEndpoints",
		Description: "Integration tests for all team-related endpoints",
		Tests: []TestEndpoint{
			{
				Name:        "TeamGameLog",
				Description: "Test TeamGameLog endpoint",
				TestFunc:    testTeamGameLog,
			},
			{
				Name:        "TeamGameLogs",
				Description: "Test TeamGameLogs endpoint",
				TestFunc:    testTeamGameLogs,
			},
			{
				Name:        "TeamInfoCommon",
				Description: "Test TeamInfoCommon endpoint",
				TestFunc:    testTeamInfoCommon,
			},
			{
				Name:        "TeamInfoCommonV2",
				Description: "Test TeamInfoCommonV2 endpoint",
				TestFunc:    testTeamInfoCommonV2,
			},
			{
				Name:        "CommonTeamRoster",
				Description: "Test CommonTeamRoster endpoint",
				TestFunc:    testCommonTeamRoster,
			},
			{
				Name:        "CommonTeamRosterV2",
				Description: "Test CommonTeamRosterV2 endpoint",
				TestFunc:    testCommonTeamRosterV2,
			},
			{
				Name:        "CommonTeamYears",
				Description: "Test CommonTeamYears endpoint",
				TestFunc:    testCommonTeamYears,
			},
			{
				Name:        "TeamYearByYearStats",
				Description: "Test TeamYearByYearStats endpoint",
				TestFunc:    testTeamYearByYearStats,
			},
			{
				Name:        "TeamDashboardByGeneralSplits",
				Description: "Test TeamDashboardByGeneralSplits endpoint",
				TestFunc:    testTeamDashboardByGeneralSplits,
			},
			{
				Name:        "TeamDashboardByShootingSplits",
				Description: "Test TeamDashboardByShootingSplits endpoint",
				TestFunc:    testTeamDashboardByShootingSplits,
			},
			{
				Name:        "TeamDashboardByYearOverYear",
				Description: "Test TeamDashboardByYearOverYear endpoint",
				TestFunc:    testTeamDashboardByYearOverYear,
			},
			{
				Name:        "TeamDashboardByOpponent",
				Description: "Test TeamDashboardByOpponent endpoint",
				TestFunc:    testTeamDashboardByOpponent,
			},
			{
				Name:        "TeamDashboardByClutch",
				Description: "Test TeamDashboardByClutch endpoint",
				TestFunc:    testTeamDashboardByClutch,
			},
			{
				Name:        "TeamDashboardByLastNGames",
				Description: "Test TeamDashboardByLastNGames endpoint",
				TestFunc:    testTeamDashboardByLastNGames,
			},
			{
				Name:        "TeamDashboardByTeamPerformance",
				Description: "Test TeamDashboardByTeamPerformance endpoint",
				TestFunc:    testTeamDashboardByTeamPerformance,
			},
			{
				Name:        "TeamDashboardByGameSplits",
				Description: "Test TeamDashboardByGameSplits endpoint",
				TestFunc:    testTeamDashboardByGameSplits,
			},
			{
				Name:        "TeamYearOverYearSplits",
				Description: "Test TeamYearOverYearSplits endpoint",
				TestFunc:    testTeamYearOverYearSplits,
			},
			{
				Name:        "TeamDetails",
				Description: "Test TeamDetails endpoint",
				TestFunc:    testTeamDetails,
			},
			{
				Name:        "TeamHistoricalLeaders",
				Description: "Test TeamHistoricalLeaders endpoint",
				TestFunc:    testTeamHistoricalLeaders,
			},
			{
				Name:        "TeamPlayerDashboard",
				Description: "Test TeamPlayerDashboard endpoint",
				TestFunc:    testTeamPlayerDashboard,
			},
			{
				Name:        "TeamVsPlayer",
				Description: "Test TeamVsPlayer endpoint",
				TestFunc:    testTeamVsPlayer,
			},
			{
				Name:        "TeamVsTeam",
				Description: "Test TeamVsTeam endpoint",
				TestFunc:    testTeamVsTeam,
			},
			{
				Name:        "TeamGameStreakFinder",
				Description: "Test TeamGameStreakFinder endpoint",
				TestFunc:    testTeamGameStreakFinder,
			},
			{
				Name:        "TeamDashPtShots",
				Description: "Test TeamDashPtShots endpoint",
				TestFunc:    testTeamDashPtShots,
			},
			{
				Name:        "TeamEstimatedMetrics",
				Description: "Test TeamEstimatedMetrics endpoint",
				TestFunc:    testTeamEstimatedMetrics,
			},
			{
				Name:        "TeamLineups",
				Description: "Test TeamLineups endpoint",
				TestFunc:    testTeamLineups,
			},
			{
				Name:        "TeamAndPlayersVsPlayers",
				Description: "Test TeamAndPlayersVsPlayers endpoint",
				TestFunc:    testTeamAndPlayersVsPlayers,
			},
			{
				Name:        "TeamNextNGames",
				Description: "Test TeamNextNGames endpoint",
				TestFunc:    testTeamNextNGames,
			},
			{
				Name:        "TeamPlayerOnOffDetails",
				Description: "Test TeamPlayerOnOffDetails endpoint",
				TestFunc:    testTeamPlayerOnOffDetails,
			},
			{
				Name:        "TeamPlayerOnOffSummary",
				Description: "Test TeamPlayerOnOffSummary endpoint",
				TestFunc:    testTeamPlayerOnOffSummary,
			},
		},
	}

	suite.Run(t)
}

// Test implementations for team endpoints

func testTeamGameLog(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamGameLogRequest{
		TeamID:     TestTeamID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
	}

	resp, err := endpoints.GetTeamGameLog(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamGameLog) == 0 {
		return fmt.Errorf("expected game log entries, got empty response")
	}

	return nil
}

func testTeamGameLogs(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamGameLogsRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetTeamGameLogs(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamGameLogs) == 0 {
		return fmt.Errorf("expected game logs, got empty response")
	}

	return nil
}

func testTeamInfoCommon(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamInfoCommonRequest{
		TeamID:   TestTeamID,
		Season:   TestSeason,
		LeagueID: TestLeagueID,
	}

	resp, err := endpoints.GetTeamInfoCommon(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamInfoCommon) == 0 {
		return fmt.Errorf("expected team info, got empty response")
	}

	return nil
}

func testTeamInfoCommonV2(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamInfoCommonV2Request{
		TeamID:   TestTeamID,
		LeagueID: TestLeagueID,
	}

	resp, err := endpoints.GetTeamInfoCommonV2(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamInfoCommon) == 0 {
		return fmt.Errorf("expected team info, got empty response")
	}

	return nil
}

func testCommonTeamRoster(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.CommonTeamRosterRequest{
		TeamID: TestTeamID,
		Season: TestSeason,
	}

	resp, err := endpoints.GetCommonTeamRoster(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.CommonTeamRoster) == 0 {
		return fmt.Errorf("expected roster data, got empty response")
	}

	return nil
}

func testCommonTeamRosterV2(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.CommonTeamRosterV2Request{
		TeamID: TestTeamID,
		Season: TestSeason,
	}

	resp, err := endpoints.GetCommonTeamRosterV2(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.CommonTeamRoster) == 0 {
		return fmt.Errorf("expected roster data, got empty response")
	}

	return nil
}

func testCommonTeamYears(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.CommonTeamYearsRequest{
		LeagueID: TestLeagueID,
	}

	resp, err := endpoints.GetCommonTeamYears(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamYears) == 0 {
		return fmt.Errorf("expected team years data, got empty response")
	}

	return nil
}

func testTeamYearByYearStats(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamYearByYearStatsRequest{
		TeamID:   TestTeamID,
		PerMode:  parameters.PerModePerGame,
		LeagueID: TestLeagueID,
	}

	resp, err := endpoints.GetTeamYearByYearStats(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamStats) == 0 {
		return fmt.Errorf("expected year by year stats, got empty response")
	}

	return nil
}

func testTeamDashboardByGeneralSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByGeneralSplitsRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByGeneralSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected dashboard data, got empty response")
	}

	return nil
}

func testTeamDashboardByShootingSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByShootingSplitsRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByShootingSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected shooting splits, got empty response")
	}

	return nil
}

func testTeamDashboardByYearOverYear(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByYearOverYearRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByYearOverYear(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected year over year data, got empty response")
	}

	return nil
}

func testTeamDashboardByOpponent(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByOpponentRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByOpponent(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected opponent data, got empty response")
	}

	return nil
}

func testTeamDashboardByClutch(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByClutchRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByClutch(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected clutch data, got empty response")
	}

	return nil
}

func testTeamDashboardByLastNGames(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByLastNGamesRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByLastNGames(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected last N games data, got empty response")
	}

	return nil
}

func testTeamDashboardByTeamPerformance(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByTeamPerformanceRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByTeamPerformance(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected team performance data, got empty response")
	}

	return nil
}

func testTeamDashboardByGameSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashboardByGameSplitsRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamDashboardByGameSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.OverallTeamDashboard) == 0 {
		return fmt.Errorf("expected game splits data, got empty response")
	}

	return nil
}

func testTeamYearOverYearSplits(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamYearOverYearSplitsRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamYearOverYearSplits(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.ByYearTeamDashboard) == 0 {
		return fmt.Errorf("expected year over year data, got empty response")
	}

	return nil
}

func testTeamDetails(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDetailsRequest{
		TeamID: TestTeamID,
	}

	resp, err := endpoints.GetTeamDetails(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamBackground) == 0 {
		return fmt.Errorf("expected team details, got empty response")
	}

	return nil
}

func testTeamHistoricalLeaders(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamHistoricalLeadersRequest{
		TeamID:   TestTeamID,
		LeagueID: TestLeagueID,
	}

	resp, err := endpoints.GetTeamHistoricalLeaders(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamHistoricalLeaders) == 0 {
		return fmt.Errorf("expected historical leaders, got empty response")
	}

	return nil
}

func testTeamPlayerDashboard(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamPlayerDashboardRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamPlayerDashboard(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.PlayersSeasonTotals) == 0 {
		return fmt.Errorf("expected player dashboard, got empty response")
	}

	return nil
}

func testTeamVsPlayer(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamVsPlayerRequest{
		TeamID:     TestTeamID,
		VsPlayerID: TestPlayerID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
	}

	resp, err := endpoints.GetTeamVsPlayer(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testTeamVsTeam(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamVsTeamRequest{
		TeamID:      TestTeamID,
		VsTeamID:    1610612739, // Cleveland Cavaliers
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamVsTeam(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.Overall) == 0 {
		return fmt.Errorf("expected matchup data, got empty response")
	}

	return nil
}

func testTeamGameStreakFinder(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamGameStreakFinderRequest{
		TeamID:     TestTeamID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
	}

	resp, err := endpoints.GetTeamGameStreakFinder(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testTeamDashPtShots(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamDashPtShotsRequest{
		TeamID:     TestTeamID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetTeamDashPtShots(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testTeamEstimatedMetrics(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamEstimatedMetricsRequest{
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		LeagueID:   TestLeagueID,
	}

	resp, err := endpoints.GetTeamEstimatedMetrics(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamEstimatedMetrics) == 0 {
		return fmt.Errorf("expected estimated metrics, got empty response")
	}

	return nil
}

func testTeamLineups(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamLineupsRequest{
		TeamID:        TestTeamID,
		Season:        TestSeason,
		SeasonType:    TestSeasonType,
		MeasureType:   "Base",
		PerMode:       parameters.PerModePerGame,
		GroupQuantity: 5,
	}

	resp, err := endpoints.GetTeamLineups(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.Lineups) == 0 {
		return fmt.Errorf("expected lineup data, got empty response")
	}

	return nil
}

func testTeamAndPlayersVsPlayers(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamAndPlayersVsPlayersRequest{
		TeamID:     TestTeamID,
		VsPlayerID: TestPlayerID,
		Season:     TestSeason,
		SeasonType: TestSeasonType,
		PerMode:    parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamAndPlayersVsPlayers(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testTeamNextNGames(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamNextNGamesRequest{
		TeamID:        TestTeamID,
		Season:        TestSeason,
		SeasonType:    TestSeasonType,
		NumberOfGames: 5,
	}

	resp, err := endpoints.GetTeamNextNGames(ctx, et.client, req)
	if err != nil {
		return err
	}

	// May be empty if no upcoming games
	_ = resp
	return nil
}

func testTeamPlayerOnOffDetails(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamPlayerOnOffDetailsRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamPlayerOnOffDetails(ctx, et.client, req)
	if err != nil {
		return err
	}

	_ = resp
	return nil
}

func testTeamPlayerOnOffSummary(et *EndpointTester) error {
	ctx, cancel := et.Context()
	defer cancel()

	req := endpoints.TeamPlayerOnOffSummaryRequest{
		TeamID:      TestTeamID,
		Season:      TestSeason,
		SeasonType:  TestSeasonType,
		MeasureType: "Base",
		PerMode:     parameters.PerModePerGame,
	}

	resp, err := endpoints.GetTeamPlayerOnOffSummary(ctx, et.client, req)
	if err != nil {
		return err
	}

	if len(resp.TeamPlayerOnOffSummary) == 0 {
		return fmt.Errorf("expected on/off summary, got empty response")
	}

	return nil
}
