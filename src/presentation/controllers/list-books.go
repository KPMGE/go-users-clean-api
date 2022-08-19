package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type ListBooksController struct {
	service usecases.ListBooksUseCase
}

func (controller *ListBooksController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	books, err := controller.service.ListBooks()
	if err != nil {
		return helpers.ServerError(err)
	}

	jsonBooks, err := json.Marshal(books)

	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonBooks)
}

func NewListBooksController(service usecases.ListBooksUseCase) *ListBooksController {
	return &ListBooksController{
		service: service,
	}
}
