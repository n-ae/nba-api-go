package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/username/nba-api-go/pkg/stats"
	"github.com/username/nba-api-go/pkg/stats/endpoints"
	"github.com/username/nba-api-go/pkg/stats/parameters"
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

	default:
		writeError(w, http.StatusNotFound, "endpoint_not_found", "Endpoint not supported: "+endpoint)
	}
}

func (h *StatsHandler) handlePlayerGameLog(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerGameLogRequest{
		PlayerID:   playerID,
		Season:     season,
		SeasonType: seasonType,
		LeagueID:   parameters.LeagueIDNBA,
	}

	resp, err := endpoints.PlayerGameLog(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleCommonAllPlayers(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	leagueID := parameters.LeagueIDNBA
	isOnlyCurrent := getQueryOrDefault(r, "IsOnlyCurrentSeason", "0")

	req := endpoints.CommonAllPlayersRequest{
		Season:              season,
		LeagueID:            &leagueID,
		IsOnlyCurrentSeason: &isOnlyCurrent,
	}

	resp, err := endpoints.GetCommonAllPlayers(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleScoreboardV2(w http.ResponseWriter, r *http.Request) {
	gameDate := r.URL.Query().Get("GameDate")
	if gameDate == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameDate is required (format: YYYY-MM-DD)")
		return
	}

	leagueID := parameters.LeagueIDNBA
	req := endpoints.ScoreboardV2Request{
		GameDate: gameDate,
		LeagueID: &leagueID,
	}

	resp, err := endpoints.GetScoreboardV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueStandings(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueStandingsRequest{
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueStandings(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleCommonTeamRoster(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.CommonTeamRosterRequest{
		TeamID:   teamID,
		Season:   &season,
		LeagueID: &leagueID,
	}

	resp, err := endpoints.GetCommonTeamRoster(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handlePlayerCareerStats(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerCareerStatsRequest{
		PlayerID: playerID,
		PerMode:  perMode,
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := endpoints.PlayerCareerStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueLeadersRequest{
		Season:     season,
		SeasonType: seasonType,
		PerMode:    perMode,
		LeagueID:   parameters.LeagueIDNBA,
	}

	resp, err := endpoints.LeagueLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleCommonPlayerInfo(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.CommonPlayerInfoRequest{
		PlayerID: playerID,
		LeagueID: parameters.LeagueIDNBA,
	}

	resp, err := endpoints.CommonPlayerInfo(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueDashTeamStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashTeamStatsRequest{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashTeamStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

func (h *StatsHandler) handleLeagueDashPlayerStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashPlayerStatsRequest{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp.Data)
}

// Player endpoint handlers (expanded)

func (h *StatsHandler) handlePlayerProfileV2(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerProfileV2Request{
		PlayerID: playerID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerProfileV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerAwards(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerAwardsRequest{
		PlayerID: playerID,
	}

	resp, err := endpoints.GetPlayerAwards(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByGeneralSplits(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByGeneralSplitsRequest{
		PlayerID:    playerID,
		Season:      season,
		SeasonType:  seasonType,
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByGeneralSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByShootingSplits(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByShootingSplitsRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByShootingSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByOpponent(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByOpponentRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByOpponent(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByClutch(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByClutchRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerGameLogs(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerGameLogsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerGameLogs(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerVsPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	vsPlayerID := r.URL.Query().Get("VsPlayerID")

	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}
	if vsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "VsPlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerVsPlayerRequest{
		PlayerID:   playerID,
		VsPlayerID: vsPlayerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(parameters.PerModePerGame),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerVsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoint handlers (expanded)

func (h *StatsHandler) handleTeamGameLog(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamGameLogRequest{
		TeamID:     teamID,
		Season:     season,
		SeasonType: seasonType,
	}

	resp, err := endpoints.GetTeamGameLog(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamInfoCommon(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamInfoCommonRequest{
		TeamID:     teamID,
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamInfoCommon(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashboardByGeneralSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByGeneralSplitsRequest{
		TeamID:      teamID,
		Season:      season,
		SeasonType:  seasonType,
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamDashboardByGeneralSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashboardByShootingSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByShootingSplitsRequest{
		TeamID:      teamID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamDashboardByShootingSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashboardByOpponent(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByOpponentRequest{
		TeamID:      teamID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamDashboardByOpponent(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDetails(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	req := endpoints.TeamDetailsRequest{
		TeamID: teamID,
	}

	resp, err := endpoints.GetTeamDetails(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamPlayerDashboard(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamPlayerDashboardRequest{
		TeamID:      teamID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
	}

	resp, err := endpoints.GetTeamPlayerDashboard(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamLineups(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamLineupsRequest{
		TeamID:        teamID,
		Season:        seasonPtr(season),
		SeasonType:    seasonTypePtr(seasonType),
		PerMode:       perModePtr(perMode),
		MeasureType:   stringPtr("Base"),
		GroupQuantity: stringPtr("5"),
	}

	resp, err := endpoints.GetTeamLineups(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Box Score endpoint handlers

func (h *StatsHandler) handleBoxScoreSummaryV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreSummaryV2Request{
		GameID: gameID,
	}

	resp, err := endpoints.GetBoxScoreSummaryV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreTraditionalV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreTraditionalV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreTraditionalV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreAdvancedV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreAdvancedV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreAdvancedV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreScoringV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreScoringV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreScoringV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreMiscV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreMiscV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreMiscV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreUsageV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreUsageV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreUsageV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreFourFactorsV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreFourFactorsV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreFourFactorsV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScorePlayerTrackV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScorePlayerTrackV2Request{
		GameID: gameID,
	}

	resp, err := endpoints.GetBoxScorePlayerTrackV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Player Tracking endpoint handlers

func (h *StatsHandler) handlePlayerTrackingShotDashboard(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingShootingEfficiencyRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingShootingEfficiency(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingPasses(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPassesRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPasses(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingDefense(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingDefenseRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingDefense(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingRebounding(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingReboundingRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingRebounding(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingSpeedDistance(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingSpeedDistanceRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingSpeedDistance(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingCatchShoot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingCatchShootRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingCatchShoot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingDrives(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingDrivesRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingDrives(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreDefensiveV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreDefensiveV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetBoxScoreDefensiveV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreHustleV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreHustleV2Request{
		GameID: gameID,
	}

	resp, err := endpoints.GetBoxScoreHustleV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Game endpoint handlers

func (h *StatsHandler) handlePlayByPlayV2(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.PlayByPlayV2Request{
		GameID:      gameID,
		StartPeriod: stringPtr("0"),
		EndPeriod:   stringPtr("10"),
	}

	resp, err := endpoints.GetPlayByPlayV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleShotChartDetail(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	teamID := getQueryOrDefault(r, "TeamID", "0")
	playerID := getQueryOrDefault(r, "PlayerID", "0")

	req := endpoints.ShotChartDetailRequest{
		Season:     season,
		SeasonType: seasonType,
		TeamID:     stringPtr(teamID),
		PlayerID:   stringPtr(playerID),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetShotChartDetail(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleGameRotation(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.GameRotationRequest{
		GameID:   gameID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetGameRotation(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// League endpoint handlers

func (h *StatsHandler) handleLeagueGameLog(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueGameLogRequest{
		Season:       season,
		SeasonType:   seasonTypePtr(seasonType),
		PlayerOrTeam: stringPtr("P"),
		LeagueID:     leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueGameLog(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayoffPicture(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.PlayoffPictureRequest{
		SeasonID: season,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayoffPicture(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashLineups(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashLineupsRequest{
		Season:        seasonPtr(season),
		SeasonType:    seasonTypePtr(seasonType),
		PerMode:       perModePtr(perMode),
		MeasureType:   stringPtr("Base"),
		GroupQuantity: stringPtr("5"),
		LeagueID:      leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashLineups(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Additional League endpoint handlers

func (h *StatsHandler) handleLeagueDashPlayerClutch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPlayerClutchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamClutch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashTeamClutchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPlayerBioStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueDashPlayerBioStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerBioStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamBioStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueDashTeamBioStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamBioStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPtStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPtStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPtStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueHustleStatsPlayer(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueHustleStatsPlayerRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueHustleStatsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueHustleStatsTeam(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueHustleStatsTeamRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueHustleStatsTeam(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPtDefend(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPtDefendRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPtDefend(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueGameFinder(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueGameFinderRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueGameFinder(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueStandingsV3(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.LeagueStandingsV3Request{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueStandingsV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Additional Player endpoint handlers

func (h *StatsHandler) handlePlayerEstimatedMetrics(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerEstimatedMetricsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerEstimatedMetrics(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerFantasyProfile(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerFantasyProfileRequest{
		PlayerID: playerID,
	}

	resp, err := endpoints.GetPlayerFantasyProfile(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashPtShots(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerDashPtShotsRequest{
		PlayerID:   playerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashPtShots(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByLastNGames(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByLastNGamesRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByLastNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByTeamPerformance(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByTeamPerformanceRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByTeamPerformance(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerDashboardByGameSplits(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByGameSplitsRequest{
		PlayerID:    playerID,
		Season:      seasonPtr(season),
		SeasonType:  seasonTypePtr(seasonType),
		PerMode:     perModePtr(perMode),
		MeasureType: stringPtr("Base"),
		LeagueID:    leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByGameSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoints (iteration 4)

func (h *StatsHandler) handleTeamGameLogs(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	teamID := r.URL.Query().Get("TeamID")

	req := endpoints.TeamGameLogsRequest{
		Season:     season,
		SeasonType: seasonType,
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}
	if teamID != "" {
		req.TeamID = stringPtr(teamID)
	}

	resp, err := endpoints.GetTeamGameLogs(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamYearByYearStats(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamYearByYearStatsRequest{
		TeamID:     teamID,
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamYearByYearStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamVsTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	vsTeamID := r.URL.Query().Get("VsTeamID")
	if teamID == "" || vsTeamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID and VsTeamID are required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamVsTeamRequest{
		TeamID:     teamID,
		VsTeamID:   vsTeamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamVsTeam(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamHistoricalLeaders(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	req := endpoints.TeamHistoricalLeadersRequest{
		TeamID:   teamID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamHistoricalLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamEstimatedMetrics(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamEstimatedMetricsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamEstimatedMetrics(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashPtShots(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamDashPtShotsRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashPtShots(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Player endpoints (iteration 4)

func (h *StatsHandler) handlePlayerDashboardByYearOverYear(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerDashboardByYearOverYearRequest{
		PlayerID:   playerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerDashboardByYearOverYear(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerCompare(w http.ResponseWriter, r *http.Request) {
	playerIDList := r.URL.Query().Get("PlayerIDList")
	if playerIDList == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerIDList is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerCompareRequest{
		PlayerIDList: playerIDList,
		Season:       seasonPtr(season),
		SeasonType:   seasonTypePtr(seasonType),
		LeagueID:     leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerCompare(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerYearByYearStats(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.PlayerYearByYearStatsRequest{
		PlayerID: playerID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerYearByYearStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Common endpoints (iteration 4)

func (h *StatsHandler) handleCommonPlayerInfoV2(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.CommonPlayerInfoV2Request{
		PlayerID: playerID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonPlayerInfoV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleCommonAllPlayersV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.CommonAllPlayersV2Request{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonAllPlayersV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleCommonTeamRosterV2(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.CommonTeamRosterV2Request{
		TeamID:   teamID,
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonTeamRosterV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleCommonPlayoffSeries(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.CommonPlayoffSeriesRequest{
		Season:   season,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonPlayoffSeries(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleCommonTeamYears(w http.ResponseWriter, r *http.Request) {
	req := endpoints.CommonTeamYearsRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCommonTeamYears(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Draft & Historical endpoints (iteration 4)

func (h *StatsHandler) handleDraftHistory(w http.ResponseWriter, r *http.Request) {
	req := endpoints.DraftHistoryRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDraftHistory(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleDraftBoard(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.DraftBoardRequest{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDraftBoard(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleDraftCombineStats(w http.ResponseWriter, r *http.Request) {
	seasonYear := getQueryOrDefault(r, "SeasonYear", "2023-24")

	req := endpoints.DraftCombineStatsRequest{
		SeasonYear: stringPtr(seasonYear),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDraftCombineStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleFranchiseHistory(w http.ResponseWriter, r *http.Request) {
	req := endpoints.FranchiseHistoryRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetFranchiseHistory(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleFranchiseLeaders(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	req := endpoints.FranchiseLeadersRequest{
		TeamID:   teamID,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetFranchiseLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// League endpoints (iteration 5)

func (h *StatsHandler) handleLeagueDashPlayerShotLocations(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPlayerShotLocationsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerShotLocations(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamShotLocations(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashTeamShotLocationsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamShotLocations(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueSeasonMatchups(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.LeagueSeasonMatchupsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueSeasonMatchups(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPtTeamDefend(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPtTeamDefendRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPtTeamDefend(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoints (iteration 5)

func (h *StatsHandler) handleTeamDashboardByClutch(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByClutchRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByClutch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashboardByLastNGames(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByLastNGamesRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByLastNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashboardByYearOverYear(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByYearOverYearRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByYearOverYear(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamVsPlayer(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	vsPlayerID := r.URL.Query().Get("VsPlayerID")
	if teamID == "" || vsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID and VsPlayerID are required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamVsPlayerRequest{
		TeamID:     teamID,
		VsPlayerID: vsPlayerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamVsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Additional endpoints (iteration 5)

func (h *StatsHandler) handleWinProbabilityPBP(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.WinProbabilityPBPRequest{
		GameID: gameID,
	}

	resp, err := endpoints.GetWinProbabilityPBP(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleInfographicFanDuelPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	req := endpoints.InfographicFanDuelPlayerRequest{
		PlayerID: playerID,
	}

	resp, err := endpoints.GetInfographicFanDuelPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleHomepageV2(w http.ResponseWriter, r *http.Request) {
	req := endpoints.HomepageV2Request{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetHomepageV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleHomepageLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.HomepageLeadersRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetHomepageLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Player Tracking endpoints (iteration 6)

func (h *StatsHandler) handlePlayerTrackingPostTouch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPostTouchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPostTouch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingPaintTouch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPaintTouchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPaintTouch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingElbowTouch(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingElbowTouchRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingElbowTouch(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingPullUpShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerTrackingPullUpShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerTrackingPullUpShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Team endpoints (iteration 6)

func (h *StatsHandler) handleTeamDashboardByGameSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByGameSplitsRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByGameSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamDashboardByTeamPerformance(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamDashboardByTeamPerformanceRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamDashboardByTeamPerformance(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamPlayerOnOffSummary(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamPlayerOnOffSummaryRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamPlayerOnOffSummary(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// League endpoints (iteration 6)

func (h *StatsHandler) handleLeagueDashPlayerPtShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashPlayerPtShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashPlayerPtShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamPtShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashTeamPtShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashTeamPtShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Advanced analytics endpoints (iteration 7)

func (h *StatsHandler) handleScoreboardV3(w http.ResponseWriter, r *http.Request) {
	gameDate := getQueryOrDefault(r, "GameDate", "")

	req := endpoints.ScoreboardV3Request{
		GameDate: gameDate,
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetScoreboardV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerIndex(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))

	req := endpoints.PlayerIndexRequest{
		Season:   seasonPtr(season),
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerIndex(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleAllTimeLeadersGrids(w http.ResponseWriter, r *http.Request) {
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.AllTimeLeadersGridsRequest{
		PerMode:    perModePtr(perMode),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetAllTimeLeadersGrids(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleDefenseHub(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.DefenseHubRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetDefenseHub(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleAssistTracker(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.AssistTrackerRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetAssistTracker(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleSynergyPlayTypes(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.SynergyPlayTypesRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetSynergyPlayTypes(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerCareerByCollege(w http.ResponseWriter, r *http.Request) {
	college := r.URL.Query().Get("College")

	req := endpoints.PlayerCareerByCollegeRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
	}
	if college != "" {
		req.College = stringPtr(college)
	}

	resp, err := endpoints.GetPlayerCareerByCollege(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleCumeStatsPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.CumeStatsPlayerRequest{
		PlayerID:   playerID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCumeStatsPlayer(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleCumeStatsTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.CumeStatsTeamRequest{
		TeamID:     teamID,
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetCumeStatsTeam(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashOppPtShot(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueDashOppPtShotRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueDashOppPtShot(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Final endpoints (iteration 8)

func (h *StatsHandler) handleLeagueLeadersV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueLeadersV2Request{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueLeadersV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerGameStreakFinder(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerGameStreakFinderRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerGameStreakFinder(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamGameStreakFinder(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.TeamGameStreakFinderRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetTeamGameStreakFinder(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleOpponentShooting(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.OpponentShootingRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetOpponentShooting(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleShootingEfficiency(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.ShootingEfficiencyRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetShootingEfficiency(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleVideoEvents(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	gameEventID := r.URL.Query().Get("GameEventID")
	if gameID == "" || gameEventID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID and GameEventID are required")
		return
	}

	req := endpoints.VideoEventsRequest{
		GameID:      gameID,
		GameEventID: gameEventID,
	}

	resp, err := endpoints.GetVideoEvents(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleMatchupRollup(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.MatchupRollupRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetMatchupRollup(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamPlayerOnOffDetails(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.TeamPlayerOnOffDetailsRequest{
		TeamID:  teamID,
		PerMode: perModePtr(perMode),
	}

	resp, err := endpoints.GetTeamPlayerOnOffDetails(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeaguePlayerOnDetails(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeaguePlayerOnDetailsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeaguePlayerOnDetails(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleAssistLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.AssistLeadersRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetAssistLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerEstimatedAdvancedStats(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))

	req := endpoints.PlayerEstimatedAdvancedStatsRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetPlayerEstimatedAdvancedStats(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueHustleStatsTeamLeaders(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.LeagueHustleStatsTeamLeadersRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		PerMode:    perModePtr(perMode),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	resp, err := endpoints.GetLeagueHustleStatsTeamLeaders(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Iteration 9 endpoints

func (h *StatsHandler) handlePlayByPlayV3(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.PlayByPlayV3Request{
		GameID: gameID,
	}

	startPeriod := r.URL.Query().Get("StartPeriod")
	if startPeriod != "" {
		req.StartPeriod = &startPeriod
	}

	endPeriod := r.URL.Query().Get("EndPeriod")
	if endPeriod != "" {
		req.EndPeriod = &endPeriod
	}

	resp, err := endpoints.GetPlayByPlayV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleBoxScoreMatchupsV3(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("GameID")
	if gameID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "GameID is required")
		return
	}

	req := endpoints.BoxScoreMatchupsV3Request{
		GameID: gameID,
	}

	startPeriod := r.URL.Query().Get("StartPeriod")
	if startPeriod != "" {
		req.StartPeriod = &startPeriod
	}

	endPeriod := r.URL.Query().Get("EndPeriod")
	if endPeriod != "" {
		req.EndPeriod = &endPeriod
	}

	resp, err := endpoints.GetBoxScoreMatchupsV3(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleShotChartLineupDetail(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	teamID := r.URL.Query().Get("TeamID")
	groupID := r.URL.Query().Get("GroupID")

	req := endpoints.ShotChartLineupDetailRequest{
		Season:     seasonPtr(season),
		SeasonType: seasonTypePtr(seasonType),
		LeagueID:   leagueIDPtr(parameters.LeagueIDNBA),
	}

	if teamID != "" {
		req.TeamID = &teamID
	}
	if groupID != "" {
		req.GroupID = &groupID
	}

	contextMeasure := r.URL.Query().Get("ContextMeasure")
	if contextMeasure != "" {
		req.ContextMeasure = &contextMeasure
	}

	resp, err := endpoints.GetShotChartLineupDetail(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerCareerByCollegeRollup(w http.ResponseWriter, r *http.Request) {
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))

	req := endpoints.PlayerCareerByCollegeRollupRequest{
		LeagueID: leagueIDPtr(parameters.LeagueIDNBA),
		PerMode:  perModePtr(perMode),
	}

	resp, err := endpoints.GetPlayerCareerByCollegeRollup(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Iteration 10 endpoints - Final SDK endpoints

func (h *StatsHandler) handleCommonPlayoffSeriesV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.CommonPlayoffSeriesV2Request{
		Season:   &season,
		LeagueID: &leagueID,
	}

	resp, err := endpoints.GetCommonPlayoffSeriesV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPlayerClutchV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashPlayerClutchV2Request{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerClutchV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashPlayerShotLocationV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashPlayerShotLocationV2Request{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashPlayerShotLocationV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleLeagueDashTeamClutchV2(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.LeagueDashTeamClutchV2Request{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetLeagueDashTeamClutchV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerNextNGames(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("PlayerID")
	if playerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.PlayerNextNGamesRequest{
		PlayerID:   playerID,
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetPlayerNextNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handlePlayerTrackingShootingEfficiency(w http.ResponseWriter, r *http.Request) {
	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.PlayerTrackingShootingEfficiencyRequest{
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetPlayerTrackingShootingEfficiency(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamAndPlayersVsPlayers(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	vsPlayerID := r.URL.Query().Get("VsPlayerID")
	if teamID == "" || vsPlayerID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID and VsPlayerID are required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamAndPlayersVsPlayersRequest{
		TeamID:     teamID,
		VsPlayerID: vsPlayerID,
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamAndPlayersVsPlayers(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamInfoCommonV2(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamInfoCommonV2Request{
		TeamID:     teamID,
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamInfoCommonV2(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamNextNGames(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamNextNGamesRequest{
		TeamID:     teamID,
		Season:     &season,
		SeasonType: &seasonType,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamNextNGames(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

func (h *StatsHandler) handleTeamYearOverYearSplits(w http.ResponseWriter, r *http.Request) {
	teamID := r.URL.Query().Get("TeamID")
	if teamID == "" {
		writeError(w, http.StatusBadRequest, "missing_parameter", "TeamID is required")
		return
	}

	season := parameters.Season(getQueryOrDefault(r, "Season", "2023-24"))
	seasonType := parameters.SeasonType(getQueryOrDefault(r, "SeasonType", "Regular Season"))
	perMode := parameters.PerMode(getQueryOrDefault(r, "PerMode", "PerGame"))
	leagueID := parameters.LeagueIDNBA

	req := endpoints.TeamYearOverYearSplitsRequest{
		TeamID:     teamID,
		Season:     &season,
		SeasonType: &seasonType,
		PerMode:    &perMode,
		LeagueID:   &leagueID,
	}

	resp, err := endpoints.GetTeamYearOverYearSplits(r.Context(), h.client, req)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "api_error", err.Error())
		return
	}

	writeSuccess(w, resp)
}

// Helper functions

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
