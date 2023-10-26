package controller

import (
	"net/http"
	"pwdmanager_api/internal/auth"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/internal/helpers"
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
	if !helpers.CompareHashes(auth_input.Password, user.Password) ||
		user.Email != auth_input.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = db.DeleteAccount(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth.DeletionEmail(user)

	c.JSON(http.StatusNotFound, gin.H{"data": ""})
}

func UpdateAccount(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var pwd_update models.AuthUserPwd
	if err := c.ShouldBindJSON(&pwd_update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if !helpers.CompareHashes(pwd_update.OldPassword, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	err = db.UpdateAccount(user, pwd_update.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth.AccountUpdated(user)

	c.JSON(http.StatusOK, gin.H{"data": "", "updated_token": c.Request.Header.Get("Updated_Token")})
}
