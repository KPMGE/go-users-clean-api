package usecases_test

import (
	"errors"
	"log"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type RemoveBookRepository interface {
	Remove(bookId string) error
}

type FindBookRepository interface {
	Find(bookId string) (*entities.Book, error)
}

type RemoveBookRepositorySpy struct {
	RemoveInput string
	RemoveError error
}

type FindBookRepositorySpy struct {
	FindInput  string
	FindOutput *entities.Book
	FindError  error
}

func (repo *RemoveBookRepositorySpy) Remove(bookId string) error {
	repo.RemoveInput = bookId
	return repo.RemoveError
}

func NewRemoveBookRepositorySpy() *RemoveBookRepositorySpy {
	return &RemoveBookRepositorySpy{}
}

func (repo *FindBookRepositorySpy) Find(bookId string) (*entities.Book, error) {
	repo.FindInput = bookId
	return repo.FindOutput, repo.FindError
}

func NewFindBookRepositorySpy() *FindBookRepositorySpy {
	return &FindBookRepositorySpy{}
}

type RemoveBookUseCase struct {
	removeBookRepo RemoveBookRepository
	findBookRepo   FindBookRepository
}

func NewRemoveBookUseCase(removeBookRepo RemoveBookRepository, findBookRepo FindBookRepository) *RemoveBookUseCase {
	return &RemoveBookUseCase{
		removeBookRepo: removeBookRepo,
		findBookRepo:   findBookRepo,
	}
}

func (useCase *RemoveBookUseCase) Remove(bookId string) error {
	foundBook, _ := useCase.findBookRepo.Find(bookId)
	if foundBook == nil {
		return errors.New("book not found!")
	}

	err := useCase.removeBookRepo.Remove(bookId)
	if err != nil {
		return err
	}
	return nil
}

func MakeRemoveBookSut() (*RemoveBookUseCase, *RemoveBookRepositorySpy, *FindBookRepositorySpy) {
	removeBookRepo := NewRemoveBookRepositorySpy()
	removeBookRepo.RemoveError = nil

	findBookRepo := NewFindBookRepositorySpy()
	findBookRepo.FindError = nil
	fakeBook, err := entities.NewBook("any_title", "any_author", "any_description", 100.2, "any_user_id")
	if err != nil {
		log.Fatal(err)
	}
	findBookRepo.FindOutput = fakeBook

	sut := NewRemoveBookUseCase(removeBookRepo, findBookRepo)
	return sut, removeBookRepo, findBookRepo
}

func TestRemoveBookUseCase_ShouldCallRepositoryWithRightBookId(t *testing.T) {
	sut, removeBookRepo, _ := MakeRemoveBookSut()

	sut.Remove("any_valid_book_id")

	require.Equal(t, "any_valid_book_id", removeBookRepo.RemoveInput)
}

func TestRemoveBookUseCase_ShouldReturnErrorIfRepositoryReturnsError(t *testing.T) {
	sut, removeBookRepo, _ := MakeRemoveBookSut()
	removeBookRepo.RemoveError = errors.New("repo error")

	err := sut.Remove("any_invalid_id")

	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
}

func TestRemoveBookUseCase_ShouldCallFindBookRepositoryWithRightBookId(t *testing.T) {
	sut, _, findBookRepo := MakeRemoveBookSut()

	sut.Remove("any_book_id")

	require.Equal(t, "any_book_id", findBookRepo.FindInput)
}

func TestRemoveBookUseCase_ShouldReturnErrorIfFindBookReturnsNil(t *testing.T) {
	sut, _, findBookRepo := MakeRemoveBookSut()
	findBookRepo.FindOutput = nil

	err := sut.Remove("any_book_id")

	require.Error(t, err)
	require.Equal(t, "book not found!", err.Error())
}
