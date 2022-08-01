package factories

import (
	"database/sql"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeAddBookController(db *sql.DB) *controllers.AddBookController {
	bookRepo := postgresrepository.NewPostgresBookRepository(db)
	userRepo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewAddBookUseCase(bookRepo, userRepo)
	controller := controllers.NewAddBookController(useCase)
	return controller
}
