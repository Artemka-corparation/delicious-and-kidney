package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	users := api.Group("user")
	users.GET("/login", h.Login)
}

func (h *AuthHandler) Login(c *gin.Context) {
	fmt.Println("Login")
}
