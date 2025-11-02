#!/usr/bin/env python3
"""
Example Python client for nba-api-go HTTP API Server
Demonstrates how to use the API from Python without needing Go
"""

import requests
import json
from typing import Dict, Any, Optional

BASE_URL = "http://localhost:8080/api/v1/stats"

class NBAApiClient:
    """Simple Python client for the nba-api-go HTTP API"""
    
    def __init__(self, base_url: str = BASE_URL):
        self.base_url = base_url
        self.session = requests.Session()
    
    def _get(self, endpoint: str, params: Optional[Dict[str, Any]] = None) -> Dict[str, Any]:
        """Make GET request to API"""
        url = f"{self.base_url}/{endpoint}"
        response = self.session.get(url, params=params)
        response.raise_for_status()
        return response.json()
    
    # Player endpoints
    def get_player_career_stats(self, player_id: str) -> Dict[str, Any]:
        """Get player career statistics"""
        return self._get("playercareerstats", {"PlayerID": player_id})
    
    def get_player_game_log(self, player_id: str, season: str = "2023-24") -> Dict[str, Any]:
        """Get player game log for a season"""
        params = {
            "PlayerID": player_id,
            "Season": season,
            "SeasonType": "Regular Season"
        }
        return self._get("playergamelog", params)
    
    def get_player_info(self, player_id: str) -> Dict[str, Any]:
        """Get player information"""
        return self._get("commonplayerinfo", {"PlayerID": player_id})
    
    # League endpoints
    def get_league_leaders(self, season: str = "2023-24", stat: str = "PTS") -> Dict[str, Any]:
        """Get league leaders for a stat"""
        params = {
            "Season": season,
            "SeasonType": "Regular Season",
            "StatCategory": stat,
            "LeagueID": "00"
        }
        return self._get("leagueleaders", params)
    
    def get_league_standings(self, season: str = "2023-24") -> Dict[str, Any]:
        """Get league standings"""
        params = {
            "Season": season,
            "SeasonType": "Regular Season",
            "LeagueID": "00"
        }
        return self._get("leaguestandings", params)
    
    # Team endpoints
    def get_team_roster(self, team_id: str, season: str = "2023-24") -> Dict[str, Any]:
        """Get team roster"""
        params = {
            "TeamID": team_id,
            "Season": season
        }
        return self._get("commonteamroster", params)
    
    def get_team_game_log(self, team_id: str, season: str = "2023-24") -> Dict[str, Any]:
        """Get team game log"""
        params = {
            "TeamID": team_id,
            "Season": season,
            "SeasonType": "Regular Season"
        }
        return self._get("teamgamelog", params)
    
    # Box scores
    def get_box_score_traditional(self, game_id: str) -> Dict[str, Any]:
        """Get traditional box score"""
        return self._get("boxscoretraditionalv2", {"GameID": game_id})
    
    def get_box_score_advanced(self, game_id: str) -> Dict[str, Any]:
        """Get advanced box score"""
        return self._get("boxscoreadvancedv2", {"GameID": game_id})
    
    # Scoreboard
    def get_scoreboard(self, game_date: str) -> Dict[str, Any]:
        """Get scoreboard for a date (YYYY-MM-DD)"""
        return self._get("scoreboardv2", {"GameDate": game_date})


def main():
    """Example usage of the NBA API client"""
    
    print("=" * 80)
    print("NBA API Go - Python Client Example")
    print("=" * 80)
    print()
    
    client = NBAApiClient()
    
    # Example 1: Player Career Stats
    print("1. Getting Nikola Jokić's career stats...")
    try:
        career = client.get_player_career_stats("203999")
        seasons = career.get("data", {}).get("SeasonTotalsRegularSeason", [])
        print(f"   Found {len(seasons)} seasons")
        if seasons:
            latest = seasons[-1]
            print(f"   Latest: {latest.get('SEASON_ID')} - {latest.get('PTS'):.1f} PPG")
    except Exception as e:
        print(f"   Error: {e}")
    print()
    
    # Example 2: League Leaders
    print("2. Getting scoring leaders...")
    try:
        leaders = client.get_league_leaders(season="2023-24", stat="PTS")
        top_scorers = leaders.get("data", {}).get("LeagueLeaders", [])[:5]
        print(f"   Top 5 Scorers:")
        for i, player in enumerate(top_scorers, 1):
            print(f"   {i}. {player.get('PLAYER')} - {player.get('PTS'):.1f} PPG")
    except Exception as e:
        print(f"   Error: {e}")
    print()
    
    # Example 3: Team Roster
    print("3. Getting Lakers roster...")
    try:
        roster = client.get_team_roster("1610612747", season="2023-24")
        players = roster.get("data", {}).get("CommonTeamRoster", [])
        print(f"   Found {len(players)} players")
        if players:
            print("   Sample players:")
            for player in players[:3]:
                print(f"   - {player.get('PLAYER')}")
    except Exception as e:
        print(f"   Error: {e}")
    print()
    
    # Example 4: Player Game Log
    print("4. Getting LeBron James game log...")
    try:
        games = client.get_player_game_log("2544", season="2023-24")
        game_log = games.get("data", {}).get("PlayerGameLog", [])[:5]
        print(f"   Last 5 games:")
        for game in game_log:
            print(f"   {game.get('GAME_DATE')} vs {game.get('MATCHUP')}: {game.get('PTS')} pts")
    except Exception as e:
        print(f"   Error: {e}")
    print()
    
    print("=" * 80)
    print("Examples complete! All 139 endpoints available via this pattern.")
    print("=" * 80)


if __name__ == "__main__":
    main()
