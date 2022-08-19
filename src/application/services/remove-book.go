package services

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type RemoveBookService struct {
	removeBookRepo protocols.RemoveBookRepository
	findBookRepo   protocols.FindBookRepository
	userRepo       protocols.UserRepository
}

func NewRemoveBookService(
	removeBookRepo protocols.RemoveBookRepository,
	findBookRepo protocols.FindBookRepository,
	userRepo protocols.UserRepository,
) *RemoveBookService {
	return &RemoveBookService{
		userRepo:       userRepo,
		removeBookRepo: removeBookRepo,
		findBookRepo:   findBookRepo,
	}
}

func removeIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func getBookIndex(books []entities.Book, bookId string) int {
	for i, book := range books {
		if book.ID == bookId {
			return i
		}
	}
	return -1
}

func (s *RemoveBookService) RemoveBook(bookId string) (*domaindto.RemoveBookUseCaseOutputDTO, error) {
	foundBook, err := s.findBookRepo.Find(bookId)
	if err != nil {
		return nil, err
	}
	if foundBook == nil {
		return nil, errors.New("book not found!")
	}

	err = s.removeBookRepo.Remove(bookId)
	if err != nil {
		return nil, err
	}

	outputDto := domaindto.RemoveBookUseCaseOutputDTO{
		Title:       foundBook.Title,
		Author:      foundBook.Author,
		Description: foundBook.Description,
		Price:       foundBook.Price,
		UserId:      foundBook.UserID,
	}

	return &outputDto, nil
}
