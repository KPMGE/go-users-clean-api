package controllers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	"github.com/stretchr/testify/require"
)

type GetBookByIdServiceMock struct {
	Output *entities.Book
	Error  error
}

func (g *GetBookByIdServiceMock) GetBookById(bookId string) (*entities.Book, error) {
	return g.Output, g.Error
}

func MakeGetBookByIdControllerSut() (*controllers.GetBookByIdController, *GetBookByIdServiceMock) {
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 203.43, "any_user_id")
	serviceMock := &GetBookByIdServiceMock{Output: fakeBook, Error: nil}
	sut := controllers.NewGetBookByIdController(serviceMock)
	return sut, serviceMock
}

func TestGetBookByIdController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceMock := MakeGetBookByIdControllerSut()

	httpResponse := sut.Handle(protocols.NewHttpRequest([]byte("any_id"), nil))

	require.Equal(t, http.StatusOK, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Output, httpResponse.Body)
}

func TestGetBookByIdController_ShouldReturnNotFoundIfServiceReturnsNull(t *testing.T) {
	sut, serviceMock := MakeGetBookByIdControllerSut()
	serviceMock.Output = nil
	expectedError := errors.New("book not found")

	httpResponse := sut.Handle(protocols.NewHttpRequest([]byte("any_id"), nil))

	require.Equal(t, http.StatusNotFound, httpResponse.StatusCode)
	require.Equal(t, expectedError.Error(), httpResponse.Body)
}

func TestGetBookByIdController_ShouldReturnServerErrorIfServiceReturnsError(t *testing.T) {
	sut, serviceMock := MakeGetBookByIdControllerSut()
	serviceMock.Error = errors.New("service error")

	httpResponse := sut.Handle(protocols.NewHttpRequest([]byte("any_id"), nil))

	require.Equal(t, http.StatusInternalServerError, httpResponse.StatusCode)
	require.Equal(t, serviceMock.Error.Error(), httpResponse.Body)
}
