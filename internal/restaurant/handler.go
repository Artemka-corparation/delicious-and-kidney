package restaurant

import (
	"delicious-and-kidney/internal/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RestaurantHandler struct {
	restaurantService service
}

func NewRestaurantHandler(restaurantService service) *RestaurantHandler {
	return &RestaurantHandler{restaurantService: restaurantService}
}

func (h *RestaurantHandler) RegisterRoutes(router *gin.Engine, authService *auth.AuthService) {
	api := router.Group("/api")
	restaurants := api.Group("/restaurants")
	{
		restaurants.GET("/:id", h.GetRestaurant)
		restaurants.GET("/my", h.GetMyRestaurants)
		restaurants.GET("/", auth.AuthMiddleware(authService), h.GetAllRestaurants)
		restaurants.GET("/search", h.SearchRestaurants)
		restaurants.POST("/", auth.AuthMiddleware(authService), h.CreateRestaurant)
		restaurants.PATCH("/:id", auth.AuthMiddleware(authService), h.UpdateRestaurant)
		restaurants.DELETE("/:id", auth.AuthMiddleware(authService), h.DeleteRestaurant)
		restaurants.PATCH("/:id/activate", auth.AuthMiddleware(authService), h.ActivateRestaurant)
		restaurants.PATCH("/:id/deactivate", auth.AuthMiddleware(authService), h.DeactivateRestaurant)
		restaurants.PATCH("/:id/featured", auth.AuthMiddleware(authService), h.SetFeaturedStatus)
	}
}

func (h *RestaurantHandler) CreateRestaurant(c *gin.Context) {
	var req CreateRestaurantRequest
	errJSON := c.ShouldBindJSON(&req)
	if errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJSON.Error()})
		return
	}
	ownerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	restaurant, err := h.restaurantService.CreateRestaurant(ownerID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}

func (h *RestaurantHandler) GetRestaurant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid restaurant ID format",
			"message": "Restaurant ID must be a number",
		})
		return
	}
	restaurantResponse, err := h.restaurantService.GetRestaurant(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": restaurantResponse})
}

func (h *RestaurantHandler) UpdateRestaurant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Restaurant ID format",
			"message": "Restaurant ID must be a number",
		})
		return
	}
	var req UpdateRestaurantRequest
	errJson := c.ShouldBindJSON(&req)
	if errJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJson.Error()})
		return
	}
	ownerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	restaurant, err := h.restaurantService.UpdateRestaurant(uint(id), ownerID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

func (h *RestaurantHandler) DeleteRestaurant(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Restaurant ID format",
			"message": "Restaurant ID must be a number",
		})
		return
	}

	ownerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err = h.restaurantService.DeleteRestaurant(uint(id), ownerID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Operation completed successfully"})

}

func (h *RestaurantHandler) ActivateRestaurant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Restaurant ID format",
			"message": "Restaurant ID must be a number",
		})
		return
	}
	ownerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err = h.restaurantService.ActivateRestaurant(uint(id), ownerID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": err})
}

func (h *RestaurantHandler) DeactivateRestaurant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Restaurant ID format",
			"message": "Restaurant ID must be a number",
		})
		return
	}
	ownerID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err = h.restaurantService.DeactivateRestaurant(uint(id), ownerID.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": err})
}

func (h *RestaurantHandler) SetFeaturedStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid Restaurant ID format",
			"message": "Restaurant ID must be a number",
		})
		return
	}
	var req struct {
		Featured bool `json:"featured" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRole, exists := c.Get("role")
	if !exists || userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
		return
	}
	err = h.restaurantService.SetFeaturedStatus(uint(id), req.Featured, userRole.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": err})
}

func (h *RestaurantHandler) GetMyRestaurants(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	restaurants, err := h.restaurantService.GetMyRestaurants(userId.(uint), limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

func (h *RestaurantHandler) GetAllRestaurants(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	restaurants, err := h.restaurantService.GetAllRestaurants(limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

func (h *RestaurantHandler) SearchRestaurants(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query required"})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	restaurants, err := h.restaurantService.SearchRestaurants(query, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}
