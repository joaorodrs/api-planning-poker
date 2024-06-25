package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err == nil {
		return true, nil
	}

	return false, err
}
