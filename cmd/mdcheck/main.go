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
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/russmiles/ai-literacy-exemplar/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mdcheck <file-or-directory>\n")
		os.Exit(2)
	}

	target := os.Args[1]
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

	broken := 0
	for _, r := range results {
		if r.Broken {
			broken++
			fmt.Printf("%s:%d: [%s](%s) — %s\n",
				r.Link.File, r.Link.Line, r.Link.Text, r.Link.Target, r.Reason)
		}
	}

	fmt.Printf("\n%d links checked, %d broken\n", len(results), broken)

	if broken > 0 {
		os.Exit(1)
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
