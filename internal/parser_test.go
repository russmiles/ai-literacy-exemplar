package internal

import (
	"testing"
)

func TestParseInlineLinks(t *testing.T) {
	content := `# Test Document

Check [example](https://example.com) and [local](./README.md).
`
	links := ParseLinks(content, "test.md")

	if len(links) != 2 {
		t.Fatalf("expected 2 links, got %d", len(links))
	}

	if links[0].Text != "example" || links[0].Target != "https://example.com" {
		t.Errorf("first link: got text=%q target=%q", links[0].Text, links[0].Target)
	}
	if links[0].LinkType != "inline" {
		t.Errorf("first link type: got %q, want %q", links[0].LinkType, "inline")
	}
	if links[0].Line != 3 {
		t.Errorf("first link line: got %d, want 3", links[0].Line)
	}

	if links[1].Text != "local" || links[1].Target != "./README.md" {
		t.Errorf("second link: got text=%q target=%q", links[1].Text, links[1].Target)
	}
}

func TestParseAutolinks(t *testing.T) {
	content := `Visit <https://example.com> for more.`

	links := ParseLinks(content, "test.md")

	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}
	if links[0].Target != "https://example.com" {
		t.Errorf("target: got %q", links[0].Target)
	}
	if links[0].LinkType != "autolink" {
		t.Errorf("type: got %q, want %q", links[0].LinkType, "autolink")
	}
}

func TestParseReferenceLinks(t *testing.T) {
	content := `See [the docs][docs] for details.

[docs]: https://docs.example.com
`
	links := ParseLinks(content, "test.md")

	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}
	if links[0].Text != "the docs" {
		t.Errorf("text: got %q", links[0].Text)
	}
	if links[0].Target != "https://docs.example.com" {
		t.Errorf("target: got %q", links[0].Target)
	}
	if links[0].LinkType != "reference" {
		t.Errorf("type: got %q, want %q", links[0].LinkType, "reference")
	}
}

func TestParseImplicitReferenceLinks(t *testing.T) {
	content := `See [docs][] for details.

[docs]: https://docs.example.com
`
	links := ParseLinks(content, "test.md")

	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}
	if links[0].Target != "https://docs.example.com" {
		t.Errorf("target: got %q", links[0].Target)
	}
}

func TestParseMixedLinkTypes(t *testing.T) {
	content := `# Mixed

An [inline](https://a.com) link, a [ref][r] link, and <https://b.com>.

[r]: https://c.com
`
	links := ParseLinks(content, "test.md")

	if len(links) != 3 {
		t.Fatalf("expected 3 links, got %d", len(links))
	}

	types := map[string]bool{}
	for _, l := range links {
		types[l.LinkType] = true
	}
	for _, expected := range []string{"inline", "reference", "autolink"} {
		if !types[expected] {
			t.Errorf("missing link type: %s", expected)
		}
	}
}

func TestParseRecordsFileAndLine(t *testing.T) {
	content := `line one
line two [link](https://example.com)
line three
`
	links := ParseLinks(content, "path/to/file.md")

	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(links))
	}
	if links[0].File != "path/to/file.md" {
		t.Errorf("file: got %q", links[0].File)
	}
	if links[0].Line != 2 {
		t.Errorf("line: got %d, want 2", links[0].Line)
	}
}

func TestParseEmptyContent(t *testing.T) {
	links := ParseLinks("", "test.md")
	if len(links) != 0 {
		t.Errorf("expected 0 links, got %d", len(links))
	}
}

func TestParseNoLinks(t *testing.T) {
	content := `# Just a heading

Some text with no links at all.
`
	links := ParseLinks(content, "test.md")
	if len(links) != 0 {
		t.Errorf("expected 0 links, got %d", len(links))
	}
}
