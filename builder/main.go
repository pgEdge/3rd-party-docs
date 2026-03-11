//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// pgdoc-converter converts PostgreSQL DocBook SGML documentation
// to Markdown for MkDocs Material.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pgEdge/postgresql-docs/builder/convert"
	"github.com/pgEdge/postgresql-docs/builder/nav"
	"github.com/pgEdge/postgresql-docs/builder/sgml"
	"github.com/pgEdge/postgresql-docs/builder/validate"
)

func main() {
	srcDir := flag.String("src", "", "Path to PostgreSQL doc/src/sgml/ directory")
	outDir := flag.String("out", "./docs", "Output directory for .md files")
	mkdocsFile := flag.String("mkdocs", "./mkdocs.yml", "Path to mkdocs.yml")
	version := flag.String("version", "", "PostgreSQL version label (e.g., 17.2)")
	doValidate := flag.Bool("validate", false, "Run link validation after conversion")
	verbose := flag.Bool("verbose", false, "Verbose output")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pgdoc-converter [flags]\n\n")
		fmt.Fprintf(os.Stderr, "Converts PostgreSQL DocBook SGML docs to Markdown.\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *srcDir == "" {
		fmt.Fprintln(os.Stderr, "error: -src is required")
		flag.Usage()
		os.Exit(1)
	}

	// Verify source directory exists
	if _, err := os.Stat(*srcDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "error: source directory %q does not exist\n", *srcDir)
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("Source: %s\n", *srcDir)
		fmt.Printf("Output: %s\n", *outDir)
		fmt.Printf("Version: %s\n", *version)
	}

	// Phase 0: Generate missing SGML files (version.sgml, errcodes, etc.)
	if *verbose {
		fmt.Println("Phase 0: Generating missing SGML files...")
	}
	genCount, genWarnings := sgml.GenerateMissingFiles(*srcDir, *version)
	if *verbose {
		fmt.Printf("  Generated %d files\n", genCount)
		for _, w := range genWarnings {
			fmt.Printf("  Warning: %s\n", w)
		}
	}
	defer sgml.CleanGeneratedFiles(*srcDir)

	// Phase 1: Entity resolution
	if *verbose {
		fmt.Println("Phase 1: Resolving entities...")
	}
	resolver := sgml.NewEntityResolver(*srcDir)
	body, err := resolver.ResolveFile("postgres.sgml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving entities: %v\n", err)
		os.Exit(1)
	}

	if *verbose {
		fmt.Printf("  Resolved %d entities\n", resolver.EntityCount())
		fmt.Printf("  Document body: %d bytes\n", len(body))
		for _, w := range resolver.Warnings() {
			fmt.Printf("  Warning: %s\n", w)
		}
	}

	// Phase 2: Parse SGML
	if *verbose {
		fmt.Println("Phase 2: Parsing SGML...")
	}
	root, warnings, err := sgml.ParseString(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing SGML: %v\n", err)
		os.Exit(1)
	}
	if *verbose && len(warnings) > 0 {
		fmt.Printf("  Parser warnings: %d\n", len(warnings))
		for _, w := range warnings {
			fmt.Printf("    %s\n", w)
		}
	}

	// Phases 3-4: Convert
	if *verbose {
		fmt.Println("Phase 3: Building ID map...")
		fmt.Println("Phase 4: Converting to Markdown...")
	}
	converter := convert.NewConverter(root, *srcDir, *outDir, *version)
	if err := converter.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "error converting: %v\n", err)
		os.Exit(1)
	}

	convWarnings := converter.Warnings()
	if *verbose && len(convWarnings) > 0 {
		fmt.Printf("  Conversion warnings: %d\n", len(convWarnings))
		for _, w := range convWarnings {
			fmt.Printf("    %s\n", w)
		}
	}

	files := converter.Files()
	if *verbose {
		fmt.Printf("  Generated %d files\n", len(files))
	}

	// Phase 5: Nav generation
	if *verbose {
		fmt.Println("Phase 5: Generating nav...")
	}
	navRoot := nav.BuildNav(files)
	navYAML := nav.GenerateYAML(navRoot)

	if *mkdocsFile != "" {
		if _, err := os.Stat(*mkdocsFile); err == nil {
			if err := nav.UpdateMkdocsYML(*mkdocsFile, navYAML, *version); err != nil {
				fmt.Fprintf(os.Stderr, "error updating mkdocs.yml: %v\n", err)
				os.Exit(1)
			}
			if *verbose {
				fmt.Printf("  Updated %s\n", *mkdocsFile)
			}
		} else if *verbose {
			fmt.Printf("  %s not found, skipping nav update\n", *mkdocsFile)
		}
	}

	// Phase 6: Validation
	if *doValidate {
		if *verbose {
			fmt.Println("Phase 6: Validating...")
		}
		result, err := validate.ValidateDir(*outDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error validating: %v\n", err)
			os.Exit(1)
		}

		if len(result.BrokenLinks) > 0 {
			fmt.Printf("\nBroken links: %d\n", len(result.BrokenLinks))
			for _, issue := range result.BrokenLinks {
				fmt.Printf("  %s:%d: %s\n", issue.File, issue.Line, issue.Message)
			}
		}
		if len(result.MissingAnchors) > 0 {
			fmt.Printf("\nMissing anchors: %d\n", len(result.MissingAnchors))
			for _, issue := range result.MissingAnchors {
				fmt.Printf("  %s:%d: %s\n", issue.File, issue.Line, issue.Message)
			}
		}
		if len(result.BrokenLinks) == 0 && len(result.MissingAnchors) == 0 {
			fmt.Println("  All links valid.")
		}
	}

	// Summary
	totalWarnings := len(warnings) + len(convWarnings)
	fmt.Printf("\nConversion complete: %d files generated", len(files))
	if totalWarnings > 0 {
		fmt.Printf(", %d warnings", totalWarnings)
	}
	fmt.Println()
}
