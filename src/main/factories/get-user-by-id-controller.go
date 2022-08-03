package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeGetUserByIdController(db *gorm.DB) *controllers.GetUserByIdController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	useCase := usecases.NewGetUserByIdUseCase(repo)
	controller := controllers.NewGetUserByIdController(useCase)
	return controller
}
