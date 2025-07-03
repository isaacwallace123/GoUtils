package stringutil

import "testing"

// Each test function below checks a specific stringutil method with a clear, documented assertion.

func TestContainsAny(t *testing.T) {
	// Should find "foo" in "foobar"
	if !ContainsAny("foobar", []string{"foo", "baz"}) {
		t.Error("ContainsAny failed")
	}
}

func TestSlugify(t *testing.T) {
	// Should create "a-test" from "A test!"
	if Slugify("A test!") != "a-test" {
		t.Errorf("Slugify failed: got %q", Slugify("A test!"))
	}
}

func TestTruncate(t *testing.T) {
	// Should cut "abcdef" to "abc" for max 3
	if Truncate("abcdef", 3) != "abc" {
		t.Error("Truncate failed")
	}
}

func TestTitleCase(t *testing.T) {
	// Should become "Hello World"
	if TitleCase("hello WORLD") != "Hello World" {
		t.Error("TitleCase failed")
	}
}

func TestSnakeCase(t *testing.T) {
	// Should become "hello_world"
	if SnakeCase("Hello World") != "hello_world" {
		t.Error("SnakeCase failed")
	}
}

func TestKebabCase(t *testing.T) {
	// Should become "hello-world"
	if KebabCase("HelloWorld!") != "hello-world" {
		t.Error("KebabCase failed")
	}
}

func TestRemoveWhitespace(t *testing.T) {
	// Should become "abcd"
	if RemoveWhitespace(" a b\tc\nd ") != "abcd" {
		t.Error("RemoveWhitespace failed")
	}
}

func TestReverse(t *testing.T) {
	// Should become "cba"
	if Reverse("abc") != "cba" {
		t.Error("Reverse failed")
	}
}

func TestContainsAll(t *testing.T) {
	// Should be true: "foobar" contains "foo" and "bar"
	if !ContainsAll("foobar", []string{"foo", "bar"}) {
		t.Error("ContainsAll failed")
	}
}

func TestStartsWithAny(t *testing.T) {
	// "foobar" starts with "foo"
	if !StartsWithAny("foobar", []string{"foo", "baz"}) {
		t.Error("StartsWithAny failed")
	}
}

func TestEndsWithAny(t *testing.T) {
	// "foobar" ends with "bar"
	if !EndsWithAny("foobar", []string{"bar", "baz"}) {
		t.Error("EndsWithAny failed")
	}
}

func TestPadLeft(t *testing.T) {
	// "**x" is "x" padded to length 3 with '*'
	if PadLeft("x", '*', 3) != "**x" {
		t.Error("PadLeft failed")
	}
}

func TestPadRight(t *testing.T) {
	// "x--" is "x" padded right to length 3 with '-'
	if PadRight("x", '-', 3) != "x--" {
		t.Error("PadRight failed")
	}
}

func TestRepeat(t *testing.T) {
	// Should return "aaa"
	if Repeat("a", 3) != "aaa" {
		t.Error("Repeat failed")
	}
}

func TestToUpper(t *testing.T) {
	// "ab" to "AB"
	if ToUpper("ab") != "AB" {
		t.Error("ToUpper failed")
	}
}

func TestToLower(t *testing.T) {
	// "AB" to "ab"
	if ToLower("AB") != "ab" {
		t.Error("ToLower failed")
	}
}

func TestTrimSpace(t *testing.T) {
	// " a " trimmed to "a"
	if TrimSpace(" a ") != "a" {
		t.Error("TrimSpace failed")
	}
}

func TestReplace(t *testing.T) {
	// Replace all "a" with "c" in "aabb" -> "ccbb"
	if Replace("aabb", "a", "c", -1) != "ccbb" {
		t.Error("Replace failed")
	}
}

func TestHasPrefix(t *testing.T) {
	// "foo" has prefix "f"
	if !HasPrefix("foo", "f") {
		t.Error("HasPrefix failed")
	}
}

func TestHasSuffix(t *testing.T) {
	// "foo" has suffix "o"
	if !HasSuffix("foo", "o") {
		t.Error("HasSuffix failed")
	}
}

func TestContains(t *testing.T) {
	// "foo" contains "o"
	if !Contains("foo", "o") {
		t.Error("Contains failed")
	}
}

func TestToInt(t *testing.T) {
	// "123" should parse to 123
	if ToInt("123", 9) != 123 {
		t.Error("ToInt failed")
	}
}

func TestToInt32(t *testing.T) {
	// "123" should parse to 123 (int32)
	if ToInt32("123", 9) != 123 {
		t.Error("ToInt32 failed")
	}
}

func TestToInt64(t *testing.T) {
	// "123" should parse to 123 (int64)
	if ToInt64("123", 9) != 123 {
		t.Error("ToInt64 failed")
	}
}

func TestToUint(t *testing.T) {
	// "42" should parse to 42 (uint)
	if ToUint("42", 2) != 42 {
		t.Error("ToUint failed")
	}
}

func TestToUint32(t *testing.T) {
	// "42" should parse to 42 (uint32)
	if ToUint32("42", 2) != 42 {
		t.Error("ToUint32 failed")
	}
}

func TestToUint64(t *testing.T) {
	// "42" should parse to 42 (uint64)
	if ToUint64("42", 2) != 42 {
		t.Error("ToUint64 failed")
	}
}

func TestToFloat32(t *testing.T) {
	// "2.5" should parse to 2.5 (float32)
	if ToFloat32("2.5", 1) != 2.5 {
		t.Error("ToFloat32 failed")
	}
}

func TestToFloat64(t *testing.T) {
	// "2.5" should parse to 2.5 (float64)
	if ToFloat64("2.5", 1) != 2.5 {
		t.Error("ToFloat64 failed")
	}
}

func TestSplit(t *testing.T) {
	// Split by "," into two elements
	if len(Split("a,b", ",")) != 2 {
		t.Error("Split failed")
	}
}

func TestJoin(t *testing.T) {
	// Join ["a", "b"] with "-" to "a-b"
	if Join([]string{"a", "b"}, "-") != "a-b" {
		t.Error("Join failed")
	}
}
