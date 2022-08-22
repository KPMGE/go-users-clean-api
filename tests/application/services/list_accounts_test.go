package usecases_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type ListAccountsOutputDTO struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt" valid:"-"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
}

type ListAccountsUseCase interface {
	ListAccounts() []ListAccountsOutputDTO
}

type ListAccountsRepository interface {
	ListAccounts() []ListAccountsOutputDTO
}

type ListAccountsRepositoryStub struct {
	Output []ListAccountsOutputDTO
}

type ListAccountsService struct {
	accountsRepo ListAccountsRepository
}

func NewListAccountsService(repo ListAccountsRepository) *ListAccountsService {
	return &ListAccountsService{
		accountsRepo: repo,
	}
}

func (l *ListAccountsService) ListAccounts() []ListAccountsOutputDTO {
	return l.accountsRepo.ListAccounts()
}

func (l *ListAccountsRepositoryStub) ListAccounts() []ListAccountsOutputDTO {
	return l.Output
}

func TestListAccounts_ShouldReturnFromRepository(t *testing.T) {
	repo := &ListAccountsRepositoryStub{
		Output: []ListAccountsOutputDTO{},
	}
	sut := NewListAccountsService(repo)

	result := sut.ListAccounts()

	require.Equal(t, []ListAccountsOutputDTO{}, result)
}
