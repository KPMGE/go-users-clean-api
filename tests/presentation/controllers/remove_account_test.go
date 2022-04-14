package controllers_test

import (
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

type RemoveAccountController struct {
	useCase *usecases.RemoveAccountUseCase
}

func (controller *RemoveAccountController) Handle(accountId string) *protocols.HttpResponse {
	return helpers.Ok("account deleted")
}

func NewRemoveAccountController(useCase *usecases.RemoveAccountUseCase) *RemoveAccountController {
	return &RemoveAccountController{
		useCase: useCase,
	}
}

func TestRemoveAccountController_WithCorrectID(t *testing.T) {
	repo := mocks_test.NewFakeAccountRepository()
	useCase := usecases.NewRemoveAccountUseCase(repo)
	sut := NewRemoveAccountController(useCase)

	httpResponse := sut.Handle("any_valid_id")

	require.Equal(t, httpResponse.StatusCode, 200)
	require.Equal(t, httpResponse.Body, "account deleted")
}
