package controllers_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListBoooksControllerSut() (*controllers.ListBooksController, *mocks_test.ListBooksRepositoryStub) {
	repo := mocks_test.NewListBooksRepositoryStub()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 200.2, "any_user_id")
	repo.Output = append(repo.Output, fakeBook)
	repo.OutputError = nil
	service := services.NewListBookService(repo)
	sut := controllers.NewListBooksController(service)
	return sut, repo
}

func TestListBooksController_shouldReturnRightDataOnSuccess(t *testing.T) {
	sut, serviceStub := MakeListBoooksControllerSut()

	fakeRequest := protocols.NewHttpRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, serviceStub.Output, httpResponse.Body)
}

func TestListBooksController_shouldServerErrorIfRepositoryReturnsError(t *testing.T) {
	sut, repo := MakeListBoooksControllerSut()
	repo.OutputError = errors.New("repo error")

	fakeRequest := protocols.NewHttpRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, repo.OutputError.Error(), httpResponse.Body)
}
