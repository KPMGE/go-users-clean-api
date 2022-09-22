package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	usecases "github.com/KPMGE/go-users-clean-api/src/domain/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	presentationerrors "github.com/KPMGE/go-users-clean-api/src/presentation/presentation-errors"
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
	if request == nil || request.Body == nil {
		return helpers.BadRequest(presentationerrors.NewBlankBodyError())
	}

	var input domaindto.LoginInputDTO
	err := json.Unmarshal(request.Body, &input)

	if err != nil {
		return helpers.ServerError(errors.New("error when parsing json body!"))
	}

	err = c.validator.Validate(&input)
	if err != nil {
		return helpers.BadRequest(err)
	}

	token, _ := c.srv.Login(&input)

	return helpers.Ok(token)
}

func NewLoginController(srv usecases.LoginUseCase, validator protocols.Validator) *LoginController {
	return &LoginController{
		srv:       srv,
		validator: validator,
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
	sut := NewLoginController(serviceMock, validatorMock)
	return sut, serviceMock, validatorMock
}

func TestLoginController_ShouldReturnOkOnSuccess(t *testing.T) {
	sut, serviceMock, _ := MakeLoginControllerSut()

	httpRequest := sut.Handle(FakeLoginRequest())

	require.Equal(t, http.StatusOK, httpRequest.StatusCode)
	require.Equal(t, serviceMock.Output, httpRequest.Body)
}

func TestLoginController_ShouldReturnBadRequestIfValidatorReturnsError(t *testing.T) {
	sut, _, validatorMock := MakeLoginControllerSut()
	validatorMock.Output = errors.New("validation error")

	httpRequest := sut.Handle(FakeLoginRequest())

	require.Equal(t, http.StatusBadRequest, httpRequest.StatusCode)
	require.Equal(t, validatorMock.Output.Error(), httpRequest.Body)
}

func TestLoginController_ShouldReturnBadRequestIfABlankRequestIsProvided(t *testing.T) {
	sut, _, validatorMock := MakeLoginControllerSut()
	validatorMock.Output = errors.New("validation error")

	httpRequest := sut.Handle(protocols.NewHttpRequest(nil, nil))

	expectedError := presentationerrors.NewBlankBodyError()

	require.Equal(t, http.StatusBadRequest, httpRequest.StatusCode)
	require.Equal(t, expectedError.Error(), httpRequest.Body)
}
