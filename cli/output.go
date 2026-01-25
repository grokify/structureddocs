package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Printer handles formatted output to stdout and stderr.
type Printer struct {
	stdout io.Writer
	stderr io.Writer
	quiet  bool
}

// NewPrinter creates a new Printer with the given writers.
// If writers are nil, os.Stdout and os.Stderr are used.
func NewPrinter(stdout, stderr io.Writer) *Printer {
	if stdout == nil {
		stdout = os.Stdout
	}
	if stderr == nil {
		stderr = os.Stderr
	}
	return &Printer{
		stdout: stdout,
		stderr: stderr,
	}
}

// DefaultPrinter returns a printer using os.Stdout and os.Stderr.
func DefaultPrinter() *Printer {
	return NewPrinter(os.Stdout, os.Stderr)
}

// SetQuiet enables or disables quiet mode.
func (p *Printer) SetQuiet(quiet bool) {
	p.quiet = quiet
}

// Success prints a success message to stdout.
// Suppressed in quiet mode.
func (p *Printer) Success(format string, args ...any) {
	if p.quiet {
		return
	}
	fmt.Fprintf(p.stdout, format+"\n", args...)
}

// Info prints an informational message to stdout.
// Suppressed in quiet mode.
func (p *Printer) Info(format string, args ...any) {
	if p.quiet {
		return
	}
	fmt.Fprintf(p.stdout, format+"\n", args...)
}

// Warning prints a warning message to stderr.
// Not suppressed in quiet mode.
func (p *Printer) Warning(format string, args ...any) {
	fmt.Fprintf(p.stderr, "Warning: "+format+"\n", args...)
}

// Error prints an error message to stderr.
// Not suppressed in quiet mode.
func (p *Printer) Error(format string, args ...any) {
	fmt.Fprintf(p.stderr, "Error: "+format+"\n", args...)
}

// List prints a bulleted list to stdout.
// Suppressed in quiet mode.
func (p *Printer) List(items []string) {
	if p.quiet {
		return
	}
	for _, item := range items {
		fmt.Fprintf(p.stdout, "  - %s\n", item)
	}
}

// Stats prints a labeled statistics list to stdout.
// Suppressed in quiet mode.
func (p *Printer) Stats(stats map[string]any) {
	if p.quiet {
		return
	}
	for k, v := range stats {
		fmt.Fprintf(p.stdout, "  %s: %v\n", k, v)
	}
}

// Table prints a simple table to stdout.
// Suppressed in quiet mode.
func (p *Printer) Table(headers []string, rows [][]string) {
	if p.quiet {
		return
	}

	// Calculate column widths
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Build format string
	formats := make([]string, len(widths))
	for i, w := range widths {
		formats[i] = fmt.Sprintf("%%-%ds", w)
	}
	format := strings.Join(formats, " | ")

	// Print headers
	headerArgs := make([]any, len(headers))
	for i, h := range headers {
		headerArgs[i] = h
	}
	fmt.Fprintf(p.stdout, format+"\n", headerArgs...)

	// Print separator
	seps := make([]string, len(widths))
	for i, w := range widths {
		seps[i] = strings.Repeat("-", w)
	}
	fmt.Fprintf(p.stdout, "%s\n", strings.Join(seps, "-+-"))

	// Print rows
	for _, row := range rows {
		rowArgs := make([]any, len(widths))
		for i := range widths {
			if i < len(row) {
				rowArgs[i] = row[i]
			} else {
				rowArgs[i] = ""
			}
		}
		fmt.Fprintf(p.stdout, format+"\n", rowArgs...)
	}
}
