package user

import (
	"delicious-and-kidney/pkg/Errors"
	"errors"
	"time"

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

func (service *UserService) CreateUser(user *User) (*UserResponse, error) {
	_, err := service.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		ID:            user.Id,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Role,
		EmailVerified: user.EmailVerified,
		PhoneVerified: user.PhoneVerified,
		CreatedAt:     time.Time{},
	}, nil
}
