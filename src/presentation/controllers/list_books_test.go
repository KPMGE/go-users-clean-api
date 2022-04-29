package controllers_test

import (
	"encoding/json"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type ListBooksController struct {
	useCase *usecases.ListBooksUseCase
}

func NewListBooksController(useCase *usecases.ListBooksUseCase) *ListBooksController {
	return &ListBooksController{
		useCase: useCase,
	}
}

func (controller *ListBooksController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	books, err := controller.useCase.List()
	if err != nil {
		return helpers.ServerError(err)
	}

	jsonBooks, err := json.Marshal(books)
	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonBooks)
}

func TestListBooksController_shouldReturnRightDataOnSuccess(t *testing.T) {
	repo := mocks_test.NewListBooksRepositoryStub()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 200.2, "any_user_id")
	repo.Output = append(repo.Output, fakeBook)
	repo.OutputError = nil
	useCase := usecases.NewListBookUseCase(repo)
	sut := NewListBooksController(useCase)

	fakeRequest := protocols.NewHtppRequest(nil, nil)
	httpResponse := sut.Handle(fakeRequest)

	var books []*entities.Book
	json.Unmarshal(httpResponse.JsonBody, &books)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, "any_title", books[0].Title)
	require.Equal(t, "any_description", books[0].Description)
	require.Equal(t, "any_user_id", books[0].UserId)
	require.Equal(t, "any_author", books[0].Author)
	require.Equal(t, 200.2, books[0].Price)
}
