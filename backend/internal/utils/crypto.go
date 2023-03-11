package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordValidation = errors.New("password validation failed")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ValidatePassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return ErrPasswordValidation
	}
	return nil
}
