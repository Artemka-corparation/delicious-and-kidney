package user

import "delicious-and-kidney/pkg/db"

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{database: database}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.database.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindById(Id uint) (*User, error) {
	var user User
	result := repo.database.First(&user, Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.database.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *User) (*User, error) {
	result := repo.database.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) UpdateFields(id uint, updates map[string]interface{}) error {
	result := repo.database.Model(&User{}).Where("id = ?", id).Updates(updates)
	return result.Error
}

func (repo *UserRepository) Delete(id uint) error {
	result := repo.database.Delete(&User{}, id)
	return result.Error
}

func (repo *UserRepository) HardDelete(id uint) error {
	result := repo.database.Unscoped().Delete(&User{}, id)
	return result.Error
}
