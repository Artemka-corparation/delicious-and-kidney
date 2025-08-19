package user

type UserRepositoryInterface interface {
	Create(user *User) (*User, error)
	FindById(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) (*User, error)
	UpdateFields(id uint, updates map[string]interface{}) error
	Delete(id uint) error
	HardDelete(id uint) error
}

type UserServiceInterface interface {
	GetProfile(Id uint) (*UserResponse, error)
	UpdateProfile(Id uint, req *UpdateUserRequest) (*UserResponse, error)
}
