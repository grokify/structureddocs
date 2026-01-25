# Structured Docs

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Coverage][coverage-svg]][coverage-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

Shared utilities for structured document projects. This package ensures consistency across:

- [structured-requirements](https://github.com/grokify/structured-requirements) - PRD, MRD, TRD
- [structured-goals](https://github.com/grokify/structured-goals) - V2MOM, OKR
- [structured-changelog](https://github.com/grokify/structured-changelog) - Changelogs
- [structured-roadmap](https://github.com/grokify/structured-roadmap) - Roadmaps

## Installation

```bash
go get github.com/grokify/structureddocs
```

## Packages

### `marp` - Marp Slide Rendering

Common utilities for generating Marp presentations from structured documents.

```go
import "github.com/grokify/structureddocs/marp"

// Get a consistent theme
theme := marp.GetTheme("corporate")

// Use common template functions
funcMap := marp.CommonFuncMap

// Render front matter
frontMatter, _ := marp.RenderFrontMatter(marp.FrontMatterData{
    Theme:    theme,
    Title:    "My Document",
    Paginate: true,
})
```

**Features:**

- Standard themes: `default`, `corporate`, `minimal`
- Common template functions: `add`, `truncate`, `progressBar`, `statusIcon`, `priorityIcon`
- Front matter and title slide templates
- Base renderer interface

### `cli` - CLI Utilities

Common patterns for generate, validate, and file handling commands.

```go
import "github.com/grokify/structureddocs/cli"

// Standard generate workflow
err := cli.RunGenerate(cli.GenerateConfig{
    InputPath:  "document.json",
    OutputPath: "document.md",
    OutputExt:  ".md",
    Parser:     myParser,
    Generator:  myGenerator,
})

// Standard validate workflow
err := cli.RunValidate(cli.ValidateConfig{
    InputPath:    "document.json",
    Parser:       myParser,
    Validator:    myValidator,
    DocumentType: "PRD",
})

// Utility functions
outputPath := cli.DeriveOutputPath("input.json", "", ".md") // Returns "input.md"
```

### `validation` - Validation Framework

Consistent error reporting and validation rules across all document types.

```go
import "github.com/grokify/structureddocs/validation"

// Create a validation result
result := validation.NewResult()

// Add errors and warnings
result.AddError("metadata.id", "id is required")
result.AddWarning("metadata.description", "description is recommended")

// Use built-in validation rules
validation.RequiredString(result, "title", doc.Title, "Title")
validation.ValidateSemVer(doc.Version)
validation.ValidateDate(doc.Date)
validation.OneOf(result, "status", doc.Status, []string{"draft", "active"}, "Status")

// Report results
reporter := validation.DefaultReporter()
reporter.Report(result, "document.json")
```

**Built-in validators:**

- `ValidateSemVer` - Semantic version strings
- `ValidateDate` - YYYY-MM-DD dates
- `ValidateCVE` - CVE identifiers
- `ValidateGHSA` - GitHub Security Advisory IDs
- `ValidateURL`, `ValidateEmail`, `ValidateID`

### `metadata` - Common Metadata Types

Shared types for document metadata, status, and priority.

```go
import "github.com/grokify/structureddocs/metadata"

// Common author type
author := metadata.Author{
    Name:  "John Doe",
    Email: "john@example.com",
    Role:  "Product Manager",
}

// Standard statuses
status := metadata.StatusInProgress
fmt.Println(status.Display())    // "In Progress"
fmt.Println(status.IsActive())   // true

// Priority with normalization
priority := metadata.PriorityP0
fmt.Println(priority.Normalize()) // "critical"
fmt.Println(priority.Level())     // 0
```

### `format` - Output Formats

Utilities for JSON, TOON, and Markdown output.

```go
import "github.com/grokify/structureddocs/format"

// TOON builder (token-efficient format for LLMs)
b := format.NewTOONBuilder()
b.Header("DOCUMENT")
b.Field("title", "My Document")
b.Field("version", "1.0.0")
output := b.String()

// Markdown builder
md := format.NewMarkdownBuilder()
md.H1("Title")
md.Paragraph("Introduction text.")
md.BulletList([]string{"Item 1", "Item 2"})
md.Table([]string{"Col1", "Col2"}, [][]string{{"a", "b"}})
output := md.String()

// Format detection
f := format.ParseFormat("json")
ext := f.FileExtension() // ".json"
```

### `schema` - JSON Schema Utilities

Helpers for embedding and validating JSON schemas.

```go
import "github.com/grokify/structureddocs/schema"

// Create a schema registry (typically with //go:embed)
registry := schema.NewRegistry(
    "https://example.com/schema/v1.json",
    "1.0.0",
    schemaBytes,
)

// Get schema as string or bytes
schemaJSON := registry.JSONString()

// Validate $schema reference
err := schema.ValidateSchemaRef(docBytes, registry.ID)
```

## Design Principles

1. **Consistency**: Same patterns across all structured document tools
2. **Composability**: Use what you need, ignore what you don't
3. **Minimal dependencies**: Only standard library (except Cobra for CLI)
4. **Well-tested**: Comprehensive test coverage

## Usage in Structured Document Projects

Each structured document project imports this package for shared functionality:

```go
// In structured-requirements
import (
    "github.com/grokify/structureddocs/marp"
    "github.com/grokify/structureddocs/validation"
)

// Use shared theme
theme := marp.GetTheme("corporate")

// Use shared validation
result := validation.NewResult()
validation.RequiredString(result, "metadata.title", doc.Metadata.Title, "Title")
```

## Contributing

Contributions are welcome. Please ensure:

1. New utilities are general enough for multiple document types
2. Tests are included for new functionality
3. Documentation is updated

## License

MIT License

 [build-status-svg]: https://github.com/grokify/structureddocs/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/structureddocs/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/structureddocs/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/structureddocs/actions/workflows/lint.yaml
 [coverage-svg]: https://img.shields.io/badge/coverage-96.1%25-brightgreen
 [coverage-url]: https://github.com/grokify/structureddocs
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/structureddocs
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/structureddocs
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/structureddocs
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/structureddocs
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fstructureddocs
 [loc-svg]: https://tokei.rs/b1/github/grokify/structureddocs
 [repo-url]: https://github.com/grokify/structureddocs
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/structureddocs/blob/master/LICENSE
