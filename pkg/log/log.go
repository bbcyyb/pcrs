package log

import (
	"bufio"
	"io"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Level uint32
type Formatter uint32
type Fields map[string]interface{}

type Entry struct {
	CoreEntry logrus.Entry
}

type EntryHandler interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

func (e *Entry) Info(args ...interface{}) {
	e.CoreEntry.Info(args...)
}

func (e *Entry) Error(args ...interface{}) {
	e.CoreEntry.Error(args...)
}

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

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func WithFields(fields Fields) EntryHandler {
	entry := *logrus.WithFields(logrus.Fields(fields))
	return &Entry{
		CoreEntry: entry,
	}
}
