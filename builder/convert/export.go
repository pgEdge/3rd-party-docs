//-------------------------------------------------------------------------
//
// pgEdge PostgreSQL Docs
//
// Copyright (c) 2026, pgEdge, Inc.
// This software is released under The PostgreSQL License
//
//-------------------------------------------------------------------------

package convert

// ExportSlugify exposes slugify for use by other packages.
func ExportSlugify(s string) string {
	return slugify(s)
}
