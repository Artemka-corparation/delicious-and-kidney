package restaurant

import (
	"delicious-and-kidney/internal/user"
	"delicious-and-kidney/pkg/Errors"
	"errors"
	"gorm.io/gorm"
)

type RestaurantService struct {
	restaurantRepo repository
	userRepo       user.Repository
}

func NewRestaurantService(repo repository, userRepo user.Repository) *RestaurantService {
	return &RestaurantService{
		restaurantRepo: repo,
		userRepo:       userRepo,
	}
}

func (s *RestaurantService) CreateRestaurant(ownerID uint, req *CreateRestaurantRequest) (*RestaurantResponse, error) {
	_, err := s.userRepo.FindById(ownerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrOwnerNotFound
		}
		return nil, err
	}
	restaurant := ToRestaurant(req, ownerID)

	savedRestaurant, err := s.restaurantRepo.Create(restaurant)
	if err != nil {
		return nil, err
	}
	response := ToRestaurantResponse(savedRestaurant)

	return response, nil

}

func (s *RestaurantService) GetRestaurant(id uint) (*RestaurantResponse, error) {
	restaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrRestaurantNotFound
		}
		return nil, err
	}
	response := ToRestaurantResponse(restaurant)

	return response, nil
}

func (s *RestaurantService) UpdateRestaurant(id uint, ownerID uint, req *UpdateRestaurantRequest) (*RestaurantResponse, error) {
	restaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrRestaurantNotFound
		}
		return nil, err
	}
	if restaurant.OwnerID != ownerID {
		return nil, Errors.ErrUnauthorized
	}
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	if req.Email != nil {
		updates["email"] = *req.Email
	}
	if req.Address != nil {
		updates["address"] = *req.Address
	}
	if req.Latitude != nil {
		updates["latitude"] = *req.Latitude
	}
	if req.Longitude != nil {
		updates["longitude"] = *req.Longitude
	}
	if req.ImageURL != nil {
		updates["image_url"] = *req.ImageURL
	}
	if req.MinOrderAmount != nil {
		updates["min_order_amount"] = *req.MinOrderAmount
	}
	if req.DeliveryFee != nil {
		updates["delivery_fee"] = *req.DeliveryFee
	}
	if req.DeliveryTimeMin != nil {
		updates["delivery_time_min"] = *req.DeliveryTimeMin
	}
	if req.DeliveryTimeMax != nil {
		updates["delivery_time_max"] = *req.DeliveryTimeMax
	}
	if req.IsActive != nil {
		updates["is_active"] = *req.IsActive
	}
	if req.IsFeatured != nil {
		updates["is_featured"] = *req.IsFeatured
	}
	err = s.restaurantRepo.UpdateFields(id, updates)
	if err != nil {
		return nil, err
	}
	updatedRestaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return ToRestaurantResponse(updatedRestaurant), nil
}

func (s *RestaurantService) DeleteRestaurant(id uint, ownerID uint) error {
	restaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrRestaurantNotFound
		}
		return err
	}
	if restaurant.OwnerID != ownerID {
		return Errors.ErrUnauthorized
	}
	err = s.restaurantRepo.Delete(id)
	return err
}

func (s *RestaurantService) ActivateRestaurant(id uint, ownerID uint) error {
	restaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrRestaurantNotFound
		}
		return err
	}
	if restaurant.OwnerID != ownerID {
		return Errors.ErrUnauthorized
	}
	updates := map[string]interface{}{
		"is_active": true,
	}
	err = s.restaurantRepo.UpdateFields(id, updates)
	return err
}

func (s *RestaurantService) DeactivateRestaurant(id uint, ownerID uint) error {
	restaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrRestaurantNotFound
		}
		return err
	}
	if restaurant.OwnerID != ownerID {
		return Errors.ErrUnauthorized
	}
	updates := map[string]interface{}{
		"deactivate": false,
	}
	err = s.restaurantRepo.UpdateFields(id, updates)
	return err
}

func (s *RestaurantService) SetFeaturedStatus(id uint, featured bool, userRole string) error {
	if userRole != "admin" {
		return Errors.ErrUnauthorized
	}
	_, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrRestaurantNotFound
		}
		return err
	}
	updates := map[string]interface{}{
		"is_featured": featured,
	}
	err = s.restaurantRepo.UpdateFields(id, updates)
	return err
}

func (s *RestaurantService) GetMyRestaurants(ownerID uint, limit, offset int) ([]RestaurantResponse, error) {
	restaurant, err := s.restaurantRepo.FindByOwnerId(ownerID)
	if err != nil {
		return nil, err
	}
	if offset >= len(restaurant) {
		return []RestaurantResponse{}, nil
	}
	restaurants := restaurant[offset:]
	if limit >= len(restaurants) {
		restaurants = restaurants[:limit]
	}
	response := make([]RestaurantResponse, len(restaurants))
	for i, rest := range restaurants {
		response[i] = *ToRestaurantResponse(&rest)
	}
	return response, nil
}

func (s *RestaurantService) GetAllRestaurants(limit, offset int) ([]RestaurantResponse, error) {
	activeRestarant, err := s.restaurantRepo.FindActiveRestaurants(limit, offset)
	if err != nil {
		return nil, err
	}
	response := make([]RestaurantResponse, len(activeRestarant))
	for i, rest := range activeRestarant {
		response[i] = *ToRestaurantResponse(&rest)
	}
	return response, nil
}

func (s *RestaurantService) SearchRestaurants(query string, limit, offset int) ([]RestaurantResponse, error) {
	result, err := s.restaurantRepo.SearchRestaurants(query, limit, offset)
	if err != nil {
		return nil, err
	}
	response := make([]RestaurantResponse, len(result))
	for i, rest := range result {
		response[i] = *ToRestaurantResponse(&rest)
	}
	return response, nil
}

func (s *RestaurantService) GetNearbyRestaurants(lat, lng, radius float64) ([]RestaurantResponse, error) {
	result, err := s.restaurantRepo.FindByLocation(lat, lng, radius)
	if err != nil {
		return nil, err
	}
	response := make([]RestaurantResponse, len(result))
	for i, rest := range result {
		response[i] = *ToRestaurantResponse(&rest)
	}
	return response, nil
}

func (s *RestaurantService) GetRestaurantStats(id uint, ownerID uint) (*RestaurantStats, error) {
	restaurant, err := s.restaurantRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrRestaurantNotFound
		}
		return nil, err
	}
	if restaurant.OwnerID != ownerID {
		return nil, Errors.ErrUnauthorized
	}
	stats := &RestaurantStats{}
	return stats, err
	//todo пока нет заказов это просто заглушка
}
