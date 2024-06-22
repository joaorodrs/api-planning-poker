package handlers

import "github.com/gofiber/fiber/v2"

// @Summary Login
// @Description Log-in
// @Tags Auth
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Router /auth/login [get]
func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "All ok"})
}
