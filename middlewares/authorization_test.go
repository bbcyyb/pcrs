package middlewares

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthorizationTestSuite struct {
	suite.Suite
	C *gin.Context
}

func TestAuthorizationSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}

func (suite *AuthorizationTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "/get", nil)
	context.Request = req

	suite.C = context
}
