package controllers

import (
	"encoding/json"
	"log"

	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
)

type ListAccountsController struct {
	service usecases.ListAccountsUseCase
}

func NewListAccountsController(service usecases.ListAccountsUseCase) *ListAccountsController {
	return &ListAccountsController{
		service: service,
	}
}

func (c *ListAccountsController) Handle(req *protocols.HttpRequest) *protocols.HttpResponse {
	accounts := c.service.ListAccounts()
	accountsJson, err := json.Marshal(&accounts)

	if err != nil {
		log.Fatalln(err)
	}

	return helpers.Ok(accountsJson)
}
