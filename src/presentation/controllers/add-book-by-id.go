package controllers

import (
	"encoding/json"

	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type GetBookByIdController struct {
	service usecases.GetBookByIdUseCase
}

func NewGetBookByIdController(service usecases.GetBookByIdUseCase) *GetBookByIdController {
	return &GetBookByIdController{
		service: service,
	}
}

func (controller *GetBookByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	book, err := controller.service.GetBookById(string(request.Params))
	if err != nil {
		return helpers.ServerError(err)
	}

	jsonBook, err := json.Marshal(book)
	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonBook)
}
