# HTTP API Client Examples

This directory contains example clients for the nba-api-go HTTP API server in multiple languages.

## Overview

The nba-api-go HTTP API server exposes all 139 NBA statistics endpoints via REST, making them accessible from any programming language. These examples demonstrate how to use the API without needing to write Go code.

## Available Examples

### Python (`python_example.py`)
Complete Python client using the `requests` library.

**Requirements:**
```bash
pip install requests
```

**Usage:**
```bash
# Start the API server first
./bin/nba-api-server

# Then run the example
python3 examples/http-api-client/python_example.py
```

### JavaScript (`javascript_example.js`)
Node.js client using native `fetch` API.

**Requirements:**
- Node.js 18+ (for native fetch support)

**Usage:**
```bash
# Start the API server first
./bin/nba-api-server

# Then run the example
node examples/http-api-client/javascript_example.js
```

### Bash (`curl_examples.sh`)
Simple curl examples for testing from command line.

**Usage:**
```bash
# Start the API server first
./bin/nba-api-server

# Then run examples
bash examples/http-api-client/curl_examples.sh
```

## API Endpoints

All 139 endpoints follow the pattern:
```
GET http://localhost:8080/api/v1/stats/{endpoint}?param1=value1&param2=value2
```

### Common Parameters

- `PlayerID`: Player identifier (e.g., "203999" for Nikola Jokić)
- `TeamID`: Team identifier (e.g., "1610612747" for Lakers)
- `Season`: Season in format "YYYY-YY" (e.g., "2023-24")
- `SeasonType`: "Regular Season", "Playoffs", or "All Star"
- `GameID`: Game identifier
- `LeagueID`: League identifier (usually "00" for NBA)

### Example Endpoints

**Player Endpoints (35 total):**
- `/playercareerstats?PlayerID=203999`
- `/playergamelog?PlayerID=2544&Season=2023-24`
- `/commonplayerinfo?PlayerID=203999`
- `/playerprofilev2?PlayerID=2544`
- `/playerdashboardbygeneralsplits?PlayerID=203999`

**Team Endpoints (30 total):**
- `/commonteamroster?TeamID=1610612747&Season=2023-24`
- `/teamgamelog?TeamID=1610612747&Season=2023-24`
- `/teaminfocommon?TeamID=1610612747`
- `/teamdashboardbygeneralsplits?TeamID=1610612747`

**League Endpoints (28 total):**
- `/leagueleaders?Season=2023-24&StatCategory=PTS`
- `/leaguestandings?Season=2023-24`
- `/leaguedashteamstats?Season=2023-24`
- `/leaguedashplayerstats?Season=2023-24`

**Box Score Endpoints (10 total):**
- `/boxscoresummaryv2?GameID=0022300001`
- `/boxscoretraditionalv2?GameID=0022300001`
- `/boxscoreadvancedv2?GameID=0022300001`
- `/boxscorescoringv2?GameID=0022300001`

**Game Endpoints (12 total):**
- `/playbyplayv2?GameID=0022300001`
- `/playbyplayv3?GameID=0022300001`
- `/shotchartdetail?GameID=0022300001`
- `/gamerotation?GameID=0022300001`

**Other Endpoints (24 total):**
- `/scoreboardv2?GameDate=2024-01-15`
- `/commonallplayers`
- `/drafthistory?LeagueID=00`
- `/franchisehistory`

See [Migration Guide](../../docs/MIGRATION_GUIDE.md) for complete endpoint list and detailed examples.

## Response Format

All endpoints return JSON in this format:

```json
{
  "data": {
    "ResultSetName": [
      { "field1": "value1", "field2": "value2" }
    ]
  },
  "metadata": {
    "endpoint": "playercareerstats",
    "timestamp": 1234567890
  }
}
```

## Error Handling

Errors return appropriate HTTP status codes:

- `400 Bad Request`: Missing or invalid parameters
- `404 Not Found`: Endpoint not found
- `500 Internal Server Error`: API error or upstream failure

Error response format:
```json
{
  "error": {
    "code": "missing_parameter",
    "message": "PlayerID is required"
  }
}
```

## Client Implementation Pattern

### Python Pattern
```python
import requests

BASE_URL = "http://localhost:8080/api/v1/stats"

def get_player_stats(player_id):
    url = f"{BASE_URL}/playercareerstats"
    params = {"PlayerID": player_id}
    response = requests.get(url, params=params)
    response.raise_for_status()
    return response.json()

data = get_player_stats("203999")
seasons = data["data"]["SeasonTotalsRegularSeason"]
```

### JavaScript Pattern
```javascript
const BASE_URL = 'http://localhost:8080/api/v1/stats';

async function getPlayerStats(playerId) {
    const url = `${BASE_URL}/playercareerstats?PlayerID=${playerId}`;
    const response = await fetch(url);
    if (!response.ok) throw new Error(`HTTP ${response.status}`);
    return response.json();
}

const data = await getPlayerStats('203999');
const seasons = data.data.SeasonTotalsRegularSeason;
```

### Curl Pattern
```bash
curl "http://localhost:8080/api/v1/stats/playercareerstats?PlayerID=203999"
```

## Testing

Start the server and test the health endpoint:

```bash
# Start server
./bin/nba-api-server &

# Check health
curl http://localhost:8080/health

# Test an endpoint
curl "http://localhost:8080/api/v1/stats/commonallplayers"
```

## Performance

The HTTP API is optimized for:
- **Low Latency**: ~10-50ms response times
- **Rate Limiting**: Respects NBA API limits (3 req/sec)
- **Concurrent Requests**: Handles 100+ concurrent connections
- **Memory Efficiency**: ~20MB resident memory

## Full Documentation

- [Migration Guide](../../docs/MIGRATION_GUIDE.md) - Complete Python/JavaScript examples
- [API Usage Guide](../../docs/API_USAGE.md) - Detailed HTTP API documentation
- [100% Coverage Achievement](../../HTTP_API_100_PERCENT_ACHIEVEMENT.md) - Complete endpoint list

## Common Use Cases

### Get Player Statistics
```python
# Python
career = client.get_player_career_stats("203999")  # Jokić
games = client.get_player_game_log("2544", "2023-24")  # LeBron
```

### Get League Data
```python
# Python
leaders = client.get_league_leaders(season="2023-24", stat="PTS")
standings = client.get_league_standings(season="2023-24")
```

### Get Team Information
```python
# Python
roster = client.get_team_roster("1610612747", "2023-24")  # Lakers
games = client.get_team_game_log("1610612747", "2023-24")
```

### Get Game Data
```python
# Python
box_score = client.get_box_score_traditional("0022300001")
play_by_play = client.get_play_by_play_v2("0022300001")
```

## Contributing

To add more examples:
1. Create a new file for your language
2. Follow the existing client pattern
3. Include error handling
4. Add usage instructions to this README

## Support

- GitHub Issues: [Report bugs](https://github.com/n-ae/nba-api-go/issues)
- Documentation: [Full docs](https://github.com/n-ae/nba-api-go)
- Migration Guide: [Python to Go](../../docs/MIGRATION_GUIDE.md)
