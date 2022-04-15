package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeAddUserController() *controllers.AddUserController {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewAddUserUseCase(repo)
	controller := controllers.NewAddUserController(useCase)
	return controller
}
