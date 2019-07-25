package misc

import (
	"time"

	"github.com/bbcyyb/pcrs/common"
	"github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type MiscController struct {
	Service IService
}

func NewMiscController(service IService) *MiscController {
	return &MIscController{
		Service: service,
	}
}

func (misc *MiscController) Test(c *gin.Context) {
	common.OK(c, gin.H{
		"message": "Hello World!",
	})
}

func (misc *MiscController) GenerateTestToken(c *gin.Context) {
	var claims jwt.Claims
	if err := c.BindJSON(&claims); err != nil {
		common.BadRequest(c, err)
		return
	}

	token, err := misc.Service.GenerateToken(nil, &claims)
	if err != nil {
		common.InternalServerError(c, err)
	} else {
		common.OK(c, gin.H{
			"token": token,
		})
	}
}
