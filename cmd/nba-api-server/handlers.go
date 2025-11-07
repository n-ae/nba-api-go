package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/n-ae/nba-api-go/pkg/stats"
	"github.com/n-ae/nba-api-go/pkg/stats/parameters"
)

type StatsHandler struct {
	client *stats.Client
}

func NewStatsHandler() *StatsHandler {
	return &StatsHandler{
		client: stats.NewDefaultClient(),
	}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET requests are supported")
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/v1/stats/")
	endpoint := strings.ToLower(path)

	switch endpoint {
	// Original 10 endpoints
	case "playergamelog":
		h.handlePlayerGameLog(w, r)
	case "commonallplayers":
		h.handleCommonAllPlayers(w, r)
	case "scoreboardv2":
		h.handleScoreboardV2(w, r)
	case "leaguestandings":
		h.handleLeagueStandings(w, r)
	case "commonteamroster":
		h.handleCommonTeamRoster(w, r)
	case "playercareerstats":
		h.handlePlayerCareerStats(w, r)
	case "leagueleaders":
		h.handleLeagueLeaders(w, r)
	case "commonplayerinfo":
		h.handleCommonPlayerInfo(w, r)
	case "leaguedashteamstats":
		h.handleLeagueDashTeamStats(w, r)
	case "leaguedashplayerstats":
		h.handleLeagueDashPlayerStats(w, r)

	// Player endpoints (expanded)
	case "playerprofilev2":
		h.handlePlayerProfileV2(w, r)
	case "playerawards":
		h.handlePlayerAwards(w, r)
	case "playerdashboardbygeneralsplits":
		h.handlePlayerDashboardByGeneralSplits(w, r)
	case "playerdashboardbyshootingsplits":
		h.handlePlayerDashboardByShootingSplits(w, r)
	case "playerdashboardbyopponent":
		h.handlePlayerDashboardByOpponent(w, r)
	case "playerdashboardbyclutch":
		h.handlePlayerDashboardByClutch(w, r)
	case "playergamelogs":
		h.handlePlayerGameLogs(w, r)
	case "playervsplayer":
		h.handlePlayerVsPlayer(w, r)

	// Team endpoints (expanded)
	case "teamgamelog":
		h.handleTeamGameLog(w, r)
	case "teaminfocommon":
		h.handleTeamInfoCommon(w, r)
	case "teamdashboardbygeneralsplits":
		h.handleTeamDashboardByGeneralSplits(w, r)
	case "teamdashboardbyshootingsplits":
		h.handleTeamDashboardByShootingSplits(w, r)
	case "teamdashboardbyopponent":
		h.handleTeamDashboardByOpponent(w, r)
	case "teamdetails":
		h.handleTeamDetails(w, r)
	case "teamplayerdashboard":
		h.handleTeamPlayerDashboard(w, r)
	case "teamlineups":
		h.handleTeamLineups(w, r)
	case "teamgamelogs":
		h.handleTeamGameLogs(w, r)
	case "teamyearbyyearstats":
		h.handleTeamYearByYearStats(w, r)
	case "teamvsteam":
		h.handleTeamVsTeam(w, r)
	case "teamhistoricalleaders":
		h.handleTeamHistoricalLeaders(w, r)
	case "teamestimatedmetrics":
		h.handleTeamEstimatedMetrics(w, r)
	case "teamdashptshots":
		h.handleTeamDashPtShots(w, r)
	case "teamdashboardbyclutch":
		h.handleTeamDashboardByClutch(w, r)
	case "teamdashboardbylastngames":
		h.handleTeamDashboardByLastNGames(w, r)
	case "teamdashboardbyyearoveryear":
		h.handleTeamDashboardByYearOverYear(w, r)
	case "teamvsplayer":
		h.handleTeamVsPlayer(w, r)
	case "teamdashboardbygamesplits":
		h.handleTeamDashboardByGameSplits(w, r)
	case "teamdashboardbyteamperformance":
		h.handleTeamDashboardByTeamPerformance(w, r)
	case "teamplayeronoffsummary":
		h.handleTeamPlayerOnOffSummary(w, r)

	// Box Score endpoints (10 total - 100% coverage!)
	case "boxscoresummaryv2":
		h.handleBoxScoreSummaryV2(w, r)
	case "boxscoretraditionalv2":
		h.handleBoxScoreTraditionalV2(w, r)
	case "boxscoreadvancedv2":
		h.handleBoxScoreAdvancedV2(w, r)
	case "boxscorescoringv2":
		h.handleBoxScoreScoringV2(w, r)
	case "boxscoremiscv2":
		h.handleBoxScoreMiscV2(w, r)
	case "boxscoreusagev2":
		h.handleBoxScoreUsageV2(w, r)
	case "boxscorefourfactorsv2":
		h.handleBoxScoreFourFactorsV2(w, r)
	case "boxscoreplayertrackv2":
		h.handleBoxScorePlayerTrackV2(w, r)
	case "boxscoredefensivev2":
		h.handleBoxScoreDefensiveV2(w, r)
	case "boxscorehustlev2":
		h.handleBoxScoreHustleV2(w, r)

	// Player Tracking endpoints
	case "playertrackingshotdashboard":
		h.handlePlayerTrackingShotDashboard(w, r)
	case "playertrackingpasses":
		h.handlePlayerTrackingPasses(w, r)
	case "playertrackingdefense":
		h.handlePlayerTrackingDefense(w, r)
	case "playertrackingrebounding":
		h.handlePlayerTrackingRebounding(w, r)
	case "playertrackingspeeddistance":
		h.handlePlayerTrackingSpeedDistance(w, r)
	case "playertrackingcatchshoot":
		h.handlePlayerTrackingCatchShoot(w, r)
	case "playertrackingdrives":
		h.handlePlayerTrackingDrives(w, r)
	case "playertrackingposttouch":
		h.handlePlayerTrackingPostTouch(w, r)
	case "playertrackingpainttouch":
		h.handlePlayerTrackingPaintTouch(w, r)
	case "playertrackingelbowtouch":
		h.handlePlayerTrackingElbowTouch(w, r)
	case "playertrackingpullupshot":
		h.handlePlayerTrackingPullUpShot(w, r)

	// Game endpoints
	case "playbyplayv2":
		h.handlePlayByPlayV2(w, r)
	case "shotchartdetail":
		h.handleShotChartDetail(w, r)
	case "gamerotation":
		h.handleGameRotation(w, r)

	// League endpoints (expanded)
	case "leaguegamelog":
		h.handleLeagueGameLog(w, r)
	case "playoffpicture":
		h.handlePlayoffPicture(w, r)
	case "leaguedashlineups":
		h.handleLeagueDashLineups(w, r)
	case "leaguedashplayerclutch":
		h.handleLeagueDashPlayerClutch(w, r)
	case "leaguedashteamclutch":
		h.handleLeagueDashTeamClutch(w, r)
	case "leaguedashplayerbiostats":
		h.handleLeagueDashPlayerBioStats(w, r)
	case "leaguedashteambiostats":
		h.handleLeagueDashTeamBioStats(w, r)
	case "leaguedashptstats":
		h.handleLeagueDashPtStats(w, r)
	case "leaguehustlestatsplayer":
		h.handleLeagueHustleStatsPlayer(w, r)
	case "leaguehustlestatsteam":
		h.handleLeagueHustleStatsTeam(w, r)
	case "leaguedashptdefend":
		h.handleLeagueDashPtDefend(w, r)
	case "leaguegamefinder":
		h.handleLeagueGameFinder(w, r)
	case "leaguestandingsv3":
		h.handleLeagueStandingsV3(w, r)
	case "leaguedashplayershotlocations":
		h.handleLeagueDashPlayerShotLocations(w, r)
	case "leaguedashteamshotlocations":
		h.handleLeagueDashTeamShotLocations(w, r)
	case "leagueseasonmatchups":
		h.handleLeagueSeasonMatchups(w, r)
	case "leaguedashptteamdefend":
		h.handleLeagueDashPtTeamDefend(w, r)
	case "leaguedashplayerptshot":
		h.handleLeagueDashPlayerPtShot(w, r)
	case "leaguedashteamptshot":
		h.handleLeagueDashTeamPtShot(w, r)

	// Additional Player endpoints
	case "playerestimatedmetrics":
		h.handlePlayerEstimatedMetrics(w, r)
	case "playerfantasyprofile":
		h.handlePlayerFantasyProfile(w, r)
	case "playerdashptshots":
		h.handlePlayerDashPtShots(w, r)
	case "playerdashboardbylastngames":
		h.handlePlayerDashboardByLastNGames(w, r)
	case "playerdashboardbyteamperformance":
		h.handlePlayerDashboardByTeamPerformance(w, r)
	case "playerdashboardbygamesplits":
		h.handlePlayerDashboardByGameSplits(w, r)
	case "playerdashboardbyyearoveryear":
		h.handlePlayerDashboardByYearOverYear(w, r)
	case "playercompare":
		h.handlePlayerCompare(w, r)
	case "playeryearbyyearstats":
		h.handlePlayerYearByYearStats(w, r)

	// Common endpoints (expanded)
	case "commonplayerinfov2":
		h.handleCommonPlayerInfoV2(w, r)
	case "commonallplayersv2":
		h.handleCommonAllPlayersV2(w, r)
	case "commonteamrosterv2":
		h.handleCommonTeamRosterV2(w, r)
	case "commonplayoffseries":
		h.handleCommonPlayoffSeries(w, r)
	case "commonteamyears":
		h.handleCommonTeamYears(w, r)

	// Draft & Historical endpoints
	case "drafthistory":
		h.handleDraftHistory(w, r)
	case "draftboard":
		h.handleDraftBoard(w, r)
	case "draftcombinestats":
		h.handleDraftCombineStats(w, r)
	case "franchisehistory":
		h.handleFranchiseHistory(w, r)
	case "franchiseleaders":
		h.handleFranchiseLeaders(w, r)

	// Additional endpoints (iteration 5)
	case "winprobabilitypbp":
		h.handleWinProbabilityPBP(w, r)
	case "infographicfanduelplayer":
		h.handleInfographicFanDuelPlayer(w, r)
	case "homepagev2":
		h.handleHomepageV2(w, r)
	case "homepageleaders":
		h.handleHomepageLeaders(w, r)

	// Advanced analytics endpoints (iteration 7)
	case "scoreboardv3":
		h.handleScoreboardV3(w, r)
	case "playerindex":
		h.handlePlayerIndex(w, r)
	case "alltimeleadersgrids":
		h.handleAllTimeLeadersGrids(w, r)
	case "defensehub":
		h.handleDefenseHub(w, r)
	case "assisttracker":
		h.handleAssistTracker(w, r)
	case "synergyplaytypes":
		h.handleSynergyPlayTypes(w, r)
	case "playercareerbycollege":
		h.handlePlayerCareerByCollege(w, r)
	case "cumestatsplayer":
		h.handleCumeStatsPlayer(w, r)
	case "cumestatsteam":
		h.handleCumeStatsTeam(w, r)
	case "leaguedashoppptshot":
		h.handleLeagueDashOppPtShot(w, r)

	// Final endpoints (iteration 8)
	case "leagueleadersv2":
		h.handleLeagueLeadersV2(w, r)
	case "playergamestreakfinder":
		h.handlePlayerGameStreakFinder(w, r)
	case "teamgamestreakfinder":
		h.handleTeamGameStreakFinder(w, r)
	case "opponentshooting":
		h.handleOpponentShooting(w, r)
	case "shootingefficiency":
		h.handleShootingEfficiency(w, r)
	case "videoevents":
		h.handleVideoEvents(w, r)
	case "matchuprollup":
		h.handleMatchupRollup(w, r)
	case "teamplayeronoffdetails":
		h.handleTeamPlayerOnOffDetails(w, r)
	case "leagueplayerondetails":
		h.handleLeaguePlayerOnDetails(w, r)
	case "assistleaders":
		h.handleAssistLeaders(w, r)
	case "playerestimatedadvancedstats":
		h.handlePlayerEstimatedAdvancedStats(w, r)
	case "leaguehustlestatsteamleaders":
		h.handleLeagueHustleStatsTeamLeaders(w, r)

	// Iteration 9 endpoints
	case "playbyplayv3":
		h.handlePlayByPlayV3(w, r)
	case "boxscorematchupsv3":
		h.handleBoxScoreMatchupsV3(w, r)
	case "shotchartlineupdetail":
		h.handleShotChartLineupDetail(w, r)
	case "playercareerbycollegerollup":
		h.handlePlayerCareerByCollegeRollup(w, r)

	// Iteration 10 endpoints - Final SDK endpoints (beyond 100%)
	case "commonplayoffseriesv2":
		h.handleCommonPlayoffSeriesV2(w, r)
	case "leaguedashplayerclutchv2":
		h.handleLeagueDashPlayerClutchV2(w, r)
	case "leaguedashplayershotlocationv2":
		h.handleLeagueDashPlayerShotLocationV2(w, r)
	case "leaguedashteamclutchv2":
		h.handleLeagueDashTeamClutchV2(w, r)
	case "playernextngames":
		h.handlePlayerNextNGames(w, r)
	case "playertrackingshootingefficiency":
		h.handlePlayerTrackingShootingEfficiency(w, r)
	case "teamandplayersvsplayers":
		h.handleTeamAndPlayersVsPlayers(w, r)
	case "teaminfocommonv2":
		h.handleTeamInfoCommonV2(w, r)
	case "teamnextngames":
		h.handleTeamNextNGames(w, r)
	case "teamyearoveryearsplits":
		h.handleTeamYearOverYearSplits(w, r)
	case "internationalbroadcasterschedule":
		h.handleInternationalBroadcasterSchedule(w, r)

	default:
		writeError(w, http.StatusNotFound, "endpoint_not_found", "Endpoint not supported: "+endpoint)
	}
}

func getQueryOrDefault(r *http.Request, key, defaultValue string) string {
	if value := r.URL.Query().Get(key); value != "" {
		return value
	}
	return defaultValue
}

func stringPtr(s string) *string {
	return &s
}

func leagueIDPtr(id parameters.LeagueID) *parameters.LeagueID {
	return &id
}

func perModePtr(pm parameters.PerMode) *parameters.PerMode {
	return &pm
}

func seasonPtr(s parameters.Season) *parameters.Season {
	return &s
}

func seasonTypePtr(st parameters.SeasonType) *parameters.SeasonType {
	return &st
}

func writeSuccess(w http.ResponseWriter, data interface{}) {
	type successResponse struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	resp := successResponse{
		Success: true,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
