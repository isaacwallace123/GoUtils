package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/isaacwallace123/GoUtils/stringutil"
)

// ToString encodes a value into a compact JSON string.
func ToString(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("json error: %v", err)
	}
	return string(data)
}

// ToStringIndent encodes a value into a pretty JSON string.
func ToStringIndent(v any) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("json error: %v", err)
	}
	return string(data)
}

// FromString decodes a JSON string into a target value.
func FromString(jsonStr string, target any) error {
	return json.Unmarshal([]byte(jsonStr), target)
}

// FromBytes decodes a JSON byte slice into a target value.
func FromBytes(data []byte, target any) error {
	return json.Unmarshal(data, target)
}

// Compact removes whitespace from a JSON string.
func Compact(jsonStr string) string {
	var buf bytes.Buffer
	if err := json.Compact(&buf, []byte(jsonStr)); err != nil {
		return jsonStr
	}
	return buf.String()
}

// Pretty formats a JSON string for readability.
func Pretty(jsonStr string) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(jsonStr), "", "  "); err != nil {
		return jsonStr
	}
	return buf.String()
}

// IsValid checks if a JSON string is valid.
func IsValid(jsonStr string) bool {
	var js any
	return json.Unmarshal([]byte(jsonStr), &js) == nil
}

// Equal checks if two JSON strings represent the same data.
func Equal(a, b string) bool {
	var o1 any
	var o2 any
	if err := json.Unmarshal([]byte(a), &o1); err != nil {
		return false
	}
	if err := json.Unmarshal([]byte(b), &o2); err != nil {
		return false
	}
	return fmt.Sprintf("%v", o1) == fmt.Sprintf("%v", o2)
}

// MustFromString panics if FromString fails.
func MustFromString(jsonStr string, target any) {
	if err := FromString(jsonStr, target); err != nil {
		panic(err)
	}
}

// MustToString panics if encoding fails.
func MustToString(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// StripComments removes // and /* */ style comments from a JSON-like string.
func StripComments(jsonStr string) string {
	lines := stringutil.Split(jsonStr, "\n")
	var cleaned []string
	for _, line := range lines {
		line = stringutil.TrimSpace(line)
		if stringutil.HasPrefix(line, "//") || line == "" {
			continue
		}
		cleaned = append(cleaned, line)
	}
	return stringutil.Join(cleaned, "\n")
}

// Append appends a new item to a JSON array string.
// If input is not a valid array, it returns the original string.
func Append(jsonArrayStr string, item any) string {
	var arr []any
	if err := json.Unmarshal([]byte(jsonArrayStr), &arr); err != nil {
		return jsonArrayStr
	}
	arr = append(arr, item)
	data, err := json.Marshal(arr)
	if err != nil {
		return jsonArrayStr
	}
	return string(data)
}

// Remove removes the item at the specified index from a JSON array string.
// If input is invalid or index is out of range, the original string is returned.
func Remove(jsonArrayStr string, index int) string {
	var arr []any
	if err := json.Unmarshal([]byte(jsonArrayStr), &arr); err != nil {
		return jsonArrayStr
	}
	if index < 0 || index >= len(arr) {
		return jsonArrayStr
	}
	arr = append(arr[:index], arr[index+1:]...)
	data, err := json.Marshal(arr)
	if err != nil {
		return jsonArrayStr
	}
	return string(data)
}

// ToObject parses a JSON string into a generic object (map[string]any).
// Returns nil if parsing fails.
func ToObject(jsonStr string) map[string]any {
	var obj map[string]any
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return nil
	}
	return obj
}
