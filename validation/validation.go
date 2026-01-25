// Package validation provides shared validation utilities for structured document projects.
// It defines common error types, severity levels, and result structures to ensure
// consistent validation behavior and reporting across all document types.
package validation

import (
	"fmt"
	"strings"
)

// Severity represents the severity level of a validation issue.
type Severity string

const (
	// SeverityError indicates a validation error that must be fixed.
	SeverityError Severity = "error"

	// SeverityWarning indicates a validation warning that should be reviewed.
	SeverityWarning Severity = "warning"

	// SeverityInfo indicates informational validation feedback.
	SeverityInfo Severity = "info"
)

// Issue represents a validation issue (error, warning, or info).
type Issue struct {
	// Path is the JSON path to the problematic field (e.g., "metadata.title", "items[0].id").
	Path string

	// Field is the field name (deprecated, use Path).
	Field string

	// Message describes the validation issue.
	Message string

	// Severity is the issue severity level.
	Severity Severity

	// Code is an optional error code for programmatic handling.
	Code string
}

// Format returns a formatted string representation of the issue.
func (i Issue) Format() string {
	location := i.Path
	if location == "" {
		location = i.Field
	}
	if location != "" {
		return fmt.Sprintf("[%s] %s", location, i.Message)
	}
	return i.Message
}

// FormatWithSeverity returns a formatted string with severity prefix.
func (i Issue) FormatWithSeverity() string {
	return fmt.Sprintf("%s: %s", strings.ToUpper(string(i.Severity)), i.Format())
}

// Error implements the error interface.
func (i Issue) Error() string {
	return i.Format()
}

// Result represents the outcome of a validation operation.
type Result struct {
	// Valid is true if validation passed (no errors).
	Valid bool

	// Errors contains validation errors.
	Errors []Issue

	// Warnings contains validation warnings.
	Warnings []Issue

	// Info contains informational messages.
	Info []Issue
}

// NewResult creates a new empty validation result.
func NewResult() *Result {
	return &Result{
		Valid:    true,
		Errors:   []Issue{},
		Warnings: []Issue{},
		Info:     []Issue{},
	}
}

// AddError adds a validation error and marks the result as invalid.
func (r *Result) AddError(path, message string) {
	r.Valid = false
	r.Errors = append(r.Errors, Issue{
		Path:     path,
		Message:  message,
		Severity: SeverityError,
	})
}

// AddErrorf adds a formatted validation error.
func (r *Result) AddErrorf(path, format string, args ...any) {
	r.AddError(path, fmt.Sprintf(format, args...))
}

// AddErrorWithCode adds a validation error with an error code.
func (r *Result) AddErrorWithCode(path, code, message string) {
	r.Valid = false
	r.Errors = append(r.Errors, Issue{
		Path:     path,
		Message:  message,
		Severity: SeverityError,
		Code:     code,
	})
}

// AddWarning adds a validation warning.
func (r *Result) AddWarning(path, message string) {
	r.Warnings = append(r.Warnings, Issue{
		Path:     path,
		Message:  message,
		Severity: SeverityWarning,
	})
}

// AddWarningf adds a formatted validation warning.
func (r *Result) AddWarningf(path, format string, args ...any) {
	r.AddWarning(path, fmt.Sprintf(format, args...))
}

// AddInfo adds an informational message.
func (r *Result) AddInfo(path, message string) {
	r.Info = append(r.Info, Issue{
		Path:     path,
		Message:  message,
		Severity: SeverityInfo,
	})
}

// Merge combines another result into this one.
func (r *Result) Merge(other *Result) {
	if other == nil {
		return
	}
	if !other.Valid {
		r.Valid = false
	}
	r.Errors = append(r.Errors, other.Errors...)
	r.Warnings = append(r.Warnings, other.Warnings...)
	r.Info = append(r.Info, other.Info...)
}

// HasWarnings returns true if there are any warnings.
func (r *Result) HasWarnings() bool {
	return len(r.Warnings) > 0
}

// HasInfo returns true if there are any info messages.
func (r *Result) HasInfo() bool {
	return len(r.Info) > 0
}

// ErrorCount returns the number of errors.
func (r *Result) ErrorCount() int {
	return len(r.Errors)
}

// WarningCount returns the number of warnings.
func (r *Result) WarningCount() int {
	return len(r.Warnings)
}

// AllIssues returns all issues (errors, warnings, info) combined.
func (r *Result) AllIssues() []Issue {
	all := make([]Issue, 0, len(r.Errors)+len(r.Warnings)+len(r.Info))
	all = append(all, r.Errors...)
	all = append(all, r.Warnings...)
	all = append(all, r.Info...)
	return all
}
