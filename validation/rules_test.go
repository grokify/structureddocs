package validation

import "testing"

func TestValidateSemVer(t *testing.T) {
	valid := []string{
		"1.0.0",
		"0.1.0",
		"1.2.3",
		"v1.0.0",
		"1.0.0-alpha",
		"1.0.0-alpha.1",
		"1.0.0+build.123",
		"1.0.0-beta+build",
	}

	invalid := []string{
		"1.0",
		"1",
		"v1",
		"1.0.0.0",
		"a.b.c",
		"",
	}

	for _, v := range valid {
		if !ValidateSemVer(v) {
			t.Errorf("ValidateSemVer(%q) = false, want true", v)
		}
	}

	for _, v := range invalid {
		if ValidateSemVer(v) {
			t.Errorf("ValidateSemVer(%q) = true, want false", v)
		}
	}
}

func TestValidateDate(t *testing.T) {
	valid := []string{
		"2024-01-15",
		"2000-12-31",
		"1999-01-01",
	}

	invalid := []string{
		"2024-13-01",   // Invalid month
		"2024-01-32",   // Invalid day
		"01-15-2024",   // Wrong format
		"2024/01/15",   // Wrong separator
		"2024-1-15",    // Missing leading zero
		"not-a-date",   // Not a date
		"",             // Empty
	}

	for _, v := range valid {
		if !ValidateDate(v) {
			t.Errorf("ValidateDate(%q) = false, want true", v)
		}
	}

	for _, v := range invalid {
		if ValidateDate(v) {
			t.Errorf("ValidateDate(%q) = true, want false", v)
		}
	}
}

func TestValidateCVE(t *testing.T) {
	valid := []string{
		"CVE-2024-12345",
		"CVE-2000-0001",
		"CVE-2024-123456789",
	}

	invalid := []string{
		"CVE-24-12345",    // Year too short
		"CVE-2024-123",    // ID too short
		"cve-2024-12345",  // Wrong case
		"CVE2024-12345",   // Missing hyphen
		"",                // Empty
	}

	for _, v := range valid {
		if !ValidateCVE(v) {
			t.Errorf("ValidateCVE(%q) = false, want true", v)
		}
	}

	for _, v := range invalid {
		if ValidateCVE(v) {
			t.Errorf("ValidateCVE(%q) = true, want false", v)
		}
	}
}

func TestValidateGHSA(t *testing.T) {
	valid := []string{
		"GHSA-abcd-1234-wxyz",
		"GHSA-0000-0000-0000",
		"GHSA-a1b2-c3d4-e5f6",
	}

	invalid := []string{
		"GHSA-abc-1234-wxyz",   // Segment too short
		"GHSA-abcde-1234-wxyz", // Segment too long
		"ghsa-abcd-1234-wxyz",  // Wrong case
		"GHSA-ABCD-1234-WXYZ",  // Uppercase not allowed
		"",                     // Empty
	}

	for _, v := range valid {
		if !ValidateGHSA(v) {
			t.Errorf("ValidateGHSA(%q) = false, want true", v)
		}
	}

	for _, v := range invalid {
		if ValidateGHSA(v) {
			t.Errorf("ValidateGHSA(%q) = true, want false", v)
		}
	}
}

func TestRequiredString(t *testing.T) {
	r := NewResult()
	RequiredString(r, "metadata.id", "", "ID")

	if r.Valid {
		t.Error("Result should be invalid after RequiredString with empty value")
	}
	if len(r.Errors) != 1 {
		t.Errorf("Should have 1 error, got %d", len(r.Errors))
	}

	r2 := NewResult()
	RequiredString(r2, "metadata.id", "some-id", "ID")

	if !r2.Valid {
		t.Error("Result should be valid after RequiredString with non-empty value")
	}
}

func TestOneOf(t *testing.T) {
	allowed := []string{"draft", "review", "approved"}

	r := NewResult()
	OneOf(r, "status", "invalid", allowed, "Status")

	if r.Valid {
		t.Error("Result should be invalid for value not in allowed list")
	}

	r2 := NewResult()
	OneOf(r2, "status", "draft", allowed, "Status")

	if !r2.Valid {
		t.Error("Result should be valid for value in allowed list")
	}
}

func TestInRange(t *testing.T) {
	r := NewResult()
	InRange(r, "score", 1.5, 0.0, 1.0, "Score")

	if r.Valid {
		t.Error("Result should be invalid for value out of range")
	}

	r2 := NewResult()
	InRange(r2, "score", 0.5, 0.0, 1.0, "Score")

	if !r2.Valid {
		t.Error("Result should be valid for value in range")
	}
}
