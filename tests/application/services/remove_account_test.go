package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

const fakeAccountId string = "any_valid_account_id"

func MakeSut() (*services.RemoveAccountService, *mocks_test.FakeAccountRepository) {
	repo := mocks_test.NewFakeAccountRepository()
	repo.DeleteAccountByIdOutput = true
	sut := services.NewRemoveAccountService(repo)
	return sut, repo
}

func TestRemoveAccount_WithCorectID(t *testing.T) {
	sut, _ := MakeSut()
	message, err := sut.RemoveAccount(fakeAccountId)

	require.Nil(t, err)
	require.Equal(t, message, "account deleted")
}

func TestRemoveAccount_WithIncorrectID(t *testing.T) {
	sut, repo := MakeSut()
	repo.DeleteAccountByIdOutput = false

	message, err := sut.RemoveAccount(fakeAccountId)

	require.Error(t, err)
	require.Equal(t, message, "")
	require.Equal(t, err.Error(), "there is no account with this id")
}
