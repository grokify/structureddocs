package validation

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Reporter handles formatted output of validation results.
type Reporter struct {
	stdout io.Writer
	stderr io.Writer
}

// NewReporter creates a new Reporter with the given writers.
// If writers are nil, os.Stdout and os.Stderr are used.
func NewReporter(stdout, stderr io.Writer) *Reporter {
	if stdout == nil {
		stdout = os.Stdout
	}
	if stderr == nil {
		stderr = os.Stderr
	}
	return &Reporter{
		stdout: stdout,
		stderr: stderr,
	}
}

// DefaultReporter returns a reporter using os.Stdout and os.Stderr.
func DefaultReporter() *Reporter {
	return NewReporter(os.Stdout, os.Stderr)
}

// Report outputs the validation result in a human-readable format.
func (r *Reporter) Report(result *Result, inputPath string) {
	// Print warnings
	if len(result.Warnings) > 0 {
		fmt.Fprintln(r.stderr, "Warnings:")
		for _, w := range result.Warnings {
			fmt.Fprintf(r.stderr, "  ⚠ %s\n", w.Format())
		}
		fmt.Fprintln(r.stderr)
	}

	// Print errors
	if len(result.Errors) > 0 {
		fmt.Fprintf(r.stderr, "Validation failed for %s:\n", inputPath)
		for _, e := range result.Errors {
			fmt.Fprintf(r.stderr, "  ✗ %s\n", e.Format())
		}
	}

	// Print info
	if len(result.Info) > 0 && result.Valid {
		fmt.Fprintln(r.stdout, "Info:")
		for _, i := range result.Info {
			fmt.Fprintf(r.stdout, "  ℹ %s\n", i.Format())
		}
		fmt.Fprintln(r.stdout)
	}
}

// ReportSummary outputs a summary of the validation result.
func (r *Reporter) ReportSummary(result *Result) {
	var parts []string
	if len(result.Errors) > 0 {
		parts = append(parts, fmt.Sprintf("%d error(s)", len(result.Errors)))
	}
	if len(result.Warnings) > 0 {
		parts = append(parts, fmt.Sprintf("%d warning(s)", len(result.Warnings)))
	}
	if len(parts) == 0 {
		fmt.Fprintln(r.stdout, "Validation passed with no issues.")
		return
	}
	fmt.Fprintf(r.stdout, "Validation summary: %s\n", strings.Join(parts, ", "))
}

// ReportSuccess outputs a success message.
func (r *Reporter) ReportSuccess(docType, inputPath string) {
	fmt.Fprintf(r.stdout, "✓ Valid %s: %s\n", docType, inputPath)
}

// FormatTOON returns the validation result in TOON format (token-efficient).
func FormatTOON(result *Result) string {
	var sb strings.Builder

	sb.WriteString("VALIDATION_RESULT\n")
	if result.Valid {
		sb.WriteString("status: VALID\n")
	} else {
		sb.WriteString("status: INVALID\n")
	}

	if len(result.Errors) > 0 {
		sb.WriteString(fmt.Sprintf("errors: %d\n", len(result.Errors)))
		for _, e := range result.Errors {
			if e.Path != "" {
				sb.WriteString(fmt.Sprintf("  - [%s] %s\n", e.Path, e.Message))
			} else {
				sb.WriteString(fmt.Sprintf("  - %s\n", e.Message))
			}
		}
	}

	if len(result.Warnings) > 0 {
		sb.WriteString(fmt.Sprintf("warnings: %d\n", len(result.Warnings)))
		for _, w := range result.Warnings {
			if w.Path != "" {
				sb.WriteString(fmt.Sprintf("  - [%s] %s\n", w.Path, w.Message))
			} else {
				sb.WriteString(fmt.Sprintf("  - %s\n", w.Message))
			}
		}
	}

	return sb.String()
}
