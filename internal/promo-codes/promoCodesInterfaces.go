package promo_codes

type Repository interface {
	Create(promo *PromoCodes) (*PromoCodes, error)
	FindById(id uint) (*PromoCodes, error)
	FindByCode(code string) (*PromoCodes, error)
	Update(promo *PromoCodes) (*PromoCodes, error)
	UpdateFields(id uint, updates map[string]interface{}) error
	Delete(id uint) error
	FindActivePromoCodes() ([]PromoCodes, error)
	FindValidPromoCodes() ([]PromoCodes, error)
	IncrementUsageCount(id uint) error
	FindByType(promoType string) ([]PromoCodes, error)
	Count() (int64, error)
	FindAll() ([]PromoCodes, error)
}

type Service interface {
	// Управление промокодами
	CreatePromoCode(req *CreatePromoCodeRequest) (*PromoCodeResponse, error)
	GetPromoCode(id uint) (*PromoCodeResponse, error)
	UpdatePromoCode(id uint, req *UpdatePromoCodeRequest) (*PromoCodeResponse, error)
	DeletePromoCode(id uint) error

	// Работа с промокодами для пользователей
	ValidatePromoCode(code string, orderAmount float64) (*ValidationResult, error)
	ApplyPromoCode(code string, orderAmount float64, userID uint) (*AppliedPromoResult, error)

	// Поиск и фильтрация
	GetAllPromoCodes() ([]PromoCodeResponse, error)
	GetActivePromoCodes() ([]PromoCodeResponse, error)
	GetPromoCodesByType(promoType string) ([]PromoCodeResponse, error)
}
