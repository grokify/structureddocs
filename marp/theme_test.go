package marp

import "testing"

func TestGetTheme(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"default", "#5a67d8"},
		{"corporate", "#1a365d"},
		{"minimal", "#2d3748"},
		{"unknown", "#5a67d8"}, // Falls back to default
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theme := GetTheme(tt.name)
			if theme.PrimaryBgColor != tt.expected {
				t.Errorf("GetTheme(%q).PrimaryBgColor = %q, want %q",
					tt.name, theme.PrimaryBgColor, tt.expected)
			}
		})
	}
}

func TestThemeNames(t *testing.T) {
	names := ThemeNames()
	if len(names) != 3 {
		t.Errorf("ThemeNames() returned %d names, want 3", len(names))
	}

	expected := map[string]bool{"default": true, "corporate": true, "minimal": true}
	for _, name := range names {
		if !expected[name] {
			t.Errorf("ThemeNames() contains unexpected name: %q", name)
		}
	}
}

func TestThemeConfigFields(t *testing.T) {
	theme := GetTheme("corporate")

	if theme.Name == "" {
		t.Error("Theme.Name is empty")
	}
	if theme.PrimaryBgColor == "" {
		t.Error("Theme.PrimaryBgColor is empty")
	}
	if theme.PrimaryTextColor == "" {
		t.Error("Theme.PrimaryTextColor is empty")
	}
	if theme.AccentColor == "" {
		t.Error("Theme.AccentColor is empty")
	}
	if theme.SuccessColor == "" {
		t.Error("Theme.SuccessColor is empty")
	}
	if theme.WarningColor == "" {
		t.Error("Theme.WarningColor is empty")
	}
	if theme.DangerColor == "" {
		t.Error("Theme.DangerColor is empty")
	}
}
