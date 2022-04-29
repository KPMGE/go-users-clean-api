package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type ListBooksRepository interface {
	List() []*entities.Book
}

type ListBooksRepositoryStub struct {
	output []*entities.Book
}

func (repo *ListBooksRepositoryStub) List() []*entities.Book {
	return repo.output
}

func NewListBooksRepositoryStub() *ListBooksRepositoryStub {
	return &ListBooksRepositoryStub{}
}

type ListBooksUseCase struct {
	listBooksRepo ListBooksRepository
}

func NewListBookUseCase(repo ListBooksRepository) *ListBooksUseCase {
	return &ListBooksUseCase{
		listBooksRepo: repo,
	}
}

func (useCase *ListBooksUseCase) List() []*entities.Book {
	return useCase.listBooksRepo.List()
}

func MakeListBooksSut() (*ListBooksUseCase, *ListBooksRepositoryStub) {
	repo := NewListBooksRepositoryStub()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 100.5, "any_user_id")
	repo.output = append(repo.output, fakeBook)
	sut := NewListBookUseCase(repo)
	return sut, repo
}

func TestListBooksUseCase_ShoulReturnRightDataFromRepository(t *testing.T) {
	sut, _ := MakeListBooksSut()

	books := sut.List()

	require.Equal(t, 1, len(books))
	require.Equal(t, "any_title", books[0].Title)
	require.Equal(t, "any_author", books[0].Author)
	require.Equal(t, "any_description", books[0].Description)
	require.Equal(t, "any_user_id", books[0].UserId)
	require.Equal(t, 100.5, books[0].Price)
}
