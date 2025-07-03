package logger

import (
	"fmt"
	"github.com/isaacwallace123/GoUtils/color"
	"github.com/isaacwallace123/GoUtils/timeutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// LogTag holds a log level name and its associated color.
type LogTag struct {
	Name  string
	Color string
}

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

var currentLevel = LevelInfo

// SetLevel sets the current log level.
func SetLevel(level Level) {
	currentLevel = level
}

// Use color semantic colors for each log tag.
var (
	info  = LogTag{"INFO", color.InfoColor}
	warn  = LogTag{"WARN", color.WarnColor}
	err   = LogTag{"ERROR", color.ErrorColor}
	fatal = LogTag{"FATAL", color.ErrorColor}
	debug = LogTag{"DEBUG", color.DebugColor}
)

const reset = color.Reset

var logLock sync.Mutex

// log prints a formatted log message with the correct color, timestamp, file, etc.
func log(t LogTag, message string, args ...interface{}) {
	logLock.Lock()
	defer logLock.Unlock()

	_, file, _, ok := runtime.Caller(2)
	if !ok {
		file = "???"
	}
	shortFile := filepath.Base(file)
	timestamp := timeutil.FormatDateTime(timeutil.NowLocal())
	formatted := fmt.Sprintf(message, args...)

	fmt.Printf("[%s] %s[%s]%s [%s] %s\n", timestamp, t.Color, t.Name, reset, shortFile, formatted)
}

// Debug logs a message at DEBUG level.
func Debug(message string, args ...interface{}) {
	if currentLevel <= LevelDebug {
		log(debug, message, args...)
	}
}

// Info logs a message at INFO level.
func Info(message string, args ...interface{}) {
	if currentLevel <= LevelInfo {
		log(info, message, args...)
	}
}

// Warn logs a message at WARN level.
func Warn(message string, args ...interface{}) {
	if currentLevel <= LevelWarn {
		log(warn, message, args...)
	}
}

// Error logs a message at ERROR level and returns an error object.
func Error(message string, args ...interface{}) error {
	if currentLevel <= LevelError {
		log(err, message, args...)
	}
	return fmt.Errorf(message, args...)
}

// Fatal logs a message and exits the application.
func Fatal(message string, args ...interface{}) {
	log(fatal, message, args...)
	os.Exit(1)
}
