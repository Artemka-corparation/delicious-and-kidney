package auth

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewClaims(userId int, email string, role string, registeredClaims jwt.RegisteredClaims) *Claims {
	return &Claims{UserId: userId, Email: email, Role: role, RegisteredClaims: registeredClaims}
}

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		if os.Getenv("GIN_MODE") != "production" {
			randomKey := generateRandomKey(32)
			log.Println(" JWT_SECRET не установлен, используется НЕБЕЗОПАСНЫЙ ключ для разработки")
			return randomKey
		}
		panic("JWT_SECRET обязательно должен быть установлен в продакшене!")
	}
	return secret
}

func generateRandomKey(length int) string {
	bates := make([]byte, length)
	_, err := rand.Read(bates)
	if err != nil {
		panic("не удалось сгенерировать рандомный ключ")
	}
	return hex.EncodeToString(bates)
}
func (j *JWTService) GenerateToken(userID int, email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := Claims{
		UserId: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (j *JWTService) ValidateToken(tokenString string) (*Claims, error) {

	//ToDO 1. Распарсить токен
	//ToDO 2. Проверить подпись
	//ToDO 3. Извлечь Claims
	//ToDO 4. Проверить что токен валидный
	return nil, nil
}
