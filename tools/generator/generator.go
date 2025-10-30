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
}

type ParameterMetadata struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

type ResultSetMetadata struct {
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
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
	for i := range metadata.Parameters {
		originalType := metadata.Parameters[i].Type
		metadata.Parameters[i].Type = toParamType(originalType)
		if metadata.Parameters[i].Type != "string" && metadata.Parameters[i].Type != originalType {
			hasParameterTypes = true
		}
	}
	metadata.HasParameterTypes = hasParameterTypes
	return metadata
}
