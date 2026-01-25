package marp

import (
	"bytes"
	"html/template"
)

// Renderer is the interface that all Marp renderers should implement.
type Renderer interface {
	// Format returns the output format name (always "marp" for Marp renderers).
	Format() string

	// FileExtension returns the output file extension (always ".md" for Marp).
	FileExtension() string
}

// Options configures Marp rendering behavior.
type Options struct {
	// Theme is the theme name ("default", "corporate", "minimal").
	Theme string

	// IncludeNotes includes speaker notes in the output.
	IncludeNotes bool

	// Terminology controls label rendering (for multi-framework documents).
	// Example values: "v2mom", "okr", "hybrid".
	Terminology string

	// CustomFuncs are additional template functions to include.
	CustomFuncs template.FuncMap
}

// DefaultOptions returns the default Marp rendering options.
func DefaultOptions() *Options {
	return &Options{
		Theme:       "default",
		Terminology: "",
	}
}

// BaseRenderer provides common Marp renderer functionality.
// Embed this in your document-specific renderer.
type BaseRenderer struct{}

// Format returns "marp".
func (r *BaseRenderer) Format() string {
	return "marp"
}

// FileExtension returns ".md".
func (r *BaseRenderer) FileExtension() string {
	return ".md"
}

// GetFuncMap returns a combined FuncMap with CommonFuncMap and any custom functions.
func GetFuncMap(custom template.FuncMap) template.FuncMap {
	result := make(template.FuncMap)
	for k, v := range CommonFuncMap {
		result[k] = v
	}
	for k, v := range custom {
		result[k] = v
	}
	return result
}

// ExecuteTemplate is a helper that executes a template and returns the result as a string.
func ExecuteTemplate(tmpl *template.Template, data any) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ExecuteTemplateToBuffer is a helper that executes a template into an existing buffer.
func ExecuteTemplateToBuffer(buf *bytes.Buffer, tmpl *template.Template, data any) error {
	return tmpl.Execute(buf, data)
}
