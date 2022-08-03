package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/providers"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeAddAccountController(db *gorm.DB) *controllers.AddAccountController {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	hasher := providers.NewBcryptHasher()
	useCase := usecases.NewAddAccountUseCase(repo, hasher)
	controller := controllers.NewAddAccountController(useCase)
	return controller
}
