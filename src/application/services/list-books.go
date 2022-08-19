package services

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type ListBooksService struct {
	listBooksRepo protocols.ListBooksRepository
}

func NewListBookService(repo protocols.ListBooksRepository) *ListBooksService {
	return &ListBooksService{
		listBooksRepo: repo,
	}
}

func (s *ListBooksService) ListBooks() ([]*entities.Book, error) {
	books, err := s.listBooksRepo.List()
	if err != nil {
		return nil, err
	}
	return books, nil
}
