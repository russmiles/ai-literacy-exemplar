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
