package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/bbcyyb/pcrs/infra/log"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Error(err.Key, err.Message)
	}
	return
}
