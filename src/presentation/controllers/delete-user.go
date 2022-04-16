package controllers

import (
  "github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
)

type DeleteUserController struct {
	useCase *usecases.DeleteUserUseCase
}

func NewDeleteUserController(useCase *usecases.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		useCase: useCase,
	}
}

func (controller *DeleteUserController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	message, err := controller.useCase.Delete(string(request.Params))
	if err != nil {
		return helpers.BadRequest(err)
	}
	return helpers.Ok([]byte(message))
}
