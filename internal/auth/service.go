package auth

type AuthService struct {
	authRepository *AuthRepository
}

func NewAuthService(repository *AuthRepository) *AuthService {
	return &AuthService{authRepository: repository}
}
