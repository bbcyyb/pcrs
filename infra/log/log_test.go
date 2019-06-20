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
		//assert.Equal(t, expectedObj, actualObj)
	}
}

func TestInfoTextLog(t *testing.T) {
	SetFormatter(TEXT)
	SetLevel(InfoLevel)
}
