package logger

import (
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/bbcyyb/pcrs/conf"
	"github.com/sirupsen/logrus"
)

// Defaults values to be used when creating a Logger without user parameters
const (
	defLevel    = logrus.InfoLevel
	PrefixField = "prefix"
)

var Log *Logger

type Logger struct {
	logrus.Logger

	mu     sync.Mutex
	prefix string
}

func init() {
	Log = newDefault()
}

func Setup() {
	Log = NewLogger()
}

func newDefault() *Logger {
	logger := &Logger{
		prefix: "",
	}
	logger.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	}
	logger.Out = os.Stdout
	logger.Level = defLevel
	return logger
}

func NewLogger() *Logger {
	prefix := conf.C.Pkg.Log.Prefix
	logger := &Logger{
		prefix: prefix,
	}

	formatter := conf.C.Pkg.Log.Formatter
	switch formatter {
	case "json":
		logger.Formatter = &logrus.JSONFormatter{}
	case "text":
		logger.Formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		}
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	if fileName := conf.C.Pkg.Log.FileName; len(fileName) > 0 {
		out, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err == nil {
			logger.Out = out
		} else {
			logger.Errorf("Cannot create log file %s. %s", fileName, err)
		}
	} else {
		logger.Out = os.Stdout
	}

	if logLevel := conf.C.Pkg.Log.Level; len(logLevel) > 0 {
		level, err := logrus.ParseLevel(logLevel)
		if err == nil {
			logger.Level = level
		} else {
			logger.Level = defLevel
		}
	} else {
		logger.Level = defLevel
	}

	return logger
}

func (logger *Logger) NewEntryWithPrefix(prefix string) *logrus.Entry {
	return logger.WithField(PrefixField, prefix)
}

func (logger *Logger) Prefix(prefix string) *logrus.Entry {
	return logger.NewEntryWithPrefix(prefix)
}

func (logger *Logger) GetPrefix() string {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.prefix
}

func (logger *Logger) SetPrefix(prefix string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.prefix = prefix
}

func (logger *Logger) GetOut() io.Writer {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.Out
}

func (logger *Logger) SetOut(out io.Writer) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.Out = out
}

func (logger *Logger) DiscardLog() {
	logger.SetOut(ioutil.Discard)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Warnf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Panicf(format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Debug(args...)
}

func (logger *Logger) Print(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Info(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Warn(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Panic(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Warnln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Fatalln(args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Panicln(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Error(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.WithField(PrefixField, logger.prefix).Warn(args...)
}
