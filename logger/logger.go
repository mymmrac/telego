package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

func CreateLogrusLogger(level logrus.Level) Logger {
	log := logrus.StandardLogger()

	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = time.RFC1123
	formatter.FullTimestamp = true
	log.SetFormatter(formatter)

	log.SetLevel(level)

	return log
}
