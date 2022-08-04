package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeListBooksController(db *gorm.DB) *controllers.ListBooksController {
	repo := postgresrepository.NewPostgresBookRepository(db)
	useCase := usecases.NewListBookUseCase(repo)
	controller := controllers.NewListBooksController(useCase)
	return controller
}
