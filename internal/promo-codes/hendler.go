package promo_codes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PromoCodeHandler struct {
	promoService Service
}

func NewPromoCodeHandler(promoService Service) *PromoCodeHandler {
	return &PromoCodeHandler{promoService: promoService}
}

func (h *PromoCodeHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	promos := api.Group("/promo-codes")
	{
		promos.POST("/", h.CreatePromoCode)
		promos.GET("/:id", h.GetPromoCode)
		promos.PUT("/:id", h.UpdatePromoCode)
		promos.DELETE("/:id", h.DeletePromoCode)
		promos.GET("/", h.GetAllPromoCodes)
		promos.GET("/active", h.GetActivePromoCodes)
		promos.GET("/type/:type", h.GetPromoCodesByType)
		promos.POST("/validate", h.ValidatePromoCode)
		promos.POST("/apply", h.ApplyPromoCode)
	}

}

func (h *PromoCodeHandler) CreatePromoCode(c *gin.Context) {
	var req CreatePromoCodeRequest
	errJson := c.ShouldBindJSON(&req)
	if errJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJson.Error()})
		return
	}

	createPromoCode, errCreate := h.promoService.CreatePromoCode(&req)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": createPromoCode})
}

func (h *PromoCodeHandler) GetPromoCode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid promo-code ID format",
			"message": "Promo-code ID must be a number",
		})
		return
	}
	promoResponse, err := h.promoService.GetPromoCode(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promo code not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": promoResponse})
}

func (h *PromoCodeHandler) UpdatePromoCode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid promo-code ID format",
			"message": "Promo-code ID must be a number",
		})
		return
	}
	var req UpdatePromoCodeRequest
	errorJSON := c.ShouldBindJSON(&req)
	if errorJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorJSON.Error()})
		return
	}
	promoUpdate, errUpdate := h.promoService.UpdatePromoCode(uint(id), &req)
	if errUpdate != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errUpdate.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": promoUpdate})
}

func (h *PromoCodeHandler) DeletePromoCode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid promo-code ID format",
			"message": "Promo-code ID must be a number",
		})
		return
	}
	err = h.promoService.DeletePromoCode(uint(id))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promo code not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Promo code deleted successfully"})
}

func (h *PromoCodeHandler) GetAllPromoCodes(c *gin.Context) {
	result, err := h.promoService.GetAllPromoCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *PromoCodeHandler) GetActivePromoCodes(c *gin.Context) {
	result, err := h.promoService.GetActivePromoCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *PromoCodeHandler) GetPromoCodesByType(c *gin.Context) {
	promoType := c.Param("type")

	result, err := h.promoService.GetPromoCodesByType(promoType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (h *PromoCodeHandler) ValidatePromoCode(c *gin.Context) {
	var req ValidatePromoCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.promoService.ValidatePromoCode(req.Code, req.OrderAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})

}

func (h *PromoCodeHandler) ApplyPromoCode(c *gin.Context) {
	var req ApplyPromoCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.promoService.ApplyPromoCode(req.Code, req.OrderAmount, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
