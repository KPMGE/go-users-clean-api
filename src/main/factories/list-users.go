package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeListUsersController(db *gorm.DB) *controllers.ListUsersController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	useCase := usecases.NewListUsersUseCase(repo)
	controller := controllers.NewListUsersController(useCase)
	return controller
}
