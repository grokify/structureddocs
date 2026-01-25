// Package marp provides shared Marp slide rendering utilities for structured document projects.
package marp

// ThemeConfig defines the color scheme and styling for Marp presentations.
type ThemeConfig struct {
	// Name is the Marp theme name (e.g., "default", "gaia", "uncover").
	Name string

	// PrimaryBgColor is the background color for title/header slides.
	PrimaryBgColor string

	// PrimaryTextColor is the text color on primary backgrounds.
	PrimaryTextColor string

	// AccentColor is used for highlights, links, and decorative elements.
	AccentColor string

	// SuccessColor indicates positive status (completed, on-track).
	SuccessColor string

	// WarningColor indicates caution status (at-risk, behind).
	WarningColor string

	// DangerColor indicates negative status (blocked, critical).
	DangerColor string
}

// DefaultThemes contains the standard theme configurations used across
// all structured document projects for consistency.
var DefaultThemes = map[string]ThemeConfig{
	"default": {
		Name:             "default",
		PrimaryBgColor:   "#4c51bf",
		PrimaryTextColor: "#ffffff",
		AccentColor:      "#667eea",
		SuccessColor:     "#38a169",
		WarningColor:     "#d69e2e",
		DangerColor:      "#e53e3e",
	},
	"corporate": {
		Name:             "default",
		PrimaryBgColor:   "#1a365d",
		PrimaryTextColor: "#ffffff",
		AccentColor:      "#2b6cb0",
		SuccessColor:     "#38a169",
		WarningColor:     "#d69e2e",
		DangerColor:      "#e53e3e",
	},
	"minimal": {
		Name:             "default",
		PrimaryBgColor:   "#2d3748",
		PrimaryTextColor: "#ffffff",
		AccentColor:      "#4a5568",
		SuccessColor:     "#48bb78",
		WarningColor:     "#ecc94b",
		DangerColor:      "#fc8181",
	},
}

// GetTheme returns the ThemeConfig for the given theme name.
// If the theme is not found, it returns the default theme.
func GetTheme(name string) ThemeConfig {
	if theme, ok := DefaultThemes[name]; ok {
		return theme
	}
	return DefaultThemes["default"]
}

// ThemeNames returns the list of available theme names.
func ThemeNames() []string {
	return []string{"default", "corporate", "minimal"}
}
