package middlewares

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/database"
	"github.com/joaorodrs/api-planning-poker/models"
	"github.com/joaorodrs/api-planning-poker/utils"
)

var ctx = context.Background()

func SessionMiddleware(c *fiber.Ctx) error {
	sessionId := c.Cookies("session_id")

	rdb := database.GetRDB()

	user, err := rdb.Get(ctx, fmt.Sprintf("session:%s", sessionId)).Result()

	if err != nil || user == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Unauthorized"})
	}

	utils.ParseBody([]byte(user), models.User{})

	c.Locals("user", user)

	return c.Next()
}
