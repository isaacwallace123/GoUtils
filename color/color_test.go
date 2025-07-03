package color

import (
	"strconv"
	"testing"
)

// TestBasicANSI checks that basic color constants have the expected ANSI values.
func TestBasicANSI(t *testing.T) {
	// Foreground ANSI codes (check a sample)
	if Red != "\033[31m" {
		t.Errorf("Red should be \\033[31m, got %q", Red)
	}
	if BrightGreen != "\033[92m" {
		t.Errorf("BrightGreen should be \\033[92m, got %q", BrightGreen)
	}
	if Reset != "\033[0m" {
		t.Errorf("Reset should be \\033[0m, got %q", Reset)
	}
}

// TestSemanticColors checks that semantic HTTP/Log color constants are mapped to the correct base colors.
func TestSemanticColors(t *testing.T) {
	if GetColor != BrightGreen {
		t.Errorf("GetColor should match BrightGreen, got %q", GetColor)
	}
	if PostColor != BrightBlue {
		t.Errorf("PostColor should match BrightBlue, got %q", PostColor)
	}
	if PutColor != BrightYellow {
		t.Errorf("PutColor should match BrightYellow, got %q", PutColor)
	}
	if PatchColor != BrightMagenta {
		t.Errorf("PatchColor should match BrightMagenta, got %q", PatchColor)
	}
	if DeleteColor != BrightRed {
		t.Errorf("DeleteColor should match BrightRed, got %q", DeleteColor)
	}
}

// TestHTTPMethodToColorMap checks the map for HTTP methods to semantic color constants.
func TestHTTPMethodToColorMap(t *testing.T) {
	if HTTPMethodToColor["GET"] != GetColor {
		t.Error("HTTPMethodToColor[GET] did not match GetColor")
	}
	if HTTPMethodToColor["DELETE"] != DeleteColor {
		t.Error("HTTPMethodToColor[DELETE] did not match DeleteColor")
	}
}

// TestNameToCode ensures color names resolve to ANSI codes (including semantic names).
func TestNameToCode(t *testing.T) {
	// Basic color
	if NameToCode["red"] != Red {
		t.Error("NameToCode[red] did not match Red")
	}
	// Semantic alias
	if NameToCode["get"] != GetColor {
		t.Error("NameToCode[get] did not match GetColor")
	}
	if NameToCode["info"] != InfoColor {
		t.Error("NameToCode[info] did not match InfoColor")
	}
}

// TestPaint ensures Paint wraps the input string in the correct color code and reset.
func TestPaint(t *testing.T) {
	s := "Hello"
	c := BrightBlue
	expected := c + s + Reset
	if Paint(s, c) != expected {
		t.Errorf("Paint failed: got %q, want %q", Paint(s, c), expected)
	}
}

// TestANSI256 generates an ANSI 256-color code for an arbitrary value.
func TestANSI256(t *testing.T) {
	for _, n := range []int{0, 1, 15, 99, 200, 255} {
		code := ANSI256(n)
		expected := "\033[38;5;" + strconv.Itoa(n) + "m"
		if code != expected {
			t.Errorf("ANSI256(%d) failed: got %q, want %q", n, code, expected)
		}
	}
}

// TestRGB generates a truecolor ANSI code and checks the string formatting.
func TestRGB(t *testing.T) {
	r, g, b := 12, 34, 56
	expected := "\033[38;2;12;34;56m"
	got := RGB(r, g, b)
	if got != expected {
		t.Errorf("RGB(%d,%d,%d) failed: got %q, want %q", r, g, b, got, expected)
	}
}
