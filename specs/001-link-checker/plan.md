# mdcheck — Go Implementation Plan

## Module Structure

```text
cmd/mdcheck/main.go     CLI entry point — argument parsing, orchestration
internal/parser.go      Markdown parsing — extract links with positions
internal/parser_test.go Tests for link extraction
internal/checker.go     Link checking — HTTP HEAD + file existence
internal/checker_test.go Tests for link verification
```

## Data Model

### Link

```go
type Link struct {
    Text     string // display text (e.g., "example")
    Target   string // URL or file path
    File     string // source Markdown file
    Line     int    // line number in source file
    LinkType string // "inline", "reference", or "autolink"
}
```

### Result

```go
type Result struct {
    Link   Link
    Broken bool
    Reason string // e.g., "HTTP 404" or "file not found"
}
```

## Parser Algorithm

The parser reads Markdown line by line and extracts links using three
patterns:

1. **Inline links**: regex `\[([^\]]+)\]\(([^)]+)\)` captures text
   and target
2. **Reference links**: regex `\[([^\]]+)\]\[([^\]]*)\]` captures text
   and ref key; a second pass finds `[ref]: url` definitions
3. **Autolinks**: regex `<(https?://[^>]+)>` captures URL

Each match records the file path, line number, extracted text, resolved
target, and link type.

## Checker Algorithm

For each Link:

1. If target starts with `http://` or `https://`, perform HTTP HEAD
   with a 10-second timeout. Non-2xx status → broken.
2. If target is a relative path, resolve relative to the source file's
   directory. Check file existence. Not found → broken.
3. Skip fragment-only links (`#heading`) — these reference anchors
   within the same file and are not checked.

## FR Mapping

| FR | Implementation |
| --- | --- |
| FR-001 | parser.go: inline link regex |
| FR-002 | parser.go: reference link regex + definition pass |
| FR-003 | parser.go: autolink regex |
| FR-004 | checker.go: HTTP HEAD with timeout |
| FR-005 | checker.go: file existence check |
| FR-006 | checker.go: Result struct with reason |
| FR-007 | main.go: argument parsing |
| FR-008 | main.go: recursive file discovery |
| FR-009 | main.go: exit code logic |
| FR-010 | main.go: summary line output |

## Test Strategy

- Parser tests: provide Markdown strings, assert correct Link structs
- Checker tests: use httptest for URL checking, temp files for local
  path checking
- Coverage target: 85% on parser.go and checker.go
