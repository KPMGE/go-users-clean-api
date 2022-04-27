package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type GetBookByIdRepositoryMock struct {
	input string
}

func (repo *GetBookByIdRepositoryMock) Get(bookId string) *entities.Book {
	repo.input = bookId
	return nil
}

func NewGetBookByIdRepositoryMock() *GetBookByIdRepositoryMock {
	return &GetBookByIdRepositoryMock{}
}

type GetBookRepository interface {
	Get(bookId string) *entities.Book
}

type GetBookByIdUseCase struct {
	getBookRepo GetBookRepository
}

func (useCase *GetBookByIdUseCase) GetById(bookId string) {
	useCase.getBookRepo.Get(bookId)
}

func NewGetBookByIdUseCase(getBookRepo GetBookRepository) *GetBookByIdUseCase {
	return &GetBookByIdUseCase{
		getBookRepo: getBookRepo,
	}
}

func MakeGetBookByIdSut() (*GetBookByIdUseCase, *GetBookByIdRepositoryMock) {
	getBookRepo := NewGetBookByIdRepositoryMock()
	sut := NewGetBookByIdUseCase(getBookRepo)
	return sut, getBookRepo
}

func TestGetBookByIdUseCase_ShouldCallRepositoryWithRightData(t *testing.T) {
	sut, getBookRepo := MakeGetBookByIdSut()
	sut.GetById("any_book_id")

	require.Equal(t, "any_book_id", getBookRepo.input)
}
