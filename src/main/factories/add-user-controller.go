package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/main/factories/validators"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeAddUserController(db *gorm.DB) *controllers.AddUserController {
	repo := postgresrepository.NewPostgresUserRepository(db)
	service := services.NewAddUserService(repo)
	validator := validators.MakeAddUserValidation()
	controller := controllers.NewAddUserController(service, validator)
	return controller
}
