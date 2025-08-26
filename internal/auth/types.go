package auth

import "time"

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2, max=100"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required, min=6"`
	Role     string `json:"role" binding:"omitepmty,oneof=customer restaurant_owner courier"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password", binding:"required"`
}

type AuthResponse struct {
	User  UserInfo `json:"user"`
	Token string   `json:"token"`
}

type UserInfo struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
	EmailVerified bool      `json:"email_verified"`
	PhoneVerified bool      `json:"phone_verified"`
	CreatedAt     time.Time `json:"created_at"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binging:"required"`
	NewPassword     string `json:"new_password" binging:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
