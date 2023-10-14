package controller

import (
	"net/http"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/pkg/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created_user := db.CreateUser(user)
	c.JSON(http.StatusCreated, gin.H{"data": created_user})
}
