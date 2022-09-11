package controllers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

type DeleteUserServiceMock struct {
	Output string
	Error  error
}

func (s *DeleteUserServiceMock) DeleteUser(userId string) (string, error) {
	return s.Output, s.Error
}

func MakeDeleteUserControllerSut() (*controllers.DeleteUserController, *DeleteUserServiceMock) {
	serviceMock := &DeleteUserServiceMock{Output: "user deleted successfully", Error: nil}
	sut := controllers.NewDeleteUserController(serviceMock)
	return sut, serviceMock
}

func TestRemoveUserController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceMock := MakeDeleteUserControllerSut()

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Output, httpResponse.Body)
}

func TestRemoveUserController_ShouldReturnBadRequestIfServiceReturnsError(t *testing.T) {
	sut, serviceMock := MakeDeleteUserControllerSut()
	serviceMock.Error = errors.New("service error")

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Error.Error(), httpResponse.Body)
}
