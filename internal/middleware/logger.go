package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UnixMicro()
		c.Next()
		duration := time.Now().UnixMicro() - start

		entry := log.WithFields(log.Fields{
			"client_ip": c.ClientIP(),
			"duration":  duration,
			"method":    c.Request.Method,
			"path":      c.Request.RequestURI,
			"status":    c.Writer.Status(),
			"referrer":  c.Request.Referer(),
		})
		if c.Writer.Status() >= 400 {
			entry.Error("")
		} else {
			entry.Info("")
		}
	}
}
