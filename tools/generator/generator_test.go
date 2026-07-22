package main

import (
	"encoding/json"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// TestGenerateFromMetadata_ProducesValidGo is a generator/template
// regression test, not a fidelity check against already-committed
// endpoint code: regenerating from tools/generator/metadata/*.json today
// produces output that differs from pkg/stats/endpoints/*.go in real,
// substantive ways (field types, not just formatting) - see
// docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md's
// v2.0.0 plan item for explicit per-field type metadata and a full
// regeneration of all 141 endpoints, which is the only safe way to
// reconcile that drift (it requires verifying corrected types against
// live NBA.com responses, not blindly trusting either side). What this
// test does check, for every committed metadata file: the generator
// doesn't crash, and its output is syntactically valid Go - so a broken
// template or a change to inferGoType that emits malformed code fails
// CI immediately instead of silently producing something nobody would
// notice was broken until they tried to build it.
func TestGenerateFromMetadata_ProducesValidGo(t *testing.T) {
	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}
	if len(metadataFiles) == 0 {
		t.Fatal("no metadata files found under metadata/ - this test would trivially pass without exercising anything")
	}

	for _, mf := range metadataFiles {
		t.Run(filepath.Base(mf), func(t *testing.T) {
			outDir := t.TempDir()
			serverOutDir := t.TempDir()
			g := NewGenerator(outDir, serverOutDir)
			if err := g.GenerateFromMetadata(mf, false); err != nil {
				t.Fatalf("GenerateFromMetadata(%s) failed: %v", mf, err)
			}

			// SDK output (outDir) is empty for a metadata file whose
			// entries are all handler_only (see
			// metadata/handwritten_handlers.json) - by design,
			// generateEndpoint is never called for those. Handler output
			// (serverOutDir) always has at least one file regardless,
			// since every entry gets a generated handler. Checking the
			// combined total, not outDir alone, keeps this test's "did it
			// silently no-op" guarantee for every metadata file without
			// special-casing handler_only ones.
			var paths []string
			for _, dir := range []string{outDir, serverOutDir} {
				entries, err := os.ReadDir(dir)
				if err != nil {
					t.Fatalf("failed to read output dir %s: %v", dir, err)
				}
				for _, entry := range entries {
					paths = append(paths, filepath.Join(dir, entry.Name()))
				}
			}
			if len(paths) == 0 {
				t.Fatalf("%s produced no output files in either output directory", mf)
			}

			fset := token.NewFileSet()
			for _, path := range paths {
				if _, err := parser.ParseFile(fset, path, nil, parser.AllErrors); err != nil {
					t.Errorf("%s: generated file %s is not valid Go: %v", mf, path, err)
				}
			}
		})
	}
}

// TestGenerateSingleEndpointLoadsMetadata prevents -endpoint from falling
// back to the old empty EndpointMetadata stub. A successful single-endpoint
// generation must include the parameters and result sets supplied by the
// matching metadata entry.
func TestGenerateSingleEndpointLoadsMetadata(t *testing.T) {
	const endpointName = "LeagueGameFinder"

	metadata, err := findEndpointMetadata(defaultMetadataDir(), endpointName)
	if err != nil {
		t.Fatalf("findEndpointMetadata(%q, %q): %v", defaultMetadataDir(), endpointName, err)
	}
	if len(metadata.Parameters) == 0 || len(metadata.ResultSets) == 0 {
		t.Fatalf("test metadata for %s is incomplete: got %d parameters and %d result sets", endpointName, len(metadata.Parameters), len(metadata.ResultSets))
	}

	outDir := t.TempDir()
	g := NewGenerator(outDir, t.TempDir())
	if err := g.GenerateSingleEndpoint(endpointName, defaultMetadataDir(), false); err != nil {
		t.Fatalf("GenerateSingleEndpoint(%q): %v", endpointName, err)
	}

	generatedFile := filepath.Join(outDir, strings.ToLower(endpointName)+".go")
	generated, err := os.ReadFile(generatedFile)
	if err != nil {
		t.Fatalf("failed to read generated endpoint: %v", err)
	}
	for _, parameter := range metadata.Parameters {
		if !strings.Contains(string(generated), "\t"+parameter.Name+" ") {
			t.Errorf("generated endpoint is missing parameter %q", parameter.Name)
		}
	}
	for _, resultSet := range metadata.ResultSets {
		if !strings.Contains(string(generated), "\t"+resultSet.Name+" []") {
			t.Errorf("generated endpoint is missing result set %q", resultSet.Name)
		}
	}
}

// TestGenerateSingleEndpointRejectsUnknownMetadata ensures a typo or a
// metadata gap fails rather than writing a compilable-but-empty endpoint.
func TestGenerateSingleEndpointRejectsUnknownMetadata(t *testing.T) {
	outDir := t.TempDir()
	g := NewGenerator(outDir, t.TempDir())
	if err := g.GenerateSingleEndpoint("NotARealEndpoint", defaultMetadataDir(), false); err == nil {
		t.Fatal("GenerateSingleEndpoint() succeeded for an endpoint without metadata")
	}

	entries, err := os.ReadDir(outDir)
	if err != nil {
		t.Fatalf("failed to read output directory: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("GenerateSingleEndpoint() wrote files for an endpoint without metadata: %v", entries)
	}
}

// TestProcessHandlerMetadata directly exercises the field resolution
// templates/handler.tmpl and templates/dispatch.tmpl depend on, rather
// than only observing it indirectly through generated output (as
// TestGenerateFromMetadata_ProducesValidGo does). Regression coverage for
// the defaulting rules: NameLower, EffectiveSDKFunction's "Get"+Name
// fallback, EffectiveResponseWrapped's true-when-nil default, each
// parameter's HandlerGoType resolution, EffectivePointer's
// !Required-when-nil fallback, and NeedsParametersImport's two triggers
// (a recognized special name, or any non-string HandlerGoType).
func TestProcessHandlerMetadata(t *testing.T) {
	g := NewGenerator(t.TempDir(), t.TempDir())

	original := EndpointMetadata{
		Name: "FooBar",
		Parameters: []ParameterMetadata{
			{Name: "Season", Type: "Season", Required: true},
			{Name: "PlayerID", Type: "string", Required: true},
			{Name: "PerMode", Type: "PerMode", Required: false},
		},
	}

	got := g.processHandlerMetadata(original)

	if got.NameLower != "foobar" {
		t.Errorf("NameLower = %q, want %q", got.NameLower, "foobar")
	}
	if got.EffectiveSDKFunction != "GetFooBar" {
		t.Errorf("EffectiveSDKFunction = %q, want %q (empty SDKFunction should default to \"Get\"+Name)", got.EffectiveSDKFunction, "GetFooBar")
	}
	if !got.EffectiveResponseWrapped {
		t.Error("EffectiveResponseWrapped = false, want true (nil ResponseWrapped should default true)")
	}
	if !got.NeedsParametersImport {
		t.Error("NeedsParametersImport = false, want true (Season/PerMode are recognized special names)")
	}

	wantEffectivePointer := map[string]bool{
		"Season":   false, // Required: true -> !Required
		"PlayerID": false,
		"PerMode":  true, // Required: false -> !Required
	}
	if len(got.Parameters) != len(original.Parameters) {
		t.Fatalf("got %d parameters, want %d", len(got.Parameters), len(original.Parameters))
	}
	for _, p := range got.Parameters {
		// toHandlerParamType is the oracle, not a repeated literal: this
		// checks processHandlerMetadata actually calls it per parameter,
		// not a hardcoded copy of its mapping table that could drift from
		// the real one independently.
		if want := toHandlerParamType(p.Type); p.HandlerGoType != want {
			t.Errorf("%s.HandlerGoType = %q, want %q (toHandlerParamType(%q))", p.Name, p.HandlerGoType, want, p.Type)
		}
		if want := wantEffectivePointer[p.Name]; p.EffectivePointer != want {
			t.Errorf("%s.EffectivePointer = %v, want %v", p.Name, p.EffectivePointer, want)
		}
	}

	// Deep-copy invariant (see processHandlerMetadata's doc comment):
	// the input's Parameters slice must not be mutated. HandlerGoType and
	// EffectivePointer are only ever set by this function, so if the
	// original still shows their zero values, the copy was real.
	for _, p := range original.Parameters {
		if p.HandlerGoType != "" {
			t.Errorf("original metadata's %s.HandlerGoType was mutated to %q - processHandlerMetadata must not mutate its input", p.Name, p.HandlerGoType)
		}
	}
}

// TestProcessHandlerMetadataExplicitOverrides covers the three fields a
// caller can set explicitly to override processHandlerMetadata's
// defaults: SDKFunction, ResponseWrapped, and a parameter's Pointer -
// needed by the hand-written-SDK (HandlerOnly) endpoints whose Request
// structs don't uniformly follow the SDK generator's own conventions.
func TestProcessHandlerMetadataExplicitOverrides(t *testing.T) {
	g := NewGenerator(t.TempDir(), t.TempDir())
	explicitFalse := false

	metadata := EndpointMetadata{
		Name:            "Baz",
		SDKFunction:     "CustomGetBaz",
		ResponseWrapped: &explicitFalse,
		Parameters: []ParameterMetadata{
			// Required: false would normally default EffectivePointer to
			// true; an explicit Pointer: false must win instead.
			{Name: "LeagueID", Type: "LeagueID", Required: false, Pointer: &explicitFalse},
		},
	}

	got := g.processHandlerMetadata(metadata)

	if got.EffectiveSDKFunction != "CustomGetBaz" {
		t.Errorf("EffectiveSDKFunction = %q, want explicit override %q", got.EffectiveSDKFunction, "CustomGetBaz")
	}
	if got.EffectiveResponseWrapped {
		t.Error("EffectiveResponseWrapped = true, want explicit override false")
	}
	if len(got.Parameters) != 1 {
		t.Fatalf("got %d parameters, want 1", len(got.Parameters))
	}
	if got.Parameters[0].EffectivePointer {
		t.Error("LeagueID.EffectivePointer = true, want explicit override false (ignoring the Required=false default)")
	}
}

// TestGenerateHandler directly exercises generateHandler - the actual
// write path templates/handler.tmpl renders through - rather than only
// observing it indirectly through GenerateFromMetadata/
// GenerateSingleEndpoint (which also call it, but bundle it with SDK
// generation). Checks the file lands at the expected path, parses as
// valid Go, and contains the specific control flow a required string
// parameter should produce.
func TestGenerateHandler(t *testing.T) {
	serverOutDir := t.TempDir()
	g := NewGenerator(t.TempDir(), serverOutDir)

	metadata := g.processHandlerMetadata(EndpointMetadata{
		Name: "FooBar",
		Parameters: []ParameterMetadata{
			{Name: "PlayerID", Type: "string", Required: true},
		},
	})

	if err := g.generateHandler(metadata, false); err != nil {
		t.Fatalf("generateHandler() error = %v", err)
	}

	path := filepath.Join(serverOutDir, "generated_foobar.go")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected generated handler at %s: %v", path, err)
	}

	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, path, data, parser.AllErrors); err != nil {
		t.Errorf("generated handler is not valid Go: %v\n%s", err, data)
	}

	content := string(data)
	if !strings.Contains(content, "func (h *StatsHandler) handleFooBar(") {
		t.Errorf("generated handler missing expected function signature, got:\n%s", content)
	}
	if !strings.Contains(content, `writeError(w, http.StatusBadRequest, "missing_parameter", "PlayerID is required")`) {
		t.Errorf("generated handler missing required-parameter validation for PlayerID, got:\n%s", content)
	}
}

// TestGenerateDispatchTable directly exercises GenerateDispatchTable
// against a small fixture metadata directory, rather than only relying on
// the real, 141-endpoint tools/generator/metadata/ being correct by
// accident. Covers the documented dedup behavior for a Name appearing in
// more than one metadata file (several real endpoints do this - see the
// function's own doc comment): the first file encountered (alphabetical,
// per filepath.Glob's sorted order) must win, both in the dispatch
// table's map entries and in which file's parameters got rendered into
// the deduped handler.
func TestGenerateDispatchTable(t *testing.T) {
	metadataDir := t.TempDir()
	serverOutDir := t.TempDir()

	writeMetadata := func(filename string, endpoints []EndpointMetadata) {
		t.Helper()
		data, err := json.MarshalIndent(endpoints, "", "  ")
		if err != nil {
			t.Fatalf("failed to marshal fixture metadata: %v", err)
		}
		if err := os.WriteFile(filepath.Join(metadataDir, filename), data, 0o600); err != nil {
			t.Fatalf("failed to write fixture metadata: %v", err)
		}
	}

	writeMetadata("a.json", []EndpointMetadata{
		{Name: "AlphaEndpoint", Parameters: []ParameterMetadata{{Name: "ID", Type: "string", Required: true}}},
		{Name: "DupeEndpoint", Parameters: []ParameterMetadata{{Name: "First", Type: "string", Required: true}}},
	})
	writeMetadata("b.json", []EndpointMetadata{
		{Name: "BetaEndpoint", Parameters: []ParameterMetadata{{Name: "ID", Type: "string", Required: true}}},
		// Same Name as a.json's DupeEndpoint but a different parameter -
		// GenerateDispatchTable must keep whichever file's entry it saw
		// first (a.json, alphabetically earlier), not silently produce a
		// duplicate map key or overwrite with this one instead.
		{Name: "DupeEndpoint", Parameters: []ParameterMetadata{{Name: "Second", Type: "string", Required: true}}},
	})

	g := NewGenerator(t.TempDir(), serverOutDir)
	if err := g.GenerateDispatchTable(metadataDir, false); err != nil {
		t.Fatalf("GenerateDispatchTable() error = %v", err)
	}

	dispatchPath := filepath.Join(serverOutDir, "generated_dispatch.go")
	dispatchData, err := os.ReadFile(dispatchPath)
	if err != nil {
		t.Fatalf("expected dispatch table at %s: %v", dispatchPath, err)
	}

	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, dispatchPath, dispatchData, parser.AllErrors); err != nil {
		t.Errorf("generated dispatch table is not valid Go: %v\n%s", err, dispatchData)
	}

	dispatchContent := string(dispatchData)
	for _, want := range []string{
		`"alphaendpoint": (*StatsHandler).handleAlphaEndpoint,`,
		`"betaendpoint": (*StatsHandler).handleBetaEndpoint,`,
		`"dupeendpoint": (*StatsHandler).handleDupeEndpoint,`,
	} {
		if !strings.Contains(dispatchContent, want) {
			t.Errorf("dispatch table missing expected entry %q, got:\n%s", want, dispatchContent)
		}
	}
	if strings.Count(dispatchContent, `"dupeendpoint":`) != 1 {
		t.Errorf("dispatch table should have exactly one %q entry despite the name appearing in two metadata files, got:\n%s", "dupeendpoint", dispatchContent)
	}

	// The deduped handler file must reflect a.json's entry (parameter
	// "First"), not b.json's ("Second").
	handlerPath := filepath.Join(serverOutDir, "generated_dupeendpoint.go")
	handlerData, err := os.ReadFile(handlerPath)
	if err != nil {
		t.Fatalf("expected handler file at %s: %v", handlerPath, err)
	}
	handlerContent := string(handlerData)
	if !strings.Contains(handlerContent, "vFirst") {
		t.Errorf("generated_dupeendpoint.go should reflect a.json's entry (parameter First, first-encountered), got:\n%s", handlerContent)
	}
	if strings.Contains(handlerContent, "vSecond") {
		t.Errorf("generated_dupeendpoint.go should NOT reflect b.json's entry (parameter Second) - first-encountered should win, got:\n%s", handlerContent)
	}
}

// TestGenerateDispatchTableRequiresMetadataFiles guards the explicit
// error GenerateDispatchTable returns for an empty metadata directory,
// rather than silently writing an empty (and useless) dispatch table.
func TestGenerateDispatchTableRequiresMetadataFiles(t *testing.T) {
	g := NewGenerator(t.TempDir(), t.TempDir())
	if err := g.GenerateDispatchTable(t.TempDir(), false); err == nil {
		t.Fatal("GenerateDispatchTable() succeeded against an empty metadata directory")
	}
}

// TestGoFieldName documents goFieldName's camelCase-to-exported-Go-identifier
// conversion using real field names from committed metadata. Most metadata
// field names are already valid, exported Go identifiers
// (SCREAMING_SNAKE_CASE, e.g. "GAME_ID") and goFieldName is a no-op for
// them; the interesting cases are the NBA Live-Data-style endpoints
// (PlayByPlayV3, ScoreboardV3, LeagueStandings) whose field names are
// camelCase and start with a lowercase letter - an unexported struct field
// generated from those would be invisible to encoding/json and
// inaccessible to any external caller, a real bug, not a style choice.
// Every case below is checked against the actual field name already
// committed (and presumably hand-fixed at some point) in the corresponding
// endpoint file, so a regression here would be a visible mismatch against
// real production code, not just an assertion this test invented.
func TestGoFieldName(t *testing.T) {
	tests := []struct {
		field string
		want  string
		note  string
	}{
		// --- Already-exported identifiers pass through unchanged ---
		{field: "GAME_ID", want: "GAME_ID"},
		{field: "PLAYER_NAME", want: "PLAYER_NAME"},

		// --- Plain camelCase words: capitalize the first letter only ---
		{field: "actionNumber", want: "ActionNumber"},
		{field: "teamTricode", want: "TeamTricode"},
		{field: "isFieldGoal", want: "IsFieldGoal", note: "three words, none an initialism"},
		{field: "xLegacy", want: "XLegacy", note: "single-letter first word"},

		// --- "Id" is a recognized initialism -> fully uppercased, matching
		//     playbyplayv3.go/scoreboardv3.go's committed field names ---
		{field: "gameId", want: "GameID"},
		{field: "teamId", want: "TeamID"},
		{field: "personId", want: "PersonID"},
		{field: "assistPersonId", want: "AssistPersonID", note: "three words, only the last is an initialism"},
		{field: "homeTeamId", want: "HomeTeamID"},

		// --- A field name that IS an initialism in its entirety ---
		{field: "uuid", want: "UUID"},

		// --- Not recognized initialisms by the general rule (outside any
		//     endpoint scope, "vl"/"surl" are NOT in fieldname_overrides.json,
		//     so they fall through to plain capitalization - see
		//     TestGoFieldNameOverridesApplyOnlyWithinTheirEndpoint for the
		//     VideoEvents-scoped case where these DO become "VL"/"SURL") ---
		{field: "vl", want: "Vl"},
		{field: "surl", want: "Surl"},

		// --- "str" prefix isn't an initialism, matching
		//     leaguestandings.go's committed "StrLongHomeStreak" ---
		{field: "strLongHomeStreak", want: "StrLongHomeStreak"},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			got := goFieldName("", "", tt.field)
			if got != tt.want {
				t.Errorf("goFieldName(%q) = %q, want %q (%s)", tt.field, got, tt.want, tt.note)
			}
		})
	}
}

// TestGoFieldNameOverridesApplyOnlyWithinTheirEndpoint proves
// fieldname_overrides.json's VideoEvents.Video entries produce VL/VT/GC/
// SURL/DURL/VURL/PURL (matching the already-committed, hand-fixed field
// names) only for that exact (endpoint, result set, field) scope - the
// same short field names anywhere else fall through to the general rule
// (plain capitalization), matching TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint's
// pattern for fieldtype_overrides.json. If this test passed with the
// override applying globally, a future endpoint that happens to have a
// field literally named "vl" for something else entirely would silently
// get "VL" instead of the field name it actually declared.
func TestGoFieldNameOverridesApplyOnlyWithinTheirEndpoint(t *testing.T) {
	cases := map[string]string{
		"vl": "VL", "vt": "VT", "gc": "GC",
		"surl": "SURL", "durl": "DURL", "vurl": "VURL", "purl": "PURL",
	}
	for field, want := range cases {
		t.Run(field, func(t *testing.T) {
			plainCapitalized := strings.ToUpper(field[:1]) + field[1:]
			if got := goFieldName("", "", field); got != plainCapitalized {
				t.Fatalf("goFieldName(%q) outside any endpoint scope = %q, want the general rule's plain capitalization %q - has the general rule changed?", field, got, plainCapitalized)
			}
			if got := goFieldName("VideoEvents", "Video", field); got != want {
				t.Errorf("goFieldName(%q, %q, %q) = %q, want %q (fieldname_overrides.json entry)", "VideoEvents", "Video", field, got, want)
			}
			if got := goFieldName("VideoEvents", "SomeOtherResultSetNotInOverrides", field); got == want {
				t.Errorf("goFieldName(%q, %q, %q) = %q - override leaked to an unlisted result set", "VideoEvents", "SomeOtherResultSetNotInOverrides", field, got)
			}
		})
	}
}

// TestGoFieldNameOverridesReferenceRealMetadata ensures every
// (endpoint, result set, field) entry in fieldname_overrides.json
// actually exists in committed metadata, mirroring
// TestFieldTypeOverridesReferenceRealMetadata's protection against a
// typo'd or stale entry silently doing nothing.
func TestGoFieldNameOverridesReferenceRealMetadata(t *testing.T) {
	overrides, err := loadFieldNameOverrides()
	if err != nil {
		t.Fatalf("loadFieldNameOverrides() failed: %v", err)
	}
	if len(overrides) == 0 {
		t.Fatal("fieldname_overrides.json is empty - this test would trivially pass without exercising anything")
	}

	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}

	realFields := map[string]map[string]map[string]bool{}
	for _, mf := range metadataFiles {
		data, err := os.ReadFile(mf)
		if err != nil {
			t.Fatalf("failed to read %s: %v", mf, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", mf, err)
		}
		for _, ep := range endpoints {
			for _, rs := range ep.ResultSets {
				if realFields[ep.Name] == nil {
					realFields[ep.Name] = map[string]map[string]bool{}
				}
				if realFields[ep.Name][rs.Name] == nil {
					realFields[ep.Name][rs.Name] = map[string]bool{}
				}
				for _, field := range rs.Fields {
					realFields[ep.Name][rs.Name][field] = true
				}
			}
		}
	}

	for endpointName, resultSets := range overrides {
		for resultSetName, fields := range resultSets {
			for field := range fields {
				if !realFields[endpointName][resultSetName][field] {
					t.Errorf("fieldname_overrides.json has %s.%s.%s, but no committed metadata file has a %q endpoint with a %q result set containing field %q",
						endpointName, resultSetName, field, endpointName, resultSetName, field)
				}
			}
		}
	}
}

// TestInferGoType documents inferGoType's current rule-by-rule heuristic
// behavior using real NBA.com field names pulled from committed generated
// code. Several cases below are marked knownWrong: true - they reproduce
// the exact data-corruption bugs cataloged in
// docs/archive/MAINTAINABLE_ARCHITECT_V4_ASSESSMENT_2026-07-19_2363f46.md (display
// names typed float64, textual range buckets typed int, decimal
// percentages typed string and truncated). These assertions exist to make
// any future change to the heuristic - the v2.0.0 explicit-type-metadata
// rework the assessment calls for - visible and intentional in a diff, not
// silent. Do not "fix" a knownWrong case by changing its want value
// without also fixing inferGoType and regenerating the affected endpoints;
// until then, this test's job is to pin current behavior, not endorse it.
//
// Two precedence quirks worth knowing when editing inferGoType or this
// test (both verified empirically, not just read from source):
//   - "STATUS_ID"/"GAME_STATUS_ID" never reach the "status substring +
//     _id suffix -> int" rule near the end of the function: the earlier,
//     unconditional "_id suffix" rule already returns "string" for any
//     _id field without "player"/"team" in it. That later rule is dead
//     code today.
//   - "CONF_RANK"/"DIV_RANK" don't reach the dedicated "rank substring ->
//     int" rule either: "_rank" is also in the stat-abbreviation list
//     checked earlier, so they fall through that loop's default
//     ("float64" for most stats) instead. A bare "RANK" (no underscore
//     before it) does reach the dedicated rule and gets "int". Similarly,
//     "SEASON_SEQUENCE" matches the season rule ("contains season, not
//     id") before it can reach the sequence/period/range rule.
func TestInferGoType(t *testing.T) {
	tests := []struct {
		field      string
		want       string
		knownWrong bool
		note       string
	}{
		// --- _PCT/_PERCENTAGE suffix -> float64 (correct) ---
		{field: "FG_PCT", want: "float64"},
		{field: "FT_PCT", want: "float64"},
		{field: "WIN_PERCENTAGE", want: "float64"},

		// --- Decimal percentage fields that DON'T end in _PCT/_PERCENTAGE
		//     fall through to later rules and land on string, losing
		//     decimal precision when toString later formats with %.0f. ---
		{field: "FG_PCT_RA", want: "string", knownWrong: true,
			note: "real committed type in leaguedashplayershotlocations.go; 0.357 becomes \"0\""},
		{field: "FG_PCT_IN_PAINT", want: "string", knownWrong: true,
			note: "same file, same bug class as FG_PCT_RA"},

		// --- _ID suffix: player/team -> int, else -> string (as designed) ---
		{field: "PLAYER_ID", want: "int"},
		{field: "TEAM_ID", want: "int"},
		{field: "GAME_ID", want: "string"},
		{field: "SEASON_ID", want: "string"},
		{field: "LEAGUE_ID", want: "string"},
		{field: "STATUS_ID", want: "string",
			note: "the later status+_id rule is dead code - see doc comment above"},
		{field: "GAME_STATUS_ID", want: "string"},

		// --- Date fields -> string (correct) ---
		{field: "GAME_DATE", want: "string"},
		{field: "GAME_DATE_EST", want: "string"},

		// --- _NAME/_TEXT/etc. suffix -> string, correct for genuine name
		//     fields, but the rule only fires on an exact suffix/substring
		//     match; PLAYER_NAME_LAST_FIRST doesn't end in one of those
		//     suffixes (it ends in "_LAST_FIRST"), so it falls through
		//     instead of matching here. ---
		{field: "TEAM_NAME", want: "string"},

		// --- Fields that read as names but don't match the name rule's
		//     suffix/substring list fall through to the stat-abbreviation
		//     loop and match "ast" hiding inside "last". ---
		{field: "PLAYER_NAME_LAST_FIRST", want: "float64", knownWrong: true,
			note: "real committed type in commonallplayers.go family; a name string becomes 0 via toFloat"},
		{field: "DISPLAY_FIRST_LAST", want: "float64", knownWrong: true,
			note: "real committed type in commonallplayers.go; \"Nikola Jokić\" becomes 0"},
		{field: "DISPLAY_LAST_COMMA_FIRST", want: "float64", knownWrong: true,
			note: "same bug class as DISPLAY_FIRST_LAST"},
		{field: "DISPLAY_FI_LAST", want: "float64", knownWrong: true,
			note: "same bug class as DISPLAY_FIRST_LAST"},

		// --- WL -> string (correct) ---
		{field: "WL", want: "string"},

		// --- Season fields without "id" -> string (correct); this rule's
		//     reach extends past plain "SEASON" to any field containing it,
		//     including ones that arguably belong to a different rule -
		//     see SEASON_SEQUENCE below. ---
		{field: "SEASON", want: "string"},
		{field: "SEASON_YEAR", want: "string"},
		{field: "SEASON_SEQUENCE", want: "string",
			note: "season rule fires before the sequence/period/range rule ever sees it"},

		// --- Stat abbreviations -> float64 by default, with sub-rules ---
		{field: "PTS", want: "float64"},
		{field: "REB", want: "float64"},
		{field: "MIN", want: "float64", note: "\"min\" substring, not \"game\" -> float64 sub-rule"},
		{field: "FGM", want: "int", note: "suffix \"m\" -> int sub-rule"},
		{field: "FGA", want: "int", note: "suffix \"a\" -> int sub-rule"},
		{field: "GP", want: "int", note: "exact \"gp\" -> int sub-rule"},
		{field: "GS", want: "int", note: "exact \"gs\" -> int sub-rule"},
		{field: "CONF_RANK", want: "float64",
			note: "\"_rank\" is itself in the stat-abbreviation list, so this never reaches the dedicated rank rule below"},
		{field: "DIV_RANK", want: "float64"},

		// --- Textual range buckets contain "range" and are caught by the
		//     sequence/period/range rule, typed int even though the actual
		//     values are text buckets like "24-22" or "Very Tight". ---
		{field: "SHOT_CLOCK_RANGE", want: "int", knownWrong: true,
			note: "real committed type in playerdashptshots.go; \"24-22\" becomes 0 via toInt"},
		{field: "CLOSE_DEF_DIST_RANGE", want: "int", knownWrong: true,
			note: "same bug class as SHOT_CLOCK_RANGE"},
		{field: "DRIBBLE_RANGE", want: "int", knownWrong: true,
			note: "same bug class as SHOT_CLOCK_RANGE"},

		// --- Age -> int (correct) ---
		{field: "PLAYER_AGE", want: "int"},
		{field: "AGE", want: "int"},

		// --- Dedicated rank rule -> int, only reached when "_rank" (with
		//     the underscore) isn't already caught by the stat-abbreviation
		//     loop above. ---
		{field: "RANK", want: "int"},

		// --- Sequence/period (non-range, non-season) -> int (correct) ---
		{field: "GAME_SEQUENCE", want: "int"},
		{field: "PERIOD", want: "int"},

		// --- Default fallback -> string ---
		{field: "SOME_UNRECOGNIZED_FIELD", want: "string"},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			got := inferGoType(tt.field)
			if got != tt.want {
				t.Errorf("inferGoType(%q) = %q, want %q (%s)", tt.field, got, tt.want, tt.note)
			}
		})
	}
}

// TestFieldTypesOverridesKnownWrongInference verifies that fieldtypes.json
// - the explicit, hand-reviewed dictionary resolveFieldGoType consults
// before ever falling back to inferGoType - actually corrects every
// knownWrong case
// documented in TestInferGoType above (plus other confirmed instances of
// the same bug classes found across the committed metadata: PCT_<STAT>
// "share of team total" fields mistyped by the FGM/FGA suffix sub-rule,
// FG_PCT_MID_RANGE mistyped by the _RANGE substring rule despite being a
// percentage rather than a text bucket, and GENERALMANAGER mistyped int
// because "generalmanager" contains the substring "age"). This is the test
// that proves the override mechanism actually fixes the corruption bugs
// inferGoType alone does not.
func TestFieldTypesOverridesKnownWrongInference(t *testing.T) {
	corrected := map[string]string{
		"CLOSE_DEF_DIST_RANGE":     "string",
		"DRIBBLE_RANGE":            "string",
		"SHOT_CLOCK_RANGE":         "string",
		"SHOT_DIST_RANGE":          "string",
		"SHOT_ZONE_RANGE":          "string",
		"TOUCH_TIME_RANGE":         "string",
		"FG_PCT_MID_RANGE":         "float64",
		"DISPLAY_FIRST_LAST":       "string",
		"DISPLAY_FI_LAST":          "string",
		"DISPLAY_LAST_COMMA_FIRST": "string",
		"PLAYER_NAME_LAST_FIRST":   "string",
		"FG_PCT_ABOVE_BREAK_3":     "float64",
		"FG_PCT_BACKCOURT":         "float64",
		"FG_PCT_IN_PAINT":          "float64",
		"FG_PCT_LEFT_CORNER_3":     "float64",
		"FG_PCT_RA":                "float64",
		"FG_PCT_RIGHT_CORNER_3":    "float64",
		"PCT":                      "float64",
		"WinPCT":                   "float64",
		"PCT_AST_2PM":              "float64",
		"PCT_AST_3PM":              "float64",
		"PCT_AST_FGM":              "float64",
		"PCT_BLKA":                 "float64",
		"PCT_FG3A":                 "float64",
		"PCT_FG3M":                 "float64",
		"PCT_FGA":                  "float64",
		"PCT_FGM":                  "float64",
		"PCT_FTA":                  "float64",
		"PCT_FTM":                  "float64",
		"PCT_UAST_2PM":             "float64",
		"PCT_UAST_3PM":             "float64",
		"PCT_UAST_FGM":             "float64",
		"PERCENTILE":               "float64",
		"GENERALMANAGER":           "string",
	}

	for field, want := range corrected {
		t.Run(field, func(t *testing.T) {
			if inferGoType(field) == want {
				t.Fatalf("inferGoType(%q) already returns %q - this case is no longer knownWrong, remove it from the correction map instead of leaving a redundant assertion", field, want)
			}
			got := resolveFieldGoType("", "", field)
			if got != want {
				t.Errorf("resolveFieldGoType(%q) = %q, want %q (fieldtypes.json entry missing or incorrect)", field, got, want)
			}
		})
	}
}

// TestFieldTypesMatchCommittedConsensus documents 14 fieldtypes.json
// corrections found by a different method than TestFieldTypesOverridesKnownWrongInference
// above: cross-referencing every field name against the Go type actually
// used for it across all 143 committed pkg/stats/endpoints/*.go files, and
// adopting the unanimous committed type where fieldtypes.json (inheriting
// inferGoType's unreviewed default) disagreed with it. Each of these fields
// has exactly one Go type across every occurrence in committed code -
// distinct from fields like OREB/REB/AST/PTS, which are float64 in ~90
// committed endpoints (per-game averages) but int in exactly 8 box-score
// endpoints (single-game counts): that's a real per-endpoint semantic
// difference a flat name->type dictionary cannot represent, not a case
// where one side is simply wrong, and is intentionally left alone here
// pending a per-endpoint override mechanism.
func TestFieldTypesMatchCommittedConsensus(t *testing.T) {
	corrected := map[string]string{
		"CONF_COUNT":           "int",
		"DIV_COUNT":            "int",
		"EVENTMSGACTIONTYPE":   "int",
		"EVENTMSGTYPE":         "int",
		"EVENTNUM":             "int",
		"PERSON1TYPE":          "int",
		"PERSON2TYPE":          "int",
		"PERSON3TYPE":          "int",
		"PLAYER_LAST":          "string",
		"PO_LOSSES":            "int",
		"PO_WINS":              "int",
		"PT_DIFF":              "float64",
		"TO":                   "int",
		"VIDEO_AVAILABLE_FLAG": "int",
	}

	for field, want := range corrected {
		t.Run(field, func(t *testing.T) {
			if inferGoType(field) == want {
				t.Fatalf("inferGoType(%q) already returns %q - this correction is redundant, remove it", field, want)
			}
			got := resolveFieldGoType("", "", field)
			if got != want {
				t.Errorf("resolveFieldGoType(%q) = %q, want %q (fieldtypes.json entry missing or incorrect)", field, got, want)
			}
		})
	}
}

// TestShotChartFieldsMatchShotChartDetailPrecedent documents 7 more
// fieldtypes.json corrections found neither by TestInferGoType's knownWrong
// cases nor by unanimous committed consensus (these fields disagreed across
// the only two committed occurrences, ShotChartDetail and
// ShotChartLineupDetail), but by strong semantic evidence: shot chart
// coordinates/distances/flags/event-IDs are unambiguously numeric (a court
// X coordinate, a shot distance in feet, an attempted/made 0-or-1 flag, an
// event sequence number), and ShotChartDetail's own committed code already
// had the correct type for every one of them - only ShotChartLineupDetail
// (regenerated in the same change as this test, from a plain fieldtypes.json
// default of "string"/"float64" inherited from inferGoType) had it wrong.
// SHOT_DISTANCE and LOC_X/LOC_Y are the same corruption class as the
// already-documented FG_PCT_RA bug: toString formats a float64 with "%.0f",
// so a shot chart X coordinate of 23.5 silently became "24" in
// ShotChartLineupDetail before this correction - a real, shipped bug caught
// only while investigating ShotChartDetail's regeneration, not by any
// dedicated fixture or contract test.
func TestShotChartFieldsMatchShotChartDetailPrecedent(t *testing.T) {
	corrected := map[string]string{
		"GAME_EVENT_ID":       "int",
		"MINUTES_REMAINING":   "int",
		"SHOT_DISTANCE":       "float64",
		"LOC_X":               "float64",
		"LOC_Y":               "float64",
		"SHOT_ATTEMPTED_FLAG": "int",
		"SHOT_MADE_FLAG":      "int",
	}

	for field, want := range corrected {
		t.Run(field, func(t *testing.T) {
			if inferGoType(field) == want {
				t.Fatalf("inferGoType(%q) already returns %q - this correction is redundant, remove it", field, want)
			}
			got := resolveFieldGoType("", "", field)
			if got != want {
				t.Errorf("resolveFieldGoType(%q) = %q, want %q (fieldtypes.json entry missing or incorrect)", field, got, want)
			}
		})
	}

	// SECONDS_REMAINING is deliberately NOT corrected globally: its two
	// committed occurrences disagree (ShotChartDetail: int, but
	// ShotChartLineupDetail and WinProbabilityPBP: string) without the
	// same unambiguous semantic evidence the fields above have (compare
	// WinProbabilityPBP's neighboring HOME_SCORE/VISITOR_SCORE, also
	// string - possibly a real, different, not-yet-investigated
	// pre-existing type gap in that endpoint, not proof SECONDS_REMAINING
	// itself is safe to correct everywhere). ShotChartDetail gets a
	// narrow override instead, matching only its own committed value.
	if got := resolveFieldGoType("ShotChartDetail", "Shot_Chart_Detail", "SECONDS_REMAINING"); got != "int" {
		t.Errorf(`resolveFieldGoType("ShotChartDetail", "Shot_Chart_Detail", "SECONDS_REMAINING") = %q, want "int" (fieldtype_overrides.json entry)`, got)
	}
	if got := resolveFieldGoType("", "", "SECONDS_REMAINING"); got != "string" {
		t.Errorf(`resolveFieldGoType("", "", "SECONDS_REMAINING") = %q, want "string" (the override must not leak into the global default)`, got)
	}
}

// TestTeamInfoCommonFieldsMatchCodebaseMajority documents a verification
// step for teaminfocommon.go done differently from every other case above:
// no fieldtypes.json entry needed correcting - the existing (unreviewed,
// inferGoType-inherited) global values already agree with the overwhelming
// committed-code majority, and teaminfocommon.go itself was the outlier.
// W/L are "int" in only 2 of 87 committed occurrences (the rest, including
// every "*DashboardBy*Splits" family endpoint, are "string"); PTS_RANK/
// REB_RANK/AST_RANK are "int" in only teaminfocommon.go among 13-14
// occurrences each; MIN_YEAR is "string" only in teaminfocommon.go, versus
// "float64" in its two siblings (commonteamyears.go, teaminfocommonv2.go).
// None of these are the corruption class fixed elsewhere in this file
// (int<->string<->float64 for a whole-number value loses no data either
// direction) - this test exists so a future reader doesn't mistake
// "teaminfocommon.go's types changed" for "nobody checked whether the
// dictionary was right first."
func TestTeamInfoCommonFieldsMatchCodebaseMajority(t *testing.T) {
	majority := []struct{ resultSet, field, want string }{
		{"TeamInfoCommon", "W", "string"},
		{"TeamInfoCommon", "L", "string"},
		{"TeamInfoCommon", "MIN_YEAR", "float64"},
		{"TeamSeasonRanks", "PTS_RANK", "float64"},
		{"TeamSeasonRanks", "REB_RANK", "float64"},
		{"TeamSeasonRanks", "AST_RANK", "float64"},
		{"TeamSeasonRanks", "OPP_PTS_RANK", "float64"},
	}
	for _, tc := range majority {
		t.Run(tc.resultSet+"/"+tc.field, func(t *testing.T) {
			got := resolveFieldGoType("TeamInfoCommon", tc.resultSet, tc.field)
			if got != tc.want {
				t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want %q", "TeamInfoCommon", tc.resultSet, tc.field, got, tc.want)
			}
		})
	}
}

// TestEndpointPathMatchesNameConvention is the automated form of the
// manual self-consistency check that originally found the 10 malformed
// endpoint paths documented in CHANGELOG.md's [3.1.0] section (embedded
// spaces, stray capitals, straightforward typos) - every metadata entry's
// "endpoint" field should equal its "name" field lowercased, the
// exceptionless convention this codebase uses everywhere but one
// documented case. This reads metadata directly, independent of
// generated output - unlike each endpoint's generated
// TestGet<Name>_Generated, whose path assertion only proves generated
// code matches its own metadata (it can't catch a typo already present
// in the metadata itself, since both sides of that assertion derive from
// the same source). This test can catch exactly that: a metadata typo,
// before it ever reaches generated code.
func TestEndpointPathMatchesNameConvention(t *testing.T) {
	// TeamYearOverYearSplits is the sole documented exception: its own
	// doc comment in pkg/stats/endpoints/teamyearoveryearsplits.go states
	// this is deliberate - a shorter Go-friendly type name for a real,
	// differently-named NBA.com endpoint
	// (teamdashboardbyyearoveryearsplits). Do not add further exceptions
	// here without an equivalent doc comment on the affected endpoint
	// file explaining why the convention doesn't apply.
	exceptions := map[string]string{
		"TeamYearOverYearSplits": "teamdashboardbyyearoveryearsplits",
	}

	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}
	if len(metadataFiles) == 0 {
		t.Fatal("no metadata files found under metadata/ - this test would trivially pass without exercising anything")
	}

	checked := 0
	for _, mf := range metadataFiles {
		data, err := os.ReadFile(mf)
		if err != nil {
			t.Fatalf("failed to read %s: %v", mf, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", mf, err)
		}
		for _, ep := range endpoints {
			if ep.HandlerOnly {
				continue // hand-written SDK entry, exists only to drive handler generation - no URL path of its own here
			}
			checked++
			want := strings.ToLower(ep.Name)
			if exception, ok := exceptions[ep.Name]; ok {
				want = exception
			}
			if ep.Endpoint != want {
				t.Errorf("%s (%s): endpoint = %q, want %q - doesn't match the lowercase(name) convention and isn't a documented exception", ep.Name, mf, ep.Endpoint, want)
			}
		}
	}
	if checked == 0 {
		t.Fatal("no non-handler-only metadata entries found - this test would trivially pass without exercising anything")
	}
}

// TestAllMetadataFieldsHaveExplicitTypes ensures every field name
// referenced by a committed metadata/*.json file has an explicit entry in
// fieldtypes.json, so generation never silently falls back to inferGoType
// for a field nobody has reviewed. A new endpoint's metadata should add its
// new field names to fieldtypes.json (verified against a live response) in
// the same change, not rely on the fallback.
func TestAllMetadataFieldsHaveExplicitTypes(t *testing.T) {
	types, err := loadFieldTypes()
	if err != nil {
		t.Fatalf("loadFieldTypes() failed: %v", err)
	}

	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}
	if len(metadataFiles) == 0 {
		t.Fatal("no metadata files found under metadata/ - this test would trivially pass without exercising anything")
	}

	missing := map[string][]string{} // field -> metadata files referencing it
	for _, mf := range metadataFiles {
		data, err := os.ReadFile(mf)
		if err != nil {
			t.Fatalf("failed to read %s: %v", mf, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", mf, err)
		}
		for _, ep := range endpoints {
			for _, rs := range ep.ResultSets {
				for _, field := range rs.Fields {
					if _, ok := types[field]; !ok {
						missing[field] = append(missing[field], mf)
					}
				}
			}
		}
	}

	if len(missing) > 0 {
		fields := make([]string, 0, len(missing))
		for f := range missing {
			fields = append(fields, f)
		}
		sort.Strings(fields)
		for _, f := range fields {
			t.Errorf("field %q referenced by %v has no entry in fieldtypes.json", f, missing[f])
		}
	}
}

// TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint proves
// fieldtype_overrides.json actually changes resolution where it's
// declared and leaves the global fieldtypes.json default untouched
// everywhere else. OREB is the motivating case: fieldtypes.json says
// float64 (correct for the ~90 per-game-average endpoints that don't
// override it), but BoxScoreTraditionalV2.PlayerStats overrides it to int
// (correct for that single-game box score). If this test passed with the
// override applying globally, every non-box-score endpoint's OREB would
// silently corrupt back to int.
func TestFieldTypeOverridesApplyOnlyWithinTheirEndpoint(t *testing.T) {
	globalDefault := resolveFieldGoType("", "", "OREB")
	if globalDefault != "float64" {
		t.Fatalf("resolveFieldGoType(_, _, %q) = %q, want %q (fieldtypes.json's global default) - has fieldtypes.json changed?", "OREB", globalDefault, "float64")
	}

	overridden := resolveFieldGoType("BoxScoreTraditionalV2", "PlayerStats", "OREB")
	if overridden != "int" {
		t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want %q (fieldtype_overrides.json entry)", "BoxScoreTraditionalV2", "PlayerStats", "OREB", overridden, "int")
	}

	// Same field, same endpoint, but a result set with no override entry:
	// must fall through to the global default, not leak the sibling
	// result set's override.
	unrelatedResultSet := resolveFieldGoType("BoxScoreTraditionalV2", "SomeOtherResultSetNotInOverrides", "OREB")
	if unrelatedResultSet != globalDefault {
		t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want the global default %q - override leaked to an unlisted result set", "BoxScoreTraditionalV2", "SomeOtherResultSetNotInOverrides", "OREB", unrelatedResultSet, globalDefault)
	}

	// Same field, an endpoint with no override entries at all: must fall
	// through to the global default. PlayerDashboardByGeneralSplits is a
	// real metadata-covered endpoint whose OverallPlayerDashboard result
	// set contains OREB as a per-game average, not a count.
	unrelatedEndpoint := resolveFieldGoType("PlayerDashboardByGeneralSplits", "OverallPlayerDashboard", "OREB")
	if unrelatedEndpoint != globalDefault {
		t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want the global default %q - override leaked to an unrelated endpoint", "PlayerDashboardByGeneralSplits", "OverallPlayerDashboard", "OREB", unrelatedEndpoint, globalDefault)
	}

	// TeamYearByYearStats.TeamStats and TeamGameLogs.TeamGameLogs carry a
	// second, larger set of overrides (season/game totals, not just the
	// OREB family) because every stat field in those two result sets is a
	// whole-number total in committed code, not a per-game average -
	// confirmed by TEAM_ID/GP/FGM/FGA also being int there, not just the
	// fields fieldtypes.json alone would get wrong.
	for _, tc := range []struct{ endpoint, resultSet, field string }{
		{"TeamYearByYearStats", "TeamStats", "WINS"},
		{"TeamYearByYearStats", "TeamStats", "LOSSES"},
		{"TeamYearByYearStats", "TeamStats", "TOV"},
		{"TeamGameLogs", "TeamGameLogs", "TOV"},
		{"TeamGameLogs", "TeamGameLogs", "DD2"},
		{"LeagueGameFinder", "LeagueGameFinderResults", "TOV"},
	} {
		if got := resolveFieldGoType(tc.endpoint, tc.resultSet, tc.field); got != "int" {
			t.Errorf("resolveFieldGoType(%q, %q, %q) = %q, want %q", tc.endpoint, tc.resultSet, tc.field, got, "int")
		}
	}
}

// TestFieldTypeOverridesReferenceRealMetadata ensures every
// (endpoint, result set, field) entry in fieldtype_overrides.json
// actually exists in committed metadata - the endpoint name matches some
// EndpointMetadata.Name, the result set name matches one of its
// ResultSets, and the field is in that result set's Fields. Without this,
// a typo'd or stale override entry (e.g. after a metadata file is
// renamed or a result set's fields change) would silently do nothing
// instead of failing loudly.
func TestFieldTypeOverridesReferenceRealMetadata(t *testing.T) {
	overrides, err := loadFieldTypeOverrides()
	if err != nil {
		t.Fatalf("loadFieldTypeOverrides() failed: %v", err)
	}
	if len(overrides) == 0 {
		t.Fatal("fieldtype_overrides.json is empty - this test would trivially pass without exercising anything")
	}

	metadataFiles, err := filepath.Glob("metadata/*.json")
	if err != nil {
		t.Fatalf("failed to glob metadata files: %v", err)
	}

	// endpoint -> result set -> set of fields, aggregated across every
	// metadata file (some endpoints are duplicated verbatim across more
	// than one committed file).
	realFields := map[string]map[string]map[string]bool{}
	for _, mf := range metadataFiles {
		data, err := os.ReadFile(mf)
		if err != nil {
			t.Fatalf("failed to read %s: %v", mf, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", mf, err)
		}
		for _, ep := range endpoints {
			for _, rs := range ep.ResultSets {
				if realFields[ep.Name] == nil {
					realFields[ep.Name] = map[string]map[string]bool{}
				}
				if realFields[ep.Name][rs.Name] == nil {
					realFields[ep.Name][rs.Name] = map[string]bool{}
				}
				for _, field := range rs.Fields {
					realFields[ep.Name][rs.Name][field] = true
				}
			}
		}
	}

	for endpointName, resultSets := range overrides {
		for resultSetName, fields := range resultSets {
			for field := range fields {
				if !realFields[endpointName][resultSetName][field] {
					t.Errorf("fieldtype_overrides.json has %s.%s.%s, but no committed metadata file has a %q endpoint with a %q result set containing field %q",
						endpointName, resultSetName, field, endpointName, resultSetName, field)
				}
			}
		}
	}
}
