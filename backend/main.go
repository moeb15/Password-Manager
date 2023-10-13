package main

import (
	"log"
	"os"

	"pwdmanager_api/controller"
	"pwdmanager_api/database"

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

	public_routes := router.Group("/auth")
	public_routes.POST("/register", controller.Register(db))

	log.Fatal(router.Run(":8080"))
}
