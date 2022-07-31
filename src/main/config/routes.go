package configuration

import (
	"database/sql"

	"github.com/KPMGE/go-users-clean-api/src/main/adapters"
	"github.com/KPMGE/go-users-clean-api/src/main/factories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) *fiber.Router {
	api := app.Group("/api")

	api.Post("/accounts", adapters.FiberRouteAdapter(factories.MakeAddAccountController(db)))
	api.Delete("/accounts/:accountId", adapters.FiberRouteAdapter(factories.MakeRemoveAccountController(db)))

	api.Post("/users", adapters.FiberRouteAdapter(factories.MakeAddUserController(db)))
	api.Get("/users/:userId", adapters.FiberRouteAdapter(factories.MakeGetUserByIdController()))
	api.Get("/users", adapters.FiberRouteAdapter(factories.MakeListUsersController(db)))
	api.Delete("/users/:userId", adapters.FiberRouteAdapter(factories.MakeDeleteUserController()))

	api.Get("/books", adapters.FiberRouteAdapter(factories.MakeListBooksController()))
	api.Get("/books/:bookId", adapters.FiberRouteAdapter(factories.MakeGetBookByIdController()))
	api.Post("/books", adapters.FiberRouteAdapter(factories.MakeAddBookController(db)))
	api.Delete("/books/:bookId", adapters.FiberRouteAdapter(factories.MakeRemoveBookController()))

	return &api
}
