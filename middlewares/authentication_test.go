package middlewares

import (
	"net/http/httptest"
	"testing"

	"github.com/bbcyyb/pcrs/common"
	infraJ "github.com/bbcyyb/pcrs/infra/jwt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthenticationTestSuite struct {
	suite.Suite
}

func TestAuthenticationSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationTestSuite))
}

func (suite *AuthenticationTestSuite) SetupSuite() {
}

func (suite *AuthenticationTestSuite) Test() {
	ass := suite.Assert()
	gin.SetMode(gin.TestMode)

	recorderPass := httptest.NewRecorder()
	contextPass, _ := gin.CreateTestContext(recorderPass)

	//recorderFail := httptest.NewRecorder()
	//contextFail, _ := gin.CreateTestContext(recorderFail)

	reqPass := httptest.NewRequest("GET", "/get", nil)
	reqPass.Header.Set("X-Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTExLCJyaWQiOiI0QjlDOTc0N0QwQzMyOEFEQjA2Nzg4REQ2MUMyRDY1OERBRTJDMEIzMDFDNjI5QUUwMjkzNjkxNjgwNUE3OTU3QjNCREUxN0JDODZENDE3RjFBMTY5MzREM0NDMkVCQjVCODI1QjY0MjM4QzNDOENBM0M3MDY4RDkxQUZEMEJCREVBMDExODdGQTdDMzQ1QzYzNTdBOTcwM0JFMkVGNTg3RTVFMTI4MUI2RkE3MzYzNENFNDZBQjM3ODMwQkRFQzEiLCJ1biI6IkRldiIsImVtIjoiZGV2QGRlbGwuY29tIiwidXIiOjIsImRlIjowLCJleHAiOjE2MDExOTU0MDAsImlzcyI6InBvd2VyY2FsY3VsYXRvciJ9.sQPjfOM1UCZehjEcN45SRQtcMSbi-DY1zWFivkADXL8")
	contextPass.Request = reqPass

	Authentication()(contextPass)

	value := contextPass.MustGet("claims")
	if ass.NotNil(value) {
		claims := value.(*infraJ.Claims)
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
