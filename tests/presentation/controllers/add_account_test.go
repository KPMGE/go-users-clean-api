package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/main/factories/validators"
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
	inputObj := domaindto.NewAddAccountInputDTO(userName, email, password, confirm)
	jsonObj, err := json.Marshal(inputObj)
	if err != nil {
		panic(err)
	}
	return protocols.NewHtppRequest(jsonObj, nil)
}

func makeSut() *controllers.AddAccountController {
	repo := repositories.NewInmemoryAccountRepository()
	hasher := NewFakeHasher()
	useCase := services.NewAddAccountService(repo, hasher)
	sut := controllers.NewAddAccountController(useCase, validators.MakeAddAccountValidation())
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

func TestAddAccountController_ShouldReturnBadRequestIfFieldsAreMissing(t *testing.T) {
	sut := makeSut()

	request := makeFakeRequest("", "any_valid_email@gmail.com", fakePassword, fakePassword)
	httpResponse := sut.Handle(request)
	require.Equal(t, httpResponse.StatusCode, 400)
	require.Equal(t, string(httpResponse.JsonBody), "Missing field UserName!")

	request = makeFakeRequest("any username", "", fakePassword, fakePassword)
	httpResponse = sut.Handle(request)
	require.Equal(t, httpResponse.StatusCode, 400)
	require.Equal(t, string(httpResponse.JsonBody), "Missing field Email!")

	request = makeFakeRequest("any username", "any_valid_email@gmail.com", fakePassword, "")
	httpResponse = sut.Handle(request)
	require.Equal(t, httpResponse.StatusCode, 400)
	require.Equal(t, string(httpResponse.JsonBody), "Missing field ConfirmPassword!")

	request = makeFakeRequest("any username", "any_valid_email@gmail.com", "", fakePassword)
	httpResponse = sut.Handle(request)
	require.Equal(t, httpResponse.StatusCode, 400)
	require.Equal(t, string(httpResponse.JsonBody), "Missing field Password!")
}
