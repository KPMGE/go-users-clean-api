package controllers

import (
	"encoding/json"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddAccountController struct {
	useCase   usecases.AddAccountUseCase
	validator protocols.Validator
}

func NewAddAccountController(useCase usecases.AddAccountUseCase, validator protocols.Validator) *AddAccountController {
	return &AddAccountController{
		useCase:   useCase,
		validator: validator,
	}
}

type AddAccountRequest struct {
	Body *domaindto.AddAccountInputDTO
}

func (c *AddAccountController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	var accountInput domaindto.AddAccountInputDTO
	err := json.Unmarshal(request.Body, &accountInput)
	if err != nil {
		return helpers.ServerError(err)
	}

	err = c.validator.Validate(&accountInput)
	if err != nil {
		return helpers.BadRequest(err)
	}

	output, err := c.useCase.AddAccount(&accountInput)
	if err != nil {
		return helpers.BadRequest(err)
	}
	outputJson, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	return helpers.Ok(outputJson)
}
