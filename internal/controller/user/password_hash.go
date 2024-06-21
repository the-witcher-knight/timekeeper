package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Wrapper function, it used for unit test
var (
	hashPasswordWrapper     = hashPassword
	comparePasswordsWrapper = comparePasswords
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func comparePasswords(hashedPwd string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	if err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
