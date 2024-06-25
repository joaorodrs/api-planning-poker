package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/models"
)

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

	user, sessionId, err := AuthInput.SignIn()

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "session_id",
		Value: sessionId,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"session_id": sessionId,
		"user":       user,
	})
}
