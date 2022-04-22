// TODO: Test outputs for use case
package usecases_test

import (
	"errors"
	"log"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/protocols"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type AddBookUseCaseInputDTO struct {
	Title       string
	Author      string
	Price       float64
	Description string
	UserId      string
}

type AddBookUseCaseOutputDTO struct {
	ID          string
	Title       string
	Author      string
	Price       float64
	Description string
	User        *entities.User
}

func NewAddBookUseCaseInputDTO(title string, author string, price float64, description string, userId string) *AddBookUseCaseInputDTO {
	return &AddBookUseCaseInputDTO{
		Title:       title,
		Author:      author,
		Price:       price,
		Description: description,
		UserId:      userId,
	}
}

type AddBookRepository interface {
	Add(newBook *entities.Book) (*entities.Book, error)
}

type AddBookRepositorySpy struct {
	input       *entities.Book
	output      *entities.Book
	outputError error
}

func (repo *AddBookRepositorySpy) Add(newBook *entities.Book) (*entities.Book, error) {
	repo.input = newBook
	repo.output = repo.input
	return repo.output, repo.outputError
}

func NewBookRepositorySpy() *AddBookRepositorySpy {
	return &AddBookRepositorySpy{}
}

type AddBookUseCase struct {
	bookRepo AddBookRepository
	userRepo protocols.UserRepository
}

func NewAddBookUseCase(bookRepo AddBookRepository, userRepo protocols.UserRepository) *AddBookUseCase {
	return &AddBookUseCase{
		bookRepo: bookRepo,
		userRepo: userRepo,
	}
}

func (useCase *AddBookUseCase) Add(input *AddBookUseCaseInputDTO) (*AddBookUseCaseOutputDTO, error) {
	foundUser, err := useCase.userRepo.GetById(input.UserId)

	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("User not found!")
	}

	newBook, _ := entities.NewBook(input.Title, input.Author, input.Description, input.Price, foundUser)
	_, err = useCase.bookRepo.Add(newBook)

	if err != nil {
		return nil, err
	}

	outputDto := AddBookUseCaseOutputDTO{
		ID:          newBook.ID,
		Title:       newBook.Title,
		Author:      newBook.Author,
		Price:       newBook.Price,
		Description: newBook.Description,
		User:        foundUser,
	}

	return &outputDto, nil
}

func MakeAddBookSut() (*AddBookUseCase, *AddBookRepositorySpy, *mocks_test.UserRepositorySpy) {
	fakeUser, err := entities.NewUser("any_name", "any_username", "any_email@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := mocks_test.NewUserRepositorySpy()
	userRepo.GetByidOutput = fakeUser
	userRepo.GetByidError = nil

	bookRepo := NewBookRepositorySpy()
	bookRepo.outputError = nil
	sut := NewAddBookUseCase(bookRepo, userRepo)

	return sut, bookRepo, userRepo
}

var FAKE_ADD_BOOK_INPUT_DTO = NewAddBookUseCaseInputDTO("any_title", "any_author", 342.2, "any_description", "any_invalid_user_id")

func TestAddBookUseCase_ShouldCallRepositoryWithRightData(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookSut()
	fakeInput := NewAddBookUseCaseInputDTO("any_title", "any_author", 342.2, "any_description", "any_valid_user_id")

	sut.Add(fakeInput)

	require.Equal(t, fakeInput.Author, bookRepo.input.Author)
	require.Equal(t, fakeInput.Description, bookRepo.input.Description)
	require.Equal(t, fakeInput.Price, bookRepo.input.Price)
	require.Equal(t, fakeInput.Title, bookRepo.input.Title)
	require.NotNil(t, bookRepo.input.ID)
	require.NotNil(t, bookRepo.input.User)
}

func TestAddBookUseCase_ShouldCallUserRepositoryWithRightUserId(t *testing.T) {
	sut, _, userRepo := MakeAddBookSut()

	sut.Add(FAKE_ADD_BOOK_INPUT_DTO)

	require.Equal(t, FAKE_ADD_BOOK_INPUT_DTO.UserId, userRepo.GetByidInput)
}

func TestAddBookUseCase_ShouldReturnErrorIfWrongUserIdIsGiven(t *testing.T) {
	sut, _, userRepo := MakeAddBookSut()
	userRepo.GetByidOutput = nil

	output, err := sut.Add(FAKE_ADD_BOOK_INPUT_DTO)

	require.Nil(t, output)
	require.Error(t, err)
	require.Equal(t, "User not found!", err.Error())
}

func TestAddBookUseCase_ShouldReturnErrorUserRepositoryReturnsError(t *testing.T) {
	sut, _, userRepo := MakeAddBookSut()
	userRepo.GetByidError = errors.New("repo error")

	output, err := sut.Add(FAKE_ADD_BOOK_INPUT_DTO)

	require.Nil(t, output)
	require.Error(t, err)
	require.Equal(t, "repo error", err.Error())
}

func TestAddBookUseCase_ShouldReturnOuputDTO(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookSut()

	output, err := sut.Add(FAKE_ADD_BOOK_INPUT_DTO)

	require.Nil(t, err)
	require.Equal(t, output.Title, bookRepo.output.Title)
	require.Equal(t, output.Author, bookRepo.output.Author)
	require.Equal(t, output.Price, bookRepo.output.Price)
	require.Equal(t, output.Description, bookRepo.output.Description)
	require.NotNil(t, output.ID)
	require.NotNil(t, output.User)
}

func TestAddBookUseCase_ShouldReturnErrorIfAddBookReturnsError(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookSut()
	bookRepo.outputError = errors.New("add book error")

	output, err := sut.Add(FAKE_ADD_BOOK_INPUT_DTO)

	require.Error(t, err)
	require.Nil(t, output)
	require.Equal(t, "add book error", err.Error())
}
