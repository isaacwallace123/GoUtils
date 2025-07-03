package env

import (
	"os"
	"testing"
)

// TestGetAndFallback checks that Get returns the environment variable's value when set,
// and returns the fallback value when not set.
func TestGetAndFallback(t *testing.T) {
	os.Setenv("FOO_TEST", "bar")
	defer os.Unsetenv("FOO_TEST") // Ensure environment is clean after test

	if got := Get("FOO_TEST", "fallback"); got != "bar" {
		t.Errorf("Get did not return set value: got %q", got)
	}

	if got := Get("NOT_SET", "fallback"); got != "fallback" {
		t.Errorf("Get did not return fallback: got %q", got)
	}
}

// TestGetRequiredPanics verifies that GetRequired panics if the variable is not set.
func TestGetRequiredPanics(t *testing.T) {
	defer func() {
		// Recover from panic and check that it actually happened
		if r := recover(); r == nil {
			t.Errorf("GetRequired did not panic for unset variable")
		}
	}()
	GetRequired("NEVER_SET_123")
}

// TestExists checks that Exists returns true for set variables and false for unset.
func TestExists(t *testing.T) {
	os.Setenv("EXIST_TEST", "abc")
	defer os.Unsetenv("EXIST_TEST")

	if !Exists("EXIST_TEST") {
		t.Error("Exists should return true for set var")
	}
	if Exists("NOPE_THIS_IS_NOT_SET") {
		t.Error("Exists should return false for unset var")
	}
}

// TestGetTrimmedUpperLower checks trimming, uppercasing, and lowercasing.
func TestGetTrimmedUpperLower(t *testing.T) {
	os.Setenv("MIX_TEST", "  VaLuE  ")
	defer os.Unsetenv("MIX_TEST")

	if got := GetTrimmed("MIX_TEST", "fallback"); got != "VaLuE" {
		t.Errorf("GetTrimmed got %q", got)
	}
	if got := GetUpper("MIX_TEST", "fallback"); got != "  VALUE  " {
		t.Errorf("GetUpper got %q", got)
	}
	if got := GetLower("MIX_TEST", "fallback"); got != "  value  " {
		t.Errorf("GetLower got %q", got)
	}
}

// TestGetNumericAndBool checks all numeric and boolean conversions for env values.
func TestGetNumericAndBool(t *testing.T) {
	os.Setenv("INT_TEST", "42")
	os.Setenv("FLOAT_TEST", "3.14")
	os.Setenv("BOOL_TEST", "true")
	defer func() {
		os.Unsetenv("INT_TEST")
		os.Unsetenv("FLOAT_TEST")
		os.Unsetenv("BOOL_TEST")
	}()

	// Test int
	if got := GetInt("INT_TEST", 0); got != 42 {
		t.Errorf("GetInt got %v", got)
	}
	// Test float64
	if got := GetFloat64("FLOAT_TEST", 0); got != 3.14 {
		t.Errorf("GetFloat64 got %v", got)
	}
	// Test bool
	if got := GetBool("BOOL_TEST", false); got != true {
		t.Errorf("GetBool got %v", got)
	}
}
