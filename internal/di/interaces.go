package di

import "os/user"

type UserRepository interface {
	FindByEmail(email string) (*user.User, error)
}
