package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"unique" validate:"required,email"`
	Phone        string `gorm:"not null" validate:"required,min=10"`
	PasswordHash string `gorm:"not null" json:"-"`
	Role         string `gorm:"default:customer;check:role IN ('customer','restaurant_owner','courier','admin')"`
}
