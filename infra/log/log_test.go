package log

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type logger struct {
	Level   string `json:"level"`
	Message string `json:"msg"`
	Time    string `json:"time"`
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestDebugJsonLog(t *testing.T) {
	ass := assert.New(t)
	buf := &bytes.Buffer{}
	SetFormatter(JSON)
	SetLevel(DebugLevel)
	SetOutput(buf)

	Errorln("Errorln")
	Warnln("Warnln")
	Infoln("Infoln")
	Debugln("Debugln")
	Traceln("Traceln")

	slice, err := refresh()
	if ass.NoError(err) {
		ass.Equal(4, len(slice))
		for index, val := range slice {
			var l logger
			json.Unmarshal([]byte(val), &l)
			switch index {
			case 0:
				ass.Equal("error", l.Level)
			case 1:
				ass.Equal("warning", l.Level)
			case 2:
				ass.Equal("info", l.Level)
			case 3:
				ass.Equal("debug", l.Level)
			}
		}
	}
}

func TestInfoTextLog(t *testing.T) {
	ass := assert.New(t)
	buf := &bytes.Buffer{}
	SetFormatter(TEXT)
	SetLevel(InfoLevel)
	SetOutput(buf)

	Error("Errorln")
	Warn("Warnln")
	Infof("Infoln %s", "with f")
	Debug("Debugln")
	Trace("Traceln")

	slice, err := refresh()
	if ass.NoError(err) {
		ass.Equal(3, len(slice))
		for index, val := range slice {
			//expected value looks like
			//time="2019-06-20T20:45:17.928749+08:00" level=error msg=Errorln
			var l logger
			json.Unmarshal([]byte(val), &l)
			switch index {
			case 0:
				ass.Contains(val, "level=error msg=Errorln")
			case 1:
				ass.Contains(val, "level=warning msg=Warnln")
			case 2:
				ass.Contains(val, "level=info msg=\"Infoln with f\"")
			}
		}
	}
}
