package models

import (
	"gorm.io/gorm"
	"time"
)

// =====================
// ОСНОВНЫЕ МОДЕЛИ
// =====================

type User struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Phone        *string   `json:"phone"`
	Role         string    `json:"role" gorm:"default:'client'"`
	PasswordHash string    `json:"-" gorm:"column:password_hash;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`

	// Связи
	Orders    []Order    `json:"orders,omitempty" gorm:"foreignKey:UserID"`
	CartItems []CartItem `json:"cart_items,omitempty" gorm:"foreignKey:UserID"`
	Addresses []Address  `json:"addresses,omitempty" gorm:"foreignKey:UserID"`
}

type Category struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null"`

	// Связи
	Products []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"not null"`
	Description *string `json:"description"`
	Price       float64 `json:"price" gorm:"type:numeric(10,2);not null"`
	ImageURL    *string `json:"image_url" gorm:"column:image_url"`
	CategoryID  *int    `json:"category_id" gorm:"column:category_id"`

	// Связи
	Category *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}

type Order struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          int       `json:"user_id" gorm:"column:user_id;not null"`
	Status          string    `json:"status" gorm:"not null;default:'pending'"`
	TotalPrice      float64   `json:"total_price" gorm:"column:total_price;type:numeric(10,2);not null"`
	DeliveryType    string    `json:"delivery_type" gorm:"column:delivery_type;default:'delivery'"`
	PaymentMethod   string    `json:"payment_method" gorm:"column:payment_method;default:'cash'"`
	DeliveryAddress *string   `json:"delivery_address" gorm:"column:delivery_address"`
	Comment         *string   `json:"comment"`
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`

	// Связи
	User       User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   int     `json:"order_id" gorm:"column:order_id;not null"`
	ProductID int     `json:"product_id" gorm:"column:product_id;not null"`
	Quantity  int     `json:"quantity" gorm:"not null;default:1"`
	Price     float64 `json:"price" gorm:"type:numeric(10,2);not null"`

	// Связи
	Order   Order   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

type CartItem struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"column:user_id;not null"`
	ProductID int       `json:"product_id" gorm:"column:product_id;not null"`
	Quantity  int       `json:"quantity" gorm:"not null;default:1"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`

	// Связи
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

type Address struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int     `json:"user_id" gorm:"column:user_id;not null"`
	Title     *string `json:"title"`
	Street    string  `json:"street" gorm:"not null"`
	House     string  `json:"house" gorm:"not null"`
	Apartment *string `json:"apartment"`
	Comment   *string `json:"comment"`
	IsDefault bool    `json:"is_default" gorm:"column:is_default;default:false"`

	// Связи
	User User `json:"user" gorm:"foreignKey:UserID"`
}

// =====================
// КОНСТАНТЫ
// =====================

// Роли пользователей
const (
	RoleClient  = "client"
	RoleAdmin   = "admin"
	RoleManager = "manager"
)

// Статусы заказов
const (
	OrderStatusPending    = "pending"
	OrderStatusConfirmed  = "confirmed"
	OrderStatusPreparing  = "preparing"
	OrderStatusReady      = "ready"
	OrderStatusDelivering = "delivering"
	OrderStatusDelivered  = "delivered"
	OrderStatusCancelled  = "cancelled"
)

// Типы доставки
const (
	DeliveryTypeDelivery = "delivery"
	DeliveryTypePickup   = "pickup"
)

// Способы оплаты
const (
	PaymentMethodCash = "cash"
	PaymentMethodCard = "card"
)

// =====================
// МЕТОДЫ
// =====================

// Методы для Order
func (o *Order) CalculateTotal() {
	var total float64
	for _, item := range o.OrderItems {
		total += item.Price * float64(item.Quantity)
	}
	o.TotalPrice = total
}

func (o *Order) IsEditable() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusConfirmed
}

func (o *Order) CanBeCancelled() bool {
	return o.Status != OrderStatusDelivered && o.Status != OrderStatusCancelled
}

// Методы для OrderItem
func (oi *OrderItem) GetTotal() float64 {
	return oi.Price * float64(oi.Quantity)
}

// Методы для CartItem
func (ci *CartItem) GetTotal(productPrice float64) float64 {
	return productPrice * float64(ci.Quantity)
}

// Методы для Address
func (a *Address) GetFullAddress() string {
	address := a.Street + ", " + a.House
	if a.Apartment != nil && *a.Apartment != "" {
		address += ", кв. " + *a.Apartment
	}
	return address
}

// =====================
// УТИЛИТЫ ДЛЯ БД
// =====================

// AutoMigrate - НЕ ИСПОЛЬЗУЕМ, так как таблицы созданы через SQL
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Category{},
		&Product{},
		&Order{},
		&OrderItem{},
		&CartItem{},
		&Address{},
	)
}

// CheckConnection - проверка подключения к БД
func CheckConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}
