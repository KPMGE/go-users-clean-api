package factories

import (
	"database/sql"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeRemoveAccountController(db *sql.DB) *controllers.RemoveAccountController {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	useCase := usecases.NewRemoveAccountUseCase(repo)
	controller := controllers.NewRemoveAccountController(useCase)
	return controller
}
