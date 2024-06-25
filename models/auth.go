package models

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/joaorodrs/api-planning-poker/database"
	"github.com/joaorodrs/api-planning-poker/utils"
)

var ctx = context.Background()

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (auth *Auth) SignIn() (*User, string, error) {
	var user User

	db.Where("Email=?", auth.Email).Find(&user)

	if user.ID == "" {
		return nil, "", errors.New("User not found")
	}

	validCredentials, _ := utils.ComparePasswords(user.Password, auth.Password)

	if !validCredentials {
		return nil, "", errors.New("Invalid credentials")
	}

	sessionId := make([]byte, 128)

	_, err := rand.Read(sessionId)
	if err != nil {
		return nil, "", errors.New("Could not generate session ID")
	}

	b, err := json.Marshal(user)
	if err != nil {
		return nil, "", errors.New("Could not parse USER")
	}

	rdb := database.GetRDB()

	stringSessionId := fmt.Sprintf("%x", sessionId)

	err = rdb.Set(ctx, fmt.Sprintf("session:%s", stringSessionId), string(b), 0).Err()

	if err != nil {
		return nil, "", errors.New("Could not store session ID in Redis")
	}

	return &user, stringSessionId, nil
}
