package stringutil

import (
	"regexp"
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
