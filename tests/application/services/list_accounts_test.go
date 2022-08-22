package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/stretchr/testify/require"
)

type ListAccountsRepositoryStub struct {
	Output []domaindto.ListAccountsOutputDTO
}

func (l *ListAccountsRepositoryStub) ListAccounts() []domaindto.ListAccountsOutputDTO {
	return l.Output
}

func TestListAccounts_ShouldReturnFromRepository(t *testing.T) {
	repo := &ListAccountsRepositoryStub{
		Output: []domaindto.ListAccountsOutputDTO{},
	}
	sut := services.NewListAccountsService(repo)

	result := sut.ListAccounts()

	require.Equal(t, []domaindto.ListAccountsOutputDTO{}, result)
}
