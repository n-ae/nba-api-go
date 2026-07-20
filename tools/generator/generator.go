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

type Generator struct {
	outputDir string
	templates map[string]*template.Template
}

func NewGenerator(outputDir string) *Generator {
	return &Generator{
		outputDir: outputDir,
		templates: make(map[string]*template.Template),
	}
}

type EndpointMetadata struct {
	Name              string              `json:"name"`
	Endpoint          string              `json:"endpoint"`
	Parameters        []ParameterMetadata `json:"parameters"`
	ResultSets        []ResultSetMetadata `json:"result_sets"`
	HasParameterTypes bool                `json:"-"`
	HasRequiredParams bool                `json:"-"`
}

type ParameterMetadata struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
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
		endpoint = g.processMetadata(endpoint)
		if err := g.generateEndpoint(endpoint, dryRun); err != nil {
			return fmt.Errorf("failed to generate %s: %w", endpoint.Name, err)
		}
		if !dryRun {
			fmt.Printf("✓ Generated %s\n", endpoint.Name)
		}
	}

	return nil
}

func (g *Generator) GenerateSingleEndpoint(name string, dryRun bool) error {
	metadata := EndpointMetadata{
		Name:     name,
		Endpoint: strings.ToLower(name),
	}

	return g.generateEndpoint(metadata, dryRun)
}

func (g *Generator) generateEndpoint(metadata EndpointMetadata, dryRun bool) (err error) {
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
	case "Season":
		return "parameters.Season"
	case "SeasonType":
		return "parameters.SeasonType"
	case "LeagueID":
		return "parameters.LeagueID"
	case "PerMode":
		return "parameters.PerMode"
	default:
		return goTypeString
	}
}

func (g *Generator) processMetadata(metadata EndpointMetadata) EndpointMetadata {
	hasParameterTypes := false
	hasRequiredParams := false
	for i := range metadata.Parameters {
		originalType := metadata.Parameters[i].Type
		metadata.Parameters[i].Type = toParamType(originalType)
		if metadata.Parameters[i].Type != goTypeString && metadata.Parameters[i].Type != originalType {
			hasParameterTypes = true
		}
		if metadata.Parameters[i].Required {
			hasRequiredParams = true
		}
	}
	metadata.HasParameterTypes = hasParameterTypes
	metadata.HasRequiredParams = hasRequiredParams

	// Process result sets to infer field types
	for i := range metadata.ResultSets {
		metadata.ResultSets[i].FieldTypes = inferFieldTypes(metadata.Name, metadata.ResultSets[i].Name, metadata.ResultSets[i].Fields)
	}

	return metadata
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
