package promo_codes

func ToPromoCodesResponse(promo *PromoCodes) *PromoCodeResponse {
	return &PromoCodeResponse{
		ID:                promo.ID,
		Code:              promo.Code,
		Name:              promo.Name,
		Description:       promo.Description,
		Type:              promo.Type,
		Value:             promo.Value,
		MinOrderAmount:    promo.MinOrderAmount,
		MaxDiscount:       promo.MaxDiscount,
		UsageLimit:        promo.UsageLimit,
		UsageLimitPerUser: promo.UsageLimitPerUser,
		ValidFrom:         promo.ValidFrom,
		ValidUntil:        promo.ValidUntil,
		IsActive:          promo.IsActive,
		UsageCount:        0,
	}
}

func ToPromoCodeResponse(promo *PromoCodes) *PromoCodeResponse {
	return &PromoCodeResponse{
		ID:                promo.ID,
		Code:              promo.Code,
		Name:              promo.Name,
		Description:       promo.Description,
		Type:              promo.Type,
		Value:             promo.Value,
		MinOrderAmount:    promo.MinOrderAmount,
		MaxDiscount:       promo.MaxDiscount,
		UsageLimit:        promo.UsageLimit,
		UsageLimitPerUser: promo.UsageLimitPerUser,
		UsageCount:        promo.UsageCount,
		ValidFrom:         promo.ValidFrom,
		ValidUntil:        promo.ValidUntil,
		IsActive:          promo.IsActive,
		CreatedAt:         promo.CreatedAt,
	}
}

func ToPromoCodes(req *CreatePromoCodeRequest) *PromoCodes {
	return &PromoCodes{
		Code:              req.Code,
		Name:              req.Name,
		Description:       req.Description,
		Type:              req.Type,
		Value:             req.Value,
		MinOrderAmount:    req.MinOrderAmount,
		MaxDiscount:       req.MaxDiscount,
		UsageLimit:        req.UsageLimit,
		UsageLimitPerUser: req.UsageLimitPerUser,
		ValidFrom:         req.ValidFrom,
		ValidUntil:        req.ValidUntil,
		IsActive:          req.IsActive,
		UsageCount:        0,
	}
}
