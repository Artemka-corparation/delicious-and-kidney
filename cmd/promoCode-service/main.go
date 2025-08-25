package main

import (
	"delicious-and-kidney/configs"
	"delicious-and-kidney/internal/promo-codes"
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
	err = database.AutoMigrate(&promo_codes.PromoCodes{})
	if err != nil {
		log.Fatal("Not to auto migrate", err)
	}
	promoRepo := promo_codes.NewPromoCodeRepository(database)
	promoService := promo_codes.NewPromoCodesService(promoRepo)
	promoHandler := promo_codes.NewPromoCodeHandler(promoService)
	router := gin.Default()
	promoHandler.RegisterRoutes(router)
	router.Run(":8082")
}
