package controllers

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type DeleteUserController struct {
	service usecases.DeleteUserUseCase
}

func NewDeleteUserController(service usecases.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{
		service: service,
	}
}

func (controller *DeleteUserController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	message, err := controller.service.DeleteUser(string(request.Params))
	if err != nil {
		return helpers.BadRequest(err)
	}
	return helpers.Ok([]byte(message))
}
