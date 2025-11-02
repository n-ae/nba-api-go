# ADR 002: HTTP API Server Architecture

## Status

Proposed

## Context

The nba-api-go project currently provides a **Go SDK** for accessing NBA statistics. To enable non-Go applications (Python, JavaScript, etc.) to use this library, we need to provide an **HTTP API wrapper**.

### Requirements

1. **SDK remains pure** - Go library for direct Go integration
2. **HTTP API available** - REST endpoints for non-Go apps
3. **Containerized deployment** - Docker image for easy deployment
4. **Minimal maintenance burden** - Solo engineer maintainability

### Constraints (Maintainable-First)

- **No frameworks** - stdlib `net/http` only (reduce dependencies)
- **No ORM** - This is a proxy, not a database app
- **Single binary** - Easy deployment and distribution
- **Boring tech** - Proven patterns, no experiments
- **Clear separation** - SDK (pkg/) and API (cmd/) are independent

## Decision

### Architecture: Three-Layer Approach

```
┌─────────────────────────────────────────┐
│  Non-Go Applications (Python, JS, etc) │
└──────────────────┬──────────────────────┘
                   │ HTTP/REST
                   ▼
┌─────────────────────────────────────────┐
│    HTTP API Server (cmd/nba-api-server) │
│  - REST endpoints                       │
│  - Request validation                   │
│  - JSON serialization                   │
│  - Thin wrapper only                    │
└──────────────────┬──────────────────────┘
                   │ Go SDK calls
                   ▼
┌─────────────────────────────────────────┐
│      Go SDK (pkg/)                      │
│  - stats.Client                         │
│  - endpoints.*                          │
│  - Type-safe operations                 │
└──────────────────┬──────────────────────┘
                   │ HTTPS
                   ▼
┌─────────────────────────────────────────┐
│         NBA.com APIs                    │
└─────────────────────────────────────────┘
```

### Implementation Strategy

#### 1. SDK Layer (Already Exists)

**Location:** `pkg/`
**Purpose:** Type-safe Go library
**Usage:** Import directly in Go applications

```go
import "github.com/n-ae/nba-api-go/pkg/stats"

client := stats.NewDefaultClient()
resp, err := endpoints.PlayerGameLog(ctx, client, req)
```

**No changes needed** - SDK is already clean and well-separated.

#### 2. HTTP API Server (New)

**Location:** `cmd/nba-api-server/`
**Tech Stack:**
- `net/http` - HTTP server (stdlib)
- `encoding/json` - JSON encoding (stdlib)
- `context` - Request context (stdlib)

**Responsibilities:**
- HTTP request handling
- JSON request/response serialization
- Parameter validation
- Calling SDK methods
- Error formatting

**NOT responsible for:**
- NBA API communication (SDK does this)
- Data transformation (SDK returns typed data)
- Caching (keep simple, add later if needed)

#### 3. Container Deployment (New)

**File:** `Containerfile` (Podman-compatible)
**Base image:** `golang:1.21-alpine` (build) → `alpine:latest` (runtime)
**Multi-stage:** Yes (small final image)
**Exposed port:** 8080

### API Design (REST)

#### Endpoint Pattern

```
GET /api/v1/stats/{endpoint}?param1=value1&param2=value2
```

Examples:
```
GET /api/v1/stats/playergamelog?PlayerID=2544&Season=2023-24
GET /api/v1/stats/commonallplayers?Season=2023-24
GET /api/v1/stats/scoreboard?GameDate=2024-01-15
```

#### Response Format

```json
{
  "success": true,
  "data": {
    "PlayerGameLog": [
      { "GameID": "...", "PTS": 30, ... }
    ]
  },
  "meta": {
    "endpoint": "playergamelog",
    "timestamp": "2024-11-01T12:00:00Z"
  }
}
```

Error format:
```json
{
  "success": false,
  "error": {
    "code": "invalid_parameter",
    "message": "PlayerID is required"
  }
}
```

#### Health Check

```
GET /health
```

Response:
```json
{
  "status": "healthy",
  "version": "0.1.0",
  "endpoints_count": 79
}
```

### File Structure

```
nba-api-go/
├── pkg/                          # SDK (no changes)
│   ├── stats/
│   ├── live/
│   └── models/
├── cmd/
│   └── nba-api-server/          # NEW - HTTP API
│       ├── main.go              # Server entrypoint
│       ├── handlers/            # HTTP handlers
│       │   ├── stats.go         # Stats endpoints
│       │   ├── health.go        # Health check
│       │   └── middleware.go    # Logging, CORS, etc.
│       └── server/
│           └── server.go        # HTTP server config
├── Containerfile                # NEW - Docker image
├── docker-compose.yml           # NEW - Local dev (optional)
└── docs/
    └── api/
        └── openapi.yaml         # FUTURE - API spec
```

### Technology Decisions

#### ✅ Use Standard Library

**Rationale:** Minimal dependencies = minimal maintenance

- `net/http` - Mature, stable, good enough
- `encoding/json` - Fast enough for this use case
- `context` - Standard cancellation

**NOT using:**
- ❌ Gin/Echo/Fiber - Adds dependency, minimal benefit
- ❌ gRPC - Overkill for REST proxy
- ❌ GraphQL - Unnecessary complexity

#### ✅ Single Binary Deployment

**Build:** `go build -o nba-api-server ./cmd/nba-api-server`
**Run:** `./nba-api-server`

**Configuration:** Environment variables only
- `PORT` - HTTP port (default: 8080)
- `NBA_API_TIMEOUT` - Request timeout (default: 30s)
- `LOG_LEVEL` - Logging level (default: info)

#### ✅ Container-First for Non-Go Users

**Why Podman/Docker:**
- Language-agnostic deployment
- Easy to integrate with Python, Node.js, etc.
- Standard deployment pattern

**Image size target:** < 20MB (multi-stage build)

### Deployment Patterns

#### Pattern 1: Direct Go Import (Recommended for Go apps)

```go
import "github.com/n-ae/nba-api-go/pkg/stats"
// Use SDK directly
```

**Pros:** Type-safe, fastest, no network hop
**Cons:** Go only

#### Pattern 2: Sidecar Container (Recommended for non-Go apps)

```yaml
# docker-compose.yml
services:
  nba-api:
    image: nba-api-go:latest
    ports:
      - "8080:8080"

  my-python-app:
    build: .
    environment:
      - NBA_API_URL=http://nba-api:8080
```

**Pros:** Language-agnostic, local network, fast
**Cons:** Requires container orchestration

#### Pattern 3: Shared API Server (Optional)

Deploy as centralized service for multiple apps.

**Pros:** One deployment for many consumers
**Cons:** SPOF, rate limiting concerns, more ops

### Security Considerations

#### For MVP (Phase 1)

- ✅ Rate limiting (inherit from SDK)
- ✅ Request timeout
- ✅ Input validation
- ✅ CORS headers (configurable)

#### Future (Phase 2+)

- Authentication (API keys)
- Request logging
- Metrics (Prometheus)
- Tracing (OpenTelemetry)

**Decision:** Start simple, add complexity only when needed.

### Maintenance Burden Analysis

#### SDK Only (Current)

- **Maintenance:** Low
- **Deps:** 2 (golang.org/x/text, golang.org/x/time)
- **Code:** ~79 endpoints + client
- **Testing:** Unit tests
- **Deployment:** `go get`

#### SDK + API Server (Proposed)

- **Maintenance:** Medium
- **Deps:** 2 (same - no new deps!)
- **Code:** +~500 LOC (handlers + server)
- **Testing:** Unit + integration (HTTP)
- **Deployment:** `go get` OR `docker run`

**Verdict:** Acceptable maintenance increase for significant value (language-agnostic access)

### Migration Path

#### Phase 1: Core API (This Session)

1. Create `cmd/nba-api-server/main.go`
2. Implement basic HTTP server
3. Add 5-10 most-used endpoints
4. Create Containerfile
5. Documentation

**Time estimate:** 2-3 hours
**Coverage:** Top 10 endpoints

#### Phase 2: Full Coverage (Future)

1. Generate handlers for all 79 endpoints
2. Add OpenAPI spec
3. Add Prometheus metrics
4. Performance testing

**Time estimate:** 4-6 hours
**Coverage:** All endpoints

#### Phase 3: Production Hardening (Future)

1. Authentication
2. Distributed tracing
3. Caching layer
4. Helm charts (if needed)

**Time estimate:** 8+ hours

### Open Questions

1. **Should API mirror all 79 endpoints?**
   - Decision: Start with top 10, add on demand

2. **Rate limiting strategy?**
   - Decision: Inherit from SDK (already implemented)

3. **Caching layer needed?**
   - Decision: No - keep simple, add later if users request

4. **Authentication required?**
   - Decision: No for MVP - NBA.com data is public

## Consequences

### Positive

1. **Language-agnostic** - Python, JS, Ruby apps can use NBA data
2. **Low complexity** - Stdlib only, no frameworks
3. **Easy deployment** - Single binary or container
4. **SDK remains clean** - No API concerns in pkg/
5. **Optional** - Go users can ignore API, use SDK directly

### Negative

1. **More code to maintain** - +500 LOC
2. **HTTP overhead** - Slower than direct SDK calls
3. **Another deployment option** - More documentation needed

### Neutral

1. **Two usage patterns** - SDK vs API (clear documentation mitigates)
2. **Container dependency** - Only for non-Go users

## Implementation Checklist

### MVP (Phase 1)

- [ ] Create `cmd/nba-api-server/main.go`
- [ ] Implement HTTP server with stdlib
- [ ] Add health check endpoint
- [ ] Add 5-10 high-priority stats endpoints
- [ ] Create Containerfile (multi-stage)
- [ ] Write API documentation
- [ ] Add usage examples (curl, Python, JavaScript)
- [ ] Test container build and run

### Future Enhancements

- [ ] Generate all 79 endpoint handlers
- [ ] OpenAPI/Swagger spec
- [ ] Prometheus metrics endpoint
- [ ] Request logging
- [ ] API authentication (optional)
- [ ] Helm chart (if needed)

## References

- [Go stdlib net/http](https://pkg.go.dev/net/http)
- [Twelve-Factor App](https://12factor.net/)
- [Boring Technology](https://boringtechnology.club/)
- [ADR 001: Go Replication Strategy](./001-go-replication-strategy.md)
