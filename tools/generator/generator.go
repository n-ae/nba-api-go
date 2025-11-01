package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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

func (g *Generator) generateEndpoint(metadata EndpointMetadata, dryRun bool) error {
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
	defer f.Close()

	return tmpl.Execute(f, metadata)
}

func (g *Generator) loadTemplate(name string) (*template.Template, error) {
	if tmpl, ok := g.templates[name]; ok {
		return tmpl, nil
	}

	tmplPath := filepath.Join("tools", "generator", "templates", name+".tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load template %s: %w", name, err)
	}

	g.templates[name] = tmpl
	return tmpl, nil
}

func toGoType(pythonType string) string {
	switch pythonType {
	case "str", "string":
		return "string"
	case "int", "integer":
		return "int"
	case "float":
		return "float64"
	case "bool", "boolean":
		return "bool"
	default:
		return "string"
	}
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
		return "string"
	}
}

func (g *Generator) processMetadata(metadata EndpointMetadata) EndpointMetadata {
	hasParameterTypes := false
	hasRequiredParams := false
	for i := range metadata.Parameters {
		originalType := metadata.Parameters[i].Type
		metadata.Parameters[i].Type = toParamType(originalType)
		if metadata.Parameters[i].Type != "string" && metadata.Parameters[i].Type != originalType {
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
		metadata.ResultSets[i].FieldTypes = inferFieldTypes(metadata.ResultSets[i].Fields)
	}

	return metadata
}

// inferFieldTypes infers Go types from NBA API field names
func inferFieldTypes(fields []string) []FieldTypeInfo {
	fieldTypes := make([]FieldTypeInfo, len(fields))
	for i, field := range fields {
		fieldTypes[i] = FieldTypeInfo{
			Name:    field,
			GoType:  inferGoType(field),
			JSONTag: field,
		}
	}
	return fieldTypes
}

// inferGoType infers the Go type from a field name using NBA API conventions
func inferGoType(fieldName string) string {
	lower := strings.ToLower(fieldName)

	// Percentage fields are always float64
	if strings.HasSuffix(lower, "_pct") || strings.HasSuffix(lower, "_percentage") {
		return "float64"
	}

	// ID fields - check for specific patterns
	if strings.HasSuffix(lower, "_id") {
		// Most IDs are strings (e.g., GAME_ID, SEASON_ID)
		// But PLAYER_ID, TEAM_ID are typically int
		if strings.Contains(lower, "player") || strings.Contains(lower, "team") {
			return "int"
		}
		return "string"
	}

	// Date fields are strings
	if strings.Contains(lower, "date") {
		return "string"
	}

	// Text/name fields are strings
	if strings.HasSuffix(lower, "_name") || strings.HasSuffix(lower, "_text") ||
		strings.HasSuffix(lower, "_abbreviation") || strings.HasSuffix(lower, "_city") ||
		strings.HasSuffix(lower, "_tricode") || strings.Contains(lower, "nickname") ||
		strings.Contains(lower, "matchup") || strings.Contains(lower, "comment") ||
		strings.Contains(lower, "position") {
		return "string"
	}

	// Win/Loss indicator
	if lower == "wl" || lower == "w_l" {
		return "string"
	}

	// Season-related fields
	if strings.Contains(lower, "season") && !strings.Contains(lower, "id") {
		return "string"
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
				return "float64"
			}
			// Made/Attempted stats can be int or float depending on context
			// For box scores, they're typically int
			if strings.HasSuffix(lower, "m") || strings.HasSuffix(lower, "a") {
				return "int"
			}
			// Most other stats are float64 (especially averages)
			if strings.Contains(lower, "avg") || strings.Contains(lower, "per") {
				return "float64"
			}
			// Game counts are int
			if lower == "gp" || lower == "gs" {
				return "int"
			}
			// Default for stats is float64
			return "float64"
		}
	}

	// Age is int
	if strings.Contains(lower, "age") {
		return "int"
	}

	// Rank is int
	if strings.Contains(lower, "rank") {
		return "int"
	}

	// Sequence/period numbers are int
	if strings.Contains(lower, "sequence") || strings.Contains(lower, "period") ||
		strings.Contains(lower, "range") {
		return "int"
	}

	// Status codes are typically int
	if strings.Contains(lower, "status") && strings.HasSuffix(lower, "_id") {
		return "int"
	}

	// Default to string for safety
	return "string"
}
