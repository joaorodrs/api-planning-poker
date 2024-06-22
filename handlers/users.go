package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/joaorodrs/api-planning-poker/models"
	"github.com/joaorodrs/api-planning-poker/utils"
)

// @Summary Get All Users
// @Description fetch all users
// @Tags Users
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
func HandleAllUsers(c *fiber.Ctx) error {
	PageStr := c.Query("Page", "1")
	Page, err := strconv.Atoi(PageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Could not convert Page param"})
	}
	TakeStr := c.Query("Take", "10")
	Take, err := strconv.Atoi(TakeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Could not convert Take param"})
	}

	fmt.Println(Page)
	fmt.Println(Take)

	Pagination := &utils.Pagination{Page: Page, Take: Take}

	users, pagination := models.GetAllUsers(Pagination)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": users, "pagination": pagination})
}

// @Summary Create User
// @Description create an user
// @Param user body models.User true "User payload"
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users [post]
func HandleCreateUser(c *fiber.Ctx) error {
	CreateUser := &models.User{}

	if err := c.BodyParser(CreateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(CreateUser.Password)

	if err != nil {
		return err
	}

	CreateUser.Password = hashedPassword

	user := CreateUser.CreateUser()

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Update User
// @Description update an user
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Router /users/:id [put]
func HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	UpdateUser := &models.User{}

	if err := c.BodyParser(UpdateUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Bad request": err.Error()})
	}

	user := models.UpdateUser(id, *UpdateUser)

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Get User By ID
// @Description get an user
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Router /users/:id [get]
func HandleGetOneUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, _ := models.GetUserById(id)

	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary Delete User
// @Description delete an user
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Router /users/:id [delete]
func HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	models.DeleteUser(id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User deleted successfully"})
}
