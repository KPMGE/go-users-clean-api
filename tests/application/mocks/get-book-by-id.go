package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type GetBookByIdRepositorySpy struct {
	Input       string
	Output      *entities.Book
	OutputError error
}

func (repo *GetBookByIdRepositorySpy) Get(bookId string) (*entities.Book, error) {
	repo.Input = bookId
	return repo.Output, repo.OutputError
}

func NewGetBookByIdRepositorySpy() *GetBookByIdRepositorySpy {
	return &GetBookByIdRepositorySpy{}
}
