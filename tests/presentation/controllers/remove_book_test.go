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

type RemoveBookServiceMock struct {
	Output *domaindto.RemoveBookUseCaseOutputDTO
	Error  error
}

func (s *RemoveBookServiceMock) RemoveBook(bookId string) (*domaindto.RemoveBookUseCaseOutputDTO, error) {
	return s.Output, s.Error
}

func MakeRemoveBookControllerSut() (*controllers.RemoveBookController, *RemoveBookServiceMock) {
	serviceMock := &RemoveBookServiceMock{}
	sut := controllers.NewRemoveBookController(serviceMock)

	return sut, serviceMock
}

func TestRemoveBookController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceMock := MakeRemoveBookControllerSut()

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Output, httpResponse.Body)
}

func TestRemoveBookController_ShouldReturnBadRequestIfServiceReturnsError(t *testing.T) {
	sut, serviceMock := MakeRemoveBookControllerSut()
	serviceMock.Error = errors.New("service error")

	httpResponse := sut.Handle(protocols.NewHttpRequest(nil, []byte("any_id")))

	require.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Error.Error(), httpResponse.Body)
}
