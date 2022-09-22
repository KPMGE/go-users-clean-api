package controllers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	fakedtos "github.com/KPMGE/go-users-clean-api/tests/domain/fake-dtos"
	controllermocks_test "github.com/KPMGE/go-users-clean-api/tests/presentation/controller-mocks"
	"github.com/stretchr/testify/require"
)

type LoginServiceMock struct {
	Input  *domaindto.LoginInputDTO
	Output string
	Error  error
}

func (s *LoginServiceMock) Login(input *domaindto.LoginInputDTO) (string, error) {
	s.Input = input
	return s.Output, s.Error
}

func NewLoginServiceMock() *LoginServiceMock {
	return &LoginServiceMock{
		Input:  nil,
		Output: "token",
		Error:  nil,
	}
}

type LoginController struct {
	srv       usecases.LoginUseCase
	validator protocols.Validator
}

func (c *LoginController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	var input domaindto.LoginInputDTO
	json.Unmarshal(request.Body, &input)

	token, _ := c.srv.Login(&input)

	return helpers.Ok(token)
}

func NewLoginController(srv usecases.LoginUseCase) *LoginController {
	return &LoginController{
		srv: srv,
	}
}

func FakeLoginRequest() *protocols.HttpRequest {
	inputJson, err := json.Marshal(fakedtos.MakeFakeLoginInputDTO())
	if err != nil {
		panic("json Marshal fail at: FakeLoginRequest")
	}
	return protocols.NewHttpRequest(inputJson, nil)
}

func MakeLoginControllerSut() (*LoginController, *LoginServiceMock, *controllermocks_test.ValidatorMock) {
	serviceMock := NewLoginServiceMock()
	validatorMock := &controllermocks_test.ValidatorMock{Output: nil}
	sut := NewLoginController(serviceMock)
	return sut, serviceMock, validatorMock
}

func TestLoginController_ShouldReturnOkOnSuccess(t *testing.T) {
	sut, serviceMock, _ := MakeLoginControllerSut()

	httpRequest := sut.Handle(FakeLoginRequest())

	require.Equal(t, http.StatusOK, httpRequest.StatusCode)
	require.Equal(t, serviceMock.Output, httpRequest.Body)
}
