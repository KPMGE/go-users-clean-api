package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeAddBookController() *controllers.AddBookController {
	bookRepo := repositories.NewInMemoryBookRepository()
	userRepo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewAddBookUseCase(bookRepo, userRepo)
	controller := controllers.NewAddBookController(useCase)
	return controller
}
