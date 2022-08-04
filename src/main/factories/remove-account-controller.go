package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeRemoveAccountController(db *gorm.DB) *controllers.RemoveAccountController {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	useCase := usecases.NewRemoveAccountUseCase(repo)
	controller := controllers.NewRemoveAccountController(useCase)
	return controller
}
