package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// TestEndpointInventory guards against the exact kind of drift fixed on
// 2026-07-19: at one point five different numbers (139/140/142/143/149)
// were all presented somewhere in this repo as "the" endpoint count. It
// statically counts the real SDK endpoint files and the server's route
// cases from source, then cross-checks both against the numbers this
// server actually reports via /health's EndpointsCount - so an endpoint
// added to the SDK without updating the route switch (or vice versa), or
// either without updating EndpointsCount, fails here instead of drifting
// silently again.
func TestEndpointInventory(t *testing.T) {
	sdkTotal := countSDKEndpointFiles(t)
	httpExposed := countHandlerRouteCases(t)
	documented := parseDocumentedEndpointsCount(t)

	if !documented.found {
		t.Fatal("could not find the EndpointsCount map literal in main.go - has handleHealth been restructured? update this test to match")
	}

	if sdkTotal != documented.sdkTotal {
		t.Errorf("pkg/stats/endpoints has %d real endpoint files, but main.go's EndpointsCount[\"sdk_total\"] says %d - update whichever is wrong", sdkTotal, documented.sdkTotal)
	}
	if httpExposed != documented.httpExposed {
		t.Errorf("handlers.go's StatsHandler switch has %d route cases, but main.go's EndpointsCount[\"http_exposed\"] says %d - update whichever is wrong", httpExposed, documented.httpExposed)
	}
}

// countSDKEndpointFiles counts pkg/stats/endpoints/*.go files that
// represent a real generated (or hand-written) endpoint, excluding test
// files and the two files that are shared helpers rather than endpoints:
// dates.go (BIRTHDATE parsing) and types.go (toInt/toFloat/toString).
func countSDKEndpointFiles(t *testing.T) int {
	t.Helper()

	dir := filepath.Join("..", "..", "pkg", "stats", "endpoints")
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("failed to read %s: %v", dir, err)
	}

	excluded := map[string]bool{"dates.go": true, "types.go": true}

	count := 0
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		if excluded[name] {
			continue
		}
		count++
	}
	return count
}

// countHandlerRouteCases parses this package's handlers.go and counts the
// case labels in the StatsHandler.ServeHTTP switch statement (the file
// has exactly one switch statement; if that ever changes, this will
// over-count and this test's failure message will point here).
func countHandlerRouteCases(t *testing.T) int {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "handlers.go", nil, 0)
	if err != nil {
		t.Fatalf("failed to parse handlers.go: %v", err)
	}

	count := 0
	ast.Inspect(file, func(n ast.Node) bool {
		sw, ok := n.(*ast.SwitchStmt)
		if !ok {
			return true
		}
		for _, stmt := range sw.Body.List {
			clause, ok := stmt.(*ast.CaseClause)
			if !ok {
				continue
			}
			// A nil List means "default:", which isn't a route.
			count += len(clause.List)
		}
		return true
	})
	return count
}

type documentedEndpointsCount struct {
	sdkTotal    int
	httpExposed int
	found       bool
}

// parseDocumentedEndpointsCount parses this package's main.go and reads
// the sdk_total/http_exposed values out of the map[string]int literal
// assigned to EndpointsCount in handleHealth, without hardcoding the
// expected numbers here - so this test only needs to change when the
// literal itself changes shape, not every time an endpoint is added.
func parseDocumentedEndpointsCount(t *testing.T) documentedEndpointsCount {
	t.Helper()

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", nil, 0)
	if err != nil {
		t.Fatalf("failed to parse main.go: %v", err)
	}

	var result documentedEndpointsCount
	ast.Inspect(file, func(n ast.Node) bool {
		lit, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}
		mapType, ok := lit.Type.(*ast.MapType)
		if !ok {
			return true
		}
		keyIdent, ok := mapType.Key.(*ast.Ident)
		if !ok || keyIdent.Name != "string" {
			return true
		}
		valIdent, ok := mapType.Value.(*ast.Ident)
		if !ok || valIdent.Name != "int" {
			return true
		}

		for _, elt := range lit.Elts {
			kv, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}
			keyLit, ok := kv.Key.(*ast.BasicLit)
			if !ok || keyLit.Kind != token.STRING {
				continue
			}
			key, err := strconv.Unquote(keyLit.Value)
			if err != nil {
				continue
			}
			valLit, ok := kv.Value.(*ast.BasicLit)
			if !ok || valLit.Kind != token.INT {
				continue
			}
			val, err := strconv.Atoi(valLit.Value)
			if err != nil {
				continue
			}

			switch key {
			case "sdk_total":
				result.sdkTotal = val
				result.found = true
			case "http_exposed":
				result.httpExposed = val
				result.found = true
			}
		}
		return true
	})
	return result
}
