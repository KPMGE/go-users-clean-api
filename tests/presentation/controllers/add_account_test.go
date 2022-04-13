package controllers_test

import (
	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/stretchr/testify/require"
	"testing"
)

type FakeHasher struct{}

func (hasher *FakeHasher) Hash(plainText string) string {
	return "some_hash"
}

func NewFakeHasher() *FakeHasher {
	return &FakeHasher{}
}

func makeFakeRequest() *controllers.AddAccountRequest {
	fakeAccount := dto.AddAccountInputDTO{
		UserName:        "any_username",
		Email:           "any_valid_email@gmail.com",
		Password:        "any_password",
		ConfirmPassword: "any_password",
	}
	return &controllers.AddAccountRequest{
		Body: &fakeAccount,
	}
}

func makeSut() *controllers.AddAccountController {
	repo := repositories.NewInmemoryAccountRepository()
	hasher := NewFakeHasher()
	useCase := usecases.NewAddAccountUseCase(repo, hasher)
	sut := controllers.NewAddAccountController(useCase)
	return sut
}

func TestAddAccountController_WithRightData(t *testing.T) {
	sut := makeSut()
	request := makeFakeRequest()
	httpResponse := sut.Handle(request)

	require.Equal(t, httpResponse.StatusCode, 200)
	require.NotNil(t, httpResponse.Body)
}

func TestAddAccountController_WithInvalidData(t *testing.T) {
	sut := makeSut()

	request := makeFakeRequest()
	request.Body.Email = "invalid_email"
	httpResponse := sut.Handle(request)

	require.Equal(t, httpResponse.StatusCode, 400)
}
