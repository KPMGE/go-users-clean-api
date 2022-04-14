package controllers

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type RemoveAccountController struct {
	useCase *usecases.RemoveAccountUseCase
}

func (controller *RemoveAccountController) Handle(accountId string) *protocols.HttpResponse {
	message, err := controller.useCase.Remove(accountId)
	if err != nil {
		return helpers.BadRequest(err)
	}
	return helpers.Ok(message)
}

func NewRemoveAccountController(useCase *usecases.RemoveAccountUseCase) *RemoveAccountController {
	return &RemoveAccountController{
		useCase: useCase,
	}
}
