package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeDeleteUserController(db *gorm.DB) *controllers.DeleteUserController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	useCase := usecases.NewDeleteUserUseCase(repo)
	controller := controllers.NewDeleteUserController(useCase)
	return controller
}
