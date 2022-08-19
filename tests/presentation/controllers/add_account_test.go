package controllers_test

import (
	"encoding/json"
	"github.com/KPMGE/go-users-clean-api/src/application/services"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
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
