package telego

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Logger represents logger used to debug or error information
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
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
	mutex       sync.Mutex
}

func newLogger() *logger {
	return &logger{
		Out:         os.Stderr,
		DebugMode:   false,
		PrintErrors: true,
	}
}

func (l *logger) prefix(mode logMode) string {
	timeNow := ansiBlue + time.Now().Local().Format(time.UnixDate) + ansiReset
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
	_, err := l.Out.Write([]byte(l.prefix(mode) + text))
	if err != nil {
		fmt.Printf("Logging error: %v\n", err)
	}
}

func (l *logger) Debug(args ...interface{}) {
	if l.DebugMode {
		l.log(debugMode, fmt.Sprintln(args...))
	}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	if l.DebugMode {
		l.log(debugMode, fmt.Sprintf(format+"\n", args...))
	}
}

func (l *logger) Error(args ...interface{}) {
	if l.PrintErrors {
		l.log(errorMode, fmt.Sprintln(args...))
	}
}

func (l *logger) Errorf(format string, args ...interface{}) {
	if l.PrintErrors {
		l.log(errorMode, fmt.Sprintf(format+"\n", args...))
	}
}
