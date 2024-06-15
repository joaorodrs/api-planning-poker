package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/joaorodrs/api-planning-poker/models"
)

// @Summary Get All Users
// @Description fetch all users
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
func HandleAllUsers(c *fiber.Ctx) error {
	users := models.GetAllUsers()

	return c.Status(fiber.StatusOK).JSON(users)
}

func HandleCreateUser(c *fiber.Ctx) error {
	CreateUser := &models.User{}

	if err := c.BodyParser(CreateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": err.Error()})
	}

	user := CreateUser.CreateUser()

	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	UpdateUser := &models.User{}

	if err := c.BodyParser(UpdateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": err.Error()})
	}

	userId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": "Could not parse ID param"})
	}

	user := models.UpdateUser(userId, *UpdateUser)

	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleGetOneUser(c *fiber.Ctx) error {
	id := c.Params("id")

	userId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": "Could not parse ID param"})
	}

	user, _ := models.GetUserById(userId)

	return c.Status(fiber.StatusOK).JSON(user)
}

func HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	userId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": "Could not parse ID param"})
	}

	models.DeleteUser(userId)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User deleted successfully"})
}
