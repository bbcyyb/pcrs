package authorizer

import (
	"net/http"
	"path"
	"testing"

	"github.com/bbcyyb/pcrs/common"
	"github.com/bbcyyb/pcrs/conf"
	"github.com/stretchr/testify/suite"
)

type AuthorizerTestSuite struct {
	suite.Suite
	Auth          *BasicAuthorizer
	PositiveCases []Case
	NegativeCases []Case
}

type Case struct {
	Role   string
	Route  string
	Method string
}

func TestAuthorizerSuite(t *testing.T) {
	suite.Run(t, new(AuthorizerTestSuite))
}

func (suite *AuthorizerTestSuite) SetupSuite() {
	ass := suite.Assert()

	dir := common.BuildRunningPath("")
	conf.C = &conf.Config{}
	conf.C.Pkg.Authorizer = conf.Authorizer{
		Policy: path.Join(dir, "auth_policy_test.csv"),
		Model:  path.Join(dir, "auth_model_test.conf"),
	}

	ass.FileExists(conf.C.Pkg.Authorizer.Model)
	ass.FileExists(conf.C.Pkg.Authorizer.Policy)

	suite.Auth = NewBasicAuthorizer()
	ass.NotNil(suite.Auth)

	suite.PositiveCases = []Case{
		Case{"admin", "/get", "GET"},
		Case{"admin", "/any", "POST"},
		Case{"user", "/get", "GET"},
		Case{"user", "/post", "POST"},
		Case{"anonymity", "/login", "GET"},
	}

	suite.NegativeCases = []Case{
		Case{"admin", "/get", "HEAD"},
		Case{"design", "/any", "POST"},
		Case{"user", "/put", "HEAD"},
		Case{"anonymity", "/post", "POST"},
		Case{"anonymity", "/login", "POST"},
	}
}

func (suite *AuthorizerTestSuite) TestCheckPermission() {
	ass := suite.Assert()

	for i, c := range suite.PositiveCases {
		req, _ := http.NewRequest(c.Method, c.Route, nil)
		ass.Truef(suite.Auth.CheckPermission(c.Role, req), "error sequence number: %v", i+1)
	}
}

func (suite *AuthorizerTestSuite) TestCheckIncorrectPermission() {
	ass := suite.Assert()

	for i, c := range suite.NegativeCases {
		req, _ := http.NewRequest(c.Method, c.Route, nil)
		ass.Falsef(suite.Auth.CheckPermission(c.Role, req), "error sequence number: %v", i)
	}
}
