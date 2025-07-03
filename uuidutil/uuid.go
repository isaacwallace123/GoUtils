package uuidutil

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// uuidRegexp is a basic RFC 4122 UUID format validator.
var uuidRegexp = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

// Generate returns a random UUIDv4 string.
func Generate() string {
	var uuid [16]byte
	_, err := rand.Read(uuid[:])
	if err != nil {
		// This should never fail, so panic is OK for a util
		panic("unable to generate random UUID: " + err.Error())
	}

	// Set version (4) and variant bits according to RFC 4122.
	uuid[6] = (uuid[6] & 0x0F) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3F) | 0x80 // Variant 10

	return Format(uuid)
}

// Format takes a [16]byte UUID and returns its string form.
func Format(uuid [16]byte) string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// IsValid returns true if the string is a valid UUID (any version).
func IsValid(u string) bool {
	return uuidRegexp.MatchString(u)
}

// Parse parses a UUID string into [16]byte.
// Returns an error if the string is not a valid UUID.
func Parse(u string) ([16]byte, error) {
	var uuid [16]byte

	if !IsValid(u) {
		return uuid, errors.New("invalid UUID format")
	}

	u = strings.ReplaceAll(u, "-", "")
	if len(u) != 32 {
		return uuid, errors.New("invalid UUID length")
	}

	b, err := hex.DecodeString(u)
	if err != nil || len(b) != 16 {
		return uuid, errors.New("invalid UUID encoding")
	}

	copy(uuid[:], b)

	return uuid, nil
}

// MustParse parses a UUID string into [16]byte and panics on error.
func MustParse(u string) [16]byte {
	id, err := Parse(u)
	if err != nil {
		panic("invalid UUID: " + err.Error())
	}

	return id
}
