# Contributing to nba-api-go

Thank you for your interest in contributing to nba-api-go! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Architecture](#architecture)
- [Adding New Endpoints](#adding-new-endpoints)
- [Testing](#testing)
- [Code Style](#code-style)
- [Submitting Changes](#submitting-changes)

## Code of Conduct

Be respectful, collaborative, and constructive in all interactions.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/nba-api-go.git`
3. Add upstream remote: `git remote add upstream https://github.com/username/nba-api-go.git`
4. Create a feature branch: `git checkout -b feature/your-feature-name`

## Development Setup

### Prerequisites

- Go 1.21 or later
- Make (optional, for convenience commands)
- golangci-lint (for linting)

### Install Dependencies

```bash
go mod download
```

### Install Development Tools

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Run Tests

```bash
make test
# or
go test ./...
```

### Run Linters

```bash
make lint
# or
golangci-lint run ./...
```

## Architecture

Please review the [Architecture Decision Record](./docs/adr/001-go-replication-strategy.md) for detailed architectural guidance.

### Key Principles

1. **Type Safety** - Use strongly typed structs for requests and responses
2. **Context Support** - All network operations accept `context.Context`
3. **Immutability** - Clients and configurations should be immutable after creation
4. **Error Handling** - Return descriptive errors with proper wrapping
5. **Testing** - All new code should have comprehensive tests

### Project Structure

```
nba-api-go/
├── pkg/
│   ├── client/          # Core HTTP client
│   ├── stats/           # NBA Stats API
│   ├── live/            # NBA Live Data API
│   ├── models/          # Shared models and errors
│   └── stats/
│       ├── endpoints/   # Stats API endpoints
│       ├── parameters/  # Parameter types
│       └── static/      # Static data (players/teams)
├── internal/
│   └── middleware/      # HTTP middleware
├── examples/            # Usage examples
└── docs/
    └── adr/            # Architecture Decision Records
```

## Adding New Endpoints

### Stats API Endpoint

1. **Analyze the Python implementation** in [nba_api](https://github.com/swar/nba_api)
   - Identify endpoint URL and parameters
   - Understand response structure

2. **Create endpoint file** in `pkg/stats/endpoints/`

```go
package endpoints

import (
    "context"
    "net/url"

    "github.com/username/nba-api-go/pkg/models"
    "github.com/username/nba-api-go/pkg/stats"
)

type YourEndpointRequest struct {
    RequiredParam string
    OptionalParam *string
}

type YourEndpointResponse struct {
    ResultSet []YourDataType `json:"resultSets"`
}

func YourEndpoint(ctx context.Context, client *stats.Client, req YourEndpointRequest) (*models.Response[*YourEndpointResponse], error) {
    params := url.Values{}
    params.Set("RequiredParam", req.RequiredParam)
    if req.OptionalParam != nil {
        params.Set("OptionalParam", *req.OptionalParam)
    }

    var resp YourEndpointResponse
    if err := client.GetJSON(ctx, "/yourendpoint", params, &resp); err != nil {
        return nil, err
    }

    return models.NewResponse(&resp, 200, "", nil), nil
}
```

3. **Add tests** in `pkg/stats/endpoints/yourendpoint_test.go`

4. **Add example** in `examples/your_endpoint/main.go`

5. **Document** in README.md

### Live API Endpoint

Follow similar pattern in `pkg/live/endpoints/`

## Testing

### Test Requirements

- Unit tests for all public functions
- Table-driven tests for parameter validation
- Mock HTTP responses for endpoint tests
- Example code that compiles and runs

### Writing Tests

```go
func TestYourFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    OutputType
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   InputType{...},
            want:    OutputType{...},
            wantErr: false,
        },
        {
            name:    "invalid input",
            input:   InputType{...},
            want:    OutputType{},
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := YourFunction(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Run Tests with Coverage

```bash
make test-coverage
# Opens coverage.html in browser
```

## Code Style

### Go Standards

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` and `goimports`
- Run `golangci-lint`

### Naming Conventions

- **Exported functions**: PascalCase (e.g., `PlayerCareerStats`)
- **Unexported functions**: camelCase (e.g., `parseResponse`)
- **Constants**: PascalCase (e.g., `DefaultTimeout`)
- **Acronyms**: Uppercase (e.g., `HTTPClient`, `URLPath`, `NBAStats`)

### Comments

- All exported types, functions, and constants must have doc comments
- Comments should be complete sentences
- Start with the name of the thing being described

```go
// PlayerCareerStats retrieves career statistics for a given player.
// It returns season-by-season stats and career totals for regular season,
// playoffs, all-star games, and college.
func PlayerCareerStats(ctx context.Context, client *stats.Client, req PlayerCareerStatsRequest) (*models.Response[*PlayerCareerStatsResponse], error) {
    // ...
}
```

### Error Messages

- Start with lowercase
- Don't end with punctuation
- Use `fmt.Errorf` with `%w` to wrap errors
- Provide context in error messages

```go
if err := validate(req); err != nil {
    return nil, fmt.Errorf("invalid request: %w", err)
}
```

### Imports

Group imports in this order:
1. Standard library
2. External packages
3. Internal packages

```go
import (
    "context"
    "fmt"

    "golang.org/x/text/transform"

    "github.com/username/nba-api-go/pkg/models"
    "github.com/username/nba-api-go/pkg/stats"
)
```

## Submitting Changes

### Before Submitting

1. **Run tests**: `make test`
2. **Run linters**: `make lint`
3. **Format code**: `make fmt`
4. **Update documentation** if needed
5. **Add/update examples** if adding new features

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat: add PlayerGameLog endpoint`
- `fix: correct parameter validation for Season`
- `docs: update README with new endpoint examples`
- `test: add tests for team search functionality`
- `refactor: simplify middleware chain construction`
- `chore: update dependencies`

### Pull Request Process

1. **Update your fork**:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

2. **Push to your fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

3. **Create Pull Request** on GitHub

4. **Describe your changes**:
   - What does this PR do?
   - Why is this change needed?
   - Any breaking changes?
   - Related issues?

5. **Wait for review** - Maintainers will review and provide feedback

### PR Checklist

- [ ] Tests pass locally (`make test`)
- [ ] Linters pass (`make lint`)
- [ ] Code is formatted (`make fmt`)
- [ ] Documentation updated
- [ ] Examples added/updated
- [ ] Commit messages follow convention
- [ ] No breaking changes (or clearly documented)

## Development Workflow

### Typical Workflow

1. Check out main and pull latest changes
   ```bash
   git checkout main
   git pull upstream main
   ```

2. Create feature branch
   ```bash
   git checkout -b feature/new-endpoint
   ```

3. Make changes and test
   ```bash
   # Make changes
   make test
   make lint
   ```

4. Commit with conventional message
   ```bash
   git add .
   git commit -m "feat: add new endpoint for player game logs"
   ```

5. Push and create PR
   ```bash
   git push origin feature/new-endpoint
   ```

## Questions?

- Open an issue for bugs or feature requests
- Start a discussion for questions or ideas
- Reference the [nba_api Python library](https://github.com/swar/nba_api) for endpoint behavior

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
