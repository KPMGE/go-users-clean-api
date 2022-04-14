package configuration

import (
	"encoding/json"
	"github.com/KPMGE/go-users-clean-api/src/main/factories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")

	api.Post("/accounts", func(c *fiber.Ctx) error {
		var accountInput controllers.AddAccountRequest

		err := json.Unmarshal(c.Body(), &accountInput.Body)
		if err != nil {
			panic(err)
		}

		controller := factories.MakeAddAccountController()
		httpResponse := controller.Handle(&accountInput)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	})

	api.Delete("/accounts/:accountId", func(c *fiber.Ctx) error {
		accountId := c.Params("accountId")
		controller := factories.MakeRemoveAccountController()
		httpResponse := controller.Handle(accountId)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	})

	return &api
}
