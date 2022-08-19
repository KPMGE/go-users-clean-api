package services

import (
	"errors"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	domaindto "github.com/KPMGE/go-users-clean-api/src/domain/domain-dto"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
)

type AddBookService struct {
	bookRepo protocols.AddBookRepository
	userRepo protocols.UserRepository
}

func NewAddBookService(bookRepo protocols.AddBookRepository, userRepo protocols.UserRepository) *AddBookService {
	return &AddBookService{
		bookRepo: bookRepo,
		userRepo: userRepo,
	}
}

func (useCase *AddBookService) AddBook(input *domaindto.AddBookUseCaseInputDTO) (*domaindto.AddBookUseCaseOutputDTO, error) {
	foundUser, err := useCase.userRepo.GetById(input.UserId)

	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("User not found!")
	}

	newBook, err := entities.NewBook(input.Title, input.Author, input.Description, input.Price, foundUser.ID)
	if err != nil {
		return nil, err
	}

	_, err = useCase.bookRepo.Add(newBook)
	if err != nil {
		return nil, err
	}

	foundUser.Books = append(foundUser.Books, *newBook)

	outputDto := domaindto.AddBookUseCaseOutputDTO{
		ID:          newBook.ID,
		Title:       newBook.Title,
		Author:      newBook.Author,
		Price:       newBook.Price,
		Description: newBook.Description,
	}

	return &outputDto, nil
}
