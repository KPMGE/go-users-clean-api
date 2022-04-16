package configuration

import (
	"github.com/KPMGE/go-users-clean-api/src/main/factories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")

	api.Post("/accounts", func(c *fiber.Ctx) error {
		controller := factories.MakeAddAccountController()
		request := protocols.NewHtppRequest(c.Body(), nil)
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).Send(httpResponse.JsonBody)
	})

	api.Delete("/accounts/:accountId", func(c *fiber.Ctx) error {
		accountId := c.Params("accountId")
		request := protocols.NewHtppRequest(nil, []byte(accountId))
		controller := factories.MakeRemoveAccountController()
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).Send(httpResponse.JsonBody)
	})

	api.Post("/users", func(c *fiber.Ctx) error {
		controller := factories.MakeAddUserController()
		request := protocols.NewHtppRequest(c.Body(), nil)
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).Send(httpResponse.JsonBody)
	})

	api.Get("/users", func(c *fiber.Ctx) error {
		controller := factories.MakeListUsersController()
		request := protocols.NewHtppRequest(nil, nil)
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).Send(httpResponse.JsonBody)
	})

	api.Delete("/users/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		request := protocols.NewHtppRequest(nil, []byte(userId))
		controller := factories.MakeDeleteUserController()
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).Send(httpResponse.JsonBody)
	})

	return &api
}
