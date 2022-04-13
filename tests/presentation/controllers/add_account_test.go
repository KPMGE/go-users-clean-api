package controllers_test

import (
	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/infrasctructure/repositories"
	"github.com/stretchr/testify/require"
	"testing"
)

type HttpResponse struct {
	StatusCode int
	Body       interface{}
}

type AddAccountController struct {
	useCase *usecases.AddAccountUseCase
}

func NewAddAccountController(useCase *usecases.AddAccountUseCase) *AddAccountController {
	return &AddAccountController{
		useCase: useCase,
	}
}

type AddAccountRequest struct {
	Body *dto.AddAccountInputDTO
}

func ok(data interface{}) *HttpResponse {
	return &HttpResponse{
		StatusCode: 200,
		Body:       data,
	}
}

func badRequest(err error) *HttpResponse {
	return &HttpResponse{
		StatusCode: 400,
		Body:       err,
	}
}

func (controller *AddAccountController) Handle(request *AddAccountRequest) *HttpResponse {
	output, err := controller.useCase.AddAccount(request.Body)
	if err != nil {
		return badRequest(err)
	}
	return ok(output)
}

type FakeHasher struct{}

func (hasher *FakeHasher) Hash(plainText string) string {
	return "some_hash"
}

func NewFakeHasher() *FakeHasher {
	return &FakeHasher{}
}

func makeFakeRequest() *AddAccountRequest {
	fakeAccount := dto.AddAccountInputDTO{
		UserName:        "any_username",
		Email:           "any_valid_email@gmail.com",
		Password:        "any_password",
		ConfirmPassword: "any_password",
	}
	return &AddAccountRequest{
		Body: &fakeAccount,
	}
}

func makeSut() *AddAccountController {
	repo := repositories.NewInmemoryAccountRepository()
	hasher := NewFakeHasher()
	useCase := usecases.NewAddAccountUseCase(repo, hasher)
	sut := NewAddAccountController(useCase)
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
