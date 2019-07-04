package controllers

import (
	"github.com/bbcyyb/pcrs/infra/app"
	"github.com/bbcyyb/pcrs/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserContent struct {
	Id       string `json:"id" valid:"Required;Numeric"`
	RsaId    string `json:"rsaid" valid:"Required;MaxSize(100)"`
	UserName string `json:"username" valid:"Required;MaxSize(100)"`
	Email    string `json:"email" valid:"Email;MaxSize(100)"`
	Role     int    `json:"role" valid:"Range(0,2)"`
}

func GetAuth(c *gin.Context) {
	var (
		appG        = app.Gin{C: c}
		userContent UserContent
	)

	httpCode, errCode := app.BindAndValid(c, &userContent)
	if errCode != 200 {
		appG.Response(httpCode, errCode, nil)
		return
	}

	data := make(map[string]interface{})
	httpCode = 400
	isExist := true // TODO: check user exists or not.
	if isExist {
		var user = middlewares.User{
			Id:       userContent.Id,
			RsaId:    userContent.RsaId,
			UserName: userContent.UserName,
			Email:    userContent.Email,
			Role:     userContent.Role,
		}
		token, err := middlewares.GenerateToken(user)
		if err != nil {
			httpCode = 400
			appG.Response(http.StatusBadRequest, httpCode, nil)
		} else {
			data["token"] = token
			httpCode = 200
			appG.Response(http.StatusOK, httpCode, data)
		}
	} else {
		httpCode = 400
		appG.Response(http.StatusBadRequest, httpCode, nil)
	}
}
