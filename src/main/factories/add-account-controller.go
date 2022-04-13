package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/providers"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeAddAccountController() *controllers.AddAccountController {
	repo := repositories.NewInmemoryAccountRepository()
	hasher := providers.NewBcryptHasher()
	useCase := usecases.NewAddAccountUseCase(repo, hasher)
	controller := controllers.NewAddAccountController(useCase)
	return controller
}
