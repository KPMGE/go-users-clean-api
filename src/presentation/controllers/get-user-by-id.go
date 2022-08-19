package controllers

import (
	"encoding/json"
	"errors"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type GetUserByIdController struct {
	service usecases.GetUserByIdUseCase
}

func NewGetUserByIdController(s usecases.GetUserByIdUseCase) *GetUserByIdController {
	return &GetUserByIdController{
		service: s,
	}
}

func (controller *GetUserByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	if len(request.Params) == 0 {
		err := errors.New("Blank userId!")
		return helpers.BadRequest(err)
	}

	foundUser, err := controller.service.GetUserById(string(request.Params))
	if err != nil {
		return helpers.ServerError(err)
	}

	outputDto := domaindto.NewGetUserByIdUseCaseOutputDTO(
		foundUser.ID,
		foundUser.Name,
		foundUser.Email,
		foundUser.UserName,
		foundUser.Books)

	jsonOutputDto, err := json.Marshal(outputDto)

	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonOutputDto)
}
