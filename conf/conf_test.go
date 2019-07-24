package conf

import (
	"testing"

	"github.com/bbcyyb/pcrs/common"
	"github.com/stretchr/testify/suite"
)

type ConfTestSuite struct {
	suite.Suite
	ExpectedConfig *Config
}

func TestConfSuite(t *testing.T) {
	suite.Run(t, new(ConfTestSuite))
}

func (suite *ConfTestSuite) SetupSuite() {
	suite.ExpectedConfig = &Config{
		DB: Database{
			Server:   "10.35.83.61:1433",
			User:     "PowerCalc",
			Password: "Power@1433",
			Database: "PowerCalcFor46",
			Source:   "file://scripts",
		},
		Middleware: Middleware{
			Authentication: AuthT{
				Enable: true,
			},
			Authorization: AuthZ{
				Enable: true,
			},
			ErrorHandler: ErrorHandler{
				Enable: true,
			},
		},
		Pkg: Pkg{
			Authorizer: Authorizer{
				Policy: "authpolicy.csv",
				Model:  "authmodel.conf",
			},
			Log: Log{
				Level:     "debug",
				Formatter: "text",
				Prefix:    "pcrs",
				FileName:  "pcrs.log",
			},
		},
	}
}

func (suite *ConfTestSuite) TestConfFileLoading() {
	ass := suite.Assert()
	configPath := common.BuildRunningPath("config.yaml")
	c := &Config{}
	initConfig(configPath, c)

	ass.EqualValues(suite.ExpectedConfig, c)
}
