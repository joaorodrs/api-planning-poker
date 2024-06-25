package router

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/database"
	"github.com/joaorodrs/api-planning-poker/handlers"
)

var ctx = context.Background()

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	users := app.Group("/users", func(c *fiber.Ctx) error {
		sessionId := c.Cookies("session_id")

		rdb := database.GetRDB()

		storedSessionId, err := rdb.Get(ctx, fmt.Sprintf("session:%s", sessionId)).Result()

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Unauthorized"})
		}

		fmt.Println(storedSessionId)

		return c.Next()
	})

	users.Get("/", handlers.HandleAllUsers)
	users.Get("/:id", handlers.HandleGetOneUser)
	users.Post("/", handlers.HandleCreateUser)
	users.Put("/:id", handlers.HandleUpdateUser)
	users.Delete("/:id", handlers.HandleDeleteUser)

	auth := app.Group("/auth")

	auth.Post("/login", handlers.Login)
}
