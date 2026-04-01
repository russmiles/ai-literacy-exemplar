// This file handles link verification — the second half of the mdcheck
// pipeline. Given a list of extracted links, the checker determines which
// are broken and why.
//
// Three verification strategies are used: HTTP HEAD requests for remote URLs
// (minimising bandwidth compared to GET), file existence checks for local
// paths (resolved relative to the source Markdown file's directory), and
// fragment validation for links that target a specific heading in a file.
//
// Fragment validation is deliberately separated from file existence checking
// because the two concerns have different failure modes and different fix
// actions. A missing file requires creating or moving a file; a missing
// heading requires editing content. Reporting them distinctly helps the user
// fix the right thing.
//
// Fragment-only links (#heading) are still skipped because they reference
// anchors within the same file and resolving "the same file" requires call-
// site context that the checker does not have.
//
// The checker does NOT extract links. That is the parser's job.
package internal

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Result records whether a single link is broken and, if so, why.
type Result struct {
	Link   Link
	Broken bool
	Reason string // e.g., "HTTP 404" or "file not found"
}

// httpClient is a shared client with a conservative timeout. HTTP HEAD
// requests should be fast; anything beyond 10 seconds is likely broken
// or unreachable.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// CheckLinks verifies each link and returns a Result for every one.
// Broken links have Broken=true and a human-readable Reason.
func CheckLinks(links []Link) []Result {
	results := make([]Result, 0, len(links))

	for _, link := range links {
		result := checkOne(link)
		results = append(results, result)
	}

	return results
}

// checkOne dispatches to the appropriate verification strategy based on
// the link's target.
func checkOne(link Link) Result {
	target := link.Target

	// Skip fragment-only links — they reference anchors within the same file
	if strings.HasPrefix(target, "#") {
		return Result{Link: link, Broken: false}
	}

	// Separate the file path from any fragment so each can be validated
	// independently. The file must exist before we can check its headings.
	filePath := target
	fragment := ""
	if idx := strings.Index(target, "#"); idx > 0 {
		filePath = target[:idx]
		fragment = target[idx+1:]
	}

	if strings.HasPrefix(filePath, "http://") || strings.HasPrefix(filePath, "https://") {
		return checkURL(link, filePath)
	}

	return checkLocalFile(link, filePath, fragment)
}

// checkURL performs an HTTP HEAD request and treats non-2xx responses as broken.
func checkURL(link Link, url string) Result {
	resp, err := httpClient.Head(url)
	if err != nil {
		return Result{
			Link:   link,
			Broken: true,
			Reason: fmt.Sprintf("request failed: %v", err),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return Result{Link: link, Broken: false}
	}

	return Result{
		Link:   link,
		Broken: true,
		Reason: fmt.Sprintf("HTTP %d", resp.StatusCode),
	}
}

// checkLocalFile resolves the path relative to the source Markdown file's
// directory and checks whether the target file exists. When a fragment is
// present, it also verifies the heading exists — a missing heading is a
// different kind of broken than a missing file, and the user needs to know
// which one to fix.
func checkLocalFile(link Link, path string, fragment string) Result {
	// Resolve relative to the directory containing the Markdown file
	dir := filepath.Dir(link.File)
	resolved := filepath.Join(dir, path)

	if _, err := os.Stat(resolved); os.IsNotExist(err) {
		return Result{
			Link:   link,
			Broken: true,
			Reason: "file not found",
		}
	}

	// File exists. If a fragment was specified, verify the heading is present.
	if fragment != "" {
		if !fileHasFragment(resolved, fragment) {
			return Result{
				Link:   link,
				Broken: true,
				Reason: fmt.Sprintf("fragment #%s not found in %s", fragment, path),
			}
		}
	}

	return Result{Link: link, Broken: false}
}

// nonAlphanumericOrHyphen matches characters that GitHub strips when
// converting heading text to URL slugs.
var nonAlphanumericOrHyphen = regexp.MustCompile(`[^a-z0-9 -]`)

// slugify converts heading text to a GitHub-compatible anchor slug.
// GitHub's algorithm: lowercase, strip non-alphanumeric (except spaces
// and hyphens), then replace spaces with hyphens. We replicate this so
// that links authored for GitHub rendering validate correctly here.
func slugify(heading string) string {
	s := strings.ToLower(heading)
	s = nonAlphanumericOrHyphen.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

// fileHasFragment reads a Markdown file line by line, extracts ATX-style
// headings (lines starting with one or more `#`), slugifies each, and
// returns true if any slug matches the target fragment. This is a linear
// scan — acceptable because Markdown files are small and the checker
// already does I/O for every local link.
func fileHasFragment(filePath string, fragment string) bool {
	f, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// ATX headings start with one or more # followed by a space
		if !strings.HasPrefix(line, "#") {
			continue
		}
		// Strip the leading #s and the space
		text := strings.TrimLeft(line, "#")
		if len(text) == 0 || text[0] != ' ' {
			continue
		}
		text = strings.TrimSpace(text)

		if slugify(text) == fragment {
			return true
		}
	}
	return false
}
