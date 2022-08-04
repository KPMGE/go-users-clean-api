package usecases_test

import (
	"errors"
	"log"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeRemoveBookSut() (
	*usecases.RemoveBookUseCase,
	*mocks_test.RemoveBookRepositorySpy,
	*mocks_test.FindBookRepositorySpy,
	*mocks_test.UserRepositorySpy,
) {
	removeBookRepo := mocks_test.NewRemoveBookRepositorySpy()
	removeBookRepo.RemoveError = nil

	findBookRepo := mocks_test.NewFindBookRepositorySpy()
	findBookRepo.FindError = nil
	fakeBook, err := entities.NewBook("any_title", "any_author", "any_description", 100.2, "any_user_id")
	if err != nil {
		log.Fatal(err)
	}
	findBookRepo.FindOutput = fakeBook

	userRepo := mocks_test.NewUserRepositorySpy()
	fakeUser, err := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fakeUser.ID = "any_user_id"
	fakeBook.ID = "any_valid_book_id"
	fakeUser.Books = append(fakeUser.Books, *fakeBook)
	userRepo.GetByidOutput = fakeUser

	sut := usecases.NewRemoveBookUseCase(removeBookRepo, findBookRepo, userRepo)
	return sut, removeBookRepo, findBookRepo, userRepo
}

func TestRemoveBookUseCase_ShouldCallRepositoryWithRightBookId(t *testing.T) {
	sut, removeBookRepo, _, _ := MakeRemoveBookSut()

	sut.Remove("any_valid_book_id")

	require.Equal(t, "any_valid_book_id", removeBookRepo.RemoveInput)
}

func TestRemoveBookUseCase_ShouldReturnErrorIfRepositoryReturnsError(t *testing.T) {
	sut, removeBookRepo, _, _ := MakeRemoveBookSut()
	removeBookRepo.RemoveError = errors.New("repo error")

	deletedBook, err := sut.Remove("any_invalid_id")

	require.Nil(t, deletedBook)
	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
}

func TestRemoveBookUseCase_ShouldCallFindBookRepositoryWithRightBookId(t *testing.T) {
	sut, _, findBookRepo, _ := MakeRemoveBookSut()

	sut.Remove("any_book_id")

	require.Equal(t, "any_book_id", findBookRepo.FindInput)
}

func TestRemoveBookUseCase_ShouldReturnErrorIfFindBookReturnsNil(t *testing.T) {
	sut, _, findBookRepo, _ := MakeRemoveBookSut()
	findBookRepo.FindOutput = nil

	deletedBook, err := sut.Remove("any_book_id")

	require.Nil(t, deletedBook)
	require.Error(t, err)
	require.Equal(t, "book not found!", err.Error())
}

func TestRemoveBookUseCase_ShouldReturnErrorIfFindBookReturnsError(t *testing.T) {
	sut, _, findBookRepo, _ := MakeRemoveBookSut()
	findBookRepo.FindError = errors.New("repo error")

	deletedBook, err := sut.Remove("any_book_id")

	require.Nil(t, deletedBook)
	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
}

func TestRemoveBookUseCase_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, _, findBookRepo, _ := MakeRemoveBookSut()

	deletedBook, err := sut.Remove("any_valid_book_id")

	require.Nil(t, err)
	require.Equal(t, findBookRepo.FindOutput.Author, deletedBook.Author)
	require.Equal(t, findBookRepo.FindOutput.Price, deletedBook.Price)
	require.Equal(t, findBookRepo.FindOutput.Description, deletedBook.Description)
	require.Equal(t, findBookRepo.FindOutput.Title, deletedBook.Title)
	require.Equal(t, findBookRepo.FindOutput.UserID, deletedBook.UserId)
}

func TestRemoveBookUseCase_ShouldCallFindUserWithCorrectUserId(t *testing.T) {
	sut, _, _, userRepo := MakeRemoveBookSut()

	sut.Remove("any_valid_book_id")

	require.Equal(t, "any_user_id", userRepo.GetByidInput)
}
