package controller

import (
	"log"
	"net/http"
	"pwdmanager_api/internal/auth"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/internal/helpers"
	"pwdmanager_api/pkg/models"

	"github.com/gin-gonic/gin"
)

func AddPassword(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var pwd models.AuthPwd
	if err := c.ShouldBindJSON(&pwd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok := helpers.CompareHashes(pwd.Key, user.MasterKey); !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid key"})
		return
	}

	enc_pwd, err := helpers.EncryptAES([]byte(pwd.Key), pwd.Password)
	if err != nil {
		log.Fatal(err)
	}

	saved_pwd := models.Password{
		UserID:      user.ID,
		Application: pwd.Application,
		Password:    enc_pwd,
	}
	user.SavedPwds = append(user.SavedPwds, saved_pwd)
	db.CreatePassword(saved_pwd, user)
	c.JSON(http.StatusCreated, gin.H{"data": ""})
}

func GetPasswords(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	pwds, err := db.FindPasswords(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusFound, gin.H{"data": pwds})
}
