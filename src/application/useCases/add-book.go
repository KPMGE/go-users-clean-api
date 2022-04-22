package usecases

import (
	"errors"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type AddBookUseCase struct {
	bookRepo protocols.AddBookRepository
	userRepo protocols.UserRepository
}

func NewAddBookUseCase(bookRepo protocols.AddBookRepository, userRepo protocols.UserRepository) *AddBookUseCase {
	return &AddBookUseCase{
		bookRepo: bookRepo,
		userRepo: userRepo,
	}
}

func (useCase *AddBookUseCase) Add(input *dto.AddBookUseCaseInputDTO) (*dto.AddBookUseCaseOutputDTO, error) {
	foundUser, err := useCase.userRepo.GetById(input.UserId)

	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("User not found!")
	}

	newBook, err := entities.NewBook(input.Title, input.Author, input.Description, input.Price, foundUser)
	if err != nil {
		return nil, err
	}

	_, err = useCase.bookRepo.Add(newBook)
	if err != nil {
		return nil, err
	}

	outputDto := dto.AddBookUseCaseOutputDTO{
		ID:          newBook.ID,
		Title:       newBook.Title,
		Author:      newBook.Author,
		Price:       newBook.Price,
		Description: newBook.Description,
		User:        foundUser,
	}

	return &outputDto, nil
}
