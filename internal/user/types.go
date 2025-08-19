package user

import "time"

type User struct {
	Id            uint      `gorm:"primarykey"`
	Name          string    `gorm:"not null" json:"name"`
	Email         string    `gorm:"unique;not null" json:"email" validate:"required,email"`
	Phone         string    `gorm:"size:20" json:"phone"`
	PasswordHash  string    `gorm:"not null" json:"-"`
	Role          string    `gorm:"default:customer;check:role IN ('customer','restaurant_owner','courier','admin')"`
	EmailVerified bool      `gorm:"default:false"`
	PhoneVerified bool      `gorm:"default:false"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
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
