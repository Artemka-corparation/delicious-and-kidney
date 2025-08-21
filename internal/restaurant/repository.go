package restaurant

import "delicious-and-kidney/pkg/db"

type RestaurantRepository struct {
	database *db.Db
}

func NewRestaurantRepository(database *db.Db) *RestaurantRepository {
	return &RestaurantRepository{database: database}
}

func (repo *RestaurantRepository) Create(restaurant *Restaurant) (*Restaurant, error) {
	result := repo.database.Create(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FindById(Id uint) (*Restaurant, error) {
	var restaurant Restaurant
	result := repo.database.First(&restaurant, Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &restaurant, nil
}

func (repo *RestaurantRepository) Update(restaurant *Restaurant) (*Restaurant, error) {
	result := repo.database.Save(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) UpdateFields(id uint, updates map[string]interface{}) error {
	result := repo.database.Model(&Restaurant{}).Where("id = ?", id).Updates(updates)
	return result.Error
}

func (repo *RestaurantRepository) Delete(id uint) error {
	result := repo.database.Delete(&Restaurant{}, id)
	return result.Error
}

func (repo *RestaurantRepository) HardDelete(id uint) error {
	result := repo.database.Unscoped().Delete(&Restaurant{}, id)
	return result.Error
}
func (repo *RestaurantRepository) FindByOwnerId(ownerID uint) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Where("owner_id = ?", ownerID).Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FindAll(limit, offset int) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Limit(limit).Offset(offset).Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FindActiveRestaurants(limit, offset int) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Where("is_active = true").Limit(limit).Offset(offset).Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FindFeaturedRestaurants(limit, offset int) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Where("is_featured = true").Limit(limit).Offset(offset).Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FindByLocation(lat, lng, radius float64) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Where("ST_DWithin(ST_Point(longitude, latitude), ST_Point(?, ?), ?)", lng, lat, radius).Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FindByName(name string) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Where("name ILIKE ?", "%"+name+"%").Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) SearchRestaurants(query string, limit, offset int) ([]Restaurant, error) {
	restaurant := make([]Restaurant, 0)
	result := repo.database.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").Limit(limit).Offset(offset).Find(&restaurant)
	if result.Error != nil {
		return nil, result.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) Count() (int64, error) {
	var count int64
	result := repo.database.Model(&Restaurant{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (repo *RestaurantRepository) CountByOwnerId(ownerID uint) (int64, error) {
	var count int64
	result := repo.database.Model(&Restaurant{}).Where("owner_id = ?", ownerID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
