package main

import (
	"delicious-and-kidney/configs"
	"delicious-and-kidney/internal/auth"
	"delicious-and-kidney/internal/user"

	"delicious-and-kidney/internal/restaurant"
	"delicious-and-kidney/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config := configs.LoadConfig()

	database, err := db.NewDb(config)
	if err != nil {
		log.Fatal("Not to connect to database", err)
	}
	err = database.AutoMigrate(&restaurant.Restaurant{})
	if err != nil {
		log.Fatal("Not to auto migrate", err)
	}

	userRepo := user.NewUserRepository(database)

	restaurantRepo := restaurant.NewRestaurantRepository(database)
	restaurantService := restaurant.NewRestaurantService(restaurantRepo, userRepo)
	authRepo := auth.NewAuthRepository(database)
	jvtService := auth.NewJWTService(config.Auth.Secret)
	authService := auth.NewAuthService(authRepo, userRepo, jvtService)
	restaurantHandler := restaurant.NewRestaurantHandler(restaurantService)
	router := gin.Default()
	restaurantHandler.RegisterRoutes(router, authService)
	router.Run(":8083")
}
