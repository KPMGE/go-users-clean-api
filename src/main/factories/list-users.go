package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeListUsersController() *controllers.ListUsersController {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewListUsersUseCase(repo)
	controller := controllers.NewListUsersController(useCase)
	return controller
}
