package middlewares

import (
	"errors"
	"net/http"

	"github.com/bbcyyb/pcrs/common"
	pkgA "github.com/bbcyyb/pcrs/pkg/authorizer"
	pkgJ "github.com/bbcyyb/pcrs/pkg/jwt"
	"github.com/bbcyyb/pcrs/pkg/log"
	"github.com/gin-gonic/gin"
)

func Authorization(auth pkgA.IAuthorizer) gin.HandlerFunc {
	log.Debug("Register middleware Authorization")

	return func(c *gin.Context) {
		role := getRole(c)
		roleName := common.GetRoleEnumMessage(role)
		if !auth.CheckPermission(roleName, c.Request) {
			err := errors.New(common.GetCodeMessage(common.ERROR_AUTHZ_CHECK_PERMISSION_FAIL))
			c.AbortWithError(http.StatusForbidden, err)
		}

		c.Next()
	}
}

func getRole(c *gin.Context) common.RoleEnum {
	if value, ok := c.Get("claims"); ok {
		claims := value.(*(pkgJ.Claims))
		role := claims.Role
		log.Debug("session role ----> ", common.GetRoleEnumMessage(role))
		return role
	}

	log.Debug("default role ---> User")
	return common.USERROLE_USER
}
