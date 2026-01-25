package format

import (
	"fmt"
	"strings"
)

// MarkdownBuilder helps construct Markdown-formatted output.
type MarkdownBuilder struct {
	sb strings.Builder
}

// NewMarkdownBuilder creates a new Markdown builder.
func NewMarkdownBuilder() *MarkdownBuilder {
	return &MarkdownBuilder{}
}

// H1 writes a level 1 heading.
func (b *MarkdownBuilder) H1(text string) {
	b.sb.WriteString("# ")
	b.sb.WriteString(text)
	b.sb.WriteString("\n\n")
}

// H2 writes a level 2 heading.
func (b *MarkdownBuilder) H2(text string) {
	b.sb.WriteString("## ")
	b.sb.WriteString(text)
	b.sb.WriteString("\n\n")
}

// H3 writes a level 3 heading.
func (b *MarkdownBuilder) H3(text string) {
	b.sb.WriteString("### ")
	b.sb.WriteString(text)
	b.sb.WriteString("\n\n")
}

// H4 writes a level 4 heading.
func (b *MarkdownBuilder) H4(text string) {
	b.sb.WriteString("#### ")
	b.sb.WriteString(text)
	b.sb.WriteString("\n\n")
}

// Paragraph writes a paragraph.
func (b *MarkdownBuilder) Paragraph(text string) {
	b.sb.WriteString(text)
	b.sb.WriteString("\n\n")
}

// Text writes text without trailing newlines.
func (b *MarkdownBuilder) Text(text string) {
	b.sb.WriteString(text)
}

// Bold writes bold text.
func (b *MarkdownBuilder) Bold(text string) {
	b.sb.WriteString("**")
	b.sb.WriteString(text)
	b.sb.WriteString("**")
}

// Italic writes italic text.
func (b *MarkdownBuilder) Italic(text string) {
	b.sb.WriteString("*")
	b.sb.WriteString(text)
	b.sb.WriteString("*")
}

// Code writes inline code.
func (b *MarkdownBuilder) Code(text string) {
	b.sb.WriteString("`")
	b.sb.WriteString(text)
	b.sb.WriteString("`")
}

// CodeBlock writes a fenced code block.
func (b *MarkdownBuilder) CodeBlock(language, code string) {
	b.sb.WriteString("```")
	b.sb.WriteString(language)
	b.sb.WriteString("\n")
	b.sb.WriteString(code)
	b.sb.WriteString("\n```\n\n")
}

// Link writes a Markdown link.
func (b *MarkdownBuilder) Link(text, url string) {
	b.sb.WriteString("[")
	b.sb.WriteString(text)
	b.sb.WriteString("](")
	b.sb.WriteString(url)
	b.sb.WriteString(")")
}

// Image writes a Markdown image.
func (b *MarkdownBuilder) Image(alt, url string) {
	b.sb.WriteString("![")
	b.sb.WriteString(alt)
	b.sb.WriteString("](")
	b.sb.WriteString(url)
	b.sb.WriteString(")")
}

// BulletList writes an unordered list.
func (b *MarkdownBuilder) BulletList(items []string) {
	for _, item := range items {
		b.sb.WriteString("- ")
		b.sb.WriteString(item)
		b.sb.WriteString("\n")
	}
	b.sb.WriteString("\n")
}

// NumberedList writes an ordered list.
func (b *MarkdownBuilder) NumberedList(items []string) {
	for i, item := range items {
		b.sb.WriteString(fmt.Sprintf("%d. ", i+1))
		b.sb.WriteString(item)
		b.sb.WriteString("\n")
	}
	b.sb.WriteString("\n")
}

// Checkbox writes a task list item.
func (b *MarkdownBuilder) Checkbox(checked bool, text string) {
	if checked {
		b.sb.WriteString("- [x] ")
	} else {
		b.sb.WriteString("- [ ] ")
	}
	b.sb.WriteString(text)
	b.sb.WriteString("\n")
}

// Blockquote writes a blockquote.
func (b *MarkdownBuilder) Blockquote(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		b.sb.WriteString("> ")
		b.sb.WriteString(line)
		b.sb.WriteString("\n")
	}
	b.sb.WriteString("\n")
}

// HorizontalRule writes a horizontal rule.
func (b *MarkdownBuilder) HorizontalRule() {
	b.sb.WriteString("---\n\n")
}

// Newline writes a newline.
func (b *MarkdownBuilder) Newline() {
	b.sb.WriteString("\n")
}

// Table writes a Markdown table.
func (b *MarkdownBuilder) Table(headers []string, rows [][]string) {
	// Headers
	b.sb.WriteString("| ")
	b.sb.WriteString(strings.Join(headers, " | "))
	b.sb.WriteString(" |\n")

	// Separator
	b.sb.WriteString("|")
	for range headers {
		b.sb.WriteString(" --- |")
	}
	b.sb.WriteString("\n")

	// Rows
	for _, row := range rows {
		b.sb.WriteString("| ")
		b.sb.WriteString(strings.Join(row, " | "))
		b.sb.WriteString(" |\n")
	}
	b.sb.WriteString("\n")
}

// String returns the built Markdown string.
func (b *MarkdownBuilder) String() string {
	return b.sb.String()
}

// Bytes returns the built Markdown as bytes.
func (b *MarkdownBuilder) Bytes() []byte {
	return []byte(b.sb.String())
}

// Reset clears the builder.
func (b *MarkdownBuilder) Reset() {
	b.sb.Reset()
}

// Escape escapes special Markdown characters.
func Escape(s string) string {
	replacer := strings.NewReplacer(
		"\\", "\\\\",
		"`", "\\`",
		"*", "\\*",
		"_", "\\_",
		"{", "\\{",
		"}", "\\}",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		".", "\\.",
		"!", "\\!",
		"|", "\\|",
	)
	return replacer.Replace(s)
}
