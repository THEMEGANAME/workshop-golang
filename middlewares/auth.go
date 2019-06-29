package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authentication - this secure section
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.Request.Header.Get("token")
		if userToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "no access token",
			})
			return
		}
		c.Next()
	}
}
