package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeRemoveBookController() *controllers.RemoveBookController {
	repo := repositories.NewInMemoryBookRepository()
	userRepo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewRemoveBookUseCase(repo, repo, userRepo)
	controller := controllers.NewRemoveBookController(useCase)
	return controller
}
