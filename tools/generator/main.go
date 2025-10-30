package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		endpoint   = flag.String("endpoint", "", "Endpoint name to generate (e.g., PlayerGameLog)")
		metadataFile = flag.String("metadata", "", "Path to metadata JSON file")
		outputDir  = flag.String("output", "pkg/stats/endpoints", "Output directory for generated files")
		dryRun     = flag.Bool("dry-run", false, "Print generated code without writing files")
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
