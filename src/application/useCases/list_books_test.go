package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListBooksSut() (*usecases.ListBooksUseCase, *mocks_test.ListBooksRepositoryStub) {
	repo := mocks_test.NewListBooksRepositoryStub()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 100.5, "any_user_id")
	repo.Output = append(repo.Output, fakeBook)
	repo.OutputError = nil
	sut := usecases.NewListBookUseCase(repo)
	return sut, repo
}

func TestListBooksUseCase_ShoulReturnRightDataFromRepository(t *testing.T) {
	sut, _ := MakeListBooksSut()

	books, _ := sut.List()

	require.Equal(t, 1, len(books))
	require.Equal(t, "any_title", books[0].Title)
	require.Equal(t, "any_author", books[0].Author)
	require.Equal(t, "any_description", books[0].Description)
	require.Equal(t, "any_user_id", books[0].UserId)
	require.Equal(t, 100.5, books[0].Price)
}

func TestListBooksUseCase_ShoulReturnErrorIfRepositoryReturnsError(t *testing.T) {
	sut, repo := MakeListBooksSut()
	repo.OutputError = errors.New("repo error")

	books, err := sut.List()

	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
	require.Nil(t, books)
}
