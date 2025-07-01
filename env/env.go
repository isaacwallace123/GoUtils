package env

import (
	"github.com/isaacwallace123/GoUtils/stringutil"
	"os"
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
	return stringutil.TrimSpace(Get(key, fallback))
}

// GetUpper returns the env variable in upper case.
func GetUpper(key, fallback string) string {
	return stringutil.ToUpper(Get(key, fallback))
}

// GetLower returns the env variable in lower case.
func GetLower(key, fallback string) string {
	return stringutil.ToLower(Get(key, fallback))
}

func GetInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i := stringutil.ToInt(val, fallback)
	return i
}

func GetInt32(key string, fallback int32) int32 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i := stringutil.ToInt64(val, int64(fallback))
	return int32(i)
}

func GetInt64(key string, fallback int64) int64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i := stringutil.ToInt64(val, fallback)
	return i
}

func GetUint(key string, fallback uint) uint {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i := stringutil.ToUint64(val, uint64(fallback))
	return uint(i)
}

func GetUint32(key string, fallback uint32) uint32 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i := stringutil.ToUint64(val, uint64(fallback))
	return uint32(i)
}

func GetUint64(key string, fallback uint64) uint64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	i := stringutil.ToUint64(val, fallback)
	return i
}

func GetFloat32(key string, fallback float32) float32 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	f := stringutil.ToFloat64(val, float64(fallback))
	return float32(f)
}

func GetFloat64(key string, fallback float64) float64 {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	f := stringutil.ToFloat64(val, fallback)
	return f
}

func GetBool(key string, fallback bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	switch stringutil.ToLower(val) {
	case "true", "1", "yes", "on":
		return true
	case "false", "0", "no", "off":
		return false
	default:
		return fallback
	}
}
