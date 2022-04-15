package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeRemoveAccountController() *controllers.RemoveAccountController {
	repo := repositories.NewInmemoryAccountRepository()
	useCase := usecases.NewRemoveAccountUseCase(repo)
	controller := controllers.NewRemoveAccountController(useCase)
	return controller
}
