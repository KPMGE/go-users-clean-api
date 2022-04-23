package controllers

import (
	"encoding/json"
	"errors"
	"log"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddBookController struct {
	useCase *usecases.AddBookUseCase
}

func NewAddBookController(useCase *usecases.AddBookUseCase) *AddBookController {
	return &AddBookController{
		useCase: useCase,
	}
}

func (controller *AddBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	if request == nil {
		newError := errors.New("Invalid body!")
		return helpers.ServerError(newError)
	}

	var inputDto dto.AddBookUseCaseInputDTO
	err := json.Unmarshal(request.Body, &inputDto)
	if err != nil {
		newError := errors.New("Invalid body!")
		return helpers.ServerError(newError)
	}

	output, err := controller.useCase.Add(&inputDto)
	if err != nil {
		return helpers.BadRequest(err)
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	return helpers.Ok(outputJson)
}
