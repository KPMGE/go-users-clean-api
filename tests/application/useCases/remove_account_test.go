package usecases_test

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
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
	return "account deleted", nil
}

func TestRemoveAccount_WithCorectID(t *testing.T) {
	repo := mocks_test.NewFakeAccountRepository()
	sut := NewRemoveAccountUseCase(repo)

	message, err := sut.Remove(fakeAccountId)

	require.Nil(t, err)
	require.Equal(t, message, "account deleted")
}
