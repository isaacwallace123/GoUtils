package jsonutil

import (
	"testing"
)

// TestToStringAndFromString verifies round-trip JSON encoding/decoding.
func TestToStringAndFromString(t *testing.T) {
	type Demo struct{ Name string }
	d := Demo{"Alice"}
	// Test encoding to JSON string
	jsonStr := ToString(d)
	if jsonStr == "" || jsonStr[0] != '{' {
		t.Errorf("ToString failed: got %s", jsonStr)
	}
	// Test decoding back from JSON string
	var out Demo
	if err := FromString(jsonStr, &out); err != nil || out.Name != "Alice" {
		t.Errorf("FromString failed: %v, out: %+v", err, out)
	}
}

// TestCompactAndPretty checks JSON pretty printing and compaction.
func TestCompactAndPretty(t *testing.T) {
	jsonMin := `{"a":1,"b":2}`
	jsonPretty := Pretty(jsonMin)
	// Ensure Pretty output is valid JSON
	if !IsValid(jsonPretty) {
		t.Errorf("Pretty returned invalid JSON: %q", jsonPretty)
	}
	// Compaction should return original minified string
	jsonComp := Compact(jsonPretty)
	if jsonComp != jsonMin {
		t.Errorf("Compact did not match: got %q, want %q", jsonComp, jsonMin)
	}
}

// TestEqual checks deep JSON equality for semantically identical objects.
func TestEqual(t *testing.T) {
	a := `{"x":1,"y":2}`
	b := `{
		"y":2,
		"x":1
	}`
	if !Equal(a, b) {
		t.Error("Equal should consider objects with same keys equal")
	}
}

// TestAppendRemove checks array append/remove manipulation in JSON strings.
func TestAppendRemove(t *testing.T) {
	start := `[1,2,3]`
	added := Append(start, 4)
	expected := `[1,2,3,4]`
	if !Equal(added, expected) {
		t.Errorf("Append failed: got %q", added)
	}
	removed := Remove(added, 2) // remove "3"
	expected2 := `[1,2,4]`
	if !Equal(removed, expected2) {
		t.Errorf("Remove failed: got %q", removed)
	}
}

// TestStripComments ensures comments are stripped from JSON-like input.
func TestStripComments(t *testing.T) {
	input := `
	// this is a comment
	{ "a": 1 }
	/* block
	comment */
	{ "b": 2 } // comment
	`
	result := StripComments(input)
	if result == "" || result[0] != '{' {
		t.Errorf("StripComments failed: got %q", result)
	}
}

// TestToObject ensures JSON is converted to a map as expected.
func TestToObject(t *testing.T) {
	jsonStr := `{"k": "v"}`
	obj := ToObject(jsonStr)
	if obj == nil || obj["k"] != "v" {
		t.Errorf("ToObject failed: got %+v", obj)
	}
}
