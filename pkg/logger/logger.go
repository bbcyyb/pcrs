package logger

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type logConfig struct {
	Environment  string
	Formatter    string
	Level        string
	ExportCaller bool
	LogFilePath  string
}

var Logger *log.Logger

var config *logConfig

func init() {
	Logger = log.New()
	initConfig()
}

func initConfig() {
	config = &logConfig{
		Environment:  viper.GetString("environment"),
		Formatter:    viper.GetString("formatter"),
		Level:        viper.GetString("level"),
		ExportCaller: viper.GetBool("exportCaller"),
		LogFilePath:  viper.GetString("logFilePath"),
	}
	log.Debug(config)
	log.Debugf("Environment : %s", config.Environment)
	switch config.Environment {
	case "DEVELOPMENT":
		initLogToStdoutDebug()
	case "TEST":
		initInfoLogToFile()
	case "STAGING":
		initWarnLogToFile()
	case "PRODUCTION":
		initWarnLogToFile()
	default:
		initLogToStdoutDebug()
	}
}

func initLogToStdoutDebug() {
	Logger.SetLevel(log.DebugLevel)
	Logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano})
	Logger.SetOutput(os.Stdout)
	Logger.SetReportCaller(true)
}

func initInfoLogToFile() {
	Logger.SetLevel(log.InfoLevel)
	Logger.SetFormatter(&log.JSONFormatter{})
	f, err := os.OpenFile(config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()
	Logger.SetOutput(f)
}

func initWarnLogToFile() {
	Logger.SetLevel(log.InfoLevel)
	Logger.SetFormatter(&log.JSONFormatter{})
	f, err := os.OpenFile(config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	writer := bufio.NewWriter(f)
	Logger.SetOutput(writer)
}
