package telego

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Logger represents logger used to debug or error information
type Logger interface {
	Debugf(format string, args ...any)
	Errorf(format string, args ...any)
}

// logMode represents logging mode
type logMode string

// Logging modes
const (
	debugMode logMode = "DEBUG"
	errorMode logMode = "ERROR"
)

// ANSI escape codes
const (
	ansiReset  = "\u001B[0m"
	ansiRed    = "\u001B[31m"
	ansiYellow = "\u001B[33m"
	ansiBlue   = "\u001B[34m"
)

// logger used to debug or error information
type logger struct {
	Out         io.Writer
	DebugMode   bool
	PrintErrors bool
	Replacer    *strings.Replacer

	mutex sync.Mutex
}

// newDefaultLogger creates new default logger
func newDefaultLogger(token string) *logger {
	return &logger{
		Out:         os.Stderr,
		DebugMode:   false,
		PrintErrors: true,
		Replacer:    defaultReplacer(token),
	}
}

// prefix returns log prefix
func (l *logger) prefix(mode logMode) string {
	timeNow := ansiBlue + time.Now().Format(time.UnixDate) + ansiReset
	switch mode {
	case debugMode:
		return fmt.Sprintf("[%s] %sDEBUG%s ", timeNow, ansiYellow, ansiReset)
	case errorMode:
		return fmt.Sprintf("[%s] %sERROR%s ", timeNow, ansiRed, ansiReset)
	default:
		return "LOGGING "
	}
}

// log logs text
func (l *logger) log(mode logMode, text string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.Replacer != nil {
		text = l.Replacer.Replace(text)
	}

	_, err := l.Out.Write([]byte(l.prefix(mode) + text))
	if err != nil {
		//nolint:forbidigo
		_, _ = fmt.Printf("Logging error: %v\n", err)
	}
}

// Debugf logs debug information
func (l *logger) Debugf(format string, args ...any) {
	if l.DebugMode {
		l.log(debugMode, fmt.Sprintf(format+"\n", args...))
	}
}

// Errorf logs error information
func (l *logger) Errorf(format string, args ...any) {
	if l.PrintErrors {
		l.log(errorMode, fmt.Sprintf(format+"\n", args...))
	}
}

// DefaultLoggerTokenReplacement used to replace bot token in logs when using default logger
const DefaultLoggerTokenReplacement = "BOT_TOKEN"

// defaultReplacer returns replacer for default logger
func defaultReplacer(token string) *strings.Replacer {
	return strings.NewReplacer(token, DefaultLoggerTokenReplacement)
}
