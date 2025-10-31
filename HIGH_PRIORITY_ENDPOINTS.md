# High-Priority Endpoints for Next Generation Batch

**Purpose:** Identify most valuable NBA API endpoints to implement next
**Criteria:** User demand, utility, and ease of implementation
**Target:** 10-15 new endpoints to reach 25-30 total (18-22% coverage)

---

## 🎯 Selection Criteria

### Value Factors
1. **User Demand** - Commonly requested in NBA data analysis
2. **Data Utility** - Provides unique/valuable insights
3. **Use Case Coverage** - Fills gaps in current functionality
4. **Complementary** - Works well with existing endpoints

### Effort Factors
1. **Complexity** - Number of parameters and result sets
2. **Data Structure** - Standard vs custom parsing needed
3. **Documentation** - Available metadata/examples
4. **Testing** - Ease of validation

---

## 🏆 Tier 1: Essential Endpoints (Highest Priority)

### 1. LeagueStandings
**Why:** Core functionality - users need current standings
**Complexity:** Low (1-2 result sets, standard fields)
**Use Cases:** Playoff race, division leaders, conference rankings
**Complementary to:** TeamGameLog, TeamInfoCommon

**Metadata Needed:**
```json
{
  "name": "LeagueStandings",
  "endpoint": "leaguestandings",
  "parameters": [
    {"name": "LeagueID", "type": "LeagueID", "required": false},
    {"name": "Season", "type": "Season", "required": true},
    {"name": "SeasonType", "type": "SeasonType", "required": true}
  ],
  "result_sets": [
    {"name": "Standings", "fields": ["TeamID", "TeamCity", "TeamName", "Conference", "Division", "W", "L", "PCT", "HOME", "ROAD", "CONF", "DIV", ...]}
  ]
}
```

### 2. PlayerAwards
**Why:** Player career highlights and accolades
**Complexity:** Low (1 result set)
**Use Cases:** Player profiles, GOAT debates, HOF analysis
**Complementary to:** PlayerCareerStats, CommonPlayerInfo

### 3. TeamYearOverYearStats
**Why:** Historical team performance trends
**Complexity:** Low (1 result set)
**Use Cases:** Dynasty analysis, franchise history
**Complementary to:** TeamYearByYearStats (already done!)

### 4. PlayoffPicture
**Why:** Real-time playoff race information
**Complexity:** Medium (2-3 result sets)
**Use Cases:** Playoff predictions, seeding scenarios
**Complementary to:** LeagueStandings

### 5. PlayerGameScoreLog
**Why:** Advanced per-game metrics (Game Score, PIE, etc.)
**Complexity:** Low (similar to PlayerGameLog)
**Use Cases:** Advanced analytics, player efficiency
**Complementary to:** PlayerGameLog

---

## 📊 Tier 2: High-Value Analytics (Strong Priority)

### 6. TeamAdvancedStats
**Why:** Advanced team metrics (ORtg, DRtg, Pace, etc.)
**Complexity:** Medium (multiple result sets)
**Use Cases:** Deep team analysis, coaching effectiveness
**Complementary to:** TeamDashboards

### 7. PlayerAdvancedStats
**Why:** Advanced player metrics (PER, TS%, USG%, etc.)
**Complexity:** Medium (multiple result sets)
**Use Cases:** MVP analysis, player comparisons
**Complementary to:** PlayerDashboards

### 8. LeagueGameLog
**Why:** All games for a season/date range
**Complexity:** Low (1 result set, similar to TeamGameLogs)
**Use Cases:** Schedule analysis, game summaries
**Complementary to:** LeagueStandings, Scoreboard

### 9. PlayerVsPlayer
**Why:** Head-to-head player matchup stats
**Complexity:** Medium (matchup-specific data)
**Use Cases:** Fantasy basketball, matchup analysis
**Complementary to:** PlayerGameLog

### 10. TeamVsTeam
**Why:** Historical team matchup data
**Complexity:** Medium (similar to LeagueGameFinder)
**Use Cases:** Rivalry analysis, prediction models
**Complementary to:** TeamGameLog

---

## 🎨 Tier 3: Enhanced Features (Good Priority)

### 11. PlayerEstimatedMetrics
**Why:** NBA's estimated advanced metrics
**Complexity:** Medium (newer endpoint)
**Use Cases:** Advanced analytics, research
**Complementary to:** PlayerAdvancedStats

### 12. TeamLineups
**Why:** Lineup combination statistics
**Complexity:** High (complex data structure)
**Use Cases:** Coaching analysis, lineup optimization
**Complementary to:** TeamDashboards

### 13. PlayerDefense
**Why:** Defensive metrics and matchup data
**Complexity:** Medium (defensive-specific stats)
**Use Cases:** DPOY analysis, defensive ratings
**Complementary to:** PlayerAdvancedStats

### 14. TeamDefense
**Why:** Team defensive metrics
**Complexity:** Medium (similar to PlayerDefense)
**Use Cases:** Defensive schemes, team rankings
**Complementary to:** TeamAdvancedStats

### 15. DraftCombineStats
**Why:** NBA Draft combine measurements
**Complexity:** Low (standard stats)
**Use Cases:** Draft analysis, prospect evaluation
**Complementary to:** Player lookups

---

## 📝 Recommended Implementation Order

### Phase 1: Core Functionality (Endpoints 1-5)
**Estimated Time:** 4-6 hours
**Value:** Fills critical gaps in basic functionality

1. LeagueStandings (1 hour)
2. PlayerAwards (45 min)
3. LeagueGameLog (45 min)
4. PlayoffPicture (1.5 hours)
5. PlayerGameScoreLog (45 min)

**Impact:** Users can now:
- Check standings
- View player accolades
- Access all games
- Track playoff race
- Analyze per-game efficiency

### Phase 2: Advanced Analytics (Endpoints 6-10)
**Estimated Time:** 6-8 hours
**Value:** Enables serious analytics use cases

6. TeamAdvancedStats (1.5 hours)
7. PlayerAdvancedStats (1.5 hours)
8. PlayerVsPlayer (1.5 hours)
9. TeamVsTeam (1 hour)
10. PlayerEstimatedMetrics (1.5 hours)

**Impact:** Library becomes viable for:
- Advanced analytics
- Research projects
- Prediction models
- Fantasy sports

### Phase 3: Enhanced Features (Endpoints 11-15)
**Estimated Time:** 6-8 hours (if needed)
**Value:** Nice-to-have features

---

## 🎯 Quick Win Strategy

**Goal:** Reach 25 endpoints (18% coverage) quickly

**Approach:**
1. Generate 5 Tier 1 endpoints (4-6 hours)
2. Test and validate (1 hour)
3. Create examples (1 hour)
4. Update documentation (30 min)

**Total Time:** 6.5-8.5 hours
**Result:**
- 20 endpoints total (15 current + 5 new)
- Core functionality complete
- Library viable for most common use cases

---

## 📊 Coverage Analysis

### Current State
- **Endpoints:** 15/139 (10.8%)
- **Categories Covered:**
  - Player stats: Good (CareerStats, GameLog, Info, Dashboard, Splits)
  - Team stats: Good (GameLog, Info, Dashboard, Splits, YearByYear)
  - Game data: Good (BoxScore, PlayByPlay, Summary, Scoreboard)
  - League data: Moderate (LeagueLeaders, GameFinder)
  - Advanced metrics: None
  - Historical: Minimal
  - Playoffs: None

### After Tier 1 Implementation
- **Endpoints:** 20/139 (14.4%)
- **New Categories:**
  - Standings ✅
  - Awards ✅
  - League-wide games ✅
  - Playoff tracking ✅
  - Enhanced game metrics ✅

### After Tier 2 Implementation
- **Endpoints:** 25/139 (18%)
- **New Categories:**
  - Advanced team metrics ✅
  - Advanced player metrics ✅
  - Matchup analysis ✅
  - Estimated metrics ✅

---

## 🔍 Metadata Extraction Strategy

Since we need metadata for new endpoints, we have options:

### Option 1: Manual Metadata Creation
**Effort:** 15-20 min per endpoint
**Quality:** High (tailored to our needs)
**Process:**
1. Find endpoint in Python nba_api
2. Extract parameters from `__init__`
3. Get result sets from `expected_data`
4. Infer types from field names
5. Create JSON file

### Option 2: Automated Extraction Script
**Effort:** 2-3 hours to build, 5 min per endpoint after
**Quality:** High (consistent)
**Process:**
1. Write Python script to analyze nba_api
2. Extract all metadata automatically
3. Generate JSON files
4. Review and adjust

### Option 3: Community/Documentation
**Effort:** 10-15 min per endpoint
**Quality:** Variable
**Process:**
1. Use NBA API documentation
2. Test live API calls
3. Manually document structure
4. Create metadata

**Recommendation:** Option 1 for immediate 5 endpoints, Option 2 if generating 20+

---

## 💡 Implementation Example: LeagueStandings

### 1. Create Metadata File

`tools/generator/metadata/leaguestandings.json`:
```json
[
  {
    "name": "LeagueStandings",
    "endpoint": "leaguestandingsv3",
    "parameters": [
      {"name": "LeagueID", "type": "LeagueID", "required": false, "default": "00"},
      {"name": "Season", "type": "Season", "required": true},
      {"name": "SeasonType", "type": "SeasonType", "required": true}
    ],
    "result_sets": [
      {
        "name": "Standings",
        "fields": [
          "LeagueID", "SeasonID", "TeamID", "TeamCity", "TeamName",
          "Conference", "ConferenceRecord", "Division", "DivisionRecord",
          "WinPCT", "HomeLosses", "HomeWins", "RoadLosses", "RoadWins",
          "Wins", "Losses", "ConferenceLosses", "ConferenceWins",
          "DivisionLosses", "DivisionWins", "PlayoffRank"
        ]
      }
    ]
  }
]
```

### 2. Generate Endpoint

```bash
go run ./tools/generator -metadata tools/generator/metadata/leaguestandings.json
```

### 3. Result

Type-safe endpoint with proper types inferred:
- TeamID → int
- TeamCity, TeamName → string
- Wins, Losses → int
- WinPCT → float64
- Conference, Division → string
- All records → string (e.g., "15-5")

### 4. Test

```go
client := stats.NewDefaultClient()
req := endpoints.LeagueStandingsRequest{
    Season:     parameters.NewSeason(2023),
    SeasonType: parameters.SeasonTypeRegular,
}
resp, err := endpoints.GetLeagueStandings(ctx, client, req)
// Access with full type safety!
```

---

## 🎉 Success Metrics

### After Phase 1 (5 new endpoints)
- ✅ 20/139 endpoints (14.4%)
- ✅ Core functionality complete
- ✅ Standings, awards, league games available
- ✅ Playoff tracking enabled

### After Phase 2 (10 new endpoints)
- ✅ 25/139 endpoints (18%)
- ✅ Advanced analytics enabled
- ✅ Matchup analysis available
- ✅ Research-grade metrics

### Quality Maintained
- ✅ All new endpoints type-safe (thanks to type inference)
- ✅ Consistent code quality across all endpoints
- ✅ Full IDE support and compile-time checking
- ✅ Production-ready from day one

---

## 📅 Timeline

**Week 1:** Complete remaining 7 regenerations (1 hour)
**Week 2:** Generate Tier 1 endpoints (4-6 hours)
**Week 3:** Generate Tier 2 endpoints (6-8 hours)
**Week 4:** Testing, examples, documentation (3-4 hours)

**Total:** ~20 hours to reach 25 high-quality, type-safe endpoints

**Alternative (Fast Track):**
- Day 1: Regenerate 7 remaining (1 hour)
- Day 2-3: Generate 5 Tier 1 (6 hours)
- Day 4: Test and document (2 hours)

**Fast Track Total:** 9 hours to reach 20 endpoints with core functionality

---

**Next Step:** Choose between:
1. Generate LeagueStandings (immediate value, 1 hour)
2. Complete remaining 7 regenerations first (finish type inference rollout, 1 hour)
3. Both in sequence (2 hours total, maximum impact)
