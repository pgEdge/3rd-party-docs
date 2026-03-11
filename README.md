# PostgreSQL Documentation

[![CI](https://github.com/pgEdge/postgresql-docs/actions/workflows/ci.yml/badge.svg)](https://github.com/pgEdge/postgresql-docs/actions/workflows/ci.yml)

MkDocs Material site for the PostgreSQL documentation, converted
from the upstream SGML sources.

The `main` branch contains the builder tooling. Additional
branches contain generated docs for each major version, named
`pg<version>` (e.g. `pg17`, `pg18`).

## Prerequisites

- Go 1.25+
- Python 3 with [MkDocs Material](https://squidfunk.github.io/mkdocs-material/)
- PostgreSQL source tree (for the SGML docs)

## Quick Start

Clone the PostgreSQL source:

```sh
git clone --depth 1 https://github.com/postgres/postgres.git postgresql
```

Build the converter and generate the docs:

```sh
make build
make convert SRC_DIR=postgresql/doc/src/sgml VERSION=19devel
```

Preview the site locally:

```sh
mkdocs serve
```

## Builder

The `builder/` directory contains a Go tool that converts
PostgreSQL's SGML/DocBook documentation to Markdown. It handles:

- SGML entity resolution and parsing
- DocBook-to-Markdown conversion (inline, block, tables, xrefs)
- `func_table_entry` tables split into proper multi-column layout
- Image copying from the PG source tree
- MkDocs nav generation from document structure

### Makefile Targets

| Target     | Description                              |
| ---------- | ---------------------------------------- |
| `setup`    | Configure git hooks (run once after clone)|
| `build`    | Compile the converter to `bin/`          |
| `test`     | Run all Go tests                         |
| `lint`     | Run `gofmt` and `go vet`                |
| `convert`  | Build and run the converter              |
| `validate` | Build and run with link validation       |
| `clean`    | Remove the compiled binary               |

### Usage

```
pgdoc-converter [flags]
  -src       Path to PostgreSQL doc/src/sgml/ directory
  -out       Output directory for .md files (default "./docs")
  -mkdocs    Path to mkdocs.yml (default "./mkdocs.yml")
  -version   PostgreSQL version label (e.g. "17.2")
  -validate  Run link validation after conversion
  -verbose   Verbose output
```

## Project Structure

```
builder/          Go converter source
  convert/          SGML-to-Markdown conversion
  sgml/             SGML tokenizer, parser, entity resolver
  nav/              MkDocs nav YAML generation
  validate/         Link validation
docs/             Generated Markdown + MkDocs support files
  img/              Site images (logo, favicon)
  overrides/        MkDocs Material template overrides
  stylesheets/      Custom CSS
mkdocs.yml        MkDocs configuration
Makefile          Build targets
```
