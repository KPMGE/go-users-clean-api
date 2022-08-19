package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type RemoveBookController struct {
	service usecases.RemoveBookUseCase
}

func NewRemoveBookController(service usecases.RemoveBookUseCase) *RemoveBookController {
	return &RemoveBookController{
		service: service,
	}
}

func (controller *RemoveBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	removedBook, err := controller.service.RemoveBook(string(request.Params))
	if err != nil {
		return helpers.BadRequest(err)
	}
	removedBookJson, err := json.Marshal(removedBook)
	if err != nil {
		return helpers.ServerError(err)
	}
	return helpers.Ok(removedBookJson)
}
