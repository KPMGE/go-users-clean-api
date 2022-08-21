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
	service   usecases.AddBookUseCase
	validator protocols.Validator
}

func (c *AddBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
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

	err = c.validator.Validate(&inputDto)
	if err != nil {
		return helpers.BadRequest(err)
	}

	output, err := c.service.AddBook(&inputDto)
	if err != nil {
		return helpers.BadRequest(err)
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	return helpers.Ok(outputJson)
}

func NewAddBookController(service usecases.AddBookUseCase, validator protocols.Validator) *AddBookController {
	return &AddBookController{
		service:   service,
		validator: validator,
	}
}
