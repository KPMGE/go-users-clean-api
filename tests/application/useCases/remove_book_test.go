package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type RemoveBookRepository interface {
	Remove(bookId string) error
}

type RemoveBookRepositorySpy struct {
	RemoveInput string
	RemoveError error
}

func (repo *RemoveBookRepositorySpy) Remove(bookId string) error {
	repo.RemoveInput = bookId
	return repo.RemoveError
}

func NewRemoveBookRepositorySpy() *RemoveBookRepositorySpy {
	return &RemoveBookRepositorySpy{}
}

type RemoveBookUseCase struct {
	bookRepo RemoveBookRepository
}

func NewRemoveBookUseCase(repo RemoveBookRepository) *RemoveBookUseCase {
	return &RemoveBookUseCase{
		bookRepo: repo,
	}
}

func (useCase *RemoveBookUseCase) Remove(bookId string) error {
	err := useCase.bookRepo.Remove(bookId)
	if err != nil {
		return err
	}
	return nil
}

func MakeRemoveBookSut() (*RemoveBookUseCase, *RemoveBookRepositorySpy) {
	bookRepo := NewRemoveBookRepositorySpy()
	bookRepo.RemoveError = nil
	sut := NewRemoveBookUseCase(bookRepo)
	return sut, bookRepo
}

func TestRemoveBookUseCase_ShouldCallRepositoryWithRightBookId(t *testing.T) {
	sut, bookRepo := MakeRemoveBookSut()

	sut.Remove("any_valid_book_id")

	require.Equal(t, "any_valid_book_id", bookRepo.RemoveInput)
}

func TestRemoveBookUseCase_ShouldReturnErrorIfRepositoryReturnsError(t *testing.T) {
	sut, bookRepo := MakeRemoveBookSut()
	bookRepo.RemoveError = errors.New("repo error")

	err := sut.Remove("any_invalid_id")

	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
}
