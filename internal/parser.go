// Package internal provides the core library for mdcheck — a Markdown link
// checker. This file handles link extraction from Markdown text.
//
// The parser reads Markdown line by line and extracts three kinds of links:
// inline links [text](url), reference links [text][ref] with [ref]: url
// definitions, and autolinks <url>. It deliberately does not build a full
// Markdown AST — regex plus a two-pass approach (extract references first,
// then resolve them) is sufficient for link extraction and avoids pulling
// in a heavyweight parsing dependency.
//
// The parser does NOT check whether links are valid. That is the checker's
// job. The parser only extracts and locates.
package internal

import (
	"bufio"
	"regexp"
	"strings"
)

// Link represents a single link extracted from a Markdown file.
type Link struct {
	Text     string // display text (e.g., "example")
	Target   string // URL or file path
	File     string // source Markdown file path
	Line     int    // 1-based line number in the source file
	LinkType string // "inline", "reference", or "autolink"
}

// Three patterns cover the Markdown link types we extract. Each is compiled
// once and reused across all parsing calls.
var (
	// Inline links: [text](target)
	inlineRe = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

	// Reference links: [text][ref] — the ref may be empty (implicit)
	referenceRe = regexp.MustCompile(`\[([^\]]+)\]\[([^\]]*)\]`)

	// Reference definitions: [ref]: url
	definitionRe = regexp.MustCompile(`^\s*\[([^\]]+)\]:\s+(.+)$`)

	// Autolinks: <https://example.com>
	autolinkRe = regexp.MustCompile(`<(https?://[^>]+)>`)
)

// ParseLinks extracts all links from Markdown content. The file parameter
// is recorded in each Link for reporting purposes — the parser does not
// read files itself.
func ParseLinks(content string, file string) []Link {
	var links []Link

	// First pass: collect reference definitions so we can resolve
	// reference links in the second pass.
	definitions := collectDefinitions(content)

	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// Extract inline links: [text](target)
		for _, match := range inlineRe.FindAllStringSubmatch(line, -1) {
			links = append(links, Link{
				Text:     match[1],
				Target:   match[2],
				File:     file,
				Line:     lineNum,
				LinkType: "inline",
			})
		}

		// Extract autolinks: <url>
		for _, match := range autolinkRe.FindAllStringSubmatch(line, -1) {
			links = append(links, Link{
				Text:     match[1],
				Target:   match[1],
				File:     file,
				Line:     lineNum,
				LinkType: "autolink",
			})
		}

		// Extract reference links: [text][ref]
		for _, match := range referenceRe.FindAllStringSubmatch(line, -1) {
			ref := match[2]
			if ref == "" {
				ref = match[1] // implicit reference: [text][] uses text as ref
			}
			ref = strings.ToLower(ref)
			if target, ok := definitions[ref]; ok {
				links = append(links, Link{
					Text:     match[1],
					Target:   target,
					File:     file,
					Line:     lineNum,
					LinkType: "reference",
				})
			}
		}
	}

	return links
}

// collectDefinitions scans the content for reference definitions of the
// form [ref]: url and returns a map from lowercase ref to url.
func collectDefinitions(content string) map[string]string {
	defs := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		if match := definitionRe.FindStringSubmatch(scanner.Text()); match != nil {
			defs[strings.ToLower(match[1])] = strings.TrimSpace(match[2])
		}
	}
	return defs
}
