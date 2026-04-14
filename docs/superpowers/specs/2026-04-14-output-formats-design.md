# mdcheck — Output Format Support

## Purpose

Add a `--format` flag to mdcheck so users can choose between human-readable
text, machine-readable JSON, and GitHub Actions annotations. This makes the
tool useful beyond the terminal — in CI pipelines, scripts, and PR review
workflows.

## Current Behaviour

mdcheck writes broken links to stdout in a fixed text format:

```
docs/guide.md:14: [example](https://example.com/gone) — HTTP 404

12 links checked, 2 broken
```

There is no way to change the output format.

## Design

### Formatter Interface

A `Formatter` interface in `internal/formatter.go`:

```go
type Formatter interface {
    Format(results []Result, w io.Writer) error
}
```

Every formatter receives the full results slice and writes to an `io.Writer`.
`main` controls where output goes; formatters don't need to know.

Three implementations, each in its own file:

| File                       | Format   | Description                          |
|----------------------------|----------|--------------------------------------|
| `internal/text_formatter.go`   | text     | Current output, reproduced exactly   |
| `internal/json_formatter.go`   | json     | Envelope with metadata + results     |
| `internal/github_formatter.go` | github   | `::error` annotation lines for PRs   |

### Text Format (default)

Identical to the current output. Existing users see no change.

```
docs/guide.md:14: [example](https://example.com/gone) — HTTP 404

12 links checked, 2 broken
```

### JSON Format

Envelope structure with summary and full results array:

```json
{
  "summary": {
    "total": 12,
    "broken": 2
  },
  "results": [
    {
      "file": "docs/guide.md",
      "line": 14,
      "text": "example",
      "target": "https://example.com/gone",
      "broken": true,
      "reason": "HTTP 404"
    },
    {
      "file": "docs/guide.md",
      "line": 22,
      "text": "setup",
      "target": "./setup.md",
      "broken": false,
      "reason": ""
    }
  ]
}
```

Key decisions:

- All results included, not just broken — consumers can filter, and knowing
  what was checked is valuable
- `reason` is empty string when not broken — keeps the schema uniform, no null
  checks needed
- No timestamp — the tool is stateless and the caller knows when they ran it

### GitHub Actions Format

Only broken links emit `::error` annotations:

```
::error file=docs/guide.md,line=14::Broken link: [example](https://example.com/gone) — HTTP 404
```

Summary printed as a `::notice`:

```
::notice::12 links checked, 2 broken
```

Key decisions:

- Valid links produce no output — annotations are for problems
- `::error` level, not `::warning` — consistent with the non-zero exit code
- Summary as `::notice` — visible in the Actions log without cluttering the PR

### CLI Flag

- **Flag:** `--format <text|json|github>` with `text` as the default
- **Short form:** `-f` as an alias
- **Invalid values:** error to stderr listing valid options, exit code 2
- **Flag parsing:** `flag` standard library — no external dependencies
- **Positional arg:** path moves to `flag.Args()` after flag parsing

### Wiring in main.go

Parse the flag, select the formatter, run the existing pipeline, call
`formatter.Format(results, os.Stdout)`. The parse and check logic is
untouched.

## Testing Strategy

Each formatter gets its own test file:

- **`text_formatter_test.go`** — verifies output matches current format
  exactly (regression safety net)
- **`json_formatter_test.go`** — round-trips through `json.Unmarshal`,
  checks envelope fields, verifies all results present
- **`github_formatter_test.go`** — verifies `::error` lines contain correct
  file/line metadata, verifies clean links produce no output

Tests use a fixed `[]Result` slice as input. No HTTP, no file system — pure
unit tests on formatting logic.

Existing tests in `parser_test.go` and `checker_test.go` remain untouched.

## Scope

### In scope

- `Formatter` interface and three implementations
- `--format` / `-f` flag in `main.go`
- Unit tests for each formatter
- Extracting current output logic into `TextFormatter`

### Out of scope

- Additional formats (JUnit XML, SARIF, etc.) — can be added later
- Coloured terminal output
- Output to file (`--output` flag)
- Any changes to link parsing or checking logic
