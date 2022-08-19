package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeGetBookByIdController(db *gorm.DB) *controllers.GetBookByIdController {
	repo := postgresrepository.NewPostgresBookRepository(db)
	service := services.NewGetBookByIdService(repo)
	controller := controllers.NewGetBookByIdController(service)
	return controller
}
