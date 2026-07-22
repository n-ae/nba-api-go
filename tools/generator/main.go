package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// defaultOutputDir resolves to <repo-root>/pkg/stats/endpoints regardless
// of the process's working directory, using the compile-time-known
// location of this source file rather than a path relative to CWD. The
// documented workflow is `cd tools/generator && go run . -endpoint X`, but
// a CWD-relative default ("pkg/stats/endpoints") would then resolve under
// tools/generator/ instead of the repo root - this file lives at
// tools/generator/main.go, so its directory's grandparent is the repo
// root.
func defaultOutputDir() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		// Practically unreachable for a normally compiled binary; fall
		// back to the old CWD-relative behavior rather than fail outright.
		return filepath.Join("pkg", "stats", "endpoints")
	}
	repoRoot := filepath.Dir(filepath.Dir(filepath.Dir(thisFile)))
	return filepath.Join(repoRoot, "pkg", "stats", "endpoints")
}

// defaultMetadataDir resolves to tools/generator/metadata regardless of
// the process's working directory, for the same reason and via the same
// runtime.Caller(0) approach as defaultOutputDir: `-endpoint NAME` (see
// GenerateSingleEndpoint) needs to reliably find the metadata directory
// to search, not assume the working directory happens to be
// tools/generator.
func defaultMetadataDir() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return "metadata"
	}
	return filepath.Join(filepath.Dir(thisFile), "metadata")
}

// defaultServerOutputDir resolves to <repo-root>/cmd/nba-api-server, via
// the same runtime.Caller(0) approach as defaultOutputDir - see its doc
// comment for why a CWD-relative default would be wrong under the
// documented `cd tools/generator && go run . ...` workflow.
func defaultServerOutputDir() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join("cmd", "nba-api-server")
	}
	repoRoot := filepath.Dir(filepath.Dir(filepath.Dir(thisFile)))
	return filepath.Join(repoRoot, "cmd", "nba-api-server")
}

func main() {
	var (
		endpoint     = flag.String("endpoint", "", "Endpoint name to generate (e.g., PlayerGameLog) - also generates its HTTP handler, see cmd/nba-api-server")
		metadataFile = flag.String("metadata", "", "Path to metadata JSON file - also generates each entry's HTTP handler")
		outputDir    = flag.String("output", defaultOutputDir(), "Output directory for generated SDK files (default: <repo-root>/pkg/stats/endpoints, resolved independently of the working directory)")
		serverOutput = flag.String("server-output", defaultServerOutputDir(), "Output directory for generated HTTP handler files (default: <repo-root>/cmd/nba-api-server, resolved independently of the working directory)")
		allHandlers  = flag.Bool("all-handlers", false, "Regenerate every endpoint's HTTP handler plus the dispatch table (cmd/nba-api-server/generated_dispatch.go) from every tools/generator/metadata/*.json file")
		dryRun       = flag.Bool("dry-run", false, "Print generated code without writing files")
	)

	flag.Parse()

	if *endpoint == "" && *metadataFile == "" && !*allHandlers {
		fmt.Println("NBA API Go - Endpoint Code Generator")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  generator -endpoint PlayerGameLog")
		fmt.Println("  generator -metadata endpoints.json")
		fmt.Println("  generator -endpoint PlayerGameLog -dry-run")
		fmt.Println("  generator -all-handlers")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	generator := NewGenerator(*outputDir, *serverOutput)

	if *metadataFile != "" {
		if err := generator.GenerateFromMetadata(*metadataFile, *dryRun); err != nil {
			log.Fatalf("Failed to generate from metadata: %v", err)
		}
	} else if *endpoint != "" {
		if err := generator.GenerateSingleEndpoint(*endpoint, defaultMetadataDir(), *dryRun); err != nil {
			log.Fatalf("Failed to generate endpoint: %v", err)
		}
	} else if *allHandlers {
		if err := generator.GenerateDispatchTable(defaultMetadataDir(), *dryRun); err != nil {
			log.Fatalf("Failed to generate dispatch table: %v", err)
		}
	}

	fmt.Println("✅ Code generation complete")
}
