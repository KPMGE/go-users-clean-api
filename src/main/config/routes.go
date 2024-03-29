package configuration

import (
	"github.com/KPMGE/go-users-clean-api/src/main/adapters"
	"github.com/KPMGE/go-users-clean-api/src/main/factories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) *fiber.Router {
	api := app.Group("/api")

	api.Get("/accounts", adapters.FiberRouteAdapter(factories.MakeListAccountsController(db)))
	api.Post("/accounts", adapters.FiberRouteAdapter(factories.MakeAddAccountController(db)))
	api.Delete("/accounts/:accountId", adapters.FiberRouteAdapter(factories.MakeRemoveAccountController(db)))

	api.Post("/users", adapters.FiberRouteAdapter(factories.MakeAddUserController(db)))
	api.Get("/users/:userId", adapters.FiberRouteAdapter(factories.MakeGetUserByIdController(db)))
	api.Get("/users", adapters.FiberRouteAdapter(factories.MakeListUsersController(db)))
	api.Delete("/users/:userId", adapters.FiberRouteAdapter(factories.MakeDeleteUserController(db)))

	api.Get("/books", adapters.FiberRouteAdapter(factories.MakeListBooksController(db)))
	api.Get("/books/:bookId", adapters.FiberRouteAdapter(factories.MakeGetBookByIdController(db)))
	api.Post("/books", adapters.FiberRouteAdapter(factories.MakeAddBookController(db)))
	api.Delete("/books/:bookId", adapters.FiberRouteAdapter(factories.MakeRemoveBookController(db)))

	return &api
}
