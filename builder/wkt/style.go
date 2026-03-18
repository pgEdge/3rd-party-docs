//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package wkt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Style defines the visual properties for rendering a WKT layer.
type Style struct {
	Name               string
	PointSize          int
	PointColor         string
	LineWidth          int
	LineColor          string
	LineStartSize      int
	LineEndSize        int
	LineArrowSize      int
	PolygonFillColor   string
	PolygonStrokeColor string
	PolygonStrokeWidth int
}

// StyleSet holds a collection of named styles.
type StyleSet struct {
	styles map[string]*Style
}

// Get returns a style by name, falling back to "Default".
func (ss *StyleSet) Get(name string) *Style {
	if s, ok := ss.styles[name]; ok {
		return s
	}
	if s, ok := ss.styles["Default"]; ok {
		return s
	}
	return &Style{
		Name:               "Default",
		PointSize:          5,
		PointColor:         "Grey",
		LineWidth:          5,
		LineColor:          "Grey",
		PolygonFillColor:   "Grey",
		PolygonStrokeColor: "Grey",
	}
}

// ParseStyles reads a styles.conf file and returns a StyleSet.
func ParseStyles(path string) (*StyleSet, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ss := &StyleSet{styles: make(map[string]*Style)}
	scanner := bufio.NewScanner(f)

	var cur *Style
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if line == "[Style]" {
			if cur != nil {
				ss.styles[cur.Name] = cur
			}
			cur = &Style{
				Name:               "Default",
				PointSize:          5,
				PointColor:         "Grey",
				LineWidth:          5,
				LineColor:          "Grey",
				PolygonFillColor:   "Grey",
				PolygonStrokeColor: "Grey",
			}
			continue
		}
		if cur == nil {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		// Strip surrounding quotes
		val = strings.Trim(val, "\"")

		switch key {
		case "styleName":
			cur.Name = val
		case "pointSize":
			cur.PointSize, _ = strconv.Atoi(val)
		case "pointColor":
			cur.PointColor = val
		case "lineWidth":
			cur.LineWidth, _ = strconv.Atoi(val)
		case "lineColor":
			cur.LineColor = val
		case "lineStartSize":
			cur.LineStartSize, _ = strconv.Atoi(val)
		case "lineEndSize":
			cur.LineEndSize, _ = strconv.Atoi(val)
		case "lineArrowSize":
			cur.LineArrowSize, _ = strconv.Atoi(val)
		case "polygonFillColor":
			cur.PolygonFillColor = val
		case "polygonStrokeColor":
			cur.PolygonStrokeColor = val
		case "polygonStrokeWidth":
			cur.PolygonStrokeWidth, _ = strconv.Atoi(val)
		}
	}
	if cur != nil {
		ss.styles[cur.Name] = cur
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ss, nil
}

// ParseColor converts a color string to SVG-compatible color and
// opacity. Handles 8-char hex (#RRGGBBAA), 6-char hex (#RRGGBB),
// and named colors.
func ParseColor(color string) (string, float64) {
	if strings.HasPrefix(color, "#") && len(color) == 9 {
		// #RRGGBBAA
		alpha, err := strconv.ParseUint(color[7:9], 16, 8)
		if err != nil {
			return color[:7], 1.0
		}
		return color[:7], float64(alpha) / 255.0
	}
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		return color, 1.0
	}
	// Named color — pass through as-is
	return color, 1.0
}

// FormatOpacity returns a string for opacity, omitting if 1.0.
func FormatOpacity(opacity float64) string {
	if opacity >= 0.999 {
		return ""
	}
	return fmt.Sprintf(" opacity=\"%.2f\"", opacity)
}
