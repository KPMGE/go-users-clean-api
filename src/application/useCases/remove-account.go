package usecases

import (
	"errors"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
)

type RemoveAccountUseCase struct {
	accountRepository protocols.AccountRepository
}

func (useCase *RemoveAccountUseCase) Remove(accountId string) (string, error) {
	foundAccount := useCase.accountRepository.DeleteAccountById(accountId)
	if !foundAccount {
		return "", errors.New("there is no account with this id")
	}
	return "account deleted", nil
}

func NewRemoveAccountUseCase(repo protocols.AccountRepository) *RemoveAccountUseCase {
	return &RemoveAccountUseCase{
		accountRepository: repo,
	}
}
