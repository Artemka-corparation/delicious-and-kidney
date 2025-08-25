package promo_codes

import (
	"delicious-and-kidney/pkg/db"
	"gorm.io/gorm"
)

type PromoCodeRepository struct {
	database *db.Db
}

func NewPromoCodeRepository(database *db.Db) *PromoCodeRepository {
	return &PromoCodeRepository{database: database}
}

func (repo *PromoCodeRepository) Create(promo *PromoCodes) (*PromoCodes, error) {
	result := repo.database.Create(promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return promo, nil
}

func (repo *PromoCodeRepository) FindById(id uint) (*PromoCodes, error) {
	var promo PromoCodes
	result := repo.database.First(&promo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &promo, nil
}

func (repo *PromoCodeRepository) FindByCode(code string) (*PromoCodes, error) {
	var promo PromoCodes
	result := repo.database.Where("code = ?", code).First(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &promo, nil
}
func (repo *PromoCodeRepository) Update(promo *PromoCodes) (*PromoCodes, error) {
	result := repo.database.Save(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return promo, nil
}

func (repo *PromoCodeRepository) UpdateFields(id uint, updates map[string]interface{}) error {
	result := repo.database.Model(&PromoCodes{}).Where("id = ?", id).Updates(updates)
	return result.Error
}

func (repo *PromoCodeRepository) Delete(id uint) error {
	result := repo.database.Delete(&PromoCodes{}, id)
	return result.Error
}

func (repo *PromoCodeRepository) FindActivePromoCodes() ([]PromoCodes, error) {
	promo := make([]PromoCodes, 0)
	result := repo.database.Where("is_active = true").Find(&promo)
	if result.Error != nil {
		return nil, result.Error
	}
	return promo, nil
}

func (repo *PromoCodeRepository) FindValidPromoCodes() ([]PromoCodes, error) {
	var promoCodes []PromoCodes
	result := repo.database.
		Where("is_active = ?", true).
		Where("valid_from <= NOW()").
		Where("valid_until IS NULL OR valid_until >= NOW()").
		Where("usage_limit IS NULL OR usage_count < usage_limit").
		Find(&promoCodes)
	if result.Error != nil {
		return nil, result.Error
	}
	return promoCodes, nil
}

func (repo *PromoCodeRepository) IncrementUsageCount(id uint) error {
	result := repo.database.Model(&PromoCodes{}).
		Where("id = ?", id).
		Update("usage_count", gorm.Expr("usage_count + ?", 1))

	return result.Error
}

func (repo *PromoCodeRepository) FindByType(promoType string) ([]PromoCodes, error) {
	var promoCodes []PromoCodes
	result := repo.database.Where("type = ?", promoType).Find(&promoCodes)
	if result.Error != nil {
		return nil, result.Error
	}
	return promoCodes, nil
}

func (repo *PromoCodeRepository) Count() (int64, error) {
	var count int64
	result := repo.database.Model(&PromoCodes{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (repo *PromoCodeRepository) FindAll() ([]PromoCodes, error) {
	var promoCodes []PromoCodes
	result := repo.database.Find(&promoCodes)
	if result.Error != nil {
		return nil, result.Error
	}
	return promoCodes, nil
}
