package controllers_test

import (
	"errors"
	"net/http"
	"testing"

	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

const FAKE_USER_ID string = "any_user_id"

type GetUserByIdServiceMock struct {
	Output *domaindto.GetUserByIdUseCaseOutputDTO
	Error  error
}

func (g *GetUserByIdServiceMock) GetUserById(userId string) (*domaindto.GetUserByIdUseCaseOutputDTO, error) {
	return g.Output, g.Error
}

func MakeGetUserByIdControllerSut() (*controllers.GetUserByIdController, *GetUserByIdServiceMock) {
	fakeOutput := domaindto.GetUserByIdUseCaseOutputDTO{ID: "any_id", Name: "any_name", Email: "any_email@gmail.com", UserName: "any_username", Books: nil}
	serviceMock := &GetUserByIdServiceMock{Error: nil, Output: &fakeOutput}
	sut := controllers.NewGetUserByIdController(serviceMock)
	return sut, serviceMock
}

func TestGetUserByIdController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceMock := MakeGetUserByIdControllerSut()

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Output, httpResponse.Body)
}

func TestGetUserByIdController_ShouldReturnNotFoundIfServiceReturnsNull(t *testing.T) {
	sut, serviceMock := MakeGetUserByIdControllerSut()
	serviceMock.Output = nil
	expectedError := errors.New("user not found")

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusNotFound, httpResponse.StatusCode)
	require.Equal(t, expectedError.Error(), httpResponse.Body)
}

func TestGetUserByIdController_ShouldReturnServerErrorIfServiceReturnsError(t *testing.T) {
	sut, serviceMock := MakeGetUserByIdControllerSut()
	serviceMock.Error = errors.New("service error")

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusInternalServerError, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Error.Error(), httpResponse.Body)
}
