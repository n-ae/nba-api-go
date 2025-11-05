# Release Checklist

This document provides a comprehensive checklist for releasing new versions of nba-api-go. Follow this process for all releases to ensure quality and consistency.

---

## Overview

**When to Release:**
- **Patch (x.x.1)**: Bug fixes, dependency updates, documentation improvements
- **Minor (x.1.0)**: New features, new endpoints, backward-compatible changes
- **Major (2.0.0)**: Breaking API changes, major architectural changes

**Release Frequency:**
- Patch: As needed (urgent bugs)
- Minor: Quarterly or when significant features accumulate
- Major: Only when absolutely necessary (avoid if possible)

---

## Pre-Release Checklist

### 1. Code Quality

- [ ] All tests pass locally
  ```bash
  go test ./...
  ```

- [ ] All examples compile
  ```bash
  make test-examples
  ```

- [ ] Linter passes
  ```bash
  make lint
  ```

- [ ] No compilation warnings
  ```bash
  go build ./...
  ```

- [ ] Code formatted
  ```bash
  gofmt -w .
  git diff  # Should show no changes
  ```

### 2. Testing

- [ ] Unit tests pass
  ```bash
  go test ./pkg/...
  go test ./cmd/...
  ```

- [ ] Integration tests pass (if possible)
  ```bash
  INTEGRATION_TESTS=1 go test ./tests/integration/... -v
  ```

- [ ] Contract tests pass (if fixtures exist)
  ```bash
  go test ./tests/contract/... -v
  ```

- [ ] All examples run without errors
  ```bash
  for dir in examples/*/; do
    echo "Testing $(basename $dir)..."
    go run "$dir/main.go" 2>&1 | head -5
  done
  ```

### 3. Documentation

- [ ] README.md is current
  - Version badges updated
  - Examples work
  - Installation instructions accurate

- [ ] CHANGELOG.md updated
  - Unreleased section populated
  - Breaking changes clearly marked
  - Contributors credited (if applicable)

- [ ] API documentation reviewed
  - docs/API_USAGE.md accurate
  - docs/DEPLOYMENT.md current
  - Code examples work

- [ ] Version references updated
  - CLAUDE.md
  - Other documentation files

### 4. Dependencies

- [ ] Dependencies are current
  ```bash
  go list -u -m all
  ```

- [ ] go.mod is tidy
  ```bash
  go mod tidy
  git diff go.mod go.sum  # Review changes
  ```

- [ ] No known security vulnerabilities
  ```bash
  go list -json -m all | nancy sleuth
  # Or use GitHub Dependabot alerts
  ```

### 5. Version Numbers

- [ ] Version number follows semver
  - MAJOR.MINOR.PATCH format
  - Increment appropriate component

- [ ] Version matches release type
  - Patch: Bug fixes only
  - Minor: New features, backward compatible
  - Major: Breaking changes

---

## Release Process

### Step 1: Prepare CHANGELOG

1. Review all commits since last release
   ```bash
   git log v0.9.0..HEAD --oneline
   ```

2. Move items from Unreleased to new version section
   ```markdown
   ## [Unreleased]

   ## [1.1.0] - 2025-XX-XX

   ### Added
   - Feature 1
   - Feature 2

   ### Changed
   - Change 1

   ### Fixed
   - Bug fix 1
   ```

3. Update version comparison links at bottom
   ```markdown
   [Unreleased]: https://github.com/n-ae/nba-api-go/compare/v1.1.0...HEAD
   [1.1.0]: https://github.com/n-ae/nba-api-go/compare/v1.0.0...v1.1.0
   ```

4. Add upgrade guide if breaking changes
   ```markdown
   ### From 1.0.0 to 1.1.0

   **Breaking Changes**: None

   **Migration Steps:**
   1. Update dependency: `go get github.com/n-ae/nba-api-go@v1.1.0`
   2. Review new features in documentation
   ```

### Step 2: Update Version References

- [ ] Update CLAUDE.md version information
- [ ] Update README.md badges
- [ ] Update any hardcoded version strings

### Step 3: Create Release Commit

1. Commit all release changes
   ```bash
   git add CHANGELOG.md CLAUDE.md README.md
   git commit -m "chore: prepare v1.1.0 release"
   ```

2. Push to main branch
   ```bash
   git push origin main
   ```

### Step 4: Create Git Tag

1. Create annotated tag
   ```bash
   git tag -a v1.1.0 -m "Release v1.1.0"
   ```

2. Push tag to remote
   ```bash
   git push origin v1.1.0
   ```

### Step 5: Create GitHub Release

1. Go to GitHub repository
2. Click "Releases" → "Draft a new release"
3. Select the tag you just created
4. Title: "v1.1.0"
5. Description: Copy from CHANGELOG.md and enhance with:
   - Overview paragraph
   - Key highlights
   - Breaking changes (if any)
   - Upgrade instructions
   - Known issues
   - Contributors

Example template:
```markdown
## What's New

Brief overview of this release...

### Highlights

- Feature 1: Description
- Feature 2: Description
- Bug fix: Description

### Breaking Changes

None. This release is fully backward compatible.

### Upgrade Instructions

\`\`\`bash
go get github.com/n-ae/nba-api-go@v1.1.0
go mod tidy
\`\`\`

See [CHANGELOG.md](CHANGELOG.md) for full details.
```

6. Attach artifacts (if any):
   - Compiled binaries for nba-api-server (optional)
   - Checksums

7. Click "Publish release"

---

## Post-Release Checklist

### 1. Verify Release

- [ ] GitHub release is published
- [ ] Tag is visible on GitHub
- [ ] Release notes are complete
- [ ] Go package is available
  ```bash
  go get github.com/n-ae/nba-api-go@v1.1.0
  ```

### 2. Update Documentation

- [ ] Update any external documentation
- [ ] Update website (if exists)
- [ ] Update package registry metadata (if applicable)

### 3. Communicate Release

- [ ] GitHub Discussions announcement (optional)
- [ ] Social media (optional)
- [ ] Email to users (if mailing list exists)

### 4. Monitor

- [ ] Watch for bug reports
- [ ] Monitor GitHub issues
- [ ] Check for user feedback
- [ ] Verify CI/CD passes

### 5. Cleanup

- [ ] Archive old release notes (if needed)
- [ ] Update project board (if used)
- [ ] Close related issues
- [ ] Close related pull requests

---

## Rollback Process

If critical issues are discovered after release:

### Option 1: Patch Release (Preferred)

1. Fix the issue
2. Create patch release (e.g., v1.1.1)
3. Follow normal release process
4. Document issue in CHANGELOG

### Option 2: Retract Release (Severe Issues)

1. Add retraction to go.mod
   ```go
   retract v1.1.0  // Critical bug, use v1.1.1 instead
   ```

2. Document in CHANGELOG
   ```markdown
   ## [1.1.0] - 2025-XX-XX [RETRACTED]

   **This version has been retracted due to [issue]. Use v1.1.1 instead.**
   ```

3. Update GitHub release to mark as retracted
4. Create fixed version immediately

### Option 3: Delete Tag (Extreme Cases Only)

Only if release was never announced and not yet used:

```bash
git tag -d v1.1.0
git push origin :refs/tags/v1.1.0
```

Then fix and re-release with same version.

---

## Release Types in Detail

### Patch Release (1.0.1, 1.0.2)

**When to Use:**
- Bug fixes
- Documentation improvements
- Security patches
- Dependency updates (patch versions)

**Requirements:**
- No new features
- No breaking changes
- Fully backward compatible

**Timeline:** As needed (can be immediate for critical bugs)

**Example CHANGELOG:**
```markdown
## [1.0.1] - 2025-11-10

### Fixed
- Fixed panic in PlayerCareerStats when response is empty
- Corrected documentation typo in README

### Changed
- Updated golang.org/x/text to v0.30.1 (security patch)
```

### Minor Release (1.1.0, 1.2.0)

**When to Use:**
- New features
- New endpoints
- New capabilities
- Deprecations (with backward compatibility)

**Requirements:**
- Backward compatible
- No breaking changes
- Comprehensive testing
- Documentation updates

**Timeline:** Quarterly or when features accumulate

**Example CHANGELOG:**
```markdown
## [1.1.0] - 2026-02-05

### Added
- OpenAPI specification for HTTP API
- 5 new example programs
- Performance profiling tools

### Changed
- Improved error messages for rate limiting
- Enhanced logging in HTTP server

### Deprecated
- OldFunction() - Use NewFunction() instead (will be removed in v2.0.0)
```

### Major Release (2.0.0, 3.0.0)

**When to Use:**
- Breaking API changes
- Major architectural changes
- Removal of deprecated features
- Go version requirement increase

**Requirements:**
- Comprehensive migration guide
- Deprecation warnings in previous version
- Extended testing period
- Clear communication to users

**Timeline:** Only when absolutely necessary (avoid if possible)

**Example CHANGELOG:**
```markdown
## [2.0.0] - 2026-XX-XX

**BREAKING CHANGES** - Read migration guide carefully.

### Added
- New architecture for parameter handling

### Changed
- **BREAKING**: Renamed all endpoint functions (Get prefix removed)
- **BREAKING**: Changed error handling approach
- **BREAKING**: Minimum Go version now 1.22+

### Removed
- **BREAKING**: Removed deprecated functions (OldFunction, etc.)
- **BREAKING**: Removed legacy parameter types

### Migration Guide

See [MIGRATION_v2.md](docs/MIGRATION_v2.md) for detailed upgrade instructions.

**Estimated Migration Time:** 2-4 hours for typical projects
```

---

## Special Release Scenarios

### Security Release

**High Priority - Release Immediately**

1. Fix security issue in private
2. Coordinate with security researchers (if external)
3. Prepare patch release
4. Test thoroughly
5. Release without pre-announcement
6. Announce after release with:
   - CVE number (if applicable)
   - Severity rating
   - Affected versions
   - Upgrade instructions
   - Workarounds (if any)

### Dependency Security Update

1. Update vulnerable dependency
2. Run full test suite
3. Create patch release
4. Document in CHANGELOG:
   ```markdown
   ### Security
   - Updated golang.org/x/text to v0.30.1 (fixes CVE-2024-XXXXX)
   ```

### NBA.com API Breaking Change

1. Update SDK endpoint
2. Update HTTP handler
3. Re-run contract tests
4. Update fixtures
5. Create patch or minor release (depends on impact)
6. Document in CHANGELOG:
   ```markdown
   ### Changed
   - Updated PlayerCareerStats to match NBA.com API changes (2025-11-XX)
   ```

---

## Automation Opportunities

**Future Improvements:**

Consider automating these steps:
- [ ] Automated version bumping
- [ ] CHANGELOG generation from commits
- [ ] GitHub release creation from tags
- [ ] Go package publishing
- [ ] Notification to users

**CI/CD Integration:**
- [ ] Automated testing on tag push
- [ ] Binary building for releases
- [ ] Container image publishing
- [ ] Documentation deployment

---

## Common Mistakes to Avoid

- [ ] Forgetting to update CHANGELOG.md
- [ ] Not testing examples before release
- [ ] Using wrong version number format
- [ ] Not updating version comparison links
- [ ] Releasing with failing tests
- [ ] Missing breaking change documentation
- [ ] Not creating annotated tags
- [ ] Forgetting to push tags to remote
- [ ] Not verifying Go package availability
- [ ] Skipping upgrade guide for major versions

---

## Quick Reference

**Version Format:** MAJOR.MINOR.PATCH

**Version Increment:**
- Patch: Bug fixes only
- Minor: New features, backward compatible
- Major: Breaking changes

**Required Files to Update:**
- CHANGELOG.md (always)
- CLAUDE.md (version references)
- README.md (badges, examples)
- Any other files with hardcoded versions

**Git Commands:**
```bash
git tag -a v1.x.x -m "Release v1.x.x"
git push origin main
git push origin v1.x.x
```

---

## Support

**Questions?** See:
- docs/MAINTENANCE.md - General maintenance guidance
- CONTRIBUTING.md - Contribution guidelines
- GitHub Discussions - Ask the community

---

**Document Version**: 1.0
**Last Updated**: 2025-11-05
**Next Review**: 2026-02-05
