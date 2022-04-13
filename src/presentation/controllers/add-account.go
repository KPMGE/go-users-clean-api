package controllers

import (
	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddAccountController struct {
	useCase *usecases.AddAccountUseCase
}

func NewAddAccountController(useCase *usecases.AddAccountUseCase) *AddAccountController {
	return &AddAccountController{
		useCase: useCase,
	}
}

type AddAccountRequest struct {
	Body *dto.AddAccountInputDTO
}

func (controller *AddAccountController) Handle(request *AddAccountRequest) *protocols.HttpResponse {
	output, err := controller.useCase.AddAccount(request.Body)
	if err != nil {
		return helpers.BadRequest(err)
	}
	return helpers.Ok(output)
}
