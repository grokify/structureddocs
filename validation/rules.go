package validation

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Common validation patterns used across structured documents.
var (
	// SemVerPattern matches semantic version strings (e.g., "1.2.3", "0.1.0-alpha").
	SemVerPattern = regexp.MustCompile(`^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

	// DatePattern matches YYYY-MM-DD date strings.
	DatePattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	// CVEPattern matches CVE identifiers (e.g., "CVE-2024-12345").
	CVEPattern = regexp.MustCompile(`^CVE-\d{4}-\d{4,}$`)

	// GHSAPattern matches GitHub Security Advisory identifiers.
	GHSAPattern = regexp.MustCompile(`^GHSA-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}$`)

	// URLPattern matches URLs (basic pattern).
	URLPattern = regexp.MustCompile(`^https?://[^\s]+$`)

	// EmailPattern matches email addresses (basic pattern).
	EmailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// IDPattern matches valid identifier strings (alphanumeric with hyphens/underscores).
	IDPattern = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]*$`)
)

// ValidateSemVer checks if a string is a valid semantic version.
func ValidateSemVer(version string) bool {
	return SemVerPattern.MatchString(version)
}

// ValidateDate checks if a string is a valid YYYY-MM-DD date.
func ValidateDate(date string) bool {
	if !DatePattern.MatchString(date) {
		return false
	}
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// ValidateCVE checks if a string is a valid CVE identifier.
func ValidateCVE(cve string) bool {
	return CVEPattern.MatchString(cve)
}

// ValidateGHSA checks if a string is a valid GHSA identifier.
func ValidateGHSA(ghsa string) bool {
	return GHSAPattern.MatchString(ghsa)
}

// ValidateURL checks if a string is a valid URL.
func ValidateURL(url string) bool {
	return URLPattern.MatchString(url)
}

// ValidateEmail checks if a string is a valid email address.
func ValidateEmail(email string) bool {
	return EmailPattern.MatchString(email)
}

// ValidateID checks if a string is a valid identifier.
func ValidateID(id string) bool {
	return IDPattern.MatchString(id)
}

// RequiredString validates that a string is non-empty.
func RequiredString(r *Result, path, value, fieldName string) bool {
	if strings.TrimSpace(value) == "" {
		r.AddError(path, fieldName+" is required")
		return false
	}
	return true
}

// RequiredStrings validates that all strings in a list are non-empty.
func RequiredStrings(r *Result, basePath string, values []string, fieldName string) bool {
	valid := true
	for i, v := range values {
		if strings.TrimSpace(v) == "" {
			r.AddErrorf(fmt.Sprintf("%s[%d]", basePath, i), "%s cannot be empty", fieldName)
			valid = false
		}
	}
	return valid
}

// UniqueStrings validates that all strings in a list are unique.
func UniqueStrings(r *Result, path string, values []string, fieldName string) bool {
	seen := make(map[string]int)
	valid := true
	for i, v := range values {
		if prev, ok := seen[v]; ok {
			r.AddErrorf(path, "duplicate %s '%s' at indices %d and %d", fieldName, v, prev, i)
			valid = false
		}
		seen[v] = i
	}
	return valid
}

// OneOf validates that a value is one of the allowed values.
func OneOf(r *Result, path, value string, allowed []string, fieldName string) bool {
	for _, a := range allowed {
		if value == a {
			return true
		}
	}
	r.AddErrorf(path, "%s must be one of: %s (got: %s)", fieldName, strings.Join(allowed, ", "), value)
	return false
}

// InRange validates that a number is within a range.
func InRange(r *Result, path string, value, min, max float64, fieldName string) bool {
	if value < min || value > max {
		r.AddErrorf(path, "%s must be between %.2f and %.2f (got: %.2f)", fieldName, min, max, value)
		return false
	}
	return true
}

// MaxLength validates that a string does not exceed a maximum length.
func MaxLength(r *Result, path, value string, max int, fieldName string) bool {
	if len(value) > max {
		r.AddErrorf(path, "%s exceeds maximum length of %d characters (got: %d)", fieldName, max, len(value))
		return false
	}
	return true
}

// MinLength validates that a string meets a minimum length.
func MinLength(r *Result, path, value string, min int, fieldName string) bool {
	if len(value) < min {
		r.AddErrorf(path, "%s must be at least %d characters (got: %d)", fieldName, min, len(value))
		return false
	}
	return true
}
