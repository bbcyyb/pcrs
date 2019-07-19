package jwt

import (
	"testing"
	"time"

	"github.com/bbcyyb/pcrs/common"
	"github.com/stretchr/testify/suite"
)

type JWTTestSuite struct {
	suite.Suite

	JWT    *JWT
	Claims Claims
	Token  string
}

func TestJWTSuite(t *testing.T) {
	suite.Run(t, new(JWTTestSuite))
}

func (suite *JWTTestSuite) SetupSuite() {

	suite.JWT = NewJWT()
	suite.JWT.SetJwtSecret([]byte("DELLEMC"))
	//suite.JWT.SetJwtSecret([]byte("bbcyyb"))

	claims := Claims{
		Id:       111,
		RsaId:    "4B9C9747D0C328ADB06788DD61C2D658DAE2C0B301C629AE02936916805A7957B3BDE17BC86D417F1A16934D3CC2EBB5B825B64238C3C8CA3C7068D91AFD0BBDEA01187FA7C345C6357A9703BE2EF587E5E1281B6FA73634CE46AB37830BDEC1",
		UserName: "Dev",
		Email:    "dev@dell.com",
		Role:     common.USERROLE_ADMIN,
		IsDebug:  0,
	}

	expireTime := time.Date(2020, 9, 27, 16, 30, 0, 0, time.Local)

	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = "powercalculator"

	suite.Claims = claims
	suite.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTExLCJyaWQiOiI0QjlDOTc0N0QwQzMyOEFEQjA2Nzg4REQ2MUMyRDY1OERBRTJDMEIzMDFDNjI5QUUwMjkzNjkxNjgwNUE3OTU3QjNCREUxN0JDODZENDE3RjFBMTY5MzREM0NDMkVCQjVCODI1QjY0MjM4QzNDOENBM0M3MDY4RDkxQUZEMEJCREVBMDExODdGQTdDMzQ1QzYzNTdBOTcwM0JFMkVGNTg3RTVFMTI4MUI2RkE3MzYzNENFNDZBQjM3ODMwQkRFQzEiLCJ1biI6IkRldiIsImVtIjoiZGV2QGRlbGwuY29tIiwidXIiOjIsImRlIjowLCJleHAiOjE2MDExOTU0MDAsImlzcyI6InBvd2VyY2FsY3VsYXRvciJ9.sQPjfOM1UCZehjEcN45SRQtcMSbi-DY1zWFivkADXL8"
}

func (suite *JWTTestSuite) TestGenerateToken() {
	ass := suite.Assert()

	token, err := suite.JWT.Generate(&(suite.Claims))

	if ass.Nil(err) {
		ass.Equal(suite.Token, token)
	}
}

func (suite *JWTTestSuite) TestGenerateTokenFailed() {
	ass := suite.Assert()

	token, err := suite.JWT.Generate(nil)

	ass.Error(err)
	ass.Empty(token)
}

func (suite *JWTTestSuite) TestParseToken() {
	ass := suite.Assert()
	claims, err := suite.JWT.Parse(suite.Token)
	expectedClaims := suite.Claims

	if ass.Nil(err) {
		ass.Equal(*claims, expectedClaims)
	}
}

func (suite *JWTTestSuite) TestParseTokenFailed() {
	ass := suite.Assert()

	incorrectToken := suite.Token + "123"

	if claims, err := suite.JWT.Parse(incorrectToken); ass.Nil(claims) && ass.Error(err) {
		ass.Equal("signature is invalid", err.Error())
	}
}
