package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type GetBookByIdRepositorySpy struct {
	input  string
	output *entities.Book
}

func (repo *GetBookByIdRepositorySpy) Get(bookId string) *entities.Book {
	repo.input = bookId
	return repo.output
}

func NewGetBookByIdRepositorySpy() *GetBookByIdRepositorySpy {
	return &GetBookByIdRepositorySpy{}
}

type GetBookRepository interface {
	Get(bookId string) *entities.Book
}

type GetBookByIdUseCase struct {
	getBookRepo GetBookRepository
}

func (useCase *GetBookByIdUseCase) GetById(bookId string) *entities.Book {
	return useCase.getBookRepo.Get(bookId)
}

func NewGetBookByIdUseCase(getBookRepo GetBookRepository) *GetBookByIdUseCase {
	return &GetBookByIdUseCase{
		getBookRepo: getBookRepo,
	}
}

func MakeGetBookByIdSut() (*GetBookByIdUseCase, *GetBookByIdRepositorySpy) {
	getBookRepo := NewGetBookByIdRepositorySpy()
	sut := NewGetBookByIdUseCase(getBookRepo)
	return sut, getBookRepo
}

func TestGetBookByIdUseCase_ShouldCallRepositoryWithRightData(t *testing.T) {
	sut, getBookRepo := MakeGetBookByIdSut()
	sut.GetById("any_book_id")
	require.Equal(t, "any_book_id", getBookRepo.input)
}

func TestGetBookByIdUseCase_ShouldRetunNilIfNoBookIsFound(t *testing.T) {
	sut, getBookRepo := MakeGetBookByIdSut()
	getBookRepo.output = nil
	foundBook := sut.GetById("any_book_id")
	require.Nil(t, foundBook)
}
