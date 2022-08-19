package controllers

import (
	"encoding/json"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddAccountController struct {
	useCase usecases.AddAccountUseCase
}

func NewAddAccountController(useCase usecases.AddAccountUseCase) *AddAccountController {
	return &AddAccountController{
		useCase: useCase,
	}
}

type AddAccountRequest struct {
	Body *domaindto.AddAccountInputDTO
}

func (controller *AddAccountController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	var accountInput domaindto.AddAccountInputDTO
	err := json.Unmarshal(request.Body, &accountInput)
	if err != nil {
		return helpers.ServerError(err)
	}

	output, err := controller.useCase.AddAccount(&accountInput)
	if err != nil {
		return helpers.BadRequest(err)
	}
	outputJson, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	return helpers.Ok(outputJson)
}
