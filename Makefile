.PHONY: help test test-coverage test-examples build build-server clean lint fmt vet examples

GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
SERVER_LDFLAGS := -X main.gitCommit=$(GIT_COMMIT) -X main.buildTime=$(BUILD_TIME)

help:
	@echo "Available targets:"
	@echo "  test          - Run all tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  test-examples - Build all example programs to verify they compile"
	@echo "  build         - Build all examples"
	@echo "  build-server  - Build nba-api-server with git commit/build time embedded"
	@echo "  clean         - Remove build artifacts"
	@echo "  lint          - Run golangci-lint"
	@echo "  fmt           - Format code with gofmt"
	@echo "  vet           - Run go vet"
	@echo "  examples      - Run all examples"

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-examples:
	@echo "Building all example programs..."
	@failed=0; \
	for dir in examples/*/; do \
		if [ -f "$$dir/main.go" ]; then \
			example=$$(basename "$$dir"); \
			printf "  %-30s" "$$example"; \
			if go build -o /dev/null "./$$dir" 2>/dev/null; then \
				echo "✓ PASS"; \
			else \
				echo "✗ FAIL"; \
				failed=$$((failed + 1)); \
			fi; \
		fi; \
	done; \
	if [ $$failed -eq 0 ]; then \
		echo "\n✓ All examples compiled successfully!"; \
	else \
		echo "\n✗ $$failed example(s) failed to compile"; \
		exit 1; \
	fi

build:
	@echo "Building examples..."
	go build -o bin/player_stats ./examples/player_stats
	go build -o bin/scoreboard ./examples/scoreboard
	go build -o bin/player_search ./examples/player_search
	@echo "Binaries built in bin/"

build-server:
	@echo "Building nba-api-server ($(GIT_COMMIT))..."
	go build -ldflags "$(SERVER_LDFLAGS)" -o bin/nba-api-server ./cmd/nba-api-server
	@echo "Binary built: bin/nba-api-server"

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean

lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest" && exit 1)
	golangci-lint run ./...

fmt:
	gofmt -s -w .

vet:
	go vet ./...

examples: build
	@echo "\n=== Running player_search example ==="
	./bin/player_search
	@echo "\n=== Running scoreboard example (may fail if no games today) ==="
	-./bin/scoreboard
	@echo "\nNote: player_stats example requires valid player ID and network access"
