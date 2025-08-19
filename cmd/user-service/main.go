package main

import (
	"delicious-and-kidney/configs"
	"delicious-and-kidney/internal/user"
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
	err = database.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Not to auto migrate", err)
	}
	userRepo := user.NewUserRepository(database)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetUser)
			users.PATCH("/:id", userHandler.UpdateUser)
		}
	}
	router.Run(":8080")
}
