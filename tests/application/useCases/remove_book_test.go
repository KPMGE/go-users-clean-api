package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type RemoveBookRepository interface {
	Remove(bookId string) error
}

type RemoveBookRepositorySpy struct {
	RemoveInput string
}

func (repo *RemoveBookRepositorySpy) Remove(bookId string) error {
	repo.RemoveInput = bookId
	return nil
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

func (useCase *RemoveBookUseCase) Remove(bookId string) {
	useCase.bookRepo.Remove(bookId)
}

func TestRemoveBookUseCase_ShouldCallRepositoryWithRightBookId(t *testing.T) {
	bookRepo := NewRemoveBookRepositorySpy()
	sut := NewRemoveBookUseCase(bookRepo)

	sut.Remove("any_valid_book_id")

	require.Equal(t, "any_valid_book_id", bookRepo.RemoveInput)
}
