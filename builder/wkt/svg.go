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
	"fmt"
	"math"
	"strings"
)

const (
	defaultWidth  = 200
	defaultHeight = 200
)

// SVGRenderer renders WKT geometries to SVG.
type SVGRenderer struct {
	width, height int
	buf           strings.Builder
}

// NewSVGRenderer creates a renderer with the given dimensions.
func NewSVGRenderer(width, height int) *SVGRenderer {
	return &SVGRenderer{width: width, height: height}
}

// svgY flips the Y axis (WKT is math-style, SVG is screen-style).
func (r *SVGRenderer) svgY(y float64) float64 {
	return float64(r.height) - y
}

// Render generates SVG for a set of layers (style+geometry pairs).
func (r *SVGRenderer) Render(layers []Layer) string {
	r.buf.Reset()
	r.buf.WriteString(fmt.Sprintf(
		`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`,
		r.width, r.height, r.width, r.height))
	r.buf.WriteByte('\n')

	// White background
	r.buf.WriteString(fmt.Sprintf(
		`<rect width="%d" height="%d" fill="white"/>`,
		r.width, r.height))
	r.buf.WriteByte('\n')

	for _, layer := range layers {
		r.renderGeometry(layer.Geom, layer.Style)
	}

	r.buf.WriteString("</svg>\n")
	return r.buf.String()
}

func (r *SVGRenderer) renderGeometry(g *Geometry, s *Style) {
	switch g.Type {
	case GeomPoint:
		r.renderPoint(g, s)
	case GeomLineString:
		r.renderLineString(g, s)
	case GeomPolygon:
		r.renderPolygon(g, s)
	case GeomMultiPoint, GeomMultiLineString, GeomMultiPolygon,
		GeomGeometryCollection:
		for _, sub := range g.Geoms {
			r.renderGeometry(sub, s)
		}
	}
}

func (r *SVGRenderer) renderPoint(g *Geometry, s *Style) {
	if len(g.Points) == 0 || s.PointSize <= 0 {
		return
	}
	pt := g.Points[0]
	color, opacity := ParseColor(s.PointColor)
	r.buf.WriteString(fmt.Sprintf(
		`<circle cx="%.2f" cy="%.2f" r="%d" fill="%s"`,
		pt.X, r.svgY(pt.Y), s.PointSize, color))
	if opacity < 0.999 {
		r.buf.WriteString(fmt.Sprintf(` fill-opacity="%.2f"`, opacity))
	}
	r.buf.WriteString("/>\n")
}

func (r *SVGRenderer) renderLineString(g *Geometry, s *Style) {
	if len(g.Points) < 2 {
		// Single point in a linestring — render as point
		if len(g.Points) == 1 {
			r.renderPoint(g, s)
		}
		return
	}

	color, opacity := ParseColor(s.LineColor)
	r.buf.WriteString(fmt.Sprintf(
		`<polyline points="%s" fill="none" stroke="%s" stroke-width="%d" stroke-linecap="round" stroke-linejoin="round"`,
		r.pointsStr(g.Points), color, s.LineWidth))
	if opacity < 0.999 {
		r.buf.WriteString(fmt.Sprintf(` stroke-opacity="%.2f"`, opacity))
	}
	r.buf.WriteString("/>\n")

	// Start point marker
	if s.LineStartSize > 0 {
		r.renderDot(g.Points[0], s.LineStartSize, color, opacity)
	}
	// End point marker
	if s.LineEndSize > 0 {
		r.renderDot(g.Points[len(g.Points)-1], s.LineEndSize, color, opacity)
	}
	// Arrow at end
	if s.LineArrowSize > 0 {
		r.renderArrow(g.Points, s.LineArrowSize, color, opacity)
	}
}

func (r *SVGRenderer) renderPolygon(g *Geometry, s *Style) {
	if len(g.Rings) == 0 {
		return
	}

	fillColor, fillOpacity := ParseColor(s.PolygonFillColor)
	strokeColor, strokeOpacity := ParseColor(s.PolygonStrokeColor)

	// Build SVG path with all rings
	var pathBuf strings.Builder
	for _, ring := range g.Rings {
		if len(ring) == 0 {
			continue
		}
		pathBuf.WriteString(fmt.Sprintf("M %.4f,%.4f",
			ring[0].X, r.svgY(ring[0].Y)))
		for _, pt := range ring[1:] {
			pathBuf.WriteString(fmt.Sprintf(" L %.4f,%.4f",
				pt.X, r.svgY(pt.Y)))
		}
		pathBuf.WriteString(" Z")
	}

	r.buf.WriteString(fmt.Sprintf(
		`<path d="%s" fill="%s"`, pathBuf.String(), fillColor))
	if fillOpacity < 0.999 {
		r.buf.WriteString(fmt.Sprintf(` fill-opacity="%.2f"`, fillOpacity))
	}
	r.buf.WriteString(fmt.Sprintf(
		` stroke="%s" stroke-width="%d"`, strokeColor, s.PolygonStrokeWidth))
	if strokeOpacity < 0.999 {
		r.buf.WriteString(fmt.Sprintf(` stroke-opacity="%.2f"`, strokeOpacity))
	}
	r.buf.WriteString("/>\n")
}

func (r *SVGRenderer) renderDot(pt Point2D, size int, color string, opacity float64) {
	r.buf.WriteString(fmt.Sprintf(
		`<circle cx="%.2f" cy="%.2f" r="%d" fill="%s"`,
		pt.X, r.svgY(pt.Y), size, color))
	if opacity < 0.999 {
		r.buf.WriteString(fmt.Sprintf(` fill-opacity="%.2f"`, opacity))
	}
	r.buf.WriteString("/>\n")
}

func (r *SVGRenderer) renderArrow(pts []Point2D, size int, color string, opacity float64) {
	if len(pts) < 2 {
		return
	}
	// Arrow at the end of the line, matching generator.c logic
	pn := pts[len(pts)-1]  // tip
	pn1 := pts[len(pts)-2] // second-to-last

	dx := pn1.X - pn.X
	dy := pn1.Y - pn.Y
	length := math.Sqrt(dx*dx + dy*dy)
	if length <= 0 {
		return
	}

	s := float64(size)
	offx := -0.5 * s * dy / length
	offy := 0.5 * s * dx / length

	p1x := pn.X + s*dx/length + offx
	p1y := pn.Y + s*dy/length + offy
	p2x := pn.X + s*dx/length - offx
	p2y := pn.Y + s*dy/length - offy

	r.buf.WriteString(fmt.Sprintf(
		`<polygon points="%.2f,%.2f %.2f,%.2f %.2f,%.2f" fill="%s" stroke="%s" stroke-width="2"`,
		pn.X, r.svgY(pn.Y),
		p1x, r.svgY(p1y),
		p2x, r.svgY(p2y),
		color, color))
	if opacity < 0.999 {
		r.buf.WriteString(fmt.Sprintf(` opacity="%.2f"`, opacity))
	}
	r.buf.WriteString("/>\n")
}

func (r *SVGRenderer) pointsStr(pts []Point2D) string {
	parts := make([]string, len(pts))
	for i, pt := range pts {
		parts[i] = fmt.Sprintf("%.2f,%.2f", pt.X, r.svgY(pt.Y))
	}
	return strings.Join(parts, " ")
}
