package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/bbcyyb/pcrs/conf"
	"github.com/sirupsen/logrus"
)

var config = &conf.Config{
	Pkg: conf.Pkg{
		Log: conf.Log{
			Level:     "debug",
			Formatter: "text",
			LineNum:   true,
			FileName:  "",
		},
	},
}

func removeTimestamp(logMessage string) string {
	noNewLine := strings.TrimSuffix(logMessage, "\n")
	noTimeStamp := strings.Split(noNewLine, " ")[1:]
	return strings.Join(noTimeStamp, " ")
}

func TestNew(t *testing.T) {
	conf.C = config
	var expectedLogMessage string
	var actualLogMessage string

	var b bytes.Buffer
	l := NewLogger()
	l.SetOut(&b)

	l.Prefix("test").WithFields(logrus.Fields{"key": "value", "env": "test testing"}).Info("Information")
	expectedLogMessage = `level=info msg=Information env="test testing" key=value prefix=test`
	fmt.Println(b.String())
	actualLogMessage = removeTimestamp(b.String())
	if actualLogMessage != expectedLogMessage {
		t.Errorf("Expected '%s', but got '%s'", expectedLogMessage, actualLogMessage)
	}
	b.Reset()

	l.Prefix("test").Warn("Warning")
	expectedLogMessage = "level=warning msg=Warning prefix=test"
	actualLogMessage = removeTimestamp(b.String())
	if actualLogMessage != expectedLogMessage {
		t.Errorf("Expected '%s', but got '%s'", expectedLogMessage, actualLogMessage)
	}
	b.Reset()

	l.Prefix("test").Error("Error")
	expectedLogMessage = "level=error msg=Error prefix=test"
	actualLogMessage = removeTimestamp(b.String())
	if actualLogMessage != expectedLogMessage {
		t.Errorf("Expected '%s', but got '%s'", expectedLogMessage, actualLogMessage)
	}
	b.Reset()
}

func TestSetLineNum(t *testing.T) {
	l := NewLogger()
	l.SetLineNum(false)
	actualValue := l.GetLineNum()
	expectedValue := false
	if actualValue != expectedValue {
		t.Errorf("Expected '%v', but got '%v", expectedValue, actualValue)
	}

	l.SetLineNum(true)
	actualValue = l.GetLineNum()
	expectedValue = true
	if actualValue != expectedValue {
		t.Errorf("Expected '%v', but got '%v", expectedValue, actualValue)
	}
}

func TestSetOut(t *testing.T) {
	l := NewLogger()
	l.SetOut(os.Stdout)
	actualValue := l.GetOut()
	expectedValue := os.Stdout
	if actualValue != expectedValue {
		t.Errorf("Expected '%v', but got '%v", expectedValue, actualValue)
	}
}
