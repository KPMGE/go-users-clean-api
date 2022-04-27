package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeGetBookByIdController() *controllers.GetBookByIdController {
	repo := repositories.NewInMemoryBookRepository()
	useCase := usecases.NewGetBookByIdUseCase(repo)
	controller := controllers.NewGetBookByIdController(useCase)
	return controller
}
