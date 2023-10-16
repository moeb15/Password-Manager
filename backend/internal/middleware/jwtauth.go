package middleware

import (
	"net/http"
	"pwdmanager_api/internal/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
