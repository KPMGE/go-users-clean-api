package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeAddBookController(db *gorm.DB) *controllers.AddBookController {
	bookRepo := postgresrepository.NewPostgresBookRepository(db)
	userRepo := postgresrepository.NewPostgresUserRepository(db)
	useCase := usecases.NewAddBookUseCase(bookRepo, userRepo)
	controller := controllers.NewAddBookController(useCase)
	return controller
}
