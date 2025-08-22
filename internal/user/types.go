package user

import "time"

type User struct {
	Id            uint      `json:"id" gorm:"primary"`
	Name          string    `json:"name" gorm:"not null"`
	Email         string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Phone         string    `json:"phone" gorm:"size:20" `
	PasswordHash  string    `json:"passwordHash" gorm:"not null"`
	Role          string    `json:"role" gorm:"default:customer;check:role IN ('customer','restaurant_owner','courier','admin')"`
	EmailVerified bool      `json:"emailVerified" gorm:"default:false"`
	PhoneVerified bool      `json:"phoneVerified" gorm:"default:false"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type UserDto struct {
	Name          string    `json:"name" validate:"required"`
	Email         string    `json:"email" validate:"required,email"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
	EmailVerified bool      `json:"email_verified"`
	PhoneVerified bool      `json:"phone_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Role          string    `json:"role"`
	EmailVerified bool      `json:"email_verified"`
	PhoneVerified bool      `json:"phone_verified"`
	CreatedAt     time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Phone *string `json:"phone,omitempty"`
}
