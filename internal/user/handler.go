package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	users := api.Group("/user")
	{
		users.GET("/:id", h.GetUser)
		users.PATCH("/:id", h.UpdateUser)
		users.POST("/", h.CreateUser)
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"message": "User ID must be a number",
		})
		return
	}
	userResponse, err := h.userService.GetProfile(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"message": "User ID must be a number",
		})
		return
	}
	var req UpdateUserRequest
	errJSON := c.ShouldBindJSON(&req)
	if errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJSON.Error()})
		return
	}
	userUpdate, errUpdate := h.userService.UpdateProfile(uint(id), &req)
	if errUpdate != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": userUpdate})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req User
	errJSON := c.ShouldBindJSON(&req)
	if errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJSON.Error()})
		return
	}
	createdUser, errCreate := h.userService.CreateUser(&req)
	if errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with created user"})
	}
	c.JSON(http.StatusOK, gin.H{"data": createdUser})
}
