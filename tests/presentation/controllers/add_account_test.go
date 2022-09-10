package controllers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	controllermocks_test "github.com/KPMGE/go-users-clean-api/tests/presentation/controller-mocks"
	"github.com/stretchr/testify/require"
)

type FakeHasher struct{}

func (hasher *FakeHasher) Hash(plainText string) string {
	return "some_hash"
}

func NewFakeHasher() *FakeHasher {
	return &FakeHasher{}
}

const fakePassword string = "any_password"

func makeSut() (*controllers.AddAccountController, *controllermocks_test.ValidatorMock) {
	repo := repositories.NewInmemoryAccountRepository()
	hasher := NewFakeHasher()
	validatorMock := &controllermocks_test.ValidatorMock{Output: nil}
	useCase := services.NewAddAccountService(repo, hasher)
	sut := controllers.NewAddAccountController(useCase, validatorMock)
	return sut, validatorMock
}

func TestAddAccountController_WithRightData(t *testing.T) {
	sut, _ := makeSut()

	fakeInput := `{ 
		"userName": "any_username",
		"email": "any_email@gmail.com",
		"password": "any_password",
		"confirmPassword": "any_password"
  }
  `

	fakeRequest := protocols.NewHttpRequest([]byte(fakeInput), nil)
	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.NotNil(t, httpResponse.Body)
}

func TestAddAccountController_ShouldReturnBadRequestValidatorReturnsError(t *testing.T) {
	sut, validatorMock := makeSut()
	validatorMock.Output = errors.New("validation error")

	fakeInput := `{ 
		"userName": "any_username",
		"email": "any_email@gmail.com",
		"password": "any_password",
		"confirmPassword": "any_password"
  }
  `

	fakeRequest := protocols.NewHttpRequest([]byte(fakeInput), nil)

	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
	require.Equal(t, validatorMock.Output.Error(), httpResponse.Body)
}
