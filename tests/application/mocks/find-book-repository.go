package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type FindBookRepositorySpy struct {
	FindInput  string
	FindOutput *entities.Book
	FindError  error
}

func (repo *FindBookRepositorySpy) Find(bookId string) (*entities.Book, error) {
	repo.FindInput = bookId
	return repo.FindOutput, repo.FindError
}

func NewFindBookRepositorySpy() *FindBookRepositorySpy {
	return &FindBookRepositorySpy{}
}
