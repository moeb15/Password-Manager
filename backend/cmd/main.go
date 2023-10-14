package main

import (
	"log"
	"os"

	"pwdmanager_api/internal/controller"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	db := database.Connect(os.Getenv("DB_URL"))
	serveApplication(db)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables")
	}
}

func serveApplication(db *database.DB) {
	router := gin.Default()
	router.Use(middleware.DBMiddleware(db))
	public_routes := router.Group("/auth")
	public_routes.POST("/register", controller.Register)

	log.Fatal(router.Run(":8080"))
}
