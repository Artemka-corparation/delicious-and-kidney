package Errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidInput       = errors.New("invalid input")
	ErrRestaurantNotFound = errors.New("restaurant not found")
	ErrOwnerNotFound      = errors.New("owner not found")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrPromoCodeNotFound  = errors.New("promo code not found")
	ErrorPromoMessage     = errors.New("Promo code is inactive")
)
