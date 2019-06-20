package log

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

type Level uint32
type Formatter uint32

const (
	_ Level = iota
	_
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

const (
	JSON Formatter = iota
	TEXT
)

var (
	buf io.ReadWriter
)

func SetFormatter(formatter Formatter) {
	if terminal.IsTerminal(unix.Stdout) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
	}

	switch formatter {
	case TEXT:
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		})
	case JSON:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

func SetOutput(readWriter io.ReadWriter) {
	buf = readWriter
	logrus.SetOutput(buf)
}

func refresh() (slice []string, err error) {
	br := bufio.NewReader(buf)
	slice = make([]string, 0)

	for {
		line, e := br.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			slice = append(slice, line)
		}

		if e != nil {
			if e != io.EOF {
				e = err
			}

			break
		}
	}

	return
}

// Default level is Info
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
