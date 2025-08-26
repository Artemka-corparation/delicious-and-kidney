package main

import (
	"delicious-and-kidney/configs"
	"delicious-and-kidney/internal/auth"
	"delicious-and-kidney/internal/user"
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
	userRepo := user.NewUserRepository(database)
	jwtService := auth.NewJWTService(configs.LoadConfig().Auth.Secret)
	authService := auth.NewAuthService(authRepo, userRepo, jwtService)
	authHandler := auth.NewAuthHandler(authService)
	router := gin.Default()
	authHandler.RegisterRoutes(router)
	router.Run(":8080")
}
