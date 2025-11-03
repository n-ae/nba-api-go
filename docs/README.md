# Documentation Index

Complete documentation for the nba-api-go project.

## 📖 User Documentation

### Getting Started
- **[Main README](../README.md)** - Project overview, features, quick start
- **[API Usage Guide](./API_USAGE.md)** - HTTP API server usage with Python/JavaScript examples
- **[Migration Guide](./MIGRATION_GUIDE.md)** - Migrating from Python nba_api to Go

### Operations
- **[Deployment Guide](../DEPLOYMENT.md)** - Production deployment (systemd, Docker, cloud platforms)
- **[Benchmarks](./BENCHMARKS.md)** - Performance analysis and optimization data

### Development
- **[Contributing Guide](../CONTRIBUTING.md)** - How to contribute to the project
- **[Code Generation Guide](../MANUAL_REGENERATION_GUIDE.md)** - Regenerating endpoint code
- **[Maintainability Assessment](../MAINTAINABILITY.md)** - Project maintainability analysis

## 🏗️ Architecture

### Architecture Decision Records (ADR)
- **[ADR-001: Go Replication Strategy](./adr/001-go-replication-strategy.md)** - Overall strategy and implementation phases
- **[ADR-002: API Server Architecture](./adr/002-api-server-architecture.md)** - HTTP API server design decisions

## 📦 Component Documentation

### Examples
- **[HTTP API Client Examples](../examples/http-api-client/README.md)** - Python, JavaScript, Bash examples
- **[Type Safety Demo](../examples/type_safety_demo/README.md)** - Type inference demonstration

### Testing
- **[Integration Tests](../tests/integration/README.md)** - SDK integration test suite
- **[HTTP API Tests](../tests/http-api/README.md)** - HTTP server integration tests

### Tools
- **[Code Generator](../tools/generator/README.md)** - Endpoint generator documentation

## 🗄️ Archived Documentation

Historical documentation moved to [archive/](./archive/) directory:
- Project summaries from development phases
- Implementation status tracking
- Roadmap (now integrated into main README)
- Type inference implementation notes
- Previous maintainability assessments

These are kept for reference but may be outdated.

## Quick Links by Use Case

### I want to...

**Use the Go SDK**
→ Start with [Main README](../README.md#quick-start---go-sdk)

**Use the HTTP API from Python/JavaScript**
→ Read [API Usage Guide](./API_USAGE.md)

**Migrate from Python nba_api**
→ Follow [Migration Guide](./MIGRATION_GUIDE.md)

**Deploy to production**
→ Use [Deployment Guide](../DEPLOYMENT.md)

**Contribute code**
→ Check [Contributing Guide](../CONTRIBUTING.md) and [ADR-001](./adr/001-go-replication-strategy.md)

**Regenerate endpoint code**
→ Follow [Code Generation Guide](../MANUAL_REGENERATION_GUIDE.md)

**Understand performance**
→ Review [Benchmarks](./BENCHMARKS.md)

**Understand architecture**
→ Read [ADR-001](./adr/001-go-replication-strategy.md) and [ADR-002](./adr/002-api-server-architecture.md)

## Documentation Standards

All documentation in this project follows these principles:

1. **Clarity over completeness** - Concise, actionable information
2. **Examples first** - Show, then explain
3. **Maintainability** - Easy to update, no duplication
4. **User-focused** - Written for developers using the library
5. **Up-to-date** - Obsolete docs moved to archive/

## Contributing to Documentation

Found an issue or want to improve documentation? Please:

1. Check if the document is in archive/ (historical reference only)
2. For active docs, submit a PR with improvements
3. Follow existing formatting and tone
4. Update this index if adding new docs

See [CONTRIBUTING.md](../CONTRIBUTING.md) for details.
