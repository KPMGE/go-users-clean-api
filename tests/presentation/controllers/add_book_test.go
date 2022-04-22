// TODO: solve problem in first test

package controllers_test

import (
	"encoding/json"
	"log"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

var FAKE_ADD_BOOK_INPUT = `{
	"title": "any_title",
	"author": "any_author",
	"price": 123.3,
	"description": "any_description",
	"userId": "any_user_id"
}`

var FAKE_REQUEST = protocols.NewHtppRequest([]byte(FAKE_ADD_BOOK_INPUT), nil)

type AddBookController struct {
	useCase *usecases.AddBookUseCase
}

func NewAddBookController(useCase *usecases.AddBookUseCase) *AddBookController {
	return &AddBookController{
		useCase: useCase,
	}
}

func (controller *AddBookController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	var inputDto dto.AddBookUseCaseInputDTO
	err := json.Unmarshal(request.Body, &inputDto)
	if err != nil {
		log.Fatal(err)
	}

	controller.useCase.Add(&inputDto)

	return nil
}

func TestAddBookController_ShouldCallUseCaseWithRightData(t *testing.T) {
	bookRepo := mocks_test.NewAddBookRepositorySpy()
	userRepo := mocks_test.NewUserRepositorySpy()
	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_email@gmail.com")
	userRepo.GetByidOutput = fakeUser
	useCase := usecases.NewAddBookUseCase(bookRepo, userRepo)
	sut := NewAddBookController(useCase)

	sut.Handle(FAKE_REQUEST)

	require.Equal(t, "any_author", bookRepo.Input.Author)
	require.Equal(t, "any_title", bookRepo.Input.Title)
	require.Equal(t, 123.3, bookRepo.Input.Price)
	require.Equal(t, "any_description", bookRepo.Input.Description)
}
