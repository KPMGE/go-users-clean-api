package controllers

import (
	"encoding/json"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddUserController struct {
	useCase *usecases.AddUserUseCase
}

func (controller *AddUserController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	var inputUser dto.AddUserInputDTO
	err := json.Unmarshal(request.Body, &inputUser)
	if err != nil {
		return helpers.ServerError(err)
	}

	newUser, err := controller.useCase.Add(&inputUser)
	if err != nil {
		return helpers.BadRequest(err)
	}

	output := dto.NewAddUserOutputDTO(newUser.ID, newUser.Name, newUser.UserName, newUser.Email)
	outputJson, err := json.Marshal(output)
	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(string(outputJson))
}

func NewAddUserController(useCase *usecases.AddUserUseCase) *AddUserController {
	return &AddUserController{
		useCase: useCase,
	}
}
