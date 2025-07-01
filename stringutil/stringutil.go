package stringutil

import (
	"regexp"
	"strconv"
	"strings"
)

// ContainsAny checks if the input string contains any of the substrings.
func ContainsAny(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// Slugify converts a string to lowercase dash-separated format.
func Slugify(s string) string {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

// Truncate shortens a string and appends "..." if it exceeds the limit.
func Truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}

// TitleCase capitalizes the first letter of each word.
func TitleCase(s string) string {
	return strings.Title(strings.ToLower(s))
}

// SnakeCase converts a string to snake_case.
func SnakeCase(s string) string {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "_")
	s = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(strings.Trim(s, "_"))
}

// KebabCase converts a string to kebab-case.
func KebabCase(s string) string {
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "-")
	s = regexp.MustCompile("([a-z0-9])([A-Z])").ReplaceAllString(s, "${1}-${2}")
	return strings.ToLower(strings.Trim(s, "-"))
}

// RemoveWhitespace removes all whitespace characters from the string.
func RemoveWhitespace(s string) string {
	return strings.Join(strings.Fields(s), "")
}

// Reverse returns the input string reversed.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ContainsAll checks if all substrings are present in the input string.
func ContainsAll(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// StartsWithAny checks if the string starts with any of the given prefixes.
func StartsWithAny(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

// EndsWithAny checks if the string ends with any of the given suffixes.
func EndsWithAny(s string, suffixes []string) bool {
	for _, sfx := range suffixes {
		if strings.HasSuffix(s, sfx) {
			return true
		}
	}
	return false
}

// PadLeft adds padding characters to the left of the string until it reaches the target length.
func PadLeft(s string, padChar rune, length int) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(padChar), length-len(s)) + s
}

// PadRight adds padding characters to the right of the string until it reaches the target length.
func PadRight(s string, padChar rune, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(string(padChar), length-len(s))
}

// Repeat returns the input string repeated n times.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// ToUpper returns the string in uppercase.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower returns the string in lowercase.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// TrimSpace trims all surrounding whitespace.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Replace replaces all occurrences of old with new.
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// HasPrefix checks if string starts with prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix checks if string ends with suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Contains checks if string contains substring.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ToInt converts string to int with fallback.
func ToInt(s string, fallback int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return i
}

// ToInt32 converts string to int32 with fallback.
func ToInt32(s string, fallback int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fallback
	}
	return int32(i)
}

// ToInt64 converts string to int64 with fallback.
func ToInt64(s string, fallback int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fallback
	}
	return i
}

// ToUint converts string to uint with fallback.
func ToUint(s string, fallback uint) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return fallback
	}
	return uint(i)
}

// ToUint32 converts string to uint32 with fallback.
func ToUint32(s string, fallback uint32) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return fallback
	}
	return uint32(i)
}

// ToUint64 converts string to uint64 with fallback.
func ToUint64(s string, fallback uint64) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return fallback
	}
	return i
}

// ToFloat32 converts string to float32 with fallback.
func ToFloat32(s string, fallback float32) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return fallback
	}
	return float32(f)
}

// ToFloat64 converts string to float64 with fallback.
func ToFloat64(s string, fallback float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fallback
	}
	return f
}

// Split splits a string into substrings separated by the given separator.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join concatenates the elements of a slice to create a single string with the given separator.
func Join(parts []string, sep string) string {
	return strings.Join(parts, sep)
}
