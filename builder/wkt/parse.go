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
	"strconv"
	"strings"
	"unicode"
)

// Point2D is a 2D coordinate.
type Point2D struct {
	X, Y float64
}

// Geometry types
const (
	GeomPoint              = "POINT"
	GeomLineString         = "LINESTRING"
	GeomPolygon            = "POLYGON"
	GeomMultiPoint         = "MULTIPOINT"
	GeomMultiLineString    = "MULTILINESTRING"
	GeomMultiPolygon       = "MULTIPOLYGON"
	GeomGeometryCollection = "GEOMETRYCOLLECTION"
)

// Geometry represents a parsed WKT geometry.
type Geometry struct {
	Type   string
	Points []Point2D   // for POINT, LINESTRING
	Rings  [][]Point2D // for POLYGON (outer + holes)
	Geoms  []*Geometry // for MULTI* and COLLECTION
}

// parser holds state for recursive descent WKT parsing.
type parser struct {
	s    string
	pos  int
	dimZ bool // has Z dimension
	dimM bool // has M dimension
}

// ParseWKT parses a WKT geometry string.
func ParseWKT(wkt string) (*Geometry, error) {
	p := &parser{s: strings.TrimSpace(wkt), pos: 0}
	g, err := p.parseGeometry()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (p *parser) parseGeometry() (*Geometry, error) {
	p.skipWS()
	typ := p.readType()
	p.skipWS()

	// Strip Z/M/ZM dimension suffixes
	upper := strings.ToUpper(typ)
	if strings.HasSuffix(upper, " ZM") {
		upper = strings.TrimSuffix(upper, " ZM")
		p.dimZ, p.dimM = true, true
	} else if strings.HasSuffix(upper, " Z") {
		upper = strings.TrimSuffix(upper, " Z")
		p.dimZ = true
	} else if strings.HasSuffix(upper, " M") {
		upper = strings.TrimSuffix(upper, " M")
		p.dimM = true
	}

	switch upper {
	case GeomPoint:
		return p.parsePoint()
	case GeomLineString:
		return p.parseLineString()
	case GeomPolygon:
		return p.parsePolygon()
	case GeomMultiPoint:
		return p.parseMultiPoint()
	case GeomMultiLineString:
		return p.parseMultiLineString()
	case GeomMultiPolygon:
		return p.parseMultiPolygon()
	case GeomGeometryCollection:
		return p.parseGeometryCollection()
	default:
		return nil, fmt.Errorf("unknown geometry type: %q at pos %d", typ, p.pos)
	}
}

func (p *parser) readType() string {
	start := p.pos
	for p.pos < len(p.s) && unicode.IsLetter(rune(p.s[p.pos])) {
		p.pos++
	}
	word := p.s[start:p.pos]

	// Check for Z/M/ZM suffix after a space (e.g. "MULTIPOLYGON Z")
	saved := p.pos
	p.skipWS()
	if p.pos < len(p.s) && unicode.IsLetter(rune(p.s[p.pos])) {
		sufStart := p.pos
		for p.pos < len(p.s) && unicode.IsLetter(rune(p.s[p.pos])) {
			p.pos++
		}
		suf := strings.ToUpper(p.s[sufStart:p.pos])
		if suf == "Z" || suf == "M" || suf == "ZM" {
			return word + " " + suf
		}
		// Not a dimension suffix — revert
		p.pos = saved
	} else {
		p.pos = saved
	}
	return word
}

func (p *parser) skipWS() {
	for p.pos < len(p.s) && (p.s[p.pos] == ' ' || p.s[p.pos] == '\t' || p.s[p.pos] == '\n' || p.s[p.pos] == '\r') {
		p.pos++
	}
}

func (p *parser) expect(ch byte) error {
	p.skipWS()
	if p.pos >= len(p.s) || p.s[p.pos] != ch {
		if p.pos >= len(p.s) {
			return fmt.Errorf("expected '%c' but got EOF", ch)
		}
		return fmt.Errorf("expected '%c' but got '%c' at pos %d", ch, p.s[p.pos], p.pos)
	}
	p.pos++
	return nil
}

func (p *parser) peek() byte {
	p.skipWS()
	if p.pos >= len(p.s) {
		return 0
	}
	return p.s[p.pos]
}

func (p *parser) readNumber() (float64, error) {
	p.skipWS()
	start := p.pos
	if p.pos < len(p.s) && (p.s[p.pos] == '-' || p.s[p.pos] == '+') {
		p.pos++
	}
	for p.pos < len(p.s) && (p.s[p.pos] >= '0' && p.s[p.pos] <= '9') {
		p.pos++
	}
	if p.pos < len(p.s) && p.s[p.pos] == '.' {
		p.pos++
		for p.pos < len(p.s) && (p.s[p.pos] >= '0' && p.s[p.pos] <= '9') {
			p.pos++
		}
	}
	// Handle scientific notation
	if p.pos < len(p.s) && (p.s[p.pos] == 'e' || p.s[p.pos] == 'E') {
		p.pos++
		if p.pos < len(p.s) && (p.s[p.pos] == '-' || p.s[p.pos] == '+') {
			p.pos++
		}
		for p.pos < len(p.s) && (p.s[p.pos] >= '0' && p.s[p.pos] <= '9') {
			p.pos++
		}
	}
	if start == p.pos {
		return 0, fmt.Errorf("expected number at pos %d", p.pos)
	}
	return strconv.ParseFloat(p.s[start:p.pos], 64)
}

func (p *parser) readCoord() (Point2D, error) {
	x, err := p.readNumber()
	if err != nil {
		return Point2D{}, err
	}
	y, err := p.readNumber()
	if err != nil {
		return Point2D{}, err
	}
	// Skip Z and/or M ordinates if present
	extra := 0
	if p.dimZ {
		extra++
	}
	if p.dimM {
		extra++
	}
	for i := 0; i < extra; i++ {
		p.skipWS()
		// Only consume if next char looks like a number
		if p.pos < len(p.s) && isNumberStart(p.s[p.pos]) {
			if _, err := p.readNumber(); err != nil {
				return Point2D{}, err
			}
		}
	}
	return Point2D{X: x, Y: y}, nil
}

func isNumberStart(ch byte) bool {
	return (ch >= '0' && ch <= '9') || ch == '-' || ch == '+' || ch == '.'
}

func (p *parser) readCoordList() ([]Point2D, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var pts []Point2D
	for {
		pt, err := p.readCoord()
		if err != nil {
			return nil, err
		}
		pts = append(pts, pt)
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return pts, nil
}

func (p *parser) parsePoint() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	pt, err := p.readCoord()
	if err != nil {
		return nil, err
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomPoint, Points: []Point2D{pt}}, nil
}

func (p *parser) parseLineString() (*Geometry, error) {
	pts, err := p.readCoordList()
	if err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomLineString, Points: pts}, nil
}

func (p *parser) parsePolygon() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var rings [][]Point2D
	for {
		ring, err := p.readCoordList()
		if err != nil {
			return nil, err
		}
		rings = append(rings, ring)
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomPolygon, Rings: rings}, nil
}

func (p *parser) parseMultiPoint() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var geoms []*Geometry
	for {
		p.skipWS()
		// MULTIPOINT can have ((x y), (x y)) or (x y, x y)
		if p.peek() == '(' {
			p.pos++
			pt, err := p.readCoord()
			if err != nil {
				return nil, err
			}
			if err := p.expect(')'); err != nil {
				return nil, err
			}
			geoms = append(geoms, &Geometry{Type: GeomPoint, Points: []Point2D{pt}})
		} else {
			pt, err := p.readCoord()
			if err != nil {
				return nil, err
			}
			geoms = append(geoms, &Geometry{Type: GeomPoint, Points: []Point2D{pt}})
		}
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomMultiPoint, Geoms: geoms}, nil
}

func (p *parser) parseMultiLineString() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var geoms []*Geometry
	for {
		pts, err := p.readCoordList()
		if err != nil {
			return nil, err
		}
		geoms = append(geoms, &Geometry{Type: GeomLineString, Points: pts})
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomMultiLineString, Geoms: geoms}, nil
}

func (p *parser) parseMultiPolygon() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var geoms []*Geometry
	for {
		g, err := p.parsePolygonBody()
		if err != nil {
			return nil, err
		}
		geoms = append(geoms, g)
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomMultiPolygon, Geoms: geoms}, nil
}

func (p *parser) parsePolygonBody() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var rings [][]Point2D
	for {
		ring, err := p.readCoordList()
		if err != nil {
			return nil, err
		}
		rings = append(rings, ring)
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomPolygon, Rings: rings}, nil
}

func (p *parser) parseGeometryCollection() (*Geometry, error) {
	if err := p.expect('('); err != nil {
		return nil, err
	}
	var geoms []*Geometry
	for {
		g, err := p.parseGeometry()
		if err != nil {
			return nil, err
		}
		geoms = append(geoms, g)
		p.skipWS()
		if p.peek() == ',' {
			p.pos++
			continue
		}
		break
	}
	if err := p.expect(')'); err != nil {
		return nil, err
	}
	return &Geometry{Type: GeomGeometryCollection, Geoms: geoms}, nil
}
