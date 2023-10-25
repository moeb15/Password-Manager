package controller

import (
	"net/http"
	"pwdmanager_api/internal/auth"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/pkg/models"

	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var auth_input models.AuthInput
	if err := c.ShouldBindJSON(&auth_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = db.DeleteAccount(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"data": ""})
}
