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

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Init() *gorm.DB {
	dbURL := "postgresql://postgres:root@localhost:5432/users"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Migrating entities...")
	err = db.AutoMigrate(&entities.Book{})
	CheckError(err)
	err = db.AutoMigrate(&entities.Account{})
	CheckError(err)
	err = db.AutoMigrate(&entities.User{})
	CheckError(err)

	return db
}

func main() {
	db := Init()
	// db1 := postgresrepository.GetPostgresConnection()
	app := fiber.New()
	configuration.SetupRoutes(app, db)
	app.Listen(":3333")
}
