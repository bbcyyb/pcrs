package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ResponseTestSuite struct {
	suite.Suite

	c   *gin.Context
	r   *httptest.ResponseRecorder
	err error
}

func TestResponseSuite(t *testing.T) {
	suite.Run(t, new(ResponseTestSuite))
}

func (suite *ResponseTestSuite) SetupTest() {
	gin.SetMode(gin.ReleaseMode)
	suite.r = httptest.NewRecorder()
	c, _ := gin.CreateTestContext(suite.r)
	suite.c = c
	err := errors.New("test error")
	suite.err = err
}

func (suite *ResponseTestSuite) TestOK() {
	ass := suite.Assert()

	OK(suite.c, gin.H{"key": "value"})
	ass.Equal(http.StatusOK, suite.c.Writer.Status(), "http status code error")

	j := suite.r.Body.Bytes()
	r := Response{}
	if err := json.Unmarshal(j, &r); ass.NoError(err) {
		ass.Equal(http.StatusOK, r.Code, "json data error")
		ass.Equal("value", r.Data.(map[string]interface{})["key"].(string))
	}
}

func (suite *ResponseTestSuite) TestNoContent() {
	ass := suite.Assert()

	NoContent(suite.c)
	ass.Equal(http.StatusNoContent, suite.c.Writer.Status(), "http status code error")

	j := suite.r.Body.Bytes()
	ass.Nil(j)
}

func (suite *ResponseTestSuite) TestBadRequest() {
	ass := suite.Assert()

	BadRequest(suite.c, suite.err)
	ass.Equal(http.StatusBadRequest, suite.c.Writer.Status(), "http status code error")

	j := suite.r.Body.Bytes()
	r := Response{}
	if err := json.Unmarshal(j, &r); ass.NoError(err) {
		ass.Equal(http.StatusBadRequest, r.Code, "json data error")
		ass.Equal("test error", r.Message)
	}
}

func (suite *ResponseTestSuite) TestUnauthorized() {
	ass := suite.Assert()

	Unauthorized(suite.c, suite.err)
	ass.Equal(http.StatusUnauthorized, suite.c.Writer.Status(), "http status code error")

	j := suite.r.Body.Bytes()
	r := Response{}
	if err := json.Unmarshal(j, &r); ass.NoError(err) {
		ass.Equal(http.StatusUnauthorized, r.Code, "json data error")
		ass.Equal("test error", r.Message)
	}
}
