package format

import "testing"

func TestParseFormat(t *testing.T) {
	tests := []struct {
		input string
		want  Format
	}{
		{"toon", FormatTOON},
		{"TOON", FormatTOON},
		{"json", FormatJSON},
		{"JSON", FormatJSON},
		{"json-compact", FormatJSONCompact},
		{"markdown", FormatMarkdown},
		{"md", FormatMarkdown},
		{"text", FormatText},
		{"txt", FormatText},
		{"", FormatText},
		{"unknown", FormatText},
	}

	for _, tt := range tests {
		got := ParseFormat(tt.input)
		if got != tt.want {
			t.Errorf("ParseFormat(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestFormatFileExtension(t *testing.T) {
	tests := []struct {
		format Format
		want   string
	}{
		{FormatJSON, ".json"},
		{FormatJSONCompact, ".json"},
		{FormatMarkdown, ".md"},
		{FormatTOON, ".toon"},
		{FormatText, ".txt"},
	}

	for _, tt := range tests {
		got := tt.format.FileExtension()
		if got != tt.want {
			t.Errorf("Format(%q).FileExtension() = %q, want %q", tt.format, got, tt.want)
		}
	}
}

func TestTOONBuilder(t *testing.T) {
	b := NewTOONBuilder()
	b.Header("test")
	b.Field("key", "value")
	b.Indent()
	b.ListItem("item1")
	b.Dedent()

	result := b.String()

	if result == "" {
		t.Error("TOONBuilder produced empty output")
	}

	// Check contains expected content
	expected := []string{"TEST", "key: value", "- item1"}
	for _, exp := range expected {
		if !contains(result, exp) {
			t.Errorf("TOONBuilder output missing %q", exp)
		}
	}
}

func TestToJSON(t *testing.T) {
	data := map[string]string{"key": "value"}
	result, err := ToJSON(data)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}

	if !contains(string(result), "\"key\"") {
		t.Error("ToJSON output missing key")
	}
	if !contains(string(result), "\"value\"") {
		t.Error("ToJSON output missing value")
	}
}

func TestToJSONCompact(t *testing.T) {
	data := map[string]string{"key": "value"}
	result, err := ToJSONCompact(data)
	if err != nil {
		t.Fatalf("ToJSONCompact() error = %v", err)
	}

	// Compact JSON should not have newlines
	if contains(string(result), "\n") {
		t.Error("ToJSONCompact output should not contain newlines")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
