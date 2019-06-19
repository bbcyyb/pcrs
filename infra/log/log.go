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
	logrus.SetLevel(logrus.Level(level))
}

func GetLevel() Level {
	l := logrus.GetLevel()
	return Level(l)
}

func Trace(args ...interface{}) {
	logrus.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

func Traceln(args ...interface{}) {
	logrus.Traceln(args...)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}
