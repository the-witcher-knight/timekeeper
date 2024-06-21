package auth

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")

	ErrInvalidHeader = errors.New("invalid header")

	ErrTokenExpired = errors.New("token expired")

	ErrInvalidSignature = errors.New("invalid signature")

	ErrInvalidUserProfile = errors.New("invalid user profile")

	ErrUnauthorized = errors.New("unauthorized")
)
