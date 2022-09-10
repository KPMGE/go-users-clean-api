package controllers_test

import (
	"errors"
	"testing"

	"github.com/KPMGE/go-users-clean-api/src/application/services"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	controllermocks_test "github.com/KPMGE/go-users-clean-api/tests/presentation/controller-mocks"
	"github.com/stretchr/testify/require"
)

var FAKE_ADD_BOOK_INPUT = `{
	"title": "any_title",
	"author": "any_author",
	"price": 123.3,
	"description": "any_description",
	"userId": "any_user_id"
}`

var FAKE_REQUEST = protocols.NewHttpRequest([]byte(FAKE_ADD_BOOK_INPUT), nil)

func MakeAddBookControllerSut() (*controllers.AddBookController, *mocks_test.AddBookRepositorySpy, *controllermocks_test.ValidatorMock) {
	bookRepo := mocks_test.NewAddBookRepositorySpy()
	userRepo := mocks_test.NewUserRepositorySpy()
	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_email@gmail.com")
	userRepo.GetByidOutput = fakeUser
	service := services.NewAddBookService(bookRepo, userRepo)
	validator := &controllermocks_test.ValidatorMock{Output: nil}
	sut := controllers.NewAddBookController(service, validator)
	return sut, bookRepo, validator
}

func TestAddBookController_ShouldCallUseCaseWithRightData(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookControllerSut()

	sut.Handle(FAKE_REQUEST)

	require.Equal(t, "any_author", bookRepo.Input.Author)
	require.Equal(t, "any_title", bookRepo.Input.Title)
	require.Equal(t, 123.3, bookRepo.Input.Price)
	require.Equal(t, "any_description", bookRepo.Input.Description)
}

func TestAddBookController_ShouldReturnErrorIfUseCaseReturnsError(t *testing.T) {
	sut, bookRepo, _ := MakeAddBookControllerSut()
	bookRepo.OutputError = errors.New("book repo error")

	httpResponse := sut.Handle(FAKE_REQUEST)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, bookRepo.OutputError.Error(), httpResponse.Body)
}

func TestAddBookController_ShouldReturnServerErrorIfValidatorReturnsError(t *testing.T) {
	sut, _, validatorMock := MakeAddBookControllerSut()
	validatorMock.Output = errors.New("validation error")

	httpResponse := sut.Handle(FAKE_REQUEST)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, validatorMock.Output.Error(), httpResponse.Body)
}
