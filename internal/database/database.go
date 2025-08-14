package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"food-delivery-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	// Добавляем найденный IP адрес контейнера
	hosts := []string{
		"127.0.0.1",
		"localhost",
	}
	port := getEnv("DB_PORT", "5433")
	user := getEnv("DB_USER", "app_user")
	password := getEnv("DB_PASSWORD", "secret")
	dbname := getEnv("DB_NAME", "app_db")
	sslmode := getEnv("DB_SSLMODE", "disable")

	var database *gorm.DB
	var err error

	for _, host := range hosts {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode,
		)

		log.Printf("🔍 Пробуем подключиться к: %s:%s", host, port)

		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err == nil {
			sqlDB, testErr := database.DB()
			if testErr == nil {
				if pingErr := sqlDB.Ping(); pingErr == nil {
					log.Printf("✅ Успешное подключение к PostgreSQL через %s:%s!", host, port)
					break
				}
			}
		}

		log.Printf("❌ Не удалось подключиться через %s: %v", host, err)
	}

	if err != nil {
		log.Fatal("❌ Не удалось подключиться ни к одному адресу PostgreSQL")
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Не удалось получить объект sql.DB:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = database
	log.Println("🎉 PostgreSQL готов к работе!")

	// Добавляем тестовые данные
	seedData()
}

func seedData() {
	log.Println("📝 Проверяем тестовые данные...")

	// Проверяем, есть ли уже данные
	var categoryCount int64
	DB.Model(&models.Category{}).Count(&categoryCount)

	if categoryCount > 0 {
		log.Println("✅ Тестовые данные уже существуют")
		return
	}

	log.Println("➕ Добавляем тестовые данные...")

	// Категории
	categories := []models.Category{
		{Name: "Бургеры"},
		{Name: "Напитки"},
		{Name: "Десерты"},
		{Name: "Закуски"},
	}

	for _, cat := range categories {
		result := DB.Create(&cat)
		if result.Error != nil {
			log.Printf("Ошибка создания категории %s: %v", cat.Name, result.Error)
		}
	}

	// Товары
	products := []models.Product{
		{Name: "Биг Мак", Description: stringPtr("Классический бургер с двумя котлетами"), Price: 299.00, CategoryID: intPtr(1)},
		{Name: "Чизбургер", Description: stringPtr("Бургер с сыром и котлетой"), Price: 149.00, CategoryID: intPtr(1)},
		{Name: "Роял Чизбургер", Description: stringPtr("Большой бургер с беконом"), Price: 259.00, CategoryID: intPtr(1)},
		{Name: "Кока-Кола", Description: stringPtr("Освежающий напиток 0.5л"), Price: 89.00, CategoryID: intPtr(2)},
		{Name: "Молочный коктейль", Description: stringPtr("Ванильный молочный коктейль"), Price: 159.00, CategoryID: intPtr(2)},
		{Name: "Апельсиновый сок", Description: stringPtr("Свежевыжатый сок 0.3л"), Price: 129.00, CategoryID: intPtr(2)},
		{Name: "Яблочный пирог", Description: stringPtr("Теплый яблочный пирог"), Price: 99.00, CategoryID: intPtr(3)},
		{Name: "Мороженое", Description: stringPtr("Ванильное мороженое"), Price: 79.00, CategoryID: intPtr(3)},
		{Name: "Картофель фри", Description: stringPtr("Хрустящий картофель фри"), Price: 129.00, CategoryID: intPtr(4)},
		{Name: "Куриные наггетсы", Description: stringPtr("6 кусочков наггетсов"), Price: 189.00, CategoryID: intPtr(4)},
	}

	for _, prod := range products {
		result := DB.Create(&prod)
		if result.Error != nil {
			log.Printf("Ошибка создания товара %s: %v", prod.Name, result.Error)
		}
	}

	log.Println("✅ Тестовые данные успешно добавлены!")
}

// Вспомогательные функции для указателей
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Ошибка при получении sql.DB: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Ошибка при закрытии БД: %v", err)
	} else {
		log.Println("🔒 Подключение к PostgreSQL закрыто")
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
