package controllers_test

import (
	"encoding/json"
	"errors"
	"testing"

	usecases "github.com/KPMGE/go-users-clean-api/src/application/useCases"
	"github.com/KPMGE/go-users-clean-api/src/presentation/helpers"
	"github.com/KPMGE/go-users-clean-api/src/presentation/protocols"
	mocks_test "github.com/KPMGE/go-users-clean-api/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

type GetBookByIdController struct {
	useCase *usecases.GetBookByIdUseCase
}

func NewGetBookByIdController(useCase *usecases.GetBookByIdUseCase) *GetBookByIdController {
	return &GetBookByIdController{
		useCase: useCase,
	}
}

func (controller *GetBookByIdController) Handle(request *protocols.HttpRequest) *protocols.HttpResponse {
	book, err := controller.useCase.GetById(string(request.Params))
	if err != nil {
		return helpers.ServerError(err)
	}

	jsonBook, err := json.Marshal(book)
	if err != nil {
		return helpers.ServerError(err)
	}

	return helpers.Ok(jsonBook)
}

func MakeGetBookByIdControllerSut() (*GetBookByIdController, *mocks_test.GetBookByIdRepositorySpy) {
	repo := mocks_test.NewGetBookByIdRepositorySpy()
	useCase := usecases.NewGetBookByIdUseCase(repo)
	sut := NewGetBookByIdController(useCase)
	return sut, repo
}

func TestGetGiftByIdController_ShouldCallUseCaseWithRightData(t *testing.T) {
	sut, repo := MakeGetBookByIdControllerSut()
	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	sut.Handle(fakeRequest)
	require.Equal(t, "any_book_id", repo.Input)
}

func TestGetGiftByIdController_ShouldReturnErrorIfUseCaseReturnsError(t *testing.T) {
	sut, repo := MakeGetBookByIdControllerSut()
	repo.OutputError = errors.New("repo error")
	fakeRequest := protocols.NewHtppRequest(nil, []byte("any_book_id"))
	httpResponse := sut.Handle(fakeRequest)
	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, "repo error", string(httpResponse.JsonBody))
}
