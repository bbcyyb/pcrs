package logger

import (
	"bytes"
	"github.com/bbcyyb/pcrs/conf"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"
	"testing"
)

type Suite struct {
	suite.Suite
	config             *conf.Config
	l                  *Logger
	expectedMsg        string
	expectedDebugMsg   string
	expectedInfoMsg    string
	expectedWarningMsg string
	expectedErrorMsg   string
}


func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	s.config = &conf.Config{
		Pkg: conf.Pkg{
			Log: conf.Log{
				Level:     "debug",
				Formatter: "text",
				LineNum:   false,
				FileName:  "",
			},
		},}
	conf.C = s.config
}

func removeTimestamp(logMessage string) string {
	noNewLine := strings.TrimSuffix(logMessage, "\n")
	noTimeStamp := strings.Split(noNewLine, " ")[1:]
	return strings.Join(noTimeStamp, " ")
}

func (s *Suite) TestNew() {
	s.l = NewLogger()
	ass := s.Assert()
	var actualLogMessage string
	var b bytes.Buffer
	s.l.SetOut(&b)

	s.l.Prefix("test").WithFields(logrus.Fields{"key": "value", "env": "test testing"}).Info("Information")
	s.expectedMsg = `level=info msg=Information env="test testing" key=value prefix=test`
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedMsg, actualLogMessage)
	b.Reset()

	s.l.Prefix("test").Warn("Warning")
	s.expectedMsg = "level=warning msg=Warning prefix=test"
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedMsg, actualLogMessage)
	b.Reset()

	s.l.Prefix("test").Error("Error")
	s.expectedMsg = "level=error msg=Error prefix=test"
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedMsg, actualLogMessage)
	b.Reset()
}


func (s *Suite) TestLog() {
	s.l = NewLogger()
	ass := s.Assert()
	var actualLogMessage string
	var b bytes.Buffer
	s.l.SetOut(&b)

	s.l.SetLineNum(false)
	s.expectedDebugMsg = `level=debug msg=debug`
	s.expectedInfoMsg = `level=info msg=info`
	s.expectedWarningMsg = `level=warning msg=warn`
	s.expectedErrorMsg = `level=error msg=error`

	s.l.Debug( "debug")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedDebugMsg, actualLogMessage)
	b.Reset()

	s.l.Info( "info")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedInfoMsg, actualLogMessage)
	b.Reset()

	s.l.Print( "info")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedInfoMsg, actualLogMessage)
	b.Reset()

	s.l.Warn( "warn")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedWarningMsg, actualLogMessage)
	b.Reset()

	s.l.Error( "error")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedErrorMsg, actualLogMessage)
	b.Reset()


	s.l.SetLineNum(false)
	s.expectedDebugMsg = `level=debug msg=debugf`
	s.expectedInfoMsg = `level=info msg=infof`
	s.expectedWarningMsg = `level=warning msg=warnf`
	s.expectedErrorMsg = `level=error msg=errorf`

	s.l.Debugf("%s", "debugf")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedDebugMsg, actualLogMessage)
	b.Reset()

	s.l.Infof("%s", "infof")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedInfoMsg, actualLogMessage)
	b.Reset()

	s.l.Printf("%s", "infof")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedInfoMsg, actualLogMessage)
	b.Reset()

	s.l.Warnf("%s", "warnf")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedWarningMsg, actualLogMessage)
	b.Reset()

	s.l.Errorf("%s", "errorf")
	actualLogMessage = removeTimestamp(b.String())
	ass.Equal(s.expectedErrorMsg, actualLogMessage)
	b.Reset()

	s.l.SetLineNum(true)
	s.expectedDebugMsg = `level=debug msg=debugf line=`
	s.expectedInfoMsg = `level=info msg=infof line=`
	s.expectedWarningMsg = `level=warning msg=warnf line=`
	s.expectedErrorMsg = `level=error msg=errorf line=`

	s.l.Debugf("%s", "debugf")
	actualLogMessage = removeTimestamp(b.String())
	contain := strings.Contains(actualLogMessage, s.expectedDebugMsg)
	s.True(contain)
	b.Reset()

	s.l.Infof("%s", "infof")
	actualLogMessage = removeTimestamp(b.String())
	contain = strings.Contains(actualLogMessage, s.expectedInfoMsg)
	s.True(contain)
	b.Reset()

	s.l.Printf("%s", "infof")
	actualLogMessage = removeTimestamp(b.String())
	contain = strings.Contains(actualLogMessage, s.expectedInfoMsg)
	s.True(contain)
	b.Reset()

	s.l.Warnf("%s", "warnf")
	actualLogMessage = removeTimestamp(b.String())
	contain = strings.Contains(actualLogMessage, s.expectedWarningMsg)
	s.True(contain)
	b.Reset()

	s.l.Errorf("%s", "errorf")
	actualLogMessage = removeTimestamp(b.String())
	contain = strings.Contains(actualLogMessage, s.expectedErrorMsg)
	s.True(contain)
	b.Reset()
}

func (s *Suite) TestSetLineNum() {
	ass := s.Assert()
	s.l.SetLineNum(false)
	actualValue := s.l.GetLineNum()
	expectedValue := false
	ass.Equal(expectedValue, actualValue)

	s.l.SetLineNum(true)
	actualValue = s.l.GetLineNum()
	expectedValue = true
	ass.EqualValues(expectedValue, actualValue)
}

func (s *Suite) TestSetOut() {
	ass := s.Assert()
	s.l.SetOut(os.Stdout)
	actualValue := s.l.GetOut()
	expectedValue := os.Stdout
	ass.Equal(expectedValue, actualValue)
}

func(s *Suite) TestGetCallerLineNum() {
	_, lineInfo := s.l.getCallerLineNum()
	s.NotEqual(":", lineInfo)
}
