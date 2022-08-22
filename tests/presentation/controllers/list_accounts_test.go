package controllers_test

import (
	"encoding/json"
	"log"
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

type ListAccountsServiceStub struct {
	Output []domaindto.ListAccountsOutputDTO
}

func (l *ListAccountsServiceStub) ListAccounts() []domaindto.ListAccountsOutputDTO {
	return l.Output
}

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

func TestListAccountsController_ShouldReturnFromService(t *testing.T) {
	service := &ListAccountsServiceStub{
		Output: []domaindto.ListAccountsOutputDTO{},
	}
	sut := NewListAccountsController(service)

	httpResponse := sut.Handle(nil)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.NotNil(t, 200, httpResponse.JsonBody)
}
