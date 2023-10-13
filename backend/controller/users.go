package controller

import (
	"net/http"
	"pwdmanager_api/database"
	"pwdmanager_api/models"

	"github.com/gin-gonic/gin"
)

func Register(db *database.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		created_user := db.CreateUser(user)
		c.JSON(http.StatusCreated, gin.H{"data": created_user})
	}

	return gin.HandlerFunc(fn)
}
