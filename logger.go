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

type logMode string

const (
	debugMode logMode = "DEBUG"
	errorMode logMode = "ERROR"

	ansiReset  = "\u001B[0m"
	ansiRed    = "\u001B[31m"
	ansiYellow = "\u001B[33m"
	ansiBlue   = "\u001B[34m"
)

type logger struct {
	Out         io.Writer
	DebugMode   bool
	PrintErrors bool
	Replacer    *strings.Replacer

	mutex sync.Mutex
}

func newDefaultLogger(token string) *logger {
	return &logger{
		Out:         os.Stderr,
		DebugMode:   false,
		PrintErrors: true,
		Replacer:    defaultReplacer(token),
	}
}

func (l *logger) prefix(mode logMode) string {
	timeNow := ansiBlue + time.Now().Format(time.UnixDate) + ansiReset
	switch mode {
	case debugMode:
		return fmt.Sprintf("[%s] %sDEBUG%s ", timeNow, ansiYellow, ansiReset)
	case errorMode:
		return fmt.Sprintf("[%s] %sERROR%s ", timeNow, ansiRed, ansiReset)
	}
	return "LOGGING "
}

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

func (l *logger) Debugf(format string, args ...any) {
	if l.DebugMode {
		l.log(debugMode, fmt.Sprintf(format+"\n", args...))
	}
}

func (l *logger) Errorf(format string, args ...any) {
	if l.PrintErrors {
		l.log(errorMode, fmt.Sprintf(format+"\n", args...))
	}
}

// DefaultLoggerTokenReplacement used to replace bot token in logs when using default logger
const DefaultLoggerTokenReplacement = "BOT_TOKEN"

func defaultReplacer(token string) *strings.Replacer {
	return strings.NewReplacer(token, DefaultLoggerTokenReplacement)
}
