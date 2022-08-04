package factories

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeRemoveBookController(db *gorm.DB) *controllers.RemoveBookController {
	repo := postgresrepository.NewPostgresBookRepository(db)
	userRepo := repositories.NewInMemoryUserRepository()
	useCase := usecases.NewRemoveBookUseCase(repo, repo, userRepo)
	controller := controllers.NewRemoveBookController(useCase)
	return controller
}
