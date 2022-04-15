package usecases_test

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

const fakeAccountId string = "any_valid_account_id"

func MakeSut() (*usecases.RemoveAccountUseCase, *mocks_test.FakeAccountRepository) {
	repo := mocks_test.NewFakeAccountRepository()
	repo.DeleteAccountByIdOutput = true
	sut := usecases.NewRemoveAccountUseCase(repo)
	return sut, repo
}

func TestRemoveAccount_WithCorectID(t *testing.T) {
	sut, _ := MakeSut()
	message, err := sut.Remove(fakeAccountId)

	require.Nil(t, err)
	require.Equal(t, message, "account deleted")
}

func TestRemoveAccount_WithIncorrectID(t *testing.T) {
	sut, repo := MakeSut()
	repo.DeleteAccountByIdOutput = false

	message, err := sut.Remove(fakeAccountId)

	require.Error(t, err)
	require.Equal(t, message, "")
	require.Equal(t, err.Error(), "there is no account with this id")
}
