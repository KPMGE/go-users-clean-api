package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/main/factories/validators"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeAddBookController(db *gorm.DB) *controllers.AddBookController {
	bookRepo := postgresrepository.NewPostgresBookRepository(db)
	userRepo := postgresrepository.NewPostgresUserRepository(db)
	service := services.NewAddBookService(bookRepo, userRepo)
	validator := validators.MakeAddBookValidation()
	controller := controllers.NewAddBookController(service, validator)
	return controller
}
