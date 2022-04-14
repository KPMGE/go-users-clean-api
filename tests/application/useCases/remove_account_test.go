package usecases_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const fakeAccountId string = "any_valid_account_id"

type RemoveAccountUseCase struct {
	accountRepository protocols.AccountRepository
}

func NewRemoveAccountUseCase(repo protocols.AccountRepository) *RemoveAccountUseCase {
	return &RemoveAccountUseCase{
		accountRepository: repo,
	}
}

func (useCase *RemoveAccountUseCase) Remove(accountId string) (string, error) {
	foundAccount := useCase.accountRepository.FindAccountById(accountId)
	if foundAccount == nil {
		return "", errors.New("there is no account with this id")
	}
	return "account deleted", nil
}

func TestRemoveAccount_WithCorectID(t *testing.T) {
	repo := mocks_test.NewFakeAccountRepository()
	sut := NewRemoveAccountUseCase(repo)
	fakeAccont, _ := entities.NewAccount("any_username", "any_valid_email@gmail.com", "any_password")
	repo.FindAccountByIdOutput = fakeAccont

	message, err := sut.Remove(fakeAccountId)

	require.Nil(t, err)
	require.Equal(t, message, "account deleted")
}

func TestRemoveAccount_WithIncorrectID(t *testing.T) {
	repo := mocks_test.NewFakeAccountRepository()
	sut := NewRemoveAccountUseCase(repo)
	repo.FindAccountByIdOutput = nil

	message, err := sut.Remove(fakeAccountId)

	require.Error(t, err)
	require.Equal(t, message, "")
	require.Equal(t, err.Error(), "there is no account with this id")
}
