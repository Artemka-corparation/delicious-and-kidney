package user

import "time"

type User struct {
	Id            uint      `gorm:"primarykey"`
	Name          string    `gorm:"not null" json:"name"`
	Email         string    `gorm:"unique" validate:"required,email"`
	Phone         string    `gorm:"size:20" json:"phone"`
	PasswordHash  string    `gorm:"not null" json:"-"`
	Role          string    `gorm:"default:customer;check:role IN ('customer','restaurant_owner','courier','admin')"`
	EmailVerified bool      `gorm:"default:false"`
	PhoneVerified bool      `gorm:"default:false"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
