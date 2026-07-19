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

func main() {
	var (
		endpoint     = flag.String("endpoint", "", "Endpoint name to generate (e.g., PlayerGameLog)")
		metadataFile = flag.String("metadata", "", "Path to metadata JSON file")
		outputDir    = flag.String("output", defaultOutputDir(), "Output directory for generated files (default: <repo-root>/pkg/stats/endpoints, resolved independently of the working directory)")
		dryRun       = flag.Bool("dry-run", false, "Print generated code without writing files")
	)

	flag.Parse()

	if *endpoint == "" && *metadataFile == "" {
		fmt.Println("NBA API Go - Endpoint Code Generator")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  generator -endpoint PlayerGameLog")
		fmt.Println("  generator -metadata endpoints.json")
		fmt.Println("  generator -endpoint PlayerGameLog -dry-run")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	generator := NewGenerator(*outputDir)

	if *metadataFile != "" {
		if err := generator.GenerateFromMetadata(*metadataFile, *dryRun); err != nil {
			log.Fatalf("Failed to generate from metadata: %v", err)
		}
	} else if *endpoint != "" {
		if err := generator.GenerateSingleEndpoint(*endpoint, *dryRun); err != nil {
			log.Fatalf("Failed to generate endpoint: %v", err)
		}
	}

	fmt.Println("✅ Code generation complete")
}
