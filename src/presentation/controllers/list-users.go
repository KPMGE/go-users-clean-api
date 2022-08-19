package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type ListUsersController struct {
	service usecases.ListUsersUseCase
}

func (controller *ListUsersController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	users := controller.service.List()
	jsonUsers, err := json.Marshal(users)

	if err != nil {
		return helpers.ServerError(err)
	}
	return helpers.Ok(jsonUsers)
}

func NewListUsersController(service usecases.ListUsersUseCase) *ListUsersController {
	return &ListUsersController{
		service: service,
	}
}
