package configuration

import (
  "github.com/gofiber/fiber/v2"
  "github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"encoding/json"
	"github.com/KPMGE/go-users-clean-api/src/main/factories"
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

	return &api
}
