# ADR 001: Go Replication Strategy for nba_api

## Status

Accepted - Implementation in Progress (Phases 1-3 Complete)

## Context

We are replicating the Python `nba_api` library in Go to provide a type-safe, high-performance NBA API client for the Go ecosystem. The original Python implementation has been analyzed to understand its architecture, patterns, and design decisions.

### Original Python Implementation Analysis

The `nba_api` (v1.10.2) is a mature Python library with the following characteristics:

**Architecture:**
- **Modular structure** with two main API domains:
  - `stats` - Official NBA Stats API (stats.nba.com) with 139 endpoint classes
  - `live` - NBA Live Data API (cdn.nba.com) for real-time game data
- **Base classes** that define common behavior for all endpoints
- **Static data sets** for frequently accessed player and team information
- **HTTP layer abstraction** with session management and debug capabilities

**Key Components:**
1. **Endpoint Classes** - Each NBA.com endpoint is wrapped in a dedicated class
   - Inherit from base `Endpoint` class
   - Define `expected_data` schema for response validation
   - Support multiple output formats (JSON, dict, pandas DataFrame)
   - Handle query parameters with type validation

2. **HTTP Layer** (`library/http.py`)
   - Shared `requests.Session` for connection pooling
   - Proxy support (including random proxy selection)
   - Request/response debugging with file caching
   - Automatic parameter sorting and URL encoding
   - Configurable timeouts and custom headers

3. **Data Models**
   - `NBAResponse` - Wraps HTTP responses with JSON parsing
   - `DataSet` - Container for tabular data with pandas integration
   - Parameter enums - Type-safe parameter values

4. **Static Data** (`stats/static/`)
   - Pre-loaded player and team datasets
   - Search functions with regex support and accent-insensitive matching
   - Active/inactive player filtering

**Design Patterns:**
- **Class-based endpoints** - One class per API endpoint
- **Lazy loading** - Data fetched on instantiation unless `get_request=False`
- **Session reuse** - Class-level session management
- **Optional dependencies** - Pandas is optional for basic usage
- **Debug mode** - File-based response caching for development

**Dependencies:**
- Core: `requests`, `numpy`
- Optional: `pandas` (for DataFrame support)
- Python 3.9+ required

## Decision

We will replicate the nba_api in Go using the following strategy:

### 1. Package Structure

```
nba-api-go/
├── pkg/
│   ├── stats/           # NBA Stats API (stats.nba.com)
│   │   ├── endpoints/   # Endpoint-specific clients
│   │   ├── parameters/  # Parameter types and enums
│   │   └── static/      # Static player/team data
│   ├── live/            # NBA Live API (cdn.nba.com)
│   │   └── endpoints/   # Live data endpoints
│   ├── client/          # Shared HTTP client
│   └── models/          # Common data structures
├── internal/            # Internal utilities
├── cmd/                 # CLI tools (optional)
├── examples/            # Usage examples
└── docs/                # Documentation and ADRs
```

### 2. Core Architecture Decisions

#### 2.1 HTTP Client Layer
- **Use `net/http.Client`** with configurable transport
- **Connection pooling** via default transport settings
- **Context-based timeouts** instead of simple timeout values
- **Middleware pattern** for:
  - Request/response logging
  - Rate limiting
  - Retry logic with exponential backoff
  - Header injection (User-Agent, Referer, custom headers)
- **Interface-based design** for easy mocking and testing

```go
type HTTPClient interface {
    Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

type Client struct {
    httpClient HTTPClient
    baseURL    string
    headers    http.Header
    timeout    time.Duration
}
```

#### 2.2 Endpoint Design
- **Functional options pattern** for endpoint configuration
- **Struct-based requests** with strong typing
- **Generated code** for repetitive endpoint definitions
- **Builder pattern** for complex queries

```go
type PlayerCareerStatsRequest struct {
    PlayerID  string
    PerMode   PerMode
    LeagueID  *string  // pointer for optional params
}

func (c *StatsClient) PlayerCareerStats(ctx context.Context, req PlayerCareerStatsRequest, opts ...RequestOption) (*PlayerCareerStatsResponse, error)
```

#### 2.3 Response Handling
- **Typed responses** using Go structs with JSON tags
- **Embedded metadata** (URL, status code, headers)
- **Generic result container** for common operations

```go
type Response[T any] struct {
    Data       T
    StatusCode int
    URL        string
    Headers    http.Header
}

type PlayerCareerStatsResponse struct {
    CareerTotalsRegularSeason []PlayerCareerStat
    CareerTotalsPostSeason    []PlayerCareerStat
    // ... other result sets
}
```

#### 2.4 Static Data
- **Embedded data files** using `go:embed` for players/teams
- **In-memory index** for fast lookups
- **Lazy initialization** with `sync.Once`
- **Search functions** with regex and fuzzy matching support

```go
//go:embed data/players.json
var playersData []byte

var (
    playersIndex map[string]Player
    indexOnce    sync.Once
)

func FindPlayerByName(name string) ([]Player, error)
func GetPlayerByID(id string) (*Player, error)
```

#### 2.5 Error Handling
- **Custom error types** for different failure modes
- **Wrapped errors** using `fmt.Errorf` with `%w`
- **Sentinel errors** for common cases

```go
var (
    ErrInvalidResponse = errors.New("invalid response format")
    ErrRateLimited     = errors.New("rate limited")
    ErrNotFound        = errors.New("resource not found")
)

type APIError struct {
    StatusCode int
    Message    string
    URL        string
}
```

#### 2.6 Parameter Validation
- **Typed enums** using string constants with custom types
- **Validation functions** in parameter types
- **Required vs optional** using pointers

```go
type PerMode string

const (
    PerModeTotals      PerMode = "Totals"
    PerModePerGame     PerMode = "PerGame"
    PerModePer36       PerMode = "Per36"
)

func (p PerMode) Validate() error {
    switch p {
    case PerModeTotals, PerModePerGame, PerModePer36:
        return nil
    default:
        return fmt.Errorf("invalid PerMode: %s", p)
    }
}
```

### 3. Divergences from Python Implementation

#### 3.1 No Pandas Dependency
- Return native Go slices and structs
- Consider optional CSV/table formatting utilities
- Focus on JSON serialization for data export

#### 3.2 Explicit Context Handling
- All network operations accept `context.Context`
- Enables cancellation and deadline propagation
- Better integration with Go server patterns

#### 3.3 Immutable Configuration
- Endpoints don't modify client state
- Thread-safe by default
- Configuration via functional options

#### 3.4 Code Generation
- Generate endpoint boilerplate from API schema
- Reduce maintenance burden for 139+ endpoints
- Ensure consistency across endpoints

#### 3.5 Testing Strategy
- Unit tests with mocked HTTP responses
- Integration tests with recorded fixtures
- Golden file testing for response parsing
- Table-driven tests for parameter validation

### 4. Development Phases

#### Phase 1: Foundation (Week 1-2) ✅ COMPLETED
- [x] HTTP client implementation with middleware
- [x] Response parsing framework
- [x] Error handling structure
- [x] Core models and types
- [x] Project scaffolding and CI/CD

#### Phase 2: Stats API Core (Week 3-4) ✅ COMPLETED
- [x] 5-10 most common stats endpoints (PlayerCareerStats implemented)
- [x] Parameter types and validation
- [x] Static player/team data (5,135 players, 30 teams)
- [x] Search functionality (accent-insensitive, regex support)
- [x] Comprehensive tests

#### Phase 3: Live API (Week 5) ✅ COMPLETED
- [x] Scoreboard endpoint
- [x] BoxScore endpoint (structure implemented)
- [x] PlayByPlay endpoint (structure implemented)
- [x] Real-time data tests

#### Phase 4: Remaining Stats Endpoints (Week 6-8) 🔄 IN PROGRESS
- [x] Code generation tooling (completed)
- [x] PlayerGameLog endpoint (completed)
- [x] CommonPlayerInfo endpoint (completed)
- [x] LeagueLeaders endpoint (completed)
- [x] TeamGameLog endpoint (completed)
- [x] TeamInfoCommon endpoint (completed)
- [x] BoxScoreSummaryV2 endpoint (completed)
- [x] ShotChartDetail endpoint (completed)
- [x] TeamYearByYearStats endpoint (completed)
- [x] PlayerDashboardByGeneralSplits endpoint (completed)
- [x] TeamDashboardByGeneralSplits endpoint (completed)
- [x] PlayByPlayV2 endpoint (completed)
- [ ] Generate remaining 127 endpoints
- [x] Integration test framework (completed)
- [x] Benchmark tests (completed)
- [ ] Documentation (in progress)

#### Phase 5: Polish (Week 9-10) 🔄 IN PROGRESS
- [ ] CLI tool (optional)
- [x] Usage examples and tutorials
- [x] Performance optimization (benchmarks added)
- [x] Rate limiting implementation
- [x] Performance benchmarking
- [ ] Release preparation

### 5. Technology Choices

- **Go version:** 1.21+ (for improved error handling and generics)
- **HTTP library:** Standard library `net/http`
- **JSON parsing:** Standard library `encoding/json` with custom unmarshaling where needed
- **Testing:** Standard library `testing` + `testify/assert` for assertions
- **Mocking:** `gomock` or interface-based manual mocks
- **Code generation:** `text/template` or dedicated code generator
- **Linting:** `golangci-lint` with strict rules
- **Documentation:** Go doc comments + examples

### 6. API Compatibility Philosophy

**Goal:** Feature parity with Python API, not API signature parity

**Approach:**
- Match endpoint coverage and capabilities
- Embrace Go idioms (contexts, errors, interfaces)
- Improve type safety where Python used dynamic typing
- Maintain response data structure compatibility
- Document migration path for Python users

## Consequences

### Positive

1. **Type Safety** - Compile-time guarantees for API parameters and responses
2. **Performance** - No GIL, better concurrency, faster JSON parsing
3. **Deployment** - Single binary distribution, no dependency management
4. **Concurrency** - Built-in goroutines for parallel API calls
5. **Maintainability** - Static typing reduces runtime errors
6. **Testing** - Excellent testing tools and mocking capabilities
7. **Documentation** - Go doc integrated with code

### Negative

1. **Initial Development Time** - More verbose than Python, type definitions required
2. **No DataFrame Support** - Users need alternative data processing tools
3. **Community Size** - Smaller NBA data analysis community in Go vs Python
4. **Code Generation Dependency** - Need tooling for 139+ endpoints

### Neutral

1. **Different Paradigm** - Go users expect different patterns than Python users
2. **Error Handling** - Explicit error checking vs Python exceptions
3. **Configuration** - Functional options vs keyword arguments

## Notes

- Consider generating OpenAPI/Swagger spec from implemented endpoints
- Monitor nba_api Python repo for endpoint changes and updates
- Evaluate creating a compatibility layer for easier Python-to-Go migration
- Document all NBA.com API quirks and undocumented behaviors discovered
- Consider rate limiting from day one to be a good API citizen

## References

- [nba_api Python Implementation](https://github.com/swar/nba_api)
- [nba_api Documentation](https://github.com/swar/nba_api/blob/master/README.md)
- [NBA.com Terms of Use](https://www.nba.com/termsofuse)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
