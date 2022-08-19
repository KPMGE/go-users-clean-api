package factories

import (
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	postgresrepository "github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories/postgres-repository"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"gorm.io/gorm"
)

func MakeListBooksController(db *gorm.DB) *controllers.ListBooksController {
	repo := postgresrepository.NewPostgresBookRepository(db)
	services := services.NewListBookService(repo)
	controller := controllers.NewListBooksController(services)
	return controller
}
