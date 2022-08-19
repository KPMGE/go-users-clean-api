package services

import (
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type GetBookByIdService struct {
	getBookRepo protocols.GetBookRepository
}

func (s *GetBookByIdService) GetBookById(bookId string) (*entities.Book, error) {
	book, err := s.getBookRepo.Get(bookId)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func NewGetBookByIdService(getBookRepo protocols.GetBookRepository) *GetBookByIdService {
	return &GetBookByIdService{
		getBookRepo: getBookRepo,
	}
}
