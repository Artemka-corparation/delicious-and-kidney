package Errors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidInput       = errors.New("invalid input")
	ErrWrongPassword      = errors.New("wrong password")
	ErrWeakPassword       = errors.New("weak password")
	ErrSamePassword       = errors.New("password same")
	ErrInvalidRole        = errors.New("invalid user role")
	ErrUnauthorized       = errors.New("unauthorized action")
	ErrRestaurantNotFound = errors.New("restaurant not found")
	ErrOwnerNotFound      = errors.New("owner not found")
	ErrPromoCodeNotFound  = errors.New("promo code not found")
	ErrorPromoMessage     = errors.New("Promo code is inactive")
)
