package logger

import (
	"fmt"
	"github.com/isaacwallace123/GoUtils/timeutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

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

func SetLevel(level Level) {
	currentLevel = level
}

var (
	info  = LogTag{"INFO", "\033[34m"}
	warn  = LogTag{"WARN", "\033[33m"}
	err   = LogTag{"ERROR", "\033[31m"}
	fatal = LogTag{"FATAL", "\033[31m"}
	debug = LogTag{"DEBUG", "\033[90m"}
)

const reset = "\033[0m"

var logLock sync.Mutex

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

func Debug(message string, args ...interface{}) {
	if currentLevel <= LevelDebug {
		log(debug, message, args...)
	}
}

func Info(message string, args ...interface{}) {
	if currentLevel <= LevelInfo {
		log(info, message, args...)
	}
}

func Warn(message string, args ...interface{}) {
	if currentLevel <= LevelWarn {
		log(warn, message, args...)
	}
}

func Error(message string, args ...interface{}) error {
	if currentLevel <= LevelError {
		log(err, message, args...)
	}
	return fmt.Errorf(message, args...)
}

func Fatal(message string, args ...interface{}) {
	log(fatal, message, args...)
	os.Exit(1)
}
