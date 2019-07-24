package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
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
	log *logrus.Logger

	mu      sync.Mutex
	lineNum bool
}

func init() {
	Log = newDefault()
}

func Setup() {
	Log = NewLogger()
}

func newDefault() *Logger {
	logger := &Logger{
		lineNum: false,
		log:     logrus.New(),
	}
	logger.log.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	}
	logger.log.Out = os.Stdout
	logger.log.Level = defLevel
	return logger
}

func NewLogger() *Logger {
	logger := &Logger{
		lineNum: conf.C.Pkg.Log.LineNum,
		log: logrus.New(),
	}

	formatter := strings.ToLower(conf.C.Pkg.Log.Formatter)
	switch formatter {
	case "json":
		logger.log.Formatter = &logrus.JSONFormatter{}
	case "text":
		logger.log.Formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		}
	default:
		logger.log.Formatter = &logrus.JSONFormatter{}
	}

	if fileName := conf.C.Pkg.Log.FileName; len(fileName) > 0 {
		out, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err == nil {
			logger.log.Out = out
		} else {
			logger.log.Out = os.Stdout
			logger.Errorf("Cannot create log file %s. %s", fileName, err)
		}
	} else {
		logger.log.Out = os.Stdout
	}

	if logLevel := conf.C.Pkg.Log.Level; len(logLevel) > 0 {
		level, err := logrus.ParseLevel(logLevel)
		if err == nil {
			logger.log.Level = level
		} else {
			logger.log.Level = defLevel
		}
	} else {
		logger.log.Level = defLevel
	}

	return logger
}

func (logger *Logger) NewEntryWithPrefix(prefix string) *logrus.Entry {
	return logger.log.WithField(PrefixField, prefix)
}

func (logger *Logger) Prefix(prefix string) *logrus.Entry {
	return logger.NewEntryWithPrefix(prefix)
}

func (logger *Logger) GetLineNum() bool {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.lineNum
}

func (logger *Logger) SetLineNum(lineNum bool) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.lineNum = lineNum
}

func (logger *Logger) GetOut() io.Writer {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.log.Out
}

func (logger *Logger) SetOut(out io.Writer) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.log.Out = out
}

func (logger *Logger) DiscardLog() {
	logger.SetOut(ioutil.Discard)
}

func (logger *Logger) getCallerLineNum() (string, string) {
	_, file, line, _ := runtime.Caller(2)
	lineInfo := fmt.Sprintf("%s:%v", file, line)
	return "line", lineInfo
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Debugf(format, args...)
	} else {
		logger.log.Debugf(format, args...)
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Infof(format, args...)
	} else {
		logger.log.Infof(format, args...)
	}
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Printf(format, args...)
	} else {
		logger.log.Printf(format, args...)
	}
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Warnf(format, args...)
	} else {
		logger.log.Warnf(format, args...)
	}
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Warningf(format, args...)
	} else {
		logger.log.Warningf(format, args...)
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Errorf(format, args...)
	} else {
		logger.log.Errorf(format, args...)
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Debug(args...)
	} else {
		logger.log.Debug(args...)
	}
}

func (logger *Logger) Print(args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Print(args...)
	} else {
		logger.log.Print(args...)
	}
}

func (logger *Logger) Warning(args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Warning(args...)
	} else {
		logger.log.Warning(args...)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Error(args...)
	} else {
		logger.log.Error(args...)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Info(args...)
	} else {
		logger.log.Info(args...)
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.lineNum {
		line, info := logger.getCallerLineNum()
		logger.log.WithField(line, info).Warn(args...)
	} else {
		logger.log.Warn(args...)
	}
}
