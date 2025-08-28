package restaurant

import (
	"delicious-and-kidney/internal/user"
	"time"
)

type Restaurant struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	OwnerID         uint      `gorm:"not null;index" json:"owner_id"`
	Name            string    `gorm:"size:255;not null" json:"name"`
	Description     string    `gorm:"type:text" json:"description"`
	Phone           string    `gorm:"size:20" json:"phone"`
	Email           string    `gorm:"size:255" json:"email"`
	Address         string    `gorm:"size:500;not null" json:"address"`
	Latitude        *float64  `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude       *float64  `gorm:"type:decimal(11,8)" json:"longitude"`
	ImageURL        string    `gorm:"size:500" json:"image_url"`
	MinOrderAmount  float64   `gorm:"type:decimal(10,2);default:0" json:"min_order_amount"`
	DeliveryFee     float64   `gorm:"type:decimal(10,2);default:0" json:"delivery_fee"`
	DeliveryTimeMin int       `gorm:"default:30" json:"delivery_time_min"`
	DeliveryTimeMax int       `gorm:"default:60" json:"delivery_time_max"`
	Rating          float64   `gorm:"type:decimal(3,2);default:0" json:"rating"`
	ReviewsCount    int       `gorm:"default:0" json:"reviews_count"`
	IsActive        bool      `gorm:"default:true" json:"is_active"`
	IsFeatured      bool      `gorm:"default:false" json:"is_featured"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Owner           user.User `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
}

type CreateRestaurantRequest struct {
	Name            string   `json:"name" binding:"required"`
	Description     string   `json:"description"`
	Email           string   `json:"email" validate:"omitempty,email"`
	Phone           string   `json:"phone" validate:"omitempty,phone"`
	Address         string   `json:"address" binding:"required"`
	Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	ImageURL        string   `json:"image_url"`
	MinOrderAmount  float64  `json:"min_order_amount"`
	DeliveryFee     float64  `json:"delivery_fee"`
	DeliveryTimeMin int      `json:"delivery_time_min"`
	DeliveryTimeMax int      `json:"delivery_time_max"`
	IsActive        *bool    `json:"is_active"`
	IsFeatured      *bool    `json:"is_featured"`
}

type UpdateRestaurantRequest struct {
	Name            *string  `json:"name,omitempty"`
	Description     *string  `json:"description,omitempty"`
	Phone           *string  `json:"phone,omitempty"`
	Email           *string  `json:"email,omitempty"`
	Address         *string  `json:"address,omitempty"`
	Latitude        *float64 `json:"latitude,omitempty"`
	Longitude       *float64 `json:"longitude,omitempty"`
	ImageURL        *string  `json:"image_url,omitempty"`
	MinOrderAmount  *float64 `json:"min_order_amount,omitempty"`
	DeliveryFee     *float64 `json:"delivery_fee,omitempty"`
	DeliveryTimeMin *int     `json:"delivery_time_min,omitempty"`
	DeliveryTimeMax *int     `json:"delivery_time_max,omitempty"`
	IsActive        *bool    `json:"is_active,omitempty"`
	IsFeatured      *bool    `json:"is_featured,omitempty"`
}

type RestaurantResponse struct {
	ID              uint      `json:"id"`
	OwnerID         uint      `json:"owner_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	Latitude        *float64  `json:"latitude"`
	Longitude       *float64  `json:"longitude"`
	ImageURL        string    `json:"image_url"`
	MinOrderAmount  float64   `json:"min_order_amount"`
	DeliveryFee     float64   `json:"delivery_fee"`
	DeliveryTimeMin int       `json:"delivery_time_min"`
	DeliveryTimeMax int       `json:"delivery_time_max"`
	Rating          float64   `json:"rating"`
	ReviewsCount    int       `json:"reviews_count"`
	IsActive        bool      `json:"is_active"`
	IsFeatured      bool      `json:"is_featured"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RestaurantStats struct {
	TotalOrders   int     `json:"total_orders"`   // всего заказов
	TotalRevenue  float64 `json:"total_revenue"`  // общая выручка
	AverageRating float64 `json:"average_rating"` // средний рейтинг
	ReviewsCount  int     `json:"reviews_count"`  // количество отзывов
}
