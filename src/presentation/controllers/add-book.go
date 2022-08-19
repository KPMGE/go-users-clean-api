package controllers

import (
	"encoding/json"
	"errors"
	"log"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type AddBookController struct {
	service usecases.AddBookUseCase
}

func (controller *AddBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	if request == nil {
		newError := errors.New("Invalid body!")
		return helpers.ServerError(newError)
	}

	var inputDto domaindto.AddBookUseCaseInputDTO
	err := json.Unmarshal(request.Body, &inputDto)
	if err != nil {
		newError := errors.New("Invalid body!")
		return helpers.ServerError(newError)
	}

	output, err := controller.service.AddBook(&inputDto)
	if err != nil {
		return helpers.BadRequest(err)
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	return helpers.Ok(outputJson)
}

func NewAddBookController(service usecases.AddBookUseCase) *AddBookController {
	return &AddBookController{
		service: service,
	}
}
