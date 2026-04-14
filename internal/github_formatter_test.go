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
