package configuration

import (
	"github.com/KPMGE/go-users-clean-api/src/main/adapters"
	"github.com/KPMGE/go-users-clean-api/src/main/factories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")

	api.Post("/accounts", adapters.FiberRouteAdapter(factories.MakeAddAccountController()))
	api.Delete("/accounts/:accountId", adapters.FiberRouteAdapter(factories.MakeRemoveAccountController()))

	api.Post("/users", adapters.FiberRouteAdapter(factories.MakeAddUserController()))
	api.Get("/users/:userId", adapters.FiberRouteAdapter(factories.MakeGetUserByIdController()))
	api.Get("/users", adapters.FiberRouteAdapter(factories.MakeListUsersController()))
	api.Delete("/users/:userId", adapters.FiberRouteAdapter(factories.MakeDeleteUserController()))

	api.Get("/books", adapters.FiberRouteAdapter(factories.MakeListBooksController()))
	api.Get("/books/:bookId", adapters.FiberRouteAdapter(factories.MakeGetBookByIdController()))
	api.Post("/books", adapters.FiberRouteAdapter(factories.MakeAddBookController()))
	api.Delete("/books/:bookId", adapters.FiberRouteAdapter(factories.MakeRemoveBookController()))

	return &api
}
