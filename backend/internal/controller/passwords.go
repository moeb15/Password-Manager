package controller

import (
	"fmt"
	"log"
	"net/http"
	"pwdmanager_api/internal/auth"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/internal/helpers"
	"pwdmanager_api/pkg/models"

	"github.com/gin-gonic/gin"
)

const MAX_PWD = 50

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
	count, err := db.NumPasswwords(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if count >= MAX_PWD {
		c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("password limit of %d reached", MAX_PWD)})
		return
	}

	enc_pwd, err := helpers.EncryptAES([]byte(pwd.Key), pwd.Password)
	if err != nil {
		log.Fatal(err)
	}

	saved_pwd := models.Password{
		UserID:      user.ID,
		Application: pwd.Application,
		Username:    pwd.Username,
		Password:    enc_pwd,
	}
	db.CreatePassword(saved_pwd, user)
	c.JSON(http.StatusCreated, gin.H{"data": "", "updated_token": c.Request.Header.Get("Updated_Token")})
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
	c.JSON(http.StatusFound, gin.H{"data": pwds, "updated_token": c.Request.Header.Get("Updated_Token")})
}

func DeletePassword(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	app_name := c.DefaultQuery("app", "")
	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	res, err := db.DeleteByApp(app_name, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{"data": "no application found"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"data": "", "updated_token": c.Request.Header.Get("Updated_Token")})
}

func GetPassword(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var pwd_req models.AuthGetPwd
	if err := c.ShouldBindJSON(&pwd_req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !helpers.CompareHashes(pwd_req.Key, user.MasterKey) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	pwd := db.RetrieveByApp(pwd_req.Application, pwd_req.Username, user.ID)
	raw_pwd, err := helpers.DecryptAES([]byte(pwd_req.Key), pwd.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// removes any characters that are invalid utf8 characters
	pwd.Password = string([]rune(raw_pwd))
	c.JSON(http.StatusFound, gin.H{"data": pwd, "updated_token": c.Request.Header.Get("Updated_Token")})
}

func UpdatePassword(c *gin.Context) {
	db := c.MustGet("db_conn").(*database.DB)
	var pwd_input models.AuthPwd
	if err := c.ShouldBindJSON(&pwd_input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := auth.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !helpers.CompareHashes(pwd_input.Key, user.MasterKey) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key"})
		return
	}

	new_pwd, err := helpers.EncryptAES([]byte(pwd_input.Key), pwd_input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = db.UpdatePassword(pwd_input.Application, new_pwd, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "", "updated_token": c.Request.Header.Get("Updated_Token")})
}
