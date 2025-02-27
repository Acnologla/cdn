package middlewares

import (
	"github.com/gin-gonic/gin"
)

func IsAdminMiddleware(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization != key {
			c.String(401, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
