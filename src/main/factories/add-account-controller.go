package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/providers"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeAddAccountController(db *gorm.DB) *controllers.AddAccountController {
	repo := postgresrepository.NewPostgresAccountRepository(db)
	hasher := providers.NewBcryptHasher()
	service := services.NewAddAccountService(repo, hasher)
	controller := controllers.NewAddAccountController(service)
	return controller
}
