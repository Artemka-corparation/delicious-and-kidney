package main

import (
	"context"
	"delicious-and-kidney/internal/database"
	"delicious-and-kidney/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println(" Запуск приложения...")
	database.InitDatabase()  //подключил
	defer database.CloseDB() //после работы программы закрыл
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis недоступен %v", err)
		log.Printf("Приложение работает без корзины!")
		rdb = nil
	} else {
		log.Printf("Redis подключен")
		defer rdb.Close()
	}

	if os.Getenv("GIN_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/health", healthCheck)
	api := router.Group("/api/v1")
	api.GET("/products", getProducts)
	api.GET("/categories", getCategories)
	api.GET("/cart", func(c *gin.Context) {})
	api.POST("/cart", func(c *gin.Context) {})

	port := getEnv("PORT", "8080")
	log.Printf("Сервер запущен на http://localhost:%s", port)
	router.Run(":" + port)

}
func healthCheck(c *gin.Context) {
	if err := models.CheckConnection(database.GetDB()); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "Server is running",
		"database": "connected",
	})
}
func getCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.GetDB().Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  categories,
		"count": len(categories),
	})
}
func getProducts(c *gin.Context) {
	var products []models.Product

	if err := database.GetDB().Preload("Category").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
