package main

import (
	"log"
	"os"
	"time"

	"pwdmanager_api/internal/controller"
	"pwdmanager_api/internal/database"
	"pwdmanager_api/internal/middleware"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Refresh"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	public_routes := router.Group("/auth")

	public_routes.POST("/register", controller.Register)
	public_routes.POST("/login", controller.Login)

	private_routes := router.Group("/api")
	private_routes.Use(middleware.JWTAuthMiddleWare())

	// Password related routes
	private_routes.GET("/pwd", controller.GetPasswords)
	private_routes.POST("/pwd", controller.AddPassword)
	private_routes.POST("/pwd/decrypt", controller.GetPassword)
	private_routes.PATCH("/pwd", controller.UpdatePassword)
	private_routes.DELETE("/pwd", controller.DeletePassword)

	// User related routes
	private_routes.POST("/user/remove", controller.DeleteAccount)
	private_routes.POST("/user/update", controller.UpdateAccount)

	log.Fatal(router.Run(":8080"))
}
