# PostgreSQL & Component Documentation

[![CI](https://github.com/pgEdge/postgresql-docs/actions/workflows/ci.yml/badge.svg)](https://github.com/pgEdge/postgresql-docs/actions/workflows/ci.yml)

MkDocs Material documentation sites converted from upstream
sources:

- **PostgreSQL** — SGML/DocBook sources converted to Markdown
- **pgAdmin 4** — reStructuredText (RST) sources converted to
  Markdown
- **PostgREST** — reStructuredText (RST) sources converted to
  Markdown

## Branch Layout

The `main` branch contains only the converter tooling and a
skeleton MkDocs configuration. All generated documentation lives
on product/version branches:

| Branch | Product | Source |
|--------|---------|--------|
| `pg16` .. `pg19` | PostgreSQL 16–19 | SGML (upstream `doc/src/sgml/`) |
| `pgadmin911` .. `pgadmin913` | pgAdmin 4 v9.11–v9.13 | RST (`docs/en_US/`) |
| `pgadminmaster` | pgAdmin 4 dev | RST (upstream `master`) |
| `postgrest145` | PostgREST v14.5 | RST (`docs/`) |

## Prerequisites

- Go 1.25+
- [yq](https://github.com/mikefarah/yq) (for `build-all.sh`)
- Python 3 with
  [MkDocs Material](https://squidfunk.github.io/mkdocs-material/)

## Quick Start

### Build All Branches

The `build-all.sh` script automates the full pipeline:
cloning/fetching upstream repos, converting docs, and
committing to each branch. All branches are defined in
`branches.yml`.

```sh
# Build everything
./build-all.sh

# Build only PostgreSQL branches
./build-all.sh --branches "pg*"

# Build specific branches
./build-all.sh --branches pg17,postgrest145

# Preview what would be built
./build-all.sh --dry-run
```

After building, the script shows a summary and prompts to
push updated branches to the remote.

### Build a Single Branch Manually

For manual builds, checkout the target branch, provide the
upstream source, and run the converter:

```sh
# PostgreSQL (SGML mode)
make convert SRC_DIR=/path/to/postgresql/doc/src/sgml VERSION=17.2

# pgAdmin 4 (RST mode)
make convert-rst SRC_DIR=/path/to/pgadmin4/docs/en_US VERSION=9.13

# PostgREST (RST mode, suppressing Sponsors section)
make convert-rst SRC_DIR=/path/to/postgrest/docs VERSION=v14.5 \
    SKIP_SECTIONS="Sponsors"
```

Preview the site locally:

```sh
mkdocs serve
```

## Builder

The `builder/` directory contains a Go tool (`pgdoc-converter`)
that converts upstream documentation to Markdown suitable for
MkDocs Material. It supports two modes:

### SGML Mode (PostgreSQL)

- Entity resolution and SGML parsing
- DocBook-to-Markdown conversion (100+ element handlers)
- `func_table_entry` tables split into multi-column layout
- Two-pass conversion: ID map then content generation
- Image copying from the PostgreSQL source tree

### RST Mode (pgAdmin, PostgREST)

- Line-by-line RST parser (headings, directives, lists,
  grid tables, labels, substitutions, literal blocks)
- Toctree resolution for hierarchical nav structure
- Directive handlers: image, code-block, admonitions,
  csv-table, list-table, grid tables (including merged
  cells), container, tabs, youtube, literalinclude, topic,
  Sphinx domain directives, and more
- Inline markup: `:ref:`, `:doc:`, external links, bold,
  italic, literal, substitutions, index entries
- Cross-reference resolution via label scanning
- HTML rendering for complex table cells (bullet lists,
  inline formatting)
- Dark/light mode image support via MkDocs Material
  `#only-dark`/`#only-light` fragments
- Section suppression (`-skip-sections` flag)
- Project name inference from Sphinx `conf.py`

### Shared

- MkDocs nav YAML generation from document structure
- Link validation (broken links, missing anchors)
- Common types (`FileEntry`, `IDEntry`, `MarkdownWriter`)

### Makefile Targets

| Target | Description |
|--------|-------------|
| `build` | Compile the converter to `bin/` |
| `test` | Run all Go tests |
| `lint` | Run `gofmt` and `go vet` |
| `convert` | Build and run the SGML converter |
| `convert-rst` | Build and run the RST converter |
| `validate` | Build and run with link validation |
| `clean` | Remove the compiled binary |
| `setup` | Configure git hooks |

### Command-Line Options

```
pgdoc-converter [flags]
  -mode           Conversion mode: sgml or rst (default "sgml")
  -src            Path to source documentation directory
  -out            Output directory for .md files (default "./docs")
  -mkdocs         Path to mkdocs.yml (default "./mkdocs.yml")
  -version        Version label (e.g. "17.2" or "9.13")
  -copyright      Copyright string (RST mode only)
  -pgadmin-src    Path to pgAdmin source tree (for
                  literalinclude directives, RST mode only)
  -skip-sections  Comma-separated section headings to suppress
                  (RST mode only, e.g. "Sponsors,Changelog")
  -validate       Run link validation after conversion
  -verbose        Show detailed progress
```

### Makefile Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `SRC_DIR` | `/doc-source` | Path to upstream documentation |
| `OUT_DIR` | `./docs` | Output directory |
| `MKDOCS` | `./mkdocs.yml` | MkDocs configuration file |
| `VERSION` | (empty) | Version label for site_name |
| `COPYRIGHT` | (empty) | Copyright string (RST mode) |
| `PGADMIN_SRC` | (empty) | pgAdmin source (RST mode) |
| `SKIP_SECTIONS` | (empty) | Sections to suppress (RST mode) |

## TODO: Additional Component Docs Sites

- [x] PostgreSQL (SGML converter)
- [x] pgAdmin 4 (RST converter)
- [ ] PgBouncer (1.24–1.25)
- [ ] pgBackRest (2.56–2.57)
- [ ] PostGIS (3.5.3–3.5.5)
- [ ] pgvector (0.8.0–0.8.1)
- [ ] pgAudit (16.1–18.0)
- [ ] psycopg2 (2.9.10)
- [x] PostgREST (14.5)

## Project Structure

```
build-all.sh        Build orchestration script
branches.yml        Branch/product configuration
builder/            Go converter source
  shared/             Shared types and Markdown writer
  convert/            SGML-to-Markdown conversion
  sgml/               SGML tokenizer, parser, entity resolver
  rst/                RST parser, converter, directive handlers
  nav/                MkDocs nav YAML generation
  validate/           Link validation
docs/               MkDocs support files (on main branch)
  img/                Site images (logo, favicon)
  overrides/          MkDocs Material template overrides
  stylesheets/        Custom CSS
mkdocs.yml          MkDocs skeleton configuration
Makefile            Build targets
```
