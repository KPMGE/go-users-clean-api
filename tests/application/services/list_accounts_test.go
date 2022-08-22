package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type ListAccountsRepositoryStub struct {
	Output []entities.Account
}

func (l *ListAccountsRepositoryStub) ListAccounts() []entities.Account {
	return l.Output
}

func TestListAccounts_ShouldReturnFromRepository(t *testing.T) {
	repo := &ListAccountsRepositoryStub{
		Output: []entities.Account{},
	}
	sut := services.NewListAccountsService(repo)

	result := sut.ListAccounts()

	require.Equal(t, []domaindto.ListAccountsOutputDTO{}, result)
}
