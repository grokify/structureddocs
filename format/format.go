// Package format provides output format utilities for structured documents.
// It includes support for TOON (Token-Oriented Object Notation), JSON, and
// Markdown output formats.
package format

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Format represents an output format type.
type Format string

// Supported output formats.
const (
	FormatTOON        Format = "toon"
	FormatJSON        Format = "json"
	FormatJSONCompact Format = "json-compact"
	FormatMarkdown    Format = "markdown"
	FormatText        Format = "text"
)

// ParseFormat parses a format string into a Format type.
// Returns FormatText for unrecognized formats.
func ParseFormat(s string) Format {
	switch strings.ToLower(s) {
	case "toon":
		return FormatTOON
	case "json":
		return FormatJSON
	case "json-compact", "jsoncompact":
		return FormatJSONCompact
	case "markdown", "md":
		return FormatMarkdown
	case "text", "txt", "":
		return FormatText
	default:
		return FormatText
	}
}

// String returns the format as a string.
func (f Format) String() string {
	return string(f)
}

// FileExtension returns the typical file extension for this format.
func (f Format) FileExtension() string {
	switch f {
	case FormatJSON, FormatJSONCompact:
		return ".json"
	case FormatMarkdown:
		return ".md"
	case FormatTOON:
		return ".toon"
	default:
		return ".txt"
	}
}

// ToJSON marshals data to JSON with indentation.
func ToJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// ToJSONCompact marshals data to JSON without indentation.
func ToJSONCompact(v any) ([]byte, error) {
	return json.Marshal(v)
}

// FromJSON unmarshals JSON data.
func FromJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// TOONBuilder helps construct TOON-formatted output.
type TOONBuilder struct {
	sb     strings.Builder
	indent int
}

// NewTOONBuilder creates a new TOON builder.
func NewTOONBuilder() *TOONBuilder {
	return &TOONBuilder{}
}

// Header writes a section header.
func (b *TOONBuilder) Header(name string) {
	b.sb.WriteString(strings.ToUpper(name))
	b.sb.WriteString("\n")
}

// Field writes a key-value field.
func (b *TOONBuilder) Field(key string, value any) {
	b.writeIndent()
	b.sb.WriteString(key)
	b.sb.WriteString(": ")
	b.sb.WriteString(fmt.Sprint(value))
	b.sb.WriteString("\n")
}

// FieldIf writes a key-value field only if the value is non-empty.
func (b *TOONBuilder) FieldIf(key string, value string) {
	if value != "" {
		b.Field(key, value)
	}
}

// List writes a list of values.
func (b *TOONBuilder) List(items []string) {
	for _, item := range items {
		b.writeIndent()
		b.sb.WriteString("- ")
		b.sb.WriteString(item)
		b.sb.WriteString("\n")
	}
}

// ListItem writes a single list item.
func (b *TOONBuilder) ListItem(item string) {
	b.writeIndent()
	b.sb.WriteString("- ")
	b.sb.WriteString(item)
	b.sb.WriteString("\n")
}

// Indent increases the indentation level.
func (b *TOONBuilder) Indent() {
	b.indent++
}

// Dedent decreases the indentation level.
func (b *TOONBuilder) Dedent() {
	if b.indent > 0 {
		b.indent--
	}
}

// Newline writes a blank line.
func (b *TOONBuilder) Newline() {
	b.sb.WriteString("\n")
}

// String returns the built TOON string.
func (b *TOONBuilder) String() string {
	return b.sb.String()
}

// Bytes returns the built TOON as bytes.
func (b *TOONBuilder) Bytes() []byte {
	return []byte(b.sb.String())
}

func (b *TOONBuilder) writeIndent() {
	for i := 0; i < b.indent; i++ {
		b.sb.WriteString("  ")
	}
}
