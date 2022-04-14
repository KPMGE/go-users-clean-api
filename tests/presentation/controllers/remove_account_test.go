package controllers_test

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func MakeSut() (*controllers.RemoveAccountController, *mocks_test.FakeAccountRepository) {
	repo := mocks_test.NewFakeAccountRepository()
	fakeAccount, _ := entities.NewAccount("any_username", "any_valid_email@gmail.com", "any_pass")
	repo.FindAccountByIdOutput = fakeAccount
	useCase := usecases.NewRemoveAccountUseCase(repo)
	sut := controllers.NewRemoveAccountController(useCase)
	return sut, repo
}

func TestRemoveAccountController_WithCorrectID(t *testing.T) {
	sut, _ := MakeSut()
	httpResponse := sut.Handle("any_valid_id")

	require.Equal(t, httpResponse.StatusCode, 200)
	require.Equal(t, httpResponse.Body, "account deleted")
}

func TestRemoveAccountController_WithWrongID(t *testing.T) {
	sut, repo := MakeSut()
	repo.FindAccountByIdOutput = nil

	httpResponse := sut.Handle("any_invalid_id")

	require.Equal(t, httpResponse.StatusCode, 400)
	require.Equal(t, httpResponse.Body, "there is no account with this id")
}
