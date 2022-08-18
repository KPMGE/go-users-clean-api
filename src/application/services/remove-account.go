package services

import (
	"errors"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
)

type RemoveAccountService struct {
	accountRepository protocols.AccountRepository
}

func (useCase *RemoveAccountService) RemoveAccount(accountId string) (string, error) {
	foundAccount := useCase.accountRepository.DeleteAccountById(accountId)
	if !foundAccount {
		return "", errors.New("there is no account with this id")
	}
	return "account deleted", nil
}

func NewRemoveAccountService(repo protocols.AccountRepository) *RemoveAccountService {
	return &RemoveAccountService{
		accountRepository: repo,
	}
}
