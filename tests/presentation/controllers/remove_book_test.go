package controllers_test

import (
	"encoding/json"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/controllers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeRemoveBookControllerSut() (*controllers.RemoveBookController, *mocks_test.FindBookRepositorySpy, *mocks_test.RemoveBookRepositorySpy) {
	removeBookRepo := mocks_test.NewRemoveBookRepositorySpy()

	findBookRepo := mocks_test.NewFindBookRepositorySpy()
	fakeBook, _ := entities.NewBook("any_title", "any_author", "any_description", 100, "any_user_id")
	findBookRepo.FindOutput = fakeBook

	userRepo := mocks_test.NewUserRepositorySpy()
	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_valid_email@gmail.com")
	fakeUser.ID = "any_user_id"
	fakeBook.ID = "any_book_id"
	fakeUser.Books = append(fakeUser.Books, *fakeBook)
	userRepo.GetByidOutput = fakeUser

	useCase := usecases.NewRemoveBookUseCase(removeBookRepo, findBookRepo, userRepo)
	sut := controllers.NewRemoveBookController(useCase)

	return sut, findBookRepo, removeBookRepo
}

func TestRemoveBookController_ShoulCallUseCaseWithRightBookId(t *testing.T) {
	sut, findBookRepo, removeBookRepo := MakeRemoveBookControllerSut()

	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	sut.Handle(fakeRequest)

	require.Equal(t, "any_book_id", findBookRepo.FindInput)
	require.Equal(t, "any_book_id", removeBookRepo.RemoveInput)
}

func TestRemoveBookController_ShoulReturnRightDataOnSuccess(t *testing.T) {
	sut, _, _ := MakeRemoveBookControllerSut()

	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	httpResponse := sut.Handle(fakeRequest)

	var removedBook dto.RemoveBookUseCaseOutputDTO
	json.Unmarshal(httpResponse.JsonBody, &removedBook)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, "any_title", removedBook.Title)
	require.Equal(t, "any_author", removedBook.Author)
	require.Equal(t, "any_description", removedBook.Description)
	require.Equal(t, 100.0, removedBook.Price)
	require.Equal(t, "any_user_id", removedBook.UserId)
}

func TestRemoveBookController_ShoulReturnErrorIfUseCaseRetunsError(t *testing.T) {
	sut, findBookRepo, _ := MakeRemoveBookControllerSut()
	findBookRepo.FindOutput = nil

	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	httpResponse := sut.Handle(fakeRequest)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, "book not found!", string(httpResponse.JsonBody))
}
