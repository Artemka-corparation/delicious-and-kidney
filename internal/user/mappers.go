package user

func ToUserResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:            user.Id,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Role:          user.Role,
		EmailVerified: user.EmailVerified,
		PhoneVerified: user.PhoneVerified,
		CreatedAt:     user.CreatedAt,
	}

}
