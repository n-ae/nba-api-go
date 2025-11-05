# Maintenance Runbook

**Audience**: Future maintainers (including future you!)
**Purpose**: Quick reference for common maintenance tasks
**Estimated maintenance time**: ~2 hours/week

---

## Quick Health Check

```bash
# 1. All tests pass
go test ./...

# 2. All examples compile
make test-examples

# 3. Linting passes
make lint

# 4. HTTP server starts
go run ./cmd/nba-api-server

# 5. Health endpoint responds
curl http://localhost:8080/health
```

**If all 5 pass**: Project is healthy ✅

---

## Common Maintenance Tasks

### 1. Update Dependencies (Quarterly)

**Frequency**: Every 3 months
**Time**: 15-30 minutes

```bash
# Check for updates
go list -u -m all

# Update golang.org/x packages
go get -u golang.org/x/text
go get -u golang.org/x/time

# Verify nothing broke
go test ./...
make test-examples

# Commit
git add go.mod go.sum
git commit -m "chore: update dependencies"
```

**Watch for**:
- Breaking changes in golang.org/x packages (rare)
- Test failures after update
- Security vulnerabilities (enable GitHub Dependabot)

---

### 2. Fix NBA.com API Changes

**Trigger**: Users report endpoint errors
**Frequency**: 2-4 times/year
**Time**: 2-4 hours

#### Symptoms
- Endpoint returns 400/404/500 errors
- Response structure changed
- New required parameters

#### Diagnosis

```bash
# 1. Test the failing endpoint
INTEGRATION_TESTS=1 go test ./tests/integration/... -v -run TestSimpleSmokeTests

# 2. Check NBA.com directly (use browser dev tools)
# Visit: https://www.nba.com/stats/players/traditional
# Open Network tab, find stats.nba.com API calls
# Compare parameters and response structure

# 3. Check endpoint file
# Example: pkg/stats/endpoints/playercareerstats.go
```

#### Fix Process

**Option A: Parameter Change**
```bash
# 1. Update request struct in endpoint file
# pkg/stats/endpoints/ENDPOINT_NAME.go

# 2. Update parameter building in function

# 3. Test
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# 4. Update HTTP handler if needed
# cmd/nba-api-server/handlers_*.go
```

**Option B: Response Structure Change**
```bash
# 1. Update response struct in endpoint file

# 2. Update JSON tags to match new field names

# 3. Test parsing
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# 4. Update HTTP handler response mapping if needed
```

**Option C: Endpoint Deprecated/Removed**
```bash
# 1. Mark endpoint as deprecated in code comments

# 2. Update documentation

# 3. Find replacement endpoint from NBA.com

# 4. Implement replacement (follow "Adding New Endpoint" below)
```

---

### 3. Adding a New Endpoint

**Trigger**: User request or NBA.com adds new endpoint
**Time**: 30-60 minutes (manual) or 5-10 minutes (generated)

#### Manual Approach (Simple Endpoints)

```bash
# 1. Create endpoint file
cp pkg/stats/endpoints/playercareerstats.go pkg/stats/endpoints/newendpoint.go

# 2. Update:
#    - Type names
#    - Request struct
#    - Response struct
#    - JSON tags
#    - URL endpoint
#    - Parameters

# 3. Test
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# 4. Add HTTP handler (optional)
# See cmd/nba-api-server/handlers_*.go for pattern
```

#### Code Generation Approach (Complex Endpoints)

```bash
# 1. Create metadata file
cd tools/generator/metadata
cp playercareerstats.json newendpoint.json

# 2. Edit JSON metadata with endpoint details

# 3. Generate
cd ../..
go run tools/generator/main.go

# 4. Test
go test ./pkg/stats/endpoints/...
```

---

### 4. Update Documentation

**Frequency**: After significant changes
**Time**: 30-60 minutes

#### Documents to Update

1. **README.md** - Feature list, quick start
2. **CHANGELOG.md** - Version history
3. **API_USAGE.md** - Usage examples
4. **ADRs** - Architectural decisions (if changed)
5. **MAINTENANCE.md** (this file) - New procedures

#### Documentation Checklist

```markdown
- [ ] README reflects current features
- [ ] CHANGELOG has unreleased changes
- [ ] Examples still work (make test-examples)
- [ ] API_USAGE shows common patterns
- [ ] Links are not broken
- [ ] Archived docs are clearly marked
```

---

### 5. Release Process

**Frequency**: When ready for new version
**Time**: 1-2 hours

#### Pre-Release Checklist

```bash
# 1. All tests pass
go test ./...
make test-examples
INTEGRATION_TESTS=1 go test ./tests/integration/... -v

# 2. Linting clean
make lint

# 3. Dependencies updated
go list -u -m all

# 4. CHANGELOG.md updated with release notes

# 5. No TODOs/FIXMEs introduced
grep -r "TODO\|FIXME" pkg/ cmd/
```

#### Release Steps

```bash
# 1. Update version in CHANGELOG.md
# Move [Unreleased] content to [X.Y.Z] section

# 2. Create git tag
git tag -a v0.10.0 -m "Release v0.10.0"

# 3. Push tag
git push origin v0.10.0

# 4. Create GitHub release
# Copy CHANGELOG.md content for this version

# 5. Announce (optional)
# - Update README badges
# - Post in relevant communities
```

---

## Troubleshooting Guide

### Problem: Tests Fail After Dependency Update

```bash
# 1. Check what changed
go mod why golang.org/x/text

# 2. Check Go version compatibility
go version  # Must be 1.21+

# 3. Rollback if needed
go get golang.org/x/text@v0.30.0

# 4. Report issue upstream if bug
```

### Problem: Container Build Fails

```bash
# 1. Test locally
go build ./cmd/nba-api-server

# 2. Check Containerfile syntax
podman build -t nba-api-go:test .

# 3. Check Go version in Containerfile
# Must match go.mod requirement (1.21+)

# 4. Common fixes:
#    - Update golang:1.25-alpine to latest
#    - Ensure go.mod/go.sum copied before build
#    - Check CGO_ENABLED=0 for static binary
```

### Problem: High Memory Usage

```bash
# 1. Profile memory
go test -memprofile=mem.prof ./...
go tool pprof mem.prof

# 2. Check for leaks in HTTP handlers
#    - Ensure contexts are canceled
#    - Check for unclosed response bodies

# 3. Review static data size
ls -lh pkg/stats/static/data/

# 4. If needed, add memory limits in deployment
```

### Problem: Rate Limiting Issues

```bash
# 1. Check current rate limit
# Default: 3 requests/5 seconds per host

# 2. NBA.com may have changed limits
# Test with longer delays

# 3. Update rate limit in code
# internal/middleware/ratelimit.go

# 4. Consider caching responses (future enhancement)
```

---

## Monitoring Recommendations

### What to Monitor (Production)

1. **Error Rate**
   - Alert if >5% of requests fail
   - Check `/metrics` endpoint

2. **Response Time**
   - NBA.com typically responds in 200-500ms
   - Alert if p95 >2 seconds

3. **Health Check**
   - `/health` should always return 200
   - Alert if down for >1 minute

4. **Dependency Vulnerabilities**
   - Enable GitHub Dependabot
   - Review alerts weekly

### Prometheus Metrics (Future)

```go
// TODO: Implement Prometheus metrics
// - http_requests_total
// - http_request_duration_seconds
// - nba_api_errors_total
// - nba_api_rate_limit_hits
```

---

## Emergency Procedures

### Production Is Down

**Response Time**: <15 minutes

```bash
# 1. Check health endpoint
curl https://your-domain.com/health

# 2. Check logs
kubectl logs deployment/nba-api  # or
journalctl -u nba-api -n 100

# 3. Common causes:
#    - NBA.com is down (check status.nba.com)
#    - Rate limited (wait 5 minutes)
#    - Out of memory (check pod/container limits)
#    - Certificate expired (renew SSL cert)

# 4. Quick fixes:
#    - Restart service: systemctl restart nba-api
#    - Rollback: kubectl rollout undo deployment/nba-api
#    - Scale down: kubectl scale deployment/nba-api --replicas=0
```

### NBA.com Changed Multiple Endpoints

**Response Time**: <4 hours

```bash
# 1. Identify affected endpoints
INTEGRATION_TESTS=1 go test ./tests/integration/... -v 2>&1 | tee test-results.txt

# 2. Prioritize by usage
# Check /metrics for most-used endpoints

# 3. Fix high-priority first
# Top 10 endpoints cover ~80% of usage:
# - PlayerCareerStats
# - PlayerGameLog
# - CommonPlayerInfo
# - LeagueLeaders
# - Scoreboard
# - CommonAllPlayers
# - LeagueStandings
# - BoxScoreSummaryV2
# - TeamGameLog
# - PlayerProfileV2

# 4. Deploy incrementally
# Fix 1-2, deploy, test, repeat
```

---

## Maintenance Calendar

### Weekly (15-30 minutes)
- [ ] Review GitHub issues
- [ ] Check Dependabot alerts
- [ ] Monitor error rates (if deployed)

### Monthly (1-2 hours)
- [ ] Run integration tests
- [ ] Review metrics and usage patterns
- [ ] Update documentation if needed
- [ ] Check for NBA.com API changes

### Quarterly (2-4 hours)
- [ ] Update dependencies
- [ ] Run full test suite
- [ ] Review and archive old issues
- [ ] Consider new feature requests
- [ ] Update CHANGELOG.md

### Annually (4-8 hours)
- [ ] Major version review
- [ ] Performance benchmarking
- [ ] Security audit
- [ ] Documentation overhaul
- [ ] Community engagement

---

## Contact & Resources

**Project Repository**: https://github.com/n-ae/nba-api-go

**Documentation**:
- README.md - Quick start
- API_USAGE.md - Usage patterns
- MIGRATION_GUIDE.md - From Python nba_api
- MAINTAINABILITY_ASSESSMENT.md - Architecture review
- ADRs in docs/adr/ - Design decisions

**NBA.com Resources**:
- Stats API: https://stats.nba.com/stats/
- Live API: https://cdn.nba.com/static/json/liveData/
- Official Site: https://www.nba.com/stats

**Community**:
- GitHub Issues - Bug reports & features
- GitHub Discussions - Questions & ideas

---

## Notes for Future Maintainers

### Design Philosophy

1. **Boring Tech Over Shiny**: stdlib > frameworks
2. **Maintainability > Features**: Can one person maintain this?
3. **Explicit > Magic**: No hidden complexity
4. **Types > Tests**: Compile-time safety first

### What NOT to Do

❌ Add frameworks (Gin, Echo, etc.)
❌ Add ORMs or databases (this is a proxy, not a data store)
❌ Complex abstractions (keep it simple)
❌ Microservices (monolith is fine for this scope)
❌ Breaking changes without major version bump

### What to Consider

✅ Caching layer (if performance becomes issue)
✅ OpenAPI spec (for better API docs)
✅ Contract tests (prevent NBA.com drift)
✅ CLI tool (if users request it)
✅ More examples (always valuable)

---

**Last Updated**: 2025-11-05
**Next Review**: 2026-02-05 (quarterly)
