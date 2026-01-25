package format

import (
	"strings"
	"testing"
)

func TestMarkdownBuilderHeadings(t *testing.T) {
	b := NewMarkdownBuilder()
	b.H1("Title")
	b.H2("Section")
	b.H3("Subsection")

	result := b.String()

	if !strings.Contains(result, "# Title") {
		t.Error("Missing H1")
	}
	if !strings.Contains(result, "## Section") {
		t.Error("Missing H2")
	}
	if !strings.Contains(result, "### Subsection") {
		t.Error("Missing H3")
	}
}

func TestMarkdownBuilderList(t *testing.T) {
	b := NewMarkdownBuilder()
	b.BulletList([]string{"item1", "item2", "item3"})

	result := b.String()

	if !strings.Contains(result, "- item1") {
		t.Error("Missing bullet item 1")
	}
	if !strings.Contains(result, "- item2") {
		t.Error("Missing bullet item 2")
	}
}

func TestMarkdownBuilderNumberedList(t *testing.T) {
	b := NewMarkdownBuilder()
	b.NumberedList([]string{"first", "second", "third"})

	result := b.String()

	if !strings.Contains(result, "1. first") {
		t.Error("Missing numbered item 1")
	}
	if !strings.Contains(result, "2. second") {
		t.Error("Missing numbered item 2")
	}
}

func TestMarkdownBuilderCheckbox(t *testing.T) {
	b := NewMarkdownBuilder()
	b.Checkbox(true, "done task")
	b.Checkbox(false, "pending task")

	result := b.String()

	if !strings.Contains(result, "- [x] done task") {
		t.Error("Missing checked checkbox")
	}
	if !strings.Contains(result, "- [ ] pending task") {
		t.Error("Missing unchecked checkbox")
	}
}

func TestMarkdownBuilderCodeBlock(t *testing.T) {
	b := NewMarkdownBuilder()
	b.CodeBlock("go", "func main() {}")

	result := b.String()

	if !strings.Contains(result, "```go") {
		t.Error("Missing code fence with language")
	}
	if !strings.Contains(result, "func main()") {
		t.Error("Missing code content")
	}
}

func TestMarkdownBuilderTable(t *testing.T) {
	b := NewMarkdownBuilder()
	headers := []string{"Name", "Value"}
	rows := [][]string{
		{"key1", "val1"},
		{"key2", "val2"},
	}
	b.Table(headers, rows)

	result := b.String()

	if !strings.Contains(result, "| Name | Value |") {
		t.Error("Missing table header")
	}
	if !strings.Contains(result, "| --- |") {
		t.Error("Missing table separator")
	}
	if !strings.Contains(result, "| key1 | val1 |") {
		t.Error("Missing table row")
	}
}

func TestMarkdownBuilderLink(t *testing.T) {
	b := NewMarkdownBuilder()
	b.Link("Click here", "https://example.com")

	result := b.String()

	if !strings.Contains(result, "[Click here](https://example.com)") {
		t.Error("Link not formatted correctly")
	}
}

func TestMarkdownBuilderBlockquote(t *testing.T) {
	b := NewMarkdownBuilder()
	b.Blockquote("This is a quote")

	result := b.String()

	if !strings.Contains(result, "> This is a quote") {
		t.Error("Blockquote not formatted correctly")
	}
}

func TestMarkdownEscape(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "hello"},
		{"*bold*", "\\*bold\\*"},
		{"_italic_", "\\_italic\\_"},
		{"`code`", "\\`code\\`"},
		{"[link]", "\\[link\\]"},
	}

	for _, tt := range tests {
		got := Escape(tt.input)
		if got != tt.want {
			t.Errorf("Escape(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestMarkdownBuilderReset(t *testing.T) {
	b := NewMarkdownBuilder()
	b.H1("Title")
	b.Reset()

	if b.String() != "" {
		t.Error("Reset() should clear the builder")
	}
}
