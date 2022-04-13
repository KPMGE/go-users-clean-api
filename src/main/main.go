package main

import (
	configuration "github.com/KPMGE/go-users-clean-api/src/main/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configuration.SetupRoutes(app)
	app.Listen(":3333")
}
