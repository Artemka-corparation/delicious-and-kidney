package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"food-delivery-app/internal/database"
	"food-delivery-app/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🚀 Запуск приложения...")

	// Инициализация базы данных
	database.InitDatabase()
	defer database.CloseDB()

	// Проверка подключения
	if err := models.CheckConnection(database.GetDB()); err != nil {
		log.Fatal("❌ Не удалось подключиться к БД:", err)
	}

	// Настройка Gin
	if os.Getenv("GIN_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := setupRouter()

	// Запуск сервера
	port := getEnv("PORT", "8080")
	log.Printf("🌐 Сервер запущен на http://localhost:%s", port)
	log.Printf("📋 Health check: http://localhost:%s/health", port)
	log.Printf("📦 API: http://localhost:%s/api/v1/", port)

	// Graceful shutdown
	go func() {
		if err := router.Run(":" + port); err != nil {
			log.Fatal("❌ Не удалось запустить сервер:", err)
		}
	}()

	// Ожидание сигнала для завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы сервера...")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Health check
	router.GET("/health", healthCheck)

	// API группы
	api := router.Group("/api/v1")
	{
		api.GET("/categories", getCategories)
		api.GET("/products", getProducts)
		api.GET("/products/:id", getProduct)
	}

	return router
}

func healthCheck(c *gin.Context) {
	// Проверка подключения к БД
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

	query := database.GetDB().Preload("Category")

	// Фильтр по категории
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"count": len(products),
	})
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.GetDB().Preload("Category").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
