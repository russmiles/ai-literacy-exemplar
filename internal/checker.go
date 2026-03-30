// This file handles link verification — the second half of the mdcheck
// pipeline. Given a list of extracted links, the checker determines which
// are broken and why.
//
// Two verification strategies are used: HTTP HEAD requests for remote URLs
// (minimising bandwidth compared to GET), and file existence checks for
// local paths (resolved relative to the source Markdown file's directory).
//
// Fragment-only links (#heading) are skipped because they reference anchors
// within the same file and would require a full Markdown heading parser to
// validate — a complexity that exceeds this tool's scope.
//
// The checker does NOT extract links. That is the parser's job.
package internal

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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

	// Strip any fragment from the target before checking
	if idx := strings.Index(target, "#"); idx > 0 {
		target = target[:idx]
	}

	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
		return checkURL(link, target)
	}

	return checkLocalFile(link, target)
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
// directory and checks whether the target file exists.
func checkLocalFile(link Link, path string) Result {
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

	return Result{Link: link, Broken: false}
}
