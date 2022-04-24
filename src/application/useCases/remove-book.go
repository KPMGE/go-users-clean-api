package usecases

import (
	"errors"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type RemoveBookUseCase struct {
	removeBookRepo protocols.RemoveBookRepository
	findBookRepo   protocols.FindBookRepository
	userRepo       protocols.UserRepository
}

func NewRemoveBookUseCase(
	removeBookRepo protocols.RemoveBookRepository,
	findBookRepo protocols.FindBookRepository,
	userRepo protocols.UserRepository,
) *RemoveBookUseCase {
	return &RemoveBookUseCase{
		userRepo:       userRepo,
		removeBookRepo: removeBookRepo,
		findBookRepo:   findBookRepo,
	}
}

func removeIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func getBookIndex(books []*entities.Book, bookId string) int {
	for i, book := range books {
		if book.ID == bookId {
			return i
		}
	}
	return -1
}

func (useCase *RemoveBookUseCase) Remove(bookId string) (*dto.RemoveBookUseCaseOutputDTO, error) {
	foundBook, err := useCase.findBookRepo.Find(bookId)
	if err != nil {
		return nil, err
	}
	if foundBook == nil {
		return nil, errors.New("book not found!")
	}

	err = useCase.removeBookRepo.Remove(bookId)
	if err != nil {
		return nil, err
	}

	foundUser, err := useCase.userRepo.GetById(foundBook.UserId)
	if err != nil {
		return nil, err
	}

	index := getBookIndex(foundUser.Books, bookId)
	if index == -1 {
		return nil, errors.New("Cannot find book in user!")
	}
	foundUser.Books = removeIndex(foundUser.Books, index)

	useCase.userRepo.Save(foundUser)

	outputDto := dto.RemoveBookUseCaseOutputDTO{
		Title:       foundBook.Title,
		Author:      foundBook.Author,
		Description: foundBook.Description,
		Price:       foundBook.Price,
		UserId:      foundBook.UserId,
	}

	return &outputDto, nil
}
