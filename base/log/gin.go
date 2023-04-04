package log

import (
	"github.com/gin-gonic/gin"
	"time"
)

func GetGinLogHandler(logger *Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		entry := logger.WithFields(map[string]interface{}{
			"latency":     time.Now().Sub(start),
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"status_code": c.Writer.Status(),
			"body_size":   c.Writer.Size(),
		})
		if raw != "" {
			entry.WithField("path", path+"?"+raw)
		} else {
			entry.WithField("path", path)
		}
		entry.Infoln()
	}
}
