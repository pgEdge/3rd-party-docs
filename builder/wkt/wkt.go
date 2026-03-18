//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

// Package wkt converts PostGIS WKT geometry files to SVG images.
// It reads .wkt source files and a styles.conf configuration file,
// producing SVG output without requiring PostGIS's C toolchain.
package wkt

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Layer pairs a geometry with its rendering style.
type Layer struct {
	Geom  *Geometry
	Style *Style
}

// RenderFile reads a .wkt file and its companion styles.conf,
// returning SVG content as a string. The dimensions default to
// 200x200 unless overridden.
func RenderFile(wktPath string, width, height int) (string, error) {
	if width <= 0 {
		width = defaultWidth
	}
	if height <= 0 {
		height = defaultHeight
	}

	// Load styles from the same directory as the WKT file
	dir := filepath.Dir(wktPath)
	stylesPath := filepath.Join(dir, "styles.conf")
	styles, err := ParseStyles(stylesPath)
	if err != nil {
		return "", fmt.Errorf("loading styles: %w", err)
	}

	// Parse the WKT file
	layers, err := ParseWKTFile(wktPath, styles)
	if err != nil {
		return "", fmt.Errorf("parsing WKT file: %w", err)
	}

	renderer := NewSVGRenderer(width, height)
	return renderer.Render(layers), nil
}

// RenderFileWithStyles reads a .wkt file using a pre-loaded StyleSet.
func RenderFileWithStyles(wktPath string, styles *StyleSet, width, height int) (string, error) {
	if width <= 0 {
		width = defaultWidth
	}
	if height <= 0 {
		height = defaultHeight
	}

	layers, err := ParseWKTFile(wktPath, styles)
	if err != nil {
		return "", fmt.Errorf("parsing WKT file: %w", err)
	}

	renderer := NewSVGRenderer(width, height)
	return renderer.Render(layers), nil
}

// ParseWKTFile reads a .wkt file and returns layers. Each non-blank
// line is "StyleName;WKT" or just "WKT" (using Default style).
func ParseWKTFile(wktPath string, styles *StyleSet) ([]Layer, error) {
	f, err := os.Open(wktPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var layers []Layer
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 256*1024), 256*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		styleName, wktStr := splitStyleLine(line)
		style := styles.Get(styleName)
		geom, err := ParseWKT(wktStr)
		if err != nil {
			return nil, fmt.Errorf("line %q: %w", line, err)
		}
		layers = append(layers, Layer{Geom: geom, Style: style})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return layers, nil
}

// splitStyleLine splits "StyleName;WKT..." into style name and WKT.
// If no semicolon, returns "Default" and the full line.
func splitStyleLine(line string) (string, string) {
	idx := strings.LastIndex(line, ";")
	if idx < 0 {
		return "Default", line
	}
	return line[:idx], line[idx+1:]
}
