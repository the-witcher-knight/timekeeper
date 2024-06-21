package user

import (
	"errors"
)

var (
	ErrEmailExisted = errors.New("email is already exists")

	ErrUserEmailNotExists = errors.New("user email does not exist")

	ErrPasswordNotMatch = errors.New("password not match")
)
