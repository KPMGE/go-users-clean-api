package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type ListBooksController struct {
	useCase *usecases.ListBooksUseCase
}

func NewListBooksController(useCase *usecases.ListBooksUseCase) *ListBooksController {
	return &ListBooksController{
		useCase: useCase,
	}
}

func (controller *ListBooksController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	books, err := controller.useCase.List()
	if err != nil {
		return helpers.ServerError(err)
	}

	jsonBooks, err := json.Marshal(books)

	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonBooks)
}
