package app

import (
	"gin-blog/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, 400
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError,500
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, 400
	}

	return http.StatusOK, e.SUCCESS
}