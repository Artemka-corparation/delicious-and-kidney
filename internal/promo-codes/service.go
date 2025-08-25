package promo_codes

import (
	"delicious-and-kidney/pkg/Errors"
	"errors"
	"gorm.io/gorm"
	"time"
)

type PromoCodesService struct {
	promoCodeRepository Repository
}

func NewPromoCodesService(promoCodeRepository Repository) *PromoCodesService {
	return &PromoCodesService{promoCodeRepository: promoCodeRepository}
}

func (s *PromoCodesService) CreatePromoCode(req *CreatePromoCodeRequest) (*PromoCodeResponse, error) {
	_, err := s.promoCodeRepository.FindByCode(req.Code)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrPromoCodeNotFound
		}
	}
	promo := ToPromoCodes(req)
	savePromo, err := s.promoCodeRepository.Create(promo)
	if err != nil {
		return nil, err
	}
	response := ToPromoCodesResponse(savePromo)
	return response, nil
}

func (s *PromoCodesService) GetPromoCode(id uint) (*PromoCodeResponse, error) {
	promo, err := s.promoCodeRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrPromoCodeNotFound
		}
		return nil, err
	}
	response := ToPromoCodeResponse(promo)
	return response, nil
}

func (s *PromoCodesService) UpdatePromoCode(id uint, req *UpdatePromoCodeRequest) (*PromoCodeResponse, error) {
	_, err := s.promoCodeRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrPromoCodeNotFound
		}
		return nil, err
	}
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = req.Name
	}
	if req.Description != nil {
		updates["description"] = req.Description
	}
	if req.Value != nil {
		updates["value"] = req.Value
	}
	if req.MinOrderAmount != nil {
		updates["min_order_amount"] = req.MinOrderAmount
	}
	if req.MaxDiscount != nil {
		updates["max_discount"] = req.MaxDiscount
	}
	if req.UsageLimit != nil {
		updates["usage_limit"] = req.UsageLimit
	}
	if req.UsageLimitPerUser != nil {
		updates["usage_limit_per_user"] = req.UsageLimitPerUser
	}
	if req.ValidUntil != nil {
		updates["valid_until"] = req.ValidUntil
	}
	if req.IsActive != nil {
		updates["is_active"] = req.IsActive
	}
	err = s.promoCodeRepository.UpdateFields(id, updates)
	if err != nil {
		return nil, err
	}
	updatePromo, err := s.promoCodeRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return ToPromoCodesResponse(updatePromo), nil
}

func (s *PromoCodesService) DeletePromoCode(id uint) error {
	_, err := s.promoCodeRepository.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrPromoCodeNotFound
		}
		return err
	}
	err = s.promoCodeRepository.Delete(id)
	return err
}

func (s *PromoCodesService) ValidatePromoCode(code string, orderAmount float64) (*ValidationResult, error) {
	promo, err := s.promoCodeRepository.FindByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ValidationResult{
				IsValid:      false,
				ErrorMessage: "Promo code not found",
			}, nil
		}
		return nil, err
	}
	if !promo.IsActive {
		return &ValidationResult{
			IsValid:      false,
			ErrorMessage: "Promo code is inactive",
		}, nil
	}
	if promo.ValidFrom.After(time.Now()) {
		return &ValidationResult{
			IsValid:      false,
			ErrorMessage: "Promo code not yet valid",
		}, nil
	}
	if promo.ValidUntil != nil && promo.ValidUntil.Before(time.Now()) {
		return &ValidationResult{
			IsValid:      false,
			ErrorMessage: "Promo code expired",
		}, nil
	}
	if orderAmount < promo.MinOrderAmount {
		return &ValidationResult{IsValid: false, ErrorMessage: "Order amount too low"}, nil
	}
	if promo.UsageLimit != nil && promo.UsageCount >= *promo.UsageLimit {
		return &ValidationResult{IsValid: false, ErrorMessage: "Usage limit exceeded"}, nil
	}
	return &ValidationResult{
		IsValid:       true,
		DiscountType:  promo.Type,
		DiscountValue: promo.Value,
	}, nil
}

func (s *PromoCodesService) ApplyPromoCode(code string, orderAmount float64, userID uint) (*AppliedPromoResult, error) {
	var discountAmount float64
	validationResult, err := s.ValidatePromoCode(code, orderAmount)
	if err != nil {
		return nil, err
	}
	if !validationResult.IsValid {
		return &AppliedPromoResult{
			OriginalAmount: orderAmount,
			DiscountAmount: 0,
			FinalAmount:    orderAmount,
			PromoCode:      code,
			DiscountType:   "",
		}, nil
	}
	promo, err := s.promoCodeRepository.FindByCode(code)
	if err != nil {
		return nil, err
	}
	if promo.Type == "percentage" {
		discountAmount = orderAmount * (promo.Value / 100)

		if promo.MaxDiscount != nil && discountAmount > *promo.MaxDiscount {
			discountAmount = *promo.MaxDiscount
		}
	} else if promo.Type == "fixed" {
		discountAmount = promo.Value
		if discountAmount > orderAmount {
			discountAmount = orderAmount
		}
	} else if promo.Type == "free_delivery" {
		discountAmount = promo.Value
	}
	finalAmount := orderAmount - discountAmount
	if finalAmount < 0 {
		finalAmount = 0
	}
	err = s.promoCodeRepository.IncrementUsageCount(promo.ID)
	if err != nil {
		return nil, err
	}
	return &AppliedPromoResult{
		OriginalAmount: orderAmount,
		DiscountAmount: discountAmount,
		FinalAmount:    finalAmount,
		PromoCode:      code,
		DiscountType:   promo.Type,
	}, nil
}

func (s *PromoCodesService) GetAllPromoCodes() ([]PromoCodeResponse, error) {
	promoCodes, err := s.promoCodeRepository.FindAll()
	if err != nil {
		return nil, err
	}
	responses := make([]PromoCodeResponse, len(promoCodes))
	for i, promo := range promoCodes {
		responses[i] = *ToPromoCodeResponse(&promo)
	}

	return responses, nil
}

func (s *PromoCodesService) GetActivePromoCodes() ([]PromoCodeResponse, error) {
	promoCodes, err := s.promoCodeRepository.FindActivePromoCodes()
	if err != nil {
		return nil, err
	}
	response := make([]PromoCodeResponse, len(promoCodes))
	for i, promo := range promoCodes {
		response[i] = *ToPromoCodeResponse(&promo)
	}
	return response, nil
}

func (s *PromoCodesService) GetPromoCodesByType(promoType string) ([]PromoCodeResponse, error) {
	promoCodes, err := s.promoCodeRepository.FindByType(promoType)
	if err != nil {
		return nil, err
	}
	responses := make([]PromoCodeResponse, len(promoCodes))
	for i, promo := range promoCodes {
		responses[i] = *ToPromoCodeResponse(&promo)
	}
	return responses, nil
}
