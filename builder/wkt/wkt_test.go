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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseColor(t *testing.T) {
	tests := []struct {
		input   string
		color   string
		opacity float64
	}{
		{"#6495ED", "#6495ED", 1.0},
		{"#C0D0F080", "#C0D0F0", 0.502},
		{"#ff000020", "#ff0000", 0.125},
		{"Grey", "Grey", 1.0},
		{"red", "red", 1.0},
	}
	for _, tt := range tests {
		c, o := ParseColor(tt.input)
		if c != tt.color {
			t.Errorf("ParseColor(%q): color = %q, want %q", tt.input, c, tt.color)
		}
		if o < tt.opacity-0.01 || o > tt.opacity+0.01 {
			t.Errorf("ParseColor(%q): opacity = %.3f, want %.3f", tt.input, o, tt.opacity)
		}
	}
}

func TestParseWKTPoint(t *testing.T) {
	g, err := ParseWKT("POINT(100 90)")
	if err != nil {
		t.Fatal(err)
	}
	if g.Type != GeomPoint {
		t.Fatalf("type = %s, want POINT", g.Type)
	}
	if len(g.Points) != 1 || g.Points[0].X != 100 || g.Points[0].Y != 90 {
		t.Errorf("points = %v", g.Points)
	}
}

func TestParseWKTPointWithSpaces(t *testing.T) {
	g, err := ParseWKT("POINT (160 40)")
	if err != nil {
		t.Fatal(err)
	}
	if g.Type != GeomPoint || g.Points[0].X != 160 {
		t.Errorf("unexpected: %+v", g)
	}
}

func TestParseWKTLineString(t *testing.T) {
	g, err := ParseWKT("LINESTRING(10 30, 50 50, 30 110)")
	if err != nil {
		t.Fatal(err)
	}
	if g.Type != GeomLineString || len(g.Points) != 3 {
		t.Errorf("unexpected: type=%s, points=%d", g.Type, len(g.Points))
	}
}

func TestParseWKTPolygon(t *testing.T) {
	g, err := ParseWKT("POLYGON((0 0, 100 0, 100 100, 0 100, 0 0))")
	if err != nil {
		t.Fatal(err)
	}
	if g.Type != GeomPolygon || len(g.Rings) != 1 || len(g.Rings[0]) != 5 {
		t.Errorf("unexpected: type=%s, rings=%d", g.Type, len(g.Rings))
	}
}

func TestParseWKTMultiPoint(t *testing.T) {
	g, err := ParseWKT("MULTIPOINT ( 60 80, 190 10 )")
	if err != nil {
		t.Fatal(err)
	}
	if g.Type != GeomMultiPoint || len(g.Geoms) != 2 {
		t.Errorf("unexpected: type=%s, geoms=%d", g.Type, len(g.Geoms))
	}
}

func TestParseWKTMultiLineString(t *testing.T) {
	g, err := ParseWKT("MULTILINESTRING((10 160, 60 120), (120 140, 60 120))")
	if err != nil {
		t.Fatal(err)
	}
	if g.Type != GeomMultiLineString || len(g.Geoms) != 2 {
		t.Errorf("unexpected: type=%s, geoms=%d", g.Type, len(g.Geoms))
	}
}

func TestParseStyles(t *testing.T) {
	dir := t.TempDir()
	conf := filepath.Join(dir, "styles.conf")
	err := os.WriteFile(conf, []byte(`
[Style]
styleName = Default
pointSize = 6
pointColor = Grey
lineWidth = 7
lineColor = Grey

[Style]
styleName = ArgA
pointSize = 5
pointColor = "#6495ED"
lineWidth = 4
lineColor = "#6495ED"
polygonFillColor = "#C0D0F080"
polygonStrokeColor = "#6495ED"
polygonStrokeWidth = 2
`), 0644)
	if err != nil {
		t.Fatal(err)
	}

	ss, err := ParseStyles(conf)
	if err != nil {
		t.Fatal(err)
	}

	s := ss.Get("ArgA")
	if s.Name != "ArgA" {
		t.Errorf("name = %q", s.Name)
	}
	if s.PointColor != "#6495ED" {
		t.Errorf("pointColor = %q", s.PointColor)
	}
	if s.PolygonStrokeWidth != 2 {
		t.Errorf("polygonStrokeWidth = %d", s.PolygonStrokeWidth)
	}

	d := ss.Get("Unknown")
	if d.Name != "Default" {
		t.Errorf("fallback name = %q", d.Name)
	}
}

func TestRenderSVGPoint(t *testing.T) {
	g := &Geometry{Type: GeomPoint, Points: []Point2D{{100, 90}}}
	s := &Style{
		PointSize:  5,
		PointColor: "#6495ED",
	}
	renderer := NewSVGRenderer(200, 200)
	svg := renderer.Render([]Layer{{Geom: g, Style: s}})

	if !strings.Contains(svg, "<svg") {
		t.Error("missing <svg> tag")
	}
	if !strings.Contains(svg, "circle") {
		t.Error("missing circle for point")
	}
	// Y should be flipped: 200 - 90 = 110
	if !strings.Contains(svg, `cy="110.00"`) {
		t.Errorf("Y not flipped correctly in: %s", svg)
	}
}

func TestRenderSVGLineString(t *testing.T) {
	g := &Geometry{
		Type:   GeomLineString,
		Points: []Point2D{{10, 30}, {50, 50}, {30, 110}},
	}
	s := &Style{
		LineWidth: 4,
		LineColor: "#6495ED",
	}
	renderer := NewSVGRenderer(200, 200)
	svg := renderer.Render([]Layer{{Geom: g, Style: s}})

	if !strings.Contains(svg, "polyline") {
		t.Error("missing polyline")
	}
}

func TestRenderSVGPolygon(t *testing.T) {
	g := &Geometry{
		Type:  GeomPolygon,
		Rings: [][]Point2D{{{0, 0}, {100, 0}, {100, 100}, {0, 100}, {0, 0}}},
	}
	s := &Style{
		PolygonFillColor:   "#C0D0F080",
		PolygonStrokeColor: "#6495ED",
		PolygonStrokeWidth: 2,
	}
	renderer := NewSVGRenderer(200, 200)
	svg := renderer.Render([]Layer{{Geom: g, Style: s}})

	if !strings.Contains(svg, "<path") {
		t.Error("missing path for polygon")
	}
	if !strings.Contains(svg, `fill-opacity=`) {
		t.Error("missing fill-opacity for semi-transparent fill")
	}
}

func TestRenderSVGArrow(t *testing.T) {
	g := &Geometry{
		Type:   GeomLineString,
		Points: []Point2D{{10, 160}, {60, 120}},
	}
	s := &Style{
		LineWidth:     4,
		LineColor:     "#6495ED",
		LineArrowSize: 10,
	}
	renderer := NewSVGRenderer(200, 200)
	svg := renderer.Render([]Layer{{Geom: g, Style: s}})

	if !strings.Contains(svg, "polygon") {
		t.Error("missing polygon for arrow")
	}
}

func TestRenderFile(t *testing.T) {
	dir := t.TempDir()

	// Write styles.conf
	err := os.WriteFile(filepath.Join(dir, "styles.conf"), []byte(`
[Style]
styleName = Default
pointSize = 5
pointColor = Grey
lineWidth = 5
lineColor = Grey

[Style]
styleName = ArgA
pointSize = 5
pointColor = "#6495ED"
lineWidth = 4
lineColor = "#6495ED"
polygonFillColor = "#C0D0F080"
polygonStrokeColor = "#6495ED"
polygonStrokeWidth = 2
`), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Write a WKT file
	err = os.WriteFile(filepath.Join(dir, "test.wkt"), []byte(
		"ArgA;POINT(100 90)\n"+
			"LINESTRING(10 30, 50 50, 30 110)\n",
	), 0644)
	if err != nil {
		t.Fatal(err)
	}

	svg, err := RenderFile(filepath.Join(dir, "test.wkt"), 200, 200)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(svg, "<svg") {
		t.Error("missing svg tag")
	}
	if !strings.Contains(svg, "circle") {
		t.Error("missing point")
	}
	if !strings.Contains(svg, "polyline") {
		t.Error("missing linestring")
	}
}
