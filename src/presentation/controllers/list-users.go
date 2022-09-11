package controllers

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type ListUsersController struct {
	service usecases.ListUsersUseCase
}

func (controller *ListUsersController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	users := controller.service.List()
	return helpers.Ok(users)
}

func NewListUsersController(service usecases.ListUsersUseCase) *ListUsersController {
	return &ListUsersController{
		service: service,
	}
}
