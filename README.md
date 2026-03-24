# PostgreSQL & Component Documentation

[![CI](https://github.com/pgEdge/postgresql-docs/actions/workflows/ci.yml/badge.svg)](https://github.com/pgEdge/postgresql-docs/actions/workflows/ci.yml)

MkDocs Material documentation sites converted from upstream
sources:

- **PostgreSQL** — SGML/DocBook sources converted to Markdown
- **pgAdmin 4** — reStructuredText (RST) sources converted to
  Markdown
- **PostgREST** — reStructuredText (RST) sources converted to
  Markdown
- **PostGIS** — XML/DocBook sources converted to Markdown
- **psycopg2** — reStructuredText (RST) sources converted to
  Markdown
- **pgBackRest** — Custom XML sources converted to Markdown
- **PgBouncer** — Markdown sources (split/copied)
- **pgvector** — Markdown sources (split by section)
- **pgAudit** — Markdown sources (split by section)
- **pg_vectorize** — MkDocs site (imported with upstream nav)
- **pg_tokenizer** — Markdown sources (multi-file copy)
- **vchord_bm25** — Markdown sources (split by section)
- **pg_cron** — Markdown sources (split by section)
- **pgmq** — MkDocs site (imported with upstream nav)
- **pg_stat_monitor** — Markdown sources (split by section)
- **pldebugger** — Markdown sources (split by section)
- **system_stats** — Markdown sources (split by section)

## How It Works

This project uses an unconventional git branching model. The
`main` branch contains **only** the Go converter tooling, a
skeleton `mkdocs.yml`, and MkDocs support files (CSS, images,
overrides). It contains no documentation content.

All generated documentation lives on **product/version
branches**. Each branch is an orphan-like branch that combines
the tooling from `main` with the converted Markdown output for
one product at one version. The `build-all.sh` script automates
this: for each branch it checks out the branch, merges tooling
from `main`, clones/fetches the upstream source, runs the
converter, and commits the result.

This means:

- **Tooling changes** go on `main` and propagate to all
  branches on the next build.
- **Generated docs** are never committed to `main` — each
  branch is self-contained with its own `docs/` and
  `mkdocs.yml`.
- Each branch can be independently deployed as a standalone
  MkDocs Material site.

## Branch Layout

| Branch | Product | Source Format |
|--------|---------|---------------|
| `pg16` .. `pg19` | PostgreSQL 16–19 | SGML (`doc/src/sgml/`) |
| `pgadmin911` .. `pgadmin913` | pgAdmin 4 v9.11–v9.13 | RST (`docs/en_US/`) |
| `pgadminmaster` | pgAdmin 4 dev | RST (`docs/en_US/`) |
| `postgrest145` | PostgREST v14.5 | RST (`docs/`) |
| `postgrestmaster` | PostgREST dev | RST (`docs/`) |
| `postgis355`, `postgis362` | PostGIS 3.5–3.6 | XML/DocBook (`doc/`) |
| `postgismaster` | PostGIS dev | XML/DocBook (`doc/`) |
| `psycopg2910` | psycopg2 v2.9.10 | RST (`doc/src/`) |
| `psycopg2master` | psycopg2 dev | RST (`doc/src/`) |
| `pgbackrest257` .. `pgbackrest258` | pgBackRest 2.57–2.58 | Custom XML (`doc/`) |
| `pgbackrestmaster` | pgBackRest dev | Custom XML (`doc/`) |
| `pgbouncer124` .. `pgbouncer125` | PgBouncer 1.24–1.25 | Markdown (`doc/`) |
| `pgbouncermaster` | PgBouncer dev | Markdown (`doc/`) |
| `pgvector080` .. `pgvector081` | pgvector 0.8.0–0.8.1 | Markdown (`README.md`) |
| `pgvectormaster` | pgvector dev | Markdown (`README.md`) |
| `pgaudit161` .. `pgaudit180` | pgAudit 16.1–18.0 | Markdown (`README.md`) |
| `pgauditmaster` | pgAudit dev | Markdown (`README.md`) |
| `pgvectorize0260` | pg_vectorize v0.26.0 | MkDocs (`docs/`) |
| `pgvectorizemaster` | pg_vectorize dev | MkDocs (`docs/`) |
| `pgtokenizer011` | pg_tokenizer v0.1.1 | Markdown (`docs/`) |
| `pgtokenizermaster` | pg_tokenizer dev | Markdown (`docs/`) |
| `vchordbm25030` | vchord_bm25 v0.3.0 | Markdown (`README.md`) |
| `vchordbm25master` | vchord_bm25 dev | Markdown (`README.md`) |
| `pgcron167` | pg_cron v1.6.7 | Markdown (`README.md`) |
| `pgcronmaster` | pg_cron dev | Markdown (`README.md`) |
| `pgmq1110` | pgmq v1.11.0 | MkDocs (`docs/`) |
| `pgmqmaster` | pgmq dev | MkDocs (`docs/`) |
| `pgstatmonitor231` | pg_stat_monitor v2.3.1 | Markdown (`README.md`) |
| `pgstatmonitormaster` | pg_stat_monitor dev | Markdown (`README.md`) |
| `pldebugger19` | pldebugger v1.9 | Markdown (`README-pldebugger.md`) |
| `pldebuggermaster` | pldebugger dev | Markdown (`README-pldebugger.md`) |
| `systemstats321` | system_stats v3.2.1 | Markdown (`README.md`) |
| `systemstatsmaster` | system_stats dev | Markdown (`README.md`) |

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
make convert SRC_DIR=/path/to/postgresql/doc/src/sgml \
    VERSION=17.2

# pgAdmin 4 (RST mode)
make convert-rst SRC_DIR=/path/to/pgadmin4/docs/en_US \
    VERSION=9.13

# PostgREST (RST mode, suppressing Sponsors section)
make convert-rst SRC_DIR=/path/to/postgrest/docs \
    VERSION=v14.5 SKIP_SECTIONS="Sponsors"

# pgBackRest (backrest mode, via binary directly)
./bin/pgdoc-converter -mode backrest \
    -src /path/to/pgbackrest/doc -version dev -verbose

# pgvector (Markdown mode)
make convert-md SRC_DIR=/path/to/pgvector \
    VERSION="pgvector v0.8.0"

# PgBouncer (Markdown mode)
make convert-md SRC_DIR=/path/to/pgbouncer/doc \
    VERSION="PgBouncer 1.25"
```

Preview the site locally:

```sh
mkdocs serve
```

## Builder

The `builder/` directory contains a Go tool
(`pgdoc-converter`) that converts upstream documentation to
Markdown suitable for MkDocs Material. It supports five
conversion modes:

### SGML Mode (PostgreSQL)

- Entity resolution and SGML parsing
- DocBook-to-Markdown conversion (100+ element handlers)
- `func_table_entry` tables split into multi-column layout
- Two-pass conversion: ID map then content generation
- Image copying from the PostgreSQL source tree

### XML Mode (PostGIS)

- Standard XML/DocBook parsing via Go's `encoding/xml`
- Entity and XInclude resolution
- WKT geometry diagrams rendered to inline SVG
- Image path rewriting for `use_directory_urls`

### RST Mode (pgAdmin, PostgREST, psycopg2)

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

### Backrest Mode (pgBackRest)

- Custom XML parser for pgBackRest's proprietary DTD
  (not DocBook)
- Entity resolution for `<!ENTITY name SYSTEM "path">`
  declarations (used heavily in release notes)
- `{[key]}` variable substitution with multi-pass
  resolution for chained references
- Block definitions (`<block-define>`/`<block>`) for
  reusable content fragments
- Executable documentation: `<execute-list>` commands
  rendered as bash code blocks with optional output
- Configuration blocks: `<backrest-config>` and
  `<postgres-config>` rendered as INI code blocks
- Self-closing brand elements (`<backrest/>`, `<postgres/>`,
  `<exe/>`)
- Semantic inline elements (`<file>`, `<path>`, `<cmd>`,
  `<br-option>`, etc.) rendered as inline code
- Cross-page link resolution with section path splitting
  (e.g. `quickstart/perform-restore`)
- Multi-page documents split by top-level `<section>`;
  single-page for small docs (FAQ, metrics, etc.)
- Nav titles derived from subtitle when title is just the
  project name

### Markdown Mode (PgBouncer, pgvector, pgAudit)

- Single-file projects split by H2 headings into separate
  pages with promoted heading levels
- Multi-file projects copied with auto-generated index page
- Internal anchor links rewritten across split files
- GitHub Alerts converted to MkDocs admonitions
- Non-doc files filtered (fragments, changelogs, etc.)

### MkDocs Mode (pg_vectorize)

- Imports upstream projects that already have an `mkdocs.yml`
- Uses upstream nav structure verbatim (no inference)
- Merges upstream `markdown_extensions` and `plugins` into
  the skeleton config (deduplicating, skipping unsupported
  plugins like `mkdocstrings`)
- Resolves `pymdownx.snippets` includes inline at copy time
- Copies entire docs tree preserving directory structure

### Shared

- MkDocs nav YAML generation from document structure
- Automatic `md_in_html` extension injection
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
| `convert-md` | Build and run the Markdown converter |
| `validate` | Build and run with link validation |
| `clean` | Remove the compiled binary |
| `setup` | Configure git hooks |

### Command-Line Options

```
pgdoc-converter [flags]
  -mode           Conversion mode: sgml, xml, rst, md,
                  or backrest (default "sgml")
  -src            Path to source documentation directory
  -out            Output directory for .md files
                  (default "./docs")
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
- [x] PgBouncer (1.24–1.25)
- [x] pgBackRest (2.57–2.58)
- [x] PostGIS (3.5.5–3.6.2)
- [x] pgvector (0.8.0–0.8.1)
- [x] pgAudit (16.1–18.0)
- [x] psycopg2 (2.9.10)
- [x] PostgREST (14.5)

### Remaining Extensions

- [x] pg_vectorize (v0.26.0)
- [x] pg_tokenizer (v0.1.1)
- [x] vchord_bm25 (v0.3.0)
- [x] pg_cron (v1.6.7)
- [x] pgmq (v1.11.0)
- [x] pg_stat_monitor (v2.3.1)
- [x] pldebugger (v1.9)
- [x] system_stats (v3.2.1)

## Project Structure

```
build-all.sh        Build orchestration script
branches.yml        Branch/product configuration
builder/            Go converter source
  backrest/           pgBackRest custom XML converter
  convert/            SGML-to-Markdown conversion
  md/                 Markdown splitter and copier
  nav/                MkDocs nav YAML generation
  rst/                RST parser, converter, directive handlers
  sgml/               SGML tokenizer, parser, entity resolver
  shared/             Shared types and Markdown writer
  validate/           Link validation
  wkt/                WKT geometry to SVG renderer (PostGIS)
docs/               MkDocs support files (on main branch)
  img/                Site images (logo, favicon)
  overrides/          MkDocs Material template overrides
  stylesheets/        Custom CSS
mkdocs.yml          MkDocs skeleton configuration
Makefile            Build targets
```
