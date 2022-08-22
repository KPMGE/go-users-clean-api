package controllers_test

import (
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/stretchr/testify/require"
)

type ListAccountsServiceStub struct {
	Output []domaindto.ListAccountsOutputDTO
}

func (l *ListAccountsServiceStub) ListAccounts() []domaindto.ListAccountsOutputDTO {
	return l.Output
}

func TestListAccountsController_ShouldReturnFromService(t *testing.T) {
	service := &ListAccountsServiceStub{
		Output: []domaindto.ListAccountsOutputDTO{},
	}
	sut := controllers.NewListAccountsController(service)

	httpResponse := sut.Handle(nil)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.NotNil(t, 200, httpResponse.JsonBody)
}
