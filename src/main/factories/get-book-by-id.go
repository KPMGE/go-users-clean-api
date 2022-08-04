package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeGetBookByIdController(db *gorm.DB) *controllers.GetBookByIdController {
	repo := postgresrepository.NewPostgresBookRepository(db)
	useCase := usecases.NewGetBookByIdUseCase(repo)
	controller := controllers.NewGetBookByIdController(useCase)
	return controller
}
