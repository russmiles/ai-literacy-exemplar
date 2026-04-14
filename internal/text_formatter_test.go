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
