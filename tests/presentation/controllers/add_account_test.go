package controllers_test

import (
	"encoding/json"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
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

func makeFakeRequest(userName string, email string, password string, confirm string) *protocols.HttpRequest {
	inputObj := dto.NewAddAccountInputDTO(userName, email, password, confirm)
	jsonObj, err := json.Marshal(inputObj)
	if err != nil {
		panic(err)
	}
	return protocols.NewHtppRequest(jsonObj, nil)
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
	request := makeFakeRequest(fakeUserName, fakeEmail, fakePassword, fakePassword)
	httpResponse := sut.Handle(request)

	require.Equal(t, httpResponse.StatusCode, 200)
	require.NotNil(t, httpResponse.JsonBody)
}

func TestAddAccountController_WithInvalidData(t *testing.T) {
	sut := makeSut()

	request := makeFakeRequest(fakeUserName, "any_invalid_email", fakePassword, fakePassword)
	httpResponse := sut.Handle(request)

	require.Equal(t, httpResponse.StatusCode, 400)
}
