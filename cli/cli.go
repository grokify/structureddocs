// Package cli provides shared CLI utilities for structured document tools.
// It defines common patterns for generate, validate, and init commands
// to ensure consistency across all structured document CLIs.
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ExitCode represents standard CLI exit codes.
type ExitCode int

const (
	// ExitSuccess indicates successful execution.
	ExitSuccess ExitCode = 0

	// ExitError indicates a general error.
	ExitError ExitCode = 1

	// ExitValidationError indicates validation failures.
	ExitValidationError ExitCode = 2

	// ExitUsageError indicates incorrect usage.
	ExitUsageError ExitCode = 64
)

// DeriveOutputPath generates an output path from an input path by changing the extension.
// If outputPath is non-empty, it is returned as-is.
// Otherwise, the input extension is replaced with newExt.
func DeriveOutputPath(inputPath, outputPath, newExt string) string {
	if outputPath != "" {
		return outputPath
	}
	ext := filepath.Ext(inputPath)
	base := strings.TrimSuffix(inputPath, ext)
	return base + newExt
}

// EnsureOutputDir creates the parent directory for the output path if it doesn't exist.
func EnsureOutputDir(outputPath string) error {
	dir := filepath.Dir(outputPath)
	if dir == "." || dir == "" {
		return nil
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating output directory: %w", err)
	}
	return nil
}

// FileExists returns true if the file exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// ReadInputFile reads a file and returns its contents.
// It provides consistent error messages across all CLI tools.
func ReadInputFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %s", path)
		}
		return nil, fmt.Errorf("reading file: %w", err)
	}
	return data, nil
}

// WriteOutputFile writes data to a file with consistent permissions.
// It creates parent directories if needed.
func WriteOutputFile(path string, data []byte) error {
	if err := EnsureOutputDir(path); err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	return nil
}
