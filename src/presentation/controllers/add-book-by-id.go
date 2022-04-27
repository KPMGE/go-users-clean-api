package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type GetBookByIdController struct {
	useCase *usecases.GetBookByIdUseCase
}

func NewGetBookByIdController(useCase *usecases.GetBookByIdUseCase) *GetBookByIdController {
	return &GetBookByIdController{
		useCase: useCase,
	}
}

func (controller *GetBookByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	book, err := controller.useCase.GetById(string(request.Params))
	if err != nil {
		return helpers.ServerError(err)
	}

	jsonBook, err := json.Marshal(book)
	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonBook)
}
