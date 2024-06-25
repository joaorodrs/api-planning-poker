package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/api-planning-poker/handlers"
	"github.com/joaorodrs/api-planning-poker/middlewares"
)

var ctx = context.Background()

func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HandleHealthCheck)

	users := app.Group("/users", middlewares.SessionMiddleware)

	users.Get("/", handlers.HandleAllUsers)
	users.Get("/:id", handlers.HandleGetOneUser)
	users.Post("/", handlers.HandleCreateUser)
	users.Put("/:id", handlers.HandleUpdateUser)
	users.Delete("/:id", handlers.HandleDeleteUser)

	auth := app.Group("/auth")

	auth.Post("/login", handlers.Login)
}
