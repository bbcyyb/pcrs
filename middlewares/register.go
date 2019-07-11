package middlewares

import (
	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"

	"bytes"
	"errors"
)

type extendResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *extendResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Register(group *gin.RouterGroup, authGroup *gin.RouterGroup) {
	log.Info("Register middlewares")

	group.Use(func(c *gin.Context) {
		w := &extendResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		originWriter := c.Writer
		c.Writer = w
		log.Info("=======> 111 begin")

		c.Next()

		body := w.body
		str := body.String()
		log.Info(body)
		log.Info(str)

		c.Writer = originWriter
		log.Info("=======> 111 end")
	})

	group.Use(func(c *gin.Context) {
		log.Info("=======> 222 begin")

		c.AbortWithError(400, errors.New("EERRRRRRRRRRRRRRRRRRRRRRRRRRROR"))
		c.Next()

		log.Info("=======> 222 end")
	})

	group.Use(func(c *gin.Context) {
		log.Info("=======> 333 begin")

		c.Next()

		log.Info("=======> 333 end")
	})

	/*
		if C.AuthTEnable {
			authGroup.Use(Authentication())
		}

		if C.AuthREnable {
			authGroup.Use(Authorization())
		}

		if C.LogEnable {
			authGroup.Use(Log())
			group.Use(Log())
		}
	*/
}
