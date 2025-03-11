package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func (mw *MiddlewareManager) RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		req := c.Request
		res := c.Writer
		status := res.Status()
		size := res.Size()
		duration := time.Since(start).String()

		mw.logger.Infof("Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			req.Method, req.URL, status, size, duration,
		)
	}
}
