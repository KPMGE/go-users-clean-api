package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type AddBookRepositorySpy struct {
	Input       *entities.Book
	Output      *entities.Book
	OutputError error
}

func (repo *AddBookRepositorySpy) Add(newBook *entities.Book) (*entities.Book, error) {
	repo.Input = newBook
	repo.Output = repo.Input
	return repo.Output, repo.OutputError
}

func NewAddBookRepositorySpy() *AddBookRepositorySpy {
	return &AddBookRepositorySpy{}
}
