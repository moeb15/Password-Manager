package auth

import (
	"errors"
	"fmt"
	"os"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/pkg/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var private_key = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user models.User) (string, error) {
	token_ttl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(token_ttl)).Unix(),
	})

	return token.SignedString(private_key)
}

func GenerateRefreshToken(user models.User) (string, error) {
	rf_ttl, _ := strconv.Atoi(os.Getenv("REFRESH_TTL"))
	rf_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(rf_ttl)).Unix(),
	})

	return rf_token.SignedString(private_key)
}

func JWTFromRefresh(rf_token *jwt.Token) (string, error) {
	claims, _ := rf_token.Claims.(jwt.MapClaims)
	user_id := claims["id"].(string)

	token_ttl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user_id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(token_ttl)).Unix(),
	})

	return token.SignedString(private_key)
}

func CurrentUser(c *gin.Context) (models.User, error) {
	err := ValidateJWT(c)
	db := c.MustGet("db_conn").(*database.DB)
	if err != nil {
		return models.User{}, err
	}
	token, _ := getToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	user_id := claims["id"].(string)

	user, err := db.FindUser(user_id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func ValidateJWT(c *gin.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func ValidateRefreshToken(c *gin.Context) error {
	rf_token, err := GetRefreshToken(c)
	if err != nil {
		return err
	}
	_, ok := rf_token.Claims.(jwt.MapClaims)
	if ok && rf_token.Valid {
		return nil
	}

	return errors.New("invalid refresh token provided")
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	token_str := getTokenFromRequest(c)
	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return private_key, nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearer_token := c.Request.Header.Get("Authorization")
	split_token := strings.Split(bearer_token, " ")
	if len(split_token) == 2 {
		return split_token[1]
	}
	return ""
}

func GetRefreshToken(c *gin.Context) (*jwt.Token, error) {
	token_str := getRefreshTokenFromRequest(c)
	rf_token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return private_key, nil
	})
	return rf_token, err
}

func getRefreshTokenFromRequest(c *gin.Context) string {
	refresh_token := c.Request.Header["Refresh"][0]
	return refresh_token
}
