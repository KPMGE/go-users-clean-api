package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"gorm.io/gorm"
)

func MakeListAccountsController(db *gorm.DB) protocols.Controller {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	service := services.NewListAccountsService(repo)
	controller := controllers.NewListAccountsController(service)
	return controller
}
