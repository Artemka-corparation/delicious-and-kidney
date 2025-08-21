package restaurant

type repository interface {
	Create(restaurant *Restaurant) (*Restaurant, error)
	FindById(id uint) (*Restaurant, error)
	Update(restaurant *Restaurant) (*Restaurant, error)
	UpdateFields(id uint, updates map[string]interface{}) error
	Delete(id uint) error
	HardDelete(id uint) error
	FindByOwnerId(ownerID uint) ([]Restaurant, error)
	FindAll(limit, offset int) ([]Restaurant, error)
	FindActiveRestaurants(limit, offset int) ([]Restaurant, error)
	FindFeaturedRestaurants(limit, offset int) ([]Restaurant, error)
	FindByLocation(lat, lng, radius float64) ([]Restaurant, error)
	FindByName(name string) ([]Restaurant, error)
	SearchRestaurants(query string, limit, offset int) ([]Restaurant, error)
	Count() (int64, error)
	CountByOwnerId(ownerID uint) (int64, error)
}

type service interface {
	CreateRestaurant(ownerID uint, req *CreateRestaurantRequest) (*RestaurantResponse, error)
	GetRestaurant(id uint) (*RestaurantResponse, error)
	UpdateRestaurant(id uint, ownerID uint, req *UpdateRestaurantRequest) (*RestaurantResponse, error)
	DeleteRestaurant(id uint, ownerID uint) error
	ActivateRestaurant(id uint, ownerID uint) error
	DeactivateRestaurant(id uint, ownerID uint) error
	SetFeaturedStatus(id uint, featured bool, userRole string) error // Только для админов
	GetMyRestaurants(ownerID uint, limit, offset int) ([]RestaurantResponse, error)
	GetAllRestaurants(limit, offset int) ([]RestaurantResponse, error)
	SearchRestaurants(query string, limit, offset int) ([]RestaurantResponse, error)
	GetNearbyRestaurants(lat, lng, radius float64) ([]RestaurantResponse, error)
	GetRestaurantStats(id uint, ownerID uint) (*RestaurantStats, error)
}
