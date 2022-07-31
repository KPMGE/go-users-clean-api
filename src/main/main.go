package main

import (
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	configuration "github.com/KPMGE/go-users-clean-api/src/main/config"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	db := postgresrepository.GetPostgresConnection()
	app := fiber.New()
	configuration.SetupRoutes(app, db)
	app.Listen(":3333")
}
