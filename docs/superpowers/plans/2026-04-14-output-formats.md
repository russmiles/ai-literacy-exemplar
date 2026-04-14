# Output Format Support Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a `--format` flag to mdcheck supporting text (default), JSON (envelope with metadata), and GitHub Actions annotation output.

**Architecture:** A `Formatter` interface in `internal/` with three implementations — `TextFormatter`, `JSONFormatter`, `GitHubFormatter`. `main.go` parses `--format`/`-f`, selects the formatter, and delegates output. The parse and check pipeline is untouched.

**Tech Stack:** Go standard library only — `encoding/json`, `flag`, `io`, `fmt`.

---

## File Map

| File | Action | Responsibility |
|------|--------|----------------|
| `internal/formatter.go` | Create | `Formatter` interface definition |
| `internal/text_formatter.go` | Create | Text output (matches current behaviour) |
| `internal/text_formatter_test.go` | Create | Tests for text output |
| `internal/json_formatter.go` | Create | JSON envelope output |
| `internal/json_formatter_test.go` | Create | Tests for JSON output |
| `internal/github_formatter.go` | Create | GitHub Actions annotation output |
| `internal/github_formatter_test.go` | Create | Tests for GitHub output |
| `cmd/mdcheck/main.go` | Modify | Add `--format` flag, wire formatter |

---

### Task 1: Formatter Interface

**Files:**
- Create: `internal/formatter.go`

- [ ] **Step 1: Create the interface file**

```go
// The Formatter interface decouples output rendering from link checking.
// Each implementation writes results in a specific format (text, JSON,
// GitHub annotations) to the provided writer. This separation means the
// checker pipeline does not need to know how results are displayed, and
// new formats can be added by implementing a single method.
//
// The interface takes an io.Writer rather than returning a string so that
// formatters can stream output and callers control the destination.
package internal

import "io"

// Formatter renders check results to a writer in a specific output format.
type Formatter interface {
	Format(results []Result, w io.Writer) error
}
```

- [ ] **Step 2: Run existing tests to confirm nothing is broken**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./...`
Expected: All existing tests PASS

- [ ] **Step 3: Commit**

```bash
git add internal/formatter.go
git commit -m "Add Formatter interface for pluggable output rendering"
```

---

### Task 2: Text Formatter (TDD)

**Files:**
- Create: `internal/text_formatter_test.go`
- Create: `internal/text_formatter.go`

- [ ] **Step 1: Write the failing test**

```go
// Tests for TextFormatter verify that the default text output matches the
// original mdcheck output exactly. This is the regression safety net —
// existing users must see no change in behaviour when --format is not
// specified.
package internal

import (
	"bytes"
	"testing"
)

func TestTextFormatter_BrokenLinks(t *testing.T) {
	results := []Result{
		{
			Link:   Link{Text: "example", Target: "https://example.com/gone", File: "docs/guide.md", Line: 14, LinkType: "inline"},
			Broken: true,
			Reason: "HTTP 404",
		},
		{
			Link:   Link{Text: "setup", Target: "./setup.md", File: "docs/guide.md", Line: 22, LinkType: "inline"},
			Broken: false,
			Reason: "",
		},
	}

	var buf bytes.Buffer
	f := TextFormatter{}
	err := f.Format(results, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()

	// Must contain the broken link line in the original format
	want := "docs/guide.md:14: [example](https://example.com/gone) — HTTP 404\n"
	if !bytes.Contains([]byte(got), []byte(want)) {
		t.Errorf("missing broken link line\ngot:\n%s\nwant substring:\n%s", got, want)
	}

	// Must contain the summary line
	wantSummary := "\n2 links checked, 1 broken\n"
	if !bytes.Contains([]byte(got), []byte(wantSummary)) {
		t.Errorf("missing summary line\ngot:\n%s\nwant substring:\n%s", got, wantSummary)
	}
}

func TestTextFormatter_NoBrokenLinks(t *testing.T) {
	results := []Result{
		{
			Link:   Link{Text: "ok", Target: "https://example.com", File: "test.md", Line: 1, LinkType: "inline"},
			Broken: false,
		},
	}

	var buf bytes.Buffer
	f := TextFormatter{}
	err := f.Format(results, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	wantSummary := "\n1 links checked, 0 broken\n"
	if got != wantSummary {
		t.Errorf("got:\n%q\nwant:\n%q", got, wantSummary)
	}
}

func TestTextFormatter_EmptyResults(t *testing.T) {
	var buf bytes.Buffer
	f := TextFormatter{}
	err := f.Format([]Result{}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	want := "\n0 links checked, 0 broken\n"
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./internal/ -run TestTextFormatter -v`
Expected: FAIL — `TextFormatter` not defined

- [ ] **Step 3: Write minimal implementation**

```go
// TextFormatter reproduces the original mdcheck output: one line per broken
// link showing file, line number, link text, target, and reason, followed
// by a summary line. Valid links produce no output lines — only the summary
// mentions them via the total count.
//
// This formatter exists so that the default behaviour is preserved exactly
// when users add --format support. If you change the output here, you change
// what every existing user sees.
package internal

import (
	"fmt"
	"io"
)

// TextFormatter writes human-readable plain text output.
type TextFormatter struct{}

// Format writes broken link details and a summary line to w.
func (f TextFormatter) Format(results []Result, w io.Writer) error {
	broken := 0
	for _, r := range results {
		if r.Broken {
			broken++
			_, err := fmt.Fprintf(w, "%s:%d: [%s](%s) — %s\n",
				r.Link.File, r.Link.Line, r.Link.Text, r.Link.Target, r.Reason)
			if err != nil {
				return err
			}
		}
	}

	_, err := fmt.Fprintf(w, "\n%d links checked, %d broken\n", len(results), broken)
	return err
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./internal/ -run TestTextFormatter -v`
Expected: All three tests PASS

- [ ] **Step 5: Run all tests to confirm no regressions**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./...`
Expected: All tests PASS

- [ ] **Step 6: Commit**

```bash
git add internal/text_formatter.go internal/text_formatter_test.go
git commit -m "Add TextFormatter that reproduces original output format"
```

---

### Task 3: JSON Formatter (TDD)

**Files:**
- Create: `internal/json_formatter_test.go`
- Create: `internal/json_formatter.go`

- [ ] **Step 1: Write the failing test**

```go
// Tests for JSONFormatter verify the envelope structure: a top-level object
// with "summary" (total and broken counts) and "results" (every checked
// link). The tests round-trip through json.Unmarshal to validate structure
// rather than comparing raw strings, which would be fragile to whitespace.
package internal

import (
	"bytes"
	"encoding/json"
	"testing"
)

// jsonOutput mirrors the JSON envelope so tests can unmarshal and inspect.
type jsonOutput struct {
	Summary struct {
		Total  int `json:"total"`
		Broken int `json:"broken"`
	} `json:"summary"`
	Results []struct {
		File   string `json:"file"`
		Line   int    `json:"line"`
		Text   string `json:"text"`
		Target string `json:"target"`
		Broken bool   `json:"broken"`
		Reason string `json:"reason"`
	} `json:"results"`
}

func TestJSONFormatter_Envelope(t *testing.T) {
	results := []Result{
		{
			Link:   Link{Text: "example", Target: "https://example.com/gone", File: "docs/guide.md", Line: 14, LinkType: "inline"},
			Broken: true,
			Reason: "HTTP 404",
		},
		{
			Link:   Link{Text: "setup", Target: "./setup.md", File: "docs/guide.md", Line: 22, LinkType: "inline"},
			Broken: false,
			Reason: "",
		},
	}

	var buf bytes.Buffer
	f := JSONFormatter{}
	err := f.Format(results, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var out jsonOutput
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("invalid JSON: %v\nraw: %s", err, buf.String())
	}

	if out.Summary.Total != 2 {
		t.Errorf("summary.total: got %d, want 2", out.Summary.Total)
	}
	if out.Summary.Broken != 1 {
		t.Errorf("summary.broken: got %d, want 1", out.Summary.Broken)
	}
	if len(out.Results) != 2 {
		t.Fatalf("results length: got %d, want 2", len(out.Results))
	}
	if out.Results[0].File != "docs/guide.md" {
		t.Errorf("results[0].file: got %q", out.Results[0].File)
	}
	if !out.Results[0].Broken {
		t.Error("results[0] should be broken")
	}
	if out.Results[0].Reason != "HTTP 404" {
		t.Errorf("results[0].reason: got %q", out.Results[0].Reason)
	}
	if out.Results[1].Broken {
		t.Error("results[1] should not be broken")
	}
	if out.Results[1].Reason != "" {
		t.Errorf("results[1].reason: got %q, want empty string", out.Results[1].Reason)
	}
}

func TestJSONFormatter_EmptyResults(t *testing.T) {
	var buf bytes.Buffer
	f := JSONFormatter{}
	err := f.Format([]Result{}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var out jsonOutput
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if out.Summary.Total != 0 {
		t.Errorf("summary.total: got %d, want 0", out.Summary.Total)
	}
	if len(out.Results) != 0 {
		t.Errorf("results length: got %d, want 0", len(out.Results))
	}
}

func TestJSONFormatter_ValidJSON(t *testing.T) {
	results := []Result{
		{
			Link:   Link{Text: "test", Target: "https://example.com", File: "test.md", Line: 1, LinkType: "inline"},
			Broken: false,
		},
	}

	var buf bytes.Buffer
	f := JSONFormatter{}
	_ = f.Format(results, &buf)

	if !json.Valid(buf.Bytes()) {
		t.Errorf("output is not valid JSON:\n%s", buf.String())
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./internal/ -run TestJSONFormatter -v`
Expected: FAIL — `JSONFormatter` not defined

- [ ] **Step 3: Write minimal implementation**

```go
// JSONFormatter outputs an envelope containing a summary object and a
// results array. Every checked link appears in the results — not just
// broken ones — because consumers often need to know what was checked,
// not just what failed. The schema is kept uniform: "reason" is always
// present as a string (empty when the link is valid) so consumers do
// not need null-handling logic.
//
// The output is indented for human readability. Machine consumers are
// unaffected since JSON parsers ignore whitespace.
package internal

import (
	"encoding/json"
	"io"
)

// JSONFormatter writes JSON envelope output with summary and results.
type JSONFormatter struct{}

// jsonResult is the per-link structure in the JSON output. It flattens
// the Link fields to avoid nested objects — the consumer does not need
// to know about mdcheck's internal Link type.
type jsonResult struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Text   string `json:"text"`
	Target string `json:"target"`
	Broken bool   `json:"broken"`
	Reason string `json:"reason"`
}

// jsonEnvelope is the top-level JSON structure.
type jsonEnvelope struct {
	Summary struct {
		Total  int `json:"total"`
		Broken int `json:"broken"`
	} `json:"summary"`
	Results []jsonResult `json:"results"`
}

// Format writes the JSON envelope to w.
func (f JSONFormatter) Format(results []Result, w io.Writer) error {
	broken := 0
	jsonResults := make([]jsonResult, 0, len(results))

	for _, r := range results {
		if r.Broken {
			broken++
		}
		jsonResults = append(jsonResults, jsonResult{
			File:   r.Link.File,
			Line:   r.Link.Line,
			Text:   r.Link.Text,
			Target: r.Link.Target,
			Broken: r.Broken,
			Reason: r.Reason,
		})
	}

	env := jsonEnvelope{
		Results: jsonResults,
	}
	env.Summary.Total = len(results)
	env.Summary.Broken = broken

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(env)
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./internal/ -run TestJSONFormatter -v`
Expected: All three tests PASS

- [ ] **Step 5: Run all tests to confirm no regressions**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./...`
Expected: All tests PASS

- [ ] **Step 6: Commit**

```bash
git add internal/json_formatter.go internal/json_formatter_test.go
git commit -m "Add JSONFormatter with envelope structure"
```

---

### Task 4: GitHub Formatter (TDD)

**Files:**
- Create: `internal/github_formatter_test.go`
- Create: `internal/github_formatter.go`

- [ ] **Step 1: Write the failing test**

```go
// Tests for GitHubFormatter verify that broken links produce ::error
// annotations with correct file and line metadata, valid links produce
// no output, and the summary appears as a ::notice.
package internal

import (
	"bytes"
	"strings"
	"testing"
)

func TestGitHubFormatter_BrokenLinks(t *testing.T) {
	results := []Result{
		{
			Link:   Link{Text: "example", Target: "https://example.com/gone", File: "docs/guide.md", Line: 14, LinkType: "inline"},
			Broken: true,
			Reason: "HTTP 404",
		},
		{
			Link:   Link{Text: "setup", Target: "./setup.md", File: "docs/guide.md", Line: 22, LinkType: "inline"},
			Broken: false,
			Reason: "",
		},
	}

	var buf bytes.Buffer
	f := GitHubFormatter{}
	err := f.Format(results, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()

	// Broken link must produce an ::error annotation
	wantError := "::error file=docs/guide.md,line=14::Broken link: [example](https://example.com/gone) — HTTP 404\n"
	if !strings.Contains(got, wantError) {
		t.Errorf("missing error annotation\ngot:\n%s\nwant substring:\n%s", got, wantError)
	}

	// Valid link must NOT produce any annotation
	if strings.Contains(got, "setup.md") {
		t.Error("valid link should not appear in output")
	}

	// Summary must appear as ::notice
	wantNotice := "::notice::2 links checked, 1 broken\n"
	if !strings.Contains(got, wantNotice) {
		t.Errorf("missing notice summary\ngot:\n%s\nwant substring:\n%s", got, wantNotice)
	}
}

func TestGitHubFormatter_NoBrokenLinks(t *testing.T) {
	results := []Result{
		{
			Link:   Link{Text: "ok", Target: "https://example.com", File: "test.md", Line: 1, LinkType: "inline"},
			Broken: false,
		},
	}

	var buf bytes.Buffer
	f := GitHubFormatter{}
	err := f.Format(results, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()

	// No ::error lines when nothing is broken
	if strings.Contains(got, "::error") {
		t.Error("should not contain ::error when no links are broken")
	}

	// Summary still appears
	wantNotice := "::notice::1 links checked, 0 broken\n"
	if got != wantNotice {
		t.Errorf("got:\n%q\nwant:\n%q", got, wantNotice)
	}
}

func TestGitHubFormatter_EmptyResults(t *testing.T) {
	var buf bytes.Buffer
	f := GitHubFormatter{}
	err := f.Format([]Result{}, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := buf.String()
	want := "::notice::0 links checked, 0 broken\n"
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./internal/ -run TestGitHubFormatter -v`
Expected: FAIL — `GitHubFormatter` not defined

- [ ] **Step 3: Write minimal implementation**

```go
// GitHubFormatter produces GitHub Actions workflow command annotations.
// GitHub's runner parses lines matching ::error file=F,line=L::message
// and renders them as inline annotations on pull request diffs. This
// makes broken links visible exactly where they occur in the source.
//
// Only broken links produce ::error annotations — valid links generate
// no output because annotations are a problem-reporting mechanism, not
// a status mechanism. The summary is emitted as a ::notice so it
// appears in the Actions log without cluttering the PR diff.
package internal

import (
	"fmt"
	"io"
)

// GitHubFormatter writes GitHub Actions annotation output.
type GitHubFormatter struct{}

// Format writes ::error annotations for broken links and a ::notice summary.
func (f GitHubFormatter) Format(results []Result, w io.Writer) error {
	broken := 0
	for _, r := range results {
		if r.Broken {
			broken++
			_, err := fmt.Fprintf(w, "::error file=%s,line=%d::Broken link: [%s](%s) — %s\n",
				r.Link.File, r.Link.Line, r.Link.Text, r.Link.Target, r.Reason)
			if err != nil {
				return err
			}
		}
	}

	_, err := fmt.Fprintf(w, "::notice::%d links checked, %d broken\n", len(results), broken)
	return err
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./internal/ -run TestGitHubFormatter -v`
Expected: All three tests PASS

- [ ] **Step 5: Run all tests to confirm no regressions**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./...`
Expected: All tests PASS

- [ ] **Step 6: Commit**

```bash
git add internal/github_formatter.go internal/github_formatter_test.go
git commit -m "Add GitHubFormatter with ::error annotations for PRs"
```

---

### Task 5: Wire --format Flag in main.go

**Files:**
- Modify: `cmd/mdcheck/main.go`

- [ ] **Step 1: Replace the full contents of main.go**

The complete updated file:

```go
// mdcheck is a command-line Markdown link checker. Given a file or directory,
// it extracts all links from Markdown files, verifies each one, and reports
// any that are broken.
//
// This tool exists because broken links in documentation erode trust —
// a reader who encounters a dead link loses confidence in the rest of the
// document. Catching broken links before they reach production is a
// documentation quality gate, analogous to how tests are a code quality gate.
//
// The tool is deliberately simple: parse, check, report. It does not attempt
// to fix links, suggest replacements, or cache results across runs. Each
// invocation is a fresh, stateless verification pass.
//
// Output format is controlled by the --format flag: "text" (default) for
// human-readable output, "json" for machine-readable envelopes, and "github"
// for GitHub Actions annotations that surface inline on PR diffs.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/russmiles/ai-literacy-exemplar/internal"
)

func main() {
	format := flag.String("format", "text", "output format: text, json, github")
	flag.StringVar(format, "f", "text", "output format (shorthand)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: mdcheck [--format text|json|github] <file-or-directory>\n")
		os.Exit(2)
	}

	formatter, err := selectFormatter(*format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	target := args[0]
	files, err := findMarkdownFiles(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No Markdown files found at %s\n", target)
		os.Exit(2)
	}

	var allLinks []internal.Link
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", file, err)
			continue
		}
		links := internal.ParseLinks(string(content), file)
		allLinks = append(allLinks, links...)
	}

	results := internal.CheckLinks(allLinks)

	if err := formatter.Format(results, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
		os.Exit(2)
	}

	broken := 0
	for _, r := range results {
		if r.Broken {
			broken++
		}
	}
	if broken > 0 {
		os.Exit(1)
	}
}

// selectFormatter returns the appropriate Formatter for the given format
// name, or an error if the name is not recognised.
func selectFormatter(name string) (internal.Formatter, error) {
	switch name {
	case "text":
		return internal.TextFormatter{}, nil
	case "json":
		return internal.JSONFormatter{}, nil
	case "github":
		return internal.GitHubFormatter{}, nil
	default:
		return nil, fmt.Errorf("unknown format %q (valid: text, json, github)", name)
	}
}

// findMarkdownFiles returns all .md files at the given path. If the path
// is a file, it returns that file. If it is a directory, it walks
// recursively and collects all .md files.
func findMarkdownFiles(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		if strings.HasSuffix(path, ".md") {
			return []string{path}, nil
		}
		return nil, fmt.Errorf("%s is not a Markdown file", path)
	}

	var files []string
	err = filepath.Walk(path, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fi.IsDir() && strings.HasSuffix(p, ".md") {
			files = append(files, p)
		}
		return nil
	})
	return files, err
}
```

- [ ] **Step 2: Run all tests to confirm no regressions**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go test ./...`
Expected: All tests PASS

- [ ] **Step 3: Build and smoke test**

Run: `cd /Users/russellmiles/code/russmiles/ai-literacy-exemplar && go build -o mdcheck ./cmd/mdcheck/`
Expected: Build succeeds with no errors

Run: `./mdcheck --format text README.md` (or any .md file present)
Expected: Original text format output

Run: `./mdcheck --format json README.md`
Expected: JSON envelope output

Run: `./mdcheck --format github README.md`
Expected: GitHub annotation output

Run: `./mdcheck --format invalid README.md`
Expected: Error message to stderr listing valid formats, exit code 2

- [ ] **Step 4: Commit**

```bash
git add cmd/mdcheck/main.go
git commit -m "Wire --format flag in main.go with text, json, github options"
```

---

### Task 6: Update Spec

**Files:**
- Modify: `specs/001-link-checker/spec.md`

- [ ] **Step 1: Add the new functional requirements to the spec**

Add the following after FR-011:

```markdown
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
```

- [ ] **Step 2: Commit**

```bash
git add specs/001-link-checker/spec.md
git commit -m "Add output format functional requirements FR-012 through FR-016"
```
