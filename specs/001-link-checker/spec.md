# mdcheck — Markdown Link Checker Specification

## Purpose

A command-line tool that checks Markdown files for broken links. Given
a file or directory, it extracts all links, verifies each one (HTTP
HEAD for URLs, file existence for local paths), and reports any that
are broken.

## User Stories

### P1: Developer checking documentation

As a developer, I want to run `mdcheck` against my Markdown files so
that I can find and fix broken links before they reach production.

### P2: CI integration

As a CI pipeline, I want `mdcheck` to exit non-zero when broken links
are found so that I can fail the build on documentation quality issues.

## Acceptance Scenarios

### Scenario 1: No broken links

Given a Markdown file with all valid links
When I run `mdcheck` against it
Then the exit code is 0
And the output reports all links checked with no failures

### Scenario 2: Broken URL

Given a Markdown file containing `[example](https://httpstat.us/404)`
When I run `mdcheck` against it
Then the exit code is 1
And the output includes the file name, line number, link text, and URL

### Scenario 3: Broken local link

Given a Markdown file containing `[readme](./nonexistent.md)`
When I run `mdcheck` against it
Then the exit code is 1
And the output includes the file name, line number, and missing path

### Scenario 4: Multiple files

Given a directory containing multiple Markdown files
When I run `mdcheck` against the directory
Then all `.md` files are checked recursively
And the report aggregates results across all files

### Scenario 5: Mixed link types

Given a Markdown file containing inline links `[text](url)`,
reference links `[text][ref]`, and autolinks `<url>`
When I run `mdcheck` against it
Then all three link types are extracted and checked

### Scenario 6: Fragment link — heading exists

Given a Markdown file with the link `[Setup](README.md#setup)`
When I run `mdcheck` against it
And README.md exists with a `## Setup` heading
Then the link is reported as valid

### Scenario 7: Fragment link — heading missing

Given a Markdown file with the link `[Setup](README.md#setup)`
When I run `mdcheck` against it
And README.md exists but has no `## Setup` heading
Then the link is reported as broken with reason "fragment #setup not found in README.md"

## Functional Requirements

- **FR-001**: Parse inline links of the form `[text](url)`
- **FR-002**: Parse reference links of the form `[text][ref]` with
  corresponding `[ref]: url` definitions
- **FR-003**: Parse autolinks of the form `<url>`
- **FR-004**: For HTTP/HTTPS URLs, perform an HTTP HEAD request and
  treat non-2xx responses as broken
- **FR-005**: For local file paths, check file existence relative to
  the Markdown file's directory
- **FR-006**: Report each broken link with: file path, line number,
  link text, target URL or path, and reason (HTTP status or "not found")
- **FR-007**: Accept a single file path or a directory path as argument
- **FR-008**: When given a directory, recursively find all `.md` files
- **FR-009**: Exit with code 0 if no broken links, code 1 if any broken
- **FR-010**: Print a summary line: "N links checked, M broken"
- **FR-011**: When a local file link contains a fragment (e.g.,
  `file.md#heading`), verify the heading exists in the target file:
  1. Verify the target file exists (existing FR-003 behaviour)
  2. Parse the target file for ATX-style headings (lines starting with `#`)
  3. Normalise fragment and heading text to lowercase, spaces replaced
     with hyphens, non-alphanumeric characters removed (GitHub-compatible slug)
  4. Report as broken if no matching heading found, with reason
     "fragment #slug not found in file"
  5. Report as valid if a matching heading is found
- **FR-012**: Accept a `--format` flag (short form `-f`) with values
  `text`, `json`, or `github`. Default to `text`.
- **FR-013**: `text` format: output broken links as
  `file:line: [text](target) — reason`, followed by a summary line
  `N links checked, M broken`
- **FR-014**: `json` format: output a JSON envelope with `summary`
  (total, broken) and `results` array (file, line, text, target,
  broken, reason for every checked link)
- **FR-015**: `github` format: output `::error file=F,line=L::message`
  for each broken link, and `::notice::N links checked, M broken` as
  summary
- **FR-016**: Reject unknown format values with an error message
  listing valid options and exit code 2

## Success Criteria

- **SC-001**: All acceptance scenarios pass as automated tests
- **SC-002**: 85% code coverage on parser.go and checker.go
- **SC-003**: All code follows literate programming conventions
- **SC-004**: All code passes CUPID review
