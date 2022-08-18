package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeRemoveAccountController(db *gorm.DB) *controllers.RemoveAccountController {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	useCase := services.NewRemoveAccountService(repo)
	controller := controllers.NewRemoveAccountController(useCase)
	return controller
}
