package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeListBooksController() *controllers.ListBooksController {
	repo := repositories.NewInMemoryBookRepository()
	useCase := usecases.NewListBookUseCase(repo)
	controller := controllers.NewListBooksController(useCase)
	return controller
}
