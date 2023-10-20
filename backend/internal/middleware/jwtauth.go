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
			c.JSON(http.StatusCreated, gin.H{"access_token": access_token})
			c.Request.Header.Set("Authroization", access_token)
		}
		c.Next()
	}
}
