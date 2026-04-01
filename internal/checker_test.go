package internal

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckURL_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	links := []Link{
		{Text: "test", Target: server.URL, File: "test.md", Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Broken {
		t.Errorf("expected link to be valid, got broken: %s", results[0].Reason)
	}
}

func TestCheckURL_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	links := []Link{
		{Text: "missing", Target: server.URL, File: "test.md", Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if !results[0].Broken {
		t.Error("expected link to be broken")
	}
	if results[0].Reason != "HTTP 404" {
		t.Errorf("reason: got %q, want %q", results[0].Reason, "HTTP 404")
	}
}

func TestCheckURL_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	links := []Link{
		{Text: "error", Target: server.URL, File: "test.md", Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if !results[0].Broken {
		t.Error("expected link to be broken")
	}
	if results[0].Reason != "HTTP 500" {
		t.Errorf("reason: got %q", results[0].Reason)
	}
}

func TestCheckLocalFile_Exists(t *testing.T) {
	// Create a temp directory with a file
	dir := t.TempDir()
	targetFile := filepath.Join(dir, "exists.md")
	if err := os.WriteFile(targetFile, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}

	sourceFile := filepath.Join(dir, "source.md")
	links := []Link{
		{Text: "local", Target: "exists.md", File: sourceFile, Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if results[0].Broken {
		t.Errorf("expected link to be valid, got broken: %s", results[0].Reason)
	}
}

func TestCheckLocalFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	sourceFile := filepath.Join(dir, "source.md")

	links := []Link{
		{Text: "missing", Target: "nonexistent.md", File: sourceFile, Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if !results[0].Broken {
		t.Error("expected link to be broken")
	}
	if results[0].Reason != "file not found" {
		t.Errorf("reason: got %q", results[0].Reason)
	}
}

func TestCheckFragmentOnly_Skipped(t *testing.T) {
	links := []Link{
		{Text: "heading", Target: "#section", File: "test.md", Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if results[0].Broken {
		t.Error("fragment-only links should not be broken")
	}
}

func TestCheckLinks_FragmentValid(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "README.md")
	os.WriteFile(target, []byte("# Title\n\n## Setup\n\nContent here.\n"), 0644)

	source := filepath.Join(dir, "doc.md")
	links := []Link{
		{Text: "Setup", Target: "README.md#setup", File: source, Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Broken {
		t.Errorf("expected valid link, got broken: %s", results[0].Reason)
	}
}

func TestCheckLinks_FragmentMissing(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "README.md")
	os.WriteFile(target, []byte("# Title\n\nNo setup heading here.\n"), 0644)

	source := filepath.Join(dir, "doc.md")
	links := []Link{
		{Text: "Setup", Target: "README.md#setup", File: source, Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !results[0].Broken {
		t.Error("expected broken link for missing fragment")
	}
	if !strings.Contains(results[0].Reason, "fragment #setup not found") {
		t.Errorf("expected fragment not found reason, got: %s", results[0].Reason)
	}
}

func TestCheckLinks_FragmentSlugNormalisation(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "README.md")
	os.WriteFile(target, []byte("# Title\n\n## Getting Started!\n\nContent.\n"), 0644)

	source := filepath.Join(dir, "doc.md")
	links := []Link{
		{Text: "Start", Target: "README.md#getting-started", File: source, Line: 1, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Broken {
		t.Errorf("expected valid link after slug normalisation, got broken: %s", results[0].Reason)
	}
}

func TestCheckMultipleLinks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	dir := t.TempDir()
	sourceFile := filepath.Join(dir, "source.md")

	links := []Link{
		{Text: "good url", Target: server.URL, File: sourceFile, Line: 1, LinkType: "inline"},
		{Text: "bad file", Target: "missing.md", File: sourceFile, Line: 2, LinkType: "inline"},
		{Text: "fragment", Target: "#heading", File: sourceFile, Line: 3, LinkType: "inline"},
	}

	results := CheckLinks(links)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	if results[0].Broken {
		t.Error("first link should be valid")
	}
	if !results[1].Broken {
		t.Error("second link should be broken")
	}
	if results[2].Broken {
		t.Error("fragment link should not be broken")
	}
}
