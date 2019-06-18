package log

import (
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

type Level uint32

const (
	PanicLevel Level = iota

	FatalLevel

	ErrorLevel

	WarnLevel

	InfoLevel

	DebugLevel

	TraceLevel
)

func InitLog() {
	if !terminal.IsTerminal(unix.Stdout) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		})
	}
}

func SetLevel(level Level) {
	logrus.SetLevel(level)
}

func Trace(args ...interface{}) {

}

func Debug(args ...interface{}) {

}

func Info(args ...interface{}) {

}

func Warn(args ...interface{}) {

}

func Error(args ...interface{}) {

}

func Fatal(args ...interface{}) {

}
