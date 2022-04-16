package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeDeleteUserController() *controllers.DeleteUserController {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewDeleteUserUseCase(repo)
	controller := controllers.NewDeleteUserController(useCase)
	return controller
}
