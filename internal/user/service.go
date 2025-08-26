package user

import (
	"delicious-and-kidney/pkg/Errors"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo Repository
}

func NewUserService(userRepo Repository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (service *UserService) GetProfile(Id uint) (*UserResponse, error) {
	user, err := service.userRepo.FindById(Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrUserNotFound
		}
		return nil, err
	}
	response := &UserResponse{
		ID:            user.Id,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		PhoneVerified: user.PhoneVerified,
		CreatedAt:     user.CreatedAt,
	}
	return response, nil
}

func (service *UserService) UpdateProfile(Id uint, req *UpdateUserRequest) (*UserResponse, error) {
	_, err := service.userRepo.FindById(Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, Errors.ErrUserNotFound
		}
		return nil, err
	}
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}
	err = service.userRepo.UpdateFields(Id, updates)
	if err != nil {
		return nil, err
	}
	return service.GetProfile(Id)
}

func (service *UserService) CreateUser(req *CreateUserRequest) (*UserResponse, error) {
	_, err := service.userRepo.FindByEmail(req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	} else {
		return nil, errors.New("email already exists")
	}
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		Role:         req.Role,
	}
	userCreate, err := service.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return ToUserResponse(userCreate), nil
}

func (service *UserService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := service.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrUserNotFound
		}
		return err
	}
	err = checkPassword(user.PasswordHash, oldPassword)
	if err != nil {
		return Errors.ErrWrongPassword
	}
	if len(newPassword) < 8 {
		return Errors.ErrWeakPassword
	}
	if oldPassword == newPassword {
		return Errors.ErrSamePassword
	}
	newhashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{
		"password_hash": newhashedPassword,
	}
	err = service.userRepo.UpdateFields(userID, updates)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) DeactivateAccount(userID uint) error {
	_, err := service.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrUserNotFound
		}
		return err
	}
	updates := map[string]interface{}{
		"is_active": false,
	}
	err = service.userRepo.UpdateFields(userID, updates)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) DeleteAccount(userID uint) error {
	_, err := service.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrUserNotFound
		}
		return err
	}
	err = service.userRepo.Delete(userID)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) ActivateAccount(userID uint) error {
	_, err := service.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrUserNotFound
		}
		return err
	}
	updates := map[string]interface{}{
		"is_active": true,
	}
	err = service.userRepo.UpdateFields(userID, updates)
	if err != nil {
		return err
	}
	return nil
}
