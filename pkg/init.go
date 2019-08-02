package pkg

import (
	"github.com/bbcyyb/pcrs/pkg/database"
	"github.com/bbcyyb/pcrs/pkg/logger"
)

func Setup() {
	logger.Setup()
	database.Setup()
}
