package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joaorodrs/api-planning-poker/config"
	"github.com/joaorodrs/api-planning-poker/database"
	"github.com/joaorodrs/api-planning-poker/router"
	_ "github.com/joho/godotenv/autoload"
)

func SetupAndRunApp() error {
	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	cookieSecret := os.Getenv("COOKIE_SECRET")
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: cookieSecret,
	}))

	app.Use(config.SetupSecurityHeaders)

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// connect redis
	database.ConnectRedis()

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
