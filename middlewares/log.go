package middlewares

import (
	"bytes"
	"io"
	"io/ioutil"
	"time"

	"github.com/bbcyyb/pcrs/infra/log"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	log.Debug("Register middleware Log")
	return func(c *gin.Context) {
		start := time.Now()
		url := c.Request.URL.String()
		body := ""
		if c.Request.Body != nil {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
			body = readBody(rdr1)
			c.Request.Body = rdr2
		}

		c.Next()

		end := time.Since(start)
		status := c.Writer.Status()

		entryHandler := log.WithFields(log.Fields{
			"url":      url,
			"method":   c.Request.Method,
			"body":     body,
			"clientIP": c.ClientIP(),
			"status":   status,
			"size":     c.Writer.Size(),
			"latency":  float64(end.Seconds()) * 1000.0,
		})

		if len(c.Errors) > 0 {
			entryHandler.Error(c.Errors.String())
		} else {
			if status > 499 {
				entryHandler.Error()
			} else {
				entryHandler.Info()
			}
		}
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
