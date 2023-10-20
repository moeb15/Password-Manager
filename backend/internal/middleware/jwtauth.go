package middleware

import (
	"fmt"
	"net/http"
	"pwdmanager_api/internal/auth"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateJWT(c)
		if err != nil {
			err = auth.ValidateRefreshToken(c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			rf_token, err := auth.GetRefreshToken(c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			access_token, err := auth.JWTFromRefresh(rf_token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))
			c.Request.Header.Set("Updated_Token", access_token)
			c.Next()
		}
		c.Request.Header.Set("Updated_Token", "")
		c.Next()
	}
}
