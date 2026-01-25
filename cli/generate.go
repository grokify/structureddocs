package cli

import (
	"fmt"
	"os"
)

// GenerateConfig configures a generate command execution.
type GenerateConfig struct {
	// InputPath is the path to the input file (required).
	InputPath string

	// OutputPath is the path to the output file (optional, derived from input if empty).
	OutputPath string

	// OutputExt is the output file extension (e.g., ".md").
	OutputExt string

	// Parser parses the input data and returns the document.
	Parser func(data []byte) (any, error)

	// Generator generates output from the document.
	Generator func(doc any) ([]byte, error)

	// Quiet suppresses success messages.
	Quiet bool
}

// RunGenerate executes a standard generate workflow:
// 1. Read input file
// 2. Parse document
// 3. Generate output
// 4. Write output file
// 5. Print success message
func RunGenerate(cfg GenerateConfig) error {
	// Validate config
	if cfg.InputPath == "" {
		return fmt.Errorf("input path is required")
	}
	if cfg.Parser == nil {
		return fmt.Errorf("parser function is required")
	}
	if cfg.Generator == nil {
		return fmt.Errorf("generator function is required")
	}

	// Determine output path
	outputPath := DeriveOutputPath(cfg.InputPath, cfg.OutputPath, cfg.OutputExt)

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

	// Generate output
	output, err := cfg.Generator(doc)
	if err != nil {
		return fmt.Errorf("generating output: %w", err)
	}

	// Write output
	if err := WriteOutputFile(outputPath, output); err != nil {
		return err
	}

	// Print success
	if !cfg.Quiet {
		fmt.Fprintf(os.Stdout, "Generated: %s\n", outputPath)
	}

	return nil
}

// GenerateFlags contains common flags for generate commands.
type GenerateFlags struct {
	// Output is the output file path (-o, --output).
	Output string

	// Quiet suppresses success messages (-q, --quiet).
	Quiet bool

	// Theme is the theme name for Marp output (--theme).
	Theme string

	// Format is the output format (--format).
	Format string
}

// DefaultGenerateFlags returns default generate flags.
func DefaultGenerateFlags() GenerateFlags {
	return GenerateFlags{
		Theme:  "default",
		Format: "markdown",
	}
}
