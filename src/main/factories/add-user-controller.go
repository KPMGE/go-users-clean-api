package factories

import (
	"database/sql"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
)

func MakeAddUserController(db *sql.DB) *controllers.AddUserController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	useCase := usecases.NewAddUserUseCase(repo)
	controller := controllers.NewAddUserController(useCase)
	return controller
}
