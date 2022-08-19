package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeDeleteUserController(db *gorm.DB) *controllers.DeleteUserController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	service := services.NewDeleteUserService(repo)
	controller := controllers.NewDeleteUserController(service)
	return controller
}
