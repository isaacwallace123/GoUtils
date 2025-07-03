package uuidutil

import (
	"strings"
	"testing"
)

// TestGenerateAndValidate checks generated UUIDs for correct format and casing.
func TestGenerateAndValidate(t *testing.T) {
	for i := 0; i < 100; i++ {
		u := Generate()
		// Should be a valid UUID
		if !IsValid(u) {
			t.Errorf("Generated UUID is invalid: %s", u)
		}
		// Should be all lower or all upper case
		if strings.ToLower(u) != u && strings.ToUpper(u) != u {
			t.Errorf("UUID should be all lower or upper case: %s", u)
		}
	}
}

// TestParseAndFormat ensures round-trip of Parse and Format returns the same UUID (case-insensitive).
func TestParseAndFormat(t *testing.T) {
	u := Generate()
	id, err := Parse(u)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	u2 := Format(id)
	// Compare lowercased forms for equivalence
	if strings.ToLower(u) != strings.ToLower(u2) {
		t.Errorf("Format(Parse(u)) mismatch. got %s, want %s", u2, u)
	}
}

// TestMustParseValid ensures MustParse returns correct [16]byte for a valid UUID.
func TestMustParseValid(t *testing.T) {
	u := Generate()
	id := MustParse(u)
	u2 := Format(id)
	if strings.ToLower(u) != strings.ToLower(u2) {
		t.Errorf("MustParse/Format mismatch. got %s, want %s", u2, u)
	}
}

// TestIsValid covers valid and invalid UUID edge cases.
func TestIsValid(t *testing.T) {
	cases := []struct {
		UUID  string
		Valid bool
	}{
		{Generate(), true},
		// This is NOT a valid UUID (bad variant), so expect false:
		{"12345678-1234-1234-1234-123456789abc", false},
		{"12345678-1234-1234-1234-123456789abz", false},
		{"12345678123412341234123456789abc", false}, // missing dashes, so expect false
		{"", false},
	}
	for _, c := range cases {
		got := IsValid(c.UUID)
		if got != c.Valid {
			t.Errorf("IsValid(%q) == %v, want %v", c.UUID, got, c.Valid)
		}
	}
}

// TestParseErrors checks that Parse returns errors for invalid UUIDs.
func TestParseErrors(t *testing.T) {
	invalid := []string{
		"", "not-a-uuid",
		"zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz",
	}
	for _, s := range invalid {
		_, err := Parse(s)
		if err == nil {
			t.Errorf("Parse(%q) did not fail as expected", s)
		}
	}
}

// TestMustParsePanics checks that MustParse panics on invalid input.
func TestMustParsePanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustParse did not panic on invalid input")
		}
	}()
	MustParse("not-a-uuid")
}
