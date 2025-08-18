package main

import (
	"delicious-and-kidney/configs"
	"delicious-and-kidney/internal/user"
	"delicious-and-kidney/pkg/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Загружаем конфигурацию
	config := configs.LoadConfig()

	// Подключаемся к БД
	database, err := db.NewDb(config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Создаем репозиторий и хендлер
	userRepo := user.NewUserRepository(database)
	userHandler := user.NewUserHandler(userRepo)

	// Настраиваем роуты
	router := gin.Default()
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", userHandler.GetUser)
			users.POST("/", userHandler.CreateUser)
		}
	}

	// Запускаем сервер
	log.Println("Server starting on :8080")
	router.Run(":8080")
}
