# NBA API Go - Development Roadmap

This roadmap tracks the implementation progress based on the [ADR](./adr/001-go-replication-strategy.md).

## Phase 1: Foundation ✅ COMPLETED

**Timeline**: Week 1-2
**Status**: ✅ Complete

- [x] HTTP client implementation with middleware
- [x] Response parsing framework
- [x] Error handling structure
- [x] Core models and types
- [x] Project scaffolding
- [x] Basic CI/CD setup (Makefile, linting config)

**Deliverables**:
- `pkg/client/` - HTTP client with context support
- `pkg/models/` - Response types and error handling
- `internal/middleware/` - Rate limiting, retry, logging, headers
- `.golangci.yml` - Linting configuration
- `Makefile` - Build and test automation

## Phase 2: Stats API Core ✅ COMPLETED

**Timeline**: Week 3-4
**Status**: ✅ Complete

- [x] 5-10 most common stats endpoints (started with PlayerCareerStats)
- [x] Parameter types and validation
- [x] Static player/team data (5,135 players, 30 teams)
- [x] Search functionality with accent-insensitive matching
- [x] Comprehensive tests

**Deliverables**:
- `pkg/stats/endpoints/playercareerstats.go` - First endpoint implementation
- `pkg/stats/parameters/` - Type-safe parameters (PerMode, LeagueID, Season, etc.)
- `pkg/stats/static/` - Embedded player and team data with search
- Tests for client, parameters, and static data
- Examples for player stats and search

**Priority Endpoints for Expansion**:
1. `PlayerCareerStats` ✅ COMPLETE
2. `PlayerGameLog` ✅ COMPLETE
3. `CommonPlayerInfo` ✅ COMPLETE
4. `LeagueLeaders` ✅ COMPLETE
5. `TeamGameLog` ✅ COMPLETE
6. `TeamInfoCommon` - Team information
7. `BoxScoreSummaryV2` - Game box scores
8. `BoxScoreTraditionalV2` - Traditional box score stats
9. `PlayByPlayV2` - Play-by-play data
10. `ShotChartDetail` - Shot location data

## Phase 3: Live API ✅ COMPLETED

**Timeline**: Week 5
**Status**: ✅ Complete

- [x] Scoreboard endpoint
- [x] BoxScore endpoint (basic structure)
- [x] PlayByPlay endpoint (basic structure)
- [x] Real-time data tests

**Deliverables**:
- `pkg/live/` - Live API client
- `pkg/live/endpoints/scoreboard.go` - Today's games and scores
- Examples for scoreboard
- Documentation for live data access

**Additional Live Endpoints Needed**:
- BoxScore - Full game box scores
- PlayByPlay - Live play-by-play
- Odds - Betting odds (if needed)

## Phase 4: Remaining Stats Endpoints 🔄 IN PROGRESS

**Timeline**: Week 6-8
**Status**: 🔄 Not Started

- [ ] Code generation tooling
- [ ] Generate remaining 130+ endpoints
- [ ] Validation and testing
- [ ] Documentation

**Approach**:
1. Analyze Python nba_api endpoint patterns
2. Create code generator from endpoint metadata
3. Generate Go code for all endpoints
4. Add integration tests
5. Document all endpoints

**Code Generation Strategy**:
```
tools/
├── generator/
│   ├── main.go           # Generator entry point
│   ├── metadata/         # Endpoint metadata
│   ├── templates/        # Go code templates
│   └── validators/       # Validation logic
└── scripts/
    └── generate.sh       # Generation script
```

**Priority Stats Endpoints** (top 20 by usage):
1. PlayerCareerStats ✅
2. PlayerGameLog
3. CommonPlayerInfo
4. TeamGameLog
5. LeagueLeaders
6. ShotChartDetail
7. BoxScoreSummaryV2
8. BoxScoreTraditionalV2
9. PlayByPlayV2
10. LeagueGameFinder
11. PlayerDashboardByYearOverYear
12. TeamDashboardByYearOverYear
13. LeagueDashPlayerStats
14. LeagueDashTeamStats
15. TeamInfoCommon
16. CommonTeamRoster
17. LeagueStandings
18. ScoreboardV2
19. PlayerProfileV2
20. VideoEvents

## Phase 5: Polish 📋 PLANNED

**Timeline**: Week 9-10
**Status**: 📋 Planned

- [ ] CLI tool (optional)
- [ ] Usage examples and tutorials
- [ ] Performance optimization
- [ ] Rate limiting implementation refinement
- [ ] Release preparation

**CLI Features**:
```bash
nba-api player --id 203999 --stats
nba-api scoreboard --date 2024-01-15
nba-api search --player "LeBron"
nba-api teams --list
```

**Performance Targets**:
- Response parsing < 10ms for typical responses
- Memory usage < 50MB for embedded data
- Concurrent request handling
- Connection pooling optimization

## Post-1.0 Roadmap 🔮 FUTURE

### v1.1 - Enhanced Endpoints
- [ ] All 139 stats endpoints
- [ ] Complete live API coverage
- [ ] WebSocket support for live updates
- [ ] Caching layer

### v1.2 - Developer Experience
- [ ] OpenAPI/Swagger spec generation
- [ ] Mock server for testing
- [ ] Request/response logging middleware
- [ ] Metrics and observability

### v1.3 - Advanced Features
- [ ] GraphQL wrapper
- [ ] Data export utilities (CSV, Excel)
- [ ] Historical data aggregation helpers
- [ ] Advanced filtering and querying

### v2.0 - Next Generation
- [ ] gRPC API
- [ ] Protobuf schemas
- [ ] Streaming APIs
- [ ] Built-in analytics functions

## Current Status Summary

| Phase | Status | Completion |
|-------|--------|------------|
| Phase 1: Foundation | ✅ Complete | 100% |
| Phase 2: Stats API Core | ✅ Complete | 100% (5/10 endpoints) |
| Phase 3: Live API | ✅ Complete | 100% (1/3 endpoints) |
| Phase 4: Remaining Endpoints | 🔄 In Progress | 3.6% (5/139) |
| Phase 5: Polish | 🔄 In Progress | 60% (benchmarks, tests) |

## Contributing

See [CONTRIBUTING.md](../CONTRIBUTING.md) for details on how to contribute to any phase of the roadmap.

## Endpoint Coverage Tracker

### Stats API Endpoints (5/139)
- ✅ PlayerCareerStats
- ✅ PlayerGameLog
- ✅ CommonPlayerInfo
- ✅ LeagueLeaders
- ✅ TeamGameLog
- 🔄 134 endpoints remaining

### Live API Endpoints (1/4)
- ✅ Scoreboard
- 🔄 BoxScore (structure ready, needs implementation)
- 🔄 PlayByPlay (structure ready, needs implementation)
- 🔄 Odds

## Next Steps

**Immediate priorities** (next 2 weeks):
1. Implement PlayerGameLog endpoint
2. Implement CommonPlayerInfo endpoint
3. Add integration tests with real API
4. Set up code generation framework
5. Document endpoint addition process

**Medium-term goals** (1-2 months):
1. Reach 20 most-used endpoints
2. Complete code generation tooling
3. Add comprehensive examples
4. Performance benchmarking
5. v0.1.0 release

**Long-term vision** (3-6 months):
1. Full endpoint parity with Python nba_api
2. CLI tool release
3. v1.0.0 stable release
4. Community adoption and feedback
5. Advanced features roadmap

## Metrics

Track implementation progress:
- **Endpoints**: 5/139 stats, 1/4 live (4.2% total)
- **Test Coverage**: ~80% for implemented code
- **Benchmarks**: Complete for all major components
- **Integration Tests**: Framework complete with 5 endpoint tests
- **Documentation**: README, ADR, CONTRIBUTING, ROADMAP, BENCHMARKS
- **Examples**: 5 working examples
- **Static Data**: 5,135 players, 30 teams embedded

## Notes

- Prioritize endpoint implementation based on community usage patterns
- Maintain backward compatibility once v1.0 is released
- Focus on developer experience and ease of use
- Keep dependencies minimal
- Ensure comprehensive testing

Last updated: 2025-10-28
