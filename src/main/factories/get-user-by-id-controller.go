package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeGetUserByIdController(db *gorm.DB) *controllers.GetUserByIdController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	service := services.NewGetUserByIdService(repo)
	controller := controllers.NewGetUserByIdController(service)
	return controller
}
