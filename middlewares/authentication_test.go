package middlewares

import (
	"net/http/httptest"
	"testing"

	"github.com/bbcyyb/pcrs/common"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthenticationTestSuite struct {
	suite.Suite
	C *gin.Context
}

type jwtParserMock struct {
	mock.Mock
}

func (m *jwtParserMock) Parse(token string) (*pkgJ.Claims, error) {
	args := m.Called(token)

	return args.Get(0).(*pkgJ.Claims), args.Error(1)
}

func (m *jwtParserMock) SetJwtSecret(jwtSecret []byte) {
	m.Called(jwtSecret)
}

func TestAuthenticationSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}

func (suite *AuthenticationTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "/get", nil)
	context.Request = req

	suite.C = context
}

func (suite *AuthenticationTestSuite) TestAuthentication() {
	ass := suite.Assert()

	suite.C.Request.Header.Set("X-Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTExLCJyaWQiOiI0QjlDOTc0N0QwQzMyOEFEQjA2Nzg4REQ2MUMyRDY1OERBRTJDMEIzMDFDNjI5QUUwMjkzNjkxNjgwNUE3OTU3QjNCREUxN0JDODZENDE3RjFBMTY5MzREM0NDMkVCQjVCODI1QjY0MjM4QzNDOENBM0M3MDY4RDkxQUZEMEJCREVBMDExODdGQTdDMzQ1QzYzNTdBOTcwM0JFMkVGNTg3RTVFMTI4MUI2RkE3MzYzNENFNDZBQjM3ODMwQkRFQzEiLCJ1biI6IkRldiIsImVtIjoiZGV2QGRlbGwuY29tIiwidXIiOjIsImRlIjowLCJleHAiOjE2MDExOTU0MDAsImlzcyI6InBvd2VyY2FsY3VsYXRvciJ9.sQPjfOM1UCZehjEcN45SRQtcMSbi-DY1zWFivkADXL8")

	jwt := new(jwtParserMock)
	Authentication(jwt)(suite.C)

	value := suite.C.MustGet("claims")
	if ass.NotNil(value) {
		claims := value.(*pkgJ.Claims)
		ass.NotNil(claims)
		ass.Equal(111, claims.Id)
		ass.Equal("Dev", claims.UserName)
		ass.Equal("4B9C9747D0C328ADB06788DD61C2D658DAE2C0B301C629AE02936916805A7957B3BDE17BC86D417F1A16934D3CC2EBB5B825B64238C3C8CA3C7068D91AFD0BBDEA01187FA7C345C6357A9703BE2EF587E5E1281B6FA73634CE46AB37830BDEC1", claims.RsaId)
		ass.Equal("dev@dell.com", claims.Email)
		ass.Equal(common.USERROLE_ADMIN, claims.Role)
		ass.Equal(0, claims.IsDebug)
		ass.Equal("powercalculator", claims.Issuer)
		ass.EqualValues(1601195400, claims.ExpiresAt)
	}
}

func (suite *AuthenticationTestSuite) TestAuthenticationToMatchLegacyTokenFormat() {
	ass := suite.Assert()

	suite.C.Request.Header.Set("X-Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTA0ODYsInJpZCI6IjExMTE2MzUiLCJ1biI6IktldmluIFlhYmluZyIsImVtIjoiS2V2aW4uWS5ZdUBlbWMuY29tIiwidXIiOjIsImV4cCI6MTU2NDEwNzc5NzI3OSwiZGUiOjB9.15eK9C0KqQWIA7JbLZVqYgz3gtdkgIykF1tLqnpg57A")

	Authentication(nil)(suite.C)

	value := suite.C.MustGet("claims")
	if ass.NotNil(value) {
		claims := value.(*pkgJ.Claims)
		ass.NotNil(claims)
		ass.Equal(10486, claims.Id)
		ass.Equal("Kevin Yabing", claims.UserName)
		ass.Equal("1111635", claims.RsaId)
		ass.Equal("Kevin.Y.Yu@emc.com", claims.Email)
		ass.Equal(common.USERROLE_ADMIN, claims.Role)
		ass.Equal(0, claims.IsDebug)
		ass.Empty(claims.Issuer)
		ass.EqualValues(1564107797, claims.ExpiresAt)
	}
}

func (suite *AuthenticationTestSuite) TestAuthenticationFail() {
	ass := suite.Assert()

	suite.C.Request.Header.Set("X-Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTExLCJyaWQiOiI0QjlDOTc0N0QwQzMyOEFEQjA2Nzg4REQ2MUMyRDY1OERBRTJDMEIzMDFDNjI5QUUwMjkzNjkxNjgwNUE3OTU3QjNCREUxN0JDODZENDE3RjFBMTY5MzREM0NDMkVCQjVCODI1QjY0MjM4QzNDOENBM0M3MDY4RDkxQUZEMEJCREVBMDExODdGQTdDMzQ1QzYzNTdBOTcwM0JFMkVGNTg3RTVFMTI4MUI2RkE3MzYzNENFNDZBQjM3ODMwQkRFQzEiLCJ1biI6IkRldiIsImVtIjoiZGV2QGRlbGwuY29tIiwidXIiOjIsImRlIjowLCJleHAiOjE2MDExOTU0MDAsImlzcyI6InBvd2VyY2FsY3VsYXRvciJ9.sQPjfOM1UCZehjEcN45SRQtcMSbi-DY1zWFivkADXLa")

	Authentication(nil)(suite.C)

	ass.Contains(suite.C.Errors.String(), common.GetCodeMessage(common.ERROR_AUTHT_CHECK_TOKEN_FAIL))

	value, exists := suite.C.Get("claims")
	ass.False(exists)
	ass.Nil(value)
}

func (suite *AuthenticationTestSuite) TestAuthenticationMiss() {
	ass := suite.Assert()

	Authentication(nil)(suite.C)

	ass.Contains(suite.C.Errors.String(), common.GetCodeMessage(common.ERROR_AUTHT_CHECK_TOKEN_MISS))

	value, exists := suite.C.Get("claims")
	ass.False(exists)
	ass.Nil(value)
}
