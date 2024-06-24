package models

import (
	"errors"

	"github.com/joaorodrs/api-planning-poker/utils"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (auth *Auth) SignIn() (*User, error) {
	var user User

	db.Where("Email=?", auth.Email).Find(&user)

	if user.ID == "" {
		return nil, errors.New("User not found")
	}

	validCredentials, _ := utils.ComparePasswords(user.Password, auth.Password)

	if !validCredentials {
		return nil, errors.New("Invalid credentials")
	}

	return &user, nil
}
