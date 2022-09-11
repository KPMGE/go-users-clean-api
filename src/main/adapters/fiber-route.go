package adapters

import (
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/gofiber/fiber/v2"
)

func FiberRouteAdapter(controller protocols.Controller) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		params := c.Route().Params
		if params == nil {
			request := protocols.NewHttpRequest(c.Body(), nil)
			httpResponse := controller.Handle(request)
			return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
		}
		paramName := params[0]
		paramValue := c.Params(paramName)
		request := protocols.NewHttpRequest(c.Body(), []byte(paramValue))
		httpResponse := controller.Handle(request)

		if httpResponse.StatusCode < 299 {
			return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
		}
		return c.Status(httpResponse.StatusCode).JSON(map[string]any{
			"error": httpResponse.Body,
		})
	}
}
