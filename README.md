# nba-api-go

A Go client library for accessing NBA.com APIs. This is a Go implementation inspired by the Python [nba_api](https://github.com/swar/nba_api) library.

## Features

- **Stats API** - Access to NBA official stats (stats.nba.com)
- **Live API** - Real-time game data and scoreboards
- **Static Data** - Pre-loaded player and team datasets with search functionality
- **Type Safety** - Strongly typed requests and responses
- **Middleware Support** - Rate limiting, retry logic, logging, and custom headers
- **Context Support** - Full support for cancellation and timeouts
- **No External Dependencies** - Uses only Go standard library and golang.org/x packages

## Installation

```bash
go get github.com/username/nba-api-go
```

## Quick Start

### Get Player Career Statistics

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/username/nba-api-go/pkg/stats"
    "github.com/username/nba-api-go/pkg/stats/endpoints"
    "github.com/username/nba-api-go/pkg/stats/parameters"
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

    "github.com/username/nba-api-go/pkg/live"
    "github.com/username/nba-api-go/pkg/live/endpoints"
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

    "github.com/username/nba-api-go/pkg/stats/static"
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

## Architecture

The library is organized into the following packages:

- **pkg/client** - Core HTTP client with middleware support
- **pkg/stats** - NBA Stats API client and endpoints
- **pkg/live** - NBA Live Data API client and endpoints
- **pkg/models** - Common data structures and error types
- **pkg/stats/static** - Static player and team data with search
- **pkg/stats/parameters** - Type-safe parameter definitions
- **internal/middleware** - HTTP middleware (rate limiting, retry, logging)

### Middleware

The client supports composable middleware for cross-cutting concerns:

```go
import (
    "github.com/username/nba-api-go/internal/middleware"
    "github.com/username/nba-api-go/pkg/client"
)

config := client.Config{
    BaseURL: "https://stats.nba.com/stats",
    Middlewares: []middleware.Middleware{
        middleware.WithUserAgent("MyApp/1.0"),
        middleware.WithReferer("https://www.nba.com/"),
        middleware.WithPerHostRateLimit(3, 5),
        middleware.WithRetry(middleware.DefaultRetryConfig()),
        middleware.WithLogging(nil), // uses default logger
    },
}

client := client.NewClient(config)
```

## Static Data

The library includes embedded static data for all NBA players and teams:

- **5,135 players** with search by name, ID, or regex
- **30 teams** with search by name, abbreviation, or city
- **Accent-insensitive search** for international player names
- **Active/inactive filtering** for players

## Parameters

All NBA API parameters are strongly typed with validation:

```go
import "github.com/username/nba-api-go/pkg/stats/parameters"

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

# Run benchmarks
go test -bench=. -benchmem ./...

# Run integration tests (requires INTEGRATION_TESTS=1)
INTEGRATION_TESTS=1 go test -tags=integration ./...
```

See [BENCHMARKS.md](./docs/BENCHMARKS.md) for detailed performance analysis.

## Roadmap

Based on the [ADR](./docs/adr/001-go-replication-strategy.md), the library is being developed in phases:

- [x] Phase 1: Foundation (HTTP client, middleware, models)
- [x] Phase 2: Core Stats API (5 endpoints implemented)
- [x] Phase 3: Live API (Scoreboard endpoint)
- [x] Phase 4: Additional Endpoints (4/139 stats, benchmarks, integration tests) - IN PROGRESS
- [x] Phase 5: Performance Optimization (benchmarks complete) - IN PROGRESS
- [ ] Phase 6: Code generation tooling
- [ ] Phase 7: CLI tool

## Contributing

Contributions are welcome! Please see the [ADR](./docs/adr/001-go-replication-strategy.md) for architectural decisions and development guidelines.

## License

MIT License - see [LICENSE](./LICENSE) file for details.

## Acknowledgments

This project is inspired by and aims for feature parity with the Python [nba_api](https://github.com/swar/nba_api) library by Swar Patel.

## Disclaimer

This library is not affiliated with or endorsed by the NBA. All data is publicly available from NBA.com. Please review the [NBA.com Terms of Use](https://www.nba.com/termsofuse) before using this library.
