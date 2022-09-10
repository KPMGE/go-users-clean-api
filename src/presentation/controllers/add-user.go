package controllers

import (
	"encoding/json"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddUserController struct {
	service   usecases.AddUserUseCase
	validator protocols.Validator
}

func (c *AddUserController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	var inputUser domaindto.AddUserInputDTO
	err := json.Unmarshal(request.Body, &inputUser)
	if err != nil {
		return helpers.ServerError(err)
	}

	err = c.validator.Validate(&inputUser)
	if err != nil {
		return helpers.BadRequest(err)
	}

	newUser, err := c.service.Add(&inputUser)
	if err != nil {
		return helpers.BadRequest(err)
	}

	output := domaindto.NewAddUserOutputDTO(newUser.ID, newUser.Name, newUser.UserName, newUser.Email)
	return helpers.Ok(output)
}

func NewAddUserController(service usecases.AddUserUseCase, validator protocols.Validator) *AddUserController {
	return &AddUserController{
		service:   service,
		validator: validator,
	}
}
