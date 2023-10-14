package middleware

import (
	"pwdmanager_api/internal/database"

	"github.com/gin-gonic/gin"
)

func DBMiddleware(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db_conn", db)
		c.Next()
	}
}
