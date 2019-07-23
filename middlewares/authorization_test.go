package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bbcyyb/pcrs/common"
	pkgA "github.com/bbcyyb/pcrs/pkg/authorizer"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthorizationTestSuite struct {
	suite.Suite
	C           *gin.Context
	AuthorizerM pkgA.IAuthorizer
}

type authorizerMock struct {
	mock.Mock
}

func (m *authorizerMock) CheckPermission(user string, r *http.Request) bool {
	args := m.Called(user, r)
	return args.Bool(0)
}

func TestAuthorizationSuite(t *testing.T) {
	suite.Run(t, new(AuthorizationTestSuite))
}

func (suite *AuthorizationTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "/get", nil)
	context.Request = req

	suite.C = context

	authorizerM := &authorizerMock{}

	authorizerM.On("CheckPermission", "user", suite.C.Request).Return(false)
	authorizerM.On("CheckPermission", "admin", suite.C.Request).Return(true)
	authorizerM.On("CheckPermission", "desiginer", suite.C.Request).Return(false)
	authorizerM.On("CheckPermission", "commnon", suite.C.Request).Return(false)

	suite.AuthorizerM = authorizerM
}

func (suite *AuthorizationTestSuite) TestAuthorization() {
	ass := suite.Assert()

	claims := &pkgJ.Claims{
		Id:       111,
		UserName: "Dev",
		RsaId:    "Dev",
		Email:    "dev@dell.com",
		Role:     common.USERROLE_ADMIN,
		IsDebug:  0,
	}

	suite.C.Set("claims", claims)

	Authorization(suite.AuthorizerM)(suite.C)

	ass.Nil(suite.C.Errors)
}

func (suite *AuthorizationTestSuite) TestAuthorizationFailed() {
	ass := suite.Assert()

	claims := &pkgJ.Claims{
		Id:       111,
		UserName: "Dev",
		RsaId:    "Dev",
		Email:    "dev@dell.com",
		Role:     common.USERROLE_USER,
		IsDebug:  0,
	}

	suite.C.Set("claims", claims)

	Authorization(suite.AuthorizerM)(suite.C)

	ass.NotNil(suite.C.Errors)
}
