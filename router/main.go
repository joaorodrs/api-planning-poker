package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	// setup the todos group
	todos := app.Group("/users")
	todos.Get("/", handlers.HandleAllUsers)
	todos.Get("/:id", handlers.HandleGetOneUser)
	todos.Post("/", handlers.HandleCreateUser)
	todos.Put("/:id", handlers.HandleUpdateUser)
	todos.Delete("/:id", handlers.HandleDeleteUser)
}
