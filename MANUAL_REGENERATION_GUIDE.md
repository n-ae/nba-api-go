# Manual Regeneration Guide - Remaining 7 Endpoints

**Purpose:** Complete type inference rollout by regenerating remaining endpoints
**Effort:** ~10 minutes per endpoint = 70 minutes total
**Pattern:** Proven with 3 successful regenerations

---

## Quick Reference: Type Inference Rules

Apply these rules when converting `interface{}` to proper types:

```go
// Percentages and ratios
_PCT, _PERCENTAGE → float64

// IDs
PLAYER_ID, TEAM_ID → int
GAME_ID, SEASON_ID, LEAGUE_ID → string
OFFICIAL_ID, VIDEO_* → string
Most other *_ID → string

// Names and text
_NAME, _TEXT, _ABBREVIATION, _CITY → string
_BROADCASTER, GAMECODE, JERSEY_NUM → string

// Dates and times
*DATE*, *TIME* → string

// Status and flags
*_STATUS, *_FLAG → string

// Numbers/Stats
PTS, REB, AST, STL, BLK, TOV, PF → int
FGM, FGA, FTM, FTA, OREB, DREB → int
MIN, PLUS_MINUS, *_PCT → float64
LEAD_CHANGES, TIMES_TIED, TURNOVERS → int
ATTENDANCE, WINS, LOSSES, POINTS → int
SEQUENCE, PERIOD, QUARTER → int

// Game outcomes
WL (Win/Loss) → string
MATCHUP → string
```

---

## Endpoint 1: BoxScoreSummaryV2

**File:** `pkg/stats/endpoints/boxscoresummaryv2.go`
**Metadata:** `tools/generator/metadata/boxscoresummaryv2.json`
**Result Sets:** 9 (complex endpoint)

### Step 1: Update GameSummary struct

```go
type BoxScoreSummaryV2GameSummary struct {
    GAME_DATE_EST                     string `json:"GAME_DATE_EST"`
    GAME_SEQUENCE                     int    `json:"GAME_SEQUENCE"`
    GAME_ID                           string `json:"GAME_ID"`
    GAME_STATUS_ID                    int    `json:"GAME_STATUS_ID"`
    GAME_STATUS_TEXT                  string `json:"GAME_STATUS_TEXT"`
    GAMECODE                          string `json:"GAMECODE"`
    HOME_TEAM_ID                      int    `json:"HOME_TEAM_ID"`
    VISITOR_TEAM_ID                   int    `json:"VISITOR_TEAM_ID"`
    SEASON                            string `json:"SEASON"`
    LIVE_PERIOD                       int    `json:"LIVE_PERIOD"`
    LIVE_PC_TIME                      string `json:"LIVE_PC_TIME"`
    NATL_TV_BROADCASTER_ABBREVIATION  string `json:"NATL_TV_BROADCASTER_ABBREVIATION"`
    LIVE_PERIOD_TIME_BCAST            string `json:"LIVE_PERIOD_TIME_BCAST"`
    WH_STATUS                         int    `json:"WH_STATUS"`
}
```

### Step 2: Update parsing (repeat for all 9 result sets)

```go
if len(rawResp.ResultSets) > 0 {
    response.GameSummary = make([]BoxScoreSummaryV2GameSummary, 0, len(rawResp.ResultSets[0].RowSet))
    for _, row := range rawResp.ResultSets[0].RowSet {
        if len(row) >= 14 {
            item := BoxScoreSummaryV2GameSummary{
                GAME_DATE_EST:                    toString(row[0]),
                GAME_SEQUENCE:                    toInt(row[1]),
                GAME_ID:                          toString(row[2]),
                GAME_STATUS_ID:                   toInt(row[3]),
                GAME_STATUS_TEXT:                 toString(row[4]),
                GAMECODE:                         toString(row[5]),
                HOME_TEAM_ID:                     toInt(row[6]),
                VISITOR_TEAM_ID:                  toInt(row[7]),
                SEASON:                           toString(row[8]),
                LIVE_PERIOD:                      toInt(row[9]),
                LIVE_PC_TIME:                     toString(row[10]),
                NATL_TV_BROADCASTER_ABBREVIATION: toString(row[11]),
                LIVE_PERIOD_TIME_BCAST:           toString(row[12]),
                WH_STATUS:                        toInt(row[13]),
            }
            response.GameSummary = append(response.GameSummary, item)
        }
    }
}
```

**Repeat this pattern for:**
- OtherStats (14 fields)
- Officials (4 fields)
- InactivePlayers (8 fields)
- GameInfo (3 fields)
- LineScore (28 fields)
- LastMeeting (13 fields)
- SeasonSeries (7 fields)
- AvailableVideo (2 fields)

---

## Endpoint 2: ShotChartDetail

**File:** `pkg/stats/endpoints/shotchartdetail.go`
**Metadata:** `tools/generator/metadata/shotchartdetail.json`

### Key Field Types

```go
type ShotChartDetailShot_Chart_Detail struct {
    GRID_TYPE          string  `json:"GRID_TYPE"`
    GAME_ID            string  `json:"GAME_ID"`
    GAME_EVENT_ID      int     `json:"GAME_EVENT_ID"`
    PLAYER_ID          int     `json:"PLAYER_ID"`
    PLAYER_NAME        string  `json:"PLAYER_NAME"`
    TEAM_ID            int     `json:"TEAM_ID"`
    TEAM_NAME          string  `json:"TEAM_NAME"`
    PERIOD             int     `json:"PERIOD"`
    MINUTES_REMAINING  int     `json:"MINUTES_REMAINING"`
    SECONDS_REMAINING  int     `json:"SECONDS_REMAINING"`
    EVENT_TYPE         string  `json:"EVENT_TYPE"`
    ACTION_TYPE        string  `json:"ACTION_TYPE"`
    SHOT_TYPE          string  `json:"SHOT_TYPE"`
    SHOT_ZONE_BASIC    string  `json:"SHOT_ZONE_BASIC"`
    SHOT_ZONE_AREA     string  `json:"SHOT_ZONE_AREA"`
    SHOT_ZONE_RANGE    string  `json:"SHOT_ZONE_RANGE"`
    SHOT_DISTANCE      float64 `json:"SHOT_DISTANCE"`
    LOC_X              float64 `json:"LOC_X"`
    LOC_Y              float64 `json:"LOC_Y"`
    SHOT_ATTEMPTED_FLAG int    `json:"SHOT_ATTEMPTED_FLAG"`
    SHOT_MADE_FLAG     int     `json:"SHOT_MADE_FLAG"`
    GAME_DATE          string  `json:"GAME_DATE"`
    HTM                string  `json:"HTM"`
    VTM                string  `json:"VTM"`
}
```

---

## Endpoint 3: TeamYearByYearStats

**File:** `pkg/stats/endpoints/teamyearbyyearstats.go`
**Metadata:** `tools/generator/metadata/teamyearbyyearstats.json`

### Key Field Types

```go
type TeamYearByYearStatsTeamStats struct {
    TEAM_ID           int     `json:"TEAM_ID"`
    TEAM_CITY         string  `json:"TEAM_CITY"`
    TEAM_NAME         string  `json:"TEAM_NAME"`
    YEAR              string  `json:"YEAR"`
    GP                int     `json:"GP"`
    WINS              int     `json:"WINS"`
    LOSSES            int     `json:"LOSSES"`
    WIN_PCT           float64 `json:"WIN_PCT"`
    CONF_RANK         int     `json:"CONF_RANK"`
    DIV_RANK          int     `json:"DIV_RANK"`
    PO_WINS           int     `json:"PO_WINS"`
    PO_LOSSES         int     `json:"PO_LOSSES"`
    CONF_COUNT        int     `json:"CONF_COUNT"`
    DIV_COUNT         int     `json:"DIV_COUNT"`
    NBA_FINALS_APPEARANCE string `json:"NBA_FINALS_APPEARANCE"`
    // ... continue for all fields
}
```

---

## Endpoint 4: PlayerDashboardByGeneralSplits

**File:** `pkg/stats/endpoints/playerdashboardbygeneralsplits.go`
**Metadata:** `tools/generator/metadata/playerdashboardbygeneralsplits.json`

### Common Dashboard Field Types

Most dashboard endpoints follow similar patterns:

```go
// Group/Split identifiers
GROUP_SET, GROUP_VALUE → string

// Basic info
PLAYER_ID, TEAM_ID → int
PLAYER_NAME, TEAM_ABBREVIATION → string

// Game counts
GP, W, L, MIN → int or float64 (MIN is float64)

// Shooting stats
FGM, FGA → int
FG_PCT, FG3_PCT, FT_PCT → float64

// Other stats
PTS, REB, AST, TOV, STL, BLK, PF → int or float64 (depends on if average)
PLUS_MINUS → float64
```

---

## Endpoint 5: TeamDashboardByGeneralSplits

**File:** `pkg/stats/endpoints/teamdashboardbygeneralsplits.go`
**Metadata:** `tools/generator/metadata/teamdashboardbygeneralsplits.json`

Similar to PlayerDashboard but with TEAM_ID instead of PLAYER_ID.

---

## Endpoint 6: PlayByPlayV2

**File:** `pkg/stats/endpoints/playbyplayv2.go`
**Metadata:** `tools/generator/metadata/playbyplayv2.json`

### Key Field Types

```go
type PlayByPlayV2PlayByPlay struct {
    GAME_ID                string `json:"GAME_ID"`
    EVENTNUM               int    `json:"EVENTNUM"`
    EVENTMSGTYPE           int    `json:"EVENTMSGTYPE"`
    EVENTMSGACTIONTYPE     int    `json:"EVENTMSGACTIONTYPE"`
    PERIOD                 int    `json:"PERIOD"`
    WCTIMESTRING           string `json:"WCTIMESTRING"`
    PCTIMESTRING           string `json:"PCTIMESTRING"`
    HOMEDESCRIPTION        string `json:"HOMEDESCRIPTION"`
    NEUTRALDESCRIPTION     string `json:"NEUTRALDESCRIPTION"`
    VISITORDESCRIPTION     string `json:"VISITORDESCRIPTION"`
    SCORE                  string `json:"SCORE"`
    SCOREMARGIN            string `json:"SCOREMARGIN"`
    PERSON1TYPE            int    `json:"PERSON1TYPE"`
    PLAYER1_ID             int    `json:"PLAYER1_ID"`
    PLAYER1_NAME           string `json:"PLAYER1_NAME"`
    PLAYER1_TEAM_ID        int    `json:"PLAYER1_TEAM_ID"`
    PLAYER1_TEAM_CITY      string `json:"PLAYER1_TEAM_CITY"`
    PLAYER1_TEAM_NICKNAME  string `json:"PLAYER1_TEAM_NICKNAME"`
    PLAYER1_TEAM_ABBREVIATION string `json:"PLAYER1_TEAM_ABBREVIATION"`
    // ... similar for PLAYER2 and PLAYER3
}
```

---

## Endpoint 7: TeamInfoCommon

**File:** `pkg/stats/endpoints/teaminfocommon.go`
**Metadata:** `tools/generator/metadata/teaminfocommon.json`

### Key Field Types

```go
type TeamInfoCommonTeamInfoCommon struct {
    TEAM_ID           int    `json:"TEAM_ID"`
    SEASON_YEAR       string `json:"SEASON_YEAR"`
    TEAM_CITY         string `json:"TEAM_CITY"`
    TEAM_NAME         string `json:"TEAM_NAME"`
    TEAM_ABBREVIATION string `json:"TEAM_ABBREVIATION"`
    TEAM_CONFERENCE   string `json:"TEAM_CONFERENCE"`
    TEAM_DIVISION     string `json:"TEAM_DIVISION"`
    TEAM_CODE         string `json:"TEAM_CODE"`
    W                 int    `json:"W"`
    L                 int    `json:"L"`
    PCT               float64 `json:"PCT"`
    CONF_RANK         int    `json:"CONF_RANK"`
    DIV_RANK          int    `json:"DIV_RANK"`
    // ... continue for all fields
}
```

---

## General Regeneration Steps

For each endpoint:

### 1. Read the metadata file
```bash
cat tools/generator/metadata/[endpoint].json
```

### 2. Identify field types
- Use the type inference rules above
- Check similar fields in already-regenerated endpoints
- When in doubt, use `string` (safest)

### 3. Update struct definition
- Replace `interface{}` with proper type
- Add ` `json:"FIELD_NAME"` ` tag to each field
- Keep field names in SCREAMING_SNAKE_CASE

### 4. Update parsing logic
- Change from `make([]Type, len(rows))` to `make([]Type, 0, len(rows))`
- Change from `response.Data[i] = ...` to `item := ... ; append()`
- Apply correct conversion function:
  - `toInt(row[N])` for int fields
  - `toFloat(row[N])` for float64 fields
  - `toString(row[N])` for string fields

### 5. Verify the pattern
- Check field count matches metadata
- Ensure array indices are correct
- Verify conversion functions match types

---

## Verification Commands

After each regeneration:

```bash
# Check the specific file
grep 'interface{}' pkg/stats/endpoints/[endpoint].go

# Should only find conversion helpers, not struct fields

# Try to compile
go build ./pkg/stats/endpoints/[endpoint].go

# Run tests
go test ./pkg/stats/endpoints -run [Endpoint]
```

---

## Common Mistakes to Avoid

1. **Wrong type for IDs**
   - PLAYER_ID, TEAM_ID → int (NOT string)
   - GAME_ID, SEASON_ID → string (NOT int)

2. **Missing JSON tags**
   - Every field needs ` `json:"FIELD_NAME"` `

3. **Forgetting to change parsing**
   - Must use `toInt()`, `toFloat()`, `toString()`
   - Can't just assign `row[N]` directly anymore

4. **Array index off-by-one**
   - If metadata has 14 fields, use `len(row) >= 14`
   - Use indices 0-13 (not 1-14)

5. **Wrong make() pattern**
   - OLD: `make([]Type, len(rows))`
   - NEW: `make([]Type, 0, len(rows))`
   - Then use `append()` not index assignment

---

## Progress Tracking

Check off as you complete each endpoint:

- [ ] BoxScoreSummaryV2 (9 result sets - most complex!)
- [ ] ShotChartDetail (1 result set, ~24 fields)
- [ ] TeamYearByYearStats (1 result set)
- [ ] PlayerDashboardByGeneralSplits (multiple result sets)
- [ ] TeamDashboardByGeneralSplits (multiple result sets)
- [ ] PlayByPlayV2 (1 large result set, ~35 fields)
- [ ] TeamInfoCommon (1 result set)

---

## Estimated Time Per Endpoint

- **Simple** (1 result set, <15 fields): 5-7 minutes
- **Medium** (1-2 result sets, 15-25 fields): 8-12 minutes
- **Complex** (3+ result sets or 25+ fields): 15-20 minutes

**Total estimated time:** 60-80 minutes

---

## Final Verification

When all 7 are complete:

```bash
# Verify NO interface{} remains in structs
grep -n 'interface{}' pkg/stats/endpoints/*.go | grep -v types.go | grep -v _test.go | grep -v _improved.go

# Should only show conversion helper functions, not struct fields

# Compile all endpoints
go build ./pkg/stats/endpoints

# Run all tests
go test ./pkg/stats/endpoints

# Check git diff to see changes
git diff pkg/stats/endpoints/
```

---

## Success Criteria

✅ All 7 endpoints regenerated
✅ No `interface{}` in struct fields (only in types.go helpers)
✅ All fields have JSON tags
✅ All parsing uses type conversion functions
✅ Code compiles without errors
✅ Tests pass

---

**Ready to proceed!** Start with the simplest endpoint (TeamInfoCommon or ShotChartDetail) to build confidence, then tackle the complex ones (BoxScoreSummaryV2).
