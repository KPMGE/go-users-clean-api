package factories

import (
	"database/sql"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/providers"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeAddAccountController(db *sql.DB) *controllers.AddAccountController {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	hasher := providers.NewBcryptHasher()
	useCase := usecases.NewAddAccountUseCase(repo, hasher)
	controller := controllers.NewAddAccountController(useCase)
	return controller
}
