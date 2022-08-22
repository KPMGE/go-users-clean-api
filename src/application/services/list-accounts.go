package services

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
)

type ListAccountsService struct {
	accountsRepo protocols.ListAccountsRepository
}

func NewListAccountsService(repo protocols.ListAccountsRepository) *ListAccountsService {
	return &ListAccountsService{
		accountsRepo: repo,
	}
}

func (l *ListAccountsService) ListAccounts() []domaindto.ListAccountsOutputDTO {
	foundAccounts := l.accountsRepo.ListAccounts()

	accountsDto := []domaindto.ListAccountsOutputDTO{}

	for _, account := range foundAccounts {
		accountsDto = append(accountsDto, domaindto.ListAccountsOutputDTO{
			UserName:  account.UserName,
			Email:     account.Email,
			ID:        account.ID,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		})
	}

	return accountsDto
}
