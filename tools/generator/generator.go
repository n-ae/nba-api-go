package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"unicode"
)

// Go type names inferGoType and toParamType return, as constants so the
// same literal isn't repeated (and potentially mistyped) at each of
// inferGoType's ~15 return sites.
const (
	goTypeString  = "string"
	goTypeInt     = "int"
	goTypeFloat64 = "float64"
)

// Recognized special parameter names, as constants so the same literal
// isn't repeated (and potentially mistyped) across toParamType,
// toHandlerParamType, and processHandlerMetadata's per-name branches.
const (
	paramNameSeason     = "Season"
	paramNameSeasonType = "SeasonType"
	paramNameLeagueID   = "LeagueID"
	paramNamePerMode    = "PerMode"
)

type Generator struct {
	outputDir       string
	serverOutputDir string
	templates       map[string]*template.Template
}

func NewGenerator(outputDir, serverOutputDir string) *Generator {
	return &Generator{
		outputDir:       outputDir,
		serverOutputDir: serverOutputDir,
		templates:       make(map[string]*template.Template),
	}
}

type EndpointMetadata struct {
	Name              string              `json:"name"`
	Endpoint          string              `json:"endpoint"`
	Parameters        []ParameterMetadata `json:"parameters"`
	ResultSets        []ResultSetMetadata `json:"result_sets"`
	HasParameterTypes bool                `json:"-"`
	HasRequiredParams bool                `json:"-"`

	// RequiredParamsNeedParametersImport is true iff at least one
	// *required* parameter resolves to a non-string type. Distinct from
	// HasParameterTypes (which considers every parameter, required or
	// not) because templates/endpoint_test.tmpl only ever sets required
	// parameters in its Request struct literal - importing "parameters"
	// because an optional field happens to need it would be an unused
	// import, since that field is never referenced.
	RequiredParamsNeedParametersImport bool `json:"-"`

	// HandlerOnly marks a metadata entry that exists solely to drive HTTP
	// handler generation (see generateHandler), for an endpoint whose SDK
	// code is intentionally hand-written - pkg/stats/endpoints has no
	// generated file for it (see CLAUDE.md's Code Generation System
	// section for which 6 endpoints these are and why). generateEndpoint
	// refuses such an entry outright; GenerateFromMetadata/
	// GenerateSingleEndpoint skip the SDK-generation call for it entirely
	// rather than relying on that refusal - both exist so a future
	// `-endpoint X` invocation can never silently overwrite a
	// hand-written SDK file with a generated stub.
	HandlerOnly bool `json:"handler_only"`

	// SDKFunction overrides the SDK function name the generated handler
	// calls. Every SDK-generated endpoint's function is named "Get"+Name
	// by convention (see endpoint.tmpl); the 6 hand-written endpoints
	// don't uniformly follow this (CommonPlayerInfo, PlayerGameLog,
	// PlayerCareerStats, LeagueLeaders have no "Get" prefix). Empty means
	// "use Get"+Name" - see EffectiveSDKFunction.
	SDKFunction string `json:"sdk_function"`

	// ResponseWrapped states whether the SDK function returns
	// (*models.Response[*XResponse], error) (true, the default - the
	// generated handler unwraps .Data) or a bare (*XResponse, error)
	// (false - the generated handler uses the return value directly).
	// Every SDK-generated endpoint and 5 of the 6 hand-written ones use
	// the wrapped convention; GetInternationalBroadcasterSchedule is the
	// sole exception. A *bool (not bool) because JSON's absent-key
	// default (false) would otherwise be indistinguishable from an
	// explicit "false" - see EffectiveResponseWrapped for the resolved,
	// template-safe value (text/template's `if` on a non-nil *bool is
	// always truthy regardless of the pointee, so the template never
	// reads this field directly).
	ResponseWrapped *bool `json:"response_wrapped"`

	// EffectiveSDKFunction, EffectiveResponseWrapped, NameLower, and
	// NeedsParametersImport are all computed by processHandlerMetadata
	// and only ever read by templates/handler.tmpl and
	// templates/dispatch.tmpl - never set directly from JSON.
	EffectiveSDKFunction     string `json:"-"`
	EffectiveResponseWrapped bool   `json:"-"`
	NameLower                string `json:"-"`
	NeedsParametersImport    bool   `json:"-"`
}

type ParameterMetadata struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`

	// Pointer overrides whether this parameter's field in the target Go
	// struct is a pointer type, for handler generation only. nil means
	// "infer from Required" (pointer iff !Required), matching the SDK
	// generator's own toParamType/endpoint.tmpl convention - true for
	// every SDK-generated Request struct. Only needed as an explicit
	// override for handler_only entries, whose hand-written Request
	// structs don't uniformly follow that convention (e.g.
	// CommonPlayerInfoRequest.LeagueID is a plain value despite being
	// optional from the HTTP caller's perspective). See
	// EffectivePointer for the resolved value templates read.
	Pointer *bool `json:"pointer"`

	// HandlerGoType and EffectivePointer are computed by
	// processHandlerMetadata and only ever read by templates/handler.tmpl.
	// HandlerGoType is a superset of toParamType's output (also
	// recognizes StatCategory, needed by LeagueLeaders's hand-written
	// metadata but never produced by SDK generation) - kept separate from
	// the SDK-generation Type/toParamType path deliberately, so extending
	// it can't change what the SDK template (endpoint.tmpl) produces.
	HandlerGoType    string `json:"-"`
	EffectivePointer bool   `json:"-"`
}

type ResultSetMetadata struct {
	Name       string          `json:"name"`
	Fields     []string        `json:"fields"`
	FieldTypes []FieldTypeInfo `json:"-"`
}

type FieldTypeInfo struct {
	Name    string
	GoType  string
	JSONTag string
}

func (g *Generator) GenerateFromMetadata(metadataFile string, dryRun bool) error {
	data, err := os.ReadFile(metadataFile)
	if err != nil {
		return fmt.Errorf("failed to read metadata file: %w", err)
	}

	var endpoints []EndpointMetadata
	if err := json.Unmarshal(data, &endpoints); err != nil {
		return fmt.Errorf("failed to parse metadata: %w", err)
	}

	for _, endpoint := range endpoints {
		// processHandlerMetadata must run before processMetadata: the
		// latter mutates ParameterMetadata.Type in place via toParamType
		// (Parameters is a slice, so the mutation is visible through any
		// other copy of the same EndpointMetadata sharing its backing
		// array), and processHandlerMetadata needs the original,
		// unconverted type name.
		handlerMeta := g.processHandlerMetadata(endpoint)
		endpoint = g.processMetadata(endpoint)
		if !endpoint.HandlerOnly {
			if err := g.generateEndpoint(endpoint, dryRun); err != nil {
				return fmt.Errorf("failed to generate %s: %w", endpoint.Name, err)
			}
			if err := g.generateEndpointTest(endpoint, dryRun); err != nil {
				return fmt.Errorf("failed to generate parsing test for %s: %w", endpoint.Name, err)
			}
		}
		if err := g.generateHandler(handlerMeta, dryRun); err != nil {
			return fmt.Errorf("failed to generate handler for %s: %w", endpoint.Name, err)
		}
		if !dryRun {
			fmt.Printf("✓ Generated %s\n", endpoint.Name)
		}
	}

	return nil
}

// GenerateSingleEndpoint generates the endpoint named name by searching
// metadataDir's *.json files for a matching entry - the documented
// `-endpoint NAME` workflow. Previously this rendered a bare
// EndpointMetadata{Name, Endpoint} with zero parameters and zero result
// sets - no metadata lookup happened at all - silently producing an empty
// stub struct and reporting success, which -dry-run did not reveal
// either. That silent-stub failure mode is worse than the crash it
// replaced (a crash stops you immediately; this cost nothing until you
// checked the diff, or didn't, and shipped it) - fixed by actually
// searching for the metadata and failing loudly when it isn't found,
// via findEndpointMetadata.
func (g *Generator) GenerateSingleEndpoint(name, metadataDir string, dryRun bool) error {
	metadata, err := findEndpointMetadata(metadataDir, name)
	if err != nil {
		return err
	}
	// See GenerateFromMetadata's identical ordering comment: handler
	// processing must see metadata's original, unconverted parameter
	// types, before processMetadata mutates them in place for SDK
	// generation.
	handlerMeta := g.processHandlerMetadata(metadata)
	metadata = g.processMetadata(metadata)
	if !metadata.HandlerOnly {
		if err := g.generateEndpoint(metadata, dryRun); err != nil {
			return err
		}
		if err := g.generateEndpointTest(metadata, dryRun); err != nil {
			return err
		}
	}
	return g.generateHandler(handlerMeta, dryRun)
}

// findEndpointMetadata searches every metadata/*.json file under
// metadataDir for an entry whose Name matches name exactly, returning an
// error - not a zero-value EndpointMetadata - if no file has a matching
// entry. GenerateSingleEndpoint relies on this erroring rather than
// silently continuing with empty metadata.
func findEndpointMetadata(metadataDir, name string) (EndpointMetadata, error) {
	files, err := filepath.Glob(filepath.Join(metadataDir, "*.json"))
	if err != nil {
		return EndpointMetadata{}, fmt.Errorf("failed to glob metadata directory %s: %w", metadataDir, err)
	}
	if len(files) == 0 {
		return EndpointMetadata{}, fmt.Errorf("no metadata/*.json files found under %s", metadataDir)
	}

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return EndpointMetadata{}, fmt.Errorf("failed to read metadata file %s: %w", f, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			return EndpointMetadata{}, fmt.Errorf("failed to parse metadata file %s: %w", f, err)
		}
		for _, ep := range endpoints {
			if ep.Name == name {
				return ep, nil
			}
		}
	}

	return EndpointMetadata{}, fmt.Errorf("no metadata found for endpoint %q in any *.json file under %s - write a metadata entry for it first (see tools/generator/README.md), or use -metadata to point at a specific file directly", name, metadataDir)
}

func (g *Generator) generateEndpoint(metadata EndpointMetadata, dryRun bool) (err error) {
	if metadata.HandlerOnly {
		return fmt.Errorf("%s is marked handler_only in its metadata - it has intentionally hand-written SDK code (see CLAUDE.md's Code Generation System section), so generating its SDK file would silently overwrite hand-written code; generate its HTTP handler instead (GenerateHandler/-endpoint already do this automatically for handler_only entries)", metadata.Name)
	}

	tmpl, err := g.loadTemplate("endpoint")
	if err != nil {
		return err
	}

	filename := filepath.Join(g.outputDir, strings.ToLower(metadata.Name)+".go")

	if dryRun {
		return tmpl.Execute(os.Stdout, metadata)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close %s: %w", filename, cerr)
		}
	}()

	return tmpl.Execute(f, metadata)
}

// generateEndpointTest renders templates/endpoint_test.tmpl for metadata,
// writing <outputDir>/generated_<lowercase name>_test.go - a response-
// parsing test synthesized from the endpoint's own result-set field
// names/types, exercising findResultSet/validateHeaders/toInt/toFloat/
// toString against real generated code instead of the request-building-
// only coverage cmd/nba-api-server's TestGeneratedHandlers gets
// cross-package. Silently skips (not an error) a HandlerOnly entry (hand-
// written SDK code - generateEndpoint already refuses those, so this
// would have nothing to test against) or one with no result sets (nothing
// meaningful to assert beyond what TestGeneratedHandlers already covers).
func (g *Generator) generateEndpointTest(metadata EndpointMetadata, dryRun bool) (err error) {
	if metadata.HandlerOnly || len(metadata.ResultSets) == 0 {
		return nil
	}

	tmpl, err := g.loadTemplate("endpoint_test")
	if err != nil {
		return err
	}

	filename := filepath.Join(g.outputDir, "generated_"+strings.ToLower(metadata.Name)+"_test.go")

	if dryRun {
		return tmpl.Execute(os.Stdout, metadata)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close %s: %w", filename, cerr)
		}
	}()

	return tmpl.Execute(f, metadata)
}

func (g *Generator) loadTemplate(name string) (*template.Template, error) {
	if tmpl, ok := g.templates[name]; ok {
		return tmpl, nil
	}

	tmpl, err := template.ParseFS(templatesFS, "templates/"+name+".tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to load template %s: %w", name, err)
	}

	g.templates[name] = tmpl
	return tmpl, nil
}

func toParamType(paramType string) string {
	switch paramType {
	case paramNameSeason:
		return "parameters.Season"
	case paramNameSeasonType:
		return "parameters.SeasonType"
	case paramNameLeagueID:
		return "parameters.LeagueID"
	case paramNamePerMode:
		return "parameters.PerMode"
	default:
		return goTypeString
	}
}

func (g *Generator) processMetadata(metadata EndpointMetadata) EndpointMetadata {
	hasParameterTypes := false
	hasRequiredParams := false
	requiredParamsNeedParametersImport := false
	for i := range metadata.Parameters {
		originalType := metadata.Parameters[i].Type
		metadata.Parameters[i].Type = toParamType(originalType)
		if metadata.Parameters[i].Type != goTypeString && metadata.Parameters[i].Type != originalType {
			hasParameterTypes = true
			if metadata.Parameters[i].Required {
				requiredParamsNeedParametersImport = true
			}
		}
		if metadata.Parameters[i].Required {
			hasRequiredParams = true
		}
	}
	metadata.HasParameterTypes = hasParameterTypes
	metadata.HasRequiredParams = hasRequiredParams
	metadata.RequiredParamsNeedParametersImport = requiredParamsNeedParametersImport

	// Process result sets to infer field types
	for i := range metadata.ResultSets {
		metadata.ResultSets[i].FieldTypes = inferFieldTypes(metadata.Name, metadata.ResultSets[i].Name, metadata.ResultSets[i].Fields)
	}

	return metadata
}

// toHandlerParamType is templates/handler.tmpl's type-resolution
// counterpart to toParamType - a superset recognizing StatCategory in
// addition to Season/SeasonType/LeagueID/PerMode, needed by
// LeagueLeaders's hand-written metadata but never produced by SDK
// generation. Deliberately a separate function, not an extension of
// toParamType itself, so extending handler-generation's recognized types
// can never change what endpoint.tmpl (SDK generation) produces.
func toHandlerParamType(paramType string) string {
	switch paramType {
	case paramNameSeason:
		return "parameters.Season"
	case paramNameSeasonType:
		return "parameters.SeasonType"
	case paramNameLeagueID:
		return "parameters.LeagueID"
	case paramNamePerMode:
		return "parameters.PerMode"
	case "StatCategory":
		return "parameters.StatCategory"
	default:
		return goTypeString
	}
}

// processHandlerMetadata resolves the fields templates/handler.tmpl and
// templates/dispatch.tmpl need. Deep-copies Parameters so its result never
// depends on call order relative to processMetadata (which mutates
// ParameterMetadata.Type in place for the same entry) even if that
// invariant slips later - callers still must call this before
// processMetadata regardless, since a copy taken after processMetadata
// already mutated the shared backing array would copy the wrong (SDK,
// not raw) type name.
func (g *Generator) processHandlerMetadata(metadata EndpointMetadata) EndpointMetadata {
	params := make([]ParameterMetadata, len(metadata.Parameters))
	copy(params, metadata.Parameters)
	metadata.Parameters = params

	needsParametersImport := false
	for i := range metadata.Parameters {
		p := &metadata.Parameters[i]
		p.HandlerGoType = toHandlerParamType(p.Type)
		if p.Pointer != nil {
			p.EffectivePointer = *p.Pointer
		} else {
			p.EffectivePointer = !p.Required
		}
		switch p.Name {
		case paramNameLeagueID, paramNameSeason, paramNameSeasonType, paramNamePerMode:
			// Always resolved via the parameters package regardless of
			// HandlerGoType - see handler.tmpl's per-name branches.
			needsParametersImport = true
		default:
			if p.HandlerGoType != goTypeString {
				needsParametersImport = true
			}
		}
	}
	metadata.NeedsParametersImport = needsParametersImport
	metadata.NameLower = strings.ToLower(metadata.Name)

	if metadata.SDKFunction != "" {
		metadata.EffectiveSDKFunction = metadata.SDKFunction
	} else {
		metadata.EffectiveSDKFunction = "Get" + metadata.Name
	}

	if metadata.ResponseWrapped != nil {
		metadata.EffectiveResponseWrapped = *metadata.ResponseWrapped
	} else {
		metadata.EffectiveResponseWrapped = true
	}

	return metadata
}

// generateHandler renders templates/handler.tmpl for metadata and writes
// <serverOutputDir>/generated_<lowercase name>.go (or prints to stdout in
// dry-run mode). Unlike generateEndpoint, this never refuses a
// HandlerOnly entry - every metadata entry, hand-written-SDK or
// generated-SDK, gets a generated HTTP handler; that's the entire purpose
// of handler_only entries existing.
func (g *Generator) generateHandler(metadata EndpointMetadata, dryRun bool) (err error) {
	tmpl, err := g.loadTemplate("handler")
	if err != nil {
		return err
	}

	filename := filepath.Join(g.serverOutputDir, "generated_"+metadata.NameLower+".go")

	if dryRun {
		return tmpl.Execute(os.Stdout, metadata)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close %s: %w", filename, cerr)
		}
	}()

	return tmpl.Execute(f, metadata)
}

// GenerateDispatchTable globs every metadata/*.json file under metadataDir
// (the same directory findEndpointMetadata searches), generates a handler
// for every entry found (idempotent - regenerating an already-generated
// handler just overwrites it identically), and renders
// templates/dispatch.tmpl once against the full, combined set - the
// generated_dispatch.go map that replaces StatsHandler.ServeHTTP's
// previously hand-maintained switch statement. Unlike GenerateFromMetadata
// (one file) or GenerateSingleEndpoint (one entry), this is deliberately
// whole-corpus: the dispatch table has to be complete in one file, not
// assembled incrementally across separate generator invocations.
func (g *Generator) GenerateDispatchTable(metadataDir string, dryRun bool) error {
	files, err := filepath.Glob(filepath.Join(metadataDir, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to glob metadata directory %s: %w", metadataDir, err)
	}
	if len(files) == 0 {
		return fmt.Errorf("no metadata/*.json files found under %s", metadataDir)
	}

	// Several endpoints (BoxScoreSummaryV2, LeagueGameFinder,
	// ShotChartDetail, TeamYearByYearStats, PlayerDashboardByGeneralSplits,
	// LeagueDashPtStats, LeagueDashLineups, TeamGameLogs,
	// TeamDashboardByGeneralSplits, BoxScoreTraditionalV2, PlayByPlayV2, at
	// least) have metadata entries in more than one *.json file - harmless
	// for SDK/handler generation (both are per-name, so a duplicate just
	// overwrites the same output file identically), but fatal for a map
	// literal, which cannot have a duplicate key. seen deduplicates by
	// Name, keeping whichever file's entry is encountered first.
	seen := make(map[string]bool)
	var all []EndpointMetadata
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("failed to read metadata file %s: %w", f, err)
		}
		var endpoints []EndpointMetadata
		if err := json.Unmarshal(data, &endpoints); err != nil {
			return fmt.Errorf("failed to parse metadata file %s: %w", f, err)
		}
		for _, endpoint := range endpoints {
			if seen[endpoint.Name] {
				continue
			}
			seen[endpoint.Name] = true
			handlerMeta := g.processHandlerMetadata(endpoint)
			if err := g.generateHandler(handlerMeta, dryRun); err != nil {
				return fmt.Errorf("failed to generate handler for %s (from %s): %w", endpoint.Name, f, err)
			}
			all = append(all, handlerMeta)
		}
	}

	tmpl, err := g.loadTemplate("dispatch")
	if err != nil {
		return err
	}

	if dryRun {
		return tmpl.Execute(os.Stdout, all)
	}

	filename := filepath.Join(g.serverOutputDir, "generated_dispatch.go")
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("failed to close %s: %w", filename, cerr)
		}
	}()

	if err := tmpl.Execute(f, all); err != nil {
		return err
	}
	fmt.Printf("✓ Generated dispatch table (%d entries)\n", len(all))
	return nil
}

var (
	fieldTypesOnce sync.Once
	fieldTypes     map[string]string
	fieldTypesErr  error

	fieldTypeOverridesOnce sync.Once
	fieldTypeOverrides     map[string]map[string]map[string]string // endpoint -> result set -> field -> Go type
	fieldTypeOverridesErr  error

	fieldNameOverridesOnce sync.Once
	fieldNameOverrides     map[string]map[string]map[string]string // endpoint -> result set -> field -> Go identifier
	fieldNameOverridesErr  error
)

// loadFieldTypes parses the embedded fieldtypes.json dictionary once and
// caches the result. This is the canonical, hand-reviewed field-name -> Go
// type mapping - see resolveFieldGoType for how it's consulted.
func loadFieldTypes() (map[string]string, error) {
	fieldTypesOnce.Do(func() {
		data, err := fieldTypesFS.ReadFile("fieldtypes.json")
		if err != nil {
			fieldTypesErr = fmt.Errorf("failed to read embedded fieldtypes.json: %w", err)
			return
		}
		fieldTypesErr = json.Unmarshal(data, &fieldTypes)
	})
	return fieldTypes, fieldTypesErr
}

// loadFieldTypeOverrides parses the embedded fieldtype_overrides.json
// dictionary once and caches the result. This holds the exceptions
// fieldtypes.json's flat field-name -> Go-type model cannot represent: a
// field name whose correct type genuinely differs by endpoint. The
// motivating case is OREB/DREB/REB/AST/STL/BLK/PF/PTS, which are float64
// in ~90 committed endpoints (per-game averages) but int in exactly a
// handful of box-score/game-log endpoints (single-game counts) - both
// correct in their own context, so neither can win in a single global
// dictionary entry. See resolveFieldGoType for the resolution order.
func loadFieldTypeOverrides() (map[string]map[string]map[string]string, error) {
	fieldTypeOverridesOnce.Do(func() {
		data, err := fieldTypeOverridesFS.ReadFile("fieldtype_overrides.json")
		if err != nil {
			fieldTypeOverridesErr = fmt.Errorf("failed to read embedded fieldtype_overrides.json: %w", err)
			return
		}
		fieldTypeOverridesErr = json.Unmarshal(data, &fieldTypeOverrides)
	})
	return fieldTypeOverrides, fieldTypeOverridesErr
}

// loadFieldNameOverrides parses the embedded fieldname_overrides.json
// dictionary once and caches the result. This holds exceptions to
// goFieldName's general camelCase-capitalization rule for field names
// that don't follow it: e.g. VideoEvents' "vl"/"vt"/"gc"/"surl"/"durl"/
// "vurl"/"purl" are hand-committed as fully-uppercase short codes
// (VL, VT, GC, SURL, DURL, VURL, PURL) despite not being recognized
// initialisms by any standard convention. Scoped per (endpoint, result
// set, field) like fieldtype_overrides.json, both to avoid a collision if
// a future endpoint reuses one of these short field names for something
// that genuinely should follow the general rule, and so a stale/typo'd
// entry is easy to spot as scoped to the wrong place. See goFieldName for
// the resolution order.
func loadFieldNameOverrides() (map[string]map[string]map[string]string, error) {
	fieldNameOverridesOnce.Do(func() {
		data, err := fieldNameOverridesFS.ReadFile("fieldname_overrides.json")
		if err != nil {
			fieldNameOverridesErr = fmt.Errorf("failed to read embedded fieldname_overrides.json: %w", err)
			return
		}
		fieldNameOverridesErr = json.Unmarshal(data, &fieldNameOverrides)
	})
	return fieldNameOverrides, fieldNameOverridesErr
}

// resolveFieldGoType returns the Go type for a field within a given
// endpoint/result-set, in order of precedence:
//  1. fieldtype_overrides.json[endpointName][resultSetName][fieldName] -
//     an explicit per-endpoint exception (see loadFieldTypeOverrides).
//  2. fieldtypes.json[fieldName] - the explicit, hand-reviewed global
//     dictionary.
//  3. inferGoType(fieldName) - a fallback for fields reviewed in neither
//     file, which prints a warning to stderr rather than silently trusting
//     the heuristic, since several of its outputs are documented,
//     confirmed-wrong data-corruption bugs (see TestInferGoType).
func resolveFieldGoType(endpointName, resultSetName, fieldName string) string {
	if overrides, err := loadFieldTypeOverrides(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to load fieldtype_overrides.json (%v); ignoring overrides for %q\n", err, fieldName)
	} else if goType, ok := overrides[endpointName][resultSetName][fieldName]; ok {
		return goType
	}

	types, err := loadFieldTypes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to load fieldtypes.json (%v); falling back to inferGoType for %q\n", err, fieldName)
		return inferGoType(fieldName)
	}
	if goType, ok := types[fieldName]; ok {
		return goType
	}
	inferred := inferGoType(fieldName)
	fmt.Fprintf(os.Stderr, "warning: %q has no explicit entry in fieldtypes.json; falling back to inferGoType's guess (%q) - verify against a live response and add it to fieldtypes.json\n", fieldName, inferred)
	return inferred
}

// goInitialisms are recognized initialisms that goFieldName fully
// uppercases when they appear as a whole camelCase word, matching the
// convention golang.org/x/lint's default initialisms list uses. This is
// also the convention the hand-fixed committed field names for NBA
// Live-Data-style endpoints (playbyplayv3.go, scoreboardv3.go,
// leaguestandings.go) already follow: "gameId" -> "GameID", not
// "GameId". Keys are lowercase for case-insensitive lookup.
var goInitialisms = map[string]bool{
	"id": true, "url": true, "uuid": true, "api": true, "html": true,
	"http": true, "https": true, "json": true, "xml": true, "ui": true,
}

// goFieldName converts a metadata field name into a valid, exported Go
// identifier, in order of precedence:
//  1. fieldname_overrides.json[endpointName][resultSetName][fieldName] -
//     an explicit per-endpoint exception for a field name the general
//     rule below gets wrong (see loadFieldNameOverrides).
//  2. The general rule: field names in this project's metadata are
//     usually already valid, exported Go identifiers
//     (SCREAMING_SNAKE_CASE, e.g. "GAME_ID") and are returned unchanged -
//     this only acts when the first rune is lowercase. Some NBA
//     Live-Data-style endpoints use camelCase field names (e.g. "gameId")
//     that are NOT valid exported Go identifiers as-is: an unexported
//     struct field is inaccessible to any external caller and invisible
//     to encoding/json in either direction, which is a real bug this
//     fixes, not a style preference.
//
// The original field name is preserved verbatim as the JSON tag (see
// FieldTypeInfo.JSONTag) regardless of which path is taken - only the Go
// identifier changes.
func goFieldName(endpointName, resultSetName, fieldName string) string {
	if overrides, err := loadFieldNameOverrides(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to load fieldname_overrides.json (%v); ignoring overrides for %q\n", err, fieldName)
	} else if name, ok := overrides[endpointName][resultSetName][fieldName]; ok {
		return name
	}

	if fieldName == "" {
		return fieldName
	}
	runes := []rune(fieldName)
	if unicode.IsUpper(runes[0]) {
		return fieldName
	}

	var words []string
	start := 0
	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) && !unicode.IsUpper(runes[i-1]) {
			words = append(words, string(runes[start:i]))
			start = i
		}
	}
	words = append(words, string(runes[start:]))

	var b strings.Builder
	for _, word := range words {
		if word == "" {
			continue
		}
		if goInitialisms[strings.ToLower(word)] {
			b.WriteString(strings.ToUpper(word))
			continue
		}
		wr := []rune(word)
		b.WriteRune(unicode.ToUpper(wr[0]))
		b.WriteString(string(wr[1:]))
	}
	return b.String()
}

// inferFieldTypes resolves Go types for NBA API field names via
// resolveFieldGoType (the name is historical; despite it, resolution now
// prefers fieldtype_overrides.json and fieldtypes.json and only infers as
// a last-resort fallback).
func inferFieldTypes(endpointName, resultSetName string, fields []string) []FieldTypeInfo {
	result := make([]FieldTypeInfo, len(fields))
	for i, field := range fields {
		result[i] = FieldTypeInfo{
			Name:    goFieldName(endpointName, resultSetName, field),
			GoType:  resolveFieldGoType(endpointName, resultSetName, field),
			JSONTag: field,
		}
	}
	return result
}

// inferGoType infers the Go type from a field name using NBA API conventions
func inferGoType(fieldName string) string {
	lower := strings.ToLower(fieldName)

	// Percentage fields are always float64
	if strings.HasSuffix(lower, "_pct") || strings.HasSuffix(lower, "_percentage") {
		return goTypeFloat64
	}

	// ID fields - check for specific patterns
	if strings.HasSuffix(lower, "_id") {
		// Most IDs are strings (e.g., GAME_ID, SEASON_ID)
		// But PLAYER_ID, TEAM_ID are typically int
		if strings.Contains(lower, "player") || strings.Contains(lower, "team") {
			return goTypeInt
		}
		return goTypeString
	}

	// Date fields are strings
	if strings.Contains(lower, "date") {
		return goTypeString
	}

	// Text/name fields are strings
	if strings.HasSuffix(lower, "_name") || strings.HasSuffix(lower, "_text") ||
		strings.HasSuffix(lower, "_abbreviation") || strings.HasSuffix(lower, "_city") ||
		strings.HasSuffix(lower, "_tricode") || strings.Contains(lower, "nickname") ||
		strings.Contains(lower, "matchup") || strings.Contains(lower, "comment") ||
		strings.Contains(lower, "position") {
		return goTypeString
	}

	// Win/Loss indicator
	if lower == "wl" || lower == "w_l" {
		return goTypeString
	}

	// Season-related fields
	if strings.Contains(lower, "season") && !strings.Contains(lower, "id") {
		return goTypeString
	}

	// Statistical fields - most are numbers
	// Common stat abbreviations
	statAbbreviations := []string{
		"pts", "reb", "ast", "stl", "blk", "tov", "pf", "fgm", "fga", "ftm", "fta",
		"fg3m", "fg3a", "oreb", "dreb", "min", "gp", "gs", "plus_minus", "pfd",
		"blka", "dd2", "td3", "fantasy", "_count", "_games", "_rank",
	}

	for _, abbrev := range statAbbreviations {
		if strings.Contains(lower, abbrev) {
			// MIN (minutes) is typically float64
			if strings.Contains(lower, "min") && !strings.Contains(lower, "game") {
				return goTypeFloat64
			}
			// Made/Attempted stats can be int or float depending on context
			// For box scores, they're typically int
			if strings.HasSuffix(lower, "m") || strings.HasSuffix(lower, "a") {
				return goTypeInt
			}
			// Most other stats are float64 (especially averages)
			if strings.Contains(lower, "avg") || strings.Contains(lower, "per") {
				return goTypeFloat64
			}
			// Game counts are int
			if lower == "gp" || lower == "gs" {
				return goTypeInt
			}
			// Default for stats is float64
			return goTypeFloat64
		}
	}

	// Age is int
	if strings.Contains(lower, "age") {
		return goTypeInt
	}

	// Rank is int
	if strings.Contains(lower, "rank") {
		return goTypeInt
	}

	// Sequence/period numbers are int
	if strings.Contains(lower, "sequence") || strings.Contains(lower, "period") ||
		strings.Contains(lower, "range") {
		return goTypeInt
	}

	// Status codes are typically int
	if strings.Contains(lower, "status") && strings.HasSuffix(lower, "_id") {
		return goTypeInt
	}

	// Default to string for safety
	return goTypeString
}
