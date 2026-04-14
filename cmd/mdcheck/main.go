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
