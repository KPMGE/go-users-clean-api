package controllers

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type RemoveAccountController struct {
	useCase usecases.RemoveAccountUseCase
}

func (controller *RemoveAccountController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	message, err := controller.useCase.RemoveAccount(string(request.Params))
	if err != nil {
		return helpers.BadRequest(err)
	}
	return helpers.Ok(message)
}

func NewRemoveAccountController(useCase usecases.RemoveAccountUseCase) *RemoveAccountController {
	return &RemoveAccountController{
		useCase: useCase,
	}
}
