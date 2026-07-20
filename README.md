# nba-api-go

[![CI](https://github.com/n-ae/nba-api-go/actions/workflows/ci.yml/badge.svg)](https://github.com/n-ae/nba-api-go/actions/workflows/ci.yml)

A type-safe Go library and HTTP API server for accessing NBA statistics from stats.nba.com. Features **100% endpoint coverage** with complete feature parity to the Python [nba_api](https://github.com/swar/nba_api) library.

## 🏆 100% Coverage Achievement

**World's first complete NBA API implementation in Go!**

- ✅ **141 endpoints** (all NBA Stats API endpoints plus international schedule)
- ✅ **All categories complete**: Box Scores, Player, Team, League, Game, Advanced, International
- ✅ **Dual access**: Native Go SDK + HTTP REST API for any language
- ✅ **Type-safe**: generated, endpoint-specific request/response types; see [CHANGELOG.md](./CHANGELOG.md) for the project's actual defect and fix history
- ✅ **Complete feature parity** with Python nba_api + additional features

See [Migration Guide](./docs/MIGRATION_GUIDE.md) to migrate from Python nba_api to Go.

## ⚠️ Type defects fixed in v2.0.0 - if you're on v1.3.0 or earlier, read this

**As of v2.0.0, all three documented type-defect families below are fully fixed** (verified by scanning every committed field against all three patterns - zero remaining instances). If you're upgrading from v1.3.0 or earlier, this is a **source-breaking change** to the affected struct fields; see `CHANGELOG.md`'s `[2.0.0]` section for the full migration guide before upgrading. **If you're still on a tagged release at v1.3.0 or earlier, everything below still applies to you.**

The generator used to infer a Go type for each field from its NBA.com column name (e.g. names containing common stat abbreviations were typed `float64`), and that heuristic got some fields wrong - silently, not as an error:

- **Display-name fields typed `float64`.** `DISPLAY_FIRST_LAST`, `DISPLAY_LAST_COMMA_FIRST`, `DISPLAY_FI_LAST`, `PLAYER_NAME_LAST_FIRST` contain a name string (e.g. `"Nikola Jokić"`); typed `float64`, parsing a name string produced `0`, not an error. Affected `CommonPlayerInfoV2`, `CommonAllPlayers`, `CommonAllPlayersV2`, and `PlayerDashPtShots`.
- **Textual range-bucket fields typed `int`.** `SHOT_CLOCK_RANGE`, `DRIBBLE_RANGE`, `CLOSE_DEF_DIST_RANGE`, `TOUCH_TIME_RANGE`, `SHOT_DIST_RANGE`, `SHOT_ZONE_RANGE` hold text buckets (e.g. `"24-22"`, `"Very Tight"`); typed `int`, parsing text produced `0`. Affected `PlayerDashPtShots`, `TeamDashPtShots`, `PlayerVsPlayer`, `ShotChartLineupDetail`, and `ShotChartDetail`.
- **Decimal percentage/rating fields typed `string`.** `FG_PCT_RA`, `FG_PCT_IN_PAINT`, `FG_PCT_ABOVE_BREAK_3`, `FG_PCT_BACKCOURT`, `FG_PCT_LEFT_CORNER_3`, `FG_PCT_RIGHT_CORNER_3` were declared `string` and formatted with `%.0f`, so a value like `0.357` was stored as `"0"` - decimal precision lost, not merely stringified. Affected `LeagueDashPlayerShotLocations` and `LeagueDashTeamShotLocations`.

This isn't a claim that generated code is now defect-free everywhere - it's specifically these three, previously-documented, previously-confirmed families. `tools/generator/fieldtypes.json`/`fieldtype_overrides.json` cover 854 fields with varying levels of verification confidence (see `CHANGELOG.md`'s `[2.0.0]` notes). `videoevents.go`'s unrelated field-*naming* issue (not a type defect - see `[2.0.0]`'s notes on it) is fixed on `main` post-v2.0.0 via `tools/generator/fieldname_overrides.json`; see `CHANGELOG.md`'s `[Unreleased]` section. See the [maintainability assessment](./docs/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-20_8549390.md) for the analysis that started this fix.

## Features

### Core Features
- **🏆 141 Stats API endpoints** - Complete feature parity with Python nba_api + international schedule
- **Go SDK** - Type-safe library for direct Go integration
- **HTTP API Server** - Complete REST API exposing all endpoints for any language
- **Docker/Podman Ready** - Multi-stage containerization (< 20MB images)
- **Live API** - Real-time game data and scoreboards
- **Static Data** - Pre-loaded player and team datasets with search (5,135 players, 30 teams)
- **Type Safety** - Automatic type inference from field names to Go types
- **Context Support** - Full support for cancellation and timeouts
- **Zero Frameworks** - HTTP API uses stdlib only (net/http, encoding/json)
- **Code Generation** - Advanced tooling for type-safe endpoint generation

### Production Features (HTTP Server)
- **⚡ Rate Limiting** - Per-IP rate limiting (100 req/s, burst 200) to prevent abuse
- **📊 Metrics & Monitoring** - Built-in `/metrics` endpoint with request stats, response times, error rates
- **🏥 Health Checks** - `/health` endpoint with NBA API connectivity status and build info
- **🔒 CORS Support** - Configurable cross-origin resource sharing
- **📝 Request Logging** - Structured logging with response times
- **💾 Low Memory** - < 100MB typical memory usage
- **🔄 Graceful Shutdown** - Safe shutdown with connection draining (10s timeout)

## Installation

**Requires Go 1.26.5 or later** (the `go` directive in `go.mod`; older toolchains cannot build this module).

### Go SDK
```bash
go get github.com/n-ae/nba-api-go/v2
```

### HTTP API Server
```bash
# Docker/Podman
podman pull nba-api-go:latest  # or build locally

# From source
go build -o nba-api-server ./cmd/nba-api-server
```

## Usage Patterns

This project provides **two ways** to access NBA data:

### Pattern 1: Go SDK (For Go Applications)

Best for: Type-safety, performance, direct Go integration

```go
import "github.com/n-ae/nba-api-go/v2/pkg/stats"
```

### Pattern 2: HTTP API (For Any Language)

Best for: Python, JavaScript, Ruby, or any language that can make HTTP requests

```bash
# Start server
docker run -p 8080:8080 nba-api-go

# Use from any language
curl "http://localhost:8080/api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24"
```

See [API Usage Documentation](./docs/API_USAGE.md) for complete HTTP API guide with Python/JavaScript examples.

---

## Quick Start - Go SDK

### Get Player Career Statistics

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/n-ae/nba-api-go/v2/pkg/stats"
    "github.com/n-ae/nba-api-go/v2/pkg/stats/endpoints"
    "github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"
)

func main() {
    client := stats.NewDefaultClient()

    req := endpoints.PlayerCareerStatsRequest{
        PlayerID: "203999", // Nikola Jokić
        PerMode:  parameters.PerModePerGame,
        LeagueID: parameters.LeagueIDNBA,
    }

    resp, err := endpoints.PlayerCareerStats(context.Background(), client, req)
    if err != nil {
        log.Fatal(err)
    }

    for _, season := range resp.Data.SeasonTotalsRegularSeason {
        fmt.Printf("Season %s: %.1f PPG, %.1f RPG, %.1f APG\n",
            season.SeasonID, season.PTS, season.REB, season.AST)
    }
}
```

### Get Player Game Log

```go
req := endpoints.PlayerGameLogRequest{
    PlayerID:   "203999",
    Season:     parameters.NewSeason(2023),
    SeasonType: parameters.SeasonTypeRegular,
}

resp, err := endpoints.PlayerGameLog(context.Background(), client, req)
if err != nil {
    log.Fatal(err)
}

for _, game := range resp.Data.PlayerGameLog {
    fmt.Printf("%s | %s | %d pts, %d reb, %d ast\n",
        game.GameDate, game.Matchup, game.PTS, game.REB, game.AST)
}
```

### Get League Leaders

```go
req := endpoints.LeagueLeadersRequest{
    Season:       parameters.NewSeason(2023),
    SeasonType:   parameters.SeasonTypeRegular,
    StatCategory: parameters.StatCategoryPoints,
    PerMode:      parameters.PerModePerGame,
}

resp, err := endpoints.LeagueLeaders(context.Background(), client, req)
if err != nil {
    log.Fatal(err)
}

for _, leader := range resp.Data.LeagueLeaders {
    fmt.Printf("%d. %s (%s) - %.1f PPG\n",
        leader.Rank, leader.Player, leader.Team, leader.PTS)
}
```

### Get Today's Scoreboard

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/n-ae/nba-api-go/v2/pkg/live"
    "github.com/n-ae/nba-api-go/v2/pkg/live/endpoints"
)

func main() {
    client := live.NewDefaultClient()

    resp, err := endpoints.Scoreboard(context.Background(), client)
    if err != nil {
        log.Fatal(err)
    }

    for _, game := range resp.Data.Scoreboard.Games {
        fmt.Printf("%s @ %s - %s\n",
            game.AwayTeam.TeamTricode,
            game.HomeTeam.TeamTricode,
            game.GameStatusText)
    }
}
```

### Search Players and Teams

```go
package main

import (
    "fmt"
    "log"

    "github.com/n-ae/nba-api-go/v2/pkg/stats/static"
)

func main() {
    players, err := static.SearchPlayers("lebron")
    if err != nil {
        log.Fatal(err)
    }

    for _, player := range players {
        fmt.Printf("%s (ID: %d)\n", player.FullName, player.ID)
    }

    player, err := static.FindPlayerByID(203999)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found: %s\n", player.FullName)

    teams, err := static.SearchTeams("lakers")
    if err != nil {
        log.Fatal(err)
    }

    for _, team := range teams {
        fmt.Printf("%s (%s)\n", team.FullName, team.Abbreviation)
    }
}
```

### Get Player Info (with Date of Birth)

```go
req := endpoints.CommonPlayerInfoRequest{
    PlayerID: "203999", // Nikola Jokić
}

resp, err := endpoints.CommonPlayerInfo(context.Background(), client, req)
if err != nil {
    log.Fatal(err)
}

info := resp.Data.CommonPlayerInfo[0]
dob, err := info.DateOfBirth()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%s was born on %s\n", info.DisplayFirstLast, dob.Format("January 2, 2006"))
```

`DateOfBirth()` parses the endpoint's raw birthdate string into a `time.Time`. It's also
available on `CommonPlayerInfoV2CommonPlayerInfo`, `CommonTeamRosterCommonTeamRoster`, and
`CommonTeamRosterV2CommonTeamRoster` — each parses that endpoint's specific date format.

## Architecture

The library is organized into the following packages:

- **pkg/client** - Core HTTP client with middleware support
- **pkg/client/middleware** - Built-in middleware (retry, rate limiting, headers, logging)
- **pkg/stats** - NBA Stats API client and endpoints
- **pkg/live** - NBA Live Data API client and endpoints
- **pkg/models** - Common data structures and error types
- **pkg/stats/static** - Static player and team data with search
- **pkg/stats/parameters** - Type-safe parameter definitions

### Middleware

The client supports composable middleware for cross-cutting concerns. The
seam itself — `client.Middleware`, `client.RoundTripper`,
`client.RoundTripperFunc`, and `client.Chain` — lives in the importable
`pkg/client` package, so you can write your own middleware without
depending on anything internal to this module. The built-in constructors
(`WithRetry`, `RetryConfig`/`DefaultRetryConfig`, `WithPerHostRateLimit`,
`WithHeaders`, `WithUserAgent`, `WithReferer`, `WithAccept`, and the
`WithLogging` family) live in the importable `pkg/client/middleware`
package too, so you can reconfigure them - not just add alongside them:

```go
import (
    "context"
    "net/http"
    "time"

    "github.com/n-ae/nba-api-go/v2/pkg/client"
    "github.com/n-ae/nba-api-go/v2/pkg/client/middleware"
    "github.com/n-ae/nba-api-go/v2/pkg/stats"
)

// A custom middleware - written entirely with exported types.
func withRequestID(id string) client.Middleware {
    return func(next client.RoundTripper) client.RoundTripper {
        return client.RoundTripperFunc(func(ctx context.Context, req *http.Request) (*http.Response, error) {
            req.Header.Set("X-Request-ID", id)
            return next.RoundTrip(ctx, req)
        })
    }
}

// AdditionalMiddlewares layers your own middleware on top of the default
// chain (retry, the headers NBA.com's API expects, and a per-host rate
// limit) without replacing it - no need to call DefaultMiddlewares()
// yourself.
statsClient := stats.NewClient(stats.Config{
    AdditionalMiddlewares: []client.Middleware{withRequestID("abc123")},
})

// To reconfigure a built-in instead of just adding to it - e.g. a
// stricter retry budget - set Middlewares explicitly (this REPLACES the
// default chain, so re-add anything else you still want from it):
retryConfig := middleware.DefaultRetryConfig()
retryConfig.MaxRetries = 1
statsClient = stats.NewClient(stats.Config{
    Middlewares: []client.Middleware{
        middleware.WithRetry(retryConfig),
        middleware.WithPerHostRateLimit(3, 5),
    },
})
```

`pkg/live` has the equivalent `live.DefaultMiddlewares()` and
`AdditionalMiddlewares`.

## Static Data

The library includes embedded static data for all NBA players and teams:

- **5,135 players** with search by name, ID, or regex
- **30 teams** with search by name, abbreviation, or city
- **Accent-insensitive search** for international player names
- **Active/inactive filtering** for players

## Parameters

All NBA API parameters are strongly typed with validation:

```go
import "github.com/n-ae/nba-api-go/v2/pkg/stats/parameters"

// Season types
parameters.SeasonTypeRegular
parameters.SeasonTypePlayoffs
parameters.SeasonTypeAllStar

// Per-mode calculations
parameters.PerModeTotals
parameters.PerModePerGame
parameters.PerModePer36

// League IDs
parameters.LeagueIDNBA
parameters.LeagueIDABA

// Create seasons
season := parameters.NewSeason(2023) // "2023-24"
```

## Examples

See the [examples](./examples) directory for complete working examples:

- [player_stats](./examples/player_stats) - Fetch player career statistics
- [game_log](./examples/game_log) - Get player game-by-game stats
- [league_leaders](./examples/league_leaders) - View statistical leaders
- [scoreboard](./examples/scoreboard) - Get today's games and scores
- [player_search](./examples/player_search) - Search players and teams

Run examples:

```bash
go run examples/player_stats/main.go
go run examples/game_log/main.go
go run examples/league_leaders/main.go
go run examples/scoreboard/main.go
go run examples/player_search/main.go
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Test example programs compile
make test-examples

# Run benchmarks
go test -bench=. -benchmem ./...

# Run integration tests (requires INTEGRATION_TESTS=1)
INTEGRATION_TESTS=1 go test -tags=integration ./...
```

See [BENCHMARKS.md](./docs/BENCHMARKS.md) for detailed performance analysis.

## Monitoring & Deployment

### Health & Metrics

The HTTP server includes built-in monitoring endpoints:

**Health Check** (`/health`):
```bash
curl http://localhost:8080/health
```

Returns:
- Server status
- NBA API connectivity status (operational/degraded)
- Build information (Go version, build time, git commit)
- Endpoint counts
- Timestamp

**Metrics** (`/metrics`):
```bash
curl http://localhost:8080/metrics
```

Returns:
- Uptime
- Total requests & errors
- Requests by status code
- Requests by path
- Response time statistics (avg, min, max)

### Deployment

See [DEPLOYMENT.md](./DEPLOYMENT.md) for complete deployment guide including:
- Systemd service setup
- Docker/Podman deployment
- Cloud platform deployment (Fly.io, Railway, etc.)
- Reverse proxy configuration (Nginx, Caddy)
- SSL/TLS setup
- Monitoring with Prometheus/UptimeRobot
- Cost estimates and recommendations

**Quick Deploy**:
```bash
# Build with version info
go build \
  -ldflags="-X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) -X main.gitCommit=$(git rev-parse --short HEAD)" \
  -o nba-api-server \
  ./cmd/nba-api-server

# Run
./nba-api-server
```

**Recommended Setup** (Solo Maintainer):
- Hetzner VPS ($5/month)
- Systemd service
- Caddy reverse proxy (auto SSL)
- UptimeRobot (free monitoring)
- **Total**: $5/month, ~1 hour maintenance/month

## Roadmap

Based on the [ADR](./docs/adr/001-go-replication-strategy.md), the library is being developed in phases:

- [x] Phase 1: Foundation (HTTP client, middleware, models)
- [x] Phase 2: Core Stats API (initial endpoints)
- [x] Phase 3: Live API (Scoreboard endpoint)
- [x] Phase 4: Additional Endpoints - ✅ **COMPLETE (141 endpoints including international schedule)**
- [x] Phase 5: Performance Optimization (benchmarks complete)
- [x] Code generation tooling (completed)
- [x] HTTP API Server - ✅ **COMPLETE (all endpoints exposed)**
- [x] Migration guide from Python nba_api - ✅ **COMPLETE (887 lines)**
- [x] HTTP API client examples - ✅ **COMPLETE (Python, JavaScript, Bash)**
- [x] Integration test suite - ✅ **COMPLETE (60+ tests)**
- [ ] CLI tool (optional - low priority)

## Contributing

Contributions are welcome! Please see the [ADR](./docs/adr/001-go-replication-strategy.md) for architectural decisions and development guidelines.

## License

MIT License - see [LICENSE](./LICENSE) file for details.

## Acknowledgments

This project is inspired by and aims for feature parity with the Python [nba_api](https://github.com/swar/nba_api) library by Swar Patel.

## Disclaimer

This library is not affiliated with or endorsed by the NBA. All data is publicly available from NBA.com. Please review the [NBA.com Terms of Use](https://www.nba.com/termsofuse) before using this library.
