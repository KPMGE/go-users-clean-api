package controllers

import (
	"encoding/json"
	"errors"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type GetUserByIdController struct {
	useCase *usecases.GetUserByIdUseCase
}

func NewGetUserByIdController(useCase *usecases.GetUserByIdUseCase) *GetUserByIdController {
	return &GetUserByIdController{
		useCase: useCase,
	}
}

func (controller *GetUserByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	if len(request.Params) == 0 {
		err := errors.New("Blank userId!")
		return helpers.BadRequest(err)
	}

	foundUser, err := controller.useCase.Get(string(request.Params))
	if err != nil {
		return helpers.ServerError(err)
	}

	outputDto := dto.NewGetUserByIdUseCaseOutputDTO(foundUser.ID, foundUser.Name, foundUser.Email, foundUser.UserName, foundUser.Books)
	jsonOutputDto, err := json.Marshal(outputDto)
	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonOutputDto)
}
