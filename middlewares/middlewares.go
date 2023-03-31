package middlewares

import (
	"culineira-backend/helpers/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, please login first")
			c.Abort()
			return
		}
		c.Next()
	}
}
