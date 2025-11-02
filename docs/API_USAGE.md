# NBA API Server - Usage Guide

## Overview

The NBA API Server provides HTTP/REST access to NBA statistics data. It's designed for non-Go applications (Python, JavaScript, etc.) to easily consume NBA data.

## Two Usage Patterns

### Pattern 1: Go SDK (Recommended for Go apps)

```go
import "github.com/username/nba-api-go/pkg/stats"

client := stats.NewDefaultClient()
resp, err := endpoints.PlayerGameLog(ctx, client, req)
```

**Pros:** Type-safe, fastest, no network overhead
**Cons:** Go only

### Pattern 2: HTTP API (For non-Go apps)

```bash
curl http://localhost:8080/api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24
```

**Pros:** Language-agnostic (Python, JavaScript, Ruby, etc.)
**Cons:** Network latency, HTTP overhead

---

## Quick Start

### Option A: Docker/Podman (Recommended)

```bash
# Using podman
podman build -t nba-api-go -f Containerfile .
podman run -p 8080:8080 nba-api-go

# Using docker
docker build -t nba-api-go -f Containerfile .
docker run -p 8080:8080 nba-api-go

# Using docker-compose/podman-compose
docker-compose up
```

### Option B: Direct Binary

```bash
# Build
go build -o nba-api-server ./cmd/nba-api-server

# Run
./nba-api-server

# Or with custom port
PORT=3000 ./nba-api-server
```

---

## API Endpoints

### Health Check

```bash
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "version": "0.1.0",
  "endpoints_count": 79
}
```

---

## Stats Endpoints

Base URL: `http://localhost:8080/api/v1/stats/`

### 1. Player Game Log

Get game-by-game stats for a player.

```bash
GET /api/v1/stats/playergamelog?PlayerID={id}&Season={season}
```

**Parameters:**
- `PlayerID` (required) - Player ID (e.g., "2544" for LeBron James)
- `Season` (optional) - Season year (default: "2023-24")
- `SeasonType` (optional) - "Regular Season" or "Playoffs" (default: "Regular Season")

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24"
```

### 2. Common All Players

Get all players for a season.

```bash
GET /api/v1/stats/commonallplayers?Season={season}
```

**Parameters:**
- `Season` (required) - Season year (e.g., "2023-24")
- `IsOnlyCurrentSeason` (optional) - "0" or "1" (default: "0")

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/commonallplayers?Season=2023-24&IsOnlyCurrentSeason=1"
```

### 3. Scoreboard

Get games for a specific date.

```bash
GET /api/v1/stats/scoreboardv2?GameDate={date}
```

**Parameters:**
- `GameDate` (required) - Date in YYYY-MM-DD format

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/scoreboardv2?GameDate=2024-01-15"
```

### 4. League Standings

Get current league standings.

```bash
GET /api/v1/stats/leaguestandings?Season={season}
```

**Parameters:**
- `Season` (optional) - Season year (default: "2023-24")
- `SeasonType` (optional) - "Regular Season" or "Playoffs" (default: "Regular Season")

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/leaguestandings?Season=2023-24"
```

### 5. Team Roster

Get roster for a specific team.

```bash
GET /api/v1/stats/commonteamroster?TeamID={id}
```

**Parameters:**
- `TeamID` (required) - Team ID (e.g., "1610612747" for Lakers)
- `Season` (optional) - Season year (default: "2023-24")

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/commonteamroster?TeamID=1610612747&Season=2023-24"
```

### 6. Player Career Stats

Get career statistics for a player.

```bash
GET /api/v1/stats/playercareerstats?PlayerID={id}
```

**Parameters:**
- `PlayerID` (required) - Player ID
- `PerMode` (optional) - "PerGame", "Totals", or "Per36" (default: "PerGame")

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/playercareerstats?PlayerID=2544&PerMode=PerGame"
```

### 7. League Leaders

Get statistical leaders.

```bash
GET /api/v1/stats/leagueleaders?Season={season}
```

**Parameters:**
- `Season` (optional) - Season year (default: "2023-24")
- `SeasonType` (optional) - "Regular Season" or "Playoffs"
- `PerMode` (optional) - "PerGame", "Totals", or "Per36"

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/leagueleaders?Season=2023-24&PerMode=PerGame"
```

### 8. Common Player Info

Get detailed player information.

```bash
GET /api/v1/stats/commonplayerinfo?PlayerID={id}
```

**Parameters:**
- `PlayerID` (required) - Player ID

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/commonplayerinfo?PlayerID=2544"
```

### 9. League Dash Team Stats

Get league-wide team statistics dashboard.

```bash
GET /api/v1/stats/leaguedashteamstats?Season={season}
```

**Parameters:**
- `Season` (optional) - Season year (default: "2023-24")
- `SeasonType` (optional) - "Regular Season" or "Playoffs"
- `PerMode` (optional) - "PerGame", "Totals", or "Per36"

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/leaguedashteamstats?Season=2023-24&PerMode=PerGame"
```

### 10. League Dash Player Stats

Get league-wide player statistics dashboard.

```bash
GET /api/v1/stats/leaguedashplayerstats?Season={season}
```

**Parameters:**
- `Season` (optional) - Season year (default: "2023-24")
- `SeasonType` (optional) - "Regular Season" or "Playoffs"
- `PerMode` (optional) - "PerGame", "Totals", or "Per36"

**Example:**
```bash
curl "http://localhost:8080/api/v1/stats/leaguedashplayerstats?Season=2023-24&PerMode=PerGame"
```

---

## Response Format

### Success Response

```json
{
  "success": true,
  "data": {
    "PlayerGameLog": [
      {
        "GameID": "0022300123",
        "GameDate": "2024-01-15",
        "PTS": 30,
        "REB": 8,
        "AST": 7
      }
    ]
  }
}
```

### Error Response

```json
{
  "success": false,
  "error": {
    "code": "missing_parameter",
    "message": "PlayerID is required"
  }
}
```

---

## Client Examples

### Python

```python
import requests

# Get player game log
response = requests.get(
    "http://localhost:8080/api/v1/stats/playergamelog",
    params={
        "PlayerID": "2544",
        "Season": "2023-24"
    }
)

data = response.json()
if data["success"]:
    games = data["data"]["PlayerGameLog"]
    for game in games:
        print(f"{game['GameDate']}: {game['PTS']} points")
```

### JavaScript (Node.js)

```javascript
const axios = require('axios');

async function getPlayerGameLog(playerId, season) {
  const response = await axios.get('http://localhost:8080/api/v1/stats/playergamelog', {
    params: {
      PlayerID: playerId,
      Season: season
    }
  });

  if (response.data.success) {
    return response.data.data.PlayerGameLog;
  } else {
    throw new Error(response.data.error.message);
  }
}

getPlayerGameLog('2544', '2023-24')
  .then(games => console.log(games))
  .catch(err => console.error(err));
```

### JavaScript (Browser)

```javascript
fetch('http://localhost:8080/api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24')
  .then(response => response.json())
  .then(data => {
    if (data.success) {
      console.log(data.data.PlayerGameLog);
    } else {
      console.error(data.error.message);
    }
  });
```

### cURL

```bash
# Get scoreboard
curl "http://localhost:8080/api/v1/stats/scoreboardv2?GameDate=2024-01-15" | jq '.'

# Get player stats
curl "http://localhost:8080/api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24" | jq '.data.PlayerGameLog[0]'

# Get standings
curl "http://localhost:8080/api/v1/stats/leaguestandings?Season=2023-24" | jq '.data.Standings[] | {team: .TeamName, wins: .WINS, losses: .LOSSES}'
```

---

## Docker Compose Integration

For applications using Docker Compose:

```yaml
version: '3.8'

services:
  # Your application
  my-app:
    build: .
    environment:
      - NBA_API_URL=http://nba-api:8080
    depends_on:
      - nba-api

  # NBA API sidecar
  nba-api:
    image: nba-api-go:latest
    ports:
      - "8080:8080"
```

Then in your app:
```python
import os
import requests

NBA_API_URL = os.getenv('NBA_API_URL', 'http://localhost:8080')

response = requests.get(f"{NBA_API_URL}/api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24")
```

---

## Environment Variables

- `PORT` - HTTP server port (default: 8080)
- `LOG_LEVEL` - Logging level: "debug", "info", "warn", "error" (default: "info")
- `NBA_API_TIMEOUT` - Timeout for NBA.com API requests (default: "30s")

---

## Common Team IDs

| Team | ID |
|------|-----|
| Lakers | 1610612747 |
| Warriors | 1610612744 |
| Celtics | 1610612738 |
| Heat | 1610612748 |
| Bulls | 1610612741 |

[Full list available in the Python nba_api documentation]

---

## Common Player IDs

| Player | ID |
|--------|-----|
| LeBron James | 2544 |
| Stephen Curry | 201939 |
| Kevin Durant | 201142 |
| Giannis Antetokounmpo | 203507 |

**Tip:** Use `/api/v1/stats/commonallplayers` to get all player IDs.

---

## Rate Limiting

The server inherits rate limiting from the Go SDK:
- **3 requests per 5 seconds** per host (to NBA.com)

**Note:** This is per API server instance, not per client.

---

## Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `missing_parameter` | 400 | Required parameter not provided |
| `invalid_parameter` | 400 | Parameter value is invalid |
| `endpoint_not_found` | 404 | Endpoint not supported |
| `api_error` | 500 | Error calling NBA.com API |
| `method_not_allowed` | 405 | Only GET requests supported |

---

## Best Practices

### 1. Use Sidecar Pattern

Deploy the NBA API as a sidecar container alongside your application for minimum latency:

```yaml
services:
  app:
    build: .
    environment:
      - NBA_API_URL=http://nba-api:8080
  nba-api:
    image: nba-api-go:latest
```

### 2. Cache Results

NBA stats don't change frequently. Cache responses client-side:

```python
import requests
from datetime import datetime, timedelta

class NBAClient:
    def __init__(self):
        self.cache = {}
        self.cache_ttl = timedelta(minutes=5)

    def get_scoreboard(self, date):
        cache_key = f"scoreboard_{date}"
        if cache_key in self.cache:
            cached_at, data = self.cache[cache_key]
            if datetime.now() - cached_at < self.cache_ttl:
                return data

        response = requests.get(f"http://localhost:8080/api/v1/stats/scoreboardv2?GameDate={date}")
        data = response.json()['data']
        self.cache[cache_key] = (datetime.now(), data)
        return data
```

### 3. Handle Errors Gracefully

```javascript
async function fetchWithRetry(url, retries = 3) {
  for (let i = 0; i < retries; i++) {
    try {
      const response = await fetch(url);
      const data = await response.json();
      if (data.success) {
        return data.data;
      }
      throw new Error(data.error.message);
    } catch (error) {
      if (i === retries - 1) throw error;
      await new Promise(resolve => setTimeout(resolve, 1000 * (i + 1)));
    }
  }
}
```

---

## Troubleshooting

### Server won't start

```bash
# Check if port is already in use
lsof -i :8080

# Use different port
PORT=3000 ./nba-api-server
```

### Container build fails

```bash
# Ensure you're in the project root
pwd  # Should show .../nba-api-go

# Use podman instead of docker (if on Linux)
podman build -t nba-api-go -f Containerfile .
```

### API returns 500 errors

Check server logs:
```bash
# Docker
docker logs <container-id>

# Podman
podman logs <container-id>

# Direct binary
# Logs output to stdout
```

---

## Next Steps

- See [ADR-002](./adr/002-api-server-architecture.md) for architecture details
- See [README.md](../README.md) for Go SDK usage
- See [examples/](../examples/) for more code samples

---

## Support

- GitHub Issues: https://github.com/username/nba-api-go/issues
- Based on Python nba_api: https://github.com/swar/nba_api
