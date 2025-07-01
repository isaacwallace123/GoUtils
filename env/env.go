package env

import (
	"os"
	"strconv"
	"strings"
)

// Get returns the environment variable or a fallback.
func Get(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// GetRequired returns the environment variable or panics if not set.
func GetRequired(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("environment variable not set: " + key)
	}
	return val
}

// Exists checks if the environment variable is set.
func Exists(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}

// GetTrimmed returns a trimmed version of the env variable.
func GetTrimmed(key, fallback string) string {
	return strings.TrimSpace(Get(key, fallback))
}

// GetUpper returns the env variable in upper case.
func GetUpper(key, fallback string) string {
	return strings.ToUpper(Get(key, fallback))
}

// GetLower returns the env variable in lower case.
func GetLower(key, fallback string) string {
	return strings.ToLower(Get(key, fallback))
}

func GetInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}

func GetInt32(key string, fallback int32) int32 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return fallback
	}
	return int32(i)
}

func GetInt64(key string, fallback int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return fallback
	}
	return i
}

func GetUint(key string, fallback uint) uint {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return fallback
	}
	return uint(i)
}

func GetUint32(key string, fallback uint32) uint32 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return fallback
	}
	return uint32(i)
}

func GetUint64(key string, fallback uint64) uint64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return fallback
	}
	return i
}

func GetFloat32(key string, fallback float32) float32 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	f, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return fallback
	}
	return float32(f)
}

func GetFloat64(key string, fallback float64) float64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fallback
	}
	return f
}

func GetBool(key string, fallback bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	switch strings.ToLower(val) {
	case "true", "1", "yes", "on":
		return true
	case "false", "0", "no", "off":
		return false
	default:
		return fallback
	}
}
