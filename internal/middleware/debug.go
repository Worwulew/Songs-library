package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func (mw *MiddlewareManager) DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if mw.cfg.Server.Debug {
			dump, err := httputil.DumpRequest(c.Request, true)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			mw.logger.Info(fmt.Sprintf("\nRequest dump begin :--------------\n\n%s\n\nRequest dump end :--------------", dump))
		}
		c.Next()
	}
}
