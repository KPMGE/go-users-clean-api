package mocks_test

import "github.com/KPMGE/go-users-clean-api/src/domain/entities"

type ListBooksRepositoryStub struct {
	Output      []*entities.Book
	OutputError error
}

func (repo *ListBooksRepositoryStub) List() ([]*entities.Book, error) {
	return repo.Output, repo.OutputError
}

func NewListBooksRepositoryStub() *ListBooksRepositoryStub {
	return &ListBooksRepositoryStub{}
}
