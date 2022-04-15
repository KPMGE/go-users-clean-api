package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type ListUsersController struct {
	useCase *usecases.ListUsersUseCase
}

func NewListUsersController(useCase *usecases.ListUsersUseCase) *ListUsersController {
	return &ListUsersController{
		useCase: useCase,
	}
}

func (controller *ListUsersController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	users := controller.useCase.List()
	jsonUsers, err := json.Marshal(users)

	if err != nil {
		return helpers.ServerError(err)
	}
	return helpers.Ok(jsonUsers)
}
