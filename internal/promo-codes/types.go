package promo_codes

import "time"

type PromoCodes struct {
	ID                uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code              string     `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name              string     `gorm:"type:varchar(255)" json:"name"`
	Description       string     `gorm:"type:text" json:"description"`
	Type              string     `gorm:"type:varchar(50);default:percentage;check:type IN ('percentage','fixed','free_delivery')" json:"type"`
	Value             float64    `gorm:"type:decimal(10,2);not null" json:"value"`
	MinOrderAmount    float64    `gorm:"type:decimal(10,2);default:0" json:"min_order_amount"`
	MaxDiscount       *float64   `gorm:"type:decimal(10,2)" json:"max_discount,omitempty"`
	UsageLimit        *int       `gorm:"type:int" json:"usage_limit,omitempty"`
	UsageLimitPerUser int        `gorm:"type:int;default:1" json:"usage_limit_per_user"`
	UsageCount        int        `gorm:"type:int;default:0" json:"usage_count"`
	ValidFrom         time.Time  `gorm:"type:timestamp;default:now()" json:"valid_from"`
	ValidUntil        *time.Time `gorm:"type:timestamp" json:"valid_until,omitempty"`
	IsActive          bool       `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time  `gorm:"type:timestamp;default:now()" json:"created_at"`
}

type CreatePromoCodeRequest struct {
	Code              string     `json:"code" binding:"required"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	Type              string     `json:"type" binding:"required"`
	Value             float64    `json:"value" binding:"required"`
	MinOrderAmount    float64    `json:"min_order_amount"`
	MaxDiscount       *float64   `json:"max_discount,omitempty"`
	UsageLimit        *int       `json:"usage_limit,omitempty"`
	UsageLimitPerUser int        `json:"usage_limit_per_user"`
	ValidFrom         time.Time  `json:"valid_from"`
	ValidUntil        *time.Time `json:"valid_until,omitempty"`
	IsActive          bool       `json:"is_active"`
}

type UpdatePromoCodeRequest struct {
	Name              *string    `json:"name,omitempty"`
	Description       *string    `json:"description,omitempty"`
	Value             *float64   `json:"value,omitempty"`
	MinOrderAmount    *float64   `json:"min_order_amount,omitempty"`
	MaxDiscount       *float64   `json:"max_discount,omitempty"`
	UsageLimit        *int       `json:"usage_limit,omitempty"`
	UsageLimitPerUser *int       `json:"usage_limit_per_user,omitempty"`
	ValidUntil        *time.Time `json:"valid_until,omitempty"`
	IsActive          *bool      `json:"is_active,omitempty"`
}

type PromoCodeResponse struct {
	ID                uint       `json:"id"`
	Code              string     `json:"code"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	Type              string     `json:"type"`
	Value             float64    `json:"value"`
	MinOrderAmount    float64    `json:"min_order_amount"`
	MaxDiscount       *float64   `json:"max_discount,omitempty"`
	UsageLimit        *int       `json:"usage_limit,omitempty"`
	UsageLimitPerUser int        `json:"usage_limit_per_user"`
	UsageCount        int        `json:"usage_count"`
	ValidFrom         time.Time  `json:"valid_from"`
	ValidUntil        *time.Time `json:"valid_until,omitempty"`
	IsActive          bool       `json:"is_active"`
	CreatedAt         time.Time  `json:"created_at"`
}

type ValidationResult struct {
	IsValid       bool    `json:"is_valid"`
	ErrorMessage  string  `json:"error_message,omitempty"`
	DiscountType  string  `json:"discount_type,omitempty"`
	DiscountValue float64 `json:"discount_value,omitempty"`
}

type AppliedPromoResult struct {
	OriginalAmount float64 `json:"original_amount"`
	DiscountAmount float64 `json:"discount_amount"`
	FinalAmount    float64 `json:"final_amount"`
	PromoCode      string  `json:"promo_code"`
	DiscountType   string  `json:"discount_type"`
}

type ValidatePromoCodeRequest struct {
	Code        string  `json:"code" binding:"required"`
	OrderAmount float64 `json:"order_amount" binding:"required,gt=0"`
}

type ApplyPromoCodeRequest struct {
	Code        string  `json:"code" binding:"required"`
	OrderAmount float64 `json:"order_amount" binding:"required,gt=0"`
	UserID      uint    `json:"user_id" binding:"required"` // Пока без JWT
}
