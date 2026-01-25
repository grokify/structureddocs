package validation

import "testing"

func TestNewResult(t *testing.T) {
	r := NewResult()
	if !r.Valid {
		t.Error("NewResult().Valid should be true")
	}
	if len(r.Errors) != 0 {
		t.Error("NewResult().Errors should be empty")
	}
	if len(r.Warnings) != 0 {
		t.Error("NewResult().Warnings should be empty")
	}
}

func TestResultAddError(t *testing.T) {
	r := NewResult()
	r.AddError("metadata.id", "id is required")

	if r.Valid {
		t.Error("Result.Valid should be false after AddError")
	}
	if len(r.Errors) != 1 {
		t.Errorf("Result.Errors should have 1 error, got %d", len(r.Errors))
	}
	if r.Errors[0].Path != "metadata.id" {
		t.Errorf("Error path = %q, want %q", r.Errors[0].Path, "metadata.id")
	}
}

func TestResultAddWarning(t *testing.T) {
	r := NewResult()
	r.AddWarning("metadata.description", "description is recommended")

	if !r.Valid {
		t.Error("Result.Valid should still be true after AddWarning")
	}
	if len(r.Warnings) != 1 {
		t.Errorf("Result.Warnings should have 1 warning, got %d", len(r.Warnings))
	}
}

func TestResultMerge(t *testing.T) {
	r1 := NewResult()
	r1.AddError("field1", "error 1")

	r2 := NewResult()
	r2.AddError("field2", "error 2")
	r2.AddWarning("field3", "warning 1")

	r1.Merge(r2)

	if r1.Valid {
		t.Error("Merged result should be invalid")
	}
	if len(r1.Errors) != 2 {
		t.Errorf("Merged result should have 2 errors, got %d", len(r1.Errors))
	}
	if len(r1.Warnings) != 1 {
		t.Errorf("Merged result should have 1 warning, got %d", len(r1.Warnings))
	}
}

func TestIssueFormat(t *testing.T) {
	tests := []struct {
		issue Issue
		want  string
	}{
		{
			Issue{Path: "metadata.id", Message: "id is required"},
			"[metadata.id] id is required",
		},
		{
			Issue{Message: "general error"},
			"general error",
		},
		{
			Issue{Field: "title", Message: "title is required"},
			"[title] title is required",
		},
	}

	for _, tt := range tests {
		got := tt.issue.Format()
		if got != tt.want {
			t.Errorf("Issue.Format() = %q, want %q", got, tt.want)
		}
	}
}

func TestResultCounts(t *testing.T) {
	r := NewResult()
	r.AddError("f1", "e1")
	r.AddError("f2", "e2")
	r.AddWarning("f3", "w1")

	if r.ErrorCount() != 2 {
		t.Errorf("ErrorCount() = %d, want 2", r.ErrorCount())
	}
	if r.WarningCount() != 1 {
		t.Errorf("WarningCount() = %d, want 1", r.WarningCount())
	}
}
