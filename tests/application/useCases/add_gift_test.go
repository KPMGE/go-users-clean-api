package usecases_test

import (
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
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
	input *entities.Book
}

func (repo *AddBookRepositorySpy) Add(newBook *entities.Book) (*entities.Book, error) {
	repo.input = newBook
	return nil, nil
}

func NewBookRepositorySpy() *AddBookRepositorySpy {
	return &AddBookRepositorySpy{}
}

type AddBookUseCase struct {
	repo AddBookRepository
}

func NewAddBookUseCase(repo AddBookRepository) *AddBookUseCase {
	return &AddBookUseCase{
		repo: repo,
	}
}

func (useCase *AddBookUseCase) Add(input *AddBookUseCaseInputDTO) (*AddBookUseCaseOutputDTO, error) {
	user, _ := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	newBook, _ := entities.NewBook(input.Title, input.Author, input.Description, input.Price, user)
	useCase.repo.Add(newBook)
	return nil, nil
}

func TestAddBookUseCase_ShouldCallRepositoryWithRightData(t *testing.T) {
	fakeInput := NewAddBookUseCaseInputDTO("any_title", "any_author", 342.2, "any_description", "any_valid_user_id")
	repo := NewBookRepositorySpy()
	sut := NewAddBookUseCase(repo)

	sut.Add(fakeInput)

	require.Equal(t, fakeInput.Author, repo.input.Author)
	require.Equal(t, fakeInput.Description, repo.input.Description)
	require.Equal(t, fakeInput.Price, repo.input.Price)
	require.Equal(t, fakeInput.Title, repo.input.Title)
	require.NotNil(t, repo.input.ID)
	require.NotNil(t, repo.input.User)
}
