package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/handlers"
)

func Example(c *fiber.Ctx) error {
	return nil
}

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/todos")
	todos.Get("/", Example)
}
