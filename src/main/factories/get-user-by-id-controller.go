package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeGetUserByIdController() *controllers.GetUserByIdController {
	repo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewGetUserByIdUseCase(repo)
	controller := controllers.NewGetUserByIdController(useCase)
	return controller
}
