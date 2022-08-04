package usecases_test

import (
	"errors"
	"log"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

var FAKE_ADD_BOOK_INPUT_DTO = &dto.AddBookUseCaseInputDTO{
	Title:       "any_title",
	Author:      "any_author",
	Price:       342.2,
	Description: "any_description",
	UserId:      "any_valid_user_id",
}

func MakeAddBookSut() (*usecases.AddBookUseCase, *mocks_test.AddBookRepositorySpy, *mocks_test.UserRepositorySpy) {
	fakeUser, err := entities.NewUser("any_name", "any_username", "any_email@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := mocks_test.NewUserRepositorySpy()
	userRepo.GetByidOutput = fakeUser
	userRepo.GetByidError = nil

	bookRepo := mocks_test.NewAddBookRepositorySpy()
	bookRepo.OutputError = nil
	sut := usecases.NewAddBookUseCase(bookRepo, userRepo)

	return sut, bookRepo, userRepo
}

func TestAddBookUseCase_ShouldCallRepositoryWithRightData(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookSut()

	sut.Add(FAKE_ADD_BOOK_INPUT_DTO)
	require.Equal(t, FAKE_ADD_BOOK_INPUT_DTO.Description, bookRepo.Input.Description)
	require.Equal(t, FAKE_ADD_BOOK_INPUT_DTO.Price, bookRepo.Input.Price)
	require.Equal(t, FAKE_ADD_BOOK_INPUT_DTO.Title, bookRepo.Input.Title)
	require.NotNil(t, bookRepo.Input.ID)
	require.NotNil(t, bookRepo.Input.UserID)
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
	require.Equal(t, output.Title, bookRepo.Output.Title)
	require.Equal(t, output.Author, bookRepo.Output.Author)
	require.Equal(t, output.Price, bookRepo.Output.Price)
	require.Equal(t, output.Description, bookRepo.Output.Description)
	require.NotNil(t, output.ID)
}

func TestAddBookUseCase_ShouldReturnErrorIfAddBookReturnsError(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookSut()
	bookRepo.OutputError = errors.New("add book error")

	output, err := sut.Add(FAKE_ADD_BOOK_INPUT_DTO)

	require.Error(t, err)
	require.Nil(t, output)
	require.Equal(t, "add book error", err.Error())
}
