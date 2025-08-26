package user

import (
	"delicious-and-kidney/pkg/Errors"
	"errors"
	"net/http"
	"strconv"
	"strings"

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
		users.POST("/:id/password", h.ChangePasswordHandler)
		users.PATCH("/:id/deactivate", h.DeactivateAccountHandler)
		users.DELETE("/:id", h.DeleteAccountHandler)
		users.PATCH("/:id/activate", h.ActivateAccountHandler)
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
	var req CreateUserRequest
	errJSON := c.ShouldBindJSON(&req)
	if errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJSON.Error()})
		return
	}
	createdUser, errCreate := h.userService.CreateUser(&req)
	if errCreate != nil {
		if strings.Contains(errCreate.Error(), "email already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": errCreate.Error()})
		} else if strings.Contains(errCreate.Error(), "password") {
			c.JSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

func (h *UserHandler) ChangePasswordHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"message": "User ID must be a number",
		})
		return
	}
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.userService.ChangePassword(uint(id), req.OldPassword, req.NewPassword)
	if err != nil {
		if errors.Is(err, Errors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else if errors.Is(err, Errors.ErrWrongPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Current password is incorrect"})
		} else if errors.Is(err, Errors.ErrWeakPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "New password is too weak"})
		} else if errors.Is(err, Errors.ErrSamePassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "New password must be different"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Password changed successfully"})
}

func (h *UserHandler) DeactivateAccountHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"message": "User ID must be a number",
		})
		return
	}
	err = h.userService.DeactivateAccount(uint(id))
	if errors.Is(err, Errors.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User deactivated"})
}

func (h *UserHandler) DeleteAccountHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"message": "User ID must be a number",
		})
		return
	}
	err = h.userService.DeleteAccount(uint(id))
	if errors.Is(err, Errors.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User deleted"})
}

func (h *UserHandler) ActivateAccountHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"message": "User ID must be a number",
		})
		return
	}
	err = h.userService.ActivateAccount(uint(id))
	if errors.Is(err, Errors.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User activated"})
}
