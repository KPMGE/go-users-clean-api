package usecases

import (
	"errors"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
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

func (useCase *RemoveBookUseCase) Remove(bookId string) (*dto.RemoveBookUseCaseOutputDTO, error) {
	foundBook, err := useCase.findBookRepo.Find(bookId)
	if err != nil {
		return nil, err
	}
	if foundBook == nil {
		return nil, errors.New("book not found!")
	}

	useCase.userRepo.GetById(foundBook.UserId)

	err = useCase.removeBookRepo.Remove(bookId)
	if err != nil {
		return nil, err
	}

	outputDto := dto.RemoveBookUseCaseOutputDTO{
		Title:       foundBook.Title,
		Author:      foundBook.Author,
		Description: foundBook.Description,
		Price:       foundBook.Price,
		UserId:      foundBook.UserId,
	}

	return &outputDto, nil
}
