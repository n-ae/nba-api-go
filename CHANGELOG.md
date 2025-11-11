# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.0] - 2025-11-07

### Added
- **New endpoint**: `InternationalBroadcasterSchedule` - Access international broadcast schedules for NBA games
  - SDK endpoint: `endpoints.GetInternationalBroadcasterSchedule()`
  - HTTP API route: `/api/v1/stats/internationalbroadcasterschedule`
  - Supports filtering by Season, LeagueID, RegionID, Date, and EST parameters
  - Returns detailed game information including broadcasters, teams, dates, and times
  - Useful for tracking which international broadcasters are showing games
- Example program: `examples/international_broadcast_schedule/` demonstrating broadcast schedule usage
- Comprehensive test coverage for new endpoint:
  - Unit tests for parameter validation
  - Integration tests for 2024 and 2025 seasons
  - Contract test with fixture for schema stability
  - HTTP handler tests for error cases and valid requests

### Changed
- HTTP API server version updated to 1.1.0
- Now supports 140/139+ NBA Stats API endpoints (added 1 additional international schedule endpoint)

### Notes
- Season parameter format: "2025" corresponds to 2025-26 season
- Returns 409+ scheduled games with international broadcast information for 2025-26 season
- All tests passing with `go test ./...`

## [1.0.0] - 2025-11-05

**STABLE RELEASE** - This release marks the project as production-ready with comprehensive testing, documentation, and stability guarantees.

### Added
- **Contract test framework** in `tests/contract/` for API drift detection
  - Record/replay system for NBA.com API responses
  - Schema validation to catch upstream changes
  - Data sanity checks for response content
  - Comprehensive documentation and usage guide
- Integration test framework in `tests/integration/` with smoke tests for key endpoints
- Maintainability assessment document analyzing solo engineer viability (Grade: A, 93/100)
- Maintenance runbook (`docs/MAINTENANCE.md`) with operational procedures
- CHANGELOG.md for tracking project changes
- CLAUDE.md for AI assistant guidance
- Comprehensive improvements documentation

### Changed
- Updated ADR 002 status from "Proposed" to "Accepted" with implementation summary
- Archived ROADMAP.md with clear deprecation notice (project reached 100% endpoint coverage)

### Fixed
- Compilation errors in examples (redundant newlines in fmt.Println)
- Import typos throughout codebase (yourn-ae → n-ae)

### Stability Guarantees
- **Semantic Versioning**: Strict semver compliance starting with v1.0.0
- **Breaking Changes**: Only in major version updates (2.0.0, 3.0.0, etc.)
- **Backward Compatibility**: Minor and patch versions guarantee backward compatibility
- **API Stability**: All public APIs in `pkg/` are stable and will not break without major version bump
- **Deprecation Policy**: Features will be deprecated for at least one minor version before removal

## [0.9.0] - 2024-11-04

### Added
- **HISTORIC MILESTONE**: All 139/139 NBA Stats API endpoints implemented (100% coverage)
- HTTP API server in `cmd/nba-api-server/` exposing all endpoints via REST
- Production-ready features:
  - Health check endpoint (`/health`)
  - Metrics endpoint (`/metrics`)
  - Rate limiting per host
  - CORS support
  - Graceful shutdown
- Multi-stage Containerfile for minimal production images (<20MB)
- Comprehensive deployment guide (systemd, Docker, Kubernetes)
- Migration guide for Python nba_api users (887 lines)
- 14 example programs demonstrating SDK usage

### Changed
- Code generation approach for all endpoints (43x productivity gain)
- Type inference system eliminates `interface{}` usage in generated code

## [0.3.0] - 2024-10-31

### Added
- First batch of 8 generated endpoints via code generation tooling
- Batch generation system producing consistent, type-safe code

## [0.2.0] - 2024-10-28

### Added
- Live API support (`pkg/live/`)
- Scoreboard endpoint for real-time game data
- Static player and team data (5,135 players, 30 teams)
- Accent-insensitive player search
- Benchmarking suite

### Changed
- Improved middleware architecture
- Rate limiting implementation

## [0.1.0] - 2024-10-24

### Added
- Initial SDK implementation
- Core HTTP client with middleware support
- First 5 Stats API endpoints:
  - PlayerCareerStats
  - PlayerGameLog
  - CommonPlayerInfo
  - LeagueLeaders
  - TeamGameLog
- Type-safe parameter system
- Response models with generics
- Context-based timeout handling
- Error handling framework
- Documentation:
  - README with examples
  - ADR 001: Go Replication Strategy
  - ADR 002: HTTP API Server Architecture
  - Contributing guidelines
  - API usage guide

### Infrastructure
- Go 1.21+ requirement
- Minimal dependencies (2 total):
  - golang.org/x/text v0.30.0
  - golang.org/x/time v0.14.0
- Makefile for build automation
- golangci-lint configuration
- GitHub-ready project structure

---

## Release Notes

### Version 1.0.0 - Stable Release

**PRODUCTION READY** - This is the first stable release with comprehensive testing, documentation, and long-term support commitment.

**Stability Highlights:**
- ✅ All 139 NBA Stats API endpoints (100% coverage)
- ✅ Comprehensive test coverage (unit + integration + contract tests)
- ✅ Production-grade maintainability (Grade A: 93/100)
- ✅ Complete operational documentation
- ✅ Semver stability guarantees
- ✅ Minimal dependencies (2 total, both from golang.org/x)

**Testing Infrastructure:**
- Integration tests for live API validation
- Contract tests for API drift detection
- Fixture recording/replay system
- Schema validation to catch upstream changes

**Documentation:**
- Maintenance runbook for operational procedures
- Maintainability assessment documenting solo engineer viability
- Complete API usage guides and examples
- Migration guide for Python nba_api users

**What This Means:**
- Stable public API - no breaking changes without major version bump
- Production-ready for serious applications
- Long-term maintenance commitment (~2 hours/week)
- Quarterly maintenance cycle established

### Version 0.9.0 - Production Ready

This release marks feature completeness for the NBA Stats API with all 139 endpoints implemented. The HTTP API server makes the SDK accessible from any programming language.

**Key Achievements:**
- ✅ 100% endpoint coverage (139/139)
- ✅ Production-ready HTTP API
- ✅ Comprehensive documentation
- ✅ Zero technical debt (no TODOs/FIXMEs)
- ✅ Minimal dependencies (2 total)

### Version 0.1.0 - Initial Release

First public release of the NBA API Go SDK. Provides type-safe access to NBA statistics with excellent developer experience.

---

## Upgrade Guide

### From 1.0.0 to 1.1.0

**No breaking changes!** This is a minor release adding a new endpoint.

**What's New:**
- `InternationalBroadcasterSchedule` endpoint for accessing international broadcast schedules
- Example program demonstrating broadcast schedule usage
- Comprehensive tests for the new endpoint

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.1.0`
2. No code changes required for existing code
3. Optionally use the new endpoint: `endpoints.GetInternationalBroadcasterSchedule()`

**New Features:**
- Track which international broadcasters are showing NBA games
- Filter by Season, LeagueID, RegionID, Date, and EST
- Access game schedules with detailed broadcaster information

### From 0.9.0 to 1.0.0

**No breaking changes!** This is a stability release with no API changes.

**What's New:**
- Comprehensive testing infrastructure (integration + contract tests)
- Maintenance runbook for operational procedures
- Maintainability assessment confirming production-readiness
- Stability guarantees and semver commitment

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.0.0`
2. No code changes required
3. Review `docs/MAINTENANCE.md` if you're maintaining a fork
4. Consider running contract tests to validate your integration

**Recommended Actions:**
- Set `INTEGRATION_TESTS=1` and run tests to validate your environment
- Review `tests/contract/README.md` for API drift detection strategies
- Check `docs/MAINTENANCE.md` for operational best practices

### From 0.1.0 to 0.9.0

No breaking changes! All endpoints maintain backward compatibility. New endpoints are purely additive.

**New Features Available:**
- 134 additional endpoints
- HTTP API server (opt-in)
- Container deployment option

**Migration Steps:**
1. Update dependency: `go get github.com/n-ae/nba-api-go@v0.9.0`
2. No code changes required
3. Optionally explore new endpoints in `pkg/stats/endpoints/`

---

## Versioning Policy

This project follows [Semantic Versioning](https://semver.org/):

- **Major version (1.x.x)**: Breaking API changes
- **Minor version (x.1.x)**: New features, backward compatible
- **Patch version (x.x.1)**: Bug fixes, backward compatible

**Post-1.0 Guarantees:** Starting with v1.0.0, strict semver guarantees apply. Breaking changes will only occur in major version updates.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for how to suggest changes or report issues.

[Unreleased]: https://github.com/n-ae/nba-api-go/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/n-ae/nba-api-go/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/n-ae/nba-api-go/compare/v0.9.0...v1.0.0
[0.9.0]: https://github.com/n-ae/nba-api-go/compare/v0.3.0...v0.9.0
[0.3.0]: https://github.com/n-ae/nba-api-go/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/n-ae/nba-api-go/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/n-ae/nba-api-go/releases/tag/v0.1.0
