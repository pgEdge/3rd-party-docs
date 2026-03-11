# PostgreSQL Docs - Development Guidelines

This document provides guidelines for Claude Code when working on this
project. These rules ensure consistency, quality, and maintainability.

## Project Purpose

This project implements an MKDocs/MKDocs-Material based site containing the 
PostgreSQL documentation, including the latest minor version of each current
and future major version.

The builder directory contains tooling required to convert the source SGML
docs into Markdown and update the MKDocs "nav" section. Tooling must always
be built into the /bin directory, which is .gitignore'd.

The main branch contains the working version of the tooling, and PostgreSQL
development docs (from the upstream master branch).

Additional branches contain the PostgreSQL documentation from released major
PostgreSQL versions, named after the corresponding PostgreSQL branch in the
upstream repo. 

## Code Style

### Indentation

- Use **4 spaces** for indentation (not tabs)
- Apply consistently across all code files

## Project Planning

### Long-Running Tasks

When working on complex, multi-step tasks:

- Store plan documents in `/.claude/` directory
- Include task breakdowns, progress tracking, and design decisions
- Use descriptive filenames (e.g., `phase-3-implementation-plan.md`)

## Documentation Standards

ALL documentation for the builder tooling must be in the top-level README.md.
The documentation under /docs (with the exception of supporting MKDocs 
stylesheets/overrides/images etc) should be entirely generated from the 
upstream PostgreSQL documentation.

### Markdown Formatting

**List Rendering:**

- Always leave a **blank line** before the first item in any list or
  sub-list
- This ensures proper rendering in tools like mkdocs

**Example:**

```markdown
This is a paragraph.

- First item
- Second item
  - Sub-item (note blank line before parent list)
```

### File Naming Conventions

**Root Directory:**

- Use UPPERCASE names for markdown files (e.g., `README.md`,
  `CONTRIBUTING.md`)
- Exception: file extensions remain lowercase

### Line Length

- Wrap markdown content at **79 characters or less**
- Exceptions:
  - URLs (don't split)
  - Code samples
  - Tables or structured content where wrapping breaks functionality

### Documentation Locations

**README.md:**

- Keep this as a brief summary for users browsing the repository
- Include: project overview, quick start, and links to full docs

**Full Documentation (`/docs`):**

- This contains the converted PostgreSQL documentation, along with supporting
    MKDocs/MKDocs-Material CSS, overrides, images, and other supporting files.

## Documentation Framework

- The documentation will be used with MKDocs/MKDocs Material
- The MKDocs configuration file is found in mkdocs.yml in the project root
- The table of contents (TOC) in the MKDocs configuration must be maintained

## Testing Requirements

### Test Coverage

**For New Functionality:**

- Always add tests to exercise new features
- Use the top-level Makefile: `make test`, `make lint`
- Ensure all tests run under the `go test` suite

### Running Tests

**Complete Validation:**

- Run ALL tests when verifying changes
- ALWAYS run gofmt if any Go code has been changed
- Check verbose output for failures or errors
- **Never** tail or trim test output (stdout and stderr)
- Capture full output to ensure nothing is missed

### Test Modifications

**When to Modify Tests:**

- Only modify tests if they are **expected to fail** due to your changes
- If a test fails unexpectedly, investigate the cause first
- Don't "fix" tests by changing expectations unless the change is
  intentional

### Test Cleanup

**Temporary Files:**

- Remove temporary files created during test runs
- Exception: Keep logs that may need review
- Ensure cleanup happens even if tests fail

## Security Requirements

### Input Validation

**SQL Injection Prevention:**

- Always escape user inputs to prevent injection attacks
- Exception: Tools explicitly designed to execute user-provided SQL
  (e.g., `query_database` tool)
- Use parameterized queries where possible

## Example Checklist

When making changes, verify:

- [ ] Code uses 4-space indentation
- [ ] Tests added for new functionality
- [ ] All tests pass (`make test`)
- [ ] Linting passes (`make lint`)
- [ ] Documentation updated in `/docs`
- [ ] Markdown files properly formatted (79 chars, blank lines before
      lists)
- [ ] Security considerations addressed
- [ ] No temporary files left behind

## Questions?

If you're unsure about any of these guidelines, refer to:

- Existing code patterns in the repository
- Documentation in `/docs`
- Recent git commits for context
