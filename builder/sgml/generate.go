//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package sgml

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GenerateMissingFiles checks for files that the PG build system
// would normally generate and creates them if missing. It handles:
//   - version.sgml from version.sgml.in + provided version string
//   - Other generated files by running their Perl scripts
//   - Stub fallbacks when scripts can't be run
//
// pgVersion should be the full version string (e.g., "17.4" or "19devel").
// srcDir is the doc/src/sgml directory.
// Returns the number of files generated and any warnings.
func GenerateMissingFiles(srcDir, pgVersion string) (int, []string) {
	var warnings []string
	generated := 0

	// Resolve srcDir to absolute so paths work when cmd.Dir is set
	srcDir, _ = filepath.Abs(srcDir)

	// Determine PG source root (3 levels up from doc/src/sgml)
	pgRoot := filepath.Join(srcDir, "..", "..", "..")

	// Handle version.sgml — always regenerate since the version
	// changes between runs (e.g., switching PG branches).
	if pgVersion != "" {
		if err := generateVersionSGML(srcDir, pgVersion); err != nil {
			warnings = append(warnings,
				fmt.Sprintf("could not generate version.sgml: %v", err))
		} else {
			generated++
		}
	} else {
		versionFile := filepath.Join(srcDir, "version.sgml")
		if _, err := os.Stat(versionFile); os.IsNotExist(err) {
			if err := generateVersionSGML(srcDir, pgVersion); err != nil {
				warnings = append(warnings,
					fmt.Sprintf("could not generate version.sgml: %v", err))
			} else {
				generated++
			}
		}
	}

	// Each generated file has its own calling convention
	generators := []struct {
		name string
		gen  func(srcDir, pgRoot string) error
		stub string
	}{
		{
			name: "features-supported.sgml",
			gen:  func(s, r string) error { return genFeatures(s, r, "YES") },
			stub: "<tbody>\n</tbody>",
		},
		{
			name: "features-unsupported.sgml",
			gen:  func(s, r string) error { return genFeatures(s, r, "NO") },
			stub: "<tbody>\n</tbody>",
		},
		{
			name: "errcodes-table.sgml",
			gen:  genErrcodes,
			stub: "<!-- errcodes table not generated -->",
		},
		{
			name: "keywords-table.sgml",
			gen:  genKeywords,
			stub: "<!-- keywords table not generated -->",
		},
		{
			name: "wait_event_types.sgml",
			gen:  genWaitEvents,
			stub: "<!-- wait event types not generated -->",
		},
		{
			name: "targets-meson.sgml",
			gen:  genTargetsMeson,
			stub: "<!-- meson targets not generated -->",
		},
	}

	for _, g := range generators {
		outPath := filepath.Join(srcDir, g.name)
		if _, err := os.Stat(outPath); err == nil {
			continue // file already exists
		}

		if err := g.gen(srcDir, pgRoot); err != nil {
			warnings = append(warnings,
				fmt.Sprintf("could not generate %s: %v; using stub",
					g.name, err))
			if wErr := os.WriteFile(outPath, []byte(g.stub+"\n"), 0644); wErr != nil {
				warnings = append(warnings,
					fmt.Sprintf("could not write stub %s: %v",
						g.name, wErr))
			}
		}
		generated++
	}

	return generated, warnings
}

// genFeatures runs mk_feature_tables.pl YES|NO <packages> <features>
func genFeatures(srcDir, pgRoot, yesno string) error {
	script := filepath.Join(srcDir, "mk_feature_tables.pl")
	pkgs := filepath.Join(pgRoot, "src/backend/catalog/sql_feature_packages.txt")
	feats := filepath.Join(pgRoot, "src/backend/catalog/sql_features.txt")

	outName := "features-supported.sgml"
	if yesno == "NO" {
		outName = "features-unsupported.sgml"
	}

	return runPerlToFile(srcDir, outName, script, yesno, pkgs, feats)
}

// genErrcodes runs generate-errcodes-table.pl <errcodes.txt>
func genErrcodes(srcDir, pgRoot string) error {
	script := filepath.Join(srcDir, "generate-errcodes-table.pl")
	input := filepath.Join(pgRoot, "src/backend/utils/errcodes.txt")

	return runPerlToFile(srcDir, "errcodes-table.sgml", script, input)
}

// genKeywords runs generate-keywords-table.pl <srcdir>
// The script expects the SGML source dir as its sole argument,
// and finds kwlist.h and keywords/*.txt relative to it.
func genKeywords(srcDir, pgRoot string) error {
	script := filepath.Join(srcDir, "generate-keywords-table.pl")

	return runPerlToFile(srcDir, "keywords-table.sgml", script, srcDir)
}

// genWaitEvents runs generate-wait_event_types.pl --docs <input>
// with --outdir pointing to the SGML dir.
func genWaitEvents(srcDir, pgRoot string) error {
	script := filepath.Join(pgRoot,
		"src/backend/utils/activity/generate-wait_event_types.pl")
	input := filepath.Join(pgRoot,
		"src/backend/utils/activity/wait_event_names.txt")

	perlPath, err := exec.LookPath("perl")
	if err != nil {
		return fmt.Errorf("perl not found in PATH")
	}

	cmd := exec.Command(perlPath, script,
		"--docs", "--outdir", srcDir, input)
	cmd.Dir = srcDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("script failed: %s", string(output))
	}
	return nil
}

// genTargetsMeson runs generate-targets-meson.pl <input>
func genTargetsMeson(srcDir, pgRoot string) error {
	script := filepath.Join(srcDir, "generate-targets-meson.pl")
	input := filepath.Join(srcDir, "targets-meson.txt")

	return runPerlToFile(srcDir, "targets-meson.sgml", script, input)
}

// runPerlToFile runs a Perl script and writes its stdout to a file.
func runPerlToFile(srcDir, outName, script string, args ...string) error {
	perlPath, err := exec.LookPath("perl")
	if err != nil {
		return fmt.Errorf("perl not found in PATH")
	}

	if _, err := os.Stat(script); os.IsNotExist(err) {
		return fmt.Errorf("script %s not found", script)
	}

	cmdArgs := append([]string{script}, args...)
	cmd := exec.Command(perlPath, cmdArgs...)
	cmd.Dir = srcDir

	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("script failed: %s", string(exitErr.Stderr))
		}
		return fmt.Errorf("running script: %w", err)
	}

	outPath := filepath.Join(srcDir, outName)
	return os.WriteFile(outPath, output, 0644)
}

// generateVersionSGML creates version.sgml from version.sgml.in.
func generateVersionSGML(srcDir, pgVersion string) error {
	inPath := filepath.Join(srcDir, "version.sgml.in")
	outPath := filepath.Join(srcDir, "version.sgml")

	data, err := os.ReadFile(inPath)
	if err != nil {
		// No template — create directly from version string
		return writeVersionSGML(outPath, pgVersion)
	}

	content := string(data)
	majorVersion := extractMajorVersion(pgVersion)

	content = strings.ReplaceAll(content, "@PG_VERSION@", "\""+pgVersion+"\"")
	content = strings.ReplaceAll(content, "@PG_MAJORVERSION@", "\""+majorVersion+"\"")

	return os.WriteFile(outPath, []byte(content), 0644)
}

// writeVersionSGML creates a version.sgml directly.
func writeVersionSGML(outPath, pgVersion string) error {
	majorVersion := extractMajorVersion(pgVersion)
	content := fmt.Sprintf(
		"<!ENTITY version \"%s\">\n<!ENTITY majorversion \"%s\">\n",
		pgVersion, majorVersion)
	return os.WriteFile(outPath, []byte(content), 0644)
}

// extractMajorVersion extracts the major version from a PG version
// string. E.g., "17.4" → "17", "19devel" → "19".
func extractMajorVersion(version string) string {
	v := version
	for _, suffix := range []string{"devel", "beta1", "beta2", "beta3",
		"beta4", "rc1", "rc2", "alpha1", "alpha2", "alpha3"} {
		v = strings.TrimSuffix(v, suffix)
	}
	if idx := strings.Index(v, "."); idx >= 0 {
		return v[:idx]
	}
	return v
}

// CleanGeneratedFiles removes files that were generated by
// GenerateMissingFiles, so they don't pollute the PG source tree.
func CleanGeneratedFiles(srcDir string) {
	// Only clean version.sgml if we generated it (check for .in file)
	inPath := filepath.Join(srcDir, "version.sgml.in")
	if _, err := os.Stat(inPath); err == nil {
		os.Remove(filepath.Join(srcDir, "version.sgml"))
	}

	names := []string{
		"features-supported.sgml",
		"features-unsupported.sgml",
		"errcodes-table.sgml",
		"keywords-table.sgml",
		"wait_event_types.sgml",
		"targets-meson.sgml",
	}
	for _, name := range names {
		os.Remove(filepath.Join(srcDir, name))
	}
}
