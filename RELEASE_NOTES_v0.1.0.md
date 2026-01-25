# Release Notes - v0.1.0

**Release Date:** January 25, 2026

## Overview

This is the initial release of **structureddocs**, a Go library providing shared utilities for rendering structured documents as presentations. It serves as the foundation for consistent document rendering across the structured-* ecosystem including structured-goals and structured-requirements.

## Highlights

- **Marp Presentation Rendering** - Generate beautiful Marp-compatible markdown slides from structured documents
- **Theme System** - Three built-in themes (default, corporate, minimal) with customizable colors
- **Template Functions** - Rich set of helper functions for progress bars, status icons, and formatting
- **Validation Framework** - Configurable document validation with rule-based checks

## What's New

### Marp Rendering Utilities

The `marp` package provides:

- `ThemeConfig` for consistent theme configuration
- `GetTheme()` for theme lookup by name
- `CommonFuncMap` with shared template functions

### Template Helper Functions

| Function | Description |
|----------|-------------|
| `progressBar` | ASCII progress bar visualization |
| `progressBarLen` | Progress bar with custom length |
| `progressPercent` | Format progress as percentage |
| `scorePercent` | Format score as percentage |
| `statusIcon` | Icon for status values |
| `statusEmoji` | Emoji for status values |
| `statusLabel` | Human-readable status labels |
| `priorityIcon` | Icon for priority levels |
| `priorityLabel` | Human-readable priority labels |
| `severityIcon` | Icon for severity levels |
| `truncate` | Truncate strings with ellipsis |

### Built-in Themes

| Theme | Primary Color | Use Case |
|-------|--------------|----------|
| `default` | #5a67d8 | General purpose presentations |
| `corporate` | #1a365d | Professional business presentations |
| `minimal` | #2d3748 | Clean, distraction-free slides |

### Validation Framework

The `validation` package provides:

- Configurable validation rules
- Error and warning level reporting
- Extensible rule system

## Installation

```bash
go get github.com/grokify/structureddocs
```

## Usage

```go
import (
    sdmarp "github.com/grokify/structureddocs/marp"
)

// Get a theme
theme := sdmarp.GetTheme("corporate")

// Use common template functions
funcMap := sdmarp.CommonFuncMap
```

## Dependencies

This release has minimal dependencies:

- Go 1.24+
- Standard library only

## Contributors

- John Wang (@grokify)
- Claude Opus 4.5 (Co-Author)

## Links

- [GitHub Repository](https://github.com/grokify/structureddocs)
- [Go Package Documentation](https://pkg.go.dev/github.com/grokify/structureddocs)
- [Changelog](CHANGELOG.md)
