package controllers_test

import (
	"encoding/json"
	"errors"
	"log"
	"testing"

	dto "github.com/KPMGE/go-users-clean-api/src/application/DTO"
	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
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

	output, err := controller.useCase.Add(&inputDto)
	if err != nil {
		return helpers.BadRequest(err)
	}

	outputJson, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	return helpers.Ok(outputJson)
}

func MakeAddBookControllerSut() (*AddBookController, *mocks_test.AddBookRepositorySpy) {
	bookRepo := mocks_test.NewAddBookRepositorySpy()
	userRepo := mocks_test.NewUserRepositorySpy()
	fakeUser, _ := entities.NewUser("any_name", "any_username", "any_email@gmail.com")
	userRepo.GetByidOutput = fakeUser
	useCase := usecases.NewAddBookUseCase(bookRepo, userRepo)
	sut := NewAddBookController(useCase)
	return sut, bookRepo
}

func TestAddBookController_ShouldCallUseCaseWithRightData(t *testing.T) {
	sut, bookRepo := MakeAddBookControllerSut()

	sut.Handle(FAKE_REQUEST)

	require.Equal(t, "any_author", bookRepo.Input.Author)
	require.Equal(t, "any_title", bookRepo.Input.Title)
	require.Equal(t, 123.3, bookRepo.Input.Price)
	require.Equal(t, "any_description", bookRepo.Input.Description)
}

func TestAddBookController_ShouldReturnErrorIfUseCaseReturnsError(t *testing.T) {
	sut, bookRepo := MakeAddBookControllerSut()
	bookRepo.OutputError = errors.New("book repo error")

	httpResponse := sut.Handle(FAKE_REQUEST)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, "book repo error", string(httpResponse.JsonBody))
}

func TestAddBookController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	sut, _ := MakeAddBookControllerSut()

	httpResponse := sut.Handle(FAKE_REQUEST)

	var outputDto dto.AddBookUseCaseOutputDTO
	err := json.Unmarshal(httpResponse.JsonBody, &outputDto)
	if err != nil {
		log.Fatal(err)
	}

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, "any_author", outputDto.Author)
	require.Equal(t, "any_description", outputDto.Description)
	require.Equal(t, "any_title", outputDto.Title)
	require.Equal(t, 123.3, outputDto.Price)
	require.Equal(t, "any_name", outputDto.User.Name)
	require.Equal(t, "any_username", outputDto.User.UserName)
	require.NotNil(t, outputDto.ID)
	require.NotNil(t, outputDto.User.ID)
}
