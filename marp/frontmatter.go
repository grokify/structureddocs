package marp

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

// FrontMatterData contains the data needed to render Marp front matter.
type FrontMatterData struct {
	// Theme is the ThemeConfig to use for styling.
	Theme ThemeConfig

	// Title is displayed in the header (optional).
	Title string

	// Subtitle is displayed below the title (optional).
	Subtitle string

	// Footer content (e.g., "Document ID | Version").
	Footer string

	// Paginate enables slide numbers.
	Paginate bool

	// CustomCSS is additional CSS to include in the style block.
	CustomCSS string
}

// frontMatterTemplate is the base Marp front matter template.
var frontMatterTemplate = template.Must(template.New("frontMatter").Parse(`---
marp: true
theme: {{.Theme.Name}}
paginate: {{.Paginate}}
{{- if .Title}}
header: "{{.Title}}{{if .Subtitle}} | {{.Subtitle}}{{end}}"
{{- end}}
{{- if .Footer}}
footer: "{{.Footer}}"
{{- end}}
style: |
  section {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  }
  section.title {
    text-align: center;
    background: linear-gradient(135deg, {{.Theme.PrimaryBgColor}} 0%, {{.Theme.AccentColor}} 100%);
    color: {{.Theme.PrimaryTextColor}};
  }
  section.title h1 {
    font-size: 2.5em;
    margin-bottom: 0.5em;
  }
  section.section-header {
    background: {{.Theme.PrimaryBgColor}};
    color: {{.Theme.PrimaryTextColor}};
  }
  h1 {
    color: {{.Theme.PrimaryBgColor}};
    border-bottom: 2px solid {{.Theme.AccentColor}};
    padding-bottom: 0.3em;
  }
  h2 {
    color: {{.Theme.PrimaryBgColor}};
  }
  table {
    font-size: 0.85em;
    width: 100%;
  }
  th {
    background: {{.Theme.PrimaryBgColor}};
    color: {{.Theme.PrimaryTextColor}};
  }
  .success { color: {{.Theme.SuccessColor}}; }
  .warning { color: {{.Theme.WarningColor}}; }
  .danger { color: {{.Theme.DangerColor}}; }
  .muted { color: #718096; }
  .small { font-size: 0.8em; }
  blockquote {
    border-left: 4px solid {{.Theme.AccentColor}};
    padding-left: 1em;
    color: #4a5568;
    font-style: italic;
  }
{{- if .CustomCSS}}
{{.CustomCSS}}
{{- end}}
---
`))

// RenderFrontMatter generates the Marp front matter YAML block.
func RenderFrontMatter(data FrontMatterData) (string, error) {
	var buf bytes.Buffer
	if err := frontMatterTemplate.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("rendering front matter: %w", err)
	}
	return buf.String(), nil
}

// TitleSlideData contains data for rendering a title slide.
type TitleSlideData struct {
	// Title is the main document title.
	Title string

	// DocumentType describes the document (e.g., "Product Requirements Document").
	DocumentType string

	// Author is the document author name.
	Author string

	// AuthorRole is the author's role (optional).
	AuthorRole string

	// Version is the document version.
	Version string

	// Status is the document status (e.g., "Draft", "Approved").
	Status string

	// Date is the document date. If zero, uses current date.
	Date time.Time
}

// titleSlideTemplate is the standard title slide template.
var titleSlideTemplate = template.Must(template.New("titleSlide").Parse(`<!-- _class: title -->

# {{.Title}}

{{- if .DocumentType}}

**{{.DocumentType}}**
{{- end}}

{{- if .Author}}

**Author:** {{.Author}}{{if .AuthorRole}} ({{.AuthorRole}}){{end}}
{{- end}}
{{- if or .Version .Status}}

{{- if .Version}}**Version:** {{.Version}}{{end}}{{if and .Version .Status}} | {{end}}{{if .Status}}**Status:** {{.Status}}{{end}}
{{- end}}

**Date:** {{.FormattedDate}}

---
`))

// FormattedDate returns the date formatted for display.
func (d TitleSlideData) FormattedDate() string {
	date := d.Date
	if date.IsZero() {
		date = time.Now()
	}
	return date.Format("January 2, 2006")
}

// RenderTitleSlide generates a standard title slide.
func RenderTitleSlide(data TitleSlideData) (string, error) {
	var buf bytes.Buffer
	if err := titleSlideTemplate.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("rendering title slide: %w", err)
	}
	return buf.String(), nil
}

// SlideSeparator is the Marp slide separator.
const SlideSeparator = "\n---\n"

// SectionHeaderClass is the CSS class for section header slides.
const SectionHeaderClass = "<!-- _class: section-header -->"
