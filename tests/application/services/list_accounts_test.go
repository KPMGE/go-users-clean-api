package usecases_test

import (
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/stretchr/testify/require"
)

type ListAccountsRepository interface {
	ListAccounts() []domaindto.ListAccountsOutputDTO
}

type ListAccountsRepositoryStub struct {
	Output []domaindto.ListAccountsOutputDTO
}

type ListAccountsService struct {
	accountsRepo ListAccountsRepository
}

func NewListAccountsService(repo ListAccountsRepository) *ListAccountsService {
	return &ListAccountsService{
		accountsRepo: repo,
	}
}

func (l *ListAccountsService) ListAccounts() []domaindto.ListAccountsOutputDTO {
	return l.accountsRepo.ListAccounts()
}

func (l *ListAccountsRepositoryStub) ListAccounts() []domaindto.ListAccountsOutputDTO {
	return l.Output
}

func TestListAccounts_ShouldReturnFromRepository(t *testing.T) {
	repo := &ListAccountsRepositoryStub{
		Output: []domaindto.ListAccountsOutputDTO{},
	}
	sut := NewListAccountsService(repo)

	result := sut.ListAccounts()

	require.Equal(t, []domaindto.ListAccountsOutputDTO{}, result)
}
