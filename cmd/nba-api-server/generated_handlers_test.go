package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/n-ae/nba-api-go/v3/pkg/stats"
)

// metadataParam and metadataEndpoint decode just enough of
// tools/generator/metadata/*.json to drive TestGeneratedHandlers below -
// name and each parameter's required-ness. Deliberately not importing
// tools/generator itself (a separate Go module, and package main besides,
// so not importable as a library) - reading its metadata files directly
// keeps this test's required-parameter list in sync with what actually
// generates the handlers under test, without a second, hand-maintained
// list that could drift from it.
type metadataParam struct {
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

type metadataEndpoint struct {
	Name       string          `json:"name"`
	Parameters []metadataParam `json:"parameters"`
}

// loadGeneratedEndpointMetadata reads every tools/generator/metadata/*.json
// file, deduplicating by Name the same way
// tools/generator's GenerateDispatchTable does (several endpoints have
// entries in more than one file - harmless there since generation is
// per-name, and harmless here since we only need the required-parameter
// set once per endpoint).
func loadGeneratedEndpointMetadata(t *testing.T) []metadataEndpoint {
	t.Helper()

	dir := filepath.Join("..", "..", "tools", "generator", "metadata")
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		t.Fatalf("failed to glob metadata dir %s: %v", dir, err)
	}
	if len(files) == 0 {
		t.Fatalf("no metadata files found under %s", dir)
	}

	seen := make(map[string]bool)
	var all []metadataEndpoint
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("failed to read %s: %v", f, err)
		}
		var endpoints []metadataEndpoint
		if err := json.Unmarshal(data, &endpoints); err != nil {
			t.Fatalf("failed to parse %s: %v", f, err)
		}
		for _, ep := range endpoints {
			if seen[ep.Name] {
				continue
			}
			seen[ep.Name] = true
			all = append(all, ep)
		}
	}
	return all
}

// dummyValueForParam returns a value for a required query parameter that
// passes pkg/stats/parameters validation, not just "non-empty" - a
// handful of parameter names carry a validated typed wrapper
// (LeagueID/Season/SeasonType/PerMode) rather than a free-form string, and
// a value like "x" would fail that validation before ever reaching the
// stub backend, turning every one of those endpoints' "valid request"
// subtest into a false failure unrelated to what's being tested.
func dummyValueForParam(name string) string {
	switch name {
	case "LeagueID":
		return "00"
	case "Season":
		return "2023-24"
	case "SeasonType":
		return "Regular Season"
	case "PerMode":
		return "PerGame"
	default:
		return "1"
	}
}

// TestGeneratedHandlers is a data-driven smoke test over every handler
// tools/generator generates (see generated_dispatch.go and
// tools/generator/templates/handler.tmpl): for each endpoint, a request
// missing a required parameter must 400, and a request with every
// required parameter present must succeed against a stub upstream. Reads
// the same tools/generator/metadata/*.json files generation itself uses,
// so a new endpoint's metadata is automatically covered here too, without
// a parallel hand-maintained list to keep in sync.
func TestGeneratedHandlers(t *testing.T) {
	endpoints := loadGeneratedEndpointMetadata(t)
	if len(endpoints) == 0 {
		t.Fatal("no endpoint metadata found - this test would trivially pass without exercising anything")
	}

	// stubUpstream stands in for stats.nba.com: returns a minimal,
	// empty-but-valid envelope for every request regardless of path or
	// query. findResultSet (pkg/stats/endpoints/types.go) treats "result
	// set not found" as "leave this field empty," not an error - and the
	// one endpoint with a genuinely different raw shape
	// (LeagueLeaders's singular "resultSet", not "resultSets") degrades
	// the same tolerant way when the key it looks for is simply absent -
	// so one generic stub can exercise all 142 differently-shaped
	// endpoints without per-endpoint fixture data.
	stubUpstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"resultSets": []}`))
	}))
	defer stubUpstream.Close()

	client, err := stats.NewClient(stats.Config{BaseURL: stubUpstream.URL})
	if err != nil {
		t.Fatalf("stats.NewClient: %v", err)
	}
	handler := NewStatsHandler(client)

	for _, ep := range endpoints {
		t.Run(ep.Name, func(t *testing.T) {
			route, ok := generatedDispatch[strings.ToLower(ep.Name)]
			if !ok {
				t.Fatalf("generatedDispatch has no entry for %q - regenerate via tools/generator (-all-handlers)", ep.Name)
			}

			var required []string
			for _, p := range ep.Parameters {
				if p.Required {
					required = append(required, p.Name)
				}
			}

			if len(required) > 0 {
				t.Run("missing required parameter returns 400", func(t *testing.T) {
					req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/"+strings.ToLower(ep.Name), nil)
					w := httptest.NewRecorder()
					route(handler, w, req)
					if w.Code != http.StatusBadRequest {
						t.Errorf("got status %d, want 400 (missing required parameter %q); body=%s", w.Code, required[0], w.Body.String())
					}
				})
			}

			t.Run("valid request succeeds", func(t *testing.T) {
				q := url.Values{}
				for _, name := range required {
					q.Set(name, dummyValueForParam(name))
				}
				target := "/api/v1/stats/" + strings.ToLower(ep.Name)
				if encoded := q.Encode(); encoded != "" {
					target += "?" + encoded
				}
				req := httptest.NewRequest(http.MethodGet, target, nil)
				w := httptest.NewRecorder()
				route(handler, w, req)
				if w.Code != http.StatusOK {
					t.Errorf("got status %d, want 200; body=%s", w.Code, w.Body.String())
				}
			})
		})
	}
}
