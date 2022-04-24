package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type RemoveBookController struct {
	useCase *usecases.RemoveBookUseCase
}

func NewRemoveBookController(useCase *usecases.RemoveBookUseCase) *RemoveBookController {
	return &RemoveBookController{
		useCase: useCase,
	}
}

func (controller *RemoveBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	removedBook, err := controller.useCase.Remove(string(request.Params))
	if err != nil {
		return helpers.BadRequest(err)
	}
	removedBookJson, err := json.Marshal(removedBook)
	if err != nil {
		return helpers.ServerError(err)
	}
	return helpers.Ok(removedBookJson)
}
