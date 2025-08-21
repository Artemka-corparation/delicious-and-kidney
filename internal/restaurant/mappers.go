package restaurant

func ToRestaurantResponse(restaurant *Restaurant) *RestaurantResponse {
	return &RestaurantResponse{
		ID:              restaurant.ID,
		OwnerID:         restaurant.OwnerID,
		Name:            restaurant.Name,
		Description:     restaurant.Description,
		Phone:           restaurant.Phone,
		Email:           restaurant.Email,
		Address:         restaurant.Address,
		Latitude:        restaurant.Latitude,
		Longitude:       restaurant.Longitude,
		ImageURL:        restaurant.ImageURL,
		MinOrderAmount:  restaurant.MinOrderAmount,
		DeliveryFee:     restaurant.DeliveryFee,
		DeliveryTimeMin: restaurant.DeliveryTimeMin,
		DeliveryTimeMax: restaurant.DeliveryTimeMax,
		Rating:          restaurant.Rating,
		ReviewsCount:    restaurant.ReviewsCount,
		IsActive:        restaurant.IsActive,
		IsFeatured:      restaurant.IsFeatured,
		CreatedAt:       restaurant.CreatedAt,
		UpdatedAt:       restaurant.UpdatedAt,
	}
}

func ToRestaurant(req *CreateRestaurantRequest, ownerID uint) *Restaurant {
	return &Restaurant{
		OwnerID:         ownerID,
		Name:            req.Name,
		Description:     req.Description,
		Phone:           req.Phone,
		Email:           req.Email,
		Address:         req.Address,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
		ImageURL:        req.ImageURL,
		MinOrderAmount:  req.MinOrderAmount,
		DeliveryFee:     req.DeliveryFee,
		DeliveryTimeMin: req.DeliveryTimeMin,
		DeliveryTimeMax: req.DeliveryTimeMax,
		Rating:          0,
		ReviewsCount:    0,
		IsActive:        true,
		IsFeatured:      false,
	}
}
