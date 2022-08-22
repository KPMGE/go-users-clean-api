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
	return l.accountsRepo.ListAccounts()
}
