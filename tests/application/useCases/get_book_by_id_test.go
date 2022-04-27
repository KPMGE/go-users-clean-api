package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeGetBookByIdSut() (*usecases.GetBookByIdUseCase, *mocks_test.GetBookByIdRepositorySpy) {
	getBookRepo := mocks_test.NewGetBookByIdRepositorySpy()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 100.23, "any_user_id")
	getBookRepo.Output = fakeBook
	sut := usecases.NewGetBookByIdUseCase(getBookRepo)
	return sut, getBookRepo
}

func TestGetBookByIdUseCase_ShouldCallRepositoryWithRightData(t *testing.T) {
	sut, getBookRepo := MakeGetBookByIdSut()
	sut.GetById("any_book_id")
	require.Equal(t, "any_book_id", getBookRepo.Input)
}

func TestGetBookByIdUseCase_ShouldRetunNilIfNoBookIsFound(t *testing.T) {
	sut, getBookRepo := MakeGetBookByIdSut()
	getBookRepo.Output = nil
	foundBook, _ := sut.GetById("any_book_id")
	require.Nil(t, foundBook)
}

func TestGetBookByIdUseCase_ShouldRetunRightBookOnSuccess(t *testing.T) {
	sut, _ := MakeGetBookByIdSut()
	foundBook, _ := sut.GetById("any_book_id")
	require.Equal(t, "any_title", foundBook.Title)
	require.Equal(t, "any_author", foundBook.Author)
	require.Equal(t, "any_user_id", foundBook.UserId)
	require.Equal(t, "any_description", foundBook.Description)
	require.Equal(t, 100.23, foundBook.Price)
	require.NotNil(t, foundBook.CreatedAt)
	require.NotNil(t, foundBook.UpdatedAt)
	require.NotNil(t, foundBook.ID)
}

func TestGetBookByIdUseCase_ShouldRetunErrorIfRepositoryReturnsError(t *testing.T) {
	sut, getBookRepo := MakeGetBookByIdSut()
	getBookRepo.OutputError = errors.New("repo error")
	foundBook, err := sut.GetById("any_book_id")
	require.Nil(t, foundBook)
	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
}
