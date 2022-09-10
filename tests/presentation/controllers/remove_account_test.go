package controllers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

type RemoveAccountServiceMock struct {
	Output string
	Error  error
}

func (r *RemoveAccountServiceMock) RemoveAccount(accountId string) (string, error) {
	return r.Output, r.Error
}

func MakeSut() (*controllers.RemoveAccountController, *RemoveAccountServiceMock) {
	serviceMock := &RemoveAccountServiceMock{Output: "account deleted", Error: nil}
	sut := controllers.NewRemoveAccountController(serviceMock)
	return sut, serviceMock
}

func TestRemoveAccountController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceMock := MakeSut()

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Output, httpResponse.Body)
}

func TestRemoveAccountController_ShouldReturnBadRequestIfServiceReturnsError(t *testing.T) {
	sut, serviceMock := MakeSut()
	serviceMock.Error = errors.New("there is not account with this id")

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Error.Error(), httpResponse.Body)
}
