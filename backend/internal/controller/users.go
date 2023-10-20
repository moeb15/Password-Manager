package controller

import (
	"net/http"
	"pwdmanager_api/internal/auth"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/internal/helpers"
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

	created_user, err := db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": *created_user})
}

func Login(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var user_auth models.AuthInput
	if err := c.ShouldBindJSON(&user_auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.FindUserByName(user_auth.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if ok := helpers.CompareHashes(user_auth.Password, user.Password); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	jwt, err := auth.GenerateJWT(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	rf_jwt, err := auth.GenerateRefreshToken(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"access_token": jwt, "refresh_token": rf_jwt})
}
