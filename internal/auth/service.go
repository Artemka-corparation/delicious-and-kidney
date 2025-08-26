package auth

import (
	"delicious-and-kidney/internal/user"
	"delicious-and-kidney/pkg/Errors"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	authRepository *AuthRepository
	userRepo       user.Repository
	jwtService     *JWTService
}

func NewAuthService(repository *AuthRepository, userRepo user.Repository, jwtService *JWTService) *AuthService {
	return &AuthService{
		authRepository: repository,
		userRepo:       userRepo,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Register(req *RegisterRequest) (*AuthResponse, error) {
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	role := req.Role
	if role == "" {
		role = "customer"
	}

	newUser := &user.User{
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: string(hashedPassword),
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdUser, err := s.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateToken(createdUser.Id, createdUser.Email, createdUser.Role)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: UserInfo{
			ID:            createdUser.Id,
			Name:          createdUser.Name,
			Email:         createdUser.Email,
			Phone:         createdUser.Phone,
			Role:          createdUser.Role,
			EmailVerified: createdUser.EmailVerified,
			PhoneVerified: createdUser.PhoneVerified,
			CreatedAt:     createdUser.CreatedAt,
		},
		Token: token,
	}, nil
}

func (s *AuthService) Login(req *LoginRequest) (*AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email of password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwtService.GenerateToken(user.Id, user.Email, user.Role)
	if err != nil {
		return nil, err
	}
	return &AuthResponse{
		User: UserInfo{
			ID:            user.Id,
			Name:          user.Name,
			Email:         user.Email,
			Phone:         user.Phone,
			Role:          user.Role,
			EmailVerified: user.EmailVerified,
			PhoneVerified: user.PhoneVerified,
			CreatedAt:     user.CreatedAt,
		},
		Token: token,
	}, nil
}

func (s *AuthService) RefreshToken(req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	newToken, err := s.jwtService.RefreshToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &RefreshTokenResponse{
		Token: newToken,
	}, nil
}

func (s *AuthService) ChangePassword(userID uint, req *ChangePasswordRequest) error {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Errors.ErrUserNotFound
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.CurrentPassword))
	if err != nil {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	updates := map[string]interface{}{
		"password_hash": string(hashedPassword),
		"update_at":     time.Now(),
	}
	return s.userRepo.UpdateFields(userID, updates)
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	return s.jwtService.ValidateToken(tokenString)
}
