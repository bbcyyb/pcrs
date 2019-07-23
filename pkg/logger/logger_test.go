package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func removeTimestamp(logMessage string) string {
	noNewLine := strings.TrimSuffix(logMessage, "\n")
	noTimeStamp := strings.Split(noNewLine, " ")[1:]
	return strings.Join(noTimeStamp, " ")
}

func newLogger(b *bytes.Buffer, v *viper.Viper) *Logger {
	v.Set(OutputKey, b)
	l := New(v)
	return l
}

func TestNew(t *testing.T) {
	var expectedLogMessage string
	var actualLogMessage string

	var b bytes.Buffer
	v := viper.New()

	v.Set(LevelKey, "debug")
	l := newLogger(&b, v)

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

func TestSetPrefix(t *testing.T) {
	v := viper.New()
	v.Set(LevelKey, "Info")
	l := New(v)
	l.SetPrefix("test")
	actualPrefix := l.GetPrefix()
	expectedPrefix := "test"
	if actualPrefix != expectedPrefix {
		t.Errorf("Expected '%v', but got '%v", expectedPrefix, actualPrefix)
	}
}

func TestSetOut(t *testing.T) {
	v := viper.New()
	v.Set(LevelKey, "Info")
	l := New(v)
	l.SetOut(os.Stdout)
	actualPrefix := l.GetOut()
	expectedPrefix := os.Stdout
	if actualPrefix != expectedPrefix {
		t.Errorf("Expected '%v', but got '%v", expectedPrefix, actualPrefix)
	}
}
