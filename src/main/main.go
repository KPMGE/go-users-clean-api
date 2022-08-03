package main

import (
	"fmt"
	"log"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	configuration "github.com/KPMGE/go-users-clean-api/src/main/config"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgresql://postgres:root@localhost:5432/users"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.Account{})

	fmt.Println("complete!!")

	return db
}

func main() {
	db := Init()
	// db1 := postgresrepository.GetPostgresConnection()
	app := fiber.New()
	configuration.SetupRoutes(app, db)
	app.Listen(":3333")
}
