#!/usr/bin/env node
/**
 * Example Node.js client for nba-api-go HTTP API Server
 * Demonstrates how to use the API from JavaScript without needing Go
 */

const BASE_URL = 'http://localhost:8080/api/v1/stats';

class NBAApiClient {
    constructor(baseUrl = BASE_URL) {
        this.baseUrl = baseUrl;
    }

    async _get(endpoint, params = {}) {
        const url = new URL(`${this.baseUrl}/${endpoint}`);
        Object.keys(params).forEach(key => url.searchParams.append(key, params[key]));
        
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`HTTP ${response.status}: ${response.statusText}`);
        }
        return response.json();
    }

    // Player endpoints
    async getPlayerCareerStats(playerId) {
        return this._get('playercareerstats', { PlayerID: playerId });
    }

    async getPlayerGameLog(playerId, season = '2023-24') {
        return this._get('playergamelog', {
            PlayerID: playerId,
            Season: season,
            SeasonType: 'Regular Season'
        });
    }

    async getPlayerInfo(playerId) {
        return this._get('commonplayerinfo', { PlayerID: playerId });
    }

    // League endpoints
    async getLeagueLeaders(season = '2023-24', stat = 'PTS') {
        return this._get('leagueleaders', {
            Season: season,
            SeasonType: 'Regular Season',
            StatCategory: stat,
            LeagueID: '00'
        });
    }

    async getLeagueStandings(season = '2023-24') {
        return this._get('leaguestandings', {
            Season: season,
            SeasonType: 'Regular Season',
            LeagueID: '00'
        });
    }

    // Team endpoints
    async getTeamRoster(teamId, season = '2023-24') {
        return this._get('commonteamroster', {
            TeamID: teamId,
            Season: season
        });
    }

    async getTeamGameLog(teamId, season = '2023-24') {
        return this._get('teamgamelog', {
            TeamID: teamId,
            Season: season,
            SeasonType: 'Regular Season'
        });
    }

    // Box scores
    async getBoxScoreTraditional(gameId) {
        return this._get('boxscoretraditionalv2', { GameID: gameId });
    }

    async getBoxScoreAdvanced(gameId) {
        return this._get('boxscoreadvancedv2', { GameID: gameId });
    }

    // Scoreboard
    async getScoreboard(gameDate) {
        return this._get('scoreboardv2', { GameDate: gameDate });
    }
}

async function main() {
    console.log('='.repeat(80));
    console.log('NBA API Go - JavaScript Client Example');
    console.log('='.repeat(80));
    console.log();

    const client = new NBAApiClient();

    // Example 1: Player Career Stats
    console.log("1. Getting Nikola Jokić's career stats...");
    try {
        const career = await client.getPlayerCareerStats('203999');
        const seasons = career.data?.SeasonTotalsRegularSeason || [];
        console.log(`   Found ${seasons.length} seasons`);
        if (seasons.length > 0) {
            const latest = seasons[seasons.length - 1];
            console.log(`   Latest: ${latest.SEASON_ID} - ${latest.PTS.toFixed(1)} PPG`);
        }
    } catch (error) {
        console.log(`   Error: ${error.message}`);
    }
    console.log();

    // Example 2: League Leaders
    console.log('2. Getting scoring leaders...');
    try {
        const leaders = await client.getLeagueLeaders('2023-24', 'PTS');
        const topScorers = (leaders.data?.LeagueLeaders || []).slice(0, 5);
        console.log('   Top 5 Scorers:');
        topScorers.forEach((player, i) => {
            console.log(`   ${i + 1}. ${player.PLAYER} - ${player.PTS.toFixed(1)} PPG`);
        });
    } catch (error) {
        console.log(`   Error: ${error.message}`);
    }
    console.log();

    // Example 3: Team Roster
    console.log('3. Getting Lakers roster...');
    try {
        const roster = await client.getTeamRoster('1610612747', '2023-24');
        const players = roster.data?.CommonTeamRoster || [];
        console.log(`   Found ${players.length} players`);
        if (players.length > 0) {
            console.log('   Sample players:');
            players.slice(0, 3).forEach(player => {
                console.log(`   - ${player.PLAYER}`);
            });
        }
    } catch (error) {
        console.log(`   Error: ${error.message}`);
    }
    console.log();

    // Example 4: Player Game Log
    console.log('4. Getting LeBron James game log...');
    try {
        const games = await client.getPlayerGameLog('2544', '2023-24');
        const gameLog = (games.data?.PlayerGameLog || []).slice(0, 5);
        console.log('   Last 5 games:');
        gameLog.forEach(game => {
            console.log(`   ${game.GAME_DATE} vs ${game.MATCHUP}: ${game.PTS} pts`);
        });
    } catch (error) {
        console.log(`   Error: ${error.message}`);
    }
    console.log();

    console.log('='.repeat(80));
    console.log('Examples complete! All 139 endpoints available via this pattern.');
    console.log('='.repeat(80));
}

// Run if called directly
if (require.main === module) {
    main().catch(console.error);
}

module.exports = NBAApiClient;
