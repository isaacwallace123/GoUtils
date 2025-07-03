package logger

import (
	"testing"
)

// TestLoggerLevels checks that log messages print at each level and Error returns an error.
func TestLoggerLevels(t *testing.T) {
	SetLevel(LevelDebug)
	Debug("Debug: %s", "visible")        // Should print
	Info("Info: %s", "visible")          // Should print
	Warn("Warn: %s", "visible")          // Should print
	err := Error("Error: %s", "visible") // Should print and return error
	if err == nil {
		t.Error("Error() should return error")
	}
	// Don't call Fatal: it exits the program
	SetLevel(LevelError)
	Info("This should not be visible") // Should NOT print due to log level
}

// TestLoggerThreadSafe runs concurrent logging to check for data races and panics.
func TestLoggerThreadSafe(t *testing.T) {
	SetLevel(LevelInfo)
	done := make(chan bool)
	go func() {
		Info("Thread 1 logging")
		done <- true
	}()
	go func() {
		Warn("Thread 2 logging")
		done <- true
	}()
	<-done
	<-done
}
