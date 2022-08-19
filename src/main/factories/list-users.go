package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeListUsersController(db *gorm.DB) *controllers.ListUsersController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	service := services.NewListUsersService(repo)
	controller := controllers.NewListUsersController(service)
	return controller
}
