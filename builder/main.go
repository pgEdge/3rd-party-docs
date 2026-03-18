//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// pgdoc-converter converts PostgreSQL DocBook SGML documentation,
// pgAdmin/PostgREST RST documentation, or upstream Markdown
// documentation to Markdown for MkDocs Material.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/pgEdge/postgresql-docs/builder/backrest"
	"github.com/pgEdge/postgresql-docs/builder/convert"
	"github.com/pgEdge/postgresql-docs/builder/md"
	"github.com/pgEdge/postgresql-docs/builder/nav"
	"github.com/pgEdge/postgresql-docs/builder/rst"
	"github.com/pgEdge/postgresql-docs/builder/sgml"
	"github.com/pgEdge/postgresql-docs/builder/validate"
)

func main() {
	mode := flag.String("mode", "sgml", "Conversion mode: sgml, xml, rst, md, or backrest")
	srcDir := flag.String("src", "", "Path to source documentation directory")
	outDir := flag.String("out", "./docs", "Output directory for .md files")
	mkdocsFile := flag.String("mkdocs", "./mkdocs.yml", "Path to mkdocs.yml")
	version := flag.String("version", "", "Version label (e.g., 17.2 or 9.13)")
	copyright := flag.String("copyright", "", "Copyright string (RST mode)")
	pgadminSrc := flag.String("pgadmin-src", "", "Path to pgAdmin source (for literalinclude)")
	skipSections := flag.String("skip-sections", "", "Comma-separated section headings to suppress")
	entryFile := flag.String("entry-file", "", "Entry file name (xml mode, default: auto-detect)")
	sitePrefix := flag.String("site-prefix", "", "Site name prefix (e.g., PostGIS)")
	doValidate := flag.Bool("validate", false, "Run link validation after conversion")
	verbose := flag.Bool("verbose", false, "Verbose output")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pgdoc-converter [flags]\n\n")
		fmt.Fprintf(os.Stderr, "Converts PostgreSQL SGML or pgAdmin RST docs to Markdown.\n\n")
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
		fmt.Printf("Mode: %s\n", *mode)
		fmt.Printf("Source: %s\n", *srcDir)
		fmt.Printf("Output: %s\n", *outDir)
		fmt.Printf("Version: %s\n", *version)
	}

	switch *mode {
	case "sgml":
		runSGML(*srcDir, *outDir, *mkdocsFile, *version, *doValidate, *verbose)
	case "xml":
		runXML(*srcDir, *outDir, *mkdocsFile, *version, *entryFile,
			*sitePrefix, *skipSections, *doValidate, *verbose)
	case "rst":
		runRST(*srcDir, *outDir, *mkdocsFile, *version, *copyright,
			*pgadminSrc, *skipSections, *doValidate, *verbose)
	case "md":
		runMD(*srcDir, *outDir, *mkdocsFile, *version,
			*doValidate, *verbose)
	case "backrest":
		runBackRest(*srcDir, *outDir, *mkdocsFile, *version,
			*doValidate, *verbose)
	default:
		fmt.Fprintf(os.Stderr, "error: unknown mode %q (use sgml, xml, rst, md, or backrest)\n", *mode)
		os.Exit(1)
	}
}

// runSGML runs the SGML conversion pipeline.
func runSGML(srcDir, outDir, mkdocsFile, version string, doValidate, verbose bool) {
	// Phase 0: Generate missing SGML files
	if verbose {
		fmt.Println("Phase 0: Generating missing SGML files...")
	}
	genCount, genWarnings := sgml.GenerateMissingFiles(srcDir, version)
	if verbose {
		fmt.Printf("  Generated %d files\n", genCount)
		for _, w := range genWarnings {
			fmt.Printf("  Warning: %s\n", w)
		}
	}
	defer sgml.CleanGeneratedFiles(srcDir)

	// Phase 1: Entity resolution
	if verbose {
		fmt.Println("Phase 1: Resolving entities...")
	}
	resolver := sgml.NewEntityResolver(srcDir)
	body, err := resolver.ResolveFile("postgres.sgml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving entities: %v\n", err)
		os.Exit(1)
	}

	if verbose {
		fmt.Printf("  Resolved %d entities\n", resolver.EntityCount())
		fmt.Printf("  Document body: %d bytes\n", len(body))
		for _, w := range resolver.Warnings() {
			fmt.Printf("  Warning: %s\n", w)
		}
	}

	// Phase 2: Parse SGML
	if verbose {
		fmt.Println("Phase 2: Parsing SGML...")
	}
	root, warnings, err := sgml.ParseString(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing SGML: %v\n", err)
		os.Exit(1)
	}
	if verbose && len(warnings) > 0 {
		fmt.Printf("  Parser warnings: %d\n", len(warnings))
		for _, w := range warnings {
			fmt.Printf("    %s\n", w)
		}
	}

	// Phases 3-4: Convert
	if verbose {
		fmt.Println("Phase 3: Building ID map...")
		fmt.Println("Phase 4: Converting to Markdown...")
	}
	converter := convert.NewConverter(root, srcDir, outDir, version)
	if err := converter.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "error converting: %v\n", err)
		os.Exit(1)
	}

	convWarnings := converter.Warnings()
	if verbose && len(convWarnings) > 0 {
		fmt.Printf("  Conversion warnings: %d\n", len(convWarnings))
		for _, w := range convWarnings {
			fmt.Printf("    %s\n", w)
		}
	}

	files := converter.Files()
	if verbose {
		fmt.Printf("  Generated %d files\n", len(files))
	}

	// Phase 5: Nav generation
	if verbose {
		fmt.Println("Phase 5: Generating nav...")
	}
	navRoot := nav.BuildNav(files)
	navYAML := nav.GenerateYAML(navRoot)

	if mkdocsFile != "" {
		if _, err := os.Stat(mkdocsFile); err == nil {
			siteName := ""
			if version != "" {
				siteName = "PostgreSQL " + version
			}
			if err := nav.UpdateMkdocsYML(mkdocsFile, navYAML, siteName); err != nil {
				fmt.Fprintf(os.Stderr, "error updating mkdocs.yml: %v\n", err)
				os.Exit(1)
			}
			if verbose {
				fmt.Printf("  Updated %s\n", mkdocsFile)
			}
		} else if verbose {
			fmt.Printf("  %s not found, skipping nav update\n", mkdocsFile)
		}
	}

	// Phase 6: Validation
	runValidation(doValidate, verbose, outDir, warnings, convWarnings, len(files))
}

// runXML runs the DocBook XML conversion pipeline, reusing the
// SGML entity resolver, parser, and converter with DocBook 5
// normalization (xml:id → id, refsection → refsect1, etc.).
func runXML(
	srcDir, outDir, mkdocsFile, version, entryFile,
	sitePrefix, skipSections string,
	doValidate, verbose bool,
) {
	// Auto-detect entry file if not specified
	if entryFile == "" {
		// Look for common DocBook XML entry files
		for _, candidate := range []string{
			"postgis.xml", "index.xml", "book.xml",
		} {
			path := srcDir + "/" + candidate
			if _, err := os.Stat(path); err == nil {
				entryFile = candidate
				break
			}
		}
		if entryFile == "" {
			fmt.Fprintln(os.Stderr,
				"error: no entry file found; use -entry-file")
			os.Exit(1)
		}
	}

	if verbose {
		fmt.Printf("Converting DocBook XML (%s)...\n", entryFile)
	}

	// Phase 1: Entity resolution
	if verbose {
		fmt.Println("Phase 1: Resolving entities...")
	}
	resolver := sgml.NewEntityResolver(srcDir)
	body, err := resolver.ResolveFile(entryFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error resolving entities: %v\n", err)
		os.Exit(1)
	}

	// Replace build-time placeholders
	if version != "" {
		body = strings.ReplaceAll(body,
			"@@LAST_RELEASE_VERSION@@", version)
	}
	body = strings.ReplaceAll(body,
		"@@POSTGIS_DOWNLOAD_URL@@",
		"https://download.osgeo.org/postgis/source")

	if verbose {
		fmt.Printf("  Resolved %d entities\n", resolver.EntityCount())
		fmt.Printf("  Document body: %d bytes\n", len(body))
		for _, w := range resolver.Warnings() {
			fmt.Printf("  Warning: %s\n", w)
		}
	}

	// Phase 2: Parse (reuses SGML parser with XML normalization)
	if verbose {
		fmt.Println("Phase 2: Parsing XML...")
	}

	// Parse skip-sections
	var skip []string
	if skipSections != "" {
		for _, s := range strings.Split(skipSections, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				skip = append(skip, s)
			}
		}
	}

	root, warnings, err := sgml.ParseString(body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing XML: %v\n", err)
		os.Exit(1)
	}
	if verbose && len(warnings) > 0 {
		fmt.Printf("  Parser warnings: %d\n", len(warnings))
		for _, w := range warnings {
			fmt.Printf("    %s\n", w)
		}
	}

	// Remove skipped sections from the tree
	if len(skip) > 0 {
		sgml.RemoveSections(root, skip)
		if verbose {
			fmt.Printf("  Skipping sections: %v\n", skip)
		}
	}

	// Phases 3-4: Convert
	if verbose {
		fmt.Println("Phase 3: Building ID map...")
		fmt.Println("Phase 4: Converting to Markdown...")
	}
	converter := convert.NewConverter(root, srcDir, outDir, version)
	if err := converter.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "error converting: %v\n", err)
		os.Exit(1)
	}

	convWarnings := converter.Warnings()
	if verbose && len(convWarnings) > 0 {
		fmt.Printf("  Conversion warnings: %d\n", len(convWarnings))
		for _, w := range convWarnings {
			fmt.Printf("    %s\n", w)
		}
	}

	files := converter.Files()
	if verbose {
		fmt.Printf("  Generated %d files\n", len(files))
	}

	// Phase 5: Nav generation
	if verbose {
		fmt.Println("Phase 5: Generating nav...")
	}
	navRoot := nav.BuildNav(files)
	navYAML := nav.GenerateYAML(navRoot)

	if mkdocsFile != "" {
		if _, err := os.Stat(mkdocsFile); err == nil {
			siteName := ""
			if sitePrefix != "" {
				siteName = sitePrefix
				if version != "" {
					siteName += " " + version
				}
			} else if version != "" {
				siteName = version
			}
			if err := nav.UpdateMkdocsYML(
				mkdocsFile, navYAML, siteName); err != nil {
				fmt.Fprintf(os.Stderr,
					"error updating mkdocs.yml: %v\n", err)
				os.Exit(1)
			}
			if verbose {
				fmt.Printf("  Updated %s\n", mkdocsFile)
			}
		} else if verbose {
			fmt.Printf("  %s not found, skipping nav update\n",
				mkdocsFile)
		}
	}

	// Phase 6: Validation
	runValidation(doValidate, verbose, outDir,
		warnings, convWarnings, len(files))
}

// runRST runs the RST conversion pipeline.
func runRST(
	srcDir, outDir, mkdocsFile, version, copyright, pgadminSrc string,
	skipSections string,
	doValidate, verbose bool,
) {
	if verbose {
		fmt.Println("Converting RST documentation...")
	}

	// Parse comma-separated skip sections
	var skip []string
	if skipSections != "" {
		for _, s := range strings.Split(skipSections, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				skip = append(skip, s)
			}
		}
		if verbose && len(skip) > 0 {
			fmt.Printf("  Skipping sections: %v\n", skip)
		}
	}

	converter := rst.NewConverter(
		srcDir, outDir, version, copyright, pgadminSrc, skip, verbose)

	if err := converter.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "error converting RST: %v\n", err)
		os.Exit(1)
	}

	convWarnings := converter.Warnings()
	if verbose && len(convWarnings) > 0 {
		fmt.Printf("  Conversion warnings: %d\n", len(convWarnings))
		for _, w := range convWarnings {
			fmt.Printf("    %s\n", w)
		}
	}

	files := converter.Files()
	if verbose {
		fmt.Printf("  Generated %d files\n", len(files))
	}

	// Nav generation
	if verbose {
		fmt.Println("Generating nav...")
	}
	navRoot := nav.BuildNav(files)
	navYAML := nav.GenerateYAML(navRoot)

	if mkdocsFile != "" {
		if _, err := os.Stat(mkdocsFile); err == nil {
			// Build site name from conf.py project name + version
			siteName := ""
			if project := converter.ProjectName(); project != "" {
				siteName = project
				if version != "" {
					siteName += " " + version
				}
			} else if version != "" {
				siteName = version
			}
			if err := nav.UpdateMkdocsYML(mkdocsFile, navYAML, siteName); err != nil {
				fmt.Fprintf(os.Stderr, "error updating mkdocs.yml: %v\n", err)
				os.Exit(1)
			}
			if verbose {
				fmt.Printf("  Updated %s\n", mkdocsFile)
			}
		} else if verbose {
			fmt.Printf("  %s not found, skipping nav update\n", mkdocsFile)
		}
	}

	// Validation
	runValidation(doValidate, verbose, outDir, nil, convWarnings, len(files))
}

// runMD runs the Markdown conversion pipeline.
func runMD(
	srcDir, outDir, mkdocsFile, version string,
	doValidate, verbose bool,
) {
	if verbose {
		fmt.Println("Converting Markdown documentation...")
	}

	converter := md.NewConverter(srcDir, outDir, version, verbose)

	if err := converter.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "error converting MD: %v\n", err)
		os.Exit(1)
	}

	convWarnings := converter.Warnings()
	if verbose && len(convWarnings) > 0 {
		fmt.Printf("  Conversion warnings: %d\n",
			len(convWarnings))
		for _, w := range convWarnings {
			fmt.Printf("    %s\n", w)
		}
	}

	files := converter.Files()
	if verbose {
		fmt.Printf("  Generated %d files\n", len(files))
	}

	// Nav generation
	if verbose {
		fmt.Println("Generating nav...")
	}
	navRoot := nav.BuildNav(files)
	navYAML := nav.GenerateYAML(navRoot)

	if mkdocsFile != "" {
		if _, err := os.Stat(mkdocsFile); err == nil {
			siteName := version
			if err := nav.UpdateMkdocsYML(
				mkdocsFile, navYAML, siteName); err != nil {
				fmt.Fprintf(os.Stderr,
					"error updating mkdocs.yml: %v\n", err)
				os.Exit(1)
			}
			if verbose {
				fmt.Printf("  Updated %s\n", mkdocsFile)
			}
		} else if verbose {
			fmt.Printf("  %s not found, skipping nav update\n",
				mkdocsFile)
		}
	}

	// Validation
	runValidation(doValidate, verbose, outDir, nil,
		convWarnings, len(files))
}

// runBackRest runs the pgBackRest custom XML conversion pipeline.
func runBackRest(
	srcDir, outDir, mkdocsFile, version string,
	doValidate, verbose bool,
) {
	if verbose {
		fmt.Println("Converting pgBackRest documentation...")
	}

	converter := backrest.NewConverter(srcDir, outDir, version, verbose)

	if err := converter.Convert(); err != nil {
		fmt.Fprintf(os.Stderr, "error converting pgBackRest: %v\n", err)
		os.Exit(1)
	}

	convWarnings := converter.Warnings()
	if verbose && len(convWarnings) > 0 {
		fmt.Printf("  Conversion warnings: %d\n", len(convWarnings))
		for _, w := range convWarnings {
			fmt.Printf("    %s\n", w)
		}
	}

	files := converter.Files()
	if verbose {
		fmt.Printf("  Generated %d files\n", len(files))
	}

	// Nav generation
	if verbose {
		fmt.Println("Generating nav...")
	}
	navRoot := nav.BuildNav(files)
	navYAML := nav.GenerateYAML(navRoot)

	if mkdocsFile != "" {
		if _, err := os.Stat(mkdocsFile); err == nil {
			siteName := "pgBackRest"
			if version != "" {
				siteName += " " + version
			}
			if err := nav.UpdateMkdocsYML(
				mkdocsFile, navYAML, siteName); err != nil {
				fmt.Fprintf(os.Stderr,
					"error updating mkdocs.yml: %v\n", err)
				os.Exit(1)
			}
			if verbose {
				fmt.Printf("  Updated %s\n", mkdocsFile)
			}
		} else if verbose {
			fmt.Printf("  %s not found, skipping nav update\n",
				mkdocsFile)
		}
	}

	// Validation
	runValidation(doValidate, verbose, outDir, nil,
		convWarnings, len(files))
}

// runValidation runs link validation if requested.
func runValidation(
	doValidate, verbose bool,
	outDir string,
	parseWarnings []string,
	convWarnings []string,
	fileCount int,
) {
	if doValidate {
		if verbose {
			fmt.Println("Validating...")
		}
		result, err := validate.ValidateDir(outDir)
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
	totalWarnings := len(parseWarnings) + len(convWarnings)
	fmt.Printf("\nConversion complete: %d files generated", fileCount)
	if totalWarnings > 0 {
		fmt.Printf(", %d warnings", totalWarnings)
	}
	fmt.Println()
}
