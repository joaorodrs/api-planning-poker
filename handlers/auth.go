package handlers

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/database"
	"github.com/joaorodrs/api-planning-poker/models"
)

var ctx = context.Background()

// @Summary Login
// @Description Log-in
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Router /auth/login [get]
func Login(c *fiber.Ctx) error {
	AuthInput := &models.Auth{}

	if err := c.BodyParser(AuthInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": err.Error()})
	}

	user, err := AuthInput.SignIn()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": err.Error()})
	}

	sessionId := make([]byte, 128)

	_, err = rand.Read(sessionId)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	b, err := json.Marshal(user)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	rdb := database.GetRDB()

	err = rdb.Set(ctx, fmt.Sprintf("session:%v", sessionId), string(b), 0).Err()

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"session_id": sessionId,
		"user":       user,
	})
}
