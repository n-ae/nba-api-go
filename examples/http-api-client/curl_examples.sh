#!/bin/bash
# Simple curl examples for testing the nba-api-go HTTP API

BASE_URL="http://localhost:8080"

echo "========================================"
echo "NBA API Go - cURL Examples"
echo "========================================"
echo ""

# Check server health
echo "1. Health Check:"
curl -s "${BASE_URL}/health" | jq '.'
echo ""

# Get all players
echo "2. Get All Players (first 5):"
curl -s "${BASE_URL}/api/v1/stats/commonallplayers" | jq '.data.CommonAllPlayers[0:5]'
echo ""

# Get player career stats
echo "3. Nikola Jokić Career Stats:"
curl -s "${BASE_URL}/api/v1/stats/playercareerstats?PlayerID=203999" | \
    jq '.data.SeasonTotalsRegularSeason[-1] | {season: .SEASON_ID, ppg: .PTS, rpg: .REB, apg: .AST}'
echo ""

# Get league leaders
echo "4. Top 5 Scorers 2023-24:"
curl -s "${BASE_URL}/api/v1/stats/leagueleaders?Season=2023-24&SeasonType=Regular+Season&StatCategory=PTS" | \
    jq '.data.LeagueLeaders[0:5] | .[] | {rank: .RANK, player: .PLAYER, ppg: .PTS}'
echo ""

# Get team roster
echo "5. Lakers Roster:"
curl -s "${BASE_URL}/api/v1/stats/commonteamroster?TeamID=1610612747&Season=2023-24" | \
    jq '.data.CommonTeamRoster[0:5] | .[] | {name: .PLAYER, number: .NUM, position: .POSITION}'
echo ""

echo "========================================"
echo "Examples complete!"
echo "All 139 endpoints available via:"
echo "  ${BASE_URL}/api/v1/stats/{endpoint}"
echo "========================================"
