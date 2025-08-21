package main

import (
	"delicious-and-kidney/configs"
	"delicious-and-kidney/internal/auth"
	"delicious-and-kidney/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.LoadConfig()

	database, err := db.NewDb(config)
	if err != nil {
		log.Fatal("Not to auto migrate", err)
	}

	authRepo := auth.NewAuthRepository(database)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	router := gin.Default()
	api := router.Group("/api")
	{
		users := api.Group("/login")
		{
			users.GET("/:id", authHandler.Login)
		}
	}
	router.Run(":8081")
}
