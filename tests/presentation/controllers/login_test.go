package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	presentationerrors "github.com/KPMGE/go-users-clean-api/src/presentation/presentation-errors"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	fakedtos "github.com/KPMGE/go-users-clean-api/tests/domain/fake-dtos"
	controllermocks_test "github.com/KPMGE/go-users-clean-api/tests/presentation/controller-mocks"
	"github.com/stretchr/testify/require"
)

func FakeLoginRequest() *protocols.HttpRequest {
	inputJson, err := json.Marshal(fakedtos.MakeFakeLoginInputDTO())
	if err != nil {
		panic("json Marshal fail at: FakeLoginRequest")
	}
	return protocols.NewHttpRequest(inputJson, nil)
}

func MakeLoginControllerSut() (*controllers.LoginController, *controllermocks_test.LoginServiceMock, *controllermocks_test.ValidatorMock) {
	serviceMock := controllermocks_test.NewLoginServiceMock()
	validatorMock := &controllermocks_test.ValidatorMock{Output: nil}
	sut := controllers.NewLoginController(serviceMock, validatorMock)
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
