package auth

import "delicious-and-kidney/pkg/db"

type AuthRepository struct {
	database *db.Db
}

func NewAuthRepository(database *db.Db) *AuthRepository {
	return &AuthRepository{database: database}
}
