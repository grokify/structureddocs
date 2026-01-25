package cli

import (
	"fmt"
	"os"

	"github.com/grokify/structureddocs/validation"
)

// ValidateConfig configures a validate command execution.
type ValidateConfig struct {
	// InputPath is the path to the input file (required).
	InputPath string

	// Parser parses the input data and returns the document.
	Parser func(data []byte) (any, error)

	// Validator validates the document and returns a result.
	Validator func(doc any) *validation.Result

	// StatsPrinter prints document statistics after successful validation (optional).
	StatsPrinter func(doc any)

	// DocumentType is the type name for display (e.g., "PRD", "V2MOM").
	DocumentType string

	// Quiet suppresses success messages.
	Quiet bool
}

// RunValidate executes a standard validate workflow:
// 1. Read input file
// 2. Parse document
// 3. Validate document
// 4. Report results
// 5. Print stats (optional)
func RunValidate(cfg ValidateConfig) error {
	// Validate config
	if cfg.InputPath == "" {
		return fmt.Errorf("input path is required")
	}
	if cfg.Parser == nil {
		return fmt.Errorf("parser function is required")
	}
	if cfg.Validator == nil {
		return fmt.Errorf("validator function is required")
	}
	if cfg.DocumentType == "" {
		cfg.DocumentType = "Document"
	}

	// Read input file
	data, err := ReadInputFile(cfg.InputPath)
	if err != nil {
		return err
	}

	// Parse document
	doc, err := cfg.Parser(data)
	if err != nil {
		return fmt.Errorf("parsing input: %w", err)
	}

	// Validate document
	result := cfg.Validator(doc)

	// Print warnings
	if len(result.Warnings) > 0 {
		fmt.Fprintln(os.Stderr, "Warnings:")
		for _, w := range result.Warnings {
			fmt.Fprintf(os.Stderr, "  - %s\n", w.Format())
		}
		fmt.Fprintln(os.Stderr)
	}

	// Print errors and return if invalid
	if !result.Valid {
		fmt.Fprintf(os.Stderr, "Validation failed for %s:\n", cfg.InputPath)
		for _, e := range result.Errors {
			fmt.Fprintf(os.Stderr, "  - %s\n", e.Format())
		}
		return fmt.Errorf("validation failed with %d error(s)", len(result.Errors))
	}

	// Print success
	if !cfg.Quiet {
		fmt.Printf("Valid %s: %s\n", cfg.DocumentType, cfg.InputPath)
	}

	// Print stats
	if cfg.StatsPrinter != nil && !cfg.Quiet {
		cfg.StatsPrinter(doc)
	}

	return nil
}

// ValidateFlags contains common flags for validate commands.
type ValidateFlags struct {
	// Quiet suppresses success messages (-q, --quiet).
	Quiet bool

	// Strict enables strict validation mode (--strict).
	Strict bool

	// Format is the output format for validation results (--format).
	Format string
}

// DefaultValidateFlags returns default validate flags.
func DefaultValidateFlags() ValidateFlags {
	return ValidateFlags{
		Format: "text",
	}
}
