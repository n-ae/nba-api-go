# Clarification: SDK vs HTTP API Server

## 📊 Current State

### SDK (Go Library) - pkg/stats/endpoints/
- **Total Endpoints:** 139/139 (100%) ✅
- **Type:** Go package for direct import
- **Usage:** `import "github.com/.../pkg/stats/endpoints"`
- **Integration Tests:** 65/139 (46.8%)
  - Player endpoints: 35/35 (100%) ✅
  - Team endpoints: 30/30 (100%) ✅
  - League/others: 74 remaining

### HTTP API Server - cmd/nba-api-server/
- **Total Endpoints:** 10/139 (7.2%) ⚠️
- **Type:** REST API for any language
- **Usage:** `curl http://localhost:8080/api/v1/stats/...`
- **Exposed Endpoints:**
  1. playergamelog
  2. commonallplayers
  3. scoreboardv2
  4. leaguestandings
  5. commonteamroster
  6. playercareerstats
  7. leagueleaders
  8. commonplayerinfo
  9. leaguedashteamstats
  10. leaguedashplayerstats

---

## 🎯 What This Means

### Integration Tests Created
The 65 integration tests I created are for the **SDK** (Go library), not the HTTP API server.

**SDK Integration Tests (65/139):**
- ✅ Test that SDK can call NBA.com API directly
- ✅ Test response parsing in Go
- ✅ Test type safety
- ✅ Validate Go package functionality

**These tests do NOT test the HTTP API server.**

### HTTP API Server Status
The API server only exposes 10 of the 139 SDK endpoints as REST endpoints.

**To use the other 129 endpoints:**
- Option 1: Import SDK directly in Go code
- Option 2: Add more handlers to API server

---

## 🔧 Two Separate Work Streams

### Stream 1: SDK Integration Tests (In Progress)
**Goal:** Test all 139 SDK endpoints
**Progress:** 65/139 (46.8%)
**Status:**
- ✅ Player endpoints: 35/35 (100%)
- ✅ Team endpoints: 30/30 (100%)
- ⏳ League endpoints: 0/28 (0%)
- ⏳ Box scores: 0/10 (0%)
- ⏳ Others: 0/46 (0%)

### Stream 2: HTTP API Server Expansion (Not Started)
**Goal:** Expose more SDK endpoints as REST APIs
**Current:** 10/139 (7.2%)
**Potential work:**
- Add 20-30 more common endpoints
- Add all 139 endpoints (comprehensive)
- Add batch/bulk operations
- Add filtering/pagination

---

## 💡 Should We Expand HTTP API Server?

### Option A: Keep Minimal (Current)
**Pros:**
- Low maintenance
- Common endpoints covered
- Users can add handlers as needed

**Cons:**
- Only 10/139 endpoints accessible via HTTP
- Non-Go users have limited access

### Option B: Add 20-30 Most Common
**Pros:**
- Covers 80% of use cases
- Still maintainable
- Good balance

**Cons:**
- ~2-3 hours of work
- More code to maintain

### Option C: Expose All 139
**Pros:**
- Complete parity with SDK
- Any language can use all endpoints
- Maximum flexibility

**Cons:**
- ~10-15 hours of work
- Significant maintenance burden
- Potentially unused endpoints

---

## 🎯 Recommendation

**Focus on SDK integration tests first** (current work stream)

Then evaluate HTTP API server expansion based on:
1. User demand - Are people asking for specific endpoints?
2. Use cases - Which endpoints are most valuable via HTTP?
3. Maintenance - Can we maintain more handlers?

### Immediate Priorities
1. ✅ Complete SDK integration tests (74 remaining)
2. ✅ Document which endpoints are HTTP-exposed
3. ⏳ Migration guide (Python nba_api → Go SDK)
4. ⏳ Consider adding 10-20 more HTTP handlers if needed

---

## 📊 Current Test Coverage Breakdown

### SDK Endpoints (139 total)
- **Integration Tests:** 65/139 (46.8%)
  - Player: 35 ✅
  - Team: 30 ✅
  - League: 0
  - Box Score: 0
  - Game: 0
  - Others: 0

### HTTP API Server (10 endpoints)
- **Unit Tests:** 5/5 (100%) ✅
- **Integration Tests:** 4 suites ✅
- **Handler Coverage:** 10/139 (7.2%)
- **Handler Tests:** 1/10 (PlayerGameLog tested)

---

## 🚀 Next Steps Options

### Option 1: Continue SDK Integration Tests
- Complete league endpoints (28 tests)
- Complete box scores (10 tests)
- Complete game endpoints (12 tests)
- **Time:** 3-4 hours
- **Value:** Validates entire SDK

### Option 2: Expand HTTP API Server
- Add 20 most-used endpoints as HTTP handlers
- Create handler tests for each
- Update API documentation
- **Time:** 3-4 hours
- **Value:** More accessible to non-Go users

### Option 3: Documentation Focus
- Create migration guide (Python → Go)
- Add usage examples (top 20 endpoints)
- API reference generation
- **Time:** 4-5 hours
- **Value:** Better user onboarding

---

## 📝 Clarification Summary

**Question:** Do integration tests include HTTP API too?

**Answer:** NO - The integration tests are for the SDK only.

- **SDK:** 139 endpoints, 65 tested (46.8%)
- **HTTP API:** 10 endpoints exposed, basic tests done

**The 139 endpoints exist in the SDK (Go package), but only 10 are exposed via the HTTP API server.**

If you want to expand the HTTP API server to expose more endpoints, that would be a separate work item.

---

**What would you like to focus on?**
1. Continue SDK integration tests (complete the 74 remaining)
2. Expand HTTP API server (add more REST endpoints)
3. Both (SDK tests + API expansion)
