package factories

import (
	"database/sql"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeListUsersController(db *sql.DB) *controllers.ListUsersController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	useCase := usecases.NewListUsersUseCase(repo)
	controller := controllers.NewListUsersController(useCase)
	return controller
}
