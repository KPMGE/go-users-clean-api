package controllers_test

import (
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeSut() (*controllers.RemoveAccountController, *mocks_test.FakeAccountRepository) {
	repo := mocks_test.NewFakeAccountRepository()
	repo.DeleteAccountByIdOutput = true
	useCase := usecases.NewRemoveAccountUseCase(repo)
	sut := controllers.NewRemoveAccountController(useCase)
	return sut, repo
}

func makeFakeRemoveAccountRequest(id string) *protocols.HttpRequest {
	return protocols.NewHtppRequest(nil, []byte(id))
}

func TestRemoveAccountController_WithCorrectID(t *testing.T) {
	sut, _ := MakeSut()
	request := makeFakeRemoveAccountRequest("any_valid_id")
	httpResponse := sut.Handle(request)

	require.Equal(t, httpResponse.StatusCode, 200)
	require.Equal(t, string(httpResponse.JsonBody), "account deleted")
}

func TestRemoveAccountController_WithWrongID(t *testing.T) {
	sut, repo := MakeSut()
	repo.DeleteAccountByIdOutput = false

	request := makeFakeRemoveAccountRequest("any_invalid_id")
	httpResponse := sut.Handle(request)

	require.Equal(t, httpResponse.StatusCode, 400)
	require.Equal(t, string(httpResponse.JsonBody), "there is no account with this id")
}
