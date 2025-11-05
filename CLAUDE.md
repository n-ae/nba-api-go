# CLAUDE.md

This file provides guidance to Claude Code when working with the nba-api-go repository.

## Repository Overview

**nba-api-go** is a production-ready Go SDK and HTTP API server providing type-safe access to all 139 NBA Stats API endpoints. The project emphasizes maintainability, minimal dependencies, and solo engineer viability.

**Current Status**: Ready for v1.0.0 release
**Grade**: A (93/100) - Production-ready with excellent maintainability
**Maintenance Burden**: ~1.6 hours/week

## Project Architecture

### Core Components

- **SDK Library** (`pkg/stats/`): Type-safe Go SDK with 139 endpoints
- **HTTP API Server** (`cmd/nba-api-server/`): REST API exposing all endpoints
- **Code Generator** (`cmd/generator/`): Generates endpoints from NBA.com API analysis
- **Static Data** (`pkg/stats/static/`): 5,135 players, 30 teams (no external DB needed)

### Key Design Principles

1. **Boring tech**: stdlib-only HTTP server, 2 dependencies total
2. **Code generation**: 139 endpoints with 43x productivity gain vs manual
3. **Type safety**: No `interface{}` in generated code, compile-time parameter validation
4. **Solo engineer optimized**: Clear docs, test safety net, minimal operational burden

## Essential Commands

### Development Workflow

```bash
# Quick health check (5 minutes)
go test ./...                    # All unit tests
make test-examples              # Verify all 14 examples compile
make lint                       # golangci-lint
go run ./cmd/nba-api-server     # Start HTTP server
curl http://localhost:8080/health

# Integration tests (requires network)
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# Contract tests (offline, requires fixtures)
go test ./tests/contract/... -v

# Record new contract test fixtures
UPDATE_FIXTURES=1 INTEGRATION_TESTS=1 go test ./tests/contract/... -v
```

### Building

```bash
# Development build
go build -o bin/nba-api-server ./cmd/nba-api-server

# Production build (optimized)
make build

# Generator tool
go build -o bin/generator ./cmd/generator
```

### Common Tasks

```bash
# Add new endpoint (if NBA.com adds one)
./bin/generator analyze https://stats.nba.com/stats/NewEndpoint
./bin/generator generate NewEndpoint
go test ./pkg/stats/endpoints/... -run TestNewEndpoint

# Update dependencies (quarterly)
go get -u golang.org/x/text golang.org/x/time
go mod tidy
go test ./...

# Format code
gofmt -w .
make lint
```

## Testing Strategy

### Test Layers

1. **Unit Tests** (`*_test.go` files throughout)
   - Fast, no network calls
   - Test parameter validation, type safety, error handling
   - Run on every commit

2. **Integration Tests** (`tests/integration/`)
   - Smoke tests for critical endpoints
   - Require live NBA.com API access
   - Skip by default (set `INTEGRATION_TESTS=1` to run)
   - Run before releases or when troubleshooting API issues

3. **Contract Tests** (`tests/contract/`)
   - Record/replay system for API responses
   - Detect NBA.com API drift (schema changes)
   - Offline testing with fixtures
   - Run regularly to catch upstream changes early

4. **Example Tests** (`examples/`)
   - Verify all example programs compile
   - Run via `make test-examples`
   - Documentation doubles as integration tests

### When Tests Fail

**Unit tests fail**: Fix immediately, likely a code bug

**Integration tests fail**: Check if NBA.com API changed or network issues
- Review error messages for API response differences
- Check NBA.com website for announcements
- May need to update SDK structs or parameters

**Contract tests fail**: NBA.com API schema changed
- Run `git diff` on fixture to see what changed
- Update SDK structs in `pkg/stats/endpoints/`
- Update HTTP handlers if needed
- Re-record fixture with `UPDATE_FIXTURES=1`
- Document breaking change in CHANGELOG.md

## Code Generation System

### How It Works

The generator (`cmd/generator/`) analyzes live NBA.com API responses and generates type-safe Go code.

**Generated files** (DO NOT EDIT MANUALLY):
```
pkg/stats/endpoints/
├── playercareerstats.go       # Generated endpoint
├── playergamelog.go           # Generated endpoint
└── ... (137 more)
```

**Template files** (edit these to change generation):
```
cmd/generator/templates/
├── endpoint.go.tmpl           # Endpoint code template
└── test.go.tmpl              # Test code template
```

### Regenerating Endpoints

```bash
# Regenerate single endpoint
./bin/generator generate PlayerCareerStats

# Regenerate all (rarely needed, takes time)
./bin/generator generate-all

# After regeneration
go test ./pkg/stats/endpoints/...
make test-examples
```

### Adding New Endpoints

When NBA.com adds a new endpoint:

1. Analyze the endpoint: `./bin/generator analyze https://stats.nba.com/stats/NewEndpoint`
2. Review analysis output for parameters and response structure
3. Generate code: `./bin/generator generate NewEndpoint`
4. Add integration test in `tests/integration/`
5. Add example in `examples/`
6. Update CHANGELOG.md

## HTTP API Server

### Running Locally

```bash
# Development mode
go run ./cmd/nba-api-server

# Production mode
./bin/nba-api-server -port 8080 -host 0.0.0.0

# With custom rate limiting
./bin/nba-api-server -rate-limit 100 -rate-burst 10
```

### Key Endpoints

- `GET /health` - Health check (always returns 200 OK if server running)
- `GET /metrics` - Prometheus-compatible metrics
- `GET /api/v1/playercareerstats` - Example stats endpoint
- `GET /api/v1/scoreboard` - Live game data

### Deployment Options

See `docs/DEPLOYMENT.md` for:
- systemd service setup
- Container deployment (Containerfile included)
- Kubernetes manifests
- Monitoring setup

## Important Development Notes

### Dependencies

**Keep minimal** (currently 2 dependencies):
- `golang.org/x/text` - Unicode normalization for player search
- `golang.org/x/time` - Rate limiting

**Before adding new dependency**:
1. Check if stdlib can do it
2. Assess maintenance burden
3. Document in ADR
4. Update dependency count in README

### API Stability

**Current: v1.0.0** - Stable with strict semver guarantees

**Breaking changes** require:
- Major version bump
- Migration guide in CHANGELOG.md
- Deprecation period if possible
- Update all examples

**Stability Promise:**
- All public APIs in `pkg/` are stable
- Minor versions (1.x.0) add features without breaking changes
- Patch versions (1.0.x) fix bugs without breaking changes
- Features deprecated for at least one minor version before removal

### NBA.com API Challenges

The upstream NBA.com API has quirks:

1. **No official documentation** - reverse engineered
2. **Changes without notice** - monitor with contract tests
3. **Rate limiting** - SDK includes automatic retry/backoff
4. **Seasonal data gaps** - some endpoints 404 in offseason
5. **Inconsistent field names** - handled by generator type inference

### Performance Considerations

**SDK Performance**:
- Target: <100ms per request (network excluded)
- JSON parsing is the bottleneck
- Consider caching for static data (players, teams)

**Server Performance**:
- Target: Handle 1000 req/sec on single core
- Rate limiting prevents NBA.com throttling
- Stateless design allows horizontal scaling

## Documentation Files

### For Users

- `README.md` - Quick start, installation, examples
- `docs/API_USAGE.md` - Detailed SDK usage guide
- `docs/PYTHON_MIGRATION.md` - For nba_api (Python) users
- `docs/DEPLOYMENT.md` - Production deployment guide
- `CHANGELOG.md` - Version history and upgrade guides
- `examples/` - 14 working example programs

### For Maintainers

- `docs/MAINTENANCE.md` - **START HERE** - Operational runbook
- `docs/MAINTAINABILITY_ASSESSMENT.md` - Solo engineer viability analysis
- `docs/IMPROVEMENTS_COMPLETED.md` - What we've fixed
- `docs/adr/` - Architecture decision records
- `CONTRIBUTING.md` - How to contribute

### Key Documents to Reference

**Before making changes**: Read `docs/MAINTENANCE.md` for procedures

**When adding features**: Check ADRs for architectural guidance

**When troubleshooting**: See `docs/MAINTENANCE.md` Emergency Procedures section

## Maintenance Calendar

### Weekly (10 minutes)
- Run quick health check: `go test ./... && make test-examples`
- Check GitHub issues/PRs
- Monitor any production deployments

### Monthly (30 minutes)
- Review dependencies for security advisories
- Check NBA.com for API announcements
- Update static player data if needed

### Quarterly (2-3 hours)
- Run full integration test suite
- Refresh contract test fixtures
- Dependency updates: `go get -u && go mod tidy`
- Review and update documentation
- Consider performance profiling

### Annually (4-6 hours)
- Full maintainability assessment review
- Major version planning if needed
- Dependency major version updates
- Archive outdated documentation

## Troubleshooting

### "Tests are failing"
1. Check if it's a single test or category (unit/integration/contract)
2. Run with `-v` flag for details: `go test -v ./path/to/test`
3. If integration tests fail, check network and NBA.com API status
4. If contract tests fail, likely NBA.com schema changed - see "When Tests Fail" above

### "Examples won't compile"
1. Verify dependencies: `go mod download`
2. Check for breaking changes in recent commits
3. Run `make test-examples` for detailed output
4. Common issue: Import path typos or missing `go mod tidy`

### "HTTP server returns errors"
1. Check if error is from SDK or NBA.com API
2. Review logs for rate limiting (429 status)
3. Verify endpoint parameters are valid
4. Test with integration tests to isolate issue

### "Generator fails"
1. Verify NBA.com endpoint URL is correct
2. Check if endpoint requires authentication (some do)
3. Review NBA.com website for endpoint deprecation
4. Try with `-debug` flag for detailed output

## Emergency Procedures

### Production Down
1. Check health endpoint: `curl http://server/health`
2. Review logs for panic/crash
3. Check NBA.com API status
4. Rollback if recent deployment
5. See `docs/MAINTENANCE.md` Emergency section for full runbook

### Multiple API Endpoints Failing
Likely NBA.com made breaking changes:
1. Run contract tests to identify affected endpoints
2. Create GitHub issue documenting failures
3. Review NBA.com website/forums for announcements
4. Plan SDK update sprint (may need to regenerate multiple endpoints)
5. Communicate timeline to users

## Architecture Decision Records (ADRs)

### Current ADRs

- **ADR 001**: Go Replication Strategy (why Go, not Python port)
- **ADR 002**: HTTP API Server Architecture (stdlib-only design)

### When to Create New ADR

Create ADR when making decisions about:
- Technology choices (new dependency, framework)
- API design (breaking changes, new patterns)
- Architectural patterns (caching layer, database)
- Development processes (testing strategy, release process)

Template: See `docs/adr/000-template.md`

## Security Considerations

### API Keys/Secrets
- NBA.com API currently doesn't require keys (public stats)
- If authentication added: Use environment variables, never commit secrets
- Document in `.env.example`, add `.env` to `.gitignore`

### Input Validation
- All user input validated via type-safe parameters
- HTTP server sanitizes query params
- No SQL injection risk (no database)
- XSS risk minimal (JSON API, no HTML rendering)

### Rate Limiting
- SDK includes automatic rate limiting
- Server enforces per-host limits
- Prevents abuse of NBA.com API

## Contributing Guidelines

See `CONTRIBUTING.md` for full guidelines. Quick checklist:

**Before submitting PR**:
- [ ] All tests pass: `go test ./...`
- [ ] Examples compile: `make test-examples`
- [ ] Code formatted: `gofmt -w .`
- [ ] Linter passes: `make lint`
- [ ] Documentation updated if needed
- [ ] CHANGELOG.md updated for user-facing changes

**PR Description Should Include**:
- What problem does this solve?
- How was it tested?
- Any breaking changes?
- Related issues/ADRs

## Quick Reference

### File Structure
```
nba-api-go/
├── cmd/
│   ├── nba-api-server/         # HTTP API server (main)
│   └── generator/              # Code generator tool
├── pkg/
│   ├── stats/
│   │   ├── endpoints/          # Generated endpoint code
│   │   ├── parameters/         # Type-safe parameters
│   │   └── static/            # Player/team data
│   └── client/                # HTTP client middleware
├── tests/
│   ├── integration/           # Live API tests
│   └── contract/              # Schema validation tests
├── examples/                  # 14 working examples
└── docs/                     # Documentation
```

### Import Paths
```go
import (
    "github.com/n-ae/nba-api-go/pkg/client"
    "github.com/n-ae/nba-api-go/pkg/stats/endpoints"
    "github.com/n-ae/nba-api-go/pkg/stats/parameters"
    "github.com/n-ae/nba-api-go/pkg/stats/static"
)
```

### Common Parameter Patterns
```go
// Using parameters package for type safety
params := &endpoints.PlayerCareerStatsParams{
    PlayerID: "203999",
    PerMode:  parameters.PerModePtr(parameters.PerModeTotals),
}

// Optional parameters use pointer types
params.Season = parameters.StringPtr("2023-24")
```

### Error Handling Pattern
```go
response, err := endpoints.PlayerCareerStats(ctx, client, params)
if err != nil {
    // Check if API error vs network error
    return fmt.Errorf("failed to fetch stats: %w", err)
}
```

## Related Projects

- **nba_api** (Python) - Original inspiration for this project
- **balldontlie** - Alternative NBA API (different data source)
- **nba-go** - Another Go implementation (different scope)

## Support and Community

- **Issues**: GitHub Issues for bugs/features
- **Discussions**: GitHub Discussions for questions
- **Documentation**: All docs in `docs/` directory
- **Examples**: See `examples/` for working code

## Version Information

**Current Version**: 1.0.0 (Stable Release)
**Go Version**: 1.21+
**Stability**: Production-ready with semver guarantees

See `CHANGELOG.md` for full version history.

---

**Last Updated**: 2025-11-05
**Maintainability Grade**: A (93/100)
**Next Review**: 2026-02-05 (quarterly)
