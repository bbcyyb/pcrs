package logger

import (
	"fmt"
	"github.com/bbcyyb/pcrs/conf"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Defaults values to be used when creating a Logger without user parameters
const (
	defLevel     = logrus.InfoLevel
	PrefixField = "prefix"
)

var Log *Logger

type Logger struct {
	log *logrus.Logger

	mu           sync.Mutex
	enableCaller bool
}

func init() {
	Log = newDefault()
}

func SetupLogger() {
	Log = NewLogger()
}

func newDefault() *Logger {
	logger := &Logger{
		enableCaller: false,
		log:          logrus.New(),
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
		enableCaller: conf.LogConf.Caller,
	}

	formatter := strings.ToLower(conf.LogConf.Formatter)
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

	if fileName := conf.LogConf.FileName; len(fileName) > 0 {
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

	logLevel, err := logrus.ParseLevel(conf.LogConf.Level)
	if err == nil {
		logger.log.Level = logLevel
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

func (logger *Logger) GetEnableCaller() bool {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	return logger.enableCaller
}

func (logger *Logger) SetEnableCaller(enableCaller bool) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.enableCaller = enableCaller
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

func (logger *Logger) getAdditionalField() (string, string) {
	_, file, line, _ := runtime.Caller(2)
	lineInfo := fmt.Sprintf("%s:%v", file, line)
	return "line", lineInfo
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Debugf(format, args...)
	} else {
		logger.log.Debugf(format, args...)
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Infof(format, args...)
	} else {
		logger.log.Infof(format, args...)
	}
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Printf(format, args...)
	} else {
		logger.log.Printf(format, args...)
	}
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Warnf(format, args...)
	} else {
		logger.log.Warnf(format, args...)
	}
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Warningf(format, args...)
	} else {
		logger.log.Warningf(format, args...)
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Errorf(format, args...)
	} else {
		logger.log.Errorf(format, args...)
	}
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Fatalf(format, args...)
	} else {
		logger.log.Fatalf(format, args...)
	}
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Panicf(format, args...)
	} else {
		logger.log.Panicf(format, args...)
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Debug(args...)
	} else {
		logger.log.Debug(args...)
	}
}

func (logger *Logger) Print(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Print(args...)
	} else {
		logger.log.Print(args...)
	}
}

func (logger *Logger) Warning(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Warning(args...)
	} else {
		logger.log.Warning(args...)
	}
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Fatal(args...)
	} else {
		logger.log.Fatal(args...)
	}
}

func (logger *Logger) Panic(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Panic(args...)
	} else {
		logger.log.Panic(args...)
	}
}

func (logger *Logger) Debugln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Debugln(args...)
	} else {
		logger.log.Debugln(args...)
	}
}

func (logger *Logger) Infoln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Infoln(args...)
	} else {
		logger.log.Infoln(args...)
	}
}

func (logger *Logger) Println(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Println(args...)
	} else {
		logger.log.Println(args...)
	}
}

func (logger *Logger) Warnln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Warnln(args...)
	} else {
		logger.log.Warnln(args...)
	}
}

func (logger *Logger) Warningln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Warningln(args...)
	} else {
		logger.log.Warningln(args...)
	}
}

func (logger *Logger) Errorln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Errorln(args...)
	} else {
		logger.log.Errorln(args...)
	}
}

func (logger *Logger) Fatalln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Fatalln(args...)
	} else {
		logger.log.Fatalln(args...)
	}
}

func (logger *Logger) Panicln(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Panicln(args...)
	} else {
		logger.log.Panicln(args...)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Error(args...)
	} else {
		logger.log.Error(args...)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Info(args...)
	} else {
		logger.log.Info(args...)
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.enableCaller {
		key, value := logger.getAdditionalField()
		logger.log.WithField(key, value).Warn(args...)
	} else {
		logger.log.Warn(args...)
	}
}
