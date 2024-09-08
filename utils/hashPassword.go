package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(HashPassword), nil
}

func VerifyPassword(HashPassword, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(Password))
}
