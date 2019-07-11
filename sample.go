package main

import (
	"fmt"
	. "github.com/bbcyyb/pcrs/infra/logger"
)

func main1() {
	Logger.Info("hello")
	fmt.Print("world")
}
